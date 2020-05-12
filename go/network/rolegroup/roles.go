package rolegroup

type Data_roles struct {
	Id       int64    `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	RoleList []string `json:"rolelist,omitempty"`
	Code     int      `json:"code"`
}

type List_roles struct {
	DataList []*Data_roles `json:"datalist"`
	Code     int           `json:"code"`
}

type Get_roles struct {
	Roles []string `json:"roles"`
	Code  int      `json:"code"`
}

type Updata_role struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Oldname string `json:"oldname"`
}
