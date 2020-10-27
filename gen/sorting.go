package gen

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Apply method
func (s DeliverySortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("deliveries"), sorts, joins)
}

// ApplyWithAlias method
func (s DeliverySortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CollectDateTime != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("collectDateTime"), Direction: s.CollectDateTime.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CollectDateTimeMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("collectDateTime") + ")", Direction: s.CollectDateTimeMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CollectDateTimeMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("collectDateTime") + ")", Direction: s.CollectDateTimeMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CollectAddress != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("collectAddress"), Direction: s.CollectAddress.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CollectAddressMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("collectAddress") + ")", Direction: s.CollectAddressMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CollectAddressMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("collectAddress") + ")", Direction: s.CollectAddressMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CollectPoint != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("collectPoint"), Direction: s.CollectPoint.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CollectPointMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("collectPoint") + ")", Direction: s.CollectPointMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CollectPointMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("collectPoint") + ")", Direction: s.CollectPointMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DropDateTime != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("dropDateTime"), Direction: s.DropDateTime.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DropDateTimeMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("dropDateTime") + ")", Direction: s.DropDateTimeMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DropDateTimeMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("dropDateTime") + ")", Direction: s.DropDateTimeMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DropAddress != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("dropAddress"), Direction: s.DropAddress.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DropAddressMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("dropAddress") + ")", Direction: s.DropAddressMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DropAddressMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("dropAddress") + ")", Direction: s.DropAddressMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DropPoint != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("dropPoint"), Direction: s.DropPoint.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DropPointMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("dropPoint") + ")", Direction: s.DropPointMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DropPointMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("dropPoint") + ")", Direction: s.DropPointMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentTotal != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("paymentTotal"), Direction: s.PaymentTotal.String()}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentTotalMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("paymentTotal") + ")", Direction: s.PaymentTotalMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentTotalMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("paymentTotal") + ")", Direction: s.PaymentTotalMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentTotalAvg != nil {
		sort := SortInfo{Field: "Avg(" + aliasPrefix + dialect.Quote("paymentTotal") + ")", Direction: s.PaymentTotalAvg.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentOnDeliver != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("paymentOnDeliver"), Direction: s.PaymentOnDeliver.String()}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentOnDeliverMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("paymentOnDeliver") + ")", Direction: s.PaymentOnDeliverMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentOnDeliverMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("paymentOnDeliver") + ")", Direction: s.PaymentOnDeliverMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ExpectedDistance != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("expectedDistance"), Direction: s.ExpectedDistance.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ExpectedDistanceMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("expectedDistance") + ")", Direction: s.ExpectedDistanceMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ExpectedDistanceMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("expectedDistance") + ")", Direction: s.ExpectedDistanceMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ExpectedCost != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("expectedCost"), Direction: s.ExpectedCost.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ExpectedCostMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("expectedCost") + ")", Direction: s.ExpectedCostMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ExpectedCostMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("expectedCost") + ")", Direction: s.ExpectedCostMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ExpectedCostAvg != nil {
		sort := SortInfo{Field: "Avg(" + aliasPrefix + dialect.Quote("expectedCost") + ")", Direction: s.ExpectedCostAvg.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Completed != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("completed"), Direction: s.Completed.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CompletedMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("completed") + ")", Direction: s.CompletedMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CompletedMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("completed") + ")", Direction: s.CompletedMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.SmsToken != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("smsToken"), Direction: s.SmsToken.String()}
		*sorts = append(*sorts, sort)
	}

	if s.SmsTokenMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("smsToken") + ")", Direction: s.SmsTokenMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.SmsTokenMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("smsToken") + ")", Direction: s.SmsTokenMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Status != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("status"), Direction: s.Status.String()}
		*sorts = append(*sorts, sort)
	}

	if s.StatusMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("status") + ")", Direction: s.StatusMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.StatusMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("status") + ")", Direction: s.StatusMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Instructions != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("instructions"), Direction: s.Instructions.String()}
		*sorts = append(*sorts, sort)
	}

	if s.InstructionsMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("instructions") + ")", Direction: s.InstructionsMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.InstructionsMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("instructions") + ")", Direction: s.InstructionsMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.SenderID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("senderId"), Direction: s.SenderID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.SenderIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("senderId") + ")", Direction: s.SenderIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.SenderIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("senderId") + ")", Direction: s.SenderIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ReceiverID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("receiverId"), Direction: s.ReceiverID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ReceiverIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("receiverId") + ")", Direction: s.ReceiverIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ReceiverIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("receiverId") + ")", Direction: s.ReceiverIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DeliverID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("deliverId"), Direction: s.DeliverID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DeliverIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("deliverId") + ")", Direction: s.DeliverIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DeliverIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("deliverId") + ")", Direction: s.DeliverIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Sender != nil {
		_alias := alias + "_sender"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("people"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("senderId"))
		err := s.Sender.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Receiver != nil {
		_alias := alias + "_receiver"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("people"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("receiverId"))
		err := s.Receiver.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Deliver != nil {
		_alias := alias + "_deliver"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("people"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("deliverId"))
		err := s.Deliver.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply method
func (s PersonSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("people"), sorts, joins)
}

// ApplyWithAlias method
func (s PersonSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Deliver != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("deliver"), Direction: s.Deliver.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DeliverMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("deliver") + ")", Direction: s.DeliverMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DeliverMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("deliver") + ")", Direction: s.DeliverMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Email != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("email"), Direction: s.Email.String()}
		*sorts = append(*sorts, sort)
	}

	if s.EmailMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("email") + ")", Direction: s.EmailMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.EmailMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("email") + ")", Direction: s.EmailMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Phone != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("phone"), Direction: s.Phone.String()}
		*sorts = append(*sorts, sort)
	}

	if s.PhoneMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("phone") + ")", Direction: s.PhoneMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PhoneMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("phone") + ")", Direction: s.PhoneMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DocumentNo != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("documentNo"), Direction: s.DocumentNo.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DocumentNoMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("documentNo") + ")", Direction: s.DocumentNoMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DocumentNoMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("documentNo") + ")", Direction: s.DocumentNoMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AvatarURL != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("avatarURL"), Direction: s.AvatarURL.String()}
		*sorts = append(*sorts, sort)
	}

	if s.AvatarURLMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("avatarURL") + ")", Direction: s.AvatarURLMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AvatarURLMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("avatarURL") + ")", Direction: s.AvatarURLMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DisplayName != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("displayName"), Direction: s.DisplayName.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DisplayNameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("displayName") + ")", Direction: s.DisplayNameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DisplayNameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("displayName") + ")", Direction: s.DisplayNameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.FirstName != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("firstName"), Direction: s.FirstName.String()}
		*sorts = append(*sorts, sort)
	}

	if s.FirstNameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("firstName") + ")", Direction: s.FirstNameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.FirstNameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("firstName") + ")", Direction: s.FirstNameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.LastName != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("lastName"), Direction: s.LastName.String()}
		*sorts = append(*sorts, sort)
	}

	if s.LastNameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("lastName") + ")", Direction: s.LastNameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.LastNameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("lastName") + ")", Direction: s.LastNameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NickName != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("nickName"), Direction: s.NickName.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NickNameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("nickName") + ")", Direction: s.NickNameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NickNameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("nickName") + ")", Direction: s.NickNameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Description != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("description"), Direction: s.Description.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Location != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("location"), Direction: s.Location.String()}
		*sorts = append(*sorts, sort)
	}

	if s.LocationMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("location") + ")", Direction: s.LocationMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.LocationMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("location") + ")", Direction: s.LocationMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UserID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("userId"), Direction: s.UserID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UserIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("userId") + ")", Direction: s.UserIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UserIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("userId") + ")", Direction: s.UserIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentStatusID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("paymentStatusId"), Direction: s.PaymentStatusID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentStatusIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("paymentStatusId") + ")", Direction: s.PaymentStatusIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentStatusIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("paymentStatusId") + ")", Direction: s.PaymentStatusIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentHistoryID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("paymentHistoryId"), Direction: s.PaymentHistoryID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentHistoryIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("paymentHistoryId") + ")", Direction: s.PaymentHistoryIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentHistoryIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("paymentHistoryId") + ")", Direction: s.PaymentHistoryIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Deliveries != nil {
		_alias := alias + "_deliveries"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("deliveries"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("deliverId")+" = "+dialect.Quote(alias)+".id")
		err := s.Deliveries.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.DeliveriesSent != nil {
		_alias := alias + "_deliveriesSent"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("deliveries"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("senderId")+" = "+dialect.Quote(alias)+".id")
		err := s.DeliveriesSent.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.DeliveriesReceived != nil {
		_alias := alias + "_deliveriesReceived"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("deliveries"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("receiverId")+" = "+dialect.Quote(alias)+".id")
		err := s.DeliveriesReceived.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.PaymentStatus != nil {
		_alias := alias + "_paymentStatus"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("payment_statuses"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("paymentStatusId"))
		err := s.PaymentStatus.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.PaymentHistory != nil {
		_alias := alias + "_paymentHistory"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("payment_histories"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("paymentHistoryId"))
		err := s.PaymentHistory.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply method
func (s DeliveryTypeSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("delivery_types"), sorts, joins)
}

// ApplyWithAlias method
func (s DeliveryTypeSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Description != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("description"), Direction: s.Description.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	return nil
}

// Apply method
func (s DeliveryChannelSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("delivery_channels"), sorts, joins)
}

// ApplyWithAlias method
func (s DeliveryChannelSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Description != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("description"), Direction: s.Description.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	return nil
}

// Apply method
func (s VehicleTypeSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("vehicle_types"), sorts, joins)
}

// ApplyWithAlias method
func (s VehicleTypeSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Description != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("description"), Direction: s.Description.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	return nil
}

// Apply method
func (s PaymentChannelSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("payment_channels"), sorts, joins)
}

// ApplyWithAlias method
func (s PaymentChannelSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Description != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("description"), Direction: s.Description.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	return nil
}

// Apply method
func (s PaymentStatusSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("payment_statuses"), sorts, joins)
}

