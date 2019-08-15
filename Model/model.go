package model

type GerageStatus struct {
	OwnerName string `json:"ownerName"`
	CarName   string `json:"CarName"`
	IDCar     string `json:"IdCar"`
}

type DetailPosition struct {
	GarageName   string `json:"GarageName"`
	Longtitude   string `json:"Longtitude"`
	Latitude     string `json:"Latitude"`
	PositionName string `json:"PositionName"`
}

type DetailCarStatus struct {
	OwnerName    string `json:"ownerNme"`
	CarName      string `json:"CarName"`
	IDCar        string `json:"IdCar"`
	GarageName   string `json:"GarageName"`
	Longtitude   string `json:"Longtitude"`
	Latitude     string `json:"Latitude"`
	PositionName string `json:"PositionName"`
}
