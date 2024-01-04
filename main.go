package main

import (
	"log"
	"os"
	campaign "tikung/analytic/internal"
	"tikung/analytic/pkg/excel"
	firestore "tikung/analytic/pkg/firebase"
)

func main() {
	var campName campaign.CampaignName = campaign.CampaignName(os.Args[1])

	fsEng, err := firestore.NewFirestoreDB("tikung")
	if err != nil {
		log.Panicln("Error init firestore project")
	}

	defer fsEng.Close()

	exclEng := excel.NewExcel()

	defer exclEng.Close()

	campRepo := campaign.NewCampaignRepo(fsEng, exclEng)

	winners, err := campRepo.GetAllWinners(campName)
	if err != nil {
		log.Panicln("Error getting the winners")
	}

	w, err := campRepo.GetWinnerSubmissions(campName, winners)
	if err != nil {
		log.Panicln("Error getting the winners submissions")
	}

	campRepo.ExportExcel(campName, w)
}
