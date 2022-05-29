package order

import (
	"context"
	"orderservice/model"
	"reflect"
	"testing"
)

func TestOrderService_GetOrder(t *testing.T) {
	type fields struct {
		order IOrderRepository
	}
	type args struct {
		ctx     context.Context
		orderID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := OrderService{
				order: tt.fields.order,
			}
			got, err := o.GetOrder(tt.args.ctx, tt.args.orderID)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderService.GetOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.GetOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
