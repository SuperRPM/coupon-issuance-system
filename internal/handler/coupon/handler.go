package coupon

import (
	"context"
	"log"

	"connectrpc.com/connect"
	couponv1 "github.com/SuperRPM/coupon-issuance-system/gen/proto/coupon/v1"
	"github.com/SuperRPM/coupon-issuance-system/internal/service/coupon"
)

type Handler struct {
	service *coupon.Service
}

func NewHandler(service *coupon.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) IssueCoupon(
	ctx context.Context,
	req *connect.Request[couponv1.IssueCouponRequest],
) (*connect.Response[couponv1.IssueCouponResponse], error) {
	log.Println("IssueCoupon called with:", req.Msg)

	c, err := h.service.IssueCoupon(ctx, int(req.Msg.CampaignId))
	if err != nil {
		return nil, err
	}

	response := &couponv1.IssueCouponResponse{
		CouponId:   int32(c.ID),
		CouponCode: c.Code,
	}

	return connect.NewResponse(response), nil
}
