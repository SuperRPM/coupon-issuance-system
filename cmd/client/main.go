package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	couponv1 "github.com/SuperRPM/coupon-issuance-system/gen/proto/coupon/v1"
	"github.com/SuperRPM/coupon-issuance-system/gen/proto/coupon/v1/couponv1connect"
)

func main() {
	client := couponv1connect.NewCouponServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)

	req := connect.NewRequest(&couponv1.IssueCouponRequest{
		CampaignId: 1,
	})

	res, err := client.IssueCoupon(context.Background(), req)
	if err != nil {
		log.Fatal("IssueCoupon 호출 실패:", err)
	}

	log.Printf("응답 받음: CouponId=%d, CouponCode=%s\n", res.Msg.CouponId, res.Msg.CouponCode)
}
