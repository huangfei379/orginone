// Code generated by entc, DO NOT EDIT.

package schema

import (
	"fmt"
	"orginone/common/schema/asrole"
	"orginone/common/tools/date"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// AsRole is the model entity for the AsRole schema.
type AsRole struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID int64 `json:"id,string"`
	// Sort holds the value of the "sort" field.
	// 排序
	Sort int64 `json:"sort"`
	// RoleAlias holds the value of the "role_alias" field.
	// 角色别名角色名的首字母拼音),可用于接口鉴权,自动生成
	RoleAlias string `json:"roleAlias"`
	// RoleName holds the value of the "role_name" field.
	// 角色名
	RoleName string `json:"roleName"`
	// RoleDescription holds the value of the "role_description" field.
	// 角色描述
	RoleDescription string `json:"roleDescription"`
	// IsDeleted holds the value of the "is_deleted" field.
	// 是否被删除
	IsDeleted int64 `json:"isDeleted"`
	// Status holds the value of the "status" field.
	// 状态
	Status int64 `json:"status"`
	// CreateUser holds the value of the "create_user" field.
	// 创建用户
	CreateUser int64 `json:"createUser"`
	// UpdateUser holds the value of the "update_user" field.
	// 更新用户
	UpdateUser int64 `json:"updateUser"`
	// CreateTime holds the value of the "create_time" field.
	// 创建时间
	CreateTime date.DateTime `json:"createTime"`
	// UpdateTime holds the value of the "update_time" field.
	// 更新时间
	UpdateTime date.DateTime `json:"updateTime"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AsRoleQuery when eager-loading is set.
	Edges AsRoleEdges `json:"edges"`
}

// AsRoleEdges holds the relations/edges for other nodes in the graph.
type AsRoleEdges struct {
	// Users holds the value of the users edge.
	Users []*AsUser `json:"users"`
	// Jobs holds the value of the jobs edge.
	Jobs []*AsJob `json:"jobs"`
	// Menus holds the value of the menus edge.
	Menus []*AsMenu `json:"menus"`
	// AttrRoles holds the value of the attrRoles edge.
	AttrRoles []*AsTenantAttrRole `json:"attrroles"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e AsRoleEdges) UsersOrErr() ([]*AsUser, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// JobsOrErr returns the Jobs value or an error if the edge
// was not loaded in eager-loading.
func (e AsRoleEdges) JobsOrErr() ([]*AsJob, error) {
	if e.loadedTypes[1] {
		return e.Jobs, nil
	}
	return nil, &NotLoadedError{edge: "jobs"}
}

// MenusOrErr returns the Menus value or an error if the edge
// was not loaded in eager-loading.
func (e AsRoleEdges) MenusOrErr() ([]*AsMenu, error) {
	if e.loadedTypes[2] {
		return e.Menus, nil
	}
	return nil, &NotLoadedError{edge: "menus"}
}

// AttrRolesOrErr returns the AttrRoles value or an error if the edge
// was not loaded in eager-loading.
func (e AsRoleEdges) AttrRolesOrErr() ([]*AsTenantAttrRole, error) {
	if e.loadedTypes[3] {
		return e.AttrRoles, nil
	}
	return nil, &NotLoadedError{edge: "attrRoles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AsRole) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case asrole.FieldID, asrole.FieldSort, asrole.FieldIsDeleted, asrole.FieldStatus, asrole.FieldCreateUser, asrole.FieldUpdateUser:
			values[i] = new(sql.NullInt64)
		case asrole.FieldRoleAlias, asrole.FieldRoleName, asrole.FieldRoleDescription:
			values[i] = new(sql.NullString)
		case asrole.FieldCreateTime, asrole.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AsRole", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AsRole fields.
func (ar *AsRole) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case asrole.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ar.ID = int64(value.Int64)
		case asrole.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				ar.Sort = value.Int64
			}
		case asrole.FieldRoleAlias:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role_alias", values[i])
			} else if value.Valid {
				ar.RoleAlias = value.String
			}
		case asrole.FieldRoleName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role_name", values[i])
			} else if value.Valid {
				ar.RoleName = value.String
			}
		case asrole.FieldRoleDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role_description", values[i])
			} else if value.Valid {
				ar.RoleDescription = value.String
			}
		case asrole.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				ar.IsDeleted = value.Int64
			}
		case asrole.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ar.Status = value.Int64
			}
		case asrole.FieldCreateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_user", values[i])
			} else if value.Valid {
				ar.CreateUser = value.Int64
			}
		case asrole.FieldUpdateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_user", values[i])
			} else if value.Valid {
				ar.UpdateUser = value.Int64
			}
		case asrole.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ar.CreateTime = date.DateTime(value.Time)
			}
		case asrole.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ar.UpdateTime = date.DateTime(value.Time)
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the AsRole entity.
func (ar *AsRole) QueryUsers() *AsUserQuery {
	return (&AsRoleClient{config: ar.config}).QueryUsers(ar)
}

// QueryJobs queries the "jobs" edge of the AsRole entity.
func (ar *AsRole) QueryJobs() *AsJobQuery {
	return (&AsRoleClient{config: ar.config}).QueryJobs(ar)
}

// QueryMenus queries the "menus" edge of the AsRole entity.
func (ar *AsRole) QueryMenus() *AsMenuQuery {
	return (&AsRoleClient{config: ar.config}).QueryMenus(ar)
}

// QueryAttrRoles queries the "attrRoles" edge of the AsRole entity.
func (ar *AsRole) QueryAttrRoles() *AsTenantAttrRoleQuery {
	return (&AsRoleClient{config: ar.config}).QueryAttrRoles(ar)
}

// Update returns a builder for updating this AsRole.
// Note that you need to call AsRole.Unwrap() before calling this method if this AsRole
// was returned from a transaction, and the transaction was committed or rolled back.
func (ar *AsRole) Update() *AsRoleUpdateOne {
	return (&AsRoleClient{config: ar.config}).UpdateOne(ar)
}

// Unwrap unwraps the AsRole entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ar *AsRole) Unwrap() *AsRole {
	tx, ok := ar.config.driver.(*txDriver)
	if !ok {
		panic("schema: AsRole is not a transactional entity")
	}
	ar.config.driver = tx.drv
	return ar
}

// String implements the fmt.Stringer.
func (ar *AsRole) String() string {
	var builder strings.Builder
	builder.WriteString("AsRole(")
	builder.WriteString(fmt.Sprintf("id=%v", ar.ID))
	builder.WriteString(", sort=")
	builder.WriteString(fmt.Sprintf("%v", ar.Sort))
	builder.WriteString(", role_alias=")
	builder.WriteString(ar.RoleAlias)
	builder.WriteString(", role_name=")
	builder.WriteString(ar.RoleName)
	builder.WriteString(", role_description=")
	builder.WriteString(ar.RoleDescription)
	builder.WriteString(", is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", ar.IsDeleted))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", ar.Status))
	builder.WriteString(", create_user=")
	builder.WriteString(fmt.Sprintf("%v", ar.CreateUser))
	builder.WriteString(", update_user=")
	builder.WriteString(fmt.Sprintf("%v", ar.UpdateUser))
	builder.WriteString(", create_time=")
	builder.WriteString(fmt.Sprintf("%v", ar.CreateTime))
	builder.WriteString(", update_time=")
	builder.WriteString(fmt.Sprintf("%v", ar.UpdateTime))
	builder.WriteByte(')')
	return builder.String()
}

// AsRoles is a parsable slice of AsRole.
type AsRoles []*AsRole

func (ar AsRoles) config(cfg config) {
	for _i := range ar {
		ar[_i].config = cfg
	}
}
