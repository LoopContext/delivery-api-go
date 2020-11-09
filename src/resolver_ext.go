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
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "deliveries", gen.JWTPermissionConstList[:1]) {
		return r.GeneratedQueryResolver.Deliveries(ctx, offset, limit, q, sort, filter)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "deliveries")
}

// CreateDelivery method
func (r *MutationResolver) CreateDelivery(ctx context.Context, input map[string]interface{}) (item *gen.Delivery, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "deliveries", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "deliveries")
	}
	return r.GeneratedMutationResolver.CreateDelivery(ctx, input)
}

// ReadDelivery method
func (r *QueryResolver) Delivery(ctx context.Context, id *string, q *string, filter *gen.DeliveryFilterType) (*gen.Delivery, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "deliveries", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "deliveries")
	}
	if !gen.HasRole(jwtClaims, "admin") {
		qd, err := r.GeneratedResolver.Handlers.QueryDelivery(ctx, r.GeneratedResolver,
			gen.QueryDeliveryHandlerOptions{
				Filter: &gen.DeliveryFilterType{
					Sender: &gen.PersonFilterType{
						UserID: &jwtClaims.Subject,
					},
					ID: id,
				},
			})
		if err != nil || qd == nil {
			return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "deliveries")
		}
	}
	return r.GeneratedQueryResolver.Delivery(ctx, id, q, filter)
}

// UpdateDelivery method
func (r *MutationResolver) UpdateDelivery(ctx context.Context, id string, input map[string]interface{}) (item *gen.Delivery, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "deliveries", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "deliveries")
	}
	if !gen.HasRole(jwtClaims, "admin") {
		completed := false
		qd, err := r.GeneratedResolver.Handlers.QueryDelivery(ctx, r.GeneratedResolver,
			gen.QueryDeliveryHandlerOptions{
				Filter: &gen.DeliveryFilterType{
					Sender: &gen.PersonFilterType{
						UserID: &jwtClaims.Subject,
					},
					Completed: &completed,
					ID:        &id,
				},
			})
		if err != nil || qd == nil {
			return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "deliveries")
		}
	}
	return r.GeneratedMutationResolver.UpdateDelivery(ctx, id, input)
}

// DeleteDelivery method
func (r *MutationResolver) DeleteDelivery(ctx context.Context, id string) (item *gen.Delivery, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "deliveries", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteDelivery(ctx, id)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "deliveries")
}

// DeleteAllDeliveries method
func (r *MutationResolver) DeleteAllDeliveries(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "deliveries", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteAllDeliveries(ctx)
	}
	return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "deliveries")
}

// People method
func (r *QueryResolver) People(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.PersonSortType, filter *gen.PersonFilterType) (*gen.PersonResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "people", gen.JWTPermissionConstList[:1]) {
		return r.GeneratedQueryResolver.People(ctx, offset, limit, q, sort, filter)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "people")
}

// CreatePerson method
func (r *MutationResolver) CreatePerson(ctx context.Context, input map[string]interface{}) (item *gen.Person, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "people", gen.JWTPermissionConstCreate[:1]) {
		return r.GeneratedMutationResolver.CreatePerson(ctx, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "people")
}

// ReadPerson method
func (r *QueryResolver) Person(ctx context.Context, id *string, q *string, filter *gen.PersonFilterType) (*gen.Person, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "people", gen.JWTPermissionConstRead[:1]) {
		return r.GeneratedQueryResolver.Person(ctx, id, q, filter)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "people")
}

// UpdatePerson method
func (r *MutationResolver) UpdatePerson(ctx context.Context, id string, input map[string]interface{}) (item *gen.Person, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "people", gen.JWTPermissionConstUpdate[:1]) {
		return r.GeneratedMutationResolver.UpdatePerson(ctx, id, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "people")
}

// DeletePerson method
func (r *MutationResolver) DeletePerson(ctx context.Context, id string) (item *gen.Person, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "people", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeletePerson(ctx, id)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "people")
}

// DeleteAllPeople method
func (r *MutationResolver) DeleteAllPeople(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "people", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteAllPeople(ctx)
	}
	return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "people")
}

// DeliveryTypes method
func (r *QueryResolver) DeliveryTypes(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.DeliveryTypeSortType, filter *gen.DeliveryTypeFilterType) (*gen.DeliveryTypeResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "delivery_types", gen.JWTPermissionConstList[:1]) {
		return r.GeneratedQueryResolver.DeliveryTypes(ctx, offset, limit, q, sort, filter)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "delivery_types")
}

// CreateDeliveryType method
func (r *MutationResolver) CreateDeliveryType(ctx context.Context, input map[string]interface{}) (item *gen.DeliveryType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "delivery_types", gen.JWTPermissionConstCreate[:1]) {
		return r.GeneratedMutationResolver.CreateDeliveryType(ctx, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "delivery_types")
}

