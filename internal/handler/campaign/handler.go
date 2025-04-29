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

type CampaignHandler struct {
	campaignService *campaign.CampaignService
	couponService   *coupon.Service
}

func NewHandler(campaignService *campaign.CampaignService, couponService *coupon.Service) *CampaignHandler {
	return &CampaignHandler{
		campaignService: campaignService,
		couponService:   couponService,
	}
}

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
		Id:        int32(c.ID),
		Name:      c.Name,
		Limit:     int32(c.Limit),
		StartDate: timestamppb.New(c.StartDate),
		EndDate:   timestamppb.New(c.EndDate),
	}

	return connect.NewResponse(response), nil
}

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
		StartDate:   timestamppb.New(c.StartDate),
		EndDate:     timestamppb.New(c.EndDate),
		CouponCodes: couponCodes,
	}

	return connect.NewResponse(response), nil
}
