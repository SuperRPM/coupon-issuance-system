package coupon

import (
	"context"
	"log"

	"connectrpc.com/connect"
	couponv1 "github.com/SuperRPM/coupon-issuance-system/gen/proto/coupon/v1"
	"github.com/SuperRPM/coupon-issuance-system/internal/service/coupon"
)

// Handler는 쿠폰 관련 HTTP 핸들러를 구현합니다.
type Handler struct {
	service *coupon.Service
}

// NewHandler는 새로운 쿠폰 핸들러를 생성합니다.
func NewHandler(service *coupon.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// IssueCoupon은 새로운 쿠폰을 발급합니다.
func (h *Handler) IssueCoupon(
	ctx context.Context,
	req *connect.Request[couponv1.IssueCouponRequest],
) (*connect.Response[couponv1.IssueCouponResponse], error) {
	log.Println("IssueCoupon called with:", req.Msg)

	c, err := h.service.IssueCoupon(ctx, int(req.Msg.CampaignId), req.Msg.Code)
	if err != nil {
		return nil, err
	}

	response := &couponv1.IssueCouponResponse{
		Id:         int32(c.ID),
		CampaignId: int32(c.CampaignID),
		Code:       c.Code,
		Used:       c.Used,
	}

	return connect.NewResponse(response), nil
}

// GetCoupon은 ID로 쿠폰을 조회합니다.
func (h *Handler) GetCoupon(
	ctx context.Context,
	req *connect.Request[couponv1.GetCouponRequest],
) (*connect.Response[couponv1.GetCouponResponse], error) {
	log.Println("GetCoupon called with:", req.Msg)

	c, err := h.service.GetCoupon(ctx, int(req.Msg.CouponId))
	if err != nil {
		return nil, err
	}

	response := &couponv1.GetCouponResponse{
		Id:         int32(c.ID),
		CampaignId: int32(c.CampaignID),
		Code:       c.Code,
		Used:       c.Used,
	}

	return connect.NewResponse(response), nil
}

// UseCoupon은 쿠폰을 사용합니다.
func (h *Handler) UseCoupon(
	ctx context.Context,
	req *connect.Request[couponv1.UseCouponRequest],
) (*connect.Response[couponv1.UseCouponResponse], error) {
	log.Println("UseCoupon called with:", req.Msg)

	c, err := h.service.UseCoupon(ctx, int(req.Msg.CouponId))
	if err != nil {
		return nil, err
	}

	response := &couponv1.UseCouponResponse{
		Id:         int32(c.ID),
		CampaignId: int32(c.CampaignID),
		Code:       c.Code,
		Used:       c.Used,
	}

	return connect.NewResponse(response), nil
}
