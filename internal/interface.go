package campaign

type CampaignName string

type CampaignRepo interface {
	GetAllWinners(CampaignName) ([]WinnerBid, error)
	GetWinnerSubmissions(CampaignName, []WinnerBid) ([]WinnerBid, error)
	ExportExcel(CampaignName, []WinnerBid) error
}
