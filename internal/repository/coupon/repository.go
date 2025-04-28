package coupon

import (
	"math/rand"
	"sync"
	"time"

	"github.com/SuperRPM/coupon-issuance-system/internal/domain/coupon"
)

// MemoryRepository는 메모리에 데이터를 저장하는 쿠폰 리포지토리 구현체입니다.
type MemoryRepository struct {
	mu        sync.RWMutex
	coupons   map[int][]string
	usedCodes map[string]bool // 사용된 코드를 추적
}

// NewMemoryRepository는 새로운 메모리 리포지토리를 생성합니다.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		coupons:   make(map[int][]string),
		usedCodes: make(map[string]bool),
	}
}

// generateCode는 한글과 숫자를 포함한 고유한 쿠폰 코드를 생성합니다.
func (r *MemoryRepository) generateCode() string {
	// 한글 초성, 중성, 종성
	initial := []rune{'ㄱ', 'ㄴ', 'ㄷ', 'ㄹ', 'ㅁ', 'ㅂ', 'ㅅ', 'ㅇ', 'ㅈ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ'}
	medial := []rune{'ㅏ', 'ㅑ', 'ㅓ', 'ㅕ', 'ㅗ', 'ㅛ', 'ㅜ', 'ㅠ', 'ㅡ', 'ㅣ'}
	final := []rune{' ', 'ㄱ', 'ㄴ', 'ㄷ', 'ㄹ', 'ㅁ', 'ㅂ', 'ㅅ', 'ㅇ', 'ㅈ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ'}

	// 랜덤 시드 설정
	rand.Seed(time.Now().UnixNano())

	// 최대 100번 시도
	for i := 0; i < 100; i++ {
		// 한글 문자 생성
		initialIdx := rand.Intn(len(initial))
		medialIdx := rand.Intn(len(medial))
		finalIdx := rand.Intn(len(final))

		// 유니코드 한글 문자 생성
		unicode := 0xAC00 + (initialIdx * 21 * 28) + (medialIdx * 28) + finalIdx
		hangul := string(rune(unicode))

		// 숫자 생성 (1-9)
		number := rand.Intn(9) + 1

		// 최종 코드 생성 (한글 + 숫자)
		code := hangul + string(rune('0'+number))

		// 코드가 이미 사용되었는지 확인
		if !r.usedCodes[code] {
			r.usedCodes[code] = true
			return code
		}
	}

	// 모든 시도가 실패하면 기본 코드 반환
	return "가1"
}

// Create는 새로운 쿠폰을 생성합니다.
func (r *MemoryRepository) Create(c *coupon.Coupon) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 코드가 비어있으면 새로 생성
	if c.Code == "" {
		c.Code = r.generateCode()
	}

	// 코드가 이미 사용되었는지 확인
	if r.usedCodes[c.Code] {
		return nil // 이미 사용된 코드는 무시
	}

	r.usedCodes[c.Code] = true
	r.coupons[c.CampaignID] = append(r.coupons[c.CampaignID], c.Code)
	return nil
}
