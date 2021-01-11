package gen

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graph-gophers/dataloader"
	"github.com/vektah/gqlparser/v2/ast"
)

// GeneratedQueryResolver struct
type GeneratedQueryResolver struct{ *GeneratedResolver }

// QueryDeliveryHandlerOptions struct
type QueryDeliveryHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *DeliveryFilterType
}

// Delivery ...
func (r *GeneratedQueryResolver) Delivery(ctx context.Context, id *string, q *string, filter *DeliveryFilterType) (*Delivery, error) {
	opts := QueryDeliveryHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryDelivery(ctx, r.GeneratedResolver, opts)
}

// QueryDeliveryHandler handler
func QueryDeliveryHandler(ctx context.Context, r *GeneratedResolver, opts QueryDeliveryHandlerOptions) (*Delivery, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := DeliveryQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &DeliveryResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if opts.ID != nil {
		qb = qb.Where(TableName("deliveries")+".id = ?", *opts.ID)
	}

	var items []*Delivery
	giOpts := GetItemsOptions{
		Alias: TableName("deliveries"),
		Preloaders: []string{
			"Sender",
			"Receiver",
			"Deliver",
			"VehicleType",
		},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryDeliveriesHandlerOptions struct
type QueryDeliveriesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*DeliverySortType
	Filter *DeliveryFilterType
}

// Deliveries handler options
func (r *GeneratedQueryResolver) Deliveries(ctx context.Context, offset *int, limit *int, q *string, sort []*DeliverySortType, filter *DeliveryFilterType) (*DeliveryResultType, error) {
	opts := QueryDeliveriesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryDeliveries(ctx, r.GeneratedResolver, opts)
}

// DeliveriesItems handler
func (r *GeneratedResolver) DeliveriesItems(ctx context.Context, opts QueryDeliveriesHandlerOptions) (res []*Delivery, err error) {
	resultType, err := r.Handlers.QueryDeliveries(ctx, r, opts)
	if err != nil {
		return
	}
	err = resultType.GetItems(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("deliveries"),
	}, &res)
	if err != nil {
		return
	}
	return
}

// DeliveriesCount handler
func (r *GeneratedResolver) DeliveriesCount(ctx context.Context, opts QueryDeliveriesHandlerOptions) (count int, err error) {
	resultType, err := r.Handlers.QueryDeliveries(ctx, r, opts)
	if err != nil {
		return
	}
	return resultType.GetCount(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("deliveries"),
	}, &Delivery{})
}

// QueryDeliveriesHandler handler
func QueryDeliveriesHandler(ctx context.Context, r *GeneratedResolver, opts QueryDeliveriesHandlerOptions) (*DeliveryResultType, error) {
	query := DeliveryQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &DeliveryResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedDeliveryResultTypeResolver struct
type GeneratedDeliveryResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedDeliveryResultTypeResolver) Items(ctx context.Context, obj *DeliveryResultType) (items []*Delivery, err error) {
	otps := GetItemsOptions{
		Alias: TableName("deliveries"),
		Preloaders: []string{
			"Sender",
			"Receiver",
			"Deliver",
			"VehicleType",
		},
	}
	err = obj.GetItems(ctx, r.GetDB(ctx), otps, &items)

	for _, item := range items {

		item.SenderPreloaded = true
		item.ReceiverPreloaded = true
		item.DeliverPreloaded = true
		item.VehicleTypePreloaded = true
	}

	uniqueItems := []*Delivery{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedDeliveryResultTypeResolver) Count(ctx context.Context, obj *DeliveryResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias: TableName("deliveries"),
		Preloaders: []string{
			"Sender",
			"Receiver",
			"Deliver",
			"VehicleType",
		},
	}
	return obj.GetCount(ctx, r.GetDB(ctx), opts, &Delivery{})
}

// GeneratedDeliveryResolver struct
type GeneratedDeliveryResolver struct{ *GeneratedResolver }

