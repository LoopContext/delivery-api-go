package gen

import (
	"fmt"
	"reflect"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mitchellh/mapstructure"
)

// DeliveryResultType struct
type DeliveryResultType struct {
	EntityResultType
}

// Delivery struct
type Delivery struct {
	ID                string     `json:"id" gorm:"column:id;primary_key"`
	PaymentID         *string    `json:"paymentId" gorm:"column:paymentId;index:pidx"`
	PaymentTotal      *float64   `json:"paymentTotal" gorm:"column:paymentTotal"`
	PaymentOnDeliver  *bool      `json:"paymentOnDeliver" gorm:"column:paymentOnDeliver"`
	CollectDateTime   *time.Time `json:"collectDateTime" gorm:"column:collectDateTime"`
	CollectAddress    *string    `json:"collectAddress" gorm:"column:collectAddress"`
	CollectPoint      *string    `json:"collectPoint" gorm:"column:collectPoint"`
	DropDateTime      *time.Time `json:"dropDateTime" gorm:"column:dropDateTime"`
	DropAddress       *string    `json:"dropAddress" gorm:"column:dropAddress"`
	DropPoint         *string    `json:"dropPoint" gorm:"column:dropPoint"`
	ExpectedDistance  *string    `json:"expectedDistance" gorm:"column:expectedDistance"`
	ExpectedCost      *float64   `json:"expectedCost" gorm:"column:expectedCost"`
	Completed         *bool      `json:"completed" gorm:"column:completed"`
	SmsToken          *string    `json:"smsToken" gorm:"column:smsToken"`
	Status            *string    `json:"status" gorm:"column:status;index:statusidx"`
	SenderID          *string    `json:"senderId" gorm:"column:senderId"`
	ReceiverID        *string    `json:"receiverId" gorm:"column:receiverId"`
	DeliverID         *string    `json:"deliverId" gorm:"column:deliverId"`
	VehicleTypeID     *string    `json:"vehicleTypeId" gorm:"column:vehicleTypeId"`
	DeliveryTypeID    *string    `json:"deliveryTypeId" gorm:"column:deliveryTypeId"`
	DeliveryChannelID *string    `json:"deliveryChannelId" gorm:"column:deliveryChannelId"`
	UpdatedAt         *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt         time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy         *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy         *string    `json:"createdBy" gorm:"column:createdBy"`

	Sender          *Person `json:"sender"`
	SenderPreloaded bool    `gorm:"-"`

	Receiver          *Person `json:"receiver"`
	ReceiverPreloaded bool    `gorm:"-"`

	Deliver          *Person `json:"deliver"`
	DeliverPreloaded bool    `gorm:"-"`

	VehicleType          *VehicleType `json:"vehicleType"`
	VehicleTypePreloaded bool         `gorm:"-"`

	DeliveryType *DeliveryType `json:"deliveryType"`

	DeliveryChannel *DeliveryChannel `json:"deliveryChannel"`
}

// IsEntity ...
func (m *Delivery) IsEntity() {}

// DeliveryChanges struct
type DeliveryChanges struct {
	ID                string
	PaymentID         *string
	PaymentTotal      *float64
	PaymentOnDeliver  *bool
	CollectDateTime   *time.Time
	CollectAddress    *string
	CollectPoint      *string
	DropDateTime      *time.Time
	DropAddress       *string
	DropPoint         *string
	ExpectedDistance  *string
	ExpectedCost      *float64
	Completed         *bool
	SmsToken          *string
	Status            *string
	SenderID          *string
	ReceiverID        *string
	DeliverID         *string
	VehicleTypeID     *string
	DeliveryTypeID    *string
	DeliveryChannelID *string
	UpdatedAt         *time.Time
	CreatedAt         time.Time
	UpdatedBy         *string
	CreatedBy         *string
}

// PersonResultType struct
type PersonResultType struct {
	EntityResultType
}

