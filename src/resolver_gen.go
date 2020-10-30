package src

import (
	"context"

	"github.com/loopcontext/deliver-api-go/gen"
)

// NewResolver ...
func NewResolver(db *gen.DB, ec *gen.EventController) *Resolver {
	handlers := gen.DefaultResolutionHandlers()
	return &Resolver{&gen.GeneratedResolver{Handlers: handlers, DB: db, EventController: ec}}
}

// Resolver ...
type Resolver struct {
	*gen.GeneratedResolver
}

// MutationResolver ...
type MutationResolver struct {
	*gen.GeneratedMutationResolver
}

// BeginTransaction ...
func (r *MutationResolver) BeginTransaction(ctx context.Context, fn func(context.Context) error) error {
	ctx = gen.EnrichContextWithMutations(ctx, r.GeneratedResolver)
	err := fn(ctx)
	if err != nil {
		tx := r.GeneratedResolver.GetDB(ctx)
		tx.Rollback()
		return err
	}
	return gen.FinishMutationContext(ctx, r.GeneratedResolver)
}

// QueryResolver ...
type QueryResolver struct {
	*gen.GeneratedQueryResolver
}

// Mutation ...
func (r *Resolver) Mutation() gen.MutationResolver {
	return &MutationResolver{
		GeneratedMutationResolver: &gen.GeneratedMutationResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// Query ...
func (r *Resolver) Query() gen.QueryResolver {
	return &QueryResolver{
		GeneratedQueryResolver: &gen.GeneratedQueryResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// DeliveryResultTypeResolver struct
type DeliveryResultTypeResolver struct {
	*gen.GeneratedDeliveryResultTypeResolver
}

// DeliveryResultType ...
func (r *Resolver) DeliveryResultType() gen.DeliveryResultTypeResolver {
	return &DeliveryResultTypeResolver{
		GeneratedDeliveryResultTypeResolver: &gen.GeneratedDeliveryResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// DeliveryResolver struct
type DeliveryResolver struct {
	*gen.GeneratedDeliveryResolver
}

// Delivery ...
func (r *Resolver) Delivery() gen.DeliveryResolver {
	return &DeliveryResolver{
		GeneratedDeliveryResolver: &gen.GeneratedDeliveryResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// PersonResultTypeResolver struct
type PersonResultTypeResolver struct {
	*gen.GeneratedPersonResultTypeResolver
}

// PersonResultType ...
func (r *Resolver) PersonResultType() gen.PersonResultTypeResolver {
	return &PersonResultTypeResolver{
		GeneratedPersonResultTypeResolver: &gen.GeneratedPersonResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// PersonResolver struct
type PersonResolver struct {
	*gen.GeneratedPersonResolver
}

// Person ...
func (r *Resolver) Person() gen.PersonResolver {
	return &PersonResolver{
		GeneratedPersonResolver: &gen.GeneratedPersonResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// DeliveryTypeResultTypeResolver struct
type DeliveryTypeResultTypeResolver struct {
	*gen.GeneratedDeliveryTypeResultTypeResolver
}

// DeliveryTypeResultType ...
func (r *Resolver) DeliveryTypeResultType() gen.DeliveryTypeResultTypeResolver {
	return &DeliveryTypeResultTypeResolver{
		GeneratedDeliveryTypeResultTypeResolver: &gen.GeneratedDeliveryTypeResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// DeliveryTypeResolver struct
type DeliveryTypeResolver struct {
	*gen.GeneratedDeliveryTypeResolver
}

// DeliveryType ...
func (r *Resolver) DeliveryType() gen.DeliveryTypeResolver {
	return &DeliveryTypeResolver{
		GeneratedDeliveryTypeResolver: &gen.GeneratedDeliveryTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// DeliveryChannelResultTypeResolver struct
type DeliveryChannelResultTypeResolver struct {
	*gen.GeneratedDeliveryChannelResultTypeResolver
}

// DeliveryChannelResultType ...
func (r *Resolver) DeliveryChannelResultType() gen.DeliveryChannelResultTypeResolver {
	return &DeliveryChannelResultTypeResolver{
		GeneratedDeliveryChannelResultTypeResolver: &gen.GeneratedDeliveryChannelResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// DeliveryChannelResolver struct
type DeliveryChannelResolver struct {
	*gen.GeneratedDeliveryChannelResolver
}

// DeliveryChannel ...
func (r *Resolver) DeliveryChannel() gen.DeliveryChannelResolver {
	return &DeliveryChannelResolver{
		GeneratedDeliveryChannelResolver: &gen.GeneratedDeliveryChannelResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// VehicleTypeResultTypeResolver struct
type VehicleTypeResultTypeResolver struct {
	*gen.GeneratedVehicleTypeResultTypeResolver
}

// VehicleTypeResultType ...
func (r *Resolver) VehicleTypeResultType() gen.VehicleTypeResultTypeResolver {
	return &VehicleTypeResultTypeResolver{
		GeneratedVehicleTypeResultTypeResolver: &gen.GeneratedVehicleTypeResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// VehicleTypeResolver struct
type VehicleTypeResolver struct {
	*gen.GeneratedVehicleTypeResolver
}

// VehicleType ...
func (r *Resolver) VehicleType() gen.VehicleTypeResolver {
	return &VehicleTypeResolver{
		GeneratedVehicleTypeResolver: &gen.GeneratedVehicleTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}
