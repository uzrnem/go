package rdb

import "context"

type Repository interface {
	Builder() *Query
	Create(c context.Context, modal any) error
	Delete(c context.Context, modal any, id int) error
	DeleteOne(c context.Context, modal any) error
	Get(c context.Context, modal any) (any, error)
	Update(c context.Context, modal any) error
	List(c context.Context, modal any) (any, error)
	FetchRow(sql string, res any)
}
