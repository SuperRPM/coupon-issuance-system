package campaign

import (
	"context"
	"time"

	"github.com/SuperRPM/coupon-issuance-system/internal/domain/campaign"
	"github.com/SuperRPM/coupon-issuance-system/internal/domain/coupon"
)

type CampaignService struct {
	campaignRepo campaign.Repository
	couponRepo   coupon.Repository
}

func NewService(campaignRepo campaign.Repository, couponRepo coupon.Repository) *CampaignService {
	return &CampaignService{
		campaignRepo: campaignRepo,
		couponRepo:   couponRepo,
	}
}

func (s *CampaignService) CreateCampaign(ctx context.Context, name string, limit int, startDate time.Time, endDate time.Time) (*campaign.Campaign, error) {
	c := campaign.NewCampaign(name, limit, startDate, endDate)
	if err := s.campaignRepo.Create(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CampaignService) GetCampaign(ctx context.Context, id int) (*campaign.Campaign, error) {
	c, err := s.campaignRepo.Get(id)
	if err != nil || c == nil {
		return nil, err
	}

	couponCount, err := s.couponRepo.GetCount(id)
	if err != nil {
		return nil, err
	}
	c.IssuedCount = couponCount

	return c, nil
}
