package grpcapi

import (
	"context"
	pb "orderservice/proto/order"
	"reflect"
	"testing"

	"google.golang.org/grpc"
)

func TestServer_GetOrder(t *testing.T) {
	type fields struct {
		UnimplementedOrderServer pb.UnimplementedOrderServer
		server                   *grpc.Server
		orderSvc                 IOrderService
	}
	type args struct {
		ctx     context.Context
		orderID *pb.OrderID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.OrderResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedOrderServer: tt.fields.UnimplementedOrderServer,
				server:                   tt.fields.server,
				orderSvc:                 tt.fields.orderSvc,
			}
			got, err := s.GetOrder(tt.args.ctx, tt.args.orderID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
