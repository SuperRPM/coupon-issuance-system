package main

import (
	"log"
	"net/http"

	campaignv1connect "github.com/SuperRPM/coupon-issuance-system/gen/proto/campaign/v1/campaignv1connect"
	couponv1connect "github.com/SuperRPM/coupon-issuance-system/gen/proto/coupon/v1/couponv1connect"
	campaignhandler "github.com/SuperRPM/coupon-issuance-system/internal/handler/campaign"
	couponhandler "github.com/SuperRPM/coupon-issuance-system/internal/handler/coupon"
	campaignservice "github.com/SuperRPM/coupon-issuance-system/internal/service/campaign"
	couponservice "github.com/SuperRPM/coupon-issuance-system/internal/service/coupon"
	"github.com/rs/cors"
)

func main() {
	// 서비스 생성
	campaignService := campaignservice.NewService(nil) // TODO: 리포지토리 구현체 주입
	couponService := couponservice.NewService(nil)     // TODO: 리포지토리 구현체 주입

	// 핸들러 생성
	campaignHandler := campaignhandler.NewHandler(campaignService)
	couponHandler := couponhandler.NewHandler(couponService)

	// HTTP 라우터 설정
	mux := http.NewServeMux()

	// 캠페인 핸들러 등록
	mux.Handle(campaignv1connect.NewCampaignServiceHandler(campaignHandler))

	// 쿠폰 핸들러 등록
	mux.Handle(couponv1connect.NewCouponServiceHandler(couponHandler))

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
