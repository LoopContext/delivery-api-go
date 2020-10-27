package gen

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/loopcontext/go-graphql-orm/events"
)

// ResolutionHandlers struct
type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *events.Event) error

	CreateDelivery      func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Delivery, err error)
	UpdateDelivery      func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Delivery, err error)
	DeleteDelivery      func(ctx context.Context, r *GeneratedResolver, id string) (item *Delivery, err error)
	DeleteAllDeliveries func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryDelivery       func(ctx context.Context, r *GeneratedResolver, opts QueryDeliveryHandlerOptions) (*Delivery, error)
	QueryDeliveries     func(ctx context.Context, r *GeneratedResolver, opts QueryDeliveriesHandlerOptions) (*DeliveryResultType, error)

	DeliveryVehicleType func(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *VehicleType, err error)

	DeliveryPaymentChannel func(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *PaymentChannel, err error)

	DeliveryDeliveryType func(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *DeliveryType, err error)

	DeliveryDeliveryChannel func(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *DeliveryChannel, err error)

	DeliverySender func(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *Person, err error)

	DeliveryReceiver func(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *Person, err error)

	DeliveryDeliver func(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *Person, err error)

	CreatePerson    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Person, err error)
	UpdatePerson    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Person, err error)
	DeletePerson    func(ctx context.Context, r *GeneratedResolver, id string) (item *Person, err error)
	DeleteAllPeople func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryPerson     func(ctx context.Context, r *GeneratedResolver, opts QueryPersonHandlerOptions) (*Person, error)
	QueryPeople     func(ctx context.Context, r *GeneratedResolver, opts QueryPeopleHandlerOptions) (*PersonResultType, error)

	PersonDeliveries func(ctx context.Context, r *GeneratedResolver, obj *Person) (res []*Delivery, err error)

	PersonDeliveriesSent func(ctx context.Context, r *GeneratedResolver, obj *Person) (res []*Delivery, err error)

	PersonDeliveriesReceived func(ctx context.Context, r *GeneratedResolver, obj *Person) (res []*Delivery, err error)

	PersonPaymentStatus func(ctx context.Context, r *GeneratedResolver, obj *Person) (res *PaymentStatus, err error)

	PersonPaymentHistory func(ctx context.Context, r *GeneratedResolver, obj *Person) (res *PaymentHistory, err error)

	CreateDeliveryType     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *DeliveryType, err error)
	UpdateDeliveryType     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *DeliveryType, err error)
	DeleteDeliveryType     func(ctx context.Context, r *GeneratedResolver, id string) (item *DeliveryType, err error)
	DeleteAllDeliveryTypes func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryDeliveryType      func(ctx context.Context, r *GeneratedResolver, opts QueryDeliveryTypeHandlerOptions) (*DeliveryType, error)
	QueryDeliveryTypes     func(ctx context.Context, r *GeneratedResolver, opts QueryDeliveryTypesHandlerOptions) (*DeliveryTypeResultType, error)

	CreateDeliveryChannel     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *DeliveryChannel, err error)
	UpdateDeliveryChannel     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *DeliveryChannel, err error)
	DeleteDeliveryChannel     func(ctx context.Context, r *GeneratedResolver, id string) (item *DeliveryChannel, err error)
	DeleteAllDeliveryChannels func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryDeliveryChannel      func(ctx context.Context, r *GeneratedResolver, opts QueryDeliveryChannelHandlerOptions) (*DeliveryChannel, error)
	QueryDeliveryChannels     func(ctx context.Context, r *GeneratedResolver, opts QueryDeliveryChannelsHandlerOptions) (*DeliveryChannelResultType, error)

	CreateVehicleType     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *VehicleType, err error)
	UpdateVehicleType     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *VehicleType, err error)
	DeleteVehicleType     func(ctx context.Context, r *GeneratedResolver, id string) (item *VehicleType, err error)
	DeleteAllVehicleTypes func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryVehicleType      func(ctx context.Context, r *GeneratedResolver, opts QueryVehicleTypeHandlerOptions) (*VehicleType, error)
	QueryVehicleTypes     func(ctx context.Context, r *GeneratedResolver, opts QueryVehicleTypesHandlerOptions) (*VehicleTypeResultType, error)

	CreatePaymentChannel     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *PaymentChannel, err error)
	UpdatePaymentChannel     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *PaymentChannel, err error)
	DeletePaymentChannel     func(ctx context.Context, r *GeneratedResolver, id string) (item *PaymentChannel, err error)
	DeleteAllPaymentChannels func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryPaymentChannel      func(ctx context.Context, r *GeneratedResolver, opts QueryPaymentChannelHandlerOptions) (*PaymentChannel, error)
	QueryPaymentChannels     func(ctx context.Context, r *GeneratedResolver, opts QueryPaymentChannelsHandlerOptions) (*PaymentChannelResultType, error)

	CreatePaymentStatus      func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *PaymentStatus, err error)
	UpdatePaymentStatus      func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *PaymentStatus, err error)
	DeletePaymentStatus      func(ctx context.Context, r *GeneratedResolver, id string) (item *PaymentStatus, err error)
	DeleteAllPaymentStatuses func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryPaymentStatus       func(ctx context.Context, r *GeneratedResolver, opts QueryPaymentStatusHandlerOptions) (*PaymentStatus, error)
	QueryPaymentStatuses     func(ctx context.Context, r *GeneratedResolver, opts QueryPaymentStatusesHandlerOptions) (*PaymentStatusResultType, error)

	PaymentStatusPerson func(ctx context.Context, r *GeneratedResolver, obj *PaymentStatus) (res *Person, err error)

	CreatePaymentHistory      func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *PaymentHistory, err error)
	UpdatePaymentHistory      func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *PaymentHistory, err error)
	DeletePaymentHistory      func(ctx context.Context, r *GeneratedResolver, id string) (item *PaymentHistory, err error)
	DeleteAllPaymentHistories func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryPaymentHistory       func(ctx context.Context, r *GeneratedResolver, opts QueryPaymentHistoryHandlerOptions) (*PaymentHistory, error)
	QueryPaymentHistories     func(ctx context.Context, r *GeneratedResolver, opts QueryPaymentHistoriesHandlerOptions) (*PaymentHistoryResultType, error)

	PaymentHistoryPaymentChannel func(ctx context.Context, r *GeneratedResolver, obj *PaymentHistory) (res *PaymentChannel, err error)

	PaymentHistoryPerson func(ctx context.Context, r *GeneratedResolver, obj *PaymentHistory) (res *Person, err error)
}

