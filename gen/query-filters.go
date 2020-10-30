package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/vektah/gqlparser/v2/ast"
)

// DeliveryQueryFilter struct
type DeliveryQueryFilter struct {
	Query *string
}

// Apply ...
func (qf *DeliveryQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(dialect, fields, part, TableName("deliveries"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *DeliveryQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["paymentTotal"]; ok {

		cast := "TEXT"
		if dialect.GetName() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+dialect.Quote("paymentTotal")+" AS %s)", dialect.Quote(alias)+".", cast)

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["collectAddress"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("collectAddress")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["collectPoint"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("collectPoint")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["dropAddress"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("dropAddress")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["dropPoint"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("dropPoint")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["expectedDistance"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("expectedDistance")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["expectedCost"]; ok {

		cast := "TEXT"
		if dialect.GetName() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+dialect.Quote("expectedCost")+" AS %s)", dialect.Quote(alias)+".", cast)

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["smsToken"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("smsToken")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["status"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("status")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if fs, ok := fieldsMap["sender"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_sender"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("people"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("senderId"))

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := PersonQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if fs, ok := fieldsMap["receiver"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_receiver"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("people"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("receiverId"))

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := PersonQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if fs, ok := fieldsMap["deliver"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_deliver"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("people"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("deliverId"))

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := PersonQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if fs, ok := fieldsMap["vehicleType"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_vehicleType"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("vehicle_types"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("vehicleTypeId"))

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := VehicleTypeQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if fs, ok := fieldsMap["deliveryType"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_deliveryType"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("delivery_types"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("deliveryTypeId"))

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := DeliveryTypeQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if fs, ok := fieldsMap["deliveryChannel"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_deliveryChannel"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("delivery_channels"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("deliveryChannelId"))

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := DeliveryChannelQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// PersonQueryFilter struct
type PersonQueryFilter struct {
	Query *string
}

// Apply ...
func (qf *PersonQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(dialect, fields, part, TableName("people"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *PersonQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["email"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("email")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["phone"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("phone")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["documentNo"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("documentNo")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["avatarURL"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("avatarURL")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["displayName"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("displayName")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["firstName"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("firstName")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["lastName"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("lastName")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["nickName"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("nickName")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["description"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("description")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["location"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("location")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["userId"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("userId")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if fs, ok := fieldsMap["deliveries"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_deliveries"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("deliveries"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("deliverId")+" = "+dialect.Quote(alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := DeliveryQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if fs, ok := fieldsMap["deliveriesSent"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_deliveriesSent"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("deliveries"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("senderId")+" = "+dialect.Quote(alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := DeliveryQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if fs, ok := fieldsMap["deliveriesReceived"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_deliveriesReceived"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("deliveries"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("receiverId")+" = "+dialect.Quote(alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := DeliveryQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// DeliveryTypeQueryFilter struct
type DeliveryTypeQueryFilter struct {
	Query *string
}

// Apply ...
func (qf *DeliveryTypeQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(dialect, fields, part, TableName("delivery_types"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *DeliveryTypeQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["name"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("name")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["description"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("description")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if fs, ok := fieldsMap["delivery"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_delivery"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("deliveries"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("deliveryId"))

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := DeliveryQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// DeliveryChannelQueryFilter struct
type DeliveryChannelQueryFilter struct {
	Query *string
}

// Apply ...
func (qf *DeliveryChannelQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(dialect, fields, part, TableName("delivery_channels"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *DeliveryChannelQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["name"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("name")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["description"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("description")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if fs, ok := fieldsMap["delivery"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_delivery"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("deliveries"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("deliveryId"))

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := DeliveryQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// VehicleTypeQueryFilter struct
type VehicleTypeQueryFilter struct {
	Query *string
}

// Apply ...
func (qf *VehicleTypeQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(dialect, fields, part, TableName("vehicle_types"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *VehicleTypeQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["name"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("name")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["description"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("description")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if fs, ok := fieldsMap["delivery"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_delivery"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("deliveries"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("deliveryId"))

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := DeliveryQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
