package constant

import "github.com/Shriharsh07/InventoryManagement/models"

var (
	Countries = []models.DataItem{
		{Name: "United States of America", Code: "US"},
		{Name: "India", Code: "IN"},
		{Name: "Germany", Code: "DE"},
	}
	DeviceTypes = []models.DataItem{
		{Name: "Laptop", Code: "LP"},
		{Name: "iPhone", Code: "iPhone"},
		{Name: "iPad", Code: "iPad"},
	}
	MachineModels = []models.DataItem{
		{Name: "Dell XPS 13", Code: "Dell"},
		{Name: "iPhone 13", Code: "Iphone"},
		{Name: "Samsung Galaxy Tab S8", Code: "SG"},
	}
	Status = []models.DataItem{
		{Name: "In Stock", Code: "InStock"},
		{Name: "Assigned", Code: "Assigned"},
		{Name: "Out Of Warranty", Code: "OutOfWarranty"},
		{Name: "In Transit", Code: "InTransit"},
		{Name: "Non-Compliant", Code: "NonCompliant"},
		{Name: "Returnable", Code: "Returnable"},
		{Name: "Recalled", Code: "Recalled"},
		{Name: "Off-Boarding", Code: "Offboarding"},
	}
	ReturnableStatus = []models.DataItem{
		{Name: "Returnable", Code: "Returnable"},
		{Name: "Recalled", Code: "Recalled"},
		{Name: "Off-Boarding", Code: "Offboarding"},
	}
	ReturnReason = []models.DataItem{
		{Name: "Damaged", Code: "Damaged"},
		{Name: "Expired", Code: "Expired"},
		{Name: "Incorrect Item", Code: "Incorrect Item"},
		{Name: "Defective", Code: "Defective"},
		{Name: "Recalled", Code: "Recalled"},
		{Name: "Overstock", Code: "Overstock"},
		{Name: "Customer Return", Code: "Customer Return"},
	}
)
