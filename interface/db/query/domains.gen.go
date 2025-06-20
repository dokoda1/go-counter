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

func newDomain(db *gorm.DB, opts ...gen.DOOption) domain {
	_domain := domain{}

	_domain.domainDo.UseDB(db, opts...)
	_domain.domainDo.UseModel(&model.Domain{})

	tableName := _domain.domainDo.TableName()
	_domain.ALL = field.NewAsterisk(tableName)
	_domain.ID = field.NewInt32(tableName, "id")
	_domain.Name = field.NewString(tableName, "name")
	_domain.URL = field.NewString(tableName, "url")
	_domain.CreatedAt = field.NewTime(tableName, "created_at")
	_domain.UpdatedAt = field.NewTime(tableName, "updated_at")

	_domain.fillFieldMap()

	return _domain
}

type domain struct {
	domainDo

	ALL       field.Asterisk
	ID        field.Int32
	Name      field.String
	URL       field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (d domain) Table(newTableName string) *domain {
	d.domainDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d domain) As(alias string) *domain {
	d.domainDo.DO = *(d.domainDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *domain) updateTableName(table string) *domain {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewInt32(table, "id")
	d.Name = field.NewString(table, "name")
	d.URL = field.NewString(table, "url")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")

	d.fillFieldMap()

	return d
}

func (d *domain) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *domain) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 5)
	d.fieldMap["id"] = d.ID
	d.fieldMap["name"] = d.Name
	d.fieldMap["url"] = d.URL
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
}

func (d domain) clone(db *gorm.DB) domain {
	d.domainDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d domain) replaceDB(db *gorm.DB) domain {
	d.domainDo.ReplaceDB(db)
	return d
}

type domainDo struct{ gen.DO }

type IDomainDo interface {
	gen.SubQuery
	Debug() IDomainDo
	WithContext(ctx context.Context) IDomainDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDomainDo
	WriteDB() IDomainDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDomainDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDomainDo
	Not(conds ...gen.Condition) IDomainDo
	Or(conds ...gen.Condition) IDomainDo
	Select(conds ...field.Expr) IDomainDo
	Where(conds ...gen.Condition) IDomainDo
	Order(conds ...field.Expr) IDomainDo
	Distinct(cols ...field.Expr) IDomainDo
	Omit(cols ...field.Expr) IDomainDo
	Join(table schema.Tabler, on ...field.Expr) IDomainDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDomainDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDomainDo
	Group(cols ...field.Expr) IDomainDo
	Having(conds ...gen.Condition) IDomainDo
	Limit(limit int) IDomainDo
	Offset(offset int) IDomainDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDomainDo
	Unscoped() IDomainDo
	Create(values ...*model.Domain) error
	CreateInBatches(values []*model.Domain, batchSize int) error
	Save(values ...*model.Domain) error
	First() (*model.Domain, error)
	Take() (*model.Domain, error)
	Last() (*model.Domain, error)
	Find() ([]*model.Domain, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Domain, err error)
	FindInBatches(result *[]*model.Domain, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Domain) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDomainDo
	Assign(attrs ...field.AssignExpr) IDomainDo
	Joins(fields ...field.RelationField) IDomainDo
	Preload(fields ...field.RelationField) IDomainDo
	FirstOrInit() (*model.Domain, error)
	FirstOrCreate() (*model.Domain, error)
	FindByPage(offset int, limit int) (result []*model.Domain, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDomainDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d domainDo) Debug() IDomainDo {
	return d.withDO(d.DO.Debug())
}

func (d domainDo) WithContext(ctx context.Context) IDomainDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d domainDo) ReadDB() IDomainDo {
	return d.Clauses(dbresolver.Read)
}

func (d domainDo) WriteDB() IDomainDo {
	return d.Clauses(dbresolver.Write)
}

func (d domainDo) Session(config *gorm.Session) IDomainDo {
	return d.withDO(d.DO.Session(config))
}

func (d domainDo) Clauses(conds ...clause.Expression) IDomainDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d domainDo) Returning(value interface{}, columns ...string) IDomainDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d domainDo) Not(conds ...gen.Condition) IDomainDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d domainDo) Or(conds ...gen.Condition) IDomainDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d domainDo) Select(conds ...field.Expr) IDomainDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d domainDo) Where(conds ...gen.Condition) IDomainDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d domainDo) Order(conds ...field.Expr) IDomainDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d domainDo) Distinct(cols ...field.Expr) IDomainDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d domainDo) Omit(cols ...field.Expr) IDomainDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d domainDo) Join(table schema.Tabler, on ...field.Expr) IDomainDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d domainDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDomainDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d domainDo) RightJoin(table schema.Tabler, on ...field.Expr) IDomainDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d domainDo) Group(cols ...field.Expr) IDomainDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d domainDo) Having(conds ...gen.Condition) IDomainDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d domainDo) Limit(limit int) IDomainDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d domainDo) Offset(offset int) IDomainDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d domainDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDomainDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d domainDo) Unscoped() IDomainDo {
	return d.withDO(d.DO.Unscoped())
}

func (d domainDo) Create(values ...*model.Domain) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d domainDo) CreateInBatches(values []*model.Domain, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d domainDo) Save(values ...*model.Domain) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d domainDo) First() (*model.Domain, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Domain), nil
	}
}

func (d domainDo) Take() (*model.Domain, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Domain), nil
	}
}

func (d domainDo) Last() (*model.Domain, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Domain), nil
	}
}

func (d domainDo) Find() ([]*model.Domain, error) {
	result, err := d.DO.Find()
	return result.([]*model.Domain), err
}

func (d domainDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Domain, err error) {
	buf := make([]*model.Domain, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d domainDo) FindInBatches(result *[]*model.Domain, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d domainDo) Attrs(attrs ...field.AssignExpr) IDomainDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d domainDo) Assign(attrs ...field.AssignExpr) IDomainDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d domainDo) Joins(fields ...field.RelationField) IDomainDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d domainDo) Preload(fields ...field.RelationField) IDomainDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d domainDo) FirstOrInit() (*model.Domain, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Domain), nil
	}
}

func (d domainDo) FirstOrCreate() (*model.Domain, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Domain), nil
	}
}

func (d domainDo) FindByPage(offset int, limit int) (result []*model.Domain, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d domainDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d domainDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d domainDo) Delete(models ...*model.Domain) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *domainDo) withDO(do gen.Dao) *domainDo {
	d.DO = *do.(*gen.DO)
	return d
}
