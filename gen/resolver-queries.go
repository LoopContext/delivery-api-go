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
		Alias:      TableName("deliveries"),
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
		Alias:      TableName("deliveries"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

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
		Alias:      TableName("deliveries"),
		Preloaders: []string{},
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

// Sender ...
func (r *GeneratedDeliveryResolver) Sender(ctx context.Context, obj *Delivery) (res *Person, err error) {
	return r.Handlers.DeliverySender(ctx, r.GeneratedResolver, obj)
}

// DeliverySenderHandler handler
func DeliverySenderHandler(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *Person, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.SenderID != nil {
		item, _err := loaders["Person"].Load(ctx, dataloader.StringKey(*obj.SenderID))()
		res, _ = item.(*Person)

		if res == nil {
			_err = fmt.Errorf("Person with id '%s' not found", *obj.SenderID)
		}
		err = _err
	}

	return
}

// Receiver ...
func (r *GeneratedDeliveryResolver) Receiver(ctx context.Context, obj *Delivery) (res *Person, err error) {
	return r.Handlers.DeliveryReceiver(ctx, r.GeneratedResolver, obj)
}

// DeliveryReceiverHandler handler
func DeliveryReceiverHandler(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *Person, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.ReceiverID != nil {
		item, _err := loaders["Person"].Load(ctx, dataloader.StringKey(*obj.ReceiverID))()
		res, _ = item.(*Person)

		if res == nil {
			_err = fmt.Errorf("Person with id '%s' not found", *obj.ReceiverID)
		}
		err = _err
	}

	return
}

// Deliver ...
func (r *GeneratedDeliveryResolver) Deliver(ctx context.Context, obj *Delivery) (res *Deliver, err error) {
	return r.Handlers.DeliveryDeliver(ctx, r.GeneratedResolver, obj)
}

// DeliveryDeliverHandler handler
func DeliveryDeliverHandler(ctx context.Context, r *GeneratedResolver, obj *Delivery) (res *Deliver, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.DeliverID != nil {
		item, _err := loaders["Deliver"].Load(ctx, dataloader.StringKey(*obj.DeliverID))()
		res, _ = item.(*Deliver)

		if res == nil {
			_err = fmt.Errorf("Deliver with id '%s' not found", *obj.DeliverID)
		}
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

// QueryDeliverHandlerOptions struct
type QueryDeliverHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *DeliverFilterType
}

// Deliver ...
func (r *GeneratedQueryResolver) Deliver(ctx context.Context, id *string, q *string, filter *DeliverFilterType) (*Deliver, error) {
	opts := QueryDeliverHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryDeliver(ctx, r.GeneratedResolver, opts)
}

// QueryDeliverHandler handler
func QueryDeliverHandler(ctx context.Context, r *GeneratedResolver, opts QueryDeliverHandlerOptions) (*Deliver, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := DeliverQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &DeliverResultType{
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
		qb = qb.Where(TableName("delivers")+".id = ?", *opts.ID)
	}

	var items []*Deliver
	giOpts := GetItemsOptions{
		Alias:      TableName("delivers"),
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

// QueryDeliversHandlerOptions struct
type QueryDeliversHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*DeliverSortType
	Filter *DeliverFilterType
}

// Delivers ...
func (r *GeneratedQueryResolver) Delivers(ctx context.Context, offset *int, limit *int, q *string, sort []*DeliverSortType, filter *DeliverFilterType) (*DeliverResultType, error) {
	opts := QueryDeliversHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryDelivers(ctx, r.GeneratedResolver, opts)
}

// QueryDeliversHandler handler
func QueryDeliversHandler(ctx context.Context, r *GeneratedResolver, opts QueryDeliversHandlerOptions) (*DeliverResultType, error) {
	query := DeliverQueryFilter{opts.Q}

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

	return &DeliverResultType{
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

// GeneratedDeliverResultTypeResolver struct
type GeneratedDeliverResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedDeliverResultTypeResolver) Items(ctx context.Context, obj *DeliverResultType) (items []*Deliver, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("delivers"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*Deliver{}
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
func (r *GeneratedDeliverResultTypeResolver) Count(ctx context.Context, obj *DeliverResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("delivers"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &Deliver{})
}

// GeneratedDeliverResolver struct
type GeneratedDeliverResolver struct{ *GeneratedResolver }

// Deliveries ...
func (r *GeneratedDeliverResolver) Deliveries(ctx context.Context, obj *Deliver) (res []*Delivery, err error) {
	return r.Handlers.DeliverDeliveries(ctx, r.GeneratedResolver, obj)
}

// DeliverDeliveriesHandler handler
func DeliverDeliveriesHandler(ctx context.Context, r *GeneratedResolver, obj *Deliver) (res []*Delivery, err error) {

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
func (r *GeneratedDeliverResolver) DeliveriesIds(ctx context.Context, obj *Deliver) (ids []string, err error) {
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
func (r *GeneratedDeliverResolver) DeliveriesConnection(ctx context.Context, obj *Deliver, offset *int, limit *int, q *string, sort []*DeliverySortType, filter *DeliveryFilterType) (res *DeliveryResultType, err error) {
	f := &DeliveryFilterType{
		Deliver: &DeliverFilterType{
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

// DeliveriesSent ...
func (r *GeneratedPersonResolver) DeliveriesSent(ctx context.Context, obj *Person) (res *Delivery, err error) {
	return r.Handlers.PersonDeliveriesSent(ctx, r.GeneratedResolver, obj)
}

// PersonDeliveriesSentHandler handler
func PersonDeliveriesSentHandler(ctx context.Context, r *GeneratedResolver, obj *Person) (res *Delivery, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.DeliveriesSentID != nil {
		item, _err := loaders["Delivery"].Load(ctx, dataloader.StringKey(*obj.DeliveriesSentID))()
		res, _ = item.(*Delivery)

		err = _err
	}

	return
}

// DeliveriesReceived ...
func (r *GeneratedPersonResolver) DeliveriesReceived(ctx context.Context, obj *Person) (res *Delivery, err error) {
	return r.Handlers.PersonDeliveriesReceived(ctx, r.GeneratedResolver, obj)
}

// PersonDeliveriesReceivedHandler handler
func PersonDeliveriesReceivedHandler(ctx context.Context, r *GeneratedResolver, obj *Person) (res *Delivery, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.DeliveriesReceivedID != nil {
		item, _err := loaders["Delivery"].Load(ctx, dataloader.StringKey(*obj.DeliveriesReceivedID))()
		res, _ = item.(*Delivery)

		err = _err
	}

	return
}
