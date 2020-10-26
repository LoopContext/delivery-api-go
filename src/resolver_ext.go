package src

import (
	"context"
	"fmt"

	"github.com/loopcontext/deliver-api-go/gen"
)

const (
	jwtTokenPermissionErrMsg = "You don't have permission to %s the entity %s"
)

// Deliveries method
func (r *QueryResolver) Deliveries(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.DeliverySortType, filter *gen.DeliveryFilterType) (*gen.DeliveryResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "deliveries", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "deliveries")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Deliveries to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Deliveries(ctx, offset, limit, q, sort, filter)
}

// CreateDelivery method
func (r *MutationResolver) CreateDelivery(ctx context.Context, input map[string]interface{}) (item *gen.Delivery, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "deliveries", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "deliveries")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Deliveries to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateDelivery(ctx, input)
}

// ReadDelivery method
func (r *QueryResolver) Delivery(ctx context.Context, id *string, q *string, filter *gen.DeliveryFilterType) (*gen.Delivery, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "deliveries", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "deliveries")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Deliveries to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Delivery(ctx, id, q, filter)
}

// UpdateDelivery method
func (r *MutationResolver) UpdateDelivery(ctx context.Context, id string, input map[string]interface{}) (item *gen.Delivery, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "deliveries", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "deliveries")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Deliveries to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateDelivery(ctx, id, input)
}

// DeleteDelivery method
func (r *MutationResolver) DeleteDelivery(ctx context.Context, id string) (item *gen.Delivery, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "deliveries", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "deliveries")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Deliveries to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteDelivery(ctx, id)
}

// DeleteAllDeliveries method
func (r *MutationResolver) DeleteAllDeliveries(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "deliveries", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "deliveries")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Deliveries to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllDeliveries(ctx)
}

// People method
func (r *QueryResolver) People(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.PersonSortType, filter *gen.PersonFilterType) (*gen.PersonResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "people", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "people")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope People to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.People(ctx, offset, limit, q, sort, filter)
}

// CreatePerson method
func (r *MutationResolver) CreatePerson(ctx context.Context, input map[string]interface{}) (item *gen.Person, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "people", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "people")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope People to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreatePerson(ctx, input)
}

// ReadPerson method
func (r *QueryResolver) Person(ctx context.Context, id *string, q *string, filter *gen.PersonFilterType) (*gen.Person, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "people", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "people")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope People to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Person(ctx, id, q, filter)
}

// UpdatePerson method
func (r *MutationResolver) UpdatePerson(ctx context.Context, id string, input map[string]interface{}) (item *gen.Person, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "people", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "people")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope People to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdatePerson(ctx, id, input)
}

// DeletePerson method
func (r *MutationResolver) DeletePerson(ctx context.Context, id string) (item *gen.Person, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "people", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "people")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope People to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeletePerson(ctx, id)
}

// DeleteAllPeople method
func (r *MutationResolver) DeleteAllPeople(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "people", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "people")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope People to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllPeople(ctx)
}

// DeliveryTypes method
func (r *QueryResolver) DeliveryTypes(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.DeliveryTypeSortType, filter *gen.DeliveryTypeFilterType) (*gen.DeliveryTypeResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "delivery_types", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "delivery_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope DeliveryTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.DeliveryTypes(ctx, offset, limit, q, sort, filter)
}

// CreateDeliveryType method
func (r *MutationResolver) CreateDeliveryType(ctx context.Context, input map[string]interface{}) (item *gen.DeliveryType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "delivery_types", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "delivery_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope DeliveryTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateDeliveryType(ctx, input)
}

// ReadDeliveryType method
func (r *QueryResolver) DeliveryType(ctx context.Context, id *string, q *string, filter *gen.DeliveryTypeFilterType) (*gen.DeliveryType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "delivery_types", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "delivery_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope DeliveryTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.DeliveryType(ctx, id, q, filter)
}

// UpdateDeliveryType method
func (r *MutationResolver) UpdateDeliveryType(ctx context.Context, id string, input map[string]interface{}) (item *gen.DeliveryType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "delivery_types", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "delivery_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope DeliveryTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateDeliveryType(ctx, id, input)
}

