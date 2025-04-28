package coupon

// Coupon은 쿠폰을 나타내는 엔티티입니다.
type Coupon struct {
	ID         string
	CampaignID int
	Code       string
	Used       bool
}

// NewCoupon은 새로운 쿠폰을 생성합니다.
func NewCoupon(campaignID int, code string) *Coupon {
	return &Coupon{
		CampaignID: campaignID,
		Code:       code,
		Used:       false,
	}
}

// Use는 쿠폰을 사용합니다.
func (c *Coupon) Use() {
	c.Used = true
}
