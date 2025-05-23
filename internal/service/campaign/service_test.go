package campaign

import (
	"context"
	"testing"
	"time"

	campaignrepo "github.com/SuperRPM/coupon-issuance-system/internal/repository/campaign"
	couponrepo "github.com/SuperRPM/coupon-issuance-system/internal/repository/coupon"
)

func TestCreateCampaign(t *testing.T) {
	// 리포지토리 및 서비스 생성
	campaignRepo := campaignrepo.NewMemoryRepository()
	couponRepo := couponrepo.NewMemoryRepository()
	svc := NewService(campaignRepo, couponRepo)

	name := "Test Campaign"
	limit := 10

	start := time.Now()
	end := start.AddDate(0, 0, 30)

	c, err := svc.CreateCampaign(context.Background(), name, limit, start, end)
	if err != nil {
		t.Fatalf("CreateCampaign 실패: %v", err)
	}

	if c.ID != 1 {
		t.Errorf("ID: 예상 1, 실제 %d", c.ID)
	}
	if c.Name != name {
		t.Errorf("Name: 예상 %q, 실제 %q", name, c.Name)
	}
	if c.Limit != limit {
		t.Errorf("Limit: 예상 %d, 실제 %d", limit, c.Limit)
	}
	if !c.StartDate.Equal(start) {
		t.Errorf("StartDate: 예상 %v, 실제 %v", start, c.StartDate)
	}
	if !c.EndDate.Equal(end) {
		t.Errorf("EndDate: 예상 %v, 실제 %v", end, c.EndDate)
	}

	stored, err := campaignRepo.Get(c.ID)
	if err != nil {
		t.Fatalf("Get 실패: %v", err)
	}
	if stored == nil {
		t.Fatal("저장된 캠페인이 존재하지 않습니다")
	}
	if stored.ID != c.ID || stored.Name != c.Name || stored.Limit != c.Limit {
		t.Errorf("저장된 캠페인 필드 불일치: %+v vs %+v", stored, c)
	}
}