// ReadDeliveryType method
func (r *QueryResolver) DeliveryType(ctx context.Context, id *string, q *string, filter *gen.DeliveryTypeFilterType) (*gen.DeliveryType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "delivery_types", gen.JWTPermissionConstRead[:1]) {
		return r.GeneratedQueryResolver.DeliveryType(ctx, id, q, filter)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "delivery_types")
}

// UpdateDeliveryType method
func (r *MutationResolver) UpdateDeliveryType(ctx context.Context, id string, input map[string]interface{}) (item *gen.DeliveryType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "delivery_types", gen.JWTPermissionConstUpdate[:1]) {
		return r.GeneratedMutationResolver.UpdateDeliveryType(ctx, id, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "delivery_types")
}

// DeleteDeliveryType method
func (r *MutationResolver) DeleteDeliveryType(ctx context.Context, id string) (item *gen.DeliveryType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "delivery_types", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteDeliveryType(ctx, id)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "delivery_types")
}

// DeleteAllDeliveryTypes method
func (r *MutationResolver) DeleteAllDeliveryTypes(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "delivery_types", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteAllDeliveryTypes(ctx)
	}
	return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "delivery_types")
}

// DeliveryChannels method
func (r *QueryResolver) DeliveryChannels(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.DeliveryChannelSortType, filter *gen.DeliveryChannelFilterType) (*gen.DeliveryChannelResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "delivery_channels", gen.JWTPermissionConstList[:1]) {
		return r.GeneratedQueryResolver.DeliveryChannels(ctx, offset, limit, q, sort, filter)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "delivery_channels")
}

// CreateDeliveryChannel method
func (r *MutationResolver) CreateDeliveryChannel(ctx context.Context, input map[string]interface{}) (item *gen.DeliveryChannel, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "delivery_channels", gen.JWTPermissionConstCreate[:1]) {
		return r.GeneratedMutationResolver.CreateDeliveryChannel(ctx, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "delivery_channels")
}

// ReadDeliveryChannel method
func (r *QueryResolver) DeliveryChannel(ctx context.Context, id *string, q *string, filter *gen.DeliveryChannelFilterType) (*gen.DeliveryChannel, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "delivery_channels", gen.JWTPermissionConstRead[:1]) {
		return r.GeneratedQueryResolver.DeliveryChannel(ctx, id, q, filter)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "delivery_channels")
}

// UpdateDeliveryChannel method
func (r *MutationResolver) UpdateDeliveryChannel(ctx context.Context, id string, input map[string]interface{}) (item *gen.DeliveryChannel, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "delivery_channels", gen.JWTPermissionConstUpdate[:1]) {
		return r.GeneratedMutationResolver.UpdateDeliveryChannel(ctx, id, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "delivery_channels")
}

// DeleteDeliveryChannel method
func (r *MutationResolver) DeleteDeliveryChannel(ctx context.Context, id string) (item *gen.DeliveryChannel, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "delivery_channels", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteDeliveryChannel(ctx, id)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "delivery_channels")
}

// DeleteAllDeliveryChannels method
func (r *MutationResolver) DeleteAllDeliveryChannels(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "delivery_channels", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteAllDeliveryChannels(ctx)
	}
	return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "delivery_channels")
}

// VehicleTypes method
func (r *QueryResolver) VehicleTypes(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.VehicleTypeSortType, filter *gen.VehicleTypeFilterType) (*gen.VehicleTypeResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "vehicle_types", gen.JWTPermissionConstList[:1]) {
		return r.GeneratedQueryResolver.VehicleTypes(ctx, offset, limit, q, sort, filter)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "vehicle_types")
}

// CreateVehicleType method
func (r *MutationResolver) CreateVehicleType(ctx context.Context, input map[string]interface{}) (item *gen.VehicleType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "vehicle_types", gen.JWTPermissionConstCreate[:1]) {
		return r.GeneratedMutationResolver.CreateVehicleType(ctx, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "vehicle_types")
}

// ReadVehicleType method
func (r *QueryResolver) VehicleType(ctx context.Context, id *string, q *string, filter *gen.VehicleTypeFilterType) (*gen.VehicleType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "vehicle_types", gen.JWTPermissionConstRead[:1]) {
		return r.GeneratedQueryResolver.VehicleType(ctx, id, q, filter)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "vehicle_types")
}

// UpdateVehicleType method
func (r *MutationResolver) UpdateVehicleType(ctx context.Context, id string, input map[string]interface{}) (item *gen.VehicleType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "vehicle_types", gen.JWTPermissionConstUpdate[:1]) {
		return r.GeneratedMutationResolver.UpdateVehicleType(ctx, id, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "vehicle_types")
}

// DeleteVehicleType method
func (r *MutationResolver) DeleteVehicleType(ctx context.Context, id string) (item *gen.VehicleType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "vehicle_types", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteVehicleType(ctx, id)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "vehicle_types")
}

// DeleteAllVehicleTypes method
func (r *MutationResolver) DeleteAllVehicleTypes(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "vehicle_types", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteAllVehicleTypes(ctx)
	}
	return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "vehicle_types")
}
