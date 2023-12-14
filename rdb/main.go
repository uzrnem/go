package main

import (
	"strings"
	"context"

	"gorm.io/gorm"
)

var (
	Repo Repository
)

type RepoStrct struct {
	DB *gorm.DB
}

func Load(db *gorm.DB) error {
	Repo = &RepoStrct{DB: db}
	return nil
}

// Create is a function that creates a new record in the repository.
//
// It takes in a `c` of type `context.Context` and `modal` of type `any` as parameters.
// It returns a value of type `any` and an `error`.
func (r *RepoStrct) Create(c context.Context, modal any) error {
	err := r.DB.Create(modal).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepoStrct) Delete(c context.Context, modal any, id int) error {
	err := r.DB.Delete(modal, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepoStrct) Get(c context.Context, modal any) (any, error) {
	err := r.DB.First(modal).Error
	if err != nil {
		return nil, err
	}
	return modal, nil
}

func (r *RepoStrct) Update(c context.Context, modal any) error {
	err := r.DB.Save(modal).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepoStrct) List(c context.Context, modal any) (any, error) {
	err := r.DB.Find(modal).Error
	if err != nil {
		return nil, err
	}
	return modal, nil
}

func (r *RepoStrct) FetchRow(sql string, res any) {
	r.DB.Debug().Raw(sql).Scan(res)
}

type Query struct {
	DB *gorm.DB
	table string
	silect string
	join string
	where string
	group string
	order string
	limit int
	offset int
	result any
}

func (r *RepoStrct) Builder() *Query {
	return &Query{DB: r.DB}
}
func (q *Query) Table(table string) *Query {
	q.table = table
	return q
}
func (q *Query) Select(silect string) *Query {
	q.silect = silect
	return q
}
func (q *Query) Join(join string) *Query {
	q.join = join
	return q
}
func (q *Query) Where(where string) *Query {
	q.where = where
	return q
}
func (q *Query) Group(group string) *Query {
	q.group = group
	return q
}
func (q *Query) Order(order string) *Query {
	q.order = order
	return q
}
func (q *Query) Limit(limit int) *Query {
	q.limit = limit
	return q
}
func (q *Query) Offset(offset int) *Query {
	q.offset = offset
	return q
}
func (q *Query) Exec(res any) error {
	db := q.DB.Debug()
	if strings.TrimSpace(q.table) != "" {
		db = db.Table(q.table)
	}
	if strings.TrimSpace(q.silect) != "" {
		db = db.Select(q.silect)
	}
	if strings.TrimSpace(q.join) != "" {
		db = db.Joins(q.join)
	}
	if strings.TrimSpace(q.where) != "" {
		db = db.Where(q.where)
	}
	if strings.TrimSpace(q.group) != "" {
		db = db.Group(q.group)
	}
	if strings.TrimSpace(q.order) != "" {
		db = db.Order(q.order)
	}
	if q.limit > 0 {
		db = db.Limit(q.limit).Offset(q.offset)
	}
	return db.Find(res).Error
}