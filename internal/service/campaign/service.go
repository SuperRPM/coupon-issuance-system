package campaign

import (
	"context"
	"time"

	"github.com/SuperRPM/coupon-issuance-system/internal/domain/campaign"
)

// Service는 캠페인 관련 비즈니스 로직을 처리합니다.
type CampaignService struct {
	repo campaign.Repository
}

// NewService는 새로운 캠페인 서비스를 생성합니다.
func NewService(repo campaign.Repository) *CampaignService {
	return &CampaignService{
		repo: repo,
	}
}

// CreateCampaign은 새로운 캠페인을 생성합니다.
func (s *CampaignService) CreateCampaign(ctx context.Context, name string, limit int, startDate time.Time, endDate time.Time) (*campaign.Campaign, error) {
	c := campaign.NewCampaign(name, limit, startDate, endDate)
	if err := s.repo.Create(c); err != nil {
		return nil, err
	}
	return c, nil
}

// GetCampaign은 ID로 캠페인을 조회합니다.
func (s *CampaignService) GetCampaign(ctx context.Context, id int) (*campaign.Campaign, error) {
	return s.repo.Get(id)
}
