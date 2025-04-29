# Coupon Issuance System

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
  - 신규 유저 전용 쿠폰, 다수의 캠페인 적용 쿠폰 등 다양한 케이스로 확장 가능

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
- 실제 서비스 환경에서는 Redis 등 외부 저장소 도입 고려

## 요구사항

1. 정확한 발급 수량 제한
   - 캠페인 한도를 초과하지 않도록 보장 (초과 발급 방지)
2. 자동 발급 시작/종료 처리
   - 시작 전 요청 시 거절, 종료 후 요청 시 거절
3. 데이터 일관성 보장
   - 동시성 제어로 발급 과정을 원자적으로 처리
4. 고유 코드 생성
   - 숫자와 한글 조합으로 최대 10글자 고유 코드 발급

## 구현 세부사항

### Coupon Service
- `IssueCoupon` 메서드는 `sync.Mutex`로 동시성 제어
- `getHangulUniqueCode` 함수로 한글/숫자 혼합 랜덤 코드 생성

### Campaign Service
- `GetCampaign` 호출 시 쿠폰 리포지토리에서 발급된 쿠폰 수를 조회하여 `IssuedCount` 동기화

## 시험 및 테스트

- **Unit Test**: `CreateCampaign`, `GetCampaign`, `IssueCoupon` 기능 검증
- **부하 테스트**: `TestConcurrentIssueCoupon`으로 1,000건 동시 요청, 한도 검증

## 엣지 케이스 및 향후 과제

- 시작일 ≥ 종료일, `limit` ≤ 0, 캠페인 이름 중복 등 입력 검증 강화
- `context` 취소/타임아웃 대응
