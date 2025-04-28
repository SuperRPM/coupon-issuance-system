package coupon

// Repository는 쿠폰 데이터를 저장하고 조회하는 인터페이스입니다.
type Repository interface {
	// Create는 새로운 쿠폰을 생성합니다.
	Create(coupon *Coupon) error
}