// Instructions ...
func (r *GeneratedDeliveryResolver) Instructions(ctx context.Context, obj *Delivery) (res *string, err error) {
	return r.Handlers.DeliveryInstructions(ctx, r.GeneratedResolver, obj)
}

// DeliveryInstructionsHandler handler
func DeliveryInstructionsHandler(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *string, err error) {

	err = fmt.Errorf("Resolver handler for DeliveryInstructions not implemented")

	return
}

// Sender ...
func (r *GeneratedDeliveryResolver) Sender(ctx context.Context, obj *Delivery) (res *Person, err error) {
	return r.Handlers.DeliverySender(ctx, r.GeneratedResolver, obj)
}

// DeliverySenderHandler handler
func DeliverySenderHandler(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *Person, err error) {

	if obj.SenderPreloaded {
		res = obj.Sender
	} else {

		loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
		if obj.SenderID != nil {
			item, _err := loaders["Person"].Load(ctx, dataloader.StringKey(*obj.SenderID))()
			res, _ = item.(*Person)

			err = _err
		}

	}

	return
}

// Receiver ...
func (r *GeneratedDeliveryResolver) Receiver(ctx context.Context, obj *Delivery) (res *Person, err error) {
	return r.Handlers.DeliveryReceiver(ctx, r.GeneratedResolver, obj)
}

// DeliveryReceiverHandler handler
func DeliveryReceiverHandler(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *Person, err error) {

	if obj.ReceiverPreloaded {
		res = obj.Receiver
	} else {

		loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
		if obj.ReceiverID != nil {
			item, _err := loaders["Person"].Load(ctx, dataloader.StringKey(*obj.ReceiverID))()
			res, _ = item.(*Person)

			err = _err
		}

	}

	return
}

// Deliver ...
func (r *GeneratedDeliveryResolver) Deliver(ctx context.Context, obj *Delivery) (res *Person, err error) {
	return r.Handlers.DeliveryDeliver(ctx, r.GeneratedResolver, obj)
}

// DeliveryDeliverHandler handler
func DeliveryDeliverHandler(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *Person, err error) {

	if obj.DeliverPreloaded {
		res = obj.Deliver
	} else {

		loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
		if obj.DeliverID != nil {
			item, _err := loaders["Person"].Load(ctx, dataloader.StringKey(*obj.DeliverID))()
			res, _ = item.(*Person)

			err = _err
		}

	}

	return
}

// VehicleType ...
func (r *GeneratedDeliveryResolver) VehicleType(ctx context.Context, obj *Delivery) (res *VehicleType, err error) {
	return r.Handlers.DeliveryVehicleType(ctx, r.GeneratedResolver, obj)
}

// DeliveryVehicleTypeHandler handler
func DeliveryVehicleTypeHandler(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *VehicleType, err error) {

	if obj.VehicleTypePreloaded {
		res = obj.VehicleType
	} else {

		loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
		if obj.VehicleTypeID != nil {
			item, _err := loaders["VehicleType"].Load(ctx, dataloader.StringKey(*obj.VehicleTypeID))()
			res, _ = item.(*VehicleType)

			if res == nil {
				_err = fmt.Errorf("VehicleType with id '%s' not found", *obj.VehicleTypeID)
			}
			err = _err
		}

	}

	return
}

// DeliveryType ...
func (r *GeneratedDeliveryResolver) DeliveryType(ctx context.Context, obj *Delivery) (res *DeliveryType, err error) {
	return r.Handlers.DeliveryDeliveryType(ctx, r.GeneratedResolver, obj)
}

// DeliveryDeliveryTypeHandler handler
func DeliveryDeliveryTypeHandler(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *DeliveryType, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.DeliveryTypeID != nil {
		item, _err := loaders["DeliveryType"].Load(ctx, dataloader.StringKey(*obj.DeliveryTypeID))()
		res, _ = item.(*DeliveryType)

		if res == nil {
			_err = fmt.Errorf("DeliveryType with id '%s' not found", *obj.DeliveryTypeID)
		}
		err = _err
	}

	return
}

