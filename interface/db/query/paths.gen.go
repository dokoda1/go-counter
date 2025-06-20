// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/dokoda1/go-counter/interface/db/model"
)

func newPath(db *gorm.DB, opts ...gen.DOOption) path {
	_path := path{}

	_path.pathDo.UseDB(db, opts...)
	_path.pathDo.UseModel(&model.Path{})

	tableName := _path.pathDo.TableName()
	_path.ALL = field.NewAsterisk(tableName)
	_path.ID = field.NewInt32(tableName, "id")
	_path.Name = field.NewString(tableName, "name")
	_path.URL = field.NewString(tableName, "url")
	_path.ParentID = field.NewInt32(tableName, "parent_id")
	_path.CreatedAt = field.NewTime(tableName, "created_at")
	_path.UpdatedAt = field.NewTime(tableName, "updated_at")

	_path.fillFieldMap()

	return _path
}

type path struct {
	pathDo

	ALL       field.Asterisk
	ID        field.Int32
	Name      field.String
	URL       field.String
	ParentID  field.Int32
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (p path) Table(newTableName string) *path {
	p.pathDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p path) As(alias string) *path {
	p.pathDo.DO = *(p.pathDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *path) updateTableName(table string) *path {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.Name = field.NewString(table, "name")
	p.URL = field.NewString(table, "url")
	p.ParentID = field.NewInt32(table, "parent_id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")

	p.fillFieldMap()

	return p
}

func (p *path) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *path) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["id"] = p.ID
	p.fieldMap["name"] = p.Name
	p.fieldMap["url"] = p.URL
	p.fieldMap["parent_id"] = p.ParentID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
}

func (p path) clone(db *gorm.DB) path {
	p.pathDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p path) replaceDB(db *gorm.DB) path {
	p.pathDo.ReplaceDB(db)
	return p
}

type pathDo struct{ gen.DO }

type IPathDo interface {
	gen.SubQuery
	Debug() IPathDo
	WithContext(ctx context.Context) IPathDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPathDo
	WriteDB() IPathDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPathDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPathDo
	Not(conds ...gen.Condition) IPathDo
	Or(conds ...gen.Condition) IPathDo
	Select(conds ...field.Expr) IPathDo
	Where(conds ...gen.Condition) IPathDo
	Order(conds ...field.Expr) IPathDo
	Distinct(cols ...field.Expr) IPathDo
	Omit(cols ...field.Expr) IPathDo
	Join(table schema.Tabler, on ...field.Expr) IPathDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPathDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPathDo
	Group(cols ...field.Expr) IPathDo
	Having(conds ...gen.Condition) IPathDo
	Limit(limit int) IPathDo
	Offset(offset int) IPathDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPathDo
	Unscoped() IPathDo
	Create(values ...*model.Path) error
	CreateInBatches(values []*model.Path, batchSize int) error
	Save(values ...*model.Path) error
	First() (*model.Path, error)
	Take() (*model.Path, error)
	Last() (*model.Path, error)
	Find() ([]*model.Path, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Path, err error)
	FindInBatches(result *[]*model.Path, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Path) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPathDo
	Assign(attrs ...field.AssignExpr) IPathDo
	Joins(fields ...field.RelationField) IPathDo
	Preload(fields ...field.RelationField) IPathDo
	FirstOrInit() (*model.Path, error)
	FirstOrCreate() (*model.Path, error)
	FindByPage(offset int, limit int) (result []*model.Path, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPathDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pathDo) Debug() IPathDo {
	return p.withDO(p.DO.Debug())
}

func (p pathDo) WithContext(ctx context.Context) IPathDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pathDo) ReadDB() IPathDo {
	return p.Clauses(dbresolver.Read)
}

func (p pathDo) WriteDB() IPathDo {
	return p.Clauses(dbresolver.Write)
}

func (p pathDo) Session(config *gorm.Session) IPathDo {
	return p.withDO(p.DO.Session(config))
}

func (p pathDo) Clauses(conds ...clause.Expression) IPathDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pathDo) Returning(value interface{}, columns ...string) IPathDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pathDo) Not(conds ...gen.Condition) IPathDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pathDo) Or(conds ...gen.Condition) IPathDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pathDo) Select(conds ...field.Expr) IPathDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pathDo) Where(conds ...gen.Condition) IPathDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pathDo) Order(conds ...field.Expr) IPathDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pathDo) Distinct(cols ...field.Expr) IPathDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pathDo) Omit(cols ...field.Expr) IPathDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pathDo) Join(table schema.Tabler, on ...field.Expr) IPathDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pathDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPathDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pathDo) RightJoin(table schema.Tabler, on ...field.Expr) IPathDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pathDo) Group(cols ...field.Expr) IPathDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pathDo) Having(conds ...gen.Condition) IPathDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pathDo) Limit(limit int) IPathDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pathDo) Offset(offset int) IPathDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pathDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPathDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pathDo) Unscoped() IPathDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pathDo) Create(values ...*model.Path) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pathDo) CreateInBatches(values []*model.Path, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pathDo) Save(values ...*model.Path) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pathDo) First() (*model.Path, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Path), nil
	}
}

func (p pathDo) Take() (*model.Path, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Path), nil
	}
}

func (p pathDo) Last() (*model.Path, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Path), nil
	}
}

func (p pathDo) Find() ([]*model.Path, error) {
	result, err := p.DO.Find()
	return result.([]*model.Path), err
}

func (p pathDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Path, err error) {
	buf := make([]*model.Path, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pathDo) FindInBatches(result *[]*model.Path, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pathDo) Attrs(attrs ...field.AssignExpr) IPathDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pathDo) Assign(attrs ...field.AssignExpr) IPathDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pathDo) Joins(fields ...field.RelationField) IPathDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pathDo) Preload(fields ...field.RelationField) IPathDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pathDo) FirstOrInit() (*model.Path, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Path), nil
	}
}

func (p pathDo) FirstOrCreate() (*model.Path, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Path), nil
	}
}

func (p pathDo) FindByPage(offset int, limit int) (result []*model.Path, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pathDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pathDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pathDo) Delete(models ...*model.Path) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pathDo) withDO(do gen.Dao) *pathDo {
	p.DO = *do.(*gen.DO)
	return p
}