// DeleteDeliveryType method
func (r *MutationResolver) DeleteDeliveryType(ctx context.Context, id string) (item *gen.DeliveryType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "delivery_types", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "delivery_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope DeliveryTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteDeliveryType(ctx, id)
}

// DeleteAllDeliveryTypes method
func (r *MutationResolver) DeleteAllDeliveryTypes(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "delivery_types", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "delivery_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope DeliveryTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllDeliveryTypes(ctx)
}

// DeliveryChannels method
func (r *QueryResolver) DeliveryChannels(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.DeliveryChannelSortType, filter *gen.DeliveryChannelFilterType) (*gen.DeliveryChannelResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "delivery_channels", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "delivery_channels")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope DeliveryChannels to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.DeliveryChannels(ctx, offset, limit, q, sort, filter)
}

// CreateDeliveryChannel method
func (r *MutationResolver) CreateDeliveryChannel(ctx context.Context, input map[string]interface{}) (item *gen.DeliveryChannel, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "delivery_channels", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "delivery_channels")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope DeliveryChannels to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateDeliveryChannel(ctx, input)
}

// ReadDeliveryChannel method
func (r *QueryResolver) DeliveryChannel(ctx context.Context, id *string, q *string, filter *gen.DeliveryChannelFilterType) (*gen.DeliveryChannel, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "delivery_channels", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "delivery_channels")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope DeliveryChannels to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.DeliveryChannel(ctx, id, q, filter)
}

// UpdateDeliveryChannel method
func (r *MutationResolver) UpdateDeliveryChannel(ctx context.Context, id string, input map[string]interface{}) (item *gen.DeliveryChannel, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "delivery_channels", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "delivery_channels")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope DeliveryChannels to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateDeliveryChannel(ctx, id, input)
}

// DeleteDeliveryChannel method
func (r *MutationResolver) DeleteDeliveryChannel(ctx context.Context, id string) (item *gen.DeliveryChannel, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "delivery_channels", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "delivery_channels")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope DeliveryChannels to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteDeliveryChannel(ctx, id)
}

// DeleteAllDeliveryChannels method
func (r *MutationResolver) DeleteAllDeliveryChannels(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "delivery_channels", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "delivery_channels")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope DeliveryChannels to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllDeliveryChannels(ctx)
}

// VehicleTypes method
func (r *QueryResolver) VehicleTypes(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.VehicleTypeSortType, filter *gen.VehicleTypeFilterType) (*gen.VehicleTypeResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "vehicle_types", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "vehicle_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope VehicleTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.VehicleTypes(ctx, offset, limit, q, sort, filter)
}

// CreateVehicleType method
func (r *MutationResolver) CreateVehicleType(ctx context.Context, input map[string]interface{}) (item *gen.VehicleType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "vehicle_types", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "vehicle_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope VehicleTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateVehicleType(ctx, input)
}

// ReadVehicleType method
func (r *QueryResolver) VehicleType(ctx context.Context, id *string, q *string, filter *gen.VehicleTypeFilterType) (*gen.VehicleType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "vehicle_types", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "vehicle_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope VehicleTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.VehicleType(ctx, id, q, filter)
}

// UpdateVehicleType method
func (r *MutationResolver) UpdateVehicleType(ctx context.Context, id string, input map[string]interface{}) (item *gen.VehicleType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "vehicle_types", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "vehicle_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope VehicleTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateVehicleType(ctx, id, input)
}

// DeleteVehicleType method
func (r *MutationResolver) DeleteVehicleType(ctx context.Context, id string) (item *gen.VehicleType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "vehicle_types", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "vehicle_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope VehicleTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteVehicleType(ctx, id)
}

// DeleteAllVehicleTypes method
func (r *MutationResolver) DeleteAllVehicleTypes(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "vehicle_types", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "vehicle_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope VehicleTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllVehicleTypes(ctx)
}

// PaymentForms method
func (r *QueryResolver) PaymentForms(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.PaymentFormSortType, filter *gen.PaymentFormFilterType) (*gen.PaymentFormResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_forms", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "payment_forms")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentForms to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.PaymentForms(ctx, offset, limit, q, sort, filter)
}

// CreatePaymentForm method
func (r *MutationResolver) CreatePaymentForm(ctx context.Context, input map[string]interface{}) (item *gen.PaymentForm, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_forms", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "payment_forms")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentForms to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreatePaymentForm(ctx, input)
}