// DeliveryChannel ...
func (r *GeneratedDeliveryResolver) DeliveryChannel(ctx context.Context, obj *Delivery) (res *DeliveryChannel, err error) {
	return r.Handlers.DeliveryDeliveryChannel(ctx, r.GeneratedResolver, obj)
}

// DeliveryDeliveryChannelHandler handler
func DeliveryDeliveryChannelHandler(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *DeliveryChannel, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.DeliveryChannelID != nil {
		item, _err := loaders["DeliveryChannel"].Load(ctx, dataloader.StringKey(*obj.DeliveryChannelID))()
		res, _ = item.(*DeliveryChannel)

		if res == nil {
			_err = fmt.Errorf("DeliveryChannel with id '%s' not found", *obj.DeliveryChannelID)
		}
		err = _err
	}

	return
}

// QueryPersonHandlerOptions struct
type QueryPersonHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *PersonFilterType
}

// Person ...
func (r *GeneratedQueryResolver) Person(ctx context.Context, id *string, q *string, filter *PersonFilterType) (*Person, error) {
	opts := QueryPersonHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryPerson(ctx, r.GeneratedResolver, opts)
}

// QueryPersonHandler handler
func QueryPersonHandler(ctx context.Context, r *GeneratedResolver, opts QueryPersonHandlerOptions) (*Person, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := PersonQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &PersonResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if opts.ID != nil {
		qb = qb.Where(TableName("people")+".id = ?", *opts.ID)
	}

	var items []*Person
	giOpts := GetItemsOptions{
		Alias:      TableName("people"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryPeopleHandlerOptions struct
type QueryPeopleHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*PersonSortType
	Filter *PersonFilterType
}

// People handler options
func (r *GeneratedQueryResolver) People(ctx context.Context, offset *int, limit *int, q *string, sort []*PersonSortType, filter *PersonFilterType) (*PersonResultType, error) {
	opts := QueryPeopleHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryPeople(ctx, r.GeneratedResolver, opts)
}

// PeopleItems handler
func (r *GeneratedResolver) PeopleItems(ctx context.Context, opts QueryPeopleHandlerOptions) (res []*Person, err error) {
	resultType, err := r.Handlers.QueryPeople(ctx, r, opts)
	if err != nil {
		return
	}
	err = resultType.GetItems(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("people"),
	}, &res)
	if err != nil {
		return
	}
	return
}

// PeopleCount handler
func (r *GeneratedResolver) PeopleCount(ctx context.Context, opts QueryPeopleHandlerOptions) (count int, err error) {
	resultType, err := r.Handlers.QueryPeople(ctx, r, opts)
	if err != nil {
		return
	}
	return resultType.GetCount(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("people"),
	}, &Person{})
}

