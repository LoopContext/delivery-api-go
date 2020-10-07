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

	DeliveryPaymentForm func(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *PaymentForm, err error)

	DeliverySender func(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *Person, err error)

	DeliveryReceiver func(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *Person, err error)

	DeliveryDeliver func(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *Deliver, err error)

	CreateVehicleType     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *VehicleType, err error)
	UpdateVehicleType     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *VehicleType, err error)
	DeleteVehicleType     func(ctx context.Context, r *GeneratedResolver, id string) (item *VehicleType, err error)
	DeleteAllVehicleTypes func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryVehicleType      func(ctx context.Context, r *GeneratedResolver, opts QueryVehicleTypeHandlerOptions) (*VehicleType, error)
	QueryVehicleTypes     func(ctx context.Context, r *GeneratedResolver, opts QueryVehicleTypesHandlerOptions) (*VehicleTypeResultType, error)

	CreatePaymentForm     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *PaymentForm, err error)
	UpdatePaymentForm     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *PaymentForm, err error)
	DeletePaymentForm     func(ctx context.Context, r *GeneratedResolver, id string) (item *PaymentForm, err error)
	DeleteAllPaymentForms func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryPaymentForm      func(ctx context.Context, r *GeneratedResolver, opts QueryPaymentFormHandlerOptions) (*PaymentForm, error)
	QueryPaymentForms     func(ctx context.Context, r *GeneratedResolver, opts QueryPaymentFormsHandlerOptions) (*PaymentFormResultType, error)

	CreateDeliver     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Deliver, err error)
	UpdateDeliver     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Deliver, err error)
	DeleteDeliver     func(ctx context.Context, r *GeneratedResolver, id string) (item *Deliver, err error)
	DeleteAllDelivers func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryDeliver      func(ctx context.Context, r *GeneratedResolver, opts QueryDeliverHandlerOptions) (*Deliver, error)
	QueryDelivers     func(ctx context.Context, r *GeneratedResolver, opts QueryDeliversHandlerOptions) (*DeliverResultType, error)

	DeliverDeliveries func(ctx context.Context, r *GeneratedResolver, obj *Deliver) (res []*Delivery, err error)

	CreatePerson    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Person, err error)
	UpdatePerson    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Person, err error)
	DeletePerson    func(ctx context.Context, r *GeneratedResolver, id string) (item *Person, err error)
	DeleteAllPeople func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryPerson     func(ctx context.Context, r *GeneratedResolver, opts QueryPersonHandlerOptions) (*Person, error)
	QueryPeople     func(ctx context.Context, r *GeneratedResolver, opts QueryPeopleHandlerOptions) (*PersonResultType, error)

	PersonDeliveriesSent func(ctx context.Context, r *GeneratedResolver, obj *Person) (res *Delivery, err error)

	PersonDeliveriesReceived func(ctx context.Context, r *GeneratedResolver, obj *Person) (res *Delivery, err error)
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

		DeliveryPaymentForm: DeliveryPaymentFormHandler,

		DeliverySender: DeliverySenderHandler,

		DeliveryReceiver: DeliveryReceiverHandler,

		DeliveryDeliver: DeliveryDeliverHandler,

		CreateVehicleType:     CreateVehicleTypeHandler,
		UpdateVehicleType:     UpdateVehicleTypeHandler,
		DeleteVehicleType:     DeleteVehicleTypeHandler,
		DeleteAllVehicleTypes: DeleteAllVehicleTypesHandler,
		QueryVehicleType:      QueryVehicleTypeHandler,
		QueryVehicleTypes:     QueryVehicleTypesHandler,

		CreatePaymentForm:     CreatePaymentFormHandler,
		UpdatePaymentForm:     UpdatePaymentFormHandler,
		DeletePaymentForm:     DeletePaymentFormHandler,
		DeleteAllPaymentForms: DeleteAllPaymentFormsHandler,
		QueryPaymentForm:      QueryPaymentFormHandler,
		QueryPaymentForms:     QueryPaymentFormsHandler,

		CreateDeliver:     CreateDeliverHandler,
		UpdateDeliver:     UpdateDeliverHandler,
		DeleteDeliver:     DeleteDeliverHandler,
		DeleteAllDelivers: DeleteAllDeliversHandler,
		QueryDeliver:      QueryDeliverHandler,
		QueryDelivers:     QueryDeliversHandler,

		DeliverDeliveries: DeliverDeliveriesHandler,

		CreatePerson:    CreatePersonHandler,
		UpdatePerson:    UpdatePersonHandler,
		DeletePerson:    DeletePersonHandler,
		DeleteAllPeople: DeleteAllPeopleHandler,
		QueryPerson:     QueryPersonHandler,
		QueryPeople:     QueryPeopleHandler,

		PersonDeliveriesSent: PersonDeliveriesSentHandler,

		PersonDeliveriesReceived: PersonDeliveriesReceivedHandler,
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
