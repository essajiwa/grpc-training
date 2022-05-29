package main

//go:generate mockgen -source=service/order/order.go -package=order -destination=service/order/order_mock_test.go
//go:generate mockgen -source=repository/order_repo/order.go -package=order_repo -destination=repository/order_repo/order_repo_mock_test.go
//go:generate mockgen -source=grpcapi/server.go -package=grpcapi -destination=grpcapi/grpcapi_mock_test.go
