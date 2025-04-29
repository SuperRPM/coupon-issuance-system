package campaign

import (
	"context"
	"testing"
	"time"

	campaignrepo "github.com/SuperRPM/coupon-issuance-system/internal/repository/campaign"
)

func TestCreateCampaign(t *testing.T) {
	// 메모리 기반 리포지토리 생성
	r := campaignrepo.NewMemoryRepository()
	svc := NewService(r)

	name := "Test Campaign"
	limit := 10

	// 캠페인 생성 호출
	c, err := svc.CreateCampaign(context.Background(), name, limit, time.Now(), time.Now().AddDate(0, 0, 30))
	if err != nil {
		t.Fatalf("CreateCampaign 실패: %v", err)
	}

	// 반환된 캠페인 필드 검증
	if c.ID != 1 {
		t.Errorf("ID: 예상 1, 실제 %d", c.ID)
	}
	if c.Name != name {
		t.Errorf("Name: 예상 %q, 실제 %q", name, c.Name)
	}
	if c.Limit != limit {
		t.Errorf("Limit: 예상 %d, 실제 %d", limit, c.Limit)
	}
	if c.IssuedCount != 0 {
		t.Errorf("IssuedCount: 예상 0, 실제 %d", c.IssuedCount)
	}

	// 저장소에 올바르게 저장되었는지 검증
	stored, err := r.Get(c.ID)
	if err != nil {
		t.Fatalf("Get 실패: %v", err)
	}
	if stored == nil {
		t.Fatal("저장된 캠페인이 존재하지 않습니다")
	}
	// 저장된 값과 반환된 값 비교
	if stored.ID != c.ID || stored.Name != c.Name || stored.Limit != c.Limit || stored.IssuedCount != c.IssuedCount {
		t.Errorf("저장된 캠페인 필드 불일치: %+v vs %+v", stored, c)
	}
}