// DefaultResolutionHandlers ...
func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *events.Event) error { return nil },

		CreateDelivery:      CreateDeliveryHandler,
		UpdateDelivery:      UpdateDeliveryHandler,
		DeleteDelivery:      DeleteDeliveryHandler,
		DeleteAllDeliveries: DeleteAllDeliveriesHandler,
		QueryDelivery:       QueryDeliveryHandler,
		QueryDeliveries:     QueryDeliveriesHandler,

		DeliveryVehicleType: DeliveryVehicleTypeHandler,

		DeliveryPaymentChannel: DeliveryPaymentChannelHandler,

		DeliveryDeliveryType: DeliveryDeliveryTypeHandler,

		DeliveryDeliveryChannel: DeliveryDeliveryChannelHandler,

		DeliverySender: DeliverySenderHandler,

		DeliveryReceiver: DeliveryReceiverHandler,

		DeliveryDeliver: DeliveryDeliverHandler,

		CreatePerson:    CreatePersonHandler,
		UpdatePerson:    UpdatePersonHandler,
		DeletePerson:    DeletePersonHandler,
		DeleteAllPeople: DeleteAllPeopleHandler,
		QueryPerson:     QueryPersonHandler,
		QueryPeople:     QueryPeopleHandler,

		PersonDeliveries: PersonDeliveriesHandler,

		PersonDeliveriesSent: PersonDeliveriesSentHandler,

		PersonDeliveriesReceived: PersonDeliveriesReceivedHandler,

		PersonPaymentStatus: PersonPaymentStatusHandler,

		PersonPaymentHistory: PersonPaymentHistoryHandler,

		CreateDeliveryType:     CreateDeliveryTypeHandler,
		UpdateDeliveryType:     UpdateDeliveryTypeHandler,
		DeleteDeliveryType:     DeleteDeliveryTypeHandler,
		DeleteAllDeliveryTypes: DeleteAllDeliveryTypesHandler,
		QueryDeliveryType:      QueryDeliveryTypeHandler,
		QueryDeliveryTypes:     QueryDeliveryTypesHandler,

		CreateDeliveryChannel:     CreateDeliveryChannelHandler,
		UpdateDeliveryChannel:     UpdateDeliveryChannelHandler,
		DeleteDeliveryChannel:     DeleteDeliveryChannelHandler,
		DeleteAllDeliveryChannels: DeleteAllDeliveryChannelsHandler,
		QueryDeliveryChannel:      QueryDeliveryChannelHandler,
		QueryDeliveryChannels:     QueryDeliveryChannelsHandler,

		CreateVehicleType:     CreateVehicleTypeHandler,
		UpdateVehicleType:     UpdateVehicleTypeHandler,
		DeleteVehicleType:     DeleteVehicleTypeHandler,
		DeleteAllVehicleTypes: DeleteAllVehicleTypesHandler,
		QueryVehicleType:      QueryVehicleTypeHandler,
		QueryVehicleTypes:     QueryVehicleTypesHandler,

		CreatePaymentChannel:     CreatePaymentChannelHandler,
		UpdatePaymentChannel:     UpdatePaymentChannelHandler,
		DeletePaymentChannel:     DeletePaymentChannelHandler,
		DeleteAllPaymentChannels: DeleteAllPaymentChannelsHandler,
		QueryPaymentChannel:      QueryPaymentChannelHandler,
		QueryPaymentChannels:     QueryPaymentChannelsHandler,

		CreatePaymentStatus:      CreatePaymentStatusHandler,
		UpdatePaymentStatus:      UpdatePaymentStatusHandler,
		DeletePaymentStatus:      DeletePaymentStatusHandler,
		DeleteAllPaymentStatuses: DeleteAllPaymentStatusesHandler,
		QueryPaymentStatus:       QueryPaymentStatusHandler,
		QueryPaymentStatuses:     QueryPaymentStatusesHandler,

		PaymentStatusPerson: PaymentStatusPersonHandler,

		CreatePaymentHistory:      CreatePaymentHistoryHandler,
		UpdatePaymentHistory:      UpdatePaymentHistoryHandler,
		DeletePaymentHistory:      DeletePaymentHistoryHandler,
		DeleteAllPaymentHistories: DeleteAllPaymentHistoriesHandler,
		QueryPaymentHistory:       QueryPaymentHistoryHandler,
		QueryPaymentHistories:     QueryPaymentHistoriesHandler,

		PaymentHistoryPaymentChannel: PaymentHistoryPaymentChannelHandler,

		PaymentHistoryPerson: PaymentHistoryPersonHandler,
	}
	return handlers
}

// GeneratedResolver struct
type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *EventController
}

// GetDB returns database connection or transaction for given context (if exists)
func (r *GeneratedResolver) GetDB(ctx context.Context) *gorm.DB {
	db, _ := ctx.Value(KeyMutationTransaction).(*gorm.DB)
	if db == nil {
		db = r.DB.Query()
	}
	return db
}
