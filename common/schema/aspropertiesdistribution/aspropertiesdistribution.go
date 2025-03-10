// Code generated by entc, DO NOT EDIT.

package aspropertiesdistribution

import (
	"orginone/common/tools/date"
)

const (
	// Label holds the string label denoting the aspropertiesdistribution type in the database.
	Label = "as_properties_distribution"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPropertiesID holds the string denoting the properties_id field in the database.
	FieldPropertiesID = "properties_id"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "is_deleted"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreateUser holds the string denoting the create_user field in the database.
	FieldCreateUser = "create_user"
	// FieldUpdateUser holds the string denoting the update_user field in the database.
	FieldUpdateUser = "update_user"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// EdgeProperties holds the string denoting the properties edge name in mutations.
	EdgeProperties = "properties"
	// EdgeTenant holds the string denoting the tenant edge name in mutations.
	EdgeTenant = "tenant"
	// Table holds the table name of the aspropertiesdistribution in the database.
	Table = "as_properties_distribution"
	// PropertiesTable is the table that holds the properties relation/edge.
	PropertiesTable = "as_properties_distribution"
	// PropertiesInverseTable is the table name for the AsProperties entity.
	// It exists in this package in order to avoid circular dependency with the "asproperties" package.
	PropertiesInverseTable = "as_properties"
	// PropertiesColumn is the table column denoting the properties relation/edge.
	PropertiesColumn = "properties_id"
	// TenantTable is the table that holds the tenant relation/edge.
	TenantTable = "as_properties_distribution"
	// TenantInverseTable is the table name for the AsTenant entity.
	// It exists in this package in order to avoid circular dependency with the "astenant" package.
	TenantInverseTable = "as_tenant"
	// TenantColumn is the table column denoting the tenant relation/edge.
	TenantColumn = "tenant_id"
)

// Columns holds all SQL columns for aspropertiesdistribution fields.
var Columns = []string{
	FieldID,
	FieldPropertiesID,
	FieldTenantID,
	FieldIsDeleted,
	FieldStatus,
	FieldCreateUser,
	FieldUpdateUser,
	FieldCreateTime,
	FieldUpdateTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsDeleted holds the default value on creation for the "is_deleted" field.
	DefaultIsDeleted int64
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int64
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() date.DateTime
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() date.DateTime
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() date.DateTime
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)
