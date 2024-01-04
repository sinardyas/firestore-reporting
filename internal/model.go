package campaign

import "time"

type TrxDetail map[string]interface{}

type WinnerSubmission struct {
	Name       string `firestore:"name,omitempty"`
	Identity   int    `firestore:"no_ktp,omitempty"`
	Phone      string `firestore:"phone,omitempty"`
	Email      string `firestore:"email,omitempty"`
	City       string `firestore:"city,omitempty"`
	District   string `firestore:"district,omitempty"`
	Province   string `firestore:"province,omitempty"`
	PostalCode int    `firestore:"postal_code,omitempty"`
	Address    string `firestore:"shipment_address,omitempty"`
	Village    string `firestore:"village,omitempty"`
}

type WinnerBid struct {
	BiddingAt    time.Time `firestore:"bidding_at,omitempty"`
	BiddingPrice int       `firestore:"bidding_price,omitempty"`
	CampaignName string    `firestore:"campaign,omitempty"`
	Phone        string    `firestore:"phone,omitempty"`
	Submitted    bool      `firestore:"submitted,omitempty"`
	User         string    `firestore:"user,omitempty"`
	TrxDetail    TrxDetail `firestore:"transactionDetails,omitempty"`
	Submission   *WinnerSubmission
}

// type CampaignWinner struct {
// 	WinnerBid
// 	WinnerSubmission
// }
