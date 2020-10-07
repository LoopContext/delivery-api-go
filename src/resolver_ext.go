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
	if !gen.HasPermission(jwtClaims, "deliveries", gen.JWTPermissionConstList) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "deliveries")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Deliveries to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Deliveries(ctx, offset, limit, q, sort, filter)
}

// CreateDelivery method
func (r *MutationResolver) CreateDelivery(ctx context.Context, input map[string]interface{}) (item *gen.Delivery, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "deliveries", gen.JWTPermissionConstCreate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "deliveries")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Deliveries to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateDelivery(ctx, input)
}

// ReadDelivery method
func (r *QueryResolver) Delivery(ctx context.Context, id *string, q *string, filter *gen.DeliveryFilterType) (*gen.Delivery, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "deliveries", gen.JWTPermissionConstRead) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "deliveries")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Deliveries to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Delivery(ctx, id, q, filter)
}

// UpdateDelivery method
func (r *MutationResolver) UpdateDelivery(ctx context.Context, id string, input map[string]interface{}) (item *gen.Delivery, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "deliveries", gen.JWTPermissionConstUpdate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "deliveries")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Deliveries to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateDelivery(ctx, id, input)
}

// DeleteDelivery method
func (r *MutationResolver) DeleteDelivery(ctx context.Context, id string) (item *gen.Delivery, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "deliveries", gen.JWTPermissionConstDelete) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "deliveries")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Deliveries to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteDelivery(ctx, id)
}

// DeleteAllDeliveries method
func (r *MutationResolver) DeleteAllDeliveries(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "deliveries", gen.JWTPermissionConstDelete) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "deliveries")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Deliveries to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllDeliveries(ctx)
}

// VehicleTypes method
func (r *QueryResolver) VehicleTypes(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.VehicleTypeSortType, filter *gen.VehicleTypeFilterType) (*gen.VehicleTypeResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "vehicle_types", gen.JWTPermissionConstList) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "vehicle_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope VehicleTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.VehicleTypes(ctx, offset, limit, q, sort, filter)
}

// CreateVehicleType method
func (r *MutationResolver) CreateVehicleType(ctx context.Context, input map[string]interface{}) (item *gen.VehicleType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "vehicle_types", gen.JWTPermissionConstCreate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "vehicle_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope VehicleTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateVehicleType(ctx, input)
}

// ReadVehicleType method
func (r *QueryResolver) VehicleType(ctx context.Context, id *string, q *string, filter *gen.VehicleTypeFilterType) (*gen.VehicleType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "vehicle_types", gen.JWTPermissionConstRead) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "vehicle_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope VehicleTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.VehicleType(ctx, id, q, filter)
}

// UpdateVehicleType method
func (r *MutationResolver) UpdateVehicleType(ctx context.Context, id string, input map[string]interface{}) (item *gen.VehicleType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "vehicle_types", gen.JWTPermissionConstUpdate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "vehicle_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope VehicleTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateVehicleType(ctx, id, input)
}

// DeleteVehicleType method
func (r *MutationResolver) DeleteVehicleType(ctx context.Context, id string) (item *gen.VehicleType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "vehicle_types", gen.JWTPermissionConstDelete) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "vehicle_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope VehicleTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteVehicleType(ctx, id)
}

// DeleteAllVehicleTypes method
func (r *MutationResolver) DeleteAllVehicleTypes(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "vehicle_types", gen.JWTPermissionConstDelete) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "vehicle_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope VehicleTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllVehicleTypes(ctx)
}

// PaymentForms method
func (r *QueryResolver) PaymentForms(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.PaymentFormSortType, filter *gen.PaymentFormFilterType) (*gen.PaymentFormResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_forms", gen.JWTPermissionConstList) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "payment_forms")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentForms to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.PaymentForms(ctx, offset, limit, q, sort, filter)
}

// CreatePaymentForm method
func (r *MutationResolver) CreatePaymentForm(ctx context.Context, input map[string]interface{}) (item *gen.PaymentForm, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_forms", gen.JWTPermissionConstCreate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "payment_forms")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentForms to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreatePaymentForm(ctx, input)
}