// ApplyWithAlias method
func (s PaymentStatusSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Amount != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("amount"), Direction: s.Amount.String()}
		*sorts = append(*sorts, sort)
	}

	if s.AmountMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("amount") + ")", Direction: s.AmountMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AmountMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("amount") + ")", Direction: s.AmountMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AmountAvg != nil {
		sort := SortInfo{Field: "Avg(" + aliasPrefix + dialect.Quote("amount") + ")", Direction: s.AmountAvg.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PersonID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("personId"), Direction: s.PersonID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.PersonIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("personId") + ")", Direction: s.PersonIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PersonIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("personId") + ")", Direction: s.PersonIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Person != nil {
		_alias := alias + "_person"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("people"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("personId"))
		err := s.Person.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply method
func (s PaymentHistorySortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("payment_histories"), sorts, joins)
}

// ApplyWithAlias method
func (s PaymentHistorySortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Amount != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("amount"), Direction: s.Amount.String()}
		*sorts = append(*sorts, sort)
	}

	if s.AmountMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("amount") + ")", Direction: s.AmountMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AmountMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("amount") + ")", Direction: s.AmountMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AmountAvg != nil {
		sort := SortInfo{Field: "Avg(" + aliasPrefix + dialect.Quote("amount") + ")", Direction: s.AmountAvg.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Concept != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("concept"), Direction: s.Concept.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ConceptMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("concept") + ")", Direction: s.ConceptMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ConceptMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("concept") + ")", Direction: s.ConceptMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PersonID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("personId"), Direction: s.PersonID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.PersonIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("personId") + ")", Direction: s.PersonIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PersonIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("personId") + ")", Direction: s.PersonIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Person != nil {
		_alias := alias + "_person"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("people"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("personId"))
		err := s.Person.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
