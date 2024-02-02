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

func newUser(db *gorm.DB, opts ...gen.DOOption) user {
	_user := user{}

	_user.userDo.UseDB(db, opts...)
	_user.userDo.UseModel(&database.User{})

	tableName := _user.userDo.TableName()
	_user.ALL = field.NewAsterisk(tableName)
	_user.ID = field.NewInt(tableName, "id")
	_user.CreatedAt = field.NewTime(tableName, "created_at")
	_user.UpdatedAt = field.NewTime(tableName, "updated_at")
	_user.DeletedAt = field.NewField(tableName, "deleted_at")
	_user.Classrooms = userHasManyClassrooms{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Classrooms", "database.UserClassrooms"),
		User: struct {
			field.RelationField
			Classrooms struct {
				field.RelationField
			}
			AssignmentRepositories struct {
				field.RelationField
				Assignment struct {
					field.RelationField
					Classroom struct {
						field.RelationField
						Owner struct {
							field.RelationField
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
				User struct {
					field.RelationField
				}
			}
		}{
			RelationField: field.NewRelation("Classrooms.User", "database.User"),
			Classrooms: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Classrooms.User.Classrooms", "database.UserClassrooms"),
			},
			AssignmentRepositories: struct {
				field.RelationField
				Assignment struct {
					field.RelationField
					Classroom struct {
						field.RelationField
						Owner struct {
							field.RelationField
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
				User struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Classrooms.User.AssignmentRepositories", "database.AssignmentProjects"),
				Assignment: struct {
					field.RelationField
					Classroom struct {
						field.RelationField
						Owner struct {
							field.RelationField
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
				}{
					RelationField: field.NewRelation("Classrooms.User.AssignmentRepositories.Assignment", "database.Assignment"),
					Classroom: struct {
						field.RelationField
						Owner struct {
							field.RelationField
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
						RelationField: field.NewRelation("Classrooms.User.AssignmentRepositories.Assignment.Classroom", "database.Classroom"),
						Owner: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("Classrooms.User.AssignmentRepositories.Assignment.Classroom.Owner", "database.User"),
						},
						Member: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("Classrooms.User.AssignmentRepositories.Assignment.Classroom.Member", "database.UserClassrooms"),
						},
						Assignments: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("Classrooms.User.AssignmentRepositories.Assignment.Classroom.Assignments", "database.Assignment"),
						},
						Invitations: struct {
							field.RelationField
							Classroom struct {
								field.RelationField
							}
						}{
							RelationField: field.NewRelation("Classrooms.User.AssignmentRepositories.Assignment.Classroom.Invitations", "database.ClassroomInvitation"),
							Classroom: struct {
								field.RelationField
							}{
								RelationField: field.NewRelation("Classrooms.User.AssignmentRepositories.Assignment.Classroom.Invitations.Classroom", "database.Classroom"),
							},
						},
					},
					Projects: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Classrooms.User.AssignmentRepositories.Assignment.Projects", "database.AssignmentProjects"),
					},
					Invitations: struct {
						field.RelationField
						Assignment struct {
							field.RelationField
						}
					}{
						RelationField: field.NewRelation("Classrooms.User.AssignmentRepositories.Assignment.Invitations", "database.AssignmentInvitation"),
						Assignment: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("Classrooms.User.AssignmentRepositories.Assignment.Invitations.Assignment", "database.Classroom"),
						},
					},
				},
				User: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Classrooms.User.AssignmentRepositories.User", "database.User"),
				},
			},
		},
		Classroom: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Classrooms.Classroom", "database.Classroom"),
		},
	}

	_user.AssignmentRepositories = userHasManyAssignmentRepositories{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("AssignmentRepositories", "database.AssignmentProjects"),
	}

	_user.fillFieldMap()

	return _user
}

type user struct {
	userDo

	ALL        field.Asterisk
	ID         field.Int
	CreatedAt  field.Time
	UpdatedAt  field.Time
	DeletedAt  field.Field
	Classrooms userHasManyClassrooms

	AssignmentRepositories userHasManyAssignmentRepositories

	fieldMap map[string]field.Expr
}

func (u user) Table(newTableName string) *user {
	u.userDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u user) As(alias string) *user {
	u.userDo.DO = *(u.userDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *user) updateTableName(table string) *user {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt(table, "id")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *user) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *user) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 6)
	u.fieldMap["id"] = u.ID
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt

}

func (u user) clone(db *gorm.DB) user {
	u.userDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u user) replaceDB(db *gorm.DB) user {
	u.userDo.ReplaceDB(db)
	return u
}

