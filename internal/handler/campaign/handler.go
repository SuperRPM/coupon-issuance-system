package campaign

import (
	"context"
	"log"

	"connectrpc.com/connect"
	campaignv1 "github.com/SuperRPM/coupon-issuance-system/gen/proto/campaign/v1"
	"github.com/SuperRPM/coupon-issuance-system/internal/service/campaign"
	"github.com/SuperRPM/coupon-issuance-system/internal/service/coupon"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Handler는 캠페인 관련 HTTP 핸들러를 구현합니다.
type CampaignHandler struct {
	campaignService *campaign.CampaignService
	couponService   *coupon.Service
}

// NewHandler는 새로운 캠페인 핸들러를 생성합니다.
func NewHandler(campaignService *campaign.CampaignService, couponService *coupon.Service) *CampaignHandler {
	return &CampaignHandler{
		campaignService: campaignService,
		couponService:   couponService,
	}
}

// CreateCampaign은 새로운 캠페인을 생성합니다.
func (h *CampaignHandler) CreateCampaign(
	ctx context.Context,
	req *connect.Request[campaignv1.CreateCampaignRequest],
) (*connect.Response[campaignv1.CreateCampaignResponse], error) {
	log.Println("CreateCampaign called with:", req.Msg)

	c, err := h.campaignService.CreateCampaign(ctx, req.Msg.Name, int(req.Msg.Limit), req.Msg.StartDate.AsTime(), req.Msg.EndDate.AsTime())
	if err != nil {
		return nil, err
	}

	response := &campaignv1.CreateCampaignResponse{
		Id:          int32(c.ID),
		Name:        c.Name,
		Limit:       int32(c.Limit),
		IssuedCount: int32(c.IssuedCount),
		StartDate:   timestamppb.New(c.StartDate),
		EndDate:     timestamppb.New(c.EndDate),
	}

	return connect.NewResponse(response), nil
}

// GetCampaign은 ID로 캠페인을 조회합니다.
func (h *CampaignHandler) GetCampaign(
	ctx context.Context,
	req *connect.Request[campaignv1.GetCampaignRequest],
) (*connect.Response[campaignv1.GetCampaignResponse], error) {
	log.Println("GetCampaign called with:", req.Msg)

	c, err := h.campaignService.GetCampaign(ctx, int(req.Msg.Id))
	if err != nil {
		return nil, err
	}

	couponCodes, err := h.couponService.GetListCodes(ctx, int(req.Msg.Id))
	if err != nil {
		return nil, err
	}

	response := &campaignv1.GetCampaignResponse{
		Id:          int32(c.ID),
		Name:        c.Name,
		Limit:       int32(c.Limit),
		IssuedCount: int32(c.IssuedCount),
		StartDate:   timestamppb.New(c.StartDate),
		EndDate:     timestamppb.New(c.EndDate),
		CouponCodes: couponCodes,
	}

	return connect.NewResponse(response), nil
}
