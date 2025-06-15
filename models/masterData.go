package models

// MasterData represents the structure of the master data.
type MasterData struct { // Correct: Uppercase M for export
	Countries        []DataItem `json:"countries"`
	DeviceTypes      []DataItem `json:"deviceTypes"`
	MachineModels    []DataItem `json:"machineModels"`
	Status           []DataItem `json:"status"`
	ReturnableStatus []DataItem `json:"returnableStatus"`
	ReturnReason     []DataItem `json:"returnReason"`
}

// DataItem represents a generic item with a name and a code.
type DataItem struct { // Correct: Uppercase D for export
	Name string `json:"name"`
	Code string `json:"code"`
}
