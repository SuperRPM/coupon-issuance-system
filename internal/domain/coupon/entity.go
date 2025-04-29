package coupon

type Coupon struct {
	ID         int
	CampaignID int
	Code       string
	Used       bool
}

func NewCoupon(campaignID int, code string) *Coupon {
	return &Coupon{
		CampaignID: campaignID,
		Code:       code,
	}
}
