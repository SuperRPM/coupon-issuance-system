package main

import (
	"log"
	"net/http"

	campaignv1connect "github.com/SuperRPM/coupon-issuance-system/gen/proto/campaign/v1/campaignv1connect"
	couponv1connect "github.com/SuperRPM/coupon-issuance-system/gen/proto/coupon/v1/couponv1connect"
	campaignhandler "github.com/SuperRPM/coupon-issuance-system/internal/handler/campaign"
	couponhandler "github.com/SuperRPM/coupon-issuance-system/internal/handler/coupon"
	campaignrepo "github.com/SuperRPM/coupon-issuance-system/internal/repository/campaign"
	couponrepo "github.com/SuperRPM/coupon-issuance-system/internal/repository/coupon"
	campaignservice "github.com/SuperRPM/coupon-issuance-system/internal/service/campaign"
	couponservice "github.com/SuperRPM/coupon-issuance-system/internal/service/coupon"
	"github.com/rs/cors"
)

func main() {
	// 리포지토리 생성
	campaignRepo := campaignrepo.NewMemoryRepository()
	couponRepo := couponrepo.NewMemoryRepository()

	// 서비스 생성
	campaignService := campaignservice.NewService(campaignRepo, couponRepo)
	couponService := couponservice.NewService(couponRepo, campaignService)

	// 핸들러 생성
	campaignHandler := campaignhandler.NewHandler(campaignService, couponService)
	couponHandler := couponhandler.NewHandler(couponService)

	// Router
	mux := http.NewServeMux()

	mux.Handle(campaignv1connect.NewCampaignServiceHandler(campaignHandler))
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

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", corsHandler); err != nil {
		log.Fatal(err)
	}
}
