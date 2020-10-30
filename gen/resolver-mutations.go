package gen

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/loopcontext/go-graphql-orm/events"
)

// GeneratedMutationResolver struct
type GeneratedMutationResolver struct{ *GeneratedResolver }

// MutationEvents struct
type MutationEvents struct {
	Events []events.Event
}

// EnrichContextWithMutations method
func EnrichContextWithMutations(ctx context.Context, r *GeneratedResolver) context.Context {
	_ctx := context.WithValue(ctx, KeyMutationTransaction, r.DB.db.Begin())
	_ctx = context.WithValue(_ctx, KeyMutationEvents, &MutationEvents{})
	return _ctx
}

// FinishMutationContext method
func FinishMutationContext(ctx context.Context, r *GeneratedResolver) (err error) {
	s := GetMutationEventStore(ctx)

	for _, event := range s.Events {
		err = r.Handlers.OnEvent(ctx, r, &event)
		if err != nil {
			return
		}
	}

	tx := r.GetDB(ctx)
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	for _, event := range s.Events {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}

// RollbackMutationContext method
func RollbackMutationContext(ctx context.Context, r *GeneratedResolver) error {
	tx := r.GetDB(ctx)
	return tx.Rollback().Error
}

// GetMutationEventStore method
func GetMutationEventStore(ctx context.Context) *MutationEvents {
	return ctx.Value(KeyMutationEvents).(*MutationEvents)
}

// AddMutationEvent method
func AddMutationEvent(ctx context.Context, e events.Event) {
	s := GetMutationEventStore(ctx)
	s.Events = append(s.Events, e)
}

// CreateDelivery method
func (r *GeneratedMutationResolver) CreateDelivery(ctx context.Context, input map[string]interface{}) (item *Delivery, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateDelivery(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateDeliveryHandler handler
func CreateDeliveryHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Delivery, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Delivery{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Delivery",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes DeliveryChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["paymentId"]; ok && (item.PaymentID != changes.PaymentID) && (item.PaymentID == nil || changes.PaymentID == nil || *item.PaymentID != *changes.PaymentID) {
		item.PaymentID = changes.PaymentID

		event.AddNewValue("paymentId", changes.PaymentID)
	}

	if _, ok := input["paymentTotal"]; ok && (item.PaymentTotal != changes.PaymentTotal) && (item.PaymentTotal == nil || changes.PaymentTotal == nil || *item.PaymentTotal != *changes.PaymentTotal) {
		item.PaymentTotal = changes.PaymentTotal

		event.AddNewValue("paymentTotal", changes.PaymentTotal)
	}

	if _, ok := input["paymentOnDeliver"]; ok && (item.PaymentOnDeliver != changes.PaymentOnDeliver) && (item.PaymentOnDeliver == nil || changes.PaymentOnDeliver == nil || *item.PaymentOnDeliver != *changes.PaymentOnDeliver) {
		item.PaymentOnDeliver = changes.PaymentOnDeliver

		event.AddNewValue("paymentOnDeliver", changes.PaymentOnDeliver)
	}

	if _, ok := input["collectDateTime"]; ok && (item.CollectDateTime != changes.CollectDateTime) && (item.CollectDateTime == nil || changes.CollectDateTime == nil || *item.CollectDateTime != *changes.CollectDateTime) {
		item.CollectDateTime = changes.CollectDateTime

		event.AddNewValue("collectDateTime", changes.CollectDateTime)
	}

	if _, ok := input["collectAddress"]; ok && (item.CollectAddress != changes.CollectAddress) && (item.CollectAddress == nil || changes.CollectAddress == nil || *item.CollectAddress != *changes.CollectAddress) {
		item.CollectAddress = changes.CollectAddress

		event.AddNewValue("collectAddress", changes.CollectAddress)
	}

	if _, ok := input["collectPoint"]; ok && (item.CollectPoint != changes.CollectPoint) && (item.CollectPoint == nil || changes.CollectPoint == nil || *item.CollectPoint != *changes.CollectPoint) {
		item.CollectPoint = changes.CollectPoint

		event.AddNewValue("collectPoint", changes.CollectPoint)
	}

	if _, ok := input["dropDateTime"]; ok && (item.DropDateTime != changes.DropDateTime) && (item.DropDateTime == nil || changes.DropDateTime == nil || *item.DropDateTime != *changes.DropDateTime) {
		item.DropDateTime = changes.DropDateTime

		event.AddNewValue("dropDateTime", changes.DropDateTime)
	}

	if _, ok := input["dropAddress"]; ok && (item.DropAddress != changes.DropAddress) && (item.DropAddress == nil || changes.DropAddress == nil || *item.DropAddress != *changes.DropAddress) {
		item.DropAddress = changes.DropAddress

		event.AddNewValue("dropAddress", changes.DropAddress)
	}

	if _, ok := input["dropPoint"]; ok && (item.DropPoint != changes.DropPoint) && (item.DropPoint == nil || changes.DropPoint == nil || *item.DropPoint != *changes.DropPoint) {
		item.DropPoint = changes.DropPoint

		event.AddNewValue("dropPoint", changes.DropPoint)
	}

	if _, ok := input["expectedDistance"]; ok && (item.ExpectedDistance != changes.ExpectedDistance) && (item.ExpectedDistance == nil || changes.ExpectedDistance == nil || *item.ExpectedDistance != *changes.ExpectedDistance) {
		item.ExpectedDistance = changes.ExpectedDistance

		event.AddNewValue("expectedDistance", changes.ExpectedDistance)
	}

	if _, ok := input["expectedCost"]; ok && (item.ExpectedCost != changes.ExpectedCost) && (item.ExpectedCost == nil || changes.ExpectedCost == nil || *item.ExpectedCost != *changes.ExpectedCost) {
		item.ExpectedCost = changes.ExpectedCost

		event.AddNewValue("expectedCost", changes.ExpectedCost)
	}

	if _, ok := input["completed"]; ok && (item.Completed != changes.Completed) && (item.Completed == nil || changes.Completed == nil || *item.Completed != *changes.Completed) {
		item.Completed = changes.Completed

		event.AddNewValue("completed", changes.Completed)
	}

	if _, ok := input["smsToken"]; ok && (item.SmsToken != changes.SmsToken) && (item.SmsToken == nil || changes.SmsToken == nil || *item.SmsToken != *changes.SmsToken) {
		item.SmsToken = changes.SmsToken

		event.AddNewValue("smsToken", changes.SmsToken)
	}

	if _, ok := input["status"]; ok && (item.Status != changes.Status) && (item.Status == nil || changes.Status == nil || *item.Status != *changes.Status) {
		item.Status = changes.Status

		event.AddNewValue("status", changes.Status)
	}

	if _, ok := input["instructions"]; ok && (item.Instructions != changes.Instructions) && (item.Instructions == nil || changes.Instructions == nil || *item.Instructions != *changes.Instructions) {
		item.Instructions = changes.Instructions

		event.AddNewValue("instructions", changes.Instructions)
	}

	if _, ok := input["senderId"]; ok && (item.SenderID != changes.SenderID) && (item.SenderID == nil || changes.SenderID == nil || *item.SenderID != *changes.SenderID) {
		item.SenderID = changes.SenderID

		event.AddNewValue("senderId", changes.SenderID)
	}

	if _, ok := input["receiverId"]; ok && (item.ReceiverID != changes.ReceiverID) && (item.ReceiverID == nil || changes.ReceiverID == nil || *item.ReceiverID != *changes.ReceiverID) {
		item.ReceiverID = changes.ReceiverID

		event.AddNewValue("receiverId", changes.ReceiverID)
	}

	if _, ok := input["deliverId"]; ok && (item.DeliverID != changes.DeliverID) && (item.DeliverID == nil || changes.DeliverID == nil || *item.DeliverID != *changes.DeliverID) {
		item.DeliverID = changes.DeliverID

		event.AddNewValue("deliverId", changes.DeliverID)
	}

	if _, ok := input["deliveryTypeId"]; ok && (item.DeliveryTypeID != changes.DeliveryTypeID) && (item.DeliveryTypeID == nil || changes.DeliveryTypeID == nil || *item.DeliveryTypeID != *changes.DeliveryTypeID) {
		item.DeliveryTypeID = changes.DeliveryTypeID

		event.AddNewValue("deliveryTypeId", changes.DeliveryTypeID)
	}

	if _, ok := input["deliveryChannelId"]; ok && (item.DeliveryChannelID != changes.DeliveryChannelID) && (item.DeliveryChannelID == nil || changes.DeliveryChannelID == nil || *item.DeliveryChannelID != *changes.DeliveryChannelID) {
		item.DeliveryChannelID = changes.DeliveryChannelID

		event.AddNewValue("deliveryChannelId", changes.DeliveryChannelID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateDelivery method
func (r *GeneratedMutationResolver) UpdateDelivery(ctx context.Context, id string, input map[string]interface{}) (item *Delivery, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateDelivery(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateDeliveryHandler handler
func UpdateDeliveryHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Delivery, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Delivery{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Delivery",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes DeliveryChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["paymentId"]; ok && (item.PaymentID != changes.PaymentID) && (item.PaymentID == nil || changes.PaymentID == nil || *item.PaymentID != *changes.PaymentID) {
		event.AddOldValue("paymentId", item.PaymentID)
		event.AddNewValue("paymentId", changes.PaymentID)
		item.PaymentID = changes.PaymentID
	}

	if _, ok := input["paymentTotal"]; ok && (item.PaymentTotal != changes.PaymentTotal) && (item.PaymentTotal == nil || changes.PaymentTotal == nil || *item.PaymentTotal != *changes.PaymentTotal) {
		event.AddOldValue("paymentTotal", item.PaymentTotal)
		event.AddNewValue("paymentTotal", changes.PaymentTotal)
		item.PaymentTotal = changes.PaymentTotal
	}

	if _, ok := input["paymentOnDeliver"]; ok && (item.PaymentOnDeliver != changes.PaymentOnDeliver) && (item.PaymentOnDeliver == nil || changes.PaymentOnDeliver == nil || *item.PaymentOnDeliver != *changes.PaymentOnDeliver) {
		event.AddOldValue("paymentOnDeliver", item.PaymentOnDeliver)
		event.AddNewValue("paymentOnDeliver", changes.PaymentOnDeliver)
		item.PaymentOnDeliver = changes.PaymentOnDeliver
	}

	if _, ok := input["collectDateTime"]; ok && (item.CollectDateTime != changes.CollectDateTime) && (item.CollectDateTime == nil || changes.CollectDateTime == nil || *item.CollectDateTime != *changes.CollectDateTime) {
		event.AddOldValue("collectDateTime", item.CollectDateTime)
		event.AddNewValue("collectDateTime", changes.CollectDateTime)
		item.CollectDateTime = changes.CollectDateTime
	}

	if _, ok := input["collectAddress"]; ok && (item.CollectAddress != changes.CollectAddress) && (item.CollectAddress == nil || changes.CollectAddress == nil || *item.CollectAddress != *changes.CollectAddress) {
		event.AddOldValue("collectAddress", item.CollectAddress)
		event.AddNewValue("collectAddress", changes.CollectAddress)
		item.CollectAddress = changes.CollectAddress
	}

	if _, ok := input["collectPoint"]; ok && (item.CollectPoint != changes.CollectPoint) && (item.CollectPoint == nil || changes.CollectPoint == nil || *item.CollectPoint != *changes.CollectPoint) {
		event.AddOldValue("collectPoint", item.CollectPoint)
		event.AddNewValue("collectPoint", changes.CollectPoint)
		item.CollectPoint = changes.CollectPoint
	}

	if _, ok := input["dropDateTime"]; ok && (item.DropDateTime != changes.DropDateTime) && (item.DropDateTime == nil || changes.DropDateTime == nil || *item.DropDateTime != *changes.DropDateTime) {
		event.AddOldValue("dropDateTime", item.DropDateTime)
		event.AddNewValue("dropDateTime", changes.DropDateTime)
		item.DropDateTime = changes.DropDateTime
	}

	if _, ok := input["dropAddress"]; ok && (item.DropAddress != changes.DropAddress) && (item.DropAddress == nil || changes.DropAddress == nil || *item.DropAddress != *changes.DropAddress) {
		event.AddOldValue("dropAddress", item.DropAddress)
		event.AddNewValue("dropAddress", changes.DropAddress)
		item.DropAddress = changes.DropAddress
	}

	if _, ok := input["dropPoint"]; ok && (item.DropPoint != changes.DropPoint) && (item.DropPoint == nil || changes.DropPoint == nil || *item.DropPoint != *changes.DropPoint) {
		event.AddOldValue("dropPoint", item.DropPoint)
		event.AddNewValue("dropPoint", changes.DropPoint)
		item.DropPoint = changes.DropPoint
	}

	if _, ok := input["expectedDistance"]; ok && (item.ExpectedDistance != changes.ExpectedDistance) && (item.ExpectedDistance == nil || changes.ExpectedDistance == nil || *item.ExpectedDistance != *changes.ExpectedDistance) {
		event.AddOldValue("expectedDistance", item.ExpectedDistance)
		event.AddNewValue("expectedDistance", changes.ExpectedDistance)
		item.ExpectedDistance = changes.ExpectedDistance
	}

	if _, ok := input["expectedCost"]; ok && (item.ExpectedCost != changes.ExpectedCost) && (item.ExpectedCost == nil || changes.ExpectedCost == nil || *item.ExpectedCost != *changes.ExpectedCost) {
		event.AddOldValue("expectedCost", item.ExpectedCost)
		event.AddNewValue("expectedCost", changes.ExpectedCost)
		item.ExpectedCost = changes.ExpectedCost
	}

	if _, ok := input["completed"]; ok && (item.Completed != changes.Completed) && (item.Completed == nil || changes.Completed == nil || *item.Completed != *changes.Completed) {
		event.AddOldValue("completed", item.Completed)
		event.AddNewValue("completed", changes.Completed)
		item.Completed = changes.Completed
	}

	if _, ok := input["smsToken"]; ok && (item.SmsToken != changes.SmsToken) && (item.SmsToken == nil || changes.SmsToken == nil || *item.SmsToken != *changes.SmsToken) {
		event.AddOldValue("smsToken", item.SmsToken)
		event.AddNewValue("smsToken", changes.SmsToken)
		item.SmsToken = changes.SmsToken
	}

	if _, ok := input["status"]; ok && (item.Status != changes.Status) && (item.Status == nil || changes.Status == nil || *item.Status != *changes.Status) {
		event.AddOldValue("status", item.Status)
		event.AddNewValue("status", changes.Status)
		item.Status = changes.Status
	}

	if _, ok := input["instructions"]; ok && (item.Instructions != changes.Instructions) && (item.Instructions == nil || changes.Instructions == nil || *item.Instructions != *changes.Instructions) {
		event.AddOldValue("instructions", item.Instructions)
		event.AddNewValue("instructions", changes.Instructions)
		item.Instructions = changes.Instructions
	}

	if _, ok := input["senderId"]; ok && (item.SenderID != changes.SenderID) && (item.SenderID == nil || changes.SenderID == nil || *item.SenderID != *changes.SenderID) {
		event.AddOldValue("senderId", item.SenderID)
		event.AddNewValue("senderId", changes.SenderID)
		item.SenderID = changes.SenderID
	}

	if _, ok := input["receiverId"]; ok && (item.ReceiverID != changes.ReceiverID) && (item.ReceiverID == nil || changes.ReceiverID == nil || *item.ReceiverID != *changes.ReceiverID) {
		event.AddOldValue("receiverId", item.ReceiverID)
		event.AddNewValue("receiverId", changes.ReceiverID)
		item.ReceiverID = changes.ReceiverID
	}

	if _, ok := input["deliverId"]; ok && (item.DeliverID != changes.DeliverID) && (item.DeliverID == nil || changes.DeliverID == nil || *item.DeliverID != *changes.DeliverID) {
		event.AddOldValue("deliverId", item.DeliverID)
		event.AddNewValue("deliverId", changes.DeliverID)
		item.DeliverID = changes.DeliverID
	}

	if _, ok := input["deliveryTypeId"]; ok && (item.DeliveryTypeID != changes.DeliveryTypeID) && (item.DeliveryTypeID == nil || changes.DeliveryTypeID == nil || *item.DeliveryTypeID != *changes.DeliveryTypeID) {
		event.AddOldValue("deliveryTypeId", item.DeliveryTypeID)
		event.AddNewValue("deliveryTypeId", changes.DeliveryTypeID)
		item.DeliveryTypeID = changes.DeliveryTypeID
	}

	if _, ok := input["deliveryChannelId"]; ok && (item.DeliveryChannelID != changes.DeliveryChannelID) && (item.DeliveryChannelID == nil || changes.DeliveryChannelID == nil || *item.DeliveryChannelID != *changes.DeliveryChannelID) {
		event.AddOldValue("deliveryChannelId", item.DeliveryChannelID)
		event.AddNewValue("deliveryChannelId", changes.DeliveryChannelID)
		item.DeliveryChannelID = changes.DeliveryChannelID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteDelivery method
func (r *GeneratedMutationResolver) DeleteDelivery(ctx context.Context, id string) (item *Delivery, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteDelivery(ctx, r.GeneratedResolver, id)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteDeliveryHandler handler
func DeleteDeliveryHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Delivery, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Delivery{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Delivery",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("deliveries")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllDeliveries method
func (r *GeneratedMutationResolver) DeleteAllDeliveries(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllDeliveries(ctx, r.GeneratedResolver)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllDeliveriesHandler handler
func DeleteAllDeliveriesHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := r.GetDB(ctx)
	err := tx.Delete(&Delivery{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreatePerson method
func (r *GeneratedMutationResolver) CreatePerson(ctx context.Context, input map[string]interface{}) (item *Person, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreatePerson(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreatePersonHandler handler
func CreatePersonHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Person, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Person{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Person",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes PersonChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["deliver"]; ok && (item.Deliver != changes.Deliver) && (item.Deliver == nil || changes.Deliver == nil || *item.Deliver != *changes.Deliver) {
		item.Deliver = changes.Deliver

		event.AddNewValue("deliver", changes.Deliver)
	}

	if _, ok := input["email"]; ok && (item.Email != changes.Email) {
		item.Email = changes.Email

		event.AddNewValue("email", changes.Email)
	}

	if _, ok := input["phone"]; ok && (item.Phone != changes.Phone) && (item.Phone == nil || changes.Phone == nil || *item.Phone != *changes.Phone) {
		item.Phone = changes.Phone

		event.AddNewValue("phone", changes.Phone)
	}

	if _, ok := input["documentNo"]; ok && (item.DocumentNo != changes.DocumentNo) && (item.DocumentNo == nil || changes.DocumentNo == nil || *item.DocumentNo != *changes.DocumentNo) {
		item.DocumentNo = changes.DocumentNo

		event.AddNewValue("documentNo", changes.DocumentNo)
	}

	if _, ok := input["avatarURL"]; ok && (item.AvatarURL != changes.AvatarURL) && (item.AvatarURL == nil || changes.AvatarURL == nil || *item.AvatarURL != *changes.AvatarURL) {
		item.AvatarURL = changes.AvatarURL

		event.AddNewValue("avatarURL", changes.AvatarURL)
	}

	if _, ok := input["displayName"]; ok && (item.DisplayName != changes.DisplayName) && (item.DisplayName == nil || changes.DisplayName == nil || *item.DisplayName != *changes.DisplayName) {
		item.DisplayName = changes.DisplayName

		event.AddNewValue("displayName", changes.DisplayName)
	}

	if _, ok := input["firstName"]; ok && (item.FirstName != changes.FirstName) && (item.FirstName == nil || changes.FirstName == nil || *item.FirstName != *changes.FirstName) {
		item.FirstName = changes.FirstName

		event.AddNewValue("firstName", changes.FirstName)
	}

	if _, ok := input["lastName"]; ok && (item.LastName != changes.LastName) && (item.LastName == nil || changes.LastName == nil || *item.LastName != *changes.LastName) {
		item.LastName = changes.LastName

		event.AddNewValue("lastName", changes.LastName)
	}

	if _, ok := input["nickName"]; ok && (item.NickName != changes.NickName) && (item.NickName == nil || changes.NickName == nil || *item.NickName != *changes.NickName) {
		item.NickName = changes.NickName

		event.AddNewValue("nickName", changes.NickName)
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		item.Description = changes.Description

		event.AddNewValue("description", changes.Description)
	}

	if _, ok := input["location"]; ok && (item.Location != changes.Location) && (item.Location == nil || changes.Location == nil || *item.Location != *changes.Location) {
		item.Location = changes.Location

		event.AddNewValue("location", changes.Location)
	}

	if _, ok := input["userId"]; ok && (item.UserID != changes.UserID) && (item.UserID == nil || changes.UserID == nil || *item.UserID != *changes.UserID) {
		item.UserID = changes.UserID

		event.AddNewValue("userId", changes.UserID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["deliveriesIds"]; exists {
		items := []Delivery{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Deliveries")
		association.Replace(items)
	}

	if ids, exists := input["deliveriesSentIds"]; exists {
		items := []Delivery{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("DeliveriesSent")
		association.Replace(items)
	}

	if ids, exists := input["deliveriesReceivedIds"]; exists {
		items := []Delivery{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("DeliveriesReceived")
		association.Replace(items)
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdatePerson method
func (r *GeneratedMutationResolver) UpdatePerson(ctx context.Context, id string, input map[string]interface{}) (item *Person, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdatePerson(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdatePersonHandler handler
func UpdatePersonHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Person, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Person{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Person",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes PersonChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["deliver"]; ok && (item.Deliver != changes.Deliver) && (item.Deliver == nil || changes.Deliver == nil || *item.Deliver != *changes.Deliver) {
		event.AddOldValue("deliver", item.Deliver)
		event.AddNewValue("deliver", changes.Deliver)
		item.Deliver = changes.Deliver
	}

	if _, ok := input["email"]; ok && (item.Email != changes.Email) {
		event.AddOldValue("email", item.Email)
		event.AddNewValue("email", changes.Email)
		item.Email = changes.Email
	}

	if _, ok := input["phone"]; ok && (item.Phone != changes.Phone) && (item.Phone == nil || changes.Phone == nil || *item.Phone != *changes.Phone) {
		event.AddOldValue("phone", item.Phone)
		event.AddNewValue("phone", changes.Phone)
		item.Phone = changes.Phone
	}

	if _, ok := input["documentNo"]; ok && (item.DocumentNo != changes.DocumentNo) && (item.DocumentNo == nil || changes.DocumentNo == nil || *item.DocumentNo != *changes.DocumentNo) {
		event.AddOldValue("documentNo", item.DocumentNo)
		event.AddNewValue("documentNo", changes.DocumentNo)
		item.DocumentNo = changes.DocumentNo
	}

	if _, ok := input["avatarURL"]; ok && (item.AvatarURL != changes.AvatarURL) && (item.AvatarURL == nil || changes.AvatarURL == nil || *item.AvatarURL != *changes.AvatarURL) {
		event.AddOldValue("avatarURL", item.AvatarURL)
		event.AddNewValue("avatarURL", changes.AvatarURL)
		item.AvatarURL = changes.AvatarURL
	}

	if _, ok := input["displayName"]; ok && (item.DisplayName != changes.DisplayName) && (item.DisplayName == nil || changes.DisplayName == nil || *item.DisplayName != *changes.DisplayName) {
		event.AddOldValue("displayName", item.DisplayName)
		event.AddNewValue("displayName", changes.DisplayName)
		item.DisplayName = changes.DisplayName
	}

	if _, ok := input["firstName"]; ok && (item.FirstName != changes.FirstName) && (item.FirstName == nil || changes.FirstName == nil || *item.FirstName != *changes.FirstName) {
		event.AddOldValue("firstName", item.FirstName)
		event.AddNewValue("firstName", changes.FirstName)
		item.FirstName = changes.FirstName
	}

	if _, ok := input["lastName"]; ok && (item.LastName != changes.LastName) && (item.LastName == nil || changes.LastName == nil || *item.LastName != *changes.LastName) {
		event.AddOldValue("lastName", item.LastName)
		event.AddNewValue("lastName", changes.LastName)
		item.LastName = changes.LastName
	}

	if _, ok := input["nickName"]; ok && (item.NickName != changes.NickName) && (item.NickName == nil || changes.NickName == nil || *item.NickName != *changes.NickName) {
		event.AddOldValue("nickName", item.NickName)
		event.AddNewValue("nickName", changes.NickName)
		item.NickName = changes.NickName
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		event.AddOldValue("description", item.Description)
		event.AddNewValue("description", changes.Description)
		item.Description = changes.Description
	}

	if _, ok := input["location"]; ok && (item.Location != changes.Location) && (item.Location == nil || changes.Location == nil || *item.Location != *changes.Location) {
		event.AddOldValue("location", item.Location)
		event.AddNewValue("location", changes.Location)
		item.Location = changes.Location
	}

	if _, ok := input["userId"]; ok && (item.UserID != changes.UserID) && (item.UserID == nil || changes.UserID == nil || *item.UserID != *changes.UserID) {
		event.AddOldValue("userId", item.UserID)
		event.AddNewValue("userId", changes.UserID)
		item.UserID = changes.UserID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["deliveriesIds"]; exists {
		items := []Delivery{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Deliveries")
		association.Replace(items)
	}

	if ids, exists := input["deliveriesSentIds"]; exists {
		items := []Delivery{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("DeliveriesSent")
		association.Replace(items)
	}

	if ids, exists := input["deliveriesReceivedIds"]; exists {
		items := []Delivery{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("DeliveriesReceived")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeletePerson method
func (r *GeneratedMutationResolver) DeletePerson(ctx context.Context, id string) (item *Person, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeletePerson(ctx, r.GeneratedResolver, id)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeletePersonHandler handler
func DeletePersonHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Person, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Person{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Person",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("people")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllPeople method
func (r *GeneratedMutationResolver) DeleteAllPeople(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllPeople(ctx, r.GeneratedResolver)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllPeopleHandler handler
func DeleteAllPeopleHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := r.GetDB(ctx)
	err := tx.Delete(&Person{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreateDeliveryType method
func (r *GeneratedMutationResolver) CreateDeliveryType(ctx context.Context, input map[string]interface{}) (item *DeliveryType, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateDeliveryType(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateDeliveryTypeHandler handler
func CreateDeliveryTypeHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *DeliveryType, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &DeliveryType{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "DeliveryType",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes DeliveryTypeChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		item.Description = changes.Description

		event.AddNewValue("description", changes.Description)
	}

	if _, ok := input["deliveryId"]; ok && (item.DeliveryID != changes.DeliveryID) && (item.DeliveryID == nil || changes.DeliveryID == nil || *item.DeliveryID != *changes.DeliveryID) {
		item.DeliveryID = changes.DeliveryID

		event.AddNewValue("deliveryId", changes.DeliveryID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateDeliveryType method
func (r *GeneratedMutationResolver) UpdateDeliveryType(ctx context.Context, id string, input map[string]interface{}) (item *DeliveryType, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateDeliveryType(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateDeliveryTypeHandler handler
func UpdateDeliveryTypeHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *DeliveryType, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &DeliveryType{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "DeliveryType",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes DeliveryTypeChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		event.AddOldValue("description", item.Description)
		event.AddNewValue("description", changes.Description)
		item.Description = changes.Description
	}

	if _, ok := input["deliveryId"]; ok && (item.DeliveryID != changes.DeliveryID) && (item.DeliveryID == nil || changes.DeliveryID == nil || *item.DeliveryID != *changes.DeliveryID) {
		event.AddOldValue("deliveryId", item.DeliveryID)
		event.AddNewValue("deliveryId", changes.DeliveryID)
		item.DeliveryID = changes.DeliveryID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteDeliveryType method
func (r *GeneratedMutationResolver) DeleteDeliveryType(ctx context.Context, id string) (item *DeliveryType, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteDeliveryType(ctx, r.GeneratedResolver, id)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteDeliveryTypeHandler handler
func DeleteDeliveryTypeHandler(ctx context.Context, r *GeneratedResolver, id string) (item *DeliveryType, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &DeliveryType{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "DeliveryType",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("delivery_types")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllDeliveryTypes method
func (r *GeneratedMutationResolver) DeleteAllDeliveryTypes(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllDeliveryTypes(ctx, r.GeneratedResolver)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllDeliveryTypesHandler handler
func DeleteAllDeliveryTypesHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := r.GetDB(ctx)
	err := tx.Delete(&DeliveryType{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreateDeliveryChannel method
func (r *GeneratedMutationResolver) CreateDeliveryChannel(ctx context.Context, input map[string]interface{}) (item *DeliveryChannel, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateDeliveryChannel(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateDeliveryChannelHandler handler
func CreateDeliveryChannelHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *DeliveryChannel, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &DeliveryChannel{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "DeliveryChannel",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes DeliveryChannelChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		item.Description = changes.Description

		event.AddNewValue("description", changes.Description)
	}

	if _, ok := input["deliveryId"]; ok && (item.DeliveryID != changes.DeliveryID) && (item.DeliveryID == nil || changes.DeliveryID == nil || *item.DeliveryID != *changes.DeliveryID) {
		item.DeliveryID = changes.DeliveryID

		event.AddNewValue("deliveryId", changes.DeliveryID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateDeliveryChannel method
func (r *GeneratedMutationResolver) UpdateDeliveryChannel(ctx context.Context, id string, input map[string]interface{}) (item *DeliveryChannel, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateDeliveryChannel(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateDeliveryChannelHandler handler
func UpdateDeliveryChannelHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *DeliveryChannel, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &DeliveryChannel{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "DeliveryChannel",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes DeliveryChannelChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		event.AddOldValue("description", item.Description)
		event.AddNewValue("description", changes.Description)
		item.Description = changes.Description
	}

	if _, ok := input["deliveryId"]; ok && (item.DeliveryID != changes.DeliveryID) && (item.DeliveryID == nil || changes.DeliveryID == nil || *item.DeliveryID != *changes.DeliveryID) {
		event.AddOldValue("deliveryId", item.DeliveryID)
		event.AddNewValue("deliveryId", changes.DeliveryID)
		item.DeliveryID = changes.DeliveryID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteDeliveryChannel method
func (r *GeneratedMutationResolver) DeleteDeliveryChannel(ctx context.Context, id string) (item *DeliveryChannel, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteDeliveryChannel(ctx, r.GeneratedResolver, id)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteDeliveryChannelHandler handler
func DeleteDeliveryChannelHandler(ctx context.Context, r *GeneratedResolver, id string) (item *DeliveryChannel, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &DeliveryChannel{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "DeliveryChannel",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("delivery_channels")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllDeliveryChannels method
func (r *GeneratedMutationResolver) DeleteAllDeliveryChannels(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllDeliveryChannels(ctx, r.GeneratedResolver)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllDeliveryChannelsHandler handler
func DeleteAllDeliveryChannelsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := r.GetDB(ctx)
	err := tx.Delete(&DeliveryChannel{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreateVehicleType method
func (r *GeneratedMutationResolver) CreateVehicleType(ctx context.Context, input map[string]interface{}) (item *VehicleType, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateVehicleType(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateVehicleTypeHandler handler
func CreateVehicleTypeHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *VehicleType, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &VehicleType{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "VehicleType",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes VehicleTypeChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		item.Description = changes.Description

		event.AddNewValue("description", changes.Description)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateVehicleType method
func (r *GeneratedMutationResolver) UpdateVehicleType(ctx context.Context, id string, input map[string]interface{}) (item *VehicleType, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateVehicleType(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateVehicleTypeHandler handler
func UpdateVehicleTypeHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *VehicleType, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &VehicleType{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "VehicleType",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes VehicleTypeChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		event.AddOldValue("description", item.Description)
		event.AddNewValue("description", changes.Description)
		item.Description = changes.Description
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteVehicleType method
func (r *GeneratedMutationResolver) DeleteVehicleType(ctx context.Context, id string) (item *VehicleType, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteVehicleType(ctx, r.GeneratedResolver, id)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteVehicleTypeHandler handler
func DeleteVehicleTypeHandler(ctx context.Context, r *GeneratedResolver, id string) (item *VehicleType, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &VehicleType{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "VehicleType",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("vehicle_types")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllVehicleTypes method
func (r *GeneratedMutationResolver) DeleteAllVehicleTypes(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllVehicleTypes(ctx, r.GeneratedResolver)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllVehicleTypesHandler handler
func DeleteAllVehicleTypesHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := r.GetDB(ctx)
	err := tx.Delete(&VehicleType{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}
