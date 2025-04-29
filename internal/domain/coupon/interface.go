package coupon

import "context"

type Handler interface {
	IssueCoupon(ctx context.Context, campaignID int) (*Coupon, error)
	// GetListCodes(ctx context.Context, campaignID int) ([]string, error)
}

type Service interface {
	IssueCoupon(ctx context.Context, campaignID int) (*Coupon, error)
	GetListCodes(ctx context.Context, campaignID int) ([]string, error)
}

type Repository interface {
	Create(coupon *Coupon) error
	GetList(campaignID int) ([]string, error)
	GetCount(campaignID int) (int, error)
}
