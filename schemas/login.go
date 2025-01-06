package schemas

type LoginUserResponse struct {
	Menssage string `json:"menssage"`
	UserId   uint   `json:"user_id,omitempty "`
	Token    string `json:"token"`
}
