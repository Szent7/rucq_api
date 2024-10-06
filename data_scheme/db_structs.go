package data_scheme

type User struct {
	Id       int64
	Login    string `pg:"type:varchar(64)" json:"login"`
	Password string `pg:"type:varchar(128)" json:"password"`
}

type Rooms struct {
	Id       int64
	Users    string `pg:"type:json"`
	Messages string `pg:"type:json"`
	Secret   string `pg:"type:varchar(128)"`
}

type Messages struct {
	MessageList []MesContainer `json:"messagelist"`
}

type MesContainer struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type MesContainerSend struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Secret   string `json:"secret"`
}

type Users struct {
	UsersList []UsersMap `json:"userslist"`
}

type UsersMap struct {
	Login    string `json:"login"`
	Username string `json:"username"`
}
