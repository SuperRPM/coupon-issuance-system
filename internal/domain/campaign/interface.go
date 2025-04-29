package campaign

import "context"

// 핸들러, 서비스, 리포지토리 의존성 관리

type Handler interface {
	CreateCampaign(ctx context.Context, campaign *Campaign) error
	GetCampaign(ctx context.Context, id int) (*Campaign, error)
}

type Service interface {
	CreateCampaign(ctx context.Context, campaign *Campaign) error
	GetCampaign(ctx context.Context, id int) (*Campaign, error)
}

type Repository interface {
	// Create는 새로운 캠페인을 생성합니다.
	Create(campaign *Campaign) error
	// Get은 ID로 캠페인을 조회합니다.
	Get(id int) (*Campaign, error)
}
