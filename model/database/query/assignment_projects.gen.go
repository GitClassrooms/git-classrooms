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

func newAssignmentProjects(db *gorm.DB, opts ...gen.DOOption) assignmentProjects {
	_assignmentProjects := assignmentProjects{}

	_assignmentProjects.assignmentProjectsDo.UseDB(db, opts...)
	_assignmentProjects.assignmentProjectsDo.UseModel(&database.AssignmentProjects{})

	tableName := _assignmentProjects.assignmentProjectsDo.TableName()
	_assignmentProjects.ALL = field.NewAsterisk(tableName)
	_assignmentProjects.ID = field.NewField(tableName, "id")
	_assignmentProjects.CreatedAt = field.NewTime(tableName, "created_at")
	_assignmentProjects.UpdatedAt = field.NewTime(tableName, "updated_at")
	_assignmentProjects.DeletedAt = field.NewField(tableName, "deleted_at")
	_assignmentProjects.AssignmentID = field.NewField(tableName, "assignment_id")
	_assignmentProjects.UserID = field.NewField(tableName, "user_id")
	_assignmentProjects.ProjectID = field.NewInt(tableName, "project_id")
	_assignmentProjects.Assignment = assignmentProjectsBelongsToAssignment{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Assignment", "database.Assignment"),
		Classroom: struct {
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
		}{
			RelationField: field.NewRelation("Assignment.Classroom", "database.Classroom"),
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
					}
					User struct {
						field.RelationField
					}
				}
			}{
				RelationField: field.NewRelation("Assignment.Classroom.Owner", "database.User"),
				Classrooms: struct {
					field.RelationField
					User struct {
						field.RelationField
					}
					Classroom struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("Assignment.Classroom.Owner.Classrooms", "database.UserClassrooms"),
					User: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Assignment.Classroom.Owner.Classrooms.User", "database.User"),
					},
					Classroom: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Assignment.Classroom.Owner.Classrooms.Classroom", "database.Classroom"),
					},
				},
				AssignmentRepositories: struct {
					field.RelationField
					Assignment struct {
						field.RelationField
					}
					User struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("Assignment.Classroom.Owner.AssignmentRepositories", "database.AssignmentProjects"),
					Assignment: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Assignment.Classroom.Owner.AssignmentRepositories.Assignment", "database.Assignment"),
					},
					User: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Assignment.Classroom.Owner.AssignmentRepositories.User", "database.User"),
					},
				},
			},
			Member: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Assignment.Classroom.Member", "database.UserClassrooms"),
			},
			Assignments: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Assignment.Classroom.Assignments", "database.Assignment"),
			},
			Invitations: struct {
				field.RelationField
				Classroom struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Assignment.Classroom.Invitations", "database.ClassroomInvitation"),
				Classroom: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Assignment.Classroom.Invitations.Classroom", "database.Classroom"),
				},
			},
		},
		Projects: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Assignment.Projects", "database.AssignmentProjects"),
		},
		Invitations: struct {
			field.RelationField
			Assignment struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Assignment.Invitations", "database.AssignmentInvitation"),
			Assignment: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Assignment.Invitations.Assignment", "database.Classroom"),
			},
		},
	}

	_assignmentProjects.User = assignmentProjectsBelongsToUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "database.User"),
	}

	_assignmentProjects.fillFieldMap()

	return _assignmentProjects
}

type assignmentProjects struct {
	assignmentProjectsDo

	ALL          field.Asterisk
	ID           field.Field
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	AssignmentID field.Field
	UserID       field.Field
	ProjectID    field.Int
	Assignment   assignmentProjectsBelongsToAssignment

	User assignmentProjectsBelongsToUser

	fieldMap map[string]field.Expr
}

func (a assignmentProjects) Table(newTableName string) *assignmentProjects {
	a.assignmentProjectsDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a assignmentProjects) As(alias string) *assignmentProjects {
	a.assignmentProjectsDo.DO = *(a.assignmentProjectsDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *assignmentProjects) updateTableName(table string) *assignmentProjects {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewField(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.AssignmentID = field.NewField(table, "assignment_id")
	a.UserID = field.NewField(table, "user_id")
	a.ProjectID = field.NewInt(table, "project_id")

	a.fillFieldMap()

	return a
}

func (a *assignmentProjects) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *assignmentProjects) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["assignment_id"] = a.AssignmentID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["project_id"] = a.ProjectID

}

func (a assignmentProjects) clone(db *gorm.DB) assignmentProjects {
	a.assignmentProjectsDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a assignmentProjects) replaceDB(db *gorm.DB) assignmentProjects {
	a.assignmentProjectsDo.ReplaceDB(db)
	return a
}

