package models

type FullNameReq struct {
	NameID int `json:"name_id"`
}

type FullNameRet struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}
