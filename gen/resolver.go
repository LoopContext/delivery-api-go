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

	DeliveryInstructions func(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *string, err error)

	DeliverySender func(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *Person, err error)

	DeliveryReceiver func(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *Person, err error)

	DeliveryDeliver func(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *Person, err error)

	DeliveryVehicleType func(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *VehicleType, err error)

	DeliveryDeliveryType func(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *DeliveryType, err error)

	DeliveryDeliveryChannel func(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *DeliveryChannel, err error)

	CreatePerson    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Person, err error)
	UpdatePerson    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Person, err error)
	DeletePerson    func(ctx context.Context, r *GeneratedResolver, id string) (item *Person, err error)
	DeleteAllPeople func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryPerson     func(ctx context.Context, r *GeneratedResolver, opts QueryPersonHandlerOptions) (*Person, error)
	QueryPeople     func(ctx context.Context, r *GeneratedResolver, opts QueryPeopleHandlerOptions) (*PersonResultType, error)

	PersonDeliveries func(ctx context.Context, r *GeneratedResolver, obj *Person) (res []*Delivery, err error)

	PersonDeliveriesSent func(ctx context.Context, r *GeneratedResolver, obj *Person) (res []*Delivery, err error)

	PersonDeliveriesReceived func(ctx context.Context, r *GeneratedResolver, obj *Person) (res []*Delivery, err error)

	CreateDeliveryType     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *DeliveryType, err error)
	UpdateDeliveryType     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *DeliveryType, err error)
	DeleteDeliveryType     func(ctx context.Context, r *GeneratedResolver, id string) (item *DeliveryType, err error)
	DeleteAllDeliveryTypes func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryDeliveryType      func(ctx context.Context, r *GeneratedResolver, opts QueryDeliveryTypeHandlerOptions) (*DeliveryType, error)
	QueryDeliveryTypes     func(ctx context.Context, r *GeneratedResolver, opts QueryDeliveryTypesHandlerOptions) (*DeliveryTypeResultType, error)

	DeliveryTypeDelivery func(ctx context.Context, r *GeneratedResolver, obj *DeliveryType) (res *Delivery, err error)

	CreateDeliveryChannel     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *DeliveryChannel, err error)
	UpdateDeliveryChannel     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *DeliveryChannel, err error)
	DeleteDeliveryChannel     func(ctx context.Context, r *GeneratedResolver, id string) (item *DeliveryChannel, err error)
	DeleteAllDeliveryChannels func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryDeliveryChannel      func(ctx context.Context, r *GeneratedResolver, opts QueryDeliveryChannelHandlerOptions) (*DeliveryChannel, error)
	QueryDeliveryChannels     func(ctx context.Context, r *GeneratedResolver, opts QueryDeliveryChannelsHandlerOptions) (*DeliveryChannelResultType, error)

	DeliveryChannelDelivery func(ctx context.Context, r *GeneratedResolver, obj *DeliveryChannel) (res *Delivery, err error)

	CreateVehicleType     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *VehicleType, err error)
	UpdateVehicleType     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *VehicleType, err error)
	DeleteVehicleType     func(ctx context.Context, r *GeneratedResolver, id string) (item *VehicleType, err error)
	DeleteAllVehicleTypes func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryVehicleType      func(ctx context.Context, r *GeneratedResolver, opts QueryVehicleTypeHandlerOptions) (*VehicleType, error)
	QueryVehicleTypes     func(ctx context.Context, r *GeneratedResolver, opts QueryVehicleTypesHandlerOptions) (*VehicleTypeResultType, error)

	VehicleTypeDelivery func(ctx context.Context, r *GeneratedResolver, obj *VehicleType) (res *Delivery, err error)
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

		DeliveryInstructions: DeliveryInstructionsHandler,

		DeliverySender: DeliverySenderHandler,

		DeliveryReceiver: DeliveryReceiverHandler,

		DeliveryDeliver: DeliveryDeliverHandler,

		DeliveryVehicleType: DeliveryVehicleTypeHandler,

		DeliveryDeliveryType: DeliveryDeliveryTypeHandler,

		DeliveryDeliveryChannel: DeliveryDeliveryChannelHandler,

		CreatePerson:    CreatePersonHandler,
		UpdatePerson:    UpdatePersonHandler,
		DeletePerson:    DeletePersonHandler,
		DeleteAllPeople: DeleteAllPeopleHandler,
		QueryPerson:     QueryPersonHandler,
		QueryPeople:     QueryPeopleHandler,

		PersonDeliveries: PersonDeliveriesHandler,

		PersonDeliveriesSent: PersonDeliveriesSentHandler,

		PersonDeliveriesReceived: PersonDeliveriesReceivedHandler,

		CreateDeliveryType:     CreateDeliveryTypeHandler,
		UpdateDeliveryType:     UpdateDeliveryTypeHandler,
		DeleteDeliveryType:     DeleteDeliveryTypeHandler,
		DeleteAllDeliveryTypes: DeleteAllDeliveryTypesHandler,
		QueryDeliveryType:      QueryDeliveryTypeHandler,
		QueryDeliveryTypes:     QueryDeliveryTypesHandler,

		DeliveryTypeDelivery: DeliveryTypeDeliveryHandler,

		CreateDeliveryChannel:     CreateDeliveryChannelHandler,
		UpdateDeliveryChannel:     UpdateDeliveryChannelHandler,
		DeleteDeliveryChannel:     DeleteDeliveryChannelHandler,
		DeleteAllDeliveryChannels: DeleteAllDeliveryChannelsHandler,
		QueryDeliveryChannel:      QueryDeliveryChannelHandler,
		QueryDeliveryChannels:     QueryDeliveryChannelsHandler,

		DeliveryChannelDelivery: DeliveryChannelDeliveryHandler,

		CreateVehicleType:     CreateVehicleTypeHandler,
		UpdateVehicleType:     UpdateVehicleTypeHandler,
		DeleteVehicleType:     DeleteVehicleTypeHandler,
		DeleteAllVehicleTypes: DeleteAllVehicleTypesHandler,
		QueryVehicleType:      QueryVehicleTypeHandler,
		QueryVehicleTypes:     QueryVehicleTypesHandler,

		VehicleTypeDelivery: VehicleTypeDeliveryHandler,
	}
	return handlers
}

// GeneratedResolver struct
type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	db              *DB
	EventController *EventController
}

func NewGeneratedResolver(handlers ResolutionHandlers, db *DB, ec *EventController) *GeneratedResolver {
	return &GeneratedResolver{Handlers: handlers, db: db, EventController: ec}
}

// GetDB returns database connection or transaction for given context (if exists)
func (r *GeneratedResolver) GetDB(ctx context.Context) *gorm.DB {
	db, _ := ctx.Value(KeyMutationTransaction).(*gorm.DB)
	if db == nil {
		db = r.db.Query()
	}
	return db
}
