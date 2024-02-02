// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gitlab.hs-flensburg.de/gitlab-classroom/model/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newAssignmentInvitation(db *gorm.DB, opts ...gen.DOOption) assignmentInvitation {
	_assignmentInvitation := assignmentInvitation{}

	_assignmentInvitation.assignmentInvitationDo.UseDB(db, opts...)
	_assignmentInvitation.assignmentInvitationDo.UseModel(&database.AssignmentInvitation{})

	tableName := _assignmentInvitation.assignmentInvitationDo.TableName()
	_assignmentInvitation.ALL = field.NewAsterisk(tableName)
	_assignmentInvitation.ID = field.NewField(tableName, "id")
	_assignmentInvitation.CreatedAt = field.NewTime(tableName, "created_at")
	_assignmentInvitation.UpdatedAt = field.NewTime(tableName, "updated_at")
	_assignmentInvitation.Status = field.NewUint8(tableName, "status")
	_assignmentInvitation.AssignmentID = field.NewField(tableName, "assignment_id")
	_assignmentInvitation.UserID = field.NewInt(tableName, "user_id")
	_assignmentInvitation.Enabled = field.NewBool(tableName, "enabled")
	_assignmentInvitation.ExpiryDate = field.NewTime(tableName, "expiry_date")
	_assignmentInvitation.Assignment = assignmentInvitationBelongsToAssignment{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Assignment", "database.Classroom"),
		Owner: struct {
			field.RelationField
			Classrooms struct {
				field.RelationField
				User struct {
					field.RelationField
				}
				Classroom struct {
					field.RelationField
				}
			}
			AssignmentRepositories struct {
				field.RelationField
				Assignment struct {
					field.RelationField
					Classroom struct {
						field.RelationField
					}
					Projects struct {
						field.RelationField
					}
					Invitations struct {
						field.RelationField
						Assignment struct {
							field.RelationField
						}
						User struct {
							field.RelationField
						}
					}
				}
				User struct {
					field.RelationField
				}
			}
		}{
			RelationField: field.NewRelation("Assignment.Owner", "database.User"),
			Classrooms: struct {
				field.RelationField
				User struct {
					field.RelationField
				}
				Classroom struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Assignment.Owner.Classrooms", "database.UserClassrooms"),
				User: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Assignment.Owner.Classrooms.User", "database.User"),
				},
				Classroom: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Assignment.Owner.Classrooms.Classroom", "database.Classroom"),
				},
			},
			AssignmentRepositories: struct {
				field.RelationField
				Assignment struct {
					field.RelationField
					Classroom struct {
						field.RelationField
					}
					Projects struct {
						field.RelationField
					}
					Invitations struct {
						field.RelationField
						Assignment struct {
							field.RelationField
						}
						User struct {
							field.RelationField
						}
					}
				}
				User struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Assignment.Owner.AssignmentRepositories", "database.AssignmentProjects"),
				Assignment: struct {
					field.RelationField
					Classroom struct {
						field.RelationField
					}
					Projects struct {
						field.RelationField
					}
					Invitations struct {
						field.RelationField
						Assignment struct {
							field.RelationField
						}
						User struct {
							field.RelationField
						}
					}
				}{
					RelationField: field.NewRelation("Assignment.Owner.AssignmentRepositories.Assignment", "database.Assignment"),
					Classroom: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Assignment.Owner.AssignmentRepositories.Assignment.Classroom", "database.Classroom"),
					},
					Projects: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Assignment.Owner.AssignmentRepositories.Assignment.Projects", "database.AssignmentProjects"),
					},
					Invitations: struct {
						field.RelationField
						Assignment struct {
							field.RelationField
						}
						User struct {
							field.RelationField
						}
					}{
						RelationField: field.NewRelation("Assignment.Owner.AssignmentRepositories.Assignment.Invitations", "database.AssignmentInvitation"),
						Assignment: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("Assignment.Owner.AssignmentRepositories.Assignment.Invitations.Assignment", "database.Classroom"),
						},
						User: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("Assignment.Owner.AssignmentRepositories.Assignment.Invitations.User", "database.User"),
						},
					},
				},
				User: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Assignment.Owner.AssignmentRepositories.User", "database.User"),
				},
			},
		},
		Member: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Assignment.Member", "database.UserClassrooms"),
		},
		Assignments: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Assignment.Assignments", "database.Assignment"),
		},
		Invitations: struct {
			field.RelationField
			Classroom struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Assignment.Invitations", "database.ClassroomInvitation"),
			Classroom: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Assignment.Invitations.Classroom", "database.Classroom"),
			},
		},
	}

	_assignmentInvitation.User = assignmentInvitationBelongsToUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "database.User"),
	}

	_assignmentInvitation.fillFieldMap()

	return _assignmentInvitation
}

