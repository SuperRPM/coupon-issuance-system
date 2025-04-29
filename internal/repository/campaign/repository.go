package campaign

import (
	"sync"

	"github.com/SuperRPM/coupon-issuance-system/internal/domain/campaign"
)

// MemoryRepository는 메모리에 데이터를 저장하는 캠페인 리포지토리 구현체입니다.
type MemoryRepository struct {
	mu        sync.RWMutex
	campaigns map[int]*campaign.Campaign
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		campaigns: make(map[int]*campaign.Campaign),
	}
}

func (r *MemoryRepository) Create(c *campaign.Campaign) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if c.ID == 0 {
		c.ID = len(r.campaigns) + 1
	}

	r.campaigns[c.ID] = c
	return nil
}

func (r *MemoryRepository) Get(id int) (*campaign.Campaign, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	c, exists := r.campaigns[id]
	if !exists {
		return nil, nil
	}

	return c, nil
}
