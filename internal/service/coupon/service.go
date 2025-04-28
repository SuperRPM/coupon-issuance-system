package coupon

import (
	"context"
	"errors"

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
func (s *Service) IssueCoupon(ctx context.Context, campaignID int, code string) (*coupon.Coupon, error) {
	c := coupon.NewCoupon(campaignID, code)
	if err := s.repo.Create(c); err != nil {
		return nil, err
	}
	return c, nil
}

// GetCoupon은 ID로 쿠폰을 조회합니다.
func (s *Service) GetCoupon(ctx context.Context, id string) (*coupon.Coupon, error) {
	return s.repo.Get(id)
}

// UseCoupon은 쿠폰을 사용합니다.
func (s *Service) UseCoupon(ctx context.Context, id string) error {
	c, err := s.repo.Get(id)
	if err != nil {
		return err
	}

	if c.Used {
		return errors.New("coupon already used")
	}

	c.Use()
	return s.repo.Update(c)
}
