package gen

import (
	"context"

	"github.com/graph-gophers/dataloader"
)

// GetLoaders ...
func GetLoaders(db *DB) map[string]*dataloader.Loader {
	loaders := map[string]*dataloader.Loader{}

	// deliveriesBatchFn batch func
	deliveriesBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Delivery{}
		res := db.Query().Find(items, "id IN (?)", ids)
		if res.Error != nil && !res.RecordNotFound() {
			return []*dataloader.Result{
				&dataloader.Result{Error: res.Error},
			}
		}

		itemMap := make(map[string]Delivery, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Delivery with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["Delivery"] = dataloader.NewBatchedLoader(deliveriesBatchFn, dataloader.WithClearCacheOnBatch())

	// peopleBatchFn batch func
	peopleBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Person{}
		res := db.Query().Find(items, "id IN (?)", ids)
		if res.Error != nil && !res.RecordNotFound() {
			return []*dataloader.Result{
				&dataloader.Result{Error: res.Error},
			}
		}

		itemMap := make(map[string]Person, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Person with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["Person"] = dataloader.NewBatchedLoader(peopleBatchFn, dataloader.WithClearCacheOnBatch())

	// deliveryTypesBatchFn batch func
	deliveryTypesBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]DeliveryType{}
		res := db.Query().Find(items, "id IN (?)", ids)
		if res.Error != nil && !res.RecordNotFound() {
			return []*dataloader.Result{
				&dataloader.Result{Error: res.Error},
			}
		}

		itemMap := make(map[string]DeliveryType, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("DeliveryType with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["DeliveryType"] = dataloader.NewBatchedLoader(deliveryTypesBatchFn, dataloader.WithClearCacheOnBatch())

	// deliveryChannelsBatchFn batch func
	deliveryChannelsBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]DeliveryChannel{}
		res := db.Query().Find(items, "id IN (?)", ids)
		if res.Error != nil && !res.RecordNotFound() {
			return []*dataloader.Result{
				&dataloader.Result{Error: res.Error},
			}
		}

		itemMap := make(map[string]DeliveryChannel, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("DeliveryChannel with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["DeliveryChannel"] = dataloader.NewBatchedLoader(deliveryChannelsBatchFn, dataloader.WithClearCacheOnBatch())

	// vehicleTypesBatchFn batch func
	vehicleTypesBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]VehicleType{}
		res := db.Query().Find(items, "id IN (?)", ids)
		if res.Error != nil && !res.RecordNotFound() {
			return []*dataloader.Result{
				&dataloader.Result{Error: res.Error},
			}
		}

		itemMap := make(map[string]VehicleType, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("VehicleType with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["VehicleType"] = dataloader.NewBatchedLoader(vehicleTypesBatchFn, dataloader.WithClearCacheOnBatch())

	return loaders
}
