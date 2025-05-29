package model

// 経路全体
type Route struct {
	Steps     []RouteStep `json:"steps"`
	TotalTime int         `json:"total_time" example:"30" description:"所要時間"`
	TotalFare int         `json:"total_fare" example:"270" description:"合計運賃(円)"`
}

// 乗り換えごとの一区間
type RouteStep struct {
	DepartureStation string `json:"departure_station" example:"新橋" description:"出発駅"`
	ArrivalStation   string `json:"arrival_station" example:"新川崎" description:"到着駅"`
	Fromtime         string `json:"from_time" example:"2025-05-28T10:01:00+09:00" description:"出発時間"`
	Totime           string `json:"to_time" example:"2025-05-28T10:13:00+09:00" description:"到着時間"`
	Movetype         string `json:"movetype" example:"local_train" description:"移動手段"`
}
