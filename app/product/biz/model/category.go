package model

import (
	"context"
	"gorm.io/gorm"
)

type Category struct {
	Base
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `json:"product" gorm:"many2many:product_category;"`
}

func (c Category) TableName() string {
	return "category"
}

type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (c *CategoryQuery) SearchProducts(query string) (any, error) {
	panic("unimplemented")
}

func (c CategoryQuery) GetProductByCategoryName(name string) (Categories []Category, err error) {
	err = c.db.WithContext(c.ctx).Model(&Category{}).Where(&Category{Name: name}).Preload("Products").Find(&Categories).Error
	return
}

func NewCategoQuery(ctx context.Context, db *gorm.DB) *CategoryQuery {
	return &CategoryQuery{
		ctx: ctx,
		db:  db,
	}
}
