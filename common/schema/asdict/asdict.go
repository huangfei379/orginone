// Code generated by entc, DO NOT EDIT.

package asdict

import (
	"orginone/common/tools/date"
)

const (
	// Label holds the string label denoting the asdict type in the database.
	Label = "as_dict"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldDictKey holds the string denoting the dict_key field in the database.
	FieldDictKey = "dict_key"
	// FieldDictValue holds the string denoting the dict_value field in the database.
	FieldDictValue = "dict_value"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldCurrversion holds the string denoting the currversion field in the database.
	FieldCurrversion = "currversion"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldDictparentID holds the string denoting the dictparent_id field in the database.
	FieldDictparentID = "dictparent_id"
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
	// EdgeParentx holds the string denoting the parentx edge name in mutations.
	EdgeParentx = "parentx"
	// EdgeChildrens holds the string denoting the childrens edge name in mutations.
	EdgeChildrens = "childrens"
	// Table holds the table name of the asdict in the database.
	Table = "as_dict"
	// ParentxTable is the table that holds the parentx relation/edge.
	ParentxTable = "as_dict"
	// ParentxColumn is the table column denoting the parentx relation/edge.
	ParentxColumn = "parent_id"
	// ChildrensTable is the table that holds the childrens relation/edge.
	ChildrensTable = "as_dict"
	// ChildrensColumn is the table column denoting the childrens relation/edge.
	ChildrensColumn = "parent_id"
)

// Columns holds all SQL columns for asdict fields.
var Columns = []string{
	FieldID,
	FieldParentID,
	FieldCode,
	FieldDictKey,
	FieldDictValue,
	FieldSort,
	FieldRemark,
	FieldCurrversion,
	FieldVersion,
	FieldDictparentID,
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
	// DefaultVersion holds the default value on creation for the "version" field.
	DefaultVersion int64
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
