# Coupon Issuance System

## 프로젝트 구조

이 프로젝트는 다음과 같은 주요 디렉토리와 파일로 구성되어 있습니다:

- `cmd/server/main.go`: 서버 실행을 위한 메인 엔트리 포인트.
- `internal/`: 비즈니스 로직과 서비스 구현을 포함하는 내부 패키지.
  - `service/`: 쿠폰 및 캠페인 관련 서비스 로직.
  - `repository/`: 메모리 기반의 데이터 저장소 구현.
  - `handler/`: HTTP 요청을 처리하는 핸들러.
  - `domain/`: 의존성 관리 및 entity설정.

## 실행 방법

### 서버 실행
```bash
go run cmd/server/main.go
```

### 테스트 실행
```bash
go test ./...
```

## 설계

### 도메인 분리
- `Campaign`과 `Coupon` 엔티티를 독립적으로 관리
  - 신규 유저 전용 쿠폰, 다수의 캠페인에 적용되는 쿠폰 등 다양한 케이스로 확장 가능

### 저장소 레이어
- 메모리 기반 Repository로 데이터와 동시성을 관리
  - **Campaign Repository**
    ```go
    type MemoryRepository struct {
        mu        sync.RWMutex
        campaigns map[int]*campaign.Campaign
    }
    ```
  - **Coupon Repository**
    ```go
    type MemoryRepository struct {
        mu      sync.RWMutex
        coupons map[int][]string
    }
    ```
- 디비를 사용해도 다수의 요청이 빠르게, 많이 들어오는 서비스에서는 인메모리 방식의 데이터 관리가 적합하다고 판단하여 별도의 디비 연결 없이 메모리에서 데이터 관리
- 실제 서비스 환경에서는 Redis로 데이터를 관리하고 DB를 백업 저장소로 이용 고려

## 요구사항

1. connectrpc
   - ```connectrpc.com/connect v1.15.0``` 사용
2. 3개의 RPC 메소드
    ```proto
    // campaign.proto
    service CampaignService {
    rpc CreateCampaign(CreateCampaignRequest) returns (CreateCampaignResponse) {}
    rpc GetCampaign(GetCampaignRequest) returns (GetCampaignResponse) {}
    }
    //coupon.proto
    service CouponService {
    rpc IssueCoupon(IssueCouponRequest) returns (IssueCouponResponse) {}
    }
    ```
3. 정확한 발급 수량 제한
   - 캠페인 한도를 초과하지 않도록 보장 (초과 발급 방지)
   ```golang
   // internal/service/coupon/service.go
   	couponCount, err := s.repo.GetCount(campaignID)
	if err != nil {
		return nil, err
	}

	if campaign.Limit <= couponCount {
		return nil, errors.New("campaign limit exceeded")
	}
   ```
2. 자동 발급 시작/종료 처리
   - 일반적으로 캠페인 종료 날짜가 있을거라고생각해 종료일 추가
   - 시작 전 요청 시 거절, 종료 후 요청 시 거절
3. 데이터 일관성 보장
   - 동시성 제어로 발급 과정을 원자적으로 처리
   ```golang
    // internal/service/coupon/service.go
    func (s *Service) IssueCoupon(ctx context.Context, campaignID int) (*coupon.Coupon, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

    // internal/repository/coupon/repository.go
    func (r *MemoryRepository) Create(c *coupon.Coupon) error {
	r.mu.Lock()
	defer r.mu.Unlock()
   ```
4. 고유 코드 생성
   - 숫자와 한글 조합으로 최대 10글자 고유 코드 발급
   - uuid에 비해 10글자는 글자 수가 적어 중복될 가능성이 있다고 판단 후 아래와 같이 약 10,000글자 이상의 한글 유니코드로 랜덤생성
   - 동일한 확률일 경우 숫자가 거의 생성되지 않아서 임의로 10% 할당
   ```golang
    func (s *Service) getHangulUniqueCode() string {
	var builder strings.Builder

	for i := 0; i < 10; i++ {
		if rand.Intn(10) != 0 {
			hangul := rune(rand.Intn(end-start+1) + start)
			builder.WriteRune(hangul)
		} else {
			number := rune(rand.Intn(10) + '0') // '0' ~ '9'
			builder.WriteRune(number)
		}
	}

	return builder.String()
    }
   ```
   

## 구현 세부사항

### Coupon Service
- `IssueCoupon` 메서드는 `sync.Mutex`로 동시성 제어
- `getHangulUniqueCode` 함수로 한글/숫자 혼합 랜덤 코드 생성

### Campaign Service
- `GetCampaign` 호출 시 쿠폰 리포지토리에서 발급된 쿠폰 수를 조회하여 `IssuedCount` 동기화

## 시험 및 테스트

- **Unit Test**: `CreateCampaign`, `GetCampaign`, `IssueCoupon` 기능 검증
- **부하 테스트**: `TestConcurrentIssueCoupon`으로 25,000건 동시 요청, 제한 검증

## 엣지 케이스
1. 쿠폰 코드 발행 메모리 관리방법(예: 쿠폰이 1,000만장 이상 발행된 경우)
 - 실제 환경이라면 go의 map에서 관리하기 보단 redis 등을 통해서 관리
 - coupon repository를 구조체로 담아 아직 사용되지 않은 쿠폰만 map에서 관리
 - 실제 코드에서는 db를 생성해도 동시성, 부하테스트를 위해 메모리에서 관리되기 때문에 이 부분은 실제 구현하지는 않았음.
2. 쿠폰 발급 수량 제한이 0인 캠페인 생성 제한
 - 쿠폰발급 수량이 정해지지 않을 경우 캠페인 생성을 제한하는 것이 "일단 캠페인 생성"으로 인한 혼란을 방지할 수 있음.
3. 현재시간보다 과거시점으로 캠페인 생성 제한
 - 휴먼에러로 생성할 경우 생성과 동시에 쿠폰 발급 가능해서 금전적인 손실입을 가능성이 있기 때문에 제한
4. 캠페인 시작시간이 10년뒤거나 캠페인 기간이 10년 이상인 경우
 - 일반적이지 않은 상황으로 별도로 운영과정에서 협의가 없는 한 1년 이내시작, 1년 이내 캠페인 종료를 validation
