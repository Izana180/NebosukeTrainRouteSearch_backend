package model

type Station struct {
	Name string `json:"name" example:"渋谷" description:"駅名"`
	ID string `json:"id" example:"urn:ucode:_00001C0000000000000100000341B4E7" description:"ODPT内の駅ID"`
}