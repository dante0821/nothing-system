package vm

type PersonInfo struct {
	PersonId    int    `json:"person_id"`
	PersonName  string `json:"person_name"`
	Sex         string `json:"sex"`
	Age         int    `json:"age"`
	IdCard      string `json:"id_card"`
	Address     string `json:"address"`
	Party       string `json:"party"`
	Phone       string `json:"phone"`
	Birthday    string `json:"birthday"`
	CreatedTime int64  `json:"created_time"`
	UpdatedTime int64  `json:"updated_time"`
}

type GetPersonInfoListRsp struct {
	List  []PersonInfo `json:"list"`
	Count int          `json:"count"`
}

type CreateOrUpdatePersonInfo struct {
	PersonName string `json:"person_name"`
	Sex        string `json:"sex"`
	Age        int    `json:"age"`
	IdCard     string `json:"id_card"`
	Address    string `json:"address"`
	Party      string `json:"party"`
	Phone      string `json:"phone"`
	Birthday   string `json:"birthday"`
}
