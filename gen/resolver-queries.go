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
	if qb == nil {
		qb = r.DB.Query()
	}
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

// Deliveries ...
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
		},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	for _, item := range items {

		item.SenderPreloaded = true
		item.ReceiverPreloaded = true
		item.DeliverPreloaded = true
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
		},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &Delivery{})
}

// GeneratedDeliveryResolver struct
type GeneratedDeliveryResolver struct{ *GeneratedResolver }

// VehicleType ...
func (r *GeneratedDeliveryResolver) VehicleType(ctx context.Context, obj *Delivery) (res *VehicleType, err error) {
	return r.Handlers.DeliveryVehicleType(ctx, r.GeneratedResolver, obj)
}

// DeliveryVehicleTypeHandler handler
func DeliveryVehicleTypeHandler(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *VehicleType, err error) {

	err = fmt.Errorf("Resolver handler for DeliveryVehicleType not implemented")

	return
}

// PaymentForm ...
func (r *GeneratedDeliveryResolver) PaymentForm(ctx context.Context, obj *Delivery) (res *PaymentForm, err error) {
	return r.Handlers.DeliveryPaymentForm(ctx, r.GeneratedResolver, obj)
}

// DeliveryPaymentFormHandler handler
func DeliveryPaymentFormHandler(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *PaymentForm, err error) {

	err = fmt.Errorf("Resolver handler for DeliveryPaymentForm not implemented")

	return
}

// DeliveryType ...
func (r *GeneratedDeliveryResolver) DeliveryType(ctx context.Context, obj *Delivery) (res *DeliveryType, err error) {
	return r.Handlers.DeliveryDeliveryType(ctx, r.GeneratedResolver, obj)
}

// DeliveryDeliveryTypeHandler handler
func DeliveryDeliveryTypeHandler(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *DeliveryType, err error) {

	err = fmt.Errorf("Resolver handler for DeliveryDeliveryType not implemented")

	return
}

// DeliveryChannel ...
func (r *GeneratedDeliveryResolver) DeliveryChannel(ctx context.Context, obj *Delivery) (res *DeliveryChannel, err error) {
	return r.Handlers.DeliveryDeliveryChannel(ctx, r.GeneratedResolver, obj)
}

// DeliveryDeliveryChannelHandler handler
func DeliveryDeliveryChannelHandler(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *DeliveryChannel, err error) {

	err = fmt.Errorf("Resolver handler for DeliveryDeliveryChannel not implemented")

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

			if res == nil {
				_err = fmt.Errorf("Person with id '%s' not found", *obj.SenderID)
			}
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

			if res == nil {
				_err = fmt.Errorf("Person with id '%s' not found", *obj.ReceiverID)
			}
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

			if res == nil {
				_err = fmt.Errorf("Person with id '%s' not found", *obj.DeliverID)
			}
			err = _err
		}

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
	if qb == nil {
		qb = r.DB.Query()
	}
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

// People ...
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
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

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
	return obj.GetCount(ctx, r.DB.db, opts, &Person{})
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
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Related(&items, "Deliveries").Error
	res = items

	return
}

