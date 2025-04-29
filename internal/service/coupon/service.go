package coupon

import (
	"context"
	"errors"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/SuperRPM/coupon-issuance-system/internal/domain/coupon"
	"github.com/SuperRPM/coupon-issuance-system/internal/service/campaign"
)

type Service struct {
	repo            coupon.Repository
	campaignService *campaign.CampaignService
	usedCodes       map[string]bool
	mu              sync.RWMutex
}

func NewService(repo coupon.Repository, campaignService *campaign.CampaignService) *Service {
	return &Service{
		repo:            repo,
		campaignService: campaignService,
		usedCodes:       make(map[string]bool),
	}
}

func (s *Service) IssueCoupon(ctx context.Context, campaignID int) (*coupon.Coupon, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	campaign, err := s.campaignService.GetCampaign(ctx, campaignID)
	if err != nil {
		return nil, err
	}

	if campaign == nil {
		return nil, errors.New("campaign not found")
	}

	// Validation
	couponCount, err := s.repo.GetCount(campaignID)
	if err != nil {
		return nil, err
	}

	if campaign.Limit <= couponCount {
		return nil, errors.New("campaign limit exceeded")
	}

	if campaign.StartDate.After(time.Now()) {
		return nil, errors.New("campaign not started")
	}

	if campaign.EndDate.Before(time.Now()) {
		return nil, errors.New("campaign expired")
	}

	code := s.getHangulUniqueCode()

	c := coupon.NewCoupon(campaignID, code)
	if err := s.repo.Create(c); err != nil {
		// 실패 시 사용된 코드 맵에서 제거
		delete(s.usedCodes, code)
		return nil, err
	}

	// 캠페인의 발급 수 증가
	campaign.IssuedCount, err = s.repo.GetCount(campaignID)
	if err != nil {
		return nil, err
	}

	return c, nil
}

const (
	start = 0xAC00
	end   = 0xD7A3
)

func (s *Service) getHangulUniqueCode() string {
	var builder strings.Builder

	for i := 0; i < 10; i++ {
		if rand.Intn(10) != 0 {
			hangul := rune(rand.Intn(end-start+1) + start)
			builder.WriteRune(hangul)
		} else {
			number := rune(rand.Intn(10) + '0') // '0' ~ '9'
			builder.WriteRune(number)
		}
	}

	return builder.String()
}

func (s *Service) GetListCodes(ctx context.Context, campaignID int) ([]string, error) {
	coupons, err := s.repo.GetList(campaignID)
	if err != nil {
		return nil, err
	}

	return coupons, nil
}
