package coupon

import (
	"sync"

	"github.com/SuperRPM/coupon-issuance-system/internal/domain/coupon"
)

type MemoryRepository struct {
	mu      sync.RWMutex
	coupons map[int][]string
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		coupons: make(map[int][]string),
	}
}

func (r *MemoryRepository) Create(c *coupon.Coupon) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.coupons[c.CampaignID] = append(r.coupons[c.CampaignID], c.Code)
	return nil
}

func (r *MemoryRepository) GetCount(campaignID int) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return len(r.coupons[campaignID]), nil
}
