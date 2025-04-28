package coupon

// Repository는 쿠폰 데이터를 저장하고 조회하는 인터페이스입니다.
type Repository interface {
	// Create는 새로운 쿠폰을 생성합니다.
	Create(coupon *Coupon) error
	// Get은 ID로 쿠폰을 조회합니다.
	Get(id string) (*Coupon, error)
	// Update는 쿠폰을 업데이트합니다.
	Update(coupon *Coupon) error
}