type userHasManyClassrooms struct {
	db *gorm.DB

	field.RelationField

	User struct {
		field.RelationField
		Classrooms struct {
			field.RelationField
		}
		AssignmentRepositories struct {
			field.RelationField
			Assignment struct {
				field.RelationField
				Classroom struct {
					field.RelationField
					Owner struct {
						field.RelationField
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
			User struct {
				field.RelationField
			}
		}
	}
	Classroom struct {
		field.RelationField
	}
}

func (a userHasManyClassrooms) Where(conds ...field.Expr) *userHasManyClassrooms {
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

func (a userHasManyClassrooms) WithContext(ctx context.Context) *userHasManyClassrooms {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userHasManyClassrooms) Session(session *gorm.Session) *userHasManyClassrooms {
	a.db = a.db.Session(session)
	return &a
}

func (a userHasManyClassrooms) Model(m *database.User) *userHasManyClassroomsTx {
	return &userHasManyClassroomsTx{a.db.Model(m).Association(a.Name())}
}

type userHasManyClassroomsTx struct{ tx *gorm.Association }

func (a userHasManyClassroomsTx) Find() (result []*database.UserClassrooms, err error) {
	return result, a.tx.Find(&result)
}

func (a userHasManyClassroomsTx) Append(values ...*database.UserClassrooms) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userHasManyClassroomsTx) Replace(values ...*database.UserClassrooms) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userHasManyClassroomsTx) Delete(values ...*database.UserClassrooms) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userHasManyClassroomsTx) Clear() error {
	return a.tx.Clear()
}

func (a userHasManyClassroomsTx) Count() int64 {
	return a.tx.Count()
}

type userHasManyAssignmentRepositories struct {
	db *gorm.DB

	field.RelationField
}

func (a userHasManyAssignmentRepositories) Where(conds ...field.Expr) *userHasManyAssignmentRepositories {
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

func (a userHasManyAssignmentRepositories) WithContext(ctx context.Context) *userHasManyAssignmentRepositories {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userHasManyAssignmentRepositories) Session(session *gorm.Session) *userHasManyAssignmentRepositories {
	a.db = a.db.Session(session)
	return &a
}

func (a userHasManyAssignmentRepositories) Model(m *database.User) *userHasManyAssignmentRepositoriesTx {
	return &userHasManyAssignmentRepositoriesTx{a.db.Model(m).Association(a.Name())}
}

type userHasManyAssignmentRepositoriesTx struct{ tx *gorm.Association }

func (a userHasManyAssignmentRepositoriesTx) Find() (result []*database.AssignmentProjects, err error) {
	return result, a.tx.Find(&result)
}

func (a userHasManyAssignmentRepositoriesTx) Append(values ...*database.AssignmentProjects) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userHasManyAssignmentRepositoriesTx) Replace(values ...*database.AssignmentProjects) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userHasManyAssignmentRepositoriesTx) Delete(values ...*database.AssignmentProjects) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userHasManyAssignmentRepositoriesTx) Clear() error {
	return a.tx.Clear()
}

func (a userHasManyAssignmentRepositoriesTx) Count() int64 {
	return a.tx.Count()
}

type userDo struct{ gen.DO }

type IUserDo interface {
	gen.SubQuery
	Debug() IUserDo
	WithContext(ctx context.Context) IUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserDo
	WriteDB() IUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserDo
	Not(conds ...gen.Condition) IUserDo
	Or(conds ...gen.Condition) IUserDo
	Select(conds ...field.Expr) IUserDo
	Where(conds ...gen.Condition) IUserDo
	Order(conds ...field.Expr) IUserDo
	Distinct(cols ...field.Expr) IUserDo
	Omit(cols ...field.Expr) IUserDo
	Join(table schema.Tabler, on ...field.Expr) IUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserDo
	Group(cols ...field.Expr) IUserDo
	Having(conds ...gen.Condition) IUserDo
	Limit(limit int) IUserDo
	Offset(offset int) IUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserDo
	Unscoped() IUserDo
	Create(values ...*database.User) error
	CreateInBatches(values []*database.User, batchSize int) error
	Save(values ...*database.User) error
	First() (*database.User, error)
	Take() (*database.User, error)
	Last() (*database.User, error)
	Find() ([]*database.User, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*database.User, err error)
	FindInBatches(result *[]*database.User, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*database.User) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserDo
	Assign(attrs ...field.AssignExpr) IUserDo
	Joins(fields ...field.RelationField) IUserDo
	Preload(fields ...field.RelationField) IUserDo
	FirstOrInit() (*database.User, error)
	FirstOrCreate() (*database.User, error)
	FindByPage(offset int, limit int) (result []*database.User, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userDo) Debug() IUserDo {
	return u.withDO(u.DO.Debug())
}

func (u userDo) WithContext(ctx context.Context) IUserDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userDo) ReadDB() IUserDo {
	return u.Clauses(dbresolver.Read)
}

func (u userDo) WriteDB() IUserDo {
	return u.Clauses(dbresolver.Write)
}

func (u userDo) Session(config *gorm.Session) IUserDo {
	return u.withDO(u.DO.Session(config))
}

func (u userDo) Clauses(conds ...clause.Expression) IUserDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userDo) Returning(value interface{}, columns ...string) IUserDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userDo) Not(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userDo) Or(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userDo) Select(conds ...field.Expr) IUserDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userDo) Where(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userDo) Order(conds ...field.Expr) IUserDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userDo) Distinct(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userDo) Omit(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userDo) Join(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userDo) Group(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userDo) Having(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userDo) Limit(limit int) IUserDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userDo) Offset(offset int) IUserDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userDo) Unscoped() IUserDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userDo) Create(values ...*database.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userDo) CreateInBatches(values []*database.User, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userDo) Save(values ...*database.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userDo) First() (*database.User, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*database.User), nil
	}
}

func (u userDo) Take() (*database.User, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*database.User), nil
	}
}

func (u userDo) Last() (*database.User, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*database.User), nil
	}
}

func (u userDo) Find() ([]*database.User, error) {
	result, err := u.DO.Find()
	return result.([]*database.User), err
}

func (u userDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*database.User, err error) {
	buf := make([]*database.User, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userDo) FindInBatches(result *[]*database.User, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userDo) Attrs(attrs ...field.AssignExpr) IUserDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userDo) Assign(attrs ...field.AssignExpr) IUserDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userDo) Joins(fields ...field.RelationField) IUserDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userDo) Preload(fields ...field.RelationField) IUserDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userDo) FirstOrInit() (*database.User, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*database.User), nil
	}
}

func (u userDo) FirstOrCreate() (*database.User, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*database.User), nil
	}
}

func (u userDo) FindByPage(offset int, limit int) (result []*database.User, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userDo) Delete(models ...*database.User) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userDo) withDO(do gen.Dao) *userDo {
	u.DO = *do.(*gen.DO)
	return u
}
