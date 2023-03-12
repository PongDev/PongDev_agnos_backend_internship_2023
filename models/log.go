package models

type Log struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Request  string `json:"request"`
	Response string `json:"response"`
}
