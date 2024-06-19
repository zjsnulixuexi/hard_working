package models

type Documents struct {
	Documentid  int    `json:"documentid"`
	Userid      int    `json:"userid"`
	Updatetime  string `json:"updateTime"`
	Filename    string `json:"filename"`
	IsCollected int    `json:"isCollected"`
	Path        string `json:"path"`
}

type Users struct {
	Userid   int    `json:"userid"`
	Password string `json:"password"`
	Username string `json:"username"`
	School   string `json:"school"`
	Sex      string `json:"sex"`
	Power    int    `json:"power"`
}

type Analysis struct {
	Documentid int    `json:"documentid"`
	Userid     int    `json:"userid"`
	Filename   string `json:"filename"`
	Path       string `json:"path"`
}

type Results struct {
	Userid     int    `json:"userid"`
	Documentid int    `json:"documentid"`
	Resultid   int    `json:"resultid"`
	Resultname string `json:"resultname"`
	Kind       int    `json:"kind"`
	Content    string `json:"content"`
}

type Visualoperations struct {
	Operationid        int    `json:"operationid"`
	Operationname      string `json:"operationname"`
	Operationintroduce string `json:"operationintroduce"`
	Operatinopath      string `json:"operationpath"`
}
