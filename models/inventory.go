package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Archive     Archive
	Employees   []EmployeeInventory `json:"employees"`
}

type RequestInventory struct {
	InventoryName        string `json:"inventory_name"`
	InventoryDescription string `json:"Inventory_description"`
	ArchiveName          string `json:"archive_name"`
	ArchiveDescription   string `json:"archive_description"`
}

type ResponseInventory struct {
	InventoryName        string `json:"inventory_name"`
	InventoryDescription string `json:"Inventory_description"`
	Archive              ResponseArchive
}

type ResponseInventoryEmployee struct {
	InventoryName        string `json:"inventory_name"`
	InventoryDescription string `json:"Inventory_description"`
	EmployeeInventory    []RespEmployeeInventory
}
