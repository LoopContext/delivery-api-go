package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// IsEmpty ...
func (f *DeliveryFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *DeliveryFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("deliveries"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *DeliveryFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Sender != nil {
		_alias := alias + "_sender"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("people"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("senderId"))
		err := f.Sender.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Receiver != nil {
		_alias := alias + "_receiver"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("people"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("receiverId"))
		err := f.Receiver.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Deliver != nil {
		_alias := alias + "_deliver"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("people"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("deliverId"))
		err := f.Deliver.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.VehicleType != nil {
		_alias := alias + "_vehicleType"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("vehicle_types"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("vehicleTypeId"))
		err := f.VehicleType.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.DeliveryType != nil {
		_alias := alias + "_deliveryType"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("delivery_types"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("deliveryTypeId"))
		err := f.DeliveryType.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.DeliveryChannel != nil {
		_alias := alias + "_deliveryChannel"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("delivery_channels"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("deliveryChannelId"))
		err := f.DeliveryChannel.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *DeliveryFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.PaymentID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" = ?")
		values = append(values, f.PaymentID)
	}

	if f.PaymentIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" != ?")
		values = append(values, f.PaymentIDNe)
	}

	if f.PaymentIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" > ?")
		values = append(values, f.PaymentIDGt)
	}

	if f.PaymentIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" < ?")
		values = append(values, f.PaymentIDLt)
	}

	if f.PaymentIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" >= ?")
		values = append(values, f.PaymentIDGte)
	}

	if f.PaymentIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" <= ?")
		values = append(values, f.PaymentIDLte)
	}

	if f.PaymentIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" IN (?)")
		values = append(values, f.PaymentIDIn)
	}

	if f.PaymentIDNull != nil {
		if *f.PaymentIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" IS NOT NULL")
		}
	}

	if f.PaymentTotal != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentTotal")+" = ?")
		values = append(values, f.PaymentTotal)
	}

	if f.PaymentTotalNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentTotal")+" != ?")
		values = append(values, f.PaymentTotalNe)
	}

	if f.PaymentTotalGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentTotal")+" > ?")
		values = append(values, f.PaymentTotalGt)
	}

	if f.PaymentTotalLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentTotal")+" < ?")
		values = append(values, f.PaymentTotalLt)
	}

	if f.PaymentTotalGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentTotal")+" >= ?")
		values = append(values, f.PaymentTotalGte)
	}

	if f.PaymentTotalLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentTotal")+" <= ?")
		values = append(values, f.PaymentTotalLte)
	}

	if f.PaymentTotalIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentTotal")+" IN (?)")
		values = append(values, f.PaymentTotalIn)
	}

	if f.PaymentTotalNull != nil {
		if *f.PaymentTotalNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("paymentTotal")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("paymentTotal")+" IS NOT NULL")
		}
	}

	if f.PaymentOnDeliver != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentOnDeliver")+" = ?")
		values = append(values, f.PaymentOnDeliver)
	}

	if f.PaymentOnDeliverNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentOnDeliver")+" != ?")
		values = append(values, f.PaymentOnDeliverNe)
	}

	if f.PaymentOnDeliverGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentOnDeliver")+" > ?")
		values = append(values, f.PaymentOnDeliverGt)
	}

	if f.PaymentOnDeliverLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentOnDeliver")+" < ?")
		values = append(values, f.PaymentOnDeliverLt)
	}

	if f.PaymentOnDeliverGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentOnDeliver")+" >= ?")
		values = append(values, f.PaymentOnDeliverGte)
	}

	if f.PaymentOnDeliverLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentOnDeliver")+" <= ?")
		values = append(values, f.PaymentOnDeliverLte)
	}

	if f.PaymentOnDeliverIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentOnDeliver")+" IN (?)")
		values = append(values, f.PaymentOnDeliverIn)
	}

	if f.PaymentOnDeliverNull != nil {
		if *f.PaymentOnDeliverNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("paymentOnDeliver")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("paymentOnDeliver")+" IS NOT NULL")
		}
	}

	if f.CollectDateTime != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectDateTime")+" = ?")
		values = append(values, f.CollectDateTime)
	}

	if f.CollectDateTimeNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectDateTime")+" != ?")
		values = append(values, f.CollectDateTimeNe)
	}

	if f.CollectDateTimeGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectDateTime")+" > ?")
		values = append(values, f.CollectDateTimeGt)
	}

	if f.CollectDateTimeLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectDateTime")+" < ?")
		values = append(values, f.CollectDateTimeLt)
	}

	if f.CollectDateTimeGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectDateTime")+" >= ?")
		values = append(values, f.CollectDateTimeGte)
	}

	if f.CollectDateTimeLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectDateTime")+" <= ?")
		values = append(values, f.CollectDateTimeLte)
	}

	if f.CollectDateTimeIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectDateTime")+" IN (?)")
		values = append(values, f.CollectDateTimeIn)
	}

	if f.CollectDateTimeNull != nil {
		if *f.CollectDateTimeNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("collectDateTime")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("collectDateTime")+" IS NOT NULL")
		}
	}

	if f.CollectAddress != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectAddress")+" = ?")
		values = append(values, f.CollectAddress)
	}

	if f.CollectAddressNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectAddress")+" != ?")
		values = append(values, f.CollectAddressNe)
	}

	if f.CollectAddressGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectAddress")+" > ?")
		values = append(values, f.CollectAddressGt)
	}

	if f.CollectAddressLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectAddress")+" < ?")
		values = append(values, f.CollectAddressLt)
	}

	if f.CollectAddressGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectAddress")+" >= ?")
		values = append(values, f.CollectAddressGte)
	}

	if f.CollectAddressLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectAddress")+" <= ?")
		values = append(values, f.CollectAddressLte)
	}

	if f.CollectAddressIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectAddress")+" IN (?)")
		values = append(values, f.CollectAddressIn)
	}

	if f.CollectAddressLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectAddress")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.CollectAddressLike, "?", "_", -1), "*", "%", -1))
	}

	if f.CollectAddressPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectAddress")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.CollectAddressPrefix))
	}

	if f.CollectAddressSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectAddress")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.CollectAddressSuffix))
	}

	if f.CollectAddressNull != nil {
		if *f.CollectAddressNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("collectAddress")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("collectAddress")+" IS NOT NULL")
		}
	}

	if f.CollectPoint != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectPoint")+" = ?")
		values = append(values, f.CollectPoint)
	}

	if f.CollectPointNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectPoint")+" != ?")
		values = append(values, f.CollectPointNe)
	}

	if f.CollectPointGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectPoint")+" > ?")
		values = append(values, f.CollectPointGt)
	}

	if f.CollectPointLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectPoint")+" < ?")
		values = append(values, f.CollectPointLt)
	}

	if f.CollectPointGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectPoint")+" >= ?")
		values = append(values, f.CollectPointGte)
	}

	if f.CollectPointLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectPoint")+" <= ?")
		values = append(values, f.CollectPointLte)
	}

	if f.CollectPointIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectPoint")+" IN (?)")
		values = append(values, f.CollectPointIn)
	}

	if f.CollectPointLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectPoint")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.CollectPointLike, "?", "_", -1), "*", "%", -1))
	}

	if f.CollectPointPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectPoint")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.CollectPointPrefix))
	}

	if f.CollectPointSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("collectPoint")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.CollectPointSuffix))
	}

	if f.CollectPointNull != nil {
		if *f.CollectPointNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("collectPoint")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("collectPoint")+" IS NOT NULL")
		}
	}

	if f.DropDateTime != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropDateTime")+" = ?")
		values = append(values, f.DropDateTime)
	}

	if f.DropDateTimeNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropDateTime")+" != ?")
		values = append(values, f.DropDateTimeNe)
	}

	if f.DropDateTimeGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropDateTime")+" > ?")
		values = append(values, f.DropDateTimeGt)
	}

	if f.DropDateTimeLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropDateTime")+" < ?")
		values = append(values, f.DropDateTimeLt)
	}

	if f.DropDateTimeGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropDateTime")+" >= ?")
		values = append(values, f.DropDateTimeGte)
	}

	if f.DropDateTimeLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropDateTime")+" <= ?")
		values = append(values, f.DropDateTimeLte)
	}

	if f.DropDateTimeIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropDateTime")+" IN (?)")
		values = append(values, f.DropDateTimeIn)
	}

	if f.DropDateTimeNull != nil {
		if *f.DropDateTimeNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("dropDateTime")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("dropDateTime")+" IS NOT NULL")
		}
	}

	if f.DropAddress != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropAddress")+" = ?")
		values = append(values, f.DropAddress)
	}

	if f.DropAddressNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropAddress")+" != ?")
		values = append(values, f.DropAddressNe)
	}

	if f.DropAddressGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropAddress")+" > ?")
		values = append(values, f.DropAddressGt)
	}

	if f.DropAddressLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropAddress")+" < ?")
		values = append(values, f.DropAddressLt)
	}

	if f.DropAddressGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropAddress")+" >= ?")
		values = append(values, f.DropAddressGte)
	}

	if f.DropAddressLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropAddress")+" <= ?")
		values = append(values, f.DropAddressLte)
	}

	if f.DropAddressIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropAddress")+" IN (?)")
		values = append(values, f.DropAddressIn)
	}

	if f.DropAddressLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropAddress")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DropAddressLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DropAddressPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropAddress")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DropAddressPrefix))
	}

	if f.DropAddressSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropAddress")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DropAddressSuffix))
	}

	if f.DropAddressNull != nil {
		if *f.DropAddressNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("dropAddress")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("dropAddress")+" IS NOT NULL")
		}
	}

	if f.DropPoint != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropPoint")+" = ?")
		values = append(values, f.DropPoint)
	}

	if f.DropPointNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropPoint")+" != ?")
		values = append(values, f.DropPointNe)
	}

	if f.DropPointGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropPoint")+" > ?")
		values = append(values, f.DropPointGt)
	}

	if f.DropPointLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropPoint")+" < ?")
		values = append(values, f.DropPointLt)
	}

	if f.DropPointGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropPoint")+" >= ?")
		values = append(values, f.DropPointGte)
	}

	if f.DropPointLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropPoint")+" <= ?")
		values = append(values, f.DropPointLte)
	}

	if f.DropPointIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropPoint")+" IN (?)")
		values = append(values, f.DropPointIn)
	}

	if f.DropPointLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropPoint")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DropPointLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DropPointPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropPoint")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DropPointPrefix))
	}

	if f.DropPointSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dropPoint")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DropPointSuffix))
	}

	if f.DropPointNull != nil {
		if *f.DropPointNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("dropPoint")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("dropPoint")+" IS NOT NULL")
		}
	}

	if f.ExpectedDistance != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("expectedDistance")+" = ?")
		values = append(values, f.ExpectedDistance)
	}

	if f.ExpectedDistanceNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("expectedDistance")+" != ?")
		values = append(values, f.ExpectedDistanceNe)
	}

	if f.ExpectedDistanceGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("expectedDistance")+" > ?")
		values = append(values, f.ExpectedDistanceGt)
	}

	if f.ExpectedDistanceLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("expectedDistance")+" < ?")
		values = append(values, f.ExpectedDistanceLt)
	}

	if f.ExpectedDistanceGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("expectedDistance")+" >= ?")
		values = append(values, f.ExpectedDistanceGte)
	}

	if f.ExpectedDistanceLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("expectedDistance")+" <= ?")
		values = append(values, f.ExpectedDistanceLte)
	}

	if f.ExpectedDistanceIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("expectedDistance")+" IN (?)")
		values = append(values, f.ExpectedDistanceIn)
	}

	if f.ExpectedDistanceLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("expectedDistance")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ExpectedDistanceLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ExpectedDistancePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("expectedDistance")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ExpectedDistancePrefix))
	}

	if f.ExpectedDistanceSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("expectedDistance")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ExpectedDistanceSuffix))
	}

	if f.ExpectedDistanceNull != nil {
		if *f.ExpectedDistanceNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("expectedDistance")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("expectedDistance")+" IS NOT NULL")
		}
	}

	if f.ExpectedCost != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("expectedCost")+" = ?")
		values = append(values, f.ExpectedCost)
	}

	if f.ExpectedCostNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("expectedCost")+" != ?")
		values = append(values, f.ExpectedCostNe)
	}

	if f.ExpectedCostGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("expectedCost")+" > ?")
		values = append(values, f.ExpectedCostGt)
	}

	if f.ExpectedCostLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("expectedCost")+" < ?")
		values = append(values, f.ExpectedCostLt)
	}

	if f.ExpectedCostGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("expectedCost")+" >= ?")
		values = append(values, f.ExpectedCostGte)
	}

	if f.ExpectedCostLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("expectedCost")+" <= ?")
		values = append(values, f.ExpectedCostLte)
	}

	if f.ExpectedCostIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("expectedCost")+" IN (?)")
		values = append(values, f.ExpectedCostIn)
	}

	if f.ExpectedCostNull != nil {
		if *f.ExpectedCostNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("expectedCost")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("expectedCost")+" IS NOT NULL")
		}
	}

	if f.Completed != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" = ?")
		values = append(values, f.Completed)
	}

	if f.CompletedNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" != ?")
		values = append(values, f.CompletedNe)
	}

	if f.CompletedGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" > ?")
		values = append(values, f.CompletedGt)
	}

	if f.CompletedLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" < ?")
		values = append(values, f.CompletedLt)
	}

	if f.CompletedGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" >= ?")
		values = append(values, f.CompletedGte)
	}

	if f.CompletedLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" <= ?")
		values = append(values, f.CompletedLte)
	}

	if f.CompletedIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" IN (?)")
		values = append(values, f.CompletedIn)
	}

	if f.CompletedNull != nil {
		if *f.CompletedNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" IS NOT NULL")
		}
	}

	if f.SmsToken != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("smsToken")+" = ?")
		values = append(values, f.SmsToken)
	}

	if f.SmsTokenNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("smsToken")+" != ?")
		values = append(values, f.SmsTokenNe)
	}

	if f.SmsTokenGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("smsToken")+" > ?")
		values = append(values, f.SmsTokenGt)
	}

	if f.SmsTokenLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("smsToken")+" < ?")
		values = append(values, f.SmsTokenLt)
	}

	if f.SmsTokenGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("smsToken")+" >= ?")
		values = append(values, f.SmsTokenGte)
	}

	if f.SmsTokenLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("smsToken")+" <= ?")
		values = append(values, f.SmsTokenLte)
	}

	if f.SmsTokenIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("smsToken")+" IN (?)")
		values = append(values, f.SmsTokenIn)
	}

	if f.SmsTokenLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("smsToken")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.SmsTokenLike, "?", "_", -1), "*", "%", -1))
	}

	if f.SmsTokenPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("smsToken")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.SmsTokenPrefix))
	}

	if f.SmsTokenSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("smsToken")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.SmsTokenSuffix))
	}

	if f.SmsTokenNull != nil {
		if *f.SmsTokenNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("smsToken")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("smsToken")+" IS NOT NULL")
		}
	}

	if f.Status != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("status")+" = ?")
		values = append(values, f.Status)
	}

	if f.StatusNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("status")+" != ?")
		values = append(values, f.StatusNe)
	}

	if f.StatusGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("status")+" > ?")
		values = append(values, f.StatusGt)
	}

	if f.StatusLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("status")+" < ?")
		values = append(values, f.StatusLt)
	}

	if f.StatusGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("status")+" >= ?")
		values = append(values, f.StatusGte)
	}

	if f.StatusLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("status")+" <= ?")
		values = append(values, f.StatusLte)
	}

	if f.StatusIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("status")+" IN (?)")
		values = append(values, f.StatusIn)
	}

	if f.StatusLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("status")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.StatusLike, "?", "_", -1), "*", "%", -1))
	}

	if f.StatusPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("status")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.StatusPrefix))
	}

	if f.StatusSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("status")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.StatusSuffix))
	}

	if f.StatusNull != nil {
		if *f.StatusNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("status")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("status")+" IS NOT NULL")
		}
	}

	if f.SenderID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("senderId")+" = ?")
		values = append(values, f.SenderID)
	}

	if f.SenderIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("senderId")+" != ?")
		values = append(values, f.SenderIDNe)
	}

	if f.SenderIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("senderId")+" > ?")
		values = append(values, f.SenderIDGt)
	}

	if f.SenderIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("senderId")+" < ?")
		values = append(values, f.SenderIDLt)
	}

	if f.SenderIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("senderId")+" >= ?")
		values = append(values, f.SenderIDGte)
	}

	if f.SenderIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("senderId")+" <= ?")
		values = append(values, f.SenderIDLte)
	}

	if f.SenderIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("senderId")+" IN (?)")
		values = append(values, f.SenderIDIn)
	}

	if f.SenderIDNull != nil {
		if *f.SenderIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("senderId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("senderId")+" IS NOT NULL")
		}
	}

	if f.ReceiverID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("receiverId")+" = ?")
		values = append(values, f.ReceiverID)
	}

	if f.ReceiverIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("receiverId")+" != ?")
		values = append(values, f.ReceiverIDNe)
	}

	if f.ReceiverIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("receiverId")+" > ?")
		values = append(values, f.ReceiverIDGt)
	}

	if f.ReceiverIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("receiverId")+" < ?")
		values = append(values, f.ReceiverIDLt)
	}

	if f.ReceiverIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("receiverId")+" >= ?")
		values = append(values, f.ReceiverIDGte)
	}

	if f.ReceiverIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("receiverId")+" <= ?")
		values = append(values, f.ReceiverIDLte)
	}

	if f.ReceiverIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("receiverId")+" IN (?)")
		values = append(values, f.ReceiverIDIn)
	}

	if f.ReceiverIDNull != nil {
		if *f.ReceiverIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("receiverId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("receiverId")+" IS NOT NULL")
		}
	}

	if f.DeliverID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliverId")+" = ?")
		values = append(values, f.DeliverID)
	}

	if f.DeliverIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliverId")+" != ?")
		values = append(values, f.DeliverIDNe)
	}

	if f.DeliverIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliverId")+" > ?")
		values = append(values, f.DeliverIDGt)
	}

	if f.DeliverIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliverId")+" < ?")
		values = append(values, f.DeliverIDLt)
	}

	if f.DeliverIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliverId")+" >= ?")
		values = append(values, f.DeliverIDGte)
	}

	if f.DeliverIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliverId")+" <= ?")
		values = append(values, f.DeliverIDLte)
	}

	if f.DeliverIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliverId")+" IN (?)")
		values = append(values, f.DeliverIDIn)
	}

	if f.DeliverIDNull != nil {
		if *f.DeliverIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("deliverId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("deliverId")+" IS NOT NULL")
		}
	}

	if f.VehicleTypeID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("vehicleTypeId")+" = ?")
		values = append(values, f.VehicleTypeID)
	}

	if f.VehicleTypeIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("vehicleTypeId")+" != ?")
		values = append(values, f.VehicleTypeIDNe)
	}

	if f.VehicleTypeIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("vehicleTypeId")+" > ?")
		values = append(values, f.VehicleTypeIDGt)
	}

	if f.VehicleTypeIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("vehicleTypeId")+" < ?")
		values = append(values, f.VehicleTypeIDLt)
	}

	if f.VehicleTypeIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("vehicleTypeId")+" >= ?")
		values = append(values, f.VehicleTypeIDGte)
	}

	if f.VehicleTypeIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("vehicleTypeId")+" <= ?")
		values = append(values, f.VehicleTypeIDLte)
	}

	if f.VehicleTypeIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("vehicleTypeId")+" IN (?)")
		values = append(values, f.VehicleTypeIDIn)
	}

	if f.VehicleTypeIDNull != nil {
		if *f.VehicleTypeIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("vehicleTypeId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("vehicleTypeId")+" IS NOT NULL")
		}
	}

	if f.DeliveryTypeID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryTypeId")+" = ?")
		values = append(values, f.DeliveryTypeID)
	}

	if f.DeliveryTypeIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryTypeId")+" != ?")
		values = append(values, f.DeliveryTypeIDNe)
	}

	if f.DeliveryTypeIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryTypeId")+" > ?")
		values = append(values, f.DeliveryTypeIDGt)
	}

	if f.DeliveryTypeIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryTypeId")+" < ?")
		values = append(values, f.DeliveryTypeIDLt)
	}

	if f.DeliveryTypeIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryTypeId")+" >= ?")
		values = append(values, f.DeliveryTypeIDGte)
	}

	if f.DeliveryTypeIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryTypeId")+" <= ?")
		values = append(values, f.DeliveryTypeIDLte)
	}

	if f.DeliveryTypeIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryTypeId")+" IN (?)")
		values = append(values, f.DeliveryTypeIDIn)
	}

	if f.DeliveryTypeIDNull != nil {
		if *f.DeliveryTypeIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryTypeId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryTypeId")+" IS NOT NULL")
		}
	}

	if f.DeliveryChannelID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryChannelId")+" = ?")
		values = append(values, f.DeliveryChannelID)
	}

	if f.DeliveryChannelIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryChannelId")+" != ?")
		values = append(values, f.DeliveryChannelIDNe)
	}

	if f.DeliveryChannelIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryChannelId")+" > ?")
		values = append(values, f.DeliveryChannelIDGt)
	}

	if f.DeliveryChannelIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryChannelId")+" < ?")
		values = append(values, f.DeliveryChannelIDLt)
	}

	if f.DeliveryChannelIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryChannelId")+" >= ?")
		values = append(values, f.DeliveryChannelIDGte)
	}

	if f.DeliveryChannelIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryChannelId")+" <= ?")
		values = append(values, f.DeliveryChannelIDLte)
	}

	if f.DeliveryChannelIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryChannelId")+" IN (?)")
		values = append(values, f.DeliveryChannelIDIn)
	}

	if f.DeliveryChannelIDNull != nil {
		if *f.DeliveryChannelIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryChannelId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryChannelId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *DeliveryFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.PaymentIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") = ?")
		values = append(values, f.PaymentIDMin)
	}

	if f.PaymentIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") = ?")
		values = append(values, f.PaymentIDMax)
	}

	if f.PaymentIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") != ?")
		values = append(values, f.PaymentIDMinNe)
	}

	if f.PaymentIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") != ?")
		values = append(values, f.PaymentIDMaxNe)
	}

	if f.PaymentIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") > ?")
		values = append(values, f.PaymentIDMinGt)
	}

	if f.PaymentIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") > ?")
		values = append(values, f.PaymentIDMaxGt)
	}

	if f.PaymentIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") < ?")
		values = append(values, f.PaymentIDMinLt)
	}

	if f.PaymentIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") < ?")
		values = append(values, f.PaymentIDMaxLt)
	}

	if f.PaymentIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") >= ?")
		values = append(values, f.PaymentIDMinGte)
	}

	if f.PaymentIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") >= ?")
		values = append(values, f.PaymentIDMaxGte)
	}

	if f.PaymentIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") <= ?")
		values = append(values, f.PaymentIDMinLte)
	}

	if f.PaymentIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") <= ?")
		values = append(values, f.PaymentIDMaxLte)
	}

	if f.PaymentIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") IN (?)")
		values = append(values, f.PaymentIDMinIn)
	}

	if f.PaymentIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") IN (?)")
		values = append(values, f.PaymentIDMaxIn)
	}

	if f.PaymentTotalMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentTotal")+") = ?")
		values = append(values, f.PaymentTotalMin)
	}

	if f.PaymentTotalMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentTotal")+") = ?")
		values = append(values, f.PaymentTotalMax)
	}

	if f.PaymentTotalAvg != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("paymentTotal")+") = ?")
		values = append(values, f.PaymentTotalAvg)
	}

	if f.PaymentTotalMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentTotal")+") != ?")
		values = append(values, f.PaymentTotalMinNe)
	}

	if f.PaymentTotalMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentTotal")+") != ?")
		values = append(values, f.PaymentTotalMaxNe)
	}

	if f.PaymentTotalAvgNe != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("paymentTotal")+") != ?")
		values = append(values, f.PaymentTotalAvgNe)
	}

	if f.PaymentTotalMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentTotal")+") > ?")
		values = append(values, f.PaymentTotalMinGt)
	}

	if f.PaymentTotalMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentTotal")+") > ?")
		values = append(values, f.PaymentTotalMaxGt)
	}

	if f.PaymentTotalAvgGt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("paymentTotal")+") > ?")
		values = append(values, f.PaymentTotalAvgGt)
	}

	if f.PaymentTotalMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentTotal")+") < ?")
		values = append(values, f.PaymentTotalMinLt)
	}

	if f.PaymentTotalMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentTotal")+") < ?")
		values = append(values, f.PaymentTotalMaxLt)
	}

	if f.PaymentTotalAvgLt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("paymentTotal")+") < ?")
		values = append(values, f.PaymentTotalAvgLt)
	}

	if f.PaymentTotalMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentTotal")+") >= ?")
		values = append(values, f.PaymentTotalMinGte)
	}

	if f.PaymentTotalMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentTotal")+") >= ?")
		values = append(values, f.PaymentTotalMaxGte)
	}

	if f.PaymentTotalAvgGte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("paymentTotal")+") >= ?")
		values = append(values, f.PaymentTotalAvgGte)
	}

	if f.PaymentTotalMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentTotal")+") <= ?")
		values = append(values, f.PaymentTotalMinLte)
	}

	if f.PaymentTotalMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentTotal")+") <= ?")
		values = append(values, f.PaymentTotalMaxLte)
	}

	if f.PaymentTotalAvgLte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("paymentTotal")+") <= ?")
		values = append(values, f.PaymentTotalAvgLte)
	}

	if f.PaymentTotalMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentTotal")+") IN (?)")
		values = append(values, f.PaymentTotalMinIn)
	}

	if f.PaymentTotalMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentTotal")+") IN (?)")
		values = append(values, f.PaymentTotalMaxIn)
	}

	if f.PaymentTotalAvgIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("paymentTotal")+") IN (?)")
		values = append(values, f.PaymentTotalAvgIn)
	}

	if f.PaymentOnDeliverMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentOnDeliver")+") = ?")
		values = append(values, f.PaymentOnDeliverMin)
	}

	if f.PaymentOnDeliverMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentOnDeliver")+") = ?")
		values = append(values, f.PaymentOnDeliverMax)
	}

	if f.PaymentOnDeliverMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentOnDeliver")+") != ?")
		values = append(values, f.PaymentOnDeliverMinNe)
	}

	if f.PaymentOnDeliverMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentOnDeliver")+") != ?")
		values = append(values, f.PaymentOnDeliverMaxNe)
	}

	if f.PaymentOnDeliverMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentOnDeliver")+") > ?")
		values = append(values, f.PaymentOnDeliverMinGt)
	}

	if f.PaymentOnDeliverMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentOnDeliver")+") > ?")
		values = append(values, f.PaymentOnDeliverMaxGt)
	}

	if f.PaymentOnDeliverMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentOnDeliver")+") < ?")
		values = append(values, f.PaymentOnDeliverMinLt)
	}

	if f.PaymentOnDeliverMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentOnDeliver")+") < ?")
		values = append(values, f.PaymentOnDeliverMaxLt)
	}

	if f.PaymentOnDeliverMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentOnDeliver")+") >= ?")
		values = append(values, f.PaymentOnDeliverMinGte)
	}

	if f.PaymentOnDeliverMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentOnDeliver")+") >= ?")
		values = append(values, f.PaymentOnDeliverMaxGte)
	}

	if f.PaymentOnDeliverMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentOnDeliver")+") <= ?")
		values = append(values, f.PaymentOnDeliverMinLte)
	}

	if f.PaymentOnDeliverMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentOnDeliver")+") <= ?")
		values = append(values, f.PaymentOnDeliverMaxLte)
	}

	if f.PaymentOnDeliverMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentOnDeliver")+") IN (?)")
		values = append(values, f.PaymentOnDeliverMinIn)
	}

	if f.PaymentOnDeliverMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentOnDeliver")+") IN (?)")
		values = append(values, f.PaymentOnDeliverMaxIn)
	}

	if f.CollectDateTimeMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectDateTime")+") = ?")
		values = append(values, f.CollectDateTimeMin)
	}

	if f.CollectDateTimeMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectDateTime")+") = ?")
		values = append(values, f.CollectDateTimeMax)
	}

	if f.CollectDateTimeMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectDateTime")+") != ?")
		values = append(values, f.CollectDateTimeMinNe)
	}

	if f.CollectDateTimeMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectDateTime")+") != ?")
		values = append(values, f.CollectDateTimeMaxNe)
	}

	if f.CollectDateTimeMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectDateTime")+") > ?")
		values = append(values, f.CollectDateTimeMinGt)
	}

	if f.CollectDateTimeMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectDateTime")+") > ?")
		values = append(values, f.CollectDateTimeMaxGt)
	}

	if f.CollectDateTimeMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectDateTime")+") < ?")
		values = append(values, f.CollectDateTimeMinLt)
	}

	if f.CollectDateTimeMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectDateTime")+") < ?")
		values = append(values, f.CollectDateTimeMaxLt)
	}

	if f.CollectDateTimeMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectDateTime")+") >= ?")
		values = append(values, f.CollectDateTimeMinGte)
	}

	if f.CollectDateTimeMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectDateTime")+") >= ?")
		values = append(values, f.CollectDateTimeMaxGte)
	}

	if f.CollectDateTimeMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectDateTime")+") <= ?")
		values = append(values, f.CollectDateTimeMinLte)
	}

	if f.CollectDateTimeMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectDateTime")+") <= ?")
		values = append(values, f.CollectDateTimeMaxLte)
	}

	if f.CollectDateTimeMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectDateTime")+") IN (?)")
		values = append(values, f.CollectDateTimeMinIn)
	}

	if f.CollectDateTimeMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectDateTime")+") IN (?)")
		values = append(values, f.CollectDateTimeMaxIn)
	}

	if f.CollectAddressMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectAddress")+") = ?")
		values = append(values, f.CollectAddressMin)
	}

	if f.CollectAddressMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectAddress")+") = ?")
		values = append(values, f.CollectAddressMax)
	}

	if f.CollectAddressMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectAddress")+") != ?")
		values = append(values, f.CollectAddressMinNe)
	}

	if f.CollectAddressMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectAddress")+") != ?")
		values = append(values, f.CollectAddressMaxNe)
	}

	if f.CollectAddressMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectAddress")+") > ?")
		values = append(values, f.CollectAddressMinGt)
	}

	if f.CollectAddressMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectAddress")+") > ?")
		values = append(values, f.CollectAddressMaxGt)
	}

	if f.CollectAddressMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectAddress")+") < ?")
		values = append(values, f.CollectAddressMinLt)
	}

	if f.CollectAddressMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectAddress")+") < ?")
		values = append(values, f.CollectAddressMaxLt)
	}

	if f.CollectAddressMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectAddress")+") >= ?")
		values = append(values, f.CollectAddressMinGte)
	}

	if f.CollectAddressMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectAddress")+") >= ?")
		values = append(values, f.CollectAddressMaxGte)
	}

	if f.CollectAddressMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectAddress")+") <= ?")
		values = append(values, f.CollectAddressMinLte)
	}

	if f.CollectAddressMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectAddress")+") <= ?")
		values = append(values, f.CollectAddressMaxLte)
	}

	if f.CollectAddressMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectAddress")+") IN (?)")
		values = append(values, f.CollectAddressMinIn)
	}

	if f.CollectAddressMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectAddress")+") IN (?)")
		values = append(values, f.CollectAddressMaxIn)
	}

	if f.CollectAddressMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectAddress")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.CollectAddressMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.CollectAddressMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectAddress")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.CollectAddressMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.CollectAddressMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectAddress")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.CollectAddressMinPrefix))
	}

	if f.CollectAddressMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectAddress")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.CollectAddressMaxPrefix))
	}

	if f.CollectAddressMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectAddress")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.CollectAddressMinSuffix))
	}

	if f.CollectAddressMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectAddress")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.CollectAddressMaxSuffix))
	}

	if f.CollectPointMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectPoint")+") = ?")
		values = append(values, f.CollectPointMin)
	}

	if f.CollectPointMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectPoint")+") = ?")
		values = append(values, f.CollectPointMax)
	}

	if f.CollectPointMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectPoint")+") != ?")
		values = append(values, f.CollectPointMinNe)
	}

	if f.CollectPointMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectPoint")+") != ?")
		values = append(values, f.CollectPointMaxNe)
	}

	if f.CollectPointMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectPoint")+") > ?")
		values = append(values, f.CollectPointMinGt)
	}

	if f.CollectPointMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectPoint")+") > ?")
		values = append(values, f.CollectPointMaxGt)
	}

	if f.CollectPointMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectPoint")+") < ?")
		values = append(values, f.CollectPointMinLt)
	}

	if f.CollectPointMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectPoint")+") < ?")
		values = append(values, f.CollectPointMaxLt)
	}

	if f.CollectPointMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectPoint")+") >= ?")
		values = append(values, f.CollectPointMinGte)
	}

	if f.CollectPointMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectPoint")+") >= ?")
		values = append(values, f.CollectPointMaxGte)
	}

	if f.CollectPointMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectPoint")+") <= ?")
		values = append(values, f.CollectPointMinLte)
	}

	if f.CollectPointMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectPoint")+") <= ?")
		values = append(values, f.CollectPointMaxLte)
	}

	if f.CollectPointMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectPoint")+") IN (?)")
		values = append(values, f.CollectPointMinIn)
	}

	if f.CollectPointMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectPoint")+") IN (?)")
		values = append(values, f.CollectPointMaxIn)
	}

	if f.CollectPointMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectPoint")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.CollectPointMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.CollectPointMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectPoint")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.CollectPointMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.CollectPointMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectPoint")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.CollectPointMinPrefix))
	}

	if f.CollectPointMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectPoint")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.CollectPointMaxPrefix))
	}

	if f.CollectPointMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("collectPoint")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.CollectPointMinSuffix))
	}

	if f.CollectPointMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("collectPoint")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.CollectPointMaxSuffix))
	}

	if f.DropDateTimeMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropDateTime")+") = ?")
		values = append(values, f.DropDateTimeMin)
	}

	if f.DropDateTimeMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropDateTime")+") = ?")
		values = append(values, f.DropDateTimeMax)
	}

	if f.DropDateTimeMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropDateTime")+") != ?")
		values = append(values, f.DropDateTimeMinNe)
	}

	if f.DropDateTimeMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropDateTime")+") != ?")
		values = append(values, f.DropDateTimeMaxNe)
	}

	if f.DropDateTimeMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropDateTime")+") > ?")
		values = append(values, f.DropDateTimeMinGt)
	}

	if f.DropDateTimeMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropDateTime")+") > ?")
		values = append(values, f.DropDateTimeMaxGt)
	}

	if f.DropDateTimeMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropDateTime")+") < ?")
		values = append(values, f.DropDateTimeMinLt)
	}

	if f.DropDateTimeMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropDateTime")+") < ?")
		values = append(values, f.DropDateTimeMaxLt)
	}

	if f.DropDateTimeMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropDateTime")+") >= ?")
		values = append(values, f.DropDateTimeMinGte)
	}

	if f.DropDateTimeMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropDateTime")+") >= ?")
		values = append(values, f.DropDateTimeMaxGte)
	}

	if f.DropDateTimeMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropDateTime")+") <= ?")
		values = append(values, f.DropDateTimeMinLte)
	}

	if f.DropDateTimeMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropDateTime")+") <= ?")
		values = append(values, f.DropDateTimeMaxLte)
	}

	if f.DropDateTimeMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropDateTime")+") IN (?)")
		values = append(values, f.DropDateTimeMinIn)
	}

	if f.DropDateTimeMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropDateTime")+") IN (?)")
		values = append(values, f.DropDateTimeMaxIn)
	}

	if f.DropAddressMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropAddress")+") = ?")
		values = append(values, f.DropAddressMin)
	}

	if f.DropAddressMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropAddress")+") = ?")
		values = append(values, f.DropAddressMax)
	}

	if f.DropAddressMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropAddress")+") != ?")
		values = append(values, f.DropAddressMinNe)
	}

	if f.DropAddressMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropAddress")+") != ?")
		values = append(values, f.DropAddressMaxNe)
	}

	if f.DropAddressMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropAddress")+") > ?")
		values = append(values, f.DropAddressMinGt)
	}

	if f.DropAddressMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropAddress")+") > ?")
		values = append(values, f.DropAddressMaxGt)
	}

	if f.DropAddressMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropAddress")+") < ?")
		values = append(values, f.DropAddressMinLt)
	}

	if f.DropAddressMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropAddress")+") < ?")
		values = append(values, f.DropAddressMaxLt)
	}

	if f.DropAddressMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropAddress")+") >= ?")
		values = append(values, f.DropAddressMinGte)
	}

	if f.DropAddressMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropAddress")+") >= ?")
		values = append(values, f.DropAddressMaxGte)
	}

	if f.DropAddressMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropAddress")+") <= ?")
		values = append(values, f.DropAddressMinLte)
	}

	if f.DropAddressMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropAddress")+") <= ?")
		values = append(values, f.DropAddressMaxLte)
	}

	if f.DropAddressMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropAddress")+") IN (?)")
		values = append(values, f.DropAddressMinIn)
	}

	if f.DropAddressMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropAddress")+") IN (?)")
		values = append(values, f.DropAddressMaxIn)
	}

	if f.DropAddressMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropAddress")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DropAddressMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DropAddressMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropAddress")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DropAddressMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DropAddressMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropAddress")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DropAddressMinPrefix))
	}

	if f.DropAddressMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropAddress")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DropAddressMaxPrefix))
	}

	if f.DropAddressMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropAddress")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DropAddressMinSuffix))
	}

	if f.DropAddressMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropAddress")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DropAddressMaxSuffix))
	}

	if f.DropPointMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropPoint")+") = ?")
		values = append(values, f.DropPointMin)
	}

	if f.DropPointMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropPoint")+") = ?")
		values = append(values, f.DropPointMax)
	}

	if f.DropPointMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropPoint")+") != ?")
		values = append(values, f.DropPointMinNe)
	}

	if f.DropPointMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropPoint")+") != ?")
		values = append(values, f.DropPointMaxNe)
	}

	if f.DropPointMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropPoint")+") > ?")
		values = append(values, f.DropPointMinGt)
	}

	if f.DropPointMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropPoint")+") > ?")
		values = append(values, f.DropPointMaxGt)
	}

	if f.DropPointMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropPoint")+") < ?")
		values = append(values, f.DropPointMinLt)
	}

	if f.DropPointMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropPoint")+") < ?")
		values = append(values, f.DropPointMaxLt)
	}

	if f.DropPointMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropPoint")+") >= ?")
		values = append(values, f.DropPointMinGte)
	}

	if f.DropPointMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropPoint")+") >= ?")
		values = append(values, f.DropPointMaxGte)
	}

	if f.DropPointMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropPoint")+") <= ?")
		values = append(values, f.DropPointMinLte)
	}

	if f.DropPointMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropPoint")+") <= ?")
		values = append(values, f.DropPointMaxLte)
	}

	if f.DropPointMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropPoint")+") IN (?)")
		values = append(values, f.DropPointMinIn)
	}

	if f.DropPointMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropPoint")+") IN (?)")
		values = append(values, f.DropPointMaxIn)
	}

	if f.DropPointMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropPoint")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DropPointMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DropPointMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropPoint")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DropPointMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DropPointMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropPoint")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DropPointMinPrefix))
	}

	if f.DropPointMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropPoint")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DropPointMaxPrefix))
	}

	if f.DropPointMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dropPoint")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DropPointMinSuffix))
	}

	if f.DropPointMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dropPoint")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DropPointMaxSuffix))
	}

	if f.ExpectedDistanceMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("expectedDistance")+") = ?")
		values = append(values, f.ExpectedDistanceMin)
	}

	if f.ExpectedDistanceMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("expectedDistance")+") = ?")
		values = append(values, f.ExpectedDistanceMax)
	}

	if f.ExpectedDistanceMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("expectedDistance")+") != ?")
		values = append(values, f.ExpectedDistanceMinNe)
	}

	if f.ExpectedDistanceMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("expectedDistance")+") != ?")
		values = append(values, f.ExpectedDistanceMaxNe)
	}

	if f.ExpectedDistanceMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("expectedDistance")+") > ?")
		values = append(values, f.ExpectedDistanceMinGt)
	}

	if f.ExpectedDistanceMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("expectedDistance")+") > ?")
		values = append(values, f.ExpectedDistanceMaxGt)
	}

	if f.ExpectedDistanceMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("expectedDistance")+") < ?")
		values = append(values, f.ExpectedDistanceMinLt)
	}

	if f.ExpectedDistanceMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("expectedDistance")+") < ?")
		values = append(values, f.ExpectedDistanceMaxLt)
	}

	if f.ExpectedDistanceMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("expectedDistance")+") >= ?")
		values = append(values, f.ExpectedDistanceMinGte)
	}

	if f.ExpectedDistanceMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("expectedDistance")+") >= ?")
		values = append(values, f.ExpectedDistanceMaxGte)
	}

	if f.ExpectedDistanceMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("expectedDistance")+") <= ?")
		values = append(values, f.ExpectedDistanceMinLte)
	}

	if f.ExpectedDistanceMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("expectedDistance")+") <= ?")
		values = append(values, f.ExpectedDistanceMaxLte)
	}

	if f.ExpectedDistanceMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("expectedDistance")+") IN (?)")
		values = append(values, f.ExpectedDistanceMinIn)
	}

	if f.ExpectedDistanceMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("expectedDistance")+") IN (?)")
		values = append(values, f.ExpectedDistanceMaxIn)
	}

	if f.ExpectedDistanceMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("expectedDistance")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ExpectedDistanceMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ExpectedDistanceMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("expectedDistance")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ExpectedDistanceMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ExpectedDistanceMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("expectedDistance")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ExpectedDistanceMinPrefix))
	}

	if f.ExpectedDistanceMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("expectedDistance")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ExpectedDistanceMaxPrefix))
	}

	if f.ExpectedDistanceMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("expectedDistance")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ExpectedDistanceMinSuffix))
	}

	if f.ExpectedDistanceMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("expectedDistance")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ExpectedDistanceMaxSuffix))
	}

	if f.ExpectedCostMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("expectedCost")+") = ?")
		values = append(values, f.ExpectedCostMin)
	}

	if f.ExpectedCostMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("expectedCost")+") = ?")
		values = append(values, f.ExpectedCostMax)
	}

	if f.ExpectedCostAvg != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("expectedCost")+") = ?")
		values = append(values, f.ExpectedCostAvg)
	}

	if f.ExpectedCostMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("expectedCost")+") != ?")
		values = append(values, f.ExpectedCostMinNe)
	}

	if f.ExpectedCostMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("expectedCost")+") != ?")
		values = append(values, f.ExpectedCostMaxNe)
	}

	if f.ExpectedCostAvgNe != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("expectedCost")+") != ?")
		values = append(values, f.ExpectedCostAvgNe)
	}

	if f.ExpectedCostMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("expectedCost")+") > ?")
		values = append(values, f.ExpectedCostMinGt)
	}

	if f.ExpectedCostMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("expectedCost")+") > ?")
		values = append(values, f.ExpectedCostMaxGt)
	}

	if f.ExpectedCostAvgGt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("expectedCost")+") > ?")
		values = append(values, f.ExpectedCostAvgGt)
	}

	if f.ExpectedCostMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("expectedCost")+") < ?")
		values = append(values, f.ExpectedCostMinLt)
	}

	if f.ExpectedCostMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("expectedCost")+") < ?")
		values = append(values, f.ExpectedCostMaxLt)
	}

	if f.ExpectedCostAvgLt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("expectedCost")+") < ?")
		values = append(values, f.ExpectedCostAvgLt)
	}

	if f.ExpectedCostMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("expectedCost")+") >= ?")
		values = append(values, f.ExpectedCostMinGte)
	}

	if f.ExpectedCostMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("expectedCost")+") >= ?")
		values = append(values, f.ExpectedCostMaxGte)
	}

	if f.ExpectedCostAvgGte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("expectedCost")+") >= ?")
		values = append(values, f.ExpectedCostAvgGte)
	}

	if f.ExpectedCostMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("expectedCost")+") <= ?")
		values = append(values, f.ExpectedCostMinLte)
	}

	if f.ExpectedCostMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("expectedCost")+") <= ?")
		values = append(values, f.ExpectedCostMaxLte)
	}

	if f.ExpectedCostAvgLte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("expectedCost")+") <= ?")
		values = append(values, f.ExpectedCostAvgLte)
	}

	if f.ExpectedCostMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("expectedCost")+") IN (?)")
		values = append(values, f.ExpectedCostMinIn)
	}

	if f.ExpectedCostMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("expectedCost")+") IN (?)")
		values = append(values, f.ExpectedCostMaxIn)
	}

	if f.ExpectedCostAvgIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("expectedCost")+") IN (?)")
		values = append(values, f.ExpectedCostAvgIn)
	}

	if f.CompletedMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("completed")+") = ?")
		values = append(values, f.CompletedMin)
	}

	if f.CompletedMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("completed")+") = ?")
		values = append(values, f.CompletedMax)
	}

	if f.CompletedMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("completed")+") != ?")
		values = append(values, f.CompletedMinNe)
	}

	if f.CompletedMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("completed")+") != ?")
		values = append(values, f.CompletedMaxNe)
	}

	if f.CompletedMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("completed")+") > ?")
		values = append(values, f.CompletedMinGt)
	}

	if f.CompletedMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("completed")+") > ?")
		values = append(values, f.CompletedMaxGt)
	}

	if f.CompletedMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("completed")+") < ?")
		values = append(values, f.CompletedMinLt)
	}

	if f.CompletedMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("completed")+") < ?")
		values = append(values, f.CompletedMaxLt)
	}

	if f.CompletedMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("completed")+") >= ?")
		values = append(values, f.CompletedMinGte)
	}

	if f.CompletedMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("completed")+") >= ?")
		values = append(values, f.CompletedMaxGte)
	}

	if f.CompletedMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("completed")+") <= ?")
		values = append(values, f.CompletedMinLte)
	}

	if f.CompletedMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("completed")+") <= ?")
		values = append(values, f.CompletedMaxLte)
	}

	if f.CompletedMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("completed")+") IN (?)")
		values = append(values, f.CompletedMinIn)
	}

	if f.CompletedMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("completed")+") IN (?)")
		values = append(values, f.CompletedMaxIn)
	}

	if f.SmsTokenMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("smsToken")+") = ?")
		values = append(values, f.SmsTokenMin)
	}

	if f.SmsTokenMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("smsToken")+") = ?")
		values = append(values, f.SmsTokenMax)
	}

	if f.SmsTokenMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("smsToken")+") != ?")
		values = append(values, f.SmsTokenMinNe)
	}

	if f.SmsTokenMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("smsToken")+") != ?")
		values = append(values, f.SmsTokenMaxNe)
	}

	if f.SmsTokenMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("smsToken")+") > ?")
		values = append(values, f.SmsTokenMinGt)
	}

	if f.SmsTokenMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("smsToken")+") > ?")
		values = append(values, f.SmsTokenMaxGt)
	}

	if f.SmsTokenMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("smsToken")+") < ?")
		values = append(values, f.SmsTokenMinLt)
	}

	if f.SmsTokenMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("smsToken")+") < ?")
		values = append(values, f.SmsTokenMaxLt)
	}

	if f.SmsTokenMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("smsToken")+") >= ?")
		values = append(values, f.SmsTokenMinGte)
	}

	if f.SmsTokenMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("smsToken")+") >= ?")
		values = append(values, f.SmsTokenMaxGte)
	}

	if f.SmsTokenMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("smsToken")+") <= ?")
		values = append(values, f.SmsTokenMinLte)
	}

	if f.SmsTokenMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("smsToken")+") <= ?")
		values = append(values, f.SmsTokenMaxLte)
	}

	if f.SmsTokenMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("smsToken")+") IN (?)")
		values = append(values, f.SmsTokenMinIn)
	}

	if f.SmsTokenMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("smsToken")+") IN (?)")
		values = append(values, f.SmsTokenMaxIn)
	}

	if f.SmsTokenMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("smsToken")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.SmsTokenMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.SmsTokenMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("smsToken")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.SmsTokenMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.SmsTokenMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("smsToken")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.SmsTokenMinPrefix))
	}

	if f.SmsTokenMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("smsToken")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.SmsTokenMaxPrefix))
	}

	if f.SmsTokenMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("smsToken")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.SmsTokenMinSuffix))
	}

	if f.SmsTokenMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("smsToken")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.SmsTokenMaxSuffix))
	}

	if f.StatusMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("status")+") = ?")
		values = append(values, f.StatusMin)
	}

	if f.StatusMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("status")+") = ?")
		values = append(values, f.StatusMax)
	}

	if f.StatusMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("status")+") != ?")
		values = append(values, f.StatusMinNe)
	}

	if f.StatusMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("status")+") != ?")
		values = append(values, f.StatusMaxNe)
	}

	if f.StatusMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("status")+") > ?")
		values = append(values, f.StatusMinGt)
	}

	if f.StatusMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("status")+") > ?")
		values = append(values, f.StatusMaxGt)
	}

	if f.StatusMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("status")+") < ?")
		values = append(values, f.StatusMinLt)
	}

	if f.StatusMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("status")+") < ?")
		values = append(values, f.StatusMaxLt)
	}

	if f.StatusMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("status")+") >= ?")
		values = append(values, f.StatusMinGte)
	}

	if f.StatusMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("status")+") >= ?")
		values = append(values, f.StatusMaxGte)
	}

	if f.StatusMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("status")+") <= ?")
		values = append(values, f.StatusMinLte)
	}

	if f.StatusMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("status")+") <= ?")
		values = append(values, f.StatusMaxLte)
	}

	if f.StatusMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("status")+") IN (?)")
		values = append(values, f.StatusMinIn)
	}

	if f.StatusMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("status")+") IN (?)")
		values = append(values, f.StatusMaxIn)
	}

	if f.StatusMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("status")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.StatusMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.StatusMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("status")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.StatusMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.StatusMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("status")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.StatusMinPrefix))
	}

	if f.StatusMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("status")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.StatusMaxPrefix))
	}

	if f.StatusMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("status")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.StatusMinSuffix))
	}

	if f.StatusMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("status")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.StatusMaxSuffix))
	}

	if f.SenderIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("senderId")+") = ?")
		values = append(values, f.SenderIDMin)
	}

	if f.SenderIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("senderId")+") = ?")
		values = append(values, f.SenderIDMax)
	}

	if f.SenderIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("senderId")+") != ?")
		values = append(values, f.SenderIDMinNe)
	}

	if f.SenderIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("senderId")+") != ?")
		values = append(values, f.SenderIDMaxNe)
	}

	if f.SenderIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("senderId")+") > ?")
		values = append(values, f.SenderIDMinGt)
	}

	if f.SenderIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("senderId")+") > ?")
		values = append(values, f.SenderIDMaxGt)
	}

	if f.SenderIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("senderId")+") < ?")
		values = append(values, f.SenderIDMinLt)
	}

	if f.SenderIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("senderId")+") < ?")
		values = append(values, f.SenderIDMaxLt)
	}

	if f.SenderIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("senderId")+") >= ?")
		values = append(values, f.SenderIDMinGte)
	}

	if f.SenderIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("senderId")+") >= ?")
		values = append(values, f.SenderIDMaxGte)
	}

	if f.SenderIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("senderId")+") <= ?")
		values = append(values, f.SenderIDMinLte)
	}

	if f.SenderIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("senderId")+") <= ?")
		values = append(values, f.SenderIDMaxLte)
	}

	if f.SenderIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("senderId")+") IN (?)")
		values = append(values, f.SenderIDMinIn)
	}

	if f.SenderIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("senderId")+") IN (?)")
		values = append(values, f.SenderIDMaxIn)
	}

	if f.ReceiverIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("receiverId")+") = ?")
		values = append(values, f.ReceiverIDMin)
	}

	if f.ReceiverIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("receiverId")+") = ?")
		values = append(values, f.ReceiverIDMax)
	}

	if f.ReceiverIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("receiverId")+") != ?")
		values = append(values, f.ReceiverIDMinNe)
	}

	if f.ReceiverIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("receiverId")+") != ?")
		values = append(values, f.ReceiverIDMaxNe)
	}

	if f.ReceiverIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("receiverId")+") > ?")
		values = append(values, f.ReceiverIDMinGt)
	}

	if f.ReceiverIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("receiverId")+") > ?")
		values = append(values, f.ReceiverIDMaxGt)
	}

	if f.ReceiverIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("receiverId")+") < ?")
		values = append(values, f.ReceiverIDMinLt)
	}

	if f.ReceiverIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("receiverId")+") < ?")
		values = append(values, f.ReceiverIDMaxLt)
	}

	if f.ReceiverIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("receiverId")+") >= ?")
		values = append(values, f.ReceiverIDMinGte)
	}

	if f.ReceiverIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("receiverId")+") >= ?")
		values = append(values, f.ReceiverIDMaxGte)
	}

	if f.ReceiverIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("receiverId")+") <= ?")
		values = append(values, f.ReceiverIDMinLte)
	}

	if f.ReceiverIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("receiverId")+") <= ?")
		values = append(values, f.ReceiverIDMaxLte)
	}

	if f.ReceiverIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("receiverId")+") IN (?)")
		values = append(values, f.ReceiverIDMinIn)
	}

	if f.ReceiverIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("receiverId")+") IN (?)")
		values = append(values, f.ReceiverIDMaxIn)
	}

	if f.DeliverIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliverId")+") = ?")
		values = append(values, f.DeliverIDMin)
	}

	if f.DeliverIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliverId")+") = ?")
		values = append(values, f.DeliverIDMax)
	}

	if f.DeliverIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliverId")+") != ?")
		values = append(values, f.DeliverIDMinNe)
	}

	if f.DeliverIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliverId")+") != ?")
		values = append(values, f.DeliverIDMaxNe)
	}

	if f.DeliverIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliverId")+") > ?")
		values = append(values, f.DeliverIDMinGt)
	}

	if f.DeliverIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliverId")+") > ?")
		values = append(values, f.DeliverIDMaxGt)
	}

	if f.DeliverIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliverId")+") < ?")
		values = append(values, f.DeliverIDMinLt)
	}

	if f.DeliverIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliverId")+") < ?")
		values = append(values, f.DeliverIDMaxLt)
	}

	if f.DeliverIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliverId")+") >= ?")
		values = append(values, f.DeliverIDMinGte)
	}

	if f.DeliverIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliverId")+") >= ?")
		values = append(values, f.DeliverIDMaxGte)
	}

	if f.DeliverIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliverId")+") <= ?")
		values = append(values, f.DeliverIDMinLte)
	}

	if f.DeliverIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliverId")+") <= ?")
		values = append(values, f.DeliverIDMaxLte)
	}

	if f.DeliverIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliverId")+") IN (?)")
		values = append(values, f.DeliverIDMinIn)
	}

	if f.DeliverIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliverId")+") IN (?)")
		values = append(values, f.DeliverIDMaxIn)
	}

	if f.VehicleTypeIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("vehicleTypeId")+") = ?")
		values = append(values, f.VehicleTypeIDMin)
	}

	if f.VehicleTypeIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("vehicleTypeId")+") = ?")
		values = append(values, f.VehicleTypeIDMax)
	}

	if f.VehicleTypeIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("vehicleTypeId")+") != ?")
		values = append(values, f.VehicleTypeIDMinNe)
	}

	if f.VehicleTypeIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("vehicleTypeId")+") != ?")
		values = append(values, f.VehicleTypeIDMaxNe)
	}

	if f.VehicleTypeIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("vehicleTypeId")+") > ?")
		values = append(values, f.VehicleTypeIDMinGt)
	}

	if f.VehicleTypeIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("vehicleTypeId")+") > ?")
		values = append(values, f.VehicleTypeIDMaxGt)
	}

	if f.VehicleTypeIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("vehicleTypeId")+") < ?")
		values = append(values, f.VehicleTypeIDMinLt)
	}

	if f.VehicleTypeIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("vehicleTypeId")+") < ?")
		values = append(values, f.VehicleTypeIDMaxLt)
	}

	if f.VehicleTypeIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("vehicleTypeId")+") >= ?")
		values = append(values, f.VehicleTypeIDMinGte)
	}

	if f.VehicleTypeIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("vehicleTypeId")+") >= ?")
		values = append(values, f.VehicleTypeIDMaxGte)
	}

	if f.VehicleTypeIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("vehicleTypeId")+") <= ?")
		values = append(values, f.VehicleTypeIDMinLte)
	}

	if f.VehicleTypeIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("vehicleTypeId")+") <= ?")
		values = append(values, f.VehicleTypeIDMaxLte)
	}

	if f.VehicleTypeIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("vehicleTypeId")+") IN (?)")
		values = append(values, f.VehicleTypeIDMinIn)
	}

	if f.VehicleTypeIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("vehicleTypeId")+") IN (?)")
		values = append(values, f.VehicleTypeIDMaxIn)
	}

	if f.DeliveryTypeIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryTypeId")+") = ?")
		values = append(values, f.DeliveryTypeIDMin)
	}

	if f.DeliveryTypeIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryTypeId")+") = ?")
		values = append(values, f.DeliveryTypeIDMax)
	}

	if f.DeliveryTypeIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryTypeId")+") != ?")
		values = append(values, f.DeliveryTypeIDMinNe)
	}

	if f.DeliveryTypeIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryTypeId")+") != ?")
		values = append(values, f.DeliveryTypeIDMaxNe)
	}

	if f.DeliveryTypeIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryTypeId")+") > ?")
		values = append(values, f.DeliveryTypeIDMinGt)
	}

	if f.DeliveryTypeIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryTypeId")+") > ?")
		values = append(values, f.DeliveryTypeIDMaxGt)
	}

	if f.DeliveryTypeIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryTypeId")+") < ?")
		values = append(values, f.DeliveryTypeIDMinLt)
	}

	if f.DeliveryTypeIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryTypeId")+") < ?")
		values = append(values, f.DeliveryTypeIDMaxLt)
	}

	if f.DeliveryTypeIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryTypeId")+") >= ?")
		values = append(values, f.DeliveryTypeIDMinGte)
	}

	if f.DeliveryTypeIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryTypeId")+") >= ?")
		values = append(values, f.DeliveryTypeIDMaxGte)
	}

	if f.DeliveryTypeIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryTypeId")+") <= ?")
		values = append(values, f.DeliveryTypeIDMinLte)
	}

	if f.DeliveryTypeIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryTypeId")+") <= ?")
		values = append(values, f.DeliveryTypeIDMaxLte)
	}

	if f.DeliveryTypeIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryTypeId")+") IN (?)")
		values = append(values, f.DeliveryTypeIDMinIn)
	}

	if f.DeliveryTypeIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryTypeId")+") IN (?)")
		values = append(values, f.DeliveryTypeIDMaxIn)
	}

	if f.DeliveryChannelIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryChannelId")+") = ?")
		values = append(values, f.DeliveryChannelIDMin)
	}

	if f.DeliveryChannelIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryChannelId")+") = ?")
		values = append(values, f.DeliveryChannelIDMax)
	}

	if f.DeliveryChannelIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryChannelId")+") != ?")
		values = append(values, f.DeliveryChannelIDMinNe)
	}

	if f.DeliveryChannelIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryChannelId")+") != ?")
		values = append(values, f.DeliveryChannelIDMaxNe)
	}

	if f.DeliveryChannelIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryChannelId")+") > ?")
		values = append(values, f.DeliveryChannelIDMinGt)
	}

	if f.DeliveryChannelIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryChannelId")+") > ?")
		values = append(values, f.DeliveryChannelIDMaxGt)
	}

	if f.DeliveryChannelIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryChannelId")+") < ?")
		values = append(values, f.DeliveryChannelIDMinLt)
	}

	if f.DeliveryChannelIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryChannelId")+") < ?")
		values = append(values, f.DeliveryChannelIDMaxLt)
	}

	if f.DeliveryChannelIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryChannelId")+") >= ?")
		values = append(values, f.DeliveryChannelIDMinGte)
	}

	if f.DeliveryChannelIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryChannelId")+") >= ?")
		values = append(values, f.DeliveryChannelIDMaxGte)
	}

	if f.DeliveryChannelIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryChannelId")+") <= ?")
		values = append(values, f.DeliveryChannelIDMinLte)
	}

	if f.DeliveryChannelIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryChannelId")+") <= ?")
		values = append(values, f.DeliveryChannelIDMaxLte)
	}

	if f.DeliveryChannelIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryChannelId")+") IN (?)")
		values = append(values, f.DeliveryChannelIDMinIn)
	}

	if f.DeliveryChannelIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryChannelId")+") IN (?)")
		values = append(values, f.DeliveryChannelIDMaxIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *DeliveryFilterType) AndWith(f2 ...*DeliveryFilterType) *DeliveryFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &DeliveryFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *DeliveryFilterType) OrWith(f2 ...*DeliveryFilterType) *DeliveryFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &DeliveryFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *PersonFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *PersonFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("people"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *PersonFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Deliveries != nil {
		_alias := alias + "_deliveries"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("deliveries"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("deliverId")+" = "+dialect.Quote(alias)+".id")
		err := f.Deliveries.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.DeliveriesSent != nil {
		_alias := alias + "_deliveriesSent"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("deliveries"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("senderId")+" = "+dialect.Quote(alias)+".id")
		err := f.DeliveriesSent.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.DeliveriesReceived != nil {
		_alias := alias + "_deliveriesReceived"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("deliveries"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("receiverId")+" = "+dialect.Quote(alias)+".id")
		err := f.DeliveriesReceived.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *PersonFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Deliver != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliver")+" = ?")
		values = append(values, f.Deliver)
	}

	if f.DeliverNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliver")+" != ?")
		values = append(values, f.DeliverNe)
	}

	if f.DeliverGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliver")+" > ?")
		values = append(values, f.DeliverGt)
	}

	if f.DeliverLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliver")+" < ?")
		values = append(values, f.DeliverLt)
	}

	if f.DeliverGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliver")+" >= ?")
		values = append(values, f.DeliverGte)
	}

	if f.DeliverLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliver")+" <= ?")
		values = append(values, f.DeliverLte)
	}

	if f.DeliverIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliver")+" IN (?)")
		values = append(values, f.DeliverIn)
	}

	if f.DeliverNull != nil {
		if *f.DeliverNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("deliver")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("deliver")+" IS NOT NULL")
		}
	}

	if f.Email != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" = ?")
		values = append(values, f.Email)
	}

	if f.EmailNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" != ?")
		values = append(values, f.EmailNe)
	}

	if f.EmailGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" > ?")
		values = append(values, f.EmailGt)
	}

	if f.EmailLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" < ?")
		values = append(values, f.EmailLt)
	}

	if f.EmailGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" >= ?")
		values = append(values, f.EmailGte)
	}

	if f.EmailLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" <= ?")
		values = append(values, f.EmailLte)
	}

	if f.EmailIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" IN (?)")
		values = append(values, f.EmailIn)
	}

	if f.EmailLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EmailLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EmailPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailPrefix))
	}

	if f.EmailSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailSuffix))
	}

	if f.EmailNull != nil {
		if *f.EmailNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" IS NOT NULL")
		}
	}

	if f.Phone != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" = ?")
		values = append(values, f.Phone)
	}

	if f.PhoneNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" != ?")
		values = append(values, f.PhoneNe)
	}

	if f.PhoneGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" > ?")
		values = append(values, f.PhoneGt)
	}

	if f.PhoneLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" < ?")
		values = append(values, f.PhoneLt)
	}

	if f.PhoneGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" >= ?")
		values = append(values, f.PhoneGte)
	}

	if f.PhoneLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" <= ?")
		values = append(values, f.PhoneLte)
	}

	if f.PhoneIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" IN (?)")
		values = append(values, f.PhoneIn)
	}

	if f.PhoneLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PhoneLike, "?", "_", -1), "*", "%", -1))
	}

	if f.PhonePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PhonePrefix))
	}

	if f.PhoneSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PhoneSuffix))
	}

	if f.PhoneNull != nil {
		if *f.PhoneNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" IS NOT NULL")
		}
	}

	if f.DocumentNo != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("documentNo")+" = ?")
		values = append(values, f.DocumentNo)
	}

	if f.DocumentNoNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("documentNo")+" != ?")
		values = append(values, f.DocumentNoNe)
	}

	if f.DocumentNoGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("documentNo")+" > ?")
		values = append(values, f.DocumentNoGt)
	}

	if f.DocumentNoLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("documentNo")+" < ?")
		values = append(values, f.DocumentNoLt)
	}

	if f.DocumentNoGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("documentNo")+" >= ?")
		values = append(values, f.DocumentNoGte)
	}

	if f.DocumentNoLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("documentNo")+" <= ?")
		values = append(values, f.DocumentNoLte)
	}

	if f.DocumentNoIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("documentNo")+" IN (?)")
		values = append(values, f.DocumentNoIn)
	}

	if f.DocumentNoLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("documentNo")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DocumentNoLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DocumentNoPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("documentNo")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DocumentNoPrefix))
	}

	if f.DocumentNoSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("documentNo")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DocumentNoSuffix))
	}

	if f.DocumentNoNull != nil {
		if *f.DocumentNoNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("documentNo")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("documentNo")+" IS NOT NULL")
		}
	}

	if f.AvatarURL != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" = ?")
		values = append(values, f.AvatarURL)
	}

	if f.AvatarURLNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" != ?")
		values = append(values, f.AvatarURLNe)
	}

	if f.AvatarURLGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" > ?")
		values = append(values, f.AvatarURLGt)
	}

	if f.AvatarURLLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" < ?")
		values = append(values, f.AvatarURLLt)
	}

	if f.AvatarURLGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" >= ?")
		values = append(values, f.AvatarURLGte)
	}

	if f.AvatarURLLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" <= ?")
		values = append(values, f.AvatarURLLte)
	}

	if f.AvatarURLIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" IN (?)")
		values = append(values, f.AvatarURLIn)
	}

	if f.AvatarURLLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.AvatarURLLike, "?", "_", -1), "*", "%", -1))
	}

	if f.AvatarURLPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AvatarURLPrefix))
	}

	if f.AvatarURLSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AvatarURLSuffix))
	}

	if f.AvatarURLNull != nil {
		if *f.AvatarURLNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" IS NOT NULL")
		}
	}

	if f.DisplayName != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" = ?")
		values = append(values, f.DisplayName)
	}

	if f.DisplayNameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" != ?")
		values = append(values, f.DisplayNameNe)
	}

	if f.DisplayNameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" > ?")
		values = append(values, f.DisplayNameGt)
	}

	if f.DisplayNameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" < ?")
		values = append(values, f.DisplayNameLt)
	}

	if f.DisplayNameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" >= ?")
		values = append(values, f.DisplayNameGte)
	}

	if f.DisplayNameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" <= ?")
		values = append(values, f.DisplayNameLte)
	}

	if f.DisplayNameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" IN (?)")
		values = append(values, f.DisplayNameIn)
	}

	if f.DisplayNameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DisplayNameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DisplayNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DisplayNamePrefix))
	}

	if f.DisplayNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DisplayNameSuffix))
	}

	if f.DisplayNameNull != nil {
		if *f.DisplayNameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" IS NOT NULL")
		}
	}

	if f.FirstName != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" = ?")
		values = append(values, f.FirstName)
	}

	if f.FirstNameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" != ?")
		values = append(values, f.FirstNameNe)
	}

	if f.FirstNameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" > ?")
		values = append(values, f.FirstNameGt)
	}

	if f.FirstNameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" < ?")
		values = append(values, f.FirstNameLt)
	}

	if f.FirstNameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" >= ?")
		values = append(values, f.FirstNameGte)
	}

	if f.FirstNameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" <= ?")
		values = append(values, f.FirstNameLte)
	}

	if f.FirstNameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" IN (?)")
		values = append(values, f.FirstNameIn)
	}

	if f.FirstNameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.FirstNameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.FirstNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.FirstNamePrefix))
	}

	if f.FirstNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.FirstNameSuffix))
	}

	if f.FirstNameNull != nil {
		if *f.FirstNameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" IS NOT NULL")
		}
	}

	if f.LastName != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" = ?")
		values = append(values, f.LastName)
	}

	if f.LastNameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" != ?")
		values = append(values, f.LastNameNe)
	}

	if f.LastNameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" > ?")
		values = append(values, f.LastNameGt)
	}

	if f.LastNameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" < ?")
		values = append(values, f.LastNameLt)
	}

	if f.LastNameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" >= ?")
		values = append(values, f.LastNameGte)
	}

	if f.LastNameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" <= ?")
		values = append(values, f.LastNameLte)
	}

	if f.LastNameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" IN (?)")
		values = append(values, f.LastNameIn)
	}

	if f.LastNameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LastNameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LastNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LastNamePrefix))
	}

	if f.LastNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LastNameSuffix))
	}

	if f.LastNameNull != nil {
		if *f.LastNameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" IS NOT NULL")
		}
	}

	if f.NickName != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" = ?")
		values = append(values, f.NickName)
	}

	if f.NickNameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" != ?")
		values = append(values, f.NickNameNe)
	}

	if f.NickNameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" > ?")
		values = append(values, f.NickNameGt)
	}

	if f.NickNameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" < ?")
		values = append(values, f.NickNameLt)
	}

	if f.NickNameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" >= ?")
		values = append(values, f.NickNameGte)
	}

	if f.NickNameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" <= ?")
		values = append(values, f.NickNameLte)
	}

	if f.NickNameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" IN (?)")
		values = append(values, f.NickNameIn)
	}

	if f.NickNameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NickNameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NickNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NickNamePrefix))
	}

	if f.NickNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NickNameSuffix))
	}

	if f.NickNameNull != nil {
		if *f.NickNameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" IS NOT NULL")
		}
	}

	if f.Description != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" = ?")
		values = append(values, f.Description)
	}

	if f.DescriptionNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" != ?")
		values = append(values, f.DescriptionNe)
	}

	if f.DescriptionGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" > ?")
		values = append(values, f.DescriptionGt)
	}

	if f.DescriptionLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" < ?")
		values = append(values, f.DescriptionLt)
	}

	if f.DescriptionGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" >= ?")
		values = append(values, f.DescriptionGte)
	}

	if f.DescriptionLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" <= ?")
		values = append(values, f.DescriptionLte)
	}

	if f.DescriptionIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IN (?)")
		values = append(values, f.DescriptionIn)
	}

	if f.DescriptionLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionPrefix))
	}

	if f.DescriptionSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionSuffix))
	}

	if f.DescriptionNull != nil {
		if *f.DescriptionNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NOT NULL")
		}
	}

	if f.Location != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" = ?")
		values = append(values, f.Location)
	}

	if f.LocationNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" != ?")
		values = append(values, f.LocationNe)
	}

	if f.LocationGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" > ?")
		values = append(values, f.LocationGt)
	}

	if f.LocationLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" < ?")
		values = append(values, f.LocationLt)
	}

	if f.LocationGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" >= ?")
		values = append(values, f.LocationGte)
	}

	if f.LocationLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" <= ?")
		values = append(values, f.LocationLte)
	}

	if f.LocationIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" IN (?)")
		values = append(values, f.LocationIn)
	}

	if f.LocationLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LocationLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LocationPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LocationPrefix))
	}

	if f.LocationSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LocationSuffix))
	}

	if f.LocationNull != nil {
		if *f.LocationNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" IS NOT NULL")
		}
	}

	if f.UserID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" = ?")
		values = append(values, f.UserID)
	}

	if f.UserIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" != ?")
		values = append(values, f.UserIDNe)
	}

	if f.UserIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" > ?")
		values = append(values, f.UserIDGt)
	}

	if f.UserIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" < ?")
		values = append(values, f.UserIDLt)
	}

	if f.UserIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" >= ?")
		values = append(values, f.UserIDGte)
	}

	if f.UserIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" <= ?")
		values = append(values, f.UserIDLte)
	}

	if f.UserIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" IN (?)")
		values = append(values, f.UserIDIn)
	}

	if f.UserIDLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.UserIDLike, "?", "_", -1), "*", "%", -1))
	}

	if f.UserIDPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.UserIDPrefix))
	}

	if f.UserIDSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.UserIDSuffix))
	}

	if f.UserIDNull != nil {
		if *f.UserIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *PersonFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.DeliverMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliver")+") = ?")
		values = append(values, f.DeliverMin)
	}

	if f.DeliverMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliver")+") = ?")
		values = append(values, f.DeliverMax)
	}

	if f.DeliverMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliver")+") != ?")
		values = append(values, f.DeliverMinNe)
	}

	if f.DeliverMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliver")+") != ?")
		values = append(values, f.DeliverMaxNe)
	}

	if f.DeliverMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliver")+") > ?")
		values = append(values, f.DeliverMinGt)
	}

	if f.DeliverMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliver")+") > ?")
		values = append(values, f.DeliverMaxGt)
	}

	if f.DeliverMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliver")+") < ?")
		values = append(values, f.DeliverMinLt)
	}

	if f.DeliverMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliver")+") < ?")
		values = append(values, f.DeliverMaxLt)
	}

	if f.DeliverMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliver")+") >= ?")
		values = append(values, f.DeliverMinGte)
	}

	if f.DeliverMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliver")+") >= ?")
		values = append(values, f.DeliverMaxGte)
	}

	if f.DeliverMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliver")+") <= ?")
		values = append(values, f.DeliverMinLte)
	}

	if f.DeliverMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliver")+") <= ?")
		values = append(values, f.DeliverMaxLte)
	}

	if f.DeliverMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliver")+") IN (?)")
		values = append(values, f.DeliverMinIn)
	}

	if f.DeliverMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliver")+") IN (?)")
		values = append(values, f.DeliverMaxIn)
	}

	if f.EmailMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") = ?")
		values = append(values, f.EmailMin)
	}

	if f.EmailMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") = ?")
		values = append(values, f.EmailMax)
	}

	if f.EmailMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") != ?")
		values = append(values, f.EmailMinNe)
	}

	if f.EmailMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") != ?")
		values = append(values, f.EmailMaxNe)
	}

	if f.EmailMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") > ?")
		values = append(values, f.EmailMinGt)
	}

	if f.EmailMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") > ?")
		values = append(values, f.EmailMaxGt)
	}

	if f.EmailMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") < ?")
		values = append(values, f.EmailMinLt)
	}

	if f.EmailMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") < ?")
		values = append(values, f.EmailMaxLt)
	}

	if f.EmailMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") >= ?")
		values = append(values, f.EmailMinGte)
	}

	if f.EmailMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") >= ?")
		values = append(values, f.EmailMaxGte)
	}

	if f.EmailMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") <= ?")
		values = append(values, f.EmailMinLte)
	}

	if f.EmailMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") <= ?")
		values = append(values, f.EmailMaxLte)
	}

	if f.EmailMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") IN (?)")
		values = append(values, f.EmailMinIn)
	}

	if f.EmailMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") IN (?)")
		values = append(values, f.EmailMaxIn)
	}

	if f.EmailMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EmailMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EmailMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EmailMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EmailMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailMinPrefix))
	}

	if f.EmailMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailMaxPrefix))
	}

	if f.EmailMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailMinSuffix))
	}

	if f.EmailMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailMaxSuffix))
	}

	if f.PhoneMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") = ?")
		values = append(values, f.PhoneMin)
	}

	if f.PhoneMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") = ?")
		values = append(values, f.PhoneMax)
	}

	if f.PhoneMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") != ?")
		values = append(values, f.PhoneMinNe)
	}

	if f.PhoneMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") != ?")
		values = append(values, f.PhoneMaxNe)
	}

	if f.PhoneMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") > ?")
		values = append(values, f.PhoneMinGt)
	}

	if f.PhoneMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") > ?")
		values = append(values, f.PhoneMaxGt)
	}

	if f.PhoneMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") < ?")
		values = append(values, f.PhoneMinLt)
	}

	if f.PhoneMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") < ?")
		values = append(values, f.PhoneMaxLt)
	}

	if f.PhoneMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") >= ?")
		values = append(values, f.PhoneMinGte)
	}

	if f.PhoneMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") >= ?")
		values = append(values, f.PhoneMaxGte)
	}

	if f.PhoneMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") <= ?")
		values = append(values, f.PhoneMinLte)
	}

	if f.PhoneMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") <= ?")
		values = append(values, f.PhoneMaxLte)
	}

	if f.PhoneMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") IN (?)")
		values = append(values, f.PhoneMinIn)
	}

	if f.PhoneMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") IN (?)")
		values = append(values, f.PhoneMaxIn)
	}

	if f.PhoneMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PhoneMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.PhoneMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PhoneMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.PhoneMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PhoneMinPrefix))
	}

	if f.PhoneMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PhoneMaxPrefix))
	}

	if f.PhoneMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PhoneMinSuffix))
	}

	if f.PhoneMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PhoneMaxSuffix))
	}

	if f.DocumentNoMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("documentNo")+") = ?")
		values = append(values, f.DocumentNoMin)
	}

	if f.DocumentNoMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("documentNo")+") = ?")
		values = append(values, f.DocumentNoMax)
	}

	if f.DocumentNoMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("documentNo")+") != ?")
		values = append(values, f.DocumentNoMinNe)
	}

	if f.DocumentNoMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("documentNo")+") != ?")
		values = append(values, f.DocumentNoMaxNe)
	}

	if f.DocumentNoMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("documentNo")+") > ?")
		values = append(values, f.DocumentNoMinGt)
	}

	if f.DocumentNoMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("documentNo")+") > ?")
		values = append(values, f.DocumentNoMaxGt)
	}

	if f.DocumentNoMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("documentNo")+") < ?")
		values = append(values, f.DocumentNoMinLt)
	}

	if f.DocumentNoMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("documentNo")+") < ?")
		values = append(values, f.DocumentNoMaxLt)
	}

	if f.DocumentNoMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("documentNo")+") >= ?")
		values = append(values, f.DocumentNoMinGte)
	}

	if f.DocumentNoMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("documentNo")+") >= ?")
		values = append(values, f.DocumentNoMaxGte)
	}

	if f.DocumentNoMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("documentNo")+") <= ?")
		values = append(values, f.DocumentNoMinLte)
	}

	if f.DocumentNoMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("documentNo")+") <= ?")
		values = append(values, f.DocumentNoMaxLte)
	}

	if f.DocumentNoMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("documentNo")+") IN (?)")
		values = append(values, f.DocumentNoMinIn)
	}

	if f.DocumentNoMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("documentNo")+") IN (?)")
		values = append(values, f.DocumentNoMaxIn)
	}

	if f.DocumentNoMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("documentNo")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DocumentNoMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DocumentNoMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("documentNo")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DocumentNoMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DocumentNoMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("documentNo")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DocumentNoMinPrefix))
	}

	if f.DocumentNoMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("documentNo")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DocumentNoMaxPrefix))
	}

	if f.DocumentNoMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("documentNo")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DocumentNoMinSuffix))
	}

	if f.DocumentNoMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("documentNo")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DocumentNoMaxSuffix))
	}

	if f.AvatarURLMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") = ?")
		values = append(values, f.AvatarURLMin)
	}

	if f.AvatarURLMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") = ?")
		values = append(values, f.AvatarURLMax)
	}

	if f.AvatarURLMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") != ?")
		values = append(values, f.AvatarURLMinNe)
	}

	if f.AvatarURLMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") != ?")
		values = append(values, f.AvatarURLMaxNe)
	}

	if f.AvatarURLMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") > ?")
		values = append(values, f.AvatarURLMinGt)
	}

	if f.AvatarURLMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") > ?")
		values = append(values, f.AvatarURLMaxGt)
	}

	if f.AvatarURLMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") < ?")
		values = append(values, f.AvatarURLMinLt)
	}

	if f.AvatarURLMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") < ?")
		values = append(values, f.AvatarURLMaxLt)
	}

	if f.AvatarURLMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") >= ?")
		values = append(values, f.AvatarURLMinGte)
	}

	if f.AvatarURLMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") >= ?")
		values = append(values, f.AvatarURLMaxGte)
	}

	if f.AvatarURLMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") <= ?")
		values = append(values, f.AvatarURLMinLte)
	}

	if f.AvatarURLMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") <= ?")
		values = append(values, f.AvatarURLMaxLte)
	}

	if f.AvatarURLMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") IN (?)")
		values = append(values, f.AvatarURLMinIn)
	}

	if f.AvatarURLMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") IN (?)")
		values = append(values, f.AvatarURLMaxIn)
	}

	if f.AvatarURLMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.AvatarURLMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.AvatarURLMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.AvatarURLMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.AvatarURLMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AvatarURLMinPrefix))
	}

	if f.AvatarURLMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AvatarURLMaxPrefix))
	}

	if f.AvatarURLMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AvatarURLMinSuffix))
	}

	if f.AvatarURLMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AvatarURLMaxSuffix))
	}

	if f.DisplayNameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") = ?")
		values = append(values, f.DisplayNameMin)
	}

	if f.DisplayNameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") = ?")
		values = append(values, f.DisplayNameMax)
	}

	if f.DisplayNameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") != ?")
		values = append(values, f.DisplayNameMinNe)
	}

	if f.DisplayNameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") != ?")
		values = append(values, f.DisplayNameMaxNe)
	}

	if f.DisplayNameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") > ?")
		values = append(values, f.DisplayNameMinGt)
	}

	if f.DisplayNameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") > ?")
		values = append(values, f.DisplayNameMaxGt)
	}

	if f.DisplayNameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") < ?")
		values = append(values, f.DisplayNameMinLt)
	}

	if f.DisplayNameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") < ?")
		values = append(values, f.DisplayNameMaxLt)
	}

	if f.DisplayNameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") >= ?")
		values = append(values, f.DisplayNameMinGte)
	}

	if f.DisplayNameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") >= ?")
		values = append(values, f.DisplayNameMaxGte)
	}

	if f.DisplayNameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") <= ?")
		values = append(values, f.DisplayNameMinLte)
	}

	if f.DisplayNameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") <= ?")
		values = append(values, f.DisplayNameMaxLte)
	}

	if f.DisplayNameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") IN (?)")
		values = append(values, f.DisplayNameMinIn)
	}

	if f.DisplayNameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") IN (?)")
		values = append(values, f.DisplayNameMaxIn)
	}

	if f.DisplayNameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DisplayNameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DisplayNameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DisplayNameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DisplayNameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DisplayNameMinPrefix))
	}

	if f.DisplayNameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DisplayNameMaxPrefix))
	}

	if f.DisplayNameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DisplayNameMinSuffix))
	}

	if f.DisplayNameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DisplayNameMaxSuffix))
	}

	if f.FirstNameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") = ?")
		values = append(values, f.FirstNameMin)
	}

	if f.FirstNameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") = ?")
		values = append(values, f.FirstNameMax)
	}

	if f.FirstNameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") != ?")
		values = append(values, f.FirstNameMinNe)
	}

	if f.FirstNameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") != ?")
		values = append(values, f.FirstNameMaxNe)
	}

	if f.FirstNameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") > ?")
		values = append(values, f.FirstNameMinGt)
	}

	if f.FirstNameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") > ?")
		values = append(values, f.FirstNameMaxGt)
	}

	if f.FirstNameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") < ?")
		values = append(values, f.FirstNameMinLt)
	}

	if f.FirstNameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") < ?")
		values = append(values, f.FirstNameMaxLt)
	}

	if f.FirstNameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") >= ?")
		values = append(values, f.FirstNameMinGte)
	}

	if f.FirstNameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") >= ?")
		values = append(values, f.FirstNameMaxGte)
	}

	if f.FirstNameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") <= ?")
		values = append(values, f.FirstNameMinLte)
	}

	if f.FirstNameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") <= ?")
		values = append(values, f.FirstNameMaxLte)
	}

	if f.FirstNameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") IN (?)")
		values = append(values, f.FirstNameMinIn)
	}

	if f.FirstNameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") IN (?)")
		values = append(values, f.FirstNameMaxIn)
	}

	if f.FirstNameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.FirstNameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.FirstNameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.FirstNameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.FirstNameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.FirstNameMinPrefix))
	}

	if f.FirstNameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.FirstNameMaxPrefix))
	}

	if f.FirstNameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.FirstNameMinSuffix))
	}

	if f.FirstNameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.FirstNameMaxSuffix))
	}

	if f.LastNameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") = ?")
		values = append(values, f.LastNameMin)
	}

	if f.LastNameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") = ?")
		values = append(values, f.LastNameMax)
	}

	if f.LastNameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") != ?")
		values = append(values, f.LastNameMinNe)
	}

	if f.LastNameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") != ?")
		values = append(values, f.LastNameMaxNe)
	}

	if f.LastNameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") > ?")
		values = append(values, f.LastNameMinGt)
	}

	if f.LastNameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") > ?")
		values = append(values, f.LastNameMaxGt)
	}

	if f.LastNameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") < ?")
		values = append(values, f.LastNameMinLt)
	}

	if f.LastNameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") < ?")
		values = append(values, f.LastNameMaxLt)
	}

	if f.LastNameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") >= ?")
		values = append(values, f.LastNameMinGte)
	}

	if f.LastNameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") >= ?")
		values = append(values, f.LastNameMaxGte)
	}

	if f.LastNameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") <= ?")
		values = append(values, f.LastNameMinLte)
	}

	if f.LastNameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") <= ?")
		values = append(values, f.LastNameMaxLte)
	}

	if f.LastNameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") IN (?)")
		values = append(values, f.LastNameMinIn)
	}

	if f.LastNameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") IN (?)")
		values = append(values, f.LastNameMaxIn)
	}

	if f.LastNameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LastNameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LastNameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LastNameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LastNameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LastNameMinPrefix))
	}

	if f.LastNameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LastNameMaxPrefix))
	}

	if f.LastNameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LastNameMinSuffix))
	}

	if f.LastNameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LastNameMaxSuffix))
	}

	if f.NickNameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") = ?")
		values = append(values, f.NickNameMin)
	}

	if f.NickNameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") = ?")
		values = append(values, f.NickNameMax)
	}

	if f.NickNameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") != ?")
		values = append(values, f.NickNameMinNe)
	}

	if f.NickNameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") != ?")
		values = append(values, f.NickNameMaxNe)
	}

	if f.NickNameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") > ?")
		values = append(values, f.NickNameMinGt)
	}

	if f.NickNameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") > ?")
		values = append(values, f.NickNameMaxGt)
	}

	if f.NickNameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") < ?")
		values = append(values, f.NickNameMinLt)
	}

	if f.NickNameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") < ?")
		values = append(values, f.NickNameMaxLt)
	}

	if f.NickNameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") >= ?")
		values = append(values, f.NickNameMinGte)
	}

	if f.NickNameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") >= ?")
		values = append(values, f.NickNameMaxGte)
	}

	if f.NickNameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") <= ?")
		values = append(values, f.NickNameMinLte)
	}

	if f.NickNameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") <= ?")
		values = append(values, f.NickNameMaxLte)
	}

	if f.NickNameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") IN (?)")
		values = append(values, f.NickNameMinIn)
	}

	if f.NickNameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") IN (?)")
		values = append(values, f.NickNameMaxIn)
	}

	if f.NickNameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NickNameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NickNameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NickNameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NickNameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NickNameMinPrefix))
	}

	if f.NickNameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NickNameMaxPrefix))
	}

	if f.NickNameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NickNameMinSuffix))
	}

	if f.NickNameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NickNameMaxSuffix))
	}

	if f.DescriptionMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMin)
	}

	if f.DescriptionMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMax)
	}

	if f.DescriptionMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMinNe)
	}

	if f.DescriptionMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMaxNe)
	}

	if f.DescriptionMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMinGt)
	}

	if f.DescriptionMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMaxGt)
	}

	if f.DescriptionMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMinLt)
	}

	if f.DescriptionMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMaxLt)
	}

	if f.DescriptionMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMinGte)
	}

	if f.DescriptionMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMaxGte)
	}

	if f.DescriptionMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMinLte)
	}

	if f.DescriptionMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMaxLte)
	}

	if f.DescriptionMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMinIn)
	}

	if f.DescriptionMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMaxIn)
	}

	if f.DescriptionMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMinPrefix))
	}

	if f.DescriptionMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMaxPrefix))
	}

	if f.DescriptionMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMinSuffix))
	}

	if f.DescriptionMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMaxSuffix))
	}

	if f.LocationMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") = ?")
		values = append(values, f.LocationMin)
	}

	if f.LocationMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") = ?")
		values = append(values, f.LocationMax)
	}

	if f.LocationMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") != ?")
		values = append(values, f.LocationMinNe)
	}

	if f.LocationMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") != ?")
		values = append(values, f.LocationMaxNe)
	}

	if f.LocationMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") > ?")
		values = append(values, f.LocationMinGt)
	}

	if f.LocationMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") > ?")
		values = append(values, f.LocationMaxGt)
	}

	if f.LocationMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") < ?")
		values = append(values, f.LocationMinLt)
	}

	if f.LocationMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") < ?")
		values = append(values, f.LocationMaxLt)
	}

	if f.LocationMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") >= ?")
		values = append(values, f.LocationMinGte)
	}

	if f.LocationMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") >= ?")
		values = append(values, f.LocationMaxGte)
	}

	if f.LocationMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") <= ?")
		values = append(values, f.LocationMinLte)
	}

	if f.LocationMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") <= ?")
		values = append(values, f.LocationMaxLte)
	}

	if f.LocationMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") IN (?)")
		values = append(values, f.LocationMinIn)
	}

	if f.LocationMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") IN (?)")
		values = append(values, f.LocationMaxIn)
	}

	if f.LocationMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LocationMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LocationMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LocationMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LocationMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LocationMinPrefix))
	}

	if f.LocationMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LocationMaxPrefix))
	}

	if f.LocationMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LocationMinSuffix))
	}

	if f.LocationMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LocationMaxSuffix))
	}

	if f.UserIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") = ?")
		values = append(values, f.UserIDMin)
	}

	if f.UserIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") = ?")
		values = append(values, f.UserIDMax)
	}

	if f.UserIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") != ?")
		values = append(values, f.UserIDMinNe)
	}

	if f.UserIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") != ?")
		values = append(values, f.UserIDMaxNe)
	}

	if f.UserIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") > ?")
		values = append(values, f.UserIDMinGt)
	}

	if f.UserIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") > ?")
		values = append(values, f.UserIDMaxGt)
	}

	if f.UserIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") < ?")
		values = append(values, f.UserIDMinLt)
	}

	if f.UserIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") < ?")
		values = append(values, f.UserIDMaxLt)
	}

	if f.UserIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") >= ?")
		values = append(values, f.UserIDMinGte)
	}

	if f.UserIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") >= ?")
		values = append(values, f.UserIDMaxGte)
	}

	if f.UserIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") <= ?")
		values = append(values, f.UserIDMinLte)
	}

	if f.UserIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") <= ?")
		values = append(values, f.UserIDMaxLte)
	}

	if f.UserIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") IN (?)")
		values = append(values, f.UserIDMinIn)
	}

	if f.UserIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") IN (?)")
		values = append(values, f.UserIDMaxIn)
	}

	if f.UserIDMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.UserIDMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.UserIDMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.UserIDMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.UserIDMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.UserIDMinPrefix))
	}

	if f.UserIDMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.UserIDMaxPrefix))
	}

	if f.UserIDMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.UserIDMinSuffix))
	}

	if f.UserIDMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.UserIDMaxSuffix))
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *PersonFilterType) AndWith(f2 ...*PersonFilterType) *PersonFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &PersonFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *PersonFilterType) OrWith(f2 ...*PersonFilterType) *PersonFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &PersonFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *DeliveryTypeFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *DeliveryTypeFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("delivery_types"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *DeliveryTypeFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Delivery != nil {
		_alias := alias + "_delivery"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("deliveries"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("deliveryId"))
		err := f.Delivery.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *DeliveryTypeFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NOT NULL")
		}
	}

	if f.Description != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" = ?")
		values = append(values, f.Description)
	}

	if f.DescriptionNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" != ?")
		values = append(values, f.DescriptionNe)
	}

	if f.DescriptionGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" > ?")
		values = append(values, f.DescriptionGt)
	}

	if f.DescriptionLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" < ?")
		values = append(values, f.DescriptionLt)
	}

	if f.DescriptionGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" >= ?")
		values = append(values, f.DescriptionGte)
	}

	if f.DescriptionLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" <= ?")
		values = append(values, f.DescriptionLte)
	}

	if f.DescriptionIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IN (?)")
		values = append(values, f.DescriptionIn)
	}

	if f.DescriptionLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionPrefix))
	}

	if f.DescriptionSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionSuffix))
	}

	if f.DescriptionNull != nil {
		if *f.DescriptionNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NOT NULL")
		}
	}

	if f.DeliveryID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" = ?")
		values = append(values, f.DeliveryID)
	}

	if f.DeliveryIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" != ?")
		values = append(values, f.DeliveryIDNe)
	}

	if f.DeliveryIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" > ?")
		values = append(values, f.DeliveryIDGt)
	}

	if f.DeliveryIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" < ?")
		values = append(values, f.DeliveryIDLt)
	}

	if f.DeliveryIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" >= ?")
		values = append(values, f.DeliveryIDGte)
	}

	if f.DeliveryIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" <= ?")
		values = append(values, f.DeliveryIDLte)
	}

	if f.DeliveryIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" IN (?)")
		values = append(values, f.DeliveryIDIn)
	}

	if f.DeliveryIDNull != nil {
		if *f.DeliveryIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *DeliveryTypeFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.NameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMin)
	}

	if f.NameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMax)
	}

	if f.NameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMinNe)
	}

	if f.NameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMaxNe)
	}

	if f.NameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMinGt)
	}

	if f.NameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMaxGt)
	}

	if f.NameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMinLt)
	}

	if f.NameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMaxLt)
	}

	if f.NameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMinGte)
	}

	if f.NameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMaxGte)
	}

	if f.NameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMinLte)
	}

	if f.NameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMaxLte)
	}

	if f.NameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMinIn)
	}

	if f.NameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMaxIn)
	}

	if f.NameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMinPrefix))
	}

	if f.NameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMaxPrefix))
	}

	if f.NameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMinSuffix))
	}

	if f.NameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMaxSuffix))
	}

	if f.DescriptionMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMin)
	}

	if f.DescriptionMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMax)
	}

	if f.DescriptionMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMinNe)
	}

	if f.DescriptionMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMaxNe)
	}

	if f.DescriptionMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMinGt)
	}

	if f.DescriptionMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMaxGt)
	}

	if f.DescriptionMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMinLt)
	}

	if f.DescriptionMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMaxLt)
	}

	if f.DescriptionMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMinGte)
	}

	if f.DescriptionMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMaxGte)
	}

	if f.DescriptionMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMinLte)
	}

	if f.DescriptionMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMaxLte)
	}

	if f.DescriptionMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMinIn)
	}

	if f.DescriptionMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMaxIn)
	}

	if f.DescriptionMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMinPrefix))
	}

	if f.DescriptionMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMaxPrefix))
	}

	if f.DescriptionMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMinSuffix))
	}

	if f.DescriptionMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMaxSuffix))
	}

	if f.DeliveryIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") = ?")
		values = append(values, f.DeliveryIDMin)
	}

	if f.DeliveryIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") = ?")
		values = append(values, f.DeliveryIDMax)
	}

	if f.DeliveryIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") != ?")
		values = append(values, f.DeliveryIDMinNe)
	}

	if f.DeliveryIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") != ?")
		values = append(values, f.DeliveryIDMaxNe)
	}

	if f.DeliveryIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") > ?")
		values = append(values, f.DeliveryIDMinGt)
	}

	if f.DeliveryIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") > ?")
		values = append(values, f.DeliveryIDMaxGt)
	}

	if f.DeliveryIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") < ?")
		values = append(values, f.DeliveryIDMinLt)
	}

	if f.DeliveryIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") < ?")
		values = append(values, f.DeliveryIDMaxLt)
	}

	if f.DeliveryIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") >= ?")
		values = append(values, f.DeliveryIDMinGte)
	}

	if f.DeliveryIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") >= ?")
		values = append(values, f.DeliveryIDMaxGte)
	}

	if f.DeliveryIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") <= ?")
		values = append(values, f.DeliveryIDMinLte)
	}

	if f.DeliveryIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") <= ?")
		values = append(values, f.DeliveryIDMaxLte)
	}

	if f.DeliveryIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") IN (?)")
		values = append(values, f.DeliveryIDMinIn)
	}

	if f.DeliveryIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") IN (?)")
		values = append(values, f.DeliveryIDMaxIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *DeliveryTypeFilterType) AndWith(f2 ...*DeliveryTypeFilterType) *DeliveryTypeFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &DeliveryTypeFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *DeliveryTypeFilterType) OrWith(f2 ...*DeliveryTypeFilterType) *DeliveryTypeFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &DeliveryTypeFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *DeliveryChannelFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *DeliveryChannelFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("delivery_channels"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *DeliveryChannelFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Delivery != nil {
		_alias := alias + "_delivery"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("deliveries"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("deliveryId"))
		err := f.Delivery.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *DeliveryChannelFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NOT NULL")
		}
	}

	if f.Description != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" = ?")
		values = append(values, f.Description)
	}

	if f.DescriptionNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" != ?")
		values = append(values, f.DescriptionNe)
	}

	if f.DescriptionGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" > ?")
		values = append(values, f.DescriptionGt)
	}

	if f.DescriptionLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" < ?")
		values = append(values, f.DescriptionLt)
	}

	if f.DescriptionGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" >= ?")
		values = append(values, f.DescriptionGte)
	}

	if f.DescriptionLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" <= ?")
		values = append(values, f.DescriptionLte)
	}

	if f.DescriptionIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IN (?)")
		values = append(values, f.DescriptionIn)
	}

	if f.DescriptionLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionPrefix))
	}

	if f.DescriptionSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionSuffix))
	}

	if f.DescriptionNull != nil {
		if *f.DescriptionNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NOT NULL")
		}
	}

	if f.DeliveryID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" = ?")
		values = append(values, f.DeliveryID)
	}

	if f.DeliveryIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" != ?")
		values = append(values, f.DeliveryIDNe)
	}

	if f.DeliveryIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" > ?")
		values = append(values, f.DeliveryIDGt)
	}

	if f.DeliveryIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" < ?")
		values = append(values, f.DeliveryIDLt)
	}

	if f.DeliveryIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" >= ?")
		values = append(values, f.DeliveryIDGte)
	}

	if f.DeliveryIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" <= ?")
		values = append(values, f.DeliveryIDLte)
	}

	if f.DeliveryIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" IN (?)")
		values = append(values, f.DeliveryIDIn)
	}

	if f.DeliveryIDNull != nil {
		if *f.DeliveryIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *DeliveryChannelFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.NameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMin)
	}

	if f.NameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMax)
	}

	if f.NameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMinNe)
	}

	if f.NameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMaxNe)
	}

	if f.NameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMinGt)
	}

	if f.NameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMaxGt)
	}

	if f.NameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMinLt)
	}

	if f.NameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMaxLt)
	}

	if f.NameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMinGte)
	}

	if f.NameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMaxGte)
	}

	if f.NameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMinLte)
	}

	if f.NameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMaxLte)
	}

	if f.NameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMinIn)
	}

	if f.NameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMaxIn)
	}

	if f.NameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMinPrefix))
	}

	if f.NameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMaxPrefix))
	}

	if f.NameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMinSuffix))
	}

	if f.NameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMaxSuffix))
	}

	if f.DescriptionMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMin)
	}

	if f.DescriptionMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMax)
	}

	if f.DescriptionMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMinNe)
	}

	if f.DescriptionMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMaxNe)
	}

	if f.DescriptionMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMinGt)
	}

	if f.DescriptionMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMaxGt)
	}

	if f.DescriptionMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMinLt)
	}

	if f.DescriptionMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMaxLt)
	}

	if f.DescriptionMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMinGte)
	}

	if f.DescriptionMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMaxGte)
	}

	if f.DescriptionMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMinLte)
	}

	if f.DescriptionMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMaxLte)
	}

	if f.DescriptionMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMinIn)
	}

	if f.DescriptionMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMaxIn)
	}

	if f.DescriptionMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMinPrefix))
	}

	if f.DescriptionMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMaxPrefix))
	}

	if f.DescriptionMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMinSuffix))
	}

	if f.DescriptionMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMaxSuffix))
	}

	if f.DeliveryIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") = ?")
		values = append(values, f.DeliveryIDMin)
	}

	if f.DeliveryIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") = ?")
		values = append(values, f.DeliveryIDMax)
	}

	if f.DeliveryIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") != ?")
		values = append(values, f.DeliveryIDMinNe)
	}

	if f.DeliveryIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") != ?")
		values = append(values, f.DeliveryIDMaxNe)
	}

	if f.DeliveryIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") > ?")
		values = append(values, f.DeliveryIDMinGt)
	}

	if f.DeliveryIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") > ?")
		values = append(values, f.DeliveryIDMaxGt)
	}

	if f.DeliveryIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") < ?")
		values = append(values, f.DeliveryIDMinLt)
	}

	if f.DeliveryIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") < ?")
		values = append(values, f.DeliveryIDMaxLt)
	}

	if f.DeliveryIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") >= ?")
		values = append(values, f.DeliveryIDMinGte)
	}

	if f.DeliveryIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") >= ?")
		values = append(values, f.DeliveryIDMaxGte)
	}

	if f.DeliveryIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") <= ?")
		values = append(values, f.DeliveryIDMinLte)
	}

	if f.DeliveryIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") <= ?")
		values = append(values, f.DeliveryIDMaxLte)
	}

	if f.DeliveryIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") IN (?)")
		values = append(values, f.DeliveryIDMinIn)
	}

	if f.DeliveryIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") IN (?)")
		values = append(values, f.DeliveryIDMaxIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *DeliveryChannelFilterType) AndWith(f2 ...*DeliveryChannelFilterType) *DeliveryChannelFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &DeliveryChannelFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *DeliveryChannelFilterType) OrWith(f2 ...*DeliveryChannelFilterType) *DeliveryChannelFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &DeliveryChannelFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *VehicleTypeFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *VehicleTypeFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("vehicle_types"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *VehicleTypeFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Delivery != nil {
		_alias := alias + "_delivery"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("deliveries"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("deliveryId"))
		err := f.Delivery.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *VehicleTypeFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NOT NULL")
		}
	}

	if f.Description != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" = ?")
		values = append(values, f.Description)
	}

	if f.DescriptionNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" != ?")
		values = append(values, f.DescriptionNe)
	}

	if f.DescriptionGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" > ?")
		values = append(values, f.DescriptionGt)
	}

	if f.DescriptionLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" < ?")
		values = append(values, f.DescriptionLt)
	}

	if f.DescriptionGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" >= ?")
		values = append(values, f.DescriptionGte)
	}

	if f.DescriptionLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" <= ?")
		values = append(values, f.DescriptionLte)
	}

	if f.DescriptionIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IN (?)")
		values = append(values, f.DescriptionIn)
	}

	if f.DescriptionLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionPrefix))
	}

	if f.DescriptionSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionSuffix))
	}

	if f.DescriptionNull != nil {
		if *f.DescriptionNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NOT NULL")
		}
	}

	if f.DeliveryID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" = ?")
		values = append(values, f.DeliveryID)
	}

	if f.DeliveryIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" != ?")
		values = append(values, f.DeliveryIDNe)
	}

	if f.DeliveryIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" > ?")
		values = append(values, f.DeliveryIDGt)
	}

	if f.DeliveryIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" < ?")
		values = append(values, f.DeliveryIDLt)
	}

	if f.DeliveryIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" >= ?")
		values = append(values, f.DeliveryIDGte)
	}

	if f.DeliveryIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" <= ?")
		values = append(values, f.DeliveryIDLte)
	}

	if f.DeliveryIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" IN (?)")
		values = append(values, f.DeliveryIDIn)
	}

	if f.DeliveryIDNull != nil {
		if *f.DeliveryIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("deliveryId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *VehicleTypeFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.NameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMin)
	}

	if f.NameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMax)
	}

	if f.NameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMinNe)
	}

	if f.NameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMaxNe)
	}

	if f.NameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMinGt)
	}

	if f.NameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMaxGt)
	}

	if f.NameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMinLt)
	}

	if f.NameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMaxLt)
	}

	if f.NameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMinGte)
	}

	if f.NameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMaxGte)
	}

	if f.NameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMinLte)
	}

	if f.NameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMaxLte)
	}

	if f.NameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMinIn)
	}

	if f.NameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMaxIn)
	}

	if f.NameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMinPrefix))
	}

	if f.NameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMaxPrefix))
	}

	if f.NameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMinSuffix))
	}

	if f.NameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMaxSuffix))
	}

	if f.DescriptionMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMin)
	}

	if f.DescriptionMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMax)
	}

	if f.DescriptionMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMinNe)
	}

	if f.DescriptionMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMaxNe)
	}

	if f.DescriptionMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMinGt)
	}

	if f.DescriptionMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMaxGt)
	}

	if f.DescriptionMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMinLt)
	}

	if f.DescriptionMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMaxLt)
	}

	if f.DescriptionMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMinGte)
	}

	if f.DescriptionMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMaxGte)
	}

	if f.DescriptionMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMinLte)
	}

	if f.DescriptionMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMaxLte)
	}

	if f.DescriptionMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMinIn)
	}

	if f.DescriptionMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMaxIn)
	}

	if f.DescriptionMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMinPrefix))
	}

	if f.DescriptionMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMaxPrefix))
	}

	if f.DescriptionMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMinSuffix))
	}

	if f.DescriptionMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMaxSuffix))
	}

	if f.DeliveryIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") = ?")
		values = append(values, f.DeliveryIDMin)
	}

	if f.DeliveryIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") = ?")
		values = append(values, f.DeliveryIDMax)
	}

	if f.DeliveryIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") != ?")
		values = append(values, f.DeliveryIDMinNe)
	}

	if f.DeliveryIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") != ?")
		values = append(values, f.DeliveryIDMaxNe)
	}

	if f.DeliveryIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") > ?")
		values = append(values, f.DeliveryIDMinGt)
	}

	if f.DeliveryIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") > ?")
		values = append(values, f.DeliveryIDMaxGt)
	}

	if f.DeliveryIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") < ?")
		values = append(values, f.DeliveryIDMinLt)
	}

	if f.DeliveryIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") < ?")
		values = append(values, f.DeliveryIDMaxLt)
	}

	if f.DeliveryIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") >= ?")
		values = append(values, f.DeliveryIDMinGte)
	}

	if f.DeliveryIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") >= ?")
		values = append(values, f.DeliveryIDMaxGte)
	}

	if f.DeliveryIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") <= ?")
		values = append(values, f.DeliveryIDMinLte)
	}

	if f.DeliveryIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") <= ?")
		values = append(values, f.DeliveryIDMaxLte)
	}

	if f.DeliveryIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("deliveryId")+") IN (?)")
		values = append(values, f.DeliveryIDMinIn)
	}

	if f.DeliveryIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("deliveryId")+") IN (?)")
		values = append(values, f.DeliveryIDMaxIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *VehicleTypeFilterType) AndWith(f2 ...*VehicleTypeFilterType) *VehicleTypeFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &VehicleTypeFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *VehicleTypeFilterType) OrWith(f2 ...*VehicleTypeFilterType) *VehicleTypeFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &VehicleTypeFilterType{
		Or: append(_f2, f),
	}
}
