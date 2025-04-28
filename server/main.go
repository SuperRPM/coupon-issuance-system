package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	campaignv1 "github.com/SuperRPM/coupon-issuance-system/gen/proto/campaign/v1"
	couponv1 "github.com/SuperRPM/coupon-issuance-system/gen/proto/coupon/v1"
	"github.com/SuperRPM/coupon-issuance-system/gen/proto/coupon/v1/couponv1connect"
	"github.com/rs/cors"
)

type CouponServer struct{}

func (s *CouponServer) IssueCoupon(
	ctx context.Context,
	req *connect.Request[couponv1.IssueCouponRequest],
) (*connect.Response[couponv1.IssueCouponResponse], error) {
	log.Println("IssueCoupon called with:", req.Msg)

	// TODO: 실제 쿠폰 발급 로직 구현
	response := &couponv1.IssueCouponResponse{
		CouponId:   1,
		CouponCode: "TEST-CODE-123",
	}

	return connect.NewResponse(response), nil
}

type CampaignServer struct{}

func (s *CampaignServer) CreateCampaign(
	ctx context.Context,
	req *connect.Request[campaignv1.CreateCampaignRequest],
) (*connect.Response[campaignv1.CreateCampaignResponse], error) {
	log.Println("CreateCampaign called with:", req.Msg)

	// TODO: 실제 캠페인 생성 로직 구현
	response := &campaignv1.CreateCampaignResponse{
		Id:          1,
		Name:        req.Msg.Name,
		Limit:       req.Msg.Limit,
		IssuedCount: 0,
	}

	return connect.NewResponse(response), nil
}

func (s *CampaignServer) GetCampaign(
	ctx context.Context,
	req *connect.Request[campaignv1.GetCampaignRequest],
) (*connect.Response[campaignv1.GetCampaignResponse], error) {
	log.Println("GetCampaign called with:", req.Msg)

	// TODO: 실제 캠페인 조회 로직 구현
	response := &campaignv1.GetCampaignResponse{
		Id:          1,
		Name:        "test-campaign",
		Limit:       100,
		IssuedCount: 0,
	}

	return connect.NewResponse(response), nil
}

func main() {
	couponServer := &CouponServer{}
	mux := http.NewServeMux()
	path, handler := couponv1connect.NewCouponServiceHandler(couponServer)
	mux.Handle(path, handler)

	// CORS 설정
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowedHeaders: []string{"*"},
	}).Handler(mux)

	log.Println("서버가 8080 포트에서 시작됩니다...")
	if err := http.ListenAndServe(":8080", corsHandler); err != nil {
		log.Fatal(err)
	}
}
