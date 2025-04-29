package campaign

import (
	"context"
	"time"

	"github.com/SuperRPM/coupon-issuance-system/internal/domain/campaign"
	"github.com/SuperRPM/coupon-issuance-system/internal/domain/coupon"
)

// Service는 캠페인 관련 비즈니스 로직을 처리합니다.
type CampaignService struct {
	campaignRepo campaign.Repository
	couponRepo   coupon.Repository
}

// NewService는 새로운 캠페인 서비스를 생성합니다.
func NewService(campaignRepo campaign.Repository, couponRepo coupon.Repository) *CampaignService {
	return &CampaignService{
		campaignRepo: campaignRepo,
		couponRepo:   couponRepo,
	}
}

// CreateCampaign은 새로운 캠페인을 생성합니다.
func (s *CampaignService) CreateCampaign(ctx context.Context, name string, limit int, startDate time.Time, endDate time.Time) (*campaign.Campaign, error) {
	c := campaign.NewCampaign(name, limit, startDate, endDate)
	if err := s.campaignRepo.Create(c); err != nil {
		return nil, err
	}
	return c, nil
}

// GetCampaign은 ID로 캠페인을 조회합니다.
func (s *CampaignService) GetCampaign(ctx context.Context, id int) (*campaign.Campaign, error) {
	c, err := s.campaignRepo.Get(id)
	if err != nil || c == nil {
		return nil, err
	}

	// 쿠폰 레포지토리에서 발급된 쿠폰 수 가져오기
	couponCount, err := s.couponRepo.GetCount(id)
	if err != nil {
		return nil, err
	}
	c.IssuedCount = couponCount

	return c, nil
}
