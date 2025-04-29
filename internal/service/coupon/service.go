package coupon

import (
	"context"
	"errors"
	"math/rand"
	"strings"
	"sync"

	"github.com/SuperRPM/coupon-issuance-system/internal/domain/coupon"
	"github.com/SuperRPM/coupon-issuance-system/internal/service/campaign"
	"github.com/google/uuid"
)

// Service는 쿠폰 관련 비즈니스 로직을 처리합니다.
type Service struct {
	repo            coupon.Repository
	campaignService *campaign.CampaignService
	usedCodes       map[string]bool
	mu              sync.RWMutex
}

// NewService는 새로운 쿠폰 서비스를 생성합니다.
func NewService(repo coupon.Repository, campaignService *campaign.CampaignService) *Service {
	return &Service{
		repo:            repo,
		campaignService: campaignService,
		usedCodes:       make(map[string]bool),
	}
}

// IssueCoupon은 새로운 쿠폰을 발급합니다.
func (s *Service) IssueCoupon(ctx context.Context, campaignID int) (*coupon.Coupon, error) {
	// 캠페인 조회
	campaign, err := s.campaignService.GetCampaign(ctx, campaignID)
	if err != nil {
		return nil, err
	}

	// 캠페인이 존재하지 않는 경우
	if campaign == nil {
		return nil, errors.New("campaign not found")
	}

	// 발급 제한 확인
	if !campaign.CanIssue() {
		return nil, errors.New("campaign limit exceeded")
	}

	// 쿠폰 코드 생성
	code := s.convertUUIDToHangul()

	// 쿠폰 생성
	c := coupon.NewCoupon(campaignID, code)
	if err := s.repo.Create(c); err != nil {
		// 실패 시 사용된 코드 맵에서 제거
		s.mu.Lock()
		delete(s.usedCodes, code)
		s.mu.Unlock()
		return nil, err
	}

	// 캠페인의 발급 수 증가
	campaign.Issue()

	return c, nil
}

var baseHangulMap = map[rune][]string{
	'a': {"가", "갸", "거", "겨", "고", "교", "구", "규", "그", "기"},
	'b': {"나", "냐", "너", "녀", "노", "뇨", "누", "뉴", "느", "니"},
	'c': {"다", "댜", "더", "뎌", "도", "됴", "두", "듀", "드", "디"},
	'd': {"라", "랴", "러", "려", "로", "료", "루", "류", "르", "리"},
	'e': {"마", "먀", "머", "며", "모", "묘", "무", "뮤", "므", "미"},
	'f': {"바", "뱌", "버", "벼", "보", "뵤", "부", "뷰", "브", "비"},
	'g': {"사", "샤", "서", "셔", "소", "쇼", "수", "슈", "스", "시"},
	'h': {"아", "야", "어", "여", "오", "요", "우", "유", "으", "이"},
	'i': {"자", "쟈", "저", "져", "조", "죠", "주", "쥬", "즈", "지"},
	'j': {"차", "챠", "처", "쳐", "초", "쵸", "추", "츄", "츠", "치"},
	'k': {"카", "캬", "커", "켜", "코", "쿄", "쿠", "큐", "크", "키"},
	'l': {"타", "탸", "터", "텨", "토", "툐", "투", "튜", "트", "티"},
	'm': {"파", "퍄", "퍼", "펴", "포", "표", "푸", "퓨", "프", "피"},
	'n': {"하", "햐", "허", "혀", "호", "효", "후", "휴", "흐", "히"},
}

var extendedHangulMap = map[rune][]string{
	'o': {"강", "걍", "겅", "경", "공", "굥", "궁", "귱", "긍", "깅"},
	'p': {"낭", "냥", "넝", "녕", "농", "뇽", "눙", "늉", "능", "닝"},
	'q': {"당", "댱", "덩", "뎅", "동", "됑", "둥", "듕", "등", "딩"},
	'r': {"랑", "랭", "렁", "렝", "롱", "룡", "룽", "륑", "릉", "링"},
	's': {"망", "먕", "멍", "명", "몽", "묭", "뭉", "묭", "믕", "밍"},
	't': {"방", "뱡", "벙", "병", "봉", "뵹", "붕", "븅", "븡", "빙"},
	'u': {"상", "샹", "성", "셩", "송", "숑", "숭", "슝", "승", "싱"},
	'v': {"앙", "양", "엉", "영", "옹", "용", "웅", "융", "응", "잉"},
	'w': {"장", "쟝", "정", "젱", "종", "죵", "중", "쥉", "증", "징"},
	'x': {"창", "챵", "청", "쳥", "총", "쵱", "충", "츙", "층", "칭"},
	'y': {"캉", "캥", "컁", "켕", "콩", "쿵", "쿙", "큥", "킁", "킹"},
	'z': {"탕", "탱", "텅", "텡", "통", "퉁", "퇑", "튕", "틍", "팅"},
}

func (s *Service) convertUUIDToHangul() string {
	uuidStr := strings.ReplaceAll(uuid.New().String(), "-", "")
	var builder strings.Builder

	var result string

	for i := 0; i < 10; i++ {
		var group string
		if i < 9 {
			group = uuidStr[i*3 : i*3+3] // 3자리씩
		} else {
			group = uuidStr[i*3:] // 마지막 5자리
		}

		sum := 0
		for _, c := range group {
			sum += int(c)
		}

		// sum을 48~57(숫자) 또는 97~122(소문자) 범위로 매핑
		var mappedChar rune
		mapped := sum % 36 // 0~35

		if mapped < 10 {
			mappedChar = rune('0' + mapped) // 0~9
		} else {
			mappedChar = rune('a' + (mapped - 10)) // a~z
		}

		result += string(mappedChar)
	}

	for _, c := range result { // 앞 10자리만 변환
		if c >= '0' && c <= '9' {
			// 숫자면 numberHangul에서 선택
			builder.WriteString(string(c))
		} else if c >= 'a' && c <= 'n' {
			choices := baseHangulMap[c]
			builder.WriteString(choices[rand.Intn(len(choices))])
		} else if c >= 'o' && c <= 'z' {
			choices := extendedHangulMap[c]
			builder.WriteString(choices[rand.Intn(len(choices))])
		}
	}

	// builder.String() 값이 이미 사용된 코드인지 확인
	// 사용된 코드라면 다시 생성
	s.mu.Lock()
	if s.usedCodes[builder.String()] {
		return s.convertUUIDToHangul()
	}

	s.usedCodes[builder.String()] = true
	s.mu.Unlock()
	return builder.String()
}