// DeliveriesIds ...
func (r *GeneratedPersonResolver) DeliveriesIds(ctx context.Context, obj *Person) (ids []string, err error) {
	ids = []string{}

	items := []*Delivery{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
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
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Related(&items, "DeliveriesSent").Error
	res = items

	return
}

// DeliveriesSentIds ...
func (r *GeneratedPersonResolver) DeliveriesSentIds(ctx context.Context, obj *Person) (ids []string, err error) {
	ids = []string{}

	items := []*Delivery{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
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
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Related(&items, "DeliveriesReceived").Error
	res = items

	return
}

// DeliveriesReceivedIds ...
func (r *GeneratedPersonResolver) DeliveriesReceivedIds(ctx context.Context, obj *Person) (ids []string, err error) {
	ids = []string{}

	items := []*Delivery{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
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

// PaymentStatus ...
func (r *GeneratedPersonResolver) PaymentStatus(ctx context.Context, obj *Person) (res *PaymentStatus, err error) {
	return r.Handlers.PersonPaymentStatus(ctx, r.GeneratedResolver, obj)
}

// PersonPaymentStatusHandler handler
func PersonPaymentStatusHandler(ctx context.Context, r *GeneratedResolver, obj *Person) (res *PaymentStatus, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.PaymentStatusID != nil {
		item, _err := loaders["PaymentStatus"].Load(ctx, dataloader.StringKey(*obj.PaymentStatusID))()
		res, _ = item.(*PaymentStatus)

		if res == nil {
			_err = fmt.Errorf("PaymentStatus with id '%s' not found", *obj.PaymentStatusID)
		}
		err = _err
	}

	return
}

// PaymentHistory ...
func (r *GeneratedPersonResolver) PaymentHistory(ctx context.Context, obj *Person) (res *PaymentHistory, err error) {
	return r.Handlers.PersonPaymentHistory(ctx, r.GeneratedResolver, obj)
}

// PersonPaymentHistoryHandler handler
func PersonPaymentHistoryHandler(ctx context.Context, r *GeneratedResolver, obj *Person) (res *PaymentHistory, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.PaymentHistoryID != nil {
		item, _err := loaders["PaymentHistory"].Load(ctx, dataloader.StringKey(*obj.PaymentHistoryID))()
		res, _ = item.(*PaymentHistory)

		if res == nil {
			_err = fmt.Errorf("PaymentHistory with id '%s' not found", *obj.PaymentHistoryID)
		}
		err = _err
	}

	return
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
	if qb == nil {
		qb = r.DB.Query()
	}
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

// DeliveryTypes ...
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
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

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
	return obj.GetCount(ctx, r.DB.db, opts, &DeliveryType{})
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
	if qb == nil {
		qb = r.DB.Query()
	}
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

// DeliveryChannels ...
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
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

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
	return obj.GetCount(ctx, r.DB.db, opts, &DeliveryChannel{})
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
	if qb == nil {
		qb = r.DB.Query()
	}
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

// VehicleTypes ...
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
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

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
	return obj.GetCount(ctx, r.DB.db, opts, &VehicleType{})
}

// QueryPaymentFormHandlerOptions struct
type QueryPaymentFormHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *PaymentFormFilterType
}

// PaymentForm ...
func (r *GeneratedQueryResolver) PaymentForm(ctx context.Context, id *string, q *string, filter *PaymentFormFilterType) (*PaymentForm, error) {
	opts := QueryPaymentFormHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryPaymentForm(ctx, r.GeneratedResolver, opts)
}

// QueryPaymentFormHandler handler
func QueryPaymentFormHandler(ctx context.Context, r *GeneratedResolver, opts QueryPaymentFormHandlerOptions) (*PaymentForm, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := PaymentFormQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &PaymentFormResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("payment_forms")+".id = ?", *opts.ID)
	}

	var items []*PaymentForm
	giOpts := GetItemsOptions{
		Alias:      TableName("payment_forms"),
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

// QueryPaymentFormsHandlerOptions struct
type QueryPaymentFormsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*PaymentFormSortType
	Filter *PaymentFormFilterType
}

// PaymentForms ...
func (r *GeneratedQueryResolver) PaymentForms(ctx context.Context, offset *int, limit *int, q *string, sort []*PaymentFormSortType, filter *PaymentFormFilterType) (*PaymentFormResultType, error) {
	opts := QueryPaymentFormsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryPaymentForms(ctx, r.GeneratedResolver, opts)
}

// QueryPaymentFormsHandler handler
func QueryPaymentFormsHandler(ctx context.Context, r *GeneratedResolver, opts QueryPaymentFormsHandlerOptions) (*PaymentFormResultType, error) {
	query := PaymentFormQueryFilter{opts.Q}

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

	return &PaymentFormResultType{
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

// GeneratedPaymentFormResultTypeResolver struct
type GeneratedPaymentFormResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedPaymentFormResultTypeResolver) Items(ctx context.Context, obj *PaymentFormResultType) (items []*PaymentForm, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("payment_forms"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*PaymentForm{}
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
func (r *GeneratedPaymentFormResultTypeResolver) Count(ctx context.Context, obj *PaymentFormResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("payment_forms"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &PaymentForm{})
}

// QueryPaymentStatusHandlerOptions struct
type QueryPaymentStatusHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *PaymentStatusFilterType
}

// PaymentStatus ...
func (r *GeneratedQueryResolver) PaymentStatus(ctx context.Context, id *string, q *string, filter *PaymentStatusFilterType) (*PaymentStatus, error) {
	opts := QueryPaymentStatusHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryPaymentStatus(ctx, r.GeneratedResolver, opts)
}

// QueryPaymentStatusHandler handler
func QueryPaymentStatusHandler(ctx context.Context, r *GeneratedResolver, opts QueryPaymentStatusHandlerOptions) (*PaymentStatus, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := PaymentStatusQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &PaymentStatusResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("payment_statuses")+".id = ?", *opts.ID)
	}

	var items []*PaymentStatus
	giOpts := GetItemsOptions{
		Alias:      TableName("payment_statuses"),
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

// QueryPaymentStatusesHandlerOptions struct
type QueryPaymentStatusesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*PaymentStatusSortType
	Filter *PaymentStatusFilterType
}

// PaymentStatuses ...
func (r *GeneratedQueryResolver) PaymentStatuses(ctx context.Context, offset *int, limit *int, q *string, sort []*PaymentStatusSortType, filter *PaymentStatusFilterType) (*PaymentStatusResultType, error) {
	opts := QueryPaymentStatusesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryPaymentStatuses(ctx, r.GeneratedResolver, opts)
}

// QueryPaymentStatusesHandler handler
func QueryPaymentStatusesHandler(ctx context.Context, r *GeneratedResolver, opts QueryPaymentStatusesHandlerOptions) (*PaymentStatusResultType, error) {
	query := PaymentStatusQueryFilter{opts.Q}

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

	return &PaymentStatusResultType{
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

// GeneratedPaymentStatusResultTypeResolver struct
type GeneratedPaymentStatusResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedPaymentStatusResultTypeResolver) Items(ctx context.Context, obj *PaymentStatusResultType) (items []*PaymentStatus, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("payment_statuses"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*PaymentStatus{}
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
func (r *GeneratedPaymentStatusResultTypeResolver) Count(ctx context.Context, obj *PaymentStatusResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("payment_statuses"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &PaymentStatus{})
}

// GeneratedPaymentStatusResolver struct
type GeneratedPaymentStatusResolver struct{ *GeneratedResolver }

// Person ...
func (r *GeneratedPaymentStatusResolver) Person(ctx context.Context, obj *PaymentStatus) (res *Person, err error) {
	return r.Handlers.PaymentStatusPerson(ctx, r.GeneratedResolver, obj)
}

// PaymentStatusPersonHandler handler
func PaymentStatusPersonHandler(ctx context.Context, r *GeneratedResolver, obj *PaymentStatus) (res *Person, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.PersonID != nil {
		item, _err := loaders["Person"].Load(ctx, dataloader.StringKey(*obj.PersonID))()
		res, _ = item.(*Person)

		if res == nil {
			_err = fmt.Errorf("Person with id '%s' not found", *obj.PersonID)
		}
		err = _err
	}

	return
}

// QueryPaymentHistoryHandlerOptions struct
type QueryPaymentHistoryHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *PaymentHistoryFilterType
}

// PaymentHistory ...
func (r *GeneratedQueryResolver) PaymentHistory(ctx context.Context, id *string, q *string, filter *PaymentHistoryFilterType) (*PaymentHistory, error) {
	opts := QueryPaymentHistoryHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryPaymentHistory(ctx, r.GeneratedResolver, opts)
}

// QueryPaymentHistoryHandler handler
func QueryPaymentHistoryHandler(ctx context.Context, r *GeneratedResolver, opts QueryPaymentHistoryHandlerOptions) (*PaymentHistory, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := PaymentHistoryQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &PaymentHistoryResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("payment_histories")+".id = ?", *opts.ID)
	}

	var items []*PaymentHistory
	giOpts := GetItemsOptions{
		Alias:      TableName("payment_histories"),
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

// QueryPaymentHistoriesHandlerOptions struct
type QueryPaymentHistoriesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*PaymentHistorySortType
	Filter *PaymentHistoryFilterType
}

// PaymentHistories ...
func (r *GeneratedQueryResolver) PaymentHistories(ctx context.Context, offset *int, limit *int, q *string, sort []*PaymentHistorySortType, filter *PaymentHistoryFilterType) (*PaymentHistoryResultType, error) {
	opts := QueryPaymentHistoriesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryPaymentHistories(ctx, r.GeneratedResolver, opts)
}

// QueryPaymentHistoriesHandler handler
func QueryPaymentHistoriesHandler(ctx context.Context, r *GeneratedResolver, opts QueryPaymentHistoriesHandlerOptions) (*PaymentHistoryResultType, error) {
	query := PaymentHistoryQueryFilter{opts.Q}

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

	return &PaymentHistoryResultType{
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

// GeneratedPaymentHistoryResultTypeResolver struct
type GeneratedPaymentHistoryResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedPaymentHistoryResultTypeResolver) Items(ctx context.Context, obj *PaymentHistoryResultType) (items []*PaymentHistory, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("payment_histories"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*PaymentHistory{}
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
func (r *GeneratedPaymentHistoryResultTypeResolver) Count(ctx context.Context, obj *PaymentHistoryResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("payment_histories"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &PaymentHistory{})
}

// GeneratedPaymentHistoryResolver struct
type GeneratedPaymentHistoryResolver struct{ *GeneratedResolver }

// PaymentForm ...
func (r *GeneratedPaymentHistoryResolver) PaymentForm(ctx context.Context, obj *PaymentHistory) (res *PaymentForm, err error) {
	return r.Handlers.PaymentHistoryPaymentForm(ctx, r.GeneratedResolver, obj)
}

// PaymentHistoryPaymentFormHandler handler
func PaymentHistoryPaymentFormHandler(ctx context.Context, r *GeneratedResolver, obj *PaymentHistory) (res *PaymentForm, err error) {

	err = fmt.Errorf("Resolver handler for PaymentHistoryPaymentForm not implemented")

	return
}

// Person ...
func (r *GeneratedPaymentHistoryResolver) Person(ctx context.Context, obj *PaymentHistory) (res *Person, err error) {
	return r.Handlers.PaymentHistoryPerson(ctx, r.GeneratedResolver, obj)
}

// PaymentHistoryPersonHandler handler
func PaymentHistoryPersonHandler(ctx context.Context, r *GeneratedResolver, obj *PaymentHistory) (res *Person, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.PersonID != nil {
		item, _err := loaders["Person"].Load(ctx, dataloader.StringKey(*obj.PersonID))()
		res, _ = item.(*Person)

		if res == nil {
			_err = fmt.Errorf("Person with id '%s' not found", *obj.PersonID)
		}
		err = _err
	}

	return
}
