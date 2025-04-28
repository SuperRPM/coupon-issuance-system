package campaign

// Campaign은 쿠폰 캠페인을 나타내는 엔티티입니다.
type Campaign struct {
	ID          int
	Name        string
	Limit       int
	IssuedCount int
}

// NewCampaign은 새로운 캠페인을 생성합니다.
func NewCampaign(name string, limit int) *Campaign {
	return &Campaign{
		Name:        name,
		Limit:       limit,
		IssuedCount: 0,
	}
}

// CanIssue는 쿠폰을 발급할 수 있는지 확인합니다.
func (c *Campaign) CanIssue() bool {
	return c.IssuedCount < c.Limit
}

// Issue는 쿠폰을 발급합니다.
func (c *Campaign) Issue() {
	c.IssuedCount++
}