type assignmentInvitation struct {
	assignmentInvitationDo

	ALL          field.Asterisk
	ID           field.Field
	CreatedAt    field.Time
	UpdatedAt    field.Time
	Status       field.Uint8
	AssignmentID field.Field
	UserID       field.Int
	Enabled      field.Bool
	ExpiryDate   field.Time
	Assignment   assignmentInvitationBelongsToAssignment

	User assignmentInvitationBelongsToUser

	fieldMap map[string]field.Expr
}

func (a assignmentInvitation) Table(newTableName string) *assignmentInvitation {
	a.assignmentInvitationDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a assignmentInvitation) As(alias string) *assignmentInvitation {
	a.assignmentInvitationDo.DO = *(a.assignmentInvitationDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *assignmentInvitation) updateTableName(table string) *assignmentInvitation {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewField(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.Status = field.NewUint8(table, "status")
	a.AssignmentID = field.NewField(table, "assignment_id")
	a.UserID = field.NewInt(table, "user_id")
	a.Enabled = field.NewBool(table, "enabled")
	a.ExpiryDate = field.NewTime(table, "expiry_date")

	a.fillFieldMap()

	return a
}

func (a *assignmentInvitation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *assignmentInvitation) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 10)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["status"] = a.Status
	a.fieldMap["assignment_id"] = a.AssignmentID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["enabled"] = a.Enabled
	a.fieldMap["expiry_date"] = a.ExpiryDate

}

func (a assignmentInvitation) clone(db *gorm.DB) assignmentInvitation {
	a.assignmentInvitationDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a assignmentInvitation) replaceDB(db *gorm.DB) assignmentInvitation {
	a.assignmentInvitationDo.ReplaceDB(db)
	return a
}

type assignmentInvitationBelongsToAssignment struct {
	db *gorm.DB

	field.RelationField

	Owner struct {
		field.RelationField
		Classrooms struct {
			field.RelationField
			User struct {
				field.RelationField
			}
			Classroom struct {
				field.RelationField
			}
		}
		AssignmentRepositories struct {
			field.RelationField
			Assignment struct {
				field.RelationField
				Classroom struct {
					field.RelationField
				}
				Projects struct {
					field.RelationField
				}
				Invitations struct {
					field.RelationField
					Assignment struct {
						field.RelationField
					}
					User struct {
						field.RelationField
					}
				}
			}
			User struct {
				field.RelationField
			}
		}
	}
	Member struct {
		field.RelationField
	}
	Assignments struct {
		field.RelationField
	}
	Invitations struct {
		field.RelationField
		Classroom struct {
			field.RelationField
		}
	}
}

func (a assignmentInvitationBelongsToAssignment) Where(conds ...field.Expr) *assignmentInvitationBelongsToAssignment {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a assignmentInvitationBelongsToAssignment) WithContext(ctx context.Context) *assignmentInvitationBelongsToAssignment {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a assignmentInvitationBelongsToAssignment) Session(session *gorm.Session) *assignmentInvitationBelongsToAssignment {
	a.db = a.db.Session(session)
	return &a
}

func (a assignmentInvitationBelongsToAssignment) Model(m *database.AssignmentInvitation) *assignmentInvitationBelongsToAssignmentTx {
	return &assignmentInvitationBelongsToAssignmentTx{a.db.Model(m).Association(a.Name())}
}

type assignmentInvitationBelongsToAssignmentTx struct{ tx *gorm.Association }

func (a assignmentInvitationBelongsToAssignmentTx) Find() (result *database.Classroom, err error) {
	return result, a.tx.Find(&result)
}

