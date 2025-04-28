package campaign

// Repository는 캠페인 데이터를 저장하고 조회하는 인터페이스입니다.
type Repository interface {
	// Create는 새로운 캠페인을 생성합니다.
	Create(campaign *Campaign) error
	// Get은 ID로 캠페인을 조회합니다.
	Get(id int) (*Campaign, error)
}
