module github.com/SuperRPM/coupon-issuance-system

go 1.21

require (
	connectrpc.com/connect v1.15.0
	github.com/rs/cors v1.10.1
	google.golang.org/protobuf v1.32.0
)

replace github.com/SuperRPM/coupon-issuance-system/gen => ./gen/proto