func (a assignmentInvitationBelongsToAssignmentTx) Append(values ...*database.Classroom) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a assignmentInvitationBelongsToAssignmentTx) Replace(values ...*database.Classroom) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a assignmentInvitationBelongsToAssignmentTx) Delete(values ...*database.Classroom) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a assignmentInvitationBelongsToAssignmentTx) Clear() error {
	return a.tx.Clear()
}

func (a assignmentInvitationBelongsToAssignmentTx) Count() int64 {
	return a.tx.Count()
}

type assignmentInvitationBelongsToUser struct {
	db *gorm.DB

	field.RelationField
}

func (a assignmentInvitationBelongsToUser) Where(conds ...field.Expr) *assignmentInvitationBelongsToUser {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a assignmentInvitationBelongsToUser) WithContext(ctx context.Context) *assignmentInvitationBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a assignmentInvitationBelongsToUser) Session(session *gorm.Session) *assignmentInvitationBelongsToUser {
	a.db = a.db.Session(session)
	return &a
}

func (a assignmentInvitationBelongsToUser) Model(m *database.AssignmentInvitation) *assignmentInvitationBelongsToUserTx {
	return &assignmentInvitationBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

type assignmentInvitationBelongsToUserTx struct{ tx *gorm.Association }

func (a assignmentInvitationBelongsToUserTx) Find() (result *database.User, err error) {
	return result, a.tx.Find(&result)
}

func (a assignmentInvitationBelongsToUserTx) Append(values ...*database.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a assignmentInvitationBelongsToUserTx) Replace(values ...*database.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a assignmentInvitationBelongsToUserTx) Delete(values ...*database.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a assignmentInvitationBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a assignmentInvitationBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

type assignmentInvitationDo struct{ gen.DO }

type IAssignmentInvitationDo interface {
	gen.SubQuery
	Debug() IAssignmentInvitationDo
	WithContext(ctx context.Context) IAssignmentInvitationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAssignmentInvitationDo
	WriteDB() IAssignmentInvitationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAssignmentInvitationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAssignmentInvitationDo
	Not(conds ...gen.Condition) IAssignmentInvitationDo
	Or(conds ...gen.Condition) IAssignmentInvitationDo
	Select(conds ...field.Expr) IAssignmentInvitationDo
	Where(conds ...gen.Condition) IAssignmentInvitationDo
	Order(conds ...field.Expr) IAssignmentInvitationDo
	Distinct(cols ...field.Expr) IAssignmentInvitationDo
	Omit(cols ...field.Expr) IAssignmentInvitationDo
	Join(table schema.Tabler, on ...field.Expr) IAssignmentInvitationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAssignmentInvitationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAssignmentInvitationDo
	Group(cols ...field.Expr) IAssignmentInvitationDo
	Having(conds ...gen.Condition) IAssignmentInvitationDo
	Limit(limit int) IAssignmentInvitationDo
	Offset(offset int) IAssignmentInvitationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAssignmentInvitationDo
	Unscoped() IAssignmentInvitationDo
	Create(values ...*database.AssignmentInvitation) error
	CreateInBatches(values []*database.AssignmentInvitation, batchSize int) error
	Save(values ...*database.AssignmentInvitation) error
	First() (*database.AssignmentInvitation, error)
	Take() (*database.AssignmentInvitation, error)
	Last() (*database.AssignmentInvitation, error)
	Find() ([]*database.AssignmentInvitation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*database.AssignmentInvitation, err error)
	FindInBatches(result *[]*database.AssignmentInvitation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*database.AssignmentInvitation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAssignmentInvitationDo
	Assign(attrs ...field.AssignExpr) IAssignmentInvitationDo
	Joins(fields ...field.RelationField) IAssignmentInvitationDo
	Preload(fields ...field.RelationField) IAssignmentInvitationDo
	FirstOrInit() (*database.AssignmentInvitation, error)
	FirstOrCreate() (*database.AssignmentInvitation, error)
	FindByPage(offset int, limit int) (result []*database.AssignmentInvitation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAssignmentInvitationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a assignmentInvitationDo) Debug() IAssignmentInvitationDo {
	return a.withDO(a.DO.Debug())
}

func (a assignmentInvitationDo) WithContext(ctx context.Context) IAssignmentInvitationDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a assignmentInvitationDo) ReadDB() IAssignmentInvitationDo {
	return a.Clauses(dbresolver.Read)
}

func (a assignmentInvitationDo) WriteDB() IAssignmentInvitationDo {
	return a.Clauses(dbresolver.Write)
}

func (a assignmentInvitationDo) Session(config *gorm.Session) IAssignmentInvitationDo {
	return a.withDO(a.DO.Session(config))
}

func (a assignmentInvitationDo) Clauses(conds ...clause.Expression) IAssignmentInvitationDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a assignmentInvitationDo) Returning(value interface{}, columns ...string) IAssignmentInvitationDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a assignmentInvitationDo) Not(conds ...gen.Condition) IAssignmentInvitationDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a assignmentInvitationDo) Or(conds ...gen.Condition) IAssignmentInvitationDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a assignmentInvitationDo) Select(conds ...field.Expr) IAssignmentInvitationDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a assignmentInvitationDo) Where(conds ...gen.Condition) IAssignmentInvitationDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a assignmentInvitationDo) Order(conds ...field.Expr) IAssignmentInvitationDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a assignmentInvitationDo) Distinct(cols ...field.Expr) IAssignmentInvitationDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a assignmentInvitationDo) Omit(cols ...field.Expr) IAssignmentInvitationDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a assignmentInvitationDo) Join(table schema.Tabler, on ...field.Expr) IAssignmentInvitationDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a assignmentInvitationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAssignmentInvitationDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a assignmentInvitationDo) RightJoin(table schema.Tabler, on ...field.Expr) IAssignmentInvitationDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a assignmentInvitationDo) Group(cols ...field.Expr) IAssignmentInvitationDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a assignmentInvitationDo) Having(conds ...gen.Condition) IAssignmentInvitationDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a assignmentInvitationDo) Limit(limit int) IAssignmentInvitationDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a assignmentInvitationDo) Offset(offset int) IAssignmentInvitationDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a assignmentInvitationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAssignmentInvitationDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a assignmentInvitationDo) Unscoped() IAssignmentInvitationDo {
	return a.withDO(a.DO.Unscoped())
}

