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

// Service는 쿠폰 관련 비즈니스 로직을 처리합니다.
type Service struct {
	repo            coupon.Repository
	campaignService *campaign.CampaignService
	usedCodes       map[string]bool
	mu              sync.RWMutex
}

// NewService는 새로운 쿠폰 서비스를 생성합니다.
func NewService(repo coupon.Repository, campaignService *campaign.CampaignService) *Service {
	return &Service{
		repo:            repo,
		campaignService: campaignService,
		usedCodes:       make(map[string]bool),
	}
}

// IssueCoupon은 새로운 쿠폰을 발급합니다.
func (s *Service) IssueCoupon(ctx context.Context, campaignID int) (*coupon.Coupon, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 캠페인 조회
	campaign, err := s.campaignService.GetCampaign(ctx, campaignID)
	if err != nil {
		return nil, err
	}

	// 캠페인이 존재하지 않는 경우
	if campaign == nil {
		return nil, errors.New("campaign not found")
	}

	// 쿠폰 레포지토리에서 발급된 쿠폰 수 가져오기
	couponCount, err := s.repo.GetCount(campaignID)
	if err != nil {
		return nil, err
	}

	// 발급 제한 확인
	if campaign.Limit <= couponCount {
		return nil, errors.New("campaign limit exceeded")
	}

	// 날짜 제한 확인
	if campaign.StartDate.After(time.Now()) {
		return nil, errors.New("campaign not started")
	}

	if campaign.EndDate.Before(time.Now()) {
		return nil, errors.New("campaign expired")
	}

	// 쿠폰 코드 생성
	code := s.getHangulUniqueCode()

	// 쿠폰 생성
	c := coupon.NewCoupon(campaignID, code)
	if err := s.repo.Create(c); err != nil {
		// 실패 시 사용된 코드 맵에서 제거
		delete(s.usedCodes, code)
		return nil, err
	}

	// 캠페인의 발급 수 증가
	campaign.Issue()

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
			// 한글 생성
			hangul := rune(rand.Intn(end-start+1) + start) // 한글 유니코드 범위: 0xAC00 ~ 0xD7A3
			builder.WriteRune(hangul)
		} else {
			// 숫자 생성
			number := rune(rand.Intn(10) + '0') // '0' ~ '9'
			builder.WriteRune(number)
		}
	}

	return builder.String()
}

// GetListCodes는 캠페인 ID를 기반으로 쿠폰 코드 목록을 반환합니다.
func (s *Service) GetListCodes(ctx context.Context, campaignID int) ([]string, error) {
	coupons, err := s.repo.GetList(campaignID)
	if err != nil {
		return nil, err
	}

	return coupons, nil
}