// ReadPaymentForm method
func (r *QueryResolver) PaymentForm(ctx context.Context, id *string, q *string, filter *gen.PaymentFormFilterType) (*gen.PaymentForm, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_forms", gen.JWTPermissionConstRead) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "payment_forms")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentForms to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.PaymentForm(ctx, id, q, filter)
}

// UpdatePaymentForm method
func (r *MutationResolver) UpdatePaymentForm(ctx context.Context, id string, input map[string]interface{}) (item *gen.PaymentForm, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_forms", gen.JWTPermissionConstUpdate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "payment_forms")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentForms to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdatePaymentForm(ctx, id, input)
}

// DeletePaymentForm method
func (r *MutationResolver) DeletePaymentForm(ctx context.Context, id string) (item *gen.PaymentForm, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payment_forms", gen.JWTPermissionConstDelete) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payment_forms")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentForms to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeletePaymentForm(ctx, id)
}

// DeleteAllPaymentForms method
func (r *MutationResolver) DeleteAllPaymentForms(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "payment_forms", gen.JWTPermissionConstDelete) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payment_forms")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentForms to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllPaymentForms(ctx)
}

// Delivers method
func (r *QueryResolver) Delivers(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.DeliverSortType, filter *gen.DeliverFilterType) (*gen.DeliverResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "delivers", gen.JWTPermissionConstList) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "delivers")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Delivers to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Delivers(ctx, offset, limit, q, sort, filter)
}

// CreateDeliver method
func (r *MutationResolver) CreateDeliver(ctx context.Context, input map[string]interface{}) (item *gen.Deliver, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "delivers", gen.JWTPermissionConstCreate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "delivers")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Delivers to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateDeliver(ctx, input)
}

// ReadDeliver method
func (r *QueryResolver) Deliver(ctx context.Context, id *string, q *string, filter *gen.DeliverFilterType) (*gen.Deliver, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "delivers", gen.JWTPermissionConstRead) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "delivers")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Delivers to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Deliver(ctx, id, q, filter)
}

// UpdateDeliver method
func (r *MutationResolver) UpdateDeliver(ctx context.Context, id string, input map[string]interface{}) (item *gen.Deliver, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "delivers", gen.JWTPermissionConstUpdate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "delivers")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Delivers to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateDeliver(ctx, id, input)
}

// DeleteDeliver method
func (r *MutationResolver) DeleteDeliver(ctx context.Context, id string) (item *gen.Deliver, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "delivers", gen.JWTPermissionConstDelete) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "delivers")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Delivers to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteDeliver(ctx, id)
}

// DeleteAllDelivers method
func (r *MutationResolver) DeleteAllDelivers(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "delivers", gen.JWTPermissionConstDelete) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "delivers")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Delivers to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllDelivers(ctx)
}

// People method
func (r *QueryResolver) People(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.PersonSortType, filter *gen.PersonFilterType) (*gen.PersonResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "people", gen.JWTPermissionConstList) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "people")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope People to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.People(ctx, offset, limit, q, sort, filter)
}

// CreatePerson method
func (r *MutationResolver) CreatePerson(ctx context.Context, input map[string]interface{}) (item *gen.Person, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "people", gen.JWTPermissionConstCreate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "people")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope People to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreatePerson(ctx, input)
}

// ReadPerson method
func (r *QueryResolver) Person(ctx context.Context, id *string, q *string, filter *gen.PersonFilterType) (*gen.Person, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "people", gen.JWTPermissionConstRead) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "people")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope People to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Person(ctx, id, q, filter)
}

// UpdatePerson method
func (r *MutationResolver) UpdatePerson(ctx context.Context, id string, input map[string]interface{}) (item *gen.Person, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "people", gen.JWTPermissionConstUpdate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "people")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope People to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdatePerson(ctx, id, input)
}

// DeletePerson method
func (r *MutationResolver) DeletePerson(ctx context.Context, id string) (item *gen.Person, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "people", gen.JWTPermissionConstDelete) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "people")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope People to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeletePerson(ctx, id)
}

// DeleteAllPeople method
func (r *MutationResolver) DeleteAllPeople(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "people", gen.JWTPermissionConstDelete) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "people")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope People to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllPeople(ctx)
}