// QueryPeopleHandler handler
func QueryPeopleHandler(ctx context.Context, r *GeneratedResolver, opts QueryPeopleHandlerOptions) (*PersonResultType, error) {
	query := PersonQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &PersonResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedPersonResultTypeResolver struct
type GeneratedPersonResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedPersonResultTypeResolver) Items(ctx context.Context, obj *PersonResultType) (items []*Person, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("people"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.GetDB(ctx), otps, &items)

	uniqueItems := []*Person{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedPersonResultTypeResolver) Count(ctx context.Context, obj *PersonResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("people"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.GetDB(ctx), opts, &Person{})
}

// GeneratedPersonResolver struct
type GeneratedPersonResolver struct{ *GeneratedResolver }

// Deliveries ...
func (r *GeneratedPersonResolver) Deliveries(ctx context.Context, obj *Person) (res []*Delivery, err error) {
	return r.Handlers.PersonDeliveries(ctx, r.GeneratedResolver, obj)
}

// PersonDeliveriesHandler handler
func PersonDeliveriesHandler(ctx context.Context, r *GeneratedResolver, obj *Person) (res []*Delivery, err error) {

	items := []*Delivery{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Related(&items, "Deliveries").Error
	res = items

	return
}

// DeliveriesIds ...
func (r *GeneratedPersonResolver) DeliveriesIds(ctx context.Context, obj *Person) (ids []string, err error) {
	ids = []string{}

	items := []*Delivery{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Select(TableName("deliveries")+".id").Related(&items, "Deliveries").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// DeliveriesConnection method
func (r *GeneratedPersonResolver) DeliveriesConnection(ctx context.Context, obj *Person, offset *int, limit *int, q *string, sort []*DeliverySortType, filter *DeliveryFilterType) (res *DeliveryResultType, err error) {
	f := &DeliveryFilterType{
		Deliver: &PersonFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &DeliveryFilterType{
			And: []*DeliveryFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryDeliveriesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryDeliveries(ctx, r.GeneratedResolver, opts)
}

// DeliveriesSent ...
func (r *GeneratedPersonResolver) DeliveriesSent(ctx context.Context, obj *Person) (res []*Delivery, err error) {
	return r.Handlers.PersonDeliveriesSent(ctx, r.GeneratedResolver, obj)
}

// PersonDeliveriesSentHandler handler
func PersonDeliveriesSentHandler(ctx context.Context, r *GeneratedResolver, obj *Person) (res []*Delivery, err error) {

	items := []*Delivery{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Related(&items, "DeliveriesSent").Error
	res = items

	return
}

// DeliveriesSentIds ...
func (r *GeneratedPersonResolver) DeliveriesSentIds(ctx context.Context, obj *Person) (ids []string, err error) {
	ids = []string{}

	items := []*Delivery{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Select(TableName("deliveries")+".id").Related(&items, "DeliveriesSent").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// DeliveriesSentConnection method
func (r *GeneratedPersonResolver) DeliveriesSentConnection(ctx context.Context, obj *Person, offset *int, limit *int, q *string, sort []*DeliverySortType, filter *DeliveryFilterType) (res *DeliveryResultType, err error) {
	f := &DeliveryFilterType{
		Sender: &PersonFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &DeliveryFilterType{
			And: []*DeliveryFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryDeliveriesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryDeliveries(ctx, r.GeneratedResolver, opts)
}

// DeliveriesReceived ...
func (r *GeneratedPersonResolver) DeliveriesReceived(ctx context.Context, obj *Person) (res []*Delivery, err error) {
	return r.Handlers.PersonDeliveriesReceived(ctx, r.GeneratedResolver, obj)
}

// PersonDeliveriesReceivedHandler handler
func PersonDeliveriesReceivedHandler(ctx context.Context, r *GeneratedResolver, obj *Person) (res []*Delivery, err error) {

	items := []*Delivery{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Related(&items, "DeliveriesReceived").Error
	res = items

	return
}

// DeliveriesReceivedIds ...
func (r *GeneratedPersonResolver) DeliveriesReceivedIds(ctx context.Context, obj *Person) (ids []string, err error) {
	ids = []string{}

	items := []*Delivery{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Select(TableName("deliveries")+".id").Related(&items, "DeliveriesReceived").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// DeliveriesReceivedConnection method
func (r *GeneratedPersonResolver) DeliveriesReceivedConnection(ctx context.Context, obj *Person, offset *int, limit *int, q *string, sort []*DeliverySortType, filter *DeliveryFilterType) (res *DeliveryResultType, err error) {
	f := &DeliveryFilterType{
		Receiver: &PersonFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &DeliveryFilterType{
			And: []*DeliveryFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryDeliveriesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryDeliveries(ctx, r.GeneratedResolver, opts)
}

// QueryDeliveryTypeHandlerOptions struct
type QueryDeliveryTypeHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *DeliveryTypeFilterType
}

// DeliveryType ...
func (r *GeneratedQueryResolver) DeliveryType(ctx context.Context, id *string, q *string, filter *DeliveryTypeFilterType) (*DeliveryType, error) {
	opts := QueryDeliveryTypeHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryDeliveryType(ctx, r.GeneratedResolver, opts)
}

// QueryDeliveryTypeHandler handler
func QueryDeliveryTypeHandler(ctx context.Context, r *GeneratedResolver, opts QueryDeliveryTypeHandlerOptions) (*DeliveryType, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := DeliveryTypeQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &DeliveryTypeResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if opts.ID != nil {
		qb = qb.Where(TableName("delivery_types")+".id = ?", *opts.ID)
	}

	var items []*DeliveryType
	giOpts := GetItemsOptions{
		Alias:      TableName("delivery_types"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryDeliveryTypesHandlerOptions struct
type QueryDeliveryTypesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*DeliveryTypeSortType
	Filter *DeliveryTypeFilterType
}

// DeliveryTypes handler options
func (r *GeneratedQueryResolver) DeliveryTypes(ctx context.Context, offset *int, limit *int, q *string, sort []*DeliveryTypeSortType, filter *DeliveryTypeFilterType) (*DeliveryTypeResultType, error) {
	opts := QueryDeliveryTypesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryDeliveryTypes(ctx, r.GeneratedResolver, opts)
}

// DeliveryTypesItems handler
func (r *GeneratedResolver) DeliveryTypesItems(ctx context.Context, opts QueryDeliveryTypesHandlerOptions) (res []*DeliveryType, err error) {
	resultType, err := r.Handlers.QueryDeliveryTypes(ctx, r, opts)
	if err != nil {
		return
	}
	err = resultType.GetItems(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("delivery_types"),
	}, &res)
	if err != nil {
		return
	}
	return
}

// DeliveryTypesCount handler
func (r *GeneratedResolver) DeliveryTypesCount(ctx context.Context, opts QueryDeliveryTypesHandlerOptions) (count int, err error) {
	resultType, err := r.Handlers.QueryDeliveryTypes(ctx, r, opts)
	if err != nil {
		return
	}
	return resultType.GetCount(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("delivery_types"),
	}, &DeliveryType{})
}

// QueryDeliveryTypesHandler handler
func QueryDeliveryTypesHandler(ctx context.Context, r *GeneratedResolver, opts QueryDeliveryTypesHandlerOptions) (*DeliveryTypeResultType, error) {
	query := DeliveryTypeQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &DeliveryTypeResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedDeliveryTypeResultTypeResolver struct
type GeneratedDeliveryTypeResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedDeliveryTypeResultTypeResolver) Items(ctx context.Context, obj *DeliveryTypeResultType) (items []*DeliveryType, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("delivery_types"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.GetDB(ctx), otps, &items)

	uniqueItems := []*DeliveryType{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedDeliveryTypeResultTypeResolver) Count(ctx context.Context, obj *DeliveryTypeResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("delivery_types"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.GetDB(ctx), opts, &DeliveryType{})
}

// GeneratedDeliveryTypeResolver struct
type GeneratedDeliveryTypeResolver struct{ *GeneratedResolver }

// Delivery ...
func (r *GeneratedDeliveryTypeResolver) Delivery(ctx context.Context, obj *DeliveryType) (res *Delivery, err error) {
	return r.Handlers.DeliveryTypeDelivery(ctx, r.GeneratedResolver, obj)
}

// DeliveryTypeDeliveryHandler handler
func DeliveryTypeDeliveryHandler(ctx context.Context, r *GeneratedResolver, obj *DeliveryType) (res *Delivery, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.DeliveryID != nil {
		item, _err := loaders["Delivery"].Load(ctx, dataloader.StringKey(*obj.DeliveryID))()
		res, _ = item.(*Delivery)

		err = _err
	}

	return
}

// QueryDeliveryChannelHandlerOptions struct
type QueryDeliveryChannelHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *DeliveryChannelFilterType
}

// DeliveryChannel ...
func (r *GeneratedQueryResolver) DeliveryChannel(ctx context.Context, id *string, q *string, filter *DeliveryChannelFilterType) (*DeliveryChannel, error) {
	opts := QueryDeliveryChannelHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryDeliveryChannel(ctx, r.GeneratedResolver, opts)
}

// QueryDeliveryChannelHandler handler
func QueryDeliveryChannelHandler(ctx context.Context, r *GeneratedResolver, opts QueryDeliveryChannelHandlerOptions) (*DeliveryChannel, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := DeliveryChannelQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &DeliveryChannelResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if opts.ID != nil {
		qb = qb.Where(TableName("delivery_channels")+".id = ?", *opts.ID)
	}

	var items []*DeliveryChannel
	giOpts := GetItemsOptions{
		Alias:      TableName("delivery_channels"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryDeliveryChannelsHandlerOptions struct
type QueryDeliveryChannelsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*DeliveryChannelSortType
	Filter *DeliveryChannelFilterType
}

// DeliveryChannels handler options
func (r *GeneratedQueryResolver) DeliveryChannels(ctx context.Context, offset *int, limit *int, q *string, sort []*DeliveryChannelSortType, filter *DeliveryChannelFilterType) (*DeliveryChannelResultType, error) {
	opts := QueryDeliveryChannelsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryDeliveryChannels(ctx, r.GeneratedResolver, opts)
}

// DeliveryChannelsItems handler
func (r *GeneratedResolver) DeliveryChannelsItems(ctx context.Context, opts QueryDeliveryChannelsHandlerOptions) (res []*DeliveryChannel, err error) {
	resultType, err := r.Handlers.QueryDeliveryChannels(ctx, r, opts)
	if err != nil {
		return
	}
	err = resultType.GetItems(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("delivery_channels"),
	}, &res)
	if err != nil {
		return
	}
	return
}

// DeliveryChannelsCount handler
func (r *GeneratedResolver) DeliveryChannelsCount(ctx context.Context, opts QueryDeliveryChannelsHandlerOptions) (count int, err error) {
	resultType, err := r.Handlers.QueryDeliveryChannels(ctx, r, opts)
	if err != nil {
		return
	}
	return resultType.GetCount(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("delivery_channels"),
	}, &DeliveryChannel{})
}

// QueryDeliveryChannelsHandler handler
func QueryDeliveryChannelsHandler(ctx context.Context, r *GeneratedResolver, opts QueryDeliveryChannelsHandlerOptions) (*DeliveryChannelResultType, error) {
	query := DeliveryChannelQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &DeliveryChannelResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedDeliveryChannelResultTypeResolver struct
type GeneratedDeliveryChannelResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedDeliveryChannelResultTypeResolver) Items(ctx context.Context, obj *DeliveryChannelResultType) (items []*DeliveryChannel, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("delivery_channels"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.GetDB(ctx), otps, &items)

	uniqueItems := []*DeliveryChannel{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedDeliveryChannelResultTypeResolver) Count(ctx context.Context, obj *DeliveryChannelResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("delivery_channels"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.GetDB(ctx), opts, &DeliveryChannel{})
}

// GeneratedDeliveryChannelResolver struct
type GeneratedDeliveryChannelResolver struct{ *GeneratedResolver }

// Delivery ...
func (r *GeneratedDeliveryChannelResolver) Delivery(ctx context.Context, obj *DeliveryChannel) (res *Delivery, err error) {
	return r.Handlers.DeliveryChannelDelivery(ctx, r.GeneratedResolver, obj)
}

// DeliveryChannelDeliveryHandler handler
func DeliveryChannelDeliveryHandler(ctx context.Context, r *GeneratedResolver, obj *DeliveryChannel) (res *Delivery, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.DeliveryID != nil {
		item, _err := loaders["Delivery"].Load(ctx, dataloader.StringKey(*obj.DeliveryID))()
		res, _ = item.(*Delivery)

		err = _err
	}

	return
}

// QueryVehicleTypeHandlerOptions struct
type QueryVehicleTypeHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *VehicleTypeFilterType
}

// VehicleType ...
func (r *GeneratedQueryResolver) VehicleType(ctx context.Context, id *string, q *string, filter *VehicleTypeFilterType) (*VehicleType, error) {
	opts := QueryVehicleTypeHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryVehicleType(ctx, r.GeneratedResolver, opts)
}

// QueryVehicleTypeHandler handler
func QueryVehicleTypeHandler(ctx context.Context, r *GeneratedResolver, opts QueryVehicleTypeHandlerOptions) (*VehicleType, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := VehicleTypeQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &VehicleTypeResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if opts.ID != nil {
		qb = qb.Where(TableName("vehicle_types")+".id = ?", *opts.ID)
	}

	var items []*VehicleType
	giOpts := GetItemsOptions{
		Alias:      TableName("vehicle_types"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryVehicleTypesHandlerOptions struct
type QueryVehicleTypesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*VehicleTypeSortType
	Filter *VehicleTypeFilterType
}

// VehicleTypes handler options
func (r *GeneratedQueryResolver) VehicleTypes(ctx context.Context, offset *int, limit *int, q *string, sort []*VehicleTypeSortType, filter *VehicleTypeFilterType) (*VehicleTypeResultType, error) {
	opts := QueryVehicleTypesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryVehicleTypes(ctx, r.GeneratedResolver, opts)
}

// VehicleTypesItems handler
func (r *GeneratedResolver) VehicleTypesItems(ctx context.Context, opts QueryVehicleTypesHandlerOptions) (res []*VehicleType, err error) {
	resultType, err := r.Handlers.QueryVehicleTypes(ctx, r, opts)
	if err != nil {
		return
	}
	err = resultType.GetItems(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("vehicle_types"),
	}, &res)
	if err != nil {
		return
	}
	return
}

// VehicleTypesCount handler
func (r *GeneratedResolver) VehicleTypesCount(ctx context.Context, opts QueryVehicleTypesHandlerOptions) (count int, err error) {
	resultType, err := r.Handlers.QueryVehicleTypes(ctx, r, opts)
	if err != nil {
		return
	}
	return resultType.GetCount(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("vehicle_types"),
	}, &VehicleType{})
}

// QueryVehicleTypesHandler handler
func QueryVehicleTypesHandler(ctx context.Context, r *GeneratedResolver, opts QueryVehicleTypesHandlerOptions) (*VehicleTypeResultType, error) {
	query := VehicleTypeQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &VehicleTypeResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedVehicleTypeResultTypeResolver struct
type GeneratedVehicleTypeResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedVehicleTypeResultTypeResolver) Items(ctx context.Context, obj *VehicleTypeResultType) (items []*VehicleType, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("vehicle_types"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.GetDB(ctx), otps, &items)

	uniqueItems := []*VehicleType{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedVehicleTypeResultTypeResolver) Count(ctx context.Context, obj *VehicleTypeResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("vehicle_types"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.GetDB(ctx), opts, &VehicleType{})
}

// GeneratedVehicleTypeResolver struct
type GeneratedVehicleTypeResolver struct{ *GeneratedResolver }

// Delivery ...
func (r *GeneratedVehicleTypeResolver) Delivery(ctx context.Context, obj *VehicleType) (res *Delivery, err error) {
	return r.Handlers.VehicleTypeDelivery(ctx, r.GeneratedResolver, obj)
}

// VehicleTypeDeliveryHandler handler
func VehicleTypeDeliveryHandler(ctx context.Context, r *GeneratedResolver, obj *VehicleType) (res *Delivery, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.DeliveryID != nil {
		item, _err := loaders["Delivery"].Load(ctx, dataloader.StringKey(*obj.DeliveryID))()
		res, _ = item.(*Delivery)

		err = _err
	}

	return
}
