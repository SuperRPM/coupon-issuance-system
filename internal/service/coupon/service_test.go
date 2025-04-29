package coupon

import (
	"context"
	"sync"
	"testing"
	"time"
	"unicode/utf8"

	campaignrepo "github.com/SuperRPM/coupon-issuance-system/internal/repository/campaign"
	couponrepo "github.com/SuperRPM/coupon-issuance-system/internal/repository/coupon"
	campaignservice "github.com/SuperRPM/coupon-issuance-system/internal/service/campaign"
)

func TestConcurrentIssueCoupon(t *testing.T) {
	campaignRepo := campaignrepo.NewMemoryRepository()
	couponRepo := couponrepo.NewMemoryRepository()

	campaignService := campaignservice.NewService(campaignRepo, couponRepo)
	couponService := NewService(couponRepo, campaignService)

	// 캠페인 생성
	limit := 20000
	start := time.Now()
	end := start.Add(time.Hour)
	camp, err := campaignService.CreateCampaign(context.Background(), "LoadTest", limit, start, end)
	if err != nil {
		t.Fatalf("CreateCampaign 실패: %v", err)
	}

	// 동시 요청 수: limit 초과하도록 설정
	requests := 25000
	var wg sync.WaitGroup
	var mu sync.Mutex
	successCount := 0
	errorCount := 0

	// 동시 발급 시뮬레이션
	for i := 0; i < requests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := couponService.IssueCoupon(context.Background(), camp.ID)
			mu.Lock()
			defer mu.Unlock()
			if err != nil {
				errorCount++
				if err.Error() != "campaign limit exceeded" {
					t.Errorf("over booking: %v", err)
				}
			} else {
				successCount++
			}
		}()
	}
	wg.Wait()

	// limit 초과 확인
	couponCount, err := couponRepo.GetCount(camp.ID)
	if err != nil {
		t.Fatalf("GetCount 실패: %v", err)
	}
	if couponCount > limit {
		t.Errorf("limit 초과: %d", couponCount)
	}

	// 성공한 발급 수 검증
	if successCount != limit {
		t.Errorf("성공한 발급 수: 예상 %d, 실제 %d", limit, successCount)
	}

	// 에러 메시지 검증
	if errorCount != requests-limit {
		t.Errorf("에러 발생 수: 예상 %d, 실제 %d", requests-limit, errorCount)
	}
}

func TestIssueCoupon(t *testing.T) {
	// 메모리 기반 리포지토리 및 서비스 생성
	campaignRepo := campaignrepo.NewMemoryRepository()
	couponRepo := couponrepo.NewMemoryRepository()
	campaignService := campaignservice.NewService(campaignRepo, couponRepo)
	couponService := NewService(couponRepo, campaignService)

	// 1) 캠페인 시작 전: 발급 거절
	startFuture := time.Now().Add(time.Hour)
	endFuture := startFuture.Add(time.Hour * 24)
	futureCamp, err := campaignService.CreateCampaign(context.Background(), "Future Campaign", 10, startFuture, endFuture)
	if err != nil {
		t.Fatalf("CreateCampaign 실패: %v", err)
	}
	_, err = couponService.IssueCoupon(context.Background(), futureCamp.ID)
	if err == nil || err.Error() != "campaign not started" {
		t.Errorf("시작 전 발급이 거절되지 않음: %v", err)
	}

	// 2) 캠페인 활성: 발급 허용 및 코드 검증
	startPast := time.Now().Add(-time.Hour)
	endPast := time.Now().Add(time.Hour * 24)
	activeCamp, err := campaignService.CreateCampaign(context.Background(), "Active Campaign", 10, startPast, endPast)
	if err != nil {
		t.Fatalf("CreateCampaign 실패: %v", err)
	}
	issuedCoupon, err := couponService.IssueCoupon(context.Background(), activeCamp.ID)
	if err != nil {
		t.Errorf("활성 캠페인 발급 실패: %v", err)
	}
	if utf8.RuneCountInString(issuedCoupon.Code) != 10 {
		t.Errorf("쿠폰 코드 글자 수: 예상 10, 실제 %d, 코드: %s", utf8.RuneCountInString(issuedCoupon.Code), issuedCoupon.Code)
	}
	codes, err := couponService.GetListCodes(context.Background(), activeCamp.ID)
	if err != nil {
		t.Fatalf("GetListCodes 실패: %v", err)
	}
	if len(codes) != 1 {
		t.Errorf("발급된 쿠폰 수: 예상 1, 실제 %d", len(codes))
	}

	// 3) 캠페인 만료: 발급 거절
	startExpired := time.Now().Add(-time.Hour * 48)
	endExpired := time.Now().Add(-time.Hour)
	expiredCamp, err := campaignService.CreateCampaign(context.Background(), "Expired Campaign", 10, startExpired, endExpired)
	if err != nil {
		t.Fatalf("CreateCampaign 실패: %v", err)
	}
	_, err = couponService.IssueCoupon(context.Background(), expiredCamp.ID)
	if err == nil || err.Error() != "campaign expired" {
		t.Errorf("만료 후 발급이 거절되지 않음: %v", err)
	}
}
