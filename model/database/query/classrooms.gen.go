// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"de.hs-flensburg.gitlab/gitlab-classroom/model/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newClassroom(db *gorm.DB, opts ...gen.DOOption) classroom {
	_classroom := classroom{}

	_classroom.classroomDo.UseDB(db, opts...)
	_classroom.classroomDo.UseModel(&database.Classroom{})

	tableName := _classroom.classroomDo.TableName()
	_classroom.ALL = field.NewAsterisk(tableName)
	_classroom.ID = field.NewField(tableName, "id")
	_classroom.CreatedAt = field.NewTime(tableName, "created_at")
	_classroom.UpdatedAt = field.NewTime(tableName, "updated_at")
	_classroom.DeletedAt = field.NewField(tableName, "deleted_at")
	_classroom.Name = field.NewString(tableName, "name")
	_classroom.OwnerID = field.NewInt(tableName, "owner_id")
	_classroom.Description = field.NewString(tableName, "description")
	_classroom.GroupID = field.NewInt(tableName, "group_id")
	_classroom.Member = classroomHasManyMember{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Member", "database.UserClassrooms"),
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
					}
					Projects struct {
						field.RelationField
					}
				}
				User struct {
					field.RelationField
				}
			}
		}{
			RelationField: field.NewRelation("Member.User", "database.User"),
			Classrooms: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Member.User.Classrooms", "database.UserClassrooms"),
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
					}
					Projects struct {
						field.RelationField
					}
				}
				User struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Member.User.AssignmentRepositories", "database.AssignmentProjects"),
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
					}
					Projects struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("Member.User.AssignmentRepositories.Assignment", "database.Assignment"),
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
					}{
						RelationField: field.NewRelation("Member.User.AssignmentRepositories.Assignment.Classroom", "database.Classroom"),
						Owner: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("Member.User.AssignmentRepositories.Assignment.Classroom.Owner", "database.User"),
						},
						Member: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("Member.User.AssignmentRepositories.Assignment.Classroom.Member", "database.UserClassrooms"),
						},
						Assignments: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("Member.User.AssignmentRepositories.Assignment.Classroom.Assignments", "database.Assignment"),
						},
					},
					Projects: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Member.User.AssignmentRepositories.Assignment.Projects", "database.AssignmentProjects"),
					},
				},
				User: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Member.User.AssignmentRepositories.User", "database.User"),
				},
			},
		},
		Classroom: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Member.Classroom", "database.Classroom"),
		},
	}

	_classroom.Assignments = classroomHasManyAssignments{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Assignments", "database.Assignment"),
	}

	_classroom.Owner = classroomBelongsToOwner{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Owner", "database.User"),
	}

	_classroom.fillFieldMap()

	return _classroom
}

type classroom struct {
	classroomDo

	ALL         field.Asterisk
	ID          field.Field
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	Name        field.String
	OwnerID     field.Int
	Description field.String
	GroupID     field.Int
	Member      classroomHasManyMember

	Assignments classroomHasManyAssignments

	Owner classroomBelongsToOwner

	fieldMap map[string]field.Expr
}

func (c classroom) Table(newTableName string) *classroom {
	c.classroomDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c classroom) As(alias string) *classroom {
	c.classroomDo.DO = *(c.classroomDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *classroom) updateTableName(table string) *classroom {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewField(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.Name = field.NewString(table, "name")
	c.OwnerID = field.NewInt(table, "owner_id")
	c.Description = field.NewString(table, "description")
	c.GroupID = field.NewInt(table, "group_id")

	c.fillFieldMap()

	return c
}

func (c *classroom) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *classroom) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 11)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["name"] = c.Name
	c.fieldMap["owner_id"] = c.OwnerID
	c.fieldMap["description"] = c.Description
	c.fieldMap["group_id"] = c.GroupID

}

func (c classroom) clone(db *gorm.DB) classroom {
	c.classroomDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c classroom) replaceDB(db *gorm.DB) classroom {
	c.classroomDo.ReplaceDB(db)
	return c
}