// Person struct
type Person struct {
	ID          string     `json:"id" gorm:"column:id;primary_key"`
	Deliver     *bool      `json:"deliver" gorm:"column:deliver;default:false"`
	Email       string     `json:"email" gorm:"column:email;unique"`
	Phone       *string    `json:"phone" gorm:"column:phone"`
	DocumentNo  *string    `json:"documentNo" gorm:"column:documentNo"`
	AvatarURL   *string    `json:"avatarURL" gorm:"column:avatarURL"`
	DisplayName *string    `json:"displayName" gorm:"column:displayName"`
	FirstName   *string    `json:"firstName" gorm:"column:firstName"`
	LastName    *string    `json:"lastName" gorm:"column:lastName"`
	NickName    *string    `json:"nickName" gorm:"column:nickName"`
	Description *string    `json:"description" gorm:"column:description"`
	Location    *string    `json:"location" gorm:"column:location"`
	UserID      *string    `json:"userId" gorm:"column:userId;index:useridx"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy   *string    `json:"createdBy" gorm:"column:createdBy"`

	Deliveries []*Delivery `json:"deliveries" gorm:"foreignkey:DeliverID"`

	DeliveriesSent []*Delivery `json:"deliveriesSent" gorm:"foreignkey:SenderID"`

	DeliveriesReceived []*Delivery `json:"deliveriesReceived" gorm:"foreignkey:ReceiverID"`
}

// IsEntity ...
func (m *Person) IsEntity() {}

// PersonChanges struct
type PersonChanges struct {
	ID          string
	Deliver     *bool
	Email       string
	Phone       *string
	DocumentNo  *string
	AvatarURL   *string
	DisplayName *string
	FirstName   *string
	LastName    *string
	NickName    *string
	Description *string
	Location    *string
	UserID      *string
	UpdatedAt   *time.Time
	CreatedAt   time.Time
	UpdatedBy   *string
	CreatedBy   *string

	DeliveriesIDs         []*string
	DeliveriesSentIDs     []*string
	DeliveriesReceivedIDs []*string
}

// DeliveryTypeResultType struct
type DeliveryTypeResultType struct {
	EntityResultType
}

// DeliveryType struct
type DeliveryType struct {
	ID          string     `json:"id" gorm:"column:id;primary_key"`
	Name        *string    `json:"name" gorm:"column:name"`
	Description *string    `json:"description" gorm:"column:description"`
	DeliveryID  *string    `json:"deliveryId" gorm:"column:deliveryId"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy   *string    `json:"createdBy" gorm:"column:createdBy"`

	Delivery *Delivery `json:"delivery"`
}

// IsEntity ...
func (m *DeliveryType) IsEntity() {}

// DeliveryTypeChanges struct
type DeliveryTypeChanges struct {
	ID          string
	Name        *string
	Description *string
	DeliveryID  *string
	UpdatedAt   *time.Time
	CreatedAt   time.Time
	UpdatedBy   *string
	CreatedBy   *string
}

// DeliveryChannelResultType struct
type DeliveryChannelResultType struct {
	EntityResultType
}

// DeliveryChannel struct
type DeliveryChannel struct {
	ID          string     `json:"id" gorm:"column:id;primary_key"`
	Name        *string    `json:"name" gorm:"column:name"`
	Description *string    `json:"description" gorm:"column:description"`
	DeliveryID  *string    `json:"deliveryId" gorm:"column:deliveryId"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy   *string    `json:"createdBy" gorm:"column:createdBy"`

	Delivery *Delivery `json:"delivery"`
}

// IsEntity ...
func (m *DeliveryChannel) IsEntity() {}

// DeliveryChannelChanges struct
type DeliveryChannelChanges struct {
	ID          string
	Name        *string
	Description *string
	DeliveryID  *string
	UpdatedAt   *time.Time
	CreatedAt   time.Time
	UpdatedBy   *string
	CreatedBy   *string
}

// VehicleTypeResultType struct
type VehicleTypeResultType struct {
	EntityResultType
}

// VehicleType struct
type VehicleType struct {
	ID          string     `json:"id" gorm:"column:id;primary_key"`
	Name        *string    `json:"name" gorm:"column:name"`
	Description *string    `json:"description" gorm:"column:description"`
	DeliveryID  *string    `json:"deliveryId" gorm:"column:deliveryId"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy   *string    `json:"createdBy" gorm:"column:createdBy"`

	Delivery *Delivery `json:"delivery"`
}

// IsEntity ...
func (m *VehicleType) IsEntity() {}

// VehicleTypeChanges struct
type VehicleTypeChanges struct {
	ID          string
	Name        *string
	Description *string
	DeliveryID  *string
	UpdatedAt   *time.Time
	CreatedAt   time.Time
	UpdatedBy   *string
	CreatedBy   *string
}

// ApplyChanges used to convert map[string]interface{} to EntityChanges struct
func ApplyChanges(changes map[string]interface{}, to interface{}) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "json",
		Result:      to,
		ZeroFields:  true,
		// This is needed to get mapstructure to call the gqlgen unmarshaler func for custom scalars (eg Date)
		DecodeHook: func(a reflect.Type, b reflect.Type, v interface{}) (interface{}, error) {

			if b == reflect.TypeOf(time.Time{}) {
				switch a.Kind() {
				case reflect.String:
					return time.Parse(time.RFC3339, v.(string))
				case reflect.Float64:
					return time.Unix(0, int64(v.(float64))*int64(time.Millisecond)), nil
				case reflect.Int64:
					return time.Unix(0, v.(int64)*int64(time.Millisecond)), nil
				default:
					return v, fmt.Errorf("Unable to parse date from %v", v)
				}
			}

			if reflect.PtrTo(b).Implements(reflect.TypeOf((*graphql.Unmarshaler)(nil)).Elem()) {
				resultType := reflect.New(b)
				result := resultType.MethodByName("UnmarshalGQL").Call([]reflect.Value{reflect.ValueOf(v)})
				err, _ := result[0].Interface().(error)
				return resultType.Elem().Interface(), err
			}

			return v, nil
		},
	})

	if err != nil {
		return err
	}

	return dec.Decode(changes)
}
