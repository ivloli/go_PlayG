package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DataMarketBase struct {
	DataType         int    `gorm:"column:data_type"`
	PackID           int    `gorm:"column:pack_id;primary_key"`
	PackType         int    `gorm:"column:pack_type"`
	PackDataSource   int    `gorm:"column:pack_data_source"`
	PackDataURL      string `gorm:"column:pack_data_url"`
	PackName         string `gorm:"column:pack_name"`
	SystemPreference int    `gorm:"column:system_preference"`
	SupplierID       int64  `gorm:"column:supplier_id"`
	IsAutoUp         int    `gorm:"column:is_auto_up"`
}

func main() {
	fmt.Println("vim-go")
}