type assignmentProjectsBelongsToAssignment struct {
	db *gorm.DB

	field.RelationField

	Classroom struct {
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
	Projects struct {
		field.RelationField
	}
	Invitations struct {
		field.RelationField
		Assignment struct {
			field.RelationField
		}
	}
}

func (a assignmentProjectsBelongsToAssignment) Where(conds ...field.Expr) *assignmentProjectsBelongsToAssignment {
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

func (a assignmentProjectsBelongsToAssignment) WithContext(ctx context.Context) *assignmentProjectsBelongsToAssignment {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a assignmentProjectsBelongsToAssignment) Session(session *gorm.Session) *assignmentProjectsBelongsToAssignment {
	a.db = a.db.Session(session)
	return &a
}

func (a assignmentProjectsBelongsToAssignment) Model(m *database.AssignmentProjects) *assignmentProjectsBelongsToAssignmentTx {
	return &assignmentProjectsBelongsToAssignmentTx{a.db.Model(m).Association(a.Name())}
}

type assignmentProjectsBelongsToAssignmentTx struct{ tx *gorm.Association }

func (a assignmentProjectsBelongsToAssignmentTx) Find() (result *database.Assignment, err error) {
	return result, a.tx.Find(&result)
}

func (a assignmentProjectsBelongsToAssignmentTx) Append(values ...*database.Assignment) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a assignmentProjectsBelongsToAssignmentTx) Replace(values ...*database.Assignment) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a assignmentProjectsBelongsToAssignmentTx) Delete(values ...*database.Assignment) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a assignmentProjectsBelongsToAssignmentTx) Clear() error {
	return a.tx.Clear()
}

func (a assignmentProjectsBelongsToAssignmentTx) Count() int64 {
	return a.tx.Count()
}

type assignmentProjectsBelongsToUser struct {
	db *gorm.DB

	field.RelationField
}

func (a assignmentProjectsBelongsToUser) Where(conds ...field.Expr) *assignmentProjectsBelongsToUser {
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

func (a assignmentProjectsBelongsToUser) WithContext(ctx context.Context) *assignmentProjectsBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a assignmentProjectsBelongsToUser) Session(session *gorm.Session) *assignmentProjectsBelongsToUser {
	a.db = a.db.Session(session)
	return &a
}

