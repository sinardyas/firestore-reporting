package campaign

import (
	"context"
	"fmt"
	"reflect"
	"tikung/analytic/pkg/excel"
	firestore "tikung/analytic/pkg/firebase"

	"google.golang.org/api/iterator"
)

var _ CampaignRepo = (*campaignRepo)(nil)

type campaignRepo struct {
	app  firestore.FirestoreEngine
	excl excel.ExcelEngine
}

func NewCampaignRepo(fs firestore.FirestoreEngine, ex excel.ExcelEngine) CampaignRepo {
	return &campaignRepo{
		app:  fs,
		excl: ex,
	}
}

func (c *campaignRepo) GetAllWinners(cName CampaignName) ([]WinnerBid, error) {
	ctx := context.Background()
	db := c.app.GetDB()

	winnerColRef := db.Collection("campaign").Doc(string(cName)).Collection("winner_details")
	winnerDocItr := winnerColRef.Documents(ctx)

	var winners []WinnerBid

	for {
		winnerSnaps, err := winnerDocItr.Next()
		if err == iterator.Done {
			break
		}

		var winnerBid WinnerBid

		winnerSnaps.DataTo(&winnerBid)

		winners = append(winners, winnerBid)
	}

	return winners, nil
}

func (c *campaignRepo) GetWinnerSubmissions(cName CampaignName, winners []WinnerBid) ([]WinnerBid, error) {
	ctx := context.Background()
	db := c.app.GetDB()

	for i, winner := range winners {
		submissionDocItr := db.Collection("campaign").Doc(string(cName)).Collection("winner_details").Doc(string(winner.Phone)).Collection("submission_detail").Documents(ctx)

		for {
			submissionSnaps, err := submissionDocItr.Next()
			if err == iterator.Done {
				break
			}

			var winnerSubmission WinnerSubmission

			submissionSnaps.DataTo(&winnerSubmission)

			v := reflect.ValueOf(winnerSubmission)
			typeOfS := v.Type()

			fmt.Println("BiddingPrice: ", winner.BiddingPrice)
			fmt.Println("BiddingAt: ", winner.BiddingAt)
			fmt.Println("CampaignName: ", winner.CampaignName)

			if winner.TrxDetail != nil {
				fmt.Println("Paid: ", true)
				fmt.Println("PaidAt: ", winner.TrxDetail["transaction_time"])
			} else {
				fmt.Println("Paid: ", false)
			}

			for i := 0; i < v.NumField(); i++ {
				fmt.Printf("%s: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
			}

			fmt.Printf("\n")

			winners[i].Submission = &winnerSubmission
		}
	}

	return winners, nil
}

func (c *campaignRepo) ExportExcel(cName CampaignName, winners []WinnerBid) error {
	c.excl.SetHeader()
	f := c.excl.GetFile()

	startRow := 2
	for _, v := range winners {

		if v.Submission != nil {
			f.SetCellValue("Sheet1", "A"+fmt.Sprint(startRow), v.CampaignName)
			f.SetCellValue("Sheet1", "B"+fmt.Sprint(startRow), v.BiddingAt)
			f.SetCellValue("Sheet1", "C"+fmt.Sprint(startRow), v.BiddingPrice)
			f.SetCellValue("Sheet1", "D"+fmt.Sprint(startRow), v.Submission.Name)
			f.SetCellValue("Sheet1", "E"+fmt.Sprint(startRow), v.Submission.Email)
			f.SetCellValue("Sheet1", "F"+fmt.Sprint(startRow), v.Submission.Phone)
			f.SetCellStr("Sheet1", "G"+fmt.Sprint(startRow), fmt.Sprint(v.Submission.Identity))
			f.SetCellValue("Sheet1", "H"+fmt.Sprint(startRow), v.Submission.Address)

			if v.TrxDetail != nil {
				f.SetCellBool("Sheet1", "I"+fmt.Sprint(startRow), true)
			} else {
				f.SetCellBool("Sheet1", "I"+fmt.Sprint(startRow), false)
			}

			startRow += 1
		}
	}

	c.excl.SaveFile(string(cName))

	return nil
}