type classroomHasManyMember struct {
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
				}
				Projects struct {
					field.RelationField
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

func (a classroomHasManyMember) Where(conds ...field.Expr) *classroomHasManyMember {
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

func (a classroomHasManyMember) WithContext(ctx context.Context) *classroomHasManyMember {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a classroomHasManyMember) Session(session *gorm.Session) *classroomHasManyMember {
	a.db = a.db.Session(session)
	return &a
}

func (a classroomHasManyMember) Model(m *database.Classroom) *classroomHasManyMemberTx {
	return &classroomHasManyMemberTx{a.db.Model(m).Association(a.Name())}
}

type classroomHasManyMemberTx struct{ tx *gorm.Association }

func (a classroomHasManyMemberTx) Find() (result []*database.UserClassrooms, err error) {
	return result, a.tx.Find(&result)
}

func (a classroomHasManyMemberTx) Append(values ...*database.UserClassrooms) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a classroomHasManyMemberTx) Replace(values ...*database.UserClassrooms) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a classroomHasManyMemberTx) Delete(values ...*database.UserClassrooms) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a classroomHasManyMemberTx) Clear() error {
	return a.tx.Clear()
}

func (a classroomHasManyMemberTx) Count() int64 {
	return a.tx.Count()
}

type classroomHasManyAssignments struct {
	db *gorm.DB

	field.RelationField
}

func (a classroomHasManyAssignments) Where(conds ...field.Expr) *classroomHasManyAssignments {
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

func (a classroomHasManyAssignments) WithContext(ctx context.Context) *classroomHasManyAssignments {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a classroomHasManyAssignments) Session(session *gorm.Session) *classroomHasManyAssignments {
	a.db = a.db.Session(session)
	return &a
}

func (a classroomHasManyAssignments) Model(m *database.Classroom) *classroomHasManyAssignmentsTx {
	return &classroomHasManyAssignmentsTx{a.db.Model(m).Association(a.Name())}
}

type classroomHasManyAssignmentsTx struct{ tx *gorm.Association }

func (a classroomHasManyAssignmentsTx) Find() (result []*database.Assignment, err error) {
	return result, a.tx.Find(&result)
}

func (a classroomHasManyAssignmentsTx) Append(values ...*database.Assignment) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a classroomHasManyAssignmentsTx) Replace(values ...*database.Assignment) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a classroomHasManyAssignmentsTx) Delete(values ...*database.Assignment) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a classroomHasManyAssignmentsTx) Clear() error {
	return a.tx.Clear()
}

func (a classroomHasManyAssignmentsTx) Count() int64 {
	return a.tx.Count()
}

type classroomBelongsToOwner struct {
	db *gorm.DB

	field.RelationField
}