func (a assignmentProjectsBelongsToUser) Model(m *database.AssignmentProjects) *assignmentProjectsBelongsToUserTx {
	return &assignmentProjectsBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

type assignmentProjectsBelongsToUserTx struct{ tx *gorm.Association }

func (a assignmentProjectsBelongsToUserTx) Find() (result *database.User, err error) {
	return result, a.tx.Find(&result)
}

func (a assignmentProjectsBelongsToUserTx) Append(values ...*database.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a assignmentProjectsBelongsToUserTx) Replace(values ...*database.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a assignmentProjectsBelongsToUserTx) Delete(values ...*database.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a assignmentProjectsBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a assignmentProjectsBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

type assignmentProjectsDo struct{ gen.DO }

type IAssignmentProjectsDo interface {
	gen.SubQuery
	Debug() IAssignmentProjectsDo
	WithContext(ctx context.Context) IAssignmentProjectsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAssignmentProjectsDo
	WriteDB() IAssignmentProjectsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAssignmentProjectsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAssignmentProjectsDo
	Not(conds ...gen.Condition) IAssignmentProjectsDo
	Or(conds ...gen.Condition) IAssignmentProjectsDo
	Select(conds ...field.Expr) IAssignmentProjectsDo
	Where(conds ...gen.Condition) IAssignmentProjectsDo
	Order(conds ...field.Expr) IAssignmentProjectsDo
	Distinct(cols ...field.Expr) IAssignmentProjectsDo
	Omit(cols ...field.Expr) IAssignmentProjectsDo
	Join(table schema.Tabler, on ...field.Expr) IAssignmentProjectsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAssignmentProjectsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAssignmentProjectsDo
	Group(cols ...field.Expr) IAssignmentProjectsDo
	Having(conds ...gen.Condition) IAssignmentProjectsDo
	Limit(limit int) IAssignmentProjectsDo
	Offset(offset int) IAssignmentProjectsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAssignmentProjectsDo
	Unscoped() IAssignmentProjectsDo
	Create(values ...*database.AssignmentProjects) error
	CreateInBatches(values []*database.AssignmentProjects, batchSize int) error
	Save(values ...*database.AssignmentProjects) error
	First() (*database.AssignmentProjects, error)
	Take() (*database.AssignmentProjects, error)
	Last() (*database.AssignmentProjects, error)
	Find() ([]*database.AssignmentProjects, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*database.AssignmentProjects, err error)
	FindInBatches(result *[]*database.AssignmentProjects, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*database.AssignmentProjects) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAssignmentProjectsDo
	Assign(attrs ...field.AssignExpr) IAssignmentProjectsDo
	Joins(fields ...field.RelationField) IAssignmentProjectsDo
	Preload(fields ...field.RelationField) IAssignmentProjectsDo
	FirstOrInit() (*database.AssignmentProjects, error)
	FirstOrCreate() (*database.AssignmentProjects, error)
	FindByPage(offset int, limit int) (result []*database.AssignmentProjects, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAssignmentProjectsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a assignmentProjectsDo) Debug() IAssignmentProjectsDo {
	return a.withDO(a.DO.Debug())
}

func (a assignmentProjectsDo) WithContext(ctx context.Context) IAssignmentProjectsDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a assignmentProjectsDo) ReadDB() IAssignmentProjectsDo {
	return a.Clauses(dbresolver.Read)
}

func (a assignmentProjectsDo) WriteDB() IAssignmentProjectsDo {
	return a.Clauses(dbresolver.Write)
}

func (a assignmentProjectsDo) Session(config *gorm.Session) IAssignmentProjectsDo {
	return a.withDO(a.DO.Session(config))
}

func (a assignmentProjectsDo) Clauses(conds ...clause.Expression) IAssignmentProjectsDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a assignmentProjectsDo) Returning(value interface{}, columns ...string) IAssignmentProjectsDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a assignmentProjectsDo) Not(conds ...gen.Condition) IAssignmentProjectsDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a assignmentProjectsDo) Or(conds ...gen.Condition) IAssignmentProjectsDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a assignmentProjectsDo) Select(conds ...field.Expr) IAssignmentProjectsDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a assignmentProjectsDo) Where(conds ...gen.Condition) IAssignmentProjectsDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a assignmentProjectsDo) Order(conds ...field.Expr) IAssignmentProjectsDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a assignmentProjectsDo) Distinct(cols ...field.Expr) IAssignmentProjectsDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a assignmentProjectsDo) Omit(cols ...field.Expr) IAssignmentProjectsDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a assignmentProjectsDo) Join(table schema.Tabler, on ...field.Expr) IAssignmentProjectsDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a assignmentProjectsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAssignmentProjectsDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a assignmentProjectsDo) RightJoin(table schema.Tabler, on ...field.Expr) IAssignmentProjectsDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a assignmentProjectsDo) Group(cols ...field.Expr) IAssignmentProjectsDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a assignmentProjectsDo) Having(conds ...gen.Condition) IAssignmentProjectsDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a assignmentProjectsDo) Limit(limit int) IAssignmentProjectsDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a assignmentProjectsDo) Offset(offset int) IAssignmentProjectsDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a assignmentProjectsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAssignmentProjectsDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a assignmentProjectsDo) Unscoped() IAssignmentProjectsDo {
	return a.withDO(a.DO.Unscoped())
}

func (a assignmentProjectsDo) Create(values ...*database.AssignmentProjects) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a assignmentProjectsDo) CreateInBatches(values []*database.AssignmentProjects, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a assignmentProjectsDo) Save(values ...*database.AssignmentProjects) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a assignmentProjectsDo) First() (*database.AssignmentProjects, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*database.AssignmentProjects), nil
	}
}

func (a assignmentProjectsDo) Take() (*database.AssignmentProjects, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*database.AssignmentProjects), nil
	}
}

func (a assignmentProjectsDo) Last() (*database.AssignmentProjects, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*database.AssignmentProjects), nil
	}
}

func (a assignmentProjectsDo) Find() ([]*database.AssignmentProjects, error) {
	result, err := a.DO.Find()
	return result.([]*database.AssignmentProjects), err
}

func (a assignmentProjectsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*database.AssignmentProjects, err error) {
	buf := make([]*database.AssignmentProjects, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a assignmentProjectsDo) FindInBatches(result *[]*database.AssignmentProjects, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a assignmentProjectsDo) Attrs(attrs ...field.AssignExpr) IAssignmentProjectsDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a assignmentProjectsDo) Assign(attrs ...field.AssignExpr) IAssignmentProjectsDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a assignmentProjectsDo) Joins(fields ...field.RelationField) IAssignmentProjectsDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a assignmentProjectsDo) Preload(fields ...field.RelationField) IAssignmentProjectsDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a assignmentProjectsDo) FirstOrInit() (*database.AssignmentProjects, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*database.AssignmentProjects), nil
	}
}

func (a assignmentProjectsDo) FirstOrCreate() (*database.AssignmentProjects, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*database.AssignmentProjects), nil
	}
}

func (a assignmentProjectsDo) FindByPage(offset int, limit int) (result []*database.AssignmentProjects, count int64, err error) {
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

func (a assignmentProjectsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a assignmentProjectsDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a assignmentProjectsDo) Delete(models ...*database.AssignmentProjects) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *assignmentProjectsDo) withDO(do gen.Dao) *assignmentProjectsDo {
	a.DO = *do.(*gen.DO)
	return a
}
