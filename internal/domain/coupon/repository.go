package coupon

// Repository는 쿠폰 데이터를 저장하고 조회하는 인터페이스입니다.
type Repository interface {
	// Create는 새로운 쿠폰을 생성합니다.
	Create(coupon *Coupon) error
	// GetList는 캠페인 ID를 기반으로 쿠폰 코드 목록을 반환합니다.
	GetList(campaignID int) ([]string, error)
	// GetCount는 캠페인 ID를 기반으로 발급된 쿠폰 수를 반환합니다.
	GetCount(campaignID int) (int, error)
}