func (a assignmentInvitationDo) Create(values ...*database.AssignmentInvitation) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a assignmentInvitationDo) CreateInBatches(values []*database.AssignmentInvitation, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a assignmentInvitationDo) Save(values ...*database.AssignmentInvitation) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a assignmentInvitationDo) First() (*database.AssignmentInvitation, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*database.AssignmentInvitation), nil
	}
}

func (a assignmentInvitationDo) Take() (*database.AssignmentInvitation, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*database.AssignmentInvitation), nil
	}
}

func (a assignmentInvitationDo) Last() (*database.AssignmentInvitation, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*database.AssignmentInvitation), nil
	}
}

func (a assignmentInvitationDo) Find() ([]*database.AssignmentInvitation, error) {
	result, err := a.DO.Find()
	return result.([]*database.AssignmentInvitation), err
}

func (a assignmentInvitationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*database.AssignmentInvitation, err error) {
	buf := make([]*database.AssignmentInvitation, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a assignmentInvitationDo) FindInBatches(result *[]*database.AssignmentInvitation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a assignmentInvitationDo) Attrs(attrs ...field.AssignExpr) IAssignmentInvitationDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a assignmentInvitationDo) Assign(attrs ...field.AssignExpr) IAssignmentInvitationDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a assignmentInvitationDo) Joins(fields ...field.RelationField) IAssignmentInvitationDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a assignmentInvitationDo) Preload(fields ...field.RelationField) IAssignmentInvitationDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a assignmentInvitationDo) FirstOrInit() (*database.AssignmentInvitation, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*database.AssignmentInvitation), nil
	}
}

func (a assignmentInvitationDo) FirstOrCreate() (*database.AssignmentInvitation, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*database.AssignmentInvitation), nil
	}
}

func (a assignmentInvitationDo) FindByPage(offset int, limit int) (result []*database.AssignmentInvitation, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a assignmentInvitationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a assignmentInvitationDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a assignmentInvitationDo) Delete(models ...*database.AssignmentInvitation) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *assignmentInvitationDo) withDO(do gen.Dao) *assignmentInvitationDo {
	a.DO = *do.(*gen.DO)
	return a
}
