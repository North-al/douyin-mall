package model

import (
	"context"
	"gorm.io/gorm"
)

type Product struct{
	Base
	Name string `json:"name"`
	Description string `json:"description"`
	Picture string `json:"picture"`
	Price float32 `json:"price"`
	Categories []Category `json:"categories" gorm:"many2many:product_categories;"`

}

func (p Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db *gorm.DB
}

func (p ProductQuery) GetById(ProductId int) (product Product, err error) {
	err=p.db.WithContext(p.ctx).Model(&Product{}).First(&product, ProductId).Error
	return 

}
	
func(p ProductQuery) SearchProducts(q string) (Products []*Product, err error){
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&Products, "name like ? or description like ?",
		"%"+q+"%","%"+q+"%",
	).Error
	return
}

func NewProductQuery(ctx context.Context,db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db: db,
	}
}