func (a classroomBelongsToOwner) Where(conds ...field.Expr) *classroomBelongsToOwner {
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

func (a classroomBelongsToOwner) WithContext(ctx context.Context) *classroomBelongsToOwner {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a classroomBelongsToOwner) Session(session *gorm.Session) *classroomBelongsToOwner {
	a.db = a.db.Session(session)
	return &a
}

func (a classroomBelongsToOwner) Model(m *database.Classroom) *classroomBelongsToOwnerTx {
	return &classroomBelongsToOwnerTx{a.db.Model(m).Association(a.Name())}
}

type classroomBelongsToOwnerTx struct{ tx *gorm.Association }

func (a classroomBelongsToOwnerTx) Find() (result *database.User, err error) {
	return result, a.tx.Find(&result)
}

func (a classroomBelongsToOwnerTx) Append(values ...*database.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a classroomBelongsToOwnerTx) Replace(values ...*database.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a classroomBelongsToOwnerTx) Delete(values ...*database.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a classroomBelongsToOwnerTx) Clear() error {
	return a.tx.Clear()
}

func (a classroomBelongsToOwnerTx) Count() int64 {
	return a.tx.Count()
}

type classroomDo struct{ gen.DO }

type IClassroomDo interface {
	gen.SubQuery
	Debug() IClassroomDo
	WithContext(ctx context.Context) IClassroomDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IClassroomDo
	WriteDB() IClassroomDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IClassroomDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IClassroomDo
	Not(conds ...gen.Condition) IClassroomDo
	Or(conds ...gen.Condition) IClassroomDo
	Select(conds ...field.Expr) IClassroomDo
	Where(conds ...gen.Condition) IClassroomDo
	Order(conds ...field.Expr) IClassroomDo
	Distinct(cols ...field.Expr) IClassroomDo
	Omit(cols ...field.Expr) IClassroomDo
	Join(table schema.Tabler, on ...field.Expr) IClassroomDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IClassroomDo
	RightJoin(table schema.Tabler, on ...field.Expr) IClassroomDo
	Group(cols ...field.Expr) IClassroomDo
	Having(conds ...gen.Condition) IClassroomDo
	Limit(limit int) IClassroomDo
	Offset(offset int) IClassroomDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IClassroomDo
	Unscoped() IClassroomDo
	Create(values ...*database.Classroom) error
	CreateInBatches(values []*database.Classroom, batchSize int) error
	Save(values ...*database.Classroom) error
	First() (*database.Classroom, error)
	Take() (*database.Classroom, error)
	Last() (*database.Classroom, error)
	Find() ([]*database.Classroom, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*database.Classroom, err error)
	FindInBatches(result *[]*database.Classroom, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*database.Classroom) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IClassroomDo
	Assign(attrs ...field.AssignExpr) IClassroomDo
	Joins(fields ...field.RelationField) IClassroomDo
	Preload(fields ...field.RelationField) IClassroomDo
	FirstOrInit() (*database.Classroom, error)
	FirstOrCreate() (*database.Classroom, error)
	FindByPage(offset int, limit int) (result []*database.Classroom, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IClassroomDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c classroomDo) Debug() IClassroomDo {
	return c.withDO(c.DO.Debug())
}

func (c classroomDo) WithContext(ctx context.Context) IClassroomDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c classroomDo) ReadDB() IClassroomDo {
	return c.Clauses(dbresolver.Read)
}

func (c classroomDo) WriteDB() IClassroomDo {
	return c.Clauses(dbresolver.Write)
}

func (c classroomDo) Session(config *gorm.Session) IClassroomDo {
	return c.withDO(c.DO.Session(config))
}

func (c classroomDo) Clauses(conds ...clause.Expression) IClassroomDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c classroomDo) Returning(value interface{}, columns ...string) IClassroomDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c classroomDo) Not(conds ...gen.Condition) IClassroomDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c classroomDo) Or(conds ...gen.Condition) IClassroomDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c classroomDo) Select(conds ...field.Expr) IClassroomDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c classroomDo) Where(conds ...gen.Condition) IClassroomDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c classroomDo) Order(conds ...field.Expr) IClassroomDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c classroomDo) Distinct(cols ...field.Expr) IClassroomDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c classroomDo) Omit(cols ...field.Expr) IClassroomDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c classroomDo) Join(table schema.Tabler, on ...field.Expr) IClassroomDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c classroomDo) LeftJoin(table schema.Tabler, on ...field.Expr) IClassroomDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c classroomDo) RightJoin(table schema.Tabler, on ...field.Expr) IClassroomDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c classroomDo) Group(cols ...field.Expr) IClassroomDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c classroomDo) Having(conds ...gen.Condition) IClassroomDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c classroomDo) Limit(limit int) IClassroomDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c classroomDo) Offset(offset int) IClassroomDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c classroomDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IClassroomDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c classroomDo) Unscoped() IClassroomDo {
	return c.withDO(c.DO.Unscoped())
}

func (c classroomDo) Create(values ...*database.Classroom) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c classroomDo) CreateInBatches(values []*database.Classroom, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c classroomDo) Save(values ...*database.Classroom) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c classroomDo) First() (*database.Classroom, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*database.Classroom), nil
	}
}

func (c classroomDo) Take() (*database.Classroom, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*database.Classroom), nil
	}
}

func (c classroomDo) Last() (*database.Classroom, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*database.Classroom), nil
	}
}

func (c classroomDo) Find() ([]*database.Classroom, error) {
	result, err := c.DO.Find()
	return result.([]*database.Classroom), err
}

func (c classroomDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*database.Classroom, err error) {
	buf := make([]*database.Classroom, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c classroomDo) FindInBatches(result *[]*database.Classroom, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c classroomDo) Attrs(attrs ...field.AssignExpr) IClassroomDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c classroomDo) Assign(attrs ...field.AssignExpr) IClassroomDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c classroomDo) Joins(fields ...field.RelationField) IClassroomDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c classroomDo) Preload(fields ...field.RelationField) IClassroomDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c classroomDo) FirstOrInit() (*database.Classroom, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*database.Classroom), nil
	}
}

func (c classroomDo) FirstOrCreate() (*database.Classroom, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*database.Classroom), nil
	}
}

func (c classroomDo) FindByPage(offset int, limit int) (result []*database.Classroom, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c classroomDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c classroomDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c classroomDo) Delete(models ...*database.Classroom) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *classroomDo) withDO(do gen.Dao) *classroomDo {
	c.DO = *do.(*gen.DO)
	return c
}
