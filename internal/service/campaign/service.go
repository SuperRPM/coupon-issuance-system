package campaign

import (
	"context"
	"errors"
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
	now := time.Now()
	maxStart := now.AddDate(1, 0, 0)
	if startDate.After(maxStart) {
		return nil, errors.New("캠페인 시작 날짜는 현재로부터 1년 이내여야 합니다")
	}
	if endDate.Before(startDate) {
		return nil, errors.New("캠페인 종료 날짜는 시작 날짜 이후여야 합니다")
	}
	maxEnd := startDate.AddDate(1, 0, 0)
	if endDate.After(maxEnd) {
		return nil, errors.New("캠페인 기간은 시작일로부터 최대 1년 이내여야 합니다")
	}
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