// ReadPaymentForm method
func (r *QueryResolver) PaymentForm(ctx context.Context, id *string, q *string, filter *gen.PaymentFormFilterType) (*gen.PaymentForm, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_forms", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "payment_forms")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentForms to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.PaymentForm(ctx, id, q, filter)
}

// UpdatePaymentForm method
func (r *MutationResolver) UpdatePaymentForm(ctx context.Context, id string, input map[string]interface{}) (item *gen.PaymentForm, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_forms", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "payment_forms")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentForms to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdatePaymentForm(ctx, id, input)
}

// DeletePaymentForm method
func (r *MutationResolver) DeletePaymentForm(ctx context.Context, id string) (item *gen.PaymentForm, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_forms", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payment_forms")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentForms to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeletePaymentForm(ctx, id)
}

// DeleteAllPaymentForms method
func (r *MutationResolver) DeleteAllPaymentForms(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "payment_forms", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payment_forms")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentForms to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllPaymentForms(ctx)
}

// PaymentStatuses method
func (r *QueryResolver) PaymentStatuses(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.PaymentStatusSortType, filter *gen.PaymentStatusFilterType) (*gen.PaymentStatusResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_statuses", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "payment_statuses")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentStatuses to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.PaymentStatuses(ctx, offset, limit, q, sort, filter)
}

// CreatePaymentStatus method
func (r *MutationResolver) CreatePaymentStatus(ctx context.Context, input map[string]interface{}) (item *gen.PaymentStatus, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_statuses", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "payment_statuses")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentStatuses to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreatePaymentStatus(ctx, input)
}

// ReadPaymentStatus method
func (r *QueryResolver) PaymentStatus(ctx context.Context, id *string, q *string, filter *gen.PaymentStatusFilterType) (*gen.PaymentStatus, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_statuses", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "payment_statuses")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentStatuses to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.PaymentStatus(ctx, id, q, filter)
}

// UpdatePaymentStatus method
func (r *MutationResolver) UpdatePaymentStatus(ctx context.Context, id string, input map[string]interface{}) (item *gen.PaymentStatus, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_statuses", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "payment_statuses")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentStatuses to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdatePaymentStatus(ctx, id, input)
}

// DeletePaymentStatus method
func (r *MutationResolver) DeletePaymentStatus(ctx context.Context, id string) (item *gen.PaymentStatus, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_statuses", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payment_statuses")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentStatuses to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeletePaymentStatus(ctx, id)
}

// DeleteAllPaymentStatuses method
func (r *MutationResolver) DeleteAllPaymentStatuses(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "payment_statuses", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payment_statuses")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentStatuses to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllPaymentStatuses(ctx)
}

// PaymentHistories method
func (r *QueryResolver) PaymentHistories(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.PaymentHistorySortType, filter *gen.PaymentHistoryFilterType) (*gen.PaymentHistoryResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_histories", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "payment_histories")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentHistories to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.PaymentHistories(ctx, offset, limit, q, sort, filter)
}

// CreatePaymentHistory method
func (r *MutationResolver) CreatePaymentHistory(ctx context.Context, input map[string]interface{}) (item *gen.PaymentHistory, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_histories", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "payment_histories")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentHistories to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreatePaymentHistory(ctx, input)
}

// ReadPaymentHistory method
func (r *QueryResolver) PaymentHistory(ctx context.Context, id *string, q *string, filter *gen.PaymentHistoryFilterType) (*gen.PaymentHistory, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_histories", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "payment_histories")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentHistories to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.PaymentHistory(ctx, id, q, filter)
}

// UpdatePaymentHistory method
func (r *MutationResolver) UpdatePaymentHistory(ctx context.Context, id string, input map[string]interface{}) (item *gen.PaymentHistory, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_histories", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "payment_histories")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentHistories to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdatePaymentHistory(ctx, id, input)
}

// DeletePaymentHistory method
func (r *MutationResolver) DeletePaymentHistory(ctx context.Context, id string) (item *gen.PaymentHistory, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_histories", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payment_histories")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentHistories to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeletePaymentHistory(ctx, id)
}

// DeleteAllPaymentHistories method
func (r *MutationResolver) DeleteAllPaymentHistories(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "payment_histories", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payment_histories")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentHistories to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllPaymentHistories(ctx)
}
