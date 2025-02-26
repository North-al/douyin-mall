
package model

type OrderItem struct {
	Base
	ProductId    uint32
	OrderIdRefer string `gorm:"size:256;index"`
	Quantity     int32
	Cost         float32
}

func (oi OrderItem) TableName() string {
	return "order_item"
}

// package model

// import "gorm.io/gorm"

// type OrderItem struct {
// 	gorm.Model
// 	ProductId    uint32  `gorm:"type:int(11)"`
// 	OrderIdRefer string  `gorm:"type:varchar(100);index"`
// 	Quantity     int32   `gorm:"type:int(11)"`
// 	Cost         float32 `gorm:"type:decimal(10,2)"`
// }

// func (OrderItem) TableName() string {
// 	return "order_item"
// }