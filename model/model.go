package model

type Row struct {
	Cid    int    `json:"ComputerID"`
	UserID int    `json:"UserID"`
	AppID  int    `json:"ApplicationID"`
	CType  string `json:"ComputerType"`
}
