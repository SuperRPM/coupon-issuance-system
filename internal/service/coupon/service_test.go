package coupon

import (
	"context"
	"sync"
	"testing"
	"time"

	campaignrepo "github.com/SuperRPM/coupon-issuance-system/internal/repository/campaign"
	couponrepo "github.com/SuperRPM/coupon-issuance-system/internal/repository/coupon"
	campaignservice "github.com/SuperRPM/coupon-issuance-system/internal/service/campaign"
)

func TestConcurrentIssueCoupon(t *testing.T) {
	// 메모리 기반 리포지토리 및 서비스 생성
	campaignRepo := campaignrepo.NewMemoryRepository()
	couponRepo := couponrepo.NewMemoryRepository()

	campaignService := campaignservice.NewService(campaignRepo, couponRepo)
	couponService := NewService(couponRepo, campaignService)

	// 캠페인 생성 (limit 설정)
	limit := 500
	start := time.Now()
	end := start.Add(time.Hour)
	camp, err := campaignService.CreateCampaign(context.Background(), "LoadTest", limit, start, end)
	if err != nil {
		t.Fatalf("CreateCampaign 실패: %v", err)
	}

	// 동시 요청 수: limit 초과하도록 설정
	requests := 1000
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
					t.Errorf("예상치 못한 에러: %v", err)
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
