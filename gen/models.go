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
	ID               string     `json:"id" gorm:"column:id;primary_key"`
	Mode             *string    `json:"mode" gorm:"column:mode"`
	CollectDateTime  *time.Time `json:"collectDateTime" gorm:"column:collectDateTime"`
	CollectAddress   *string    `json:"collectAddress" gorm:"column:collectAddress"`
	CollectPoint     *string    `json:"collectPoint" gorm:"column:collectPoint"`
	DropDateTime     *time.Time `json:"dropDateTime" gorm:"column:dropDateTime"`
	DropAddress      *string    `json:"dropAddress" gorm:"column:dropAddress"`
	DropPoint        *string    `json:"dropPoint" gorm:"column:dropPoint"`
	PaymentTotal     *float64   `json:"paymentTotal" gorm:"column:paymentTotal"`
	PaymentOnDeliver *bool      `json:"paymentOnDeliver" gorm:"column:paymentOnDeliver"`
	ExpectedDistance *string    `json:"expectedDistance" gorm:"column:expectedDistance"`
	ExpectedCost     *float64   `json:"expectedCost" gorm:"column:expectedCost"`
	Status           *string    `json:"status" gorm:"column:status"`
	Completed        *bool      `json:"completed" gorm:"column:completed"`
	SmsToken         *string    `json:"smsToken" gorm:"column:smsToken"`
	Instructions     *string    `json:"instructions" gorm:"column:instructions;type:text"`
	SenderID         *string    `json:"senderId" gorm:"column:senderId"`
	ReceiverID       *string    `json:"receiverId" gorm:"column:receiverId"`
	DeliverID        *string    `json:"deliverId" gorm:"column:deliverId"`
	UpdatedAt        *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt        time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy        *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy        *string    `json:"createdBy" gorm:"column:createdBy"`

	Sender *Person `json:"sender"`

	Receiver *Person `json:"receiver"`

	Deliver *Deliver `json:"deliver"`
}

// IsEntity ...
func (m *Delivery) IsEntity() {}

// DeliveryChanges struct
type DeliveryChanges struct {
	ID               string
	Mode             *string
	CollectDateTime  *time.Time
	CollectAddress   *string
	CollectPoint     *string
	DropDateTime     *time.Time
	DropAddress      *string
	DropPoint        *string
	PaymentTotal     *float64
	PaymentOnDeliver *bool
	ExpectedDistance *string
	ExpectedCost     *float64
	Status           *string
	Completed        *bool
	SmsToken         *string
	Instructions     *string
	SenderID         *string
	ReceiverID       *string
	DeliverID        *string
	UpdatedAt        *time.Time
	CreatedAt        time.Time
	UpdatedBy        *string
	CreatedBy        *string
}

// VehicleTypeResultType struct
type VehicleTypeResultType struct {
	EntityResultType
}

// VehicleType struct
type VehicleType struct {
	ID          string     `json:"id" gorm:"column:id;primary_key"`
	Name        *string    `json:"name" gorm:"column:name"`
	Description *string    `json:"description" gorm:"column:description;type:text"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy   *string    `json:"createdBy" gorm:"column:createdBy"`
}

// IsEntity ...
func (m *VehicleType) IsEntity() {}

// VehicleTypeChanges struct
type VehicleTypeChanges struct {
	ID          string
	Name        *string
	Description *string
	UpdatedAt   *time.Time
	CreatedAt   time.Time
	UpdatedBy   *string
	CreatedBy   *string
}

// PaymentFormResultType struct
type PaymentFormResultType struct {
	EntityResultType
}

// PaymentForm struct
type PaymentForm struct {
	ID          string     `json:"id" gorm:"column:id;primary_key"`
	Name        *string    `json:"name" gorm:"column:name"`
	Description *string    `json:"description" gorm:"column:description;type:text"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy   *string    `json:"createdBy" gorm:"column:createdBy"`
}

// IsEntity ...
func (m *PaymentForm) IsEntity() {}

// PaymentFormChanges struct
type PaymentFormChanges struct {
	ID          string
	Name        *string
	Description *string
	UpdatedAt   *time.Time
	CreatedAt   time.Time
	UpdatedBy   *string
	CreatedBy   *string
}

// DeliverResultType struct
type DeliverResultType struct {
	EntityResultType
}

// Deliver struct
type Deliver struct {
	ID          string     `json:"id" gorm:"column:id;primary_key"`
	Email       string     `json:"email" gorm:"column:email;unique"`
	Phone       *string    `json:"phone" gorm:"column:phone"`
	AvatarURL   *string    `json:"avatarURL" gorm:"column:avatarURL;type:text"`
	DisplayName *string    `json:"displayName" gorm:"column:displayName"`
	FirstName   *string    `json:"firstName" gorm:"column:firstName"`
	LastName    *string    `json:"lastName" gorm:"column:lastName"`
	NickName    *string    `json:"nickName" gorm:"column:nickName"`
	Description *string    `json:"description" gorm:"column:description;type:text"`
	Location    *string    `json:"location" gorm:"column:location"`
	UserID      *string    `json:"userId" gorm:"column:userId"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy   *string    `json:"createdBy" gorm:"column:createdBy"`

	Deliveries []*Delivery `json:"deliveries" gorm:"foreignkey:DeliverID"`
}

// IsEntity ...
func (m *Deliver) IsEntity() {}

// DeliverChanges struct
type DeliverChanges struct {
	ID          string
	Email       string
	Phone       *string
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

	DeliveriesIDs []*string
}

// PersonResultType struct
type PersonResultType struct {
	EntityResultType
}

// Person struct
type Person struct {
	ID                   string     `json:"id" gorm:"column:id;primary_key"`
	Name                 *string    `json:"name" gorm:"column:name"`
	Phone                *string    `json:"phone" gorm:"column:phone"`
	Email                string     `json:"email" gorm:"column:email;unique"`
	DocumentNo           *string    `json:"documentNo" gorm:"column:documentNo"`
	UserID               *string    `json:"userId" gorm:"column:userId"`
	DeliveriesSentID     *string    `json:"deliveriesSentId" gorm:"column:deliveriesSentId"`
	DeliveriesReceivedID *string    `json:"deliveriesReceivedId" gorm:"column:deliveriesReceivedId"`
	UpdatedAt            *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt            time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy            *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy            *string    `json:"createdBy" gorm:"column:createdBy"`

	DeliveriesSent *Delivery `json:"deliveriesSent"`

	DeliveriesReceived *Delivery `json:"deliveriesReceived"`
}

// IsEntity ...
func (m *Person) IsEntity() {}

// PersonChanges struct
type PersonChanges struct {
	ID                   string
	Name                 *string
	Phone                *string
	Email                string
	DocumentNo           *string
	UserID               *string
	DeliveriesSentID     *string
	DeliveriesReceivedID *string
	UpdatedAt            *time.Time
	CreatedAt            time.Time
	UpdatedBy            *string
	CreatedBy            *string
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
