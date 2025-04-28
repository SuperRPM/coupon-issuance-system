package coupon

import (
	"context"

	"github.com/SuperRPM/coupon-issuance-system/internal/domain/coupon"
)

// Service는 쿠폰 관련 비즈니스 로직을 처리합니다.
type Service struct {
	repo coupon.Repository
}

// NewService는 새로운 쿠폰 서비스를 생성합니다.
func NewService(repo coupon.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// IssueCoupon은 새로운 쿠폰을 발급합니다.
func (s *Service) IssueCoupon(ctx context.Context, campaignID int) (*coupon.Coupon, error) {
	c := coupon.NewCoupon(campaignID, code)
	if err := s.repo.Create(c); err != nil {
		return nil, err
	}
	return c, nil
}
