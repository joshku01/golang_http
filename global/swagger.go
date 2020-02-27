package global

type adminRegisterBody struct {
	// 帳號
	Account string `json:"account" example:"user02"`
	// 密碼
	Password string `json:"pwd" example:"qwe123"`
	// 再次確認密碼
	PasswordConfirmation string `json:"pwd_confirmation" example:"qwe123"`
	// 帳號權限 0:一般,1:超級使用者
	GroupID int `json:"group_id" example:"1"`
	// 部門名稱
	Department string `json:"department" example:"研發三處"`
}

// APIResult 回傳API格式
type APIResult struct {
	ErrorCode   int         `json:"error_code"`
	ErrorMsg    string      `json:"error_msg"`
	LogIDentity string      `json:"log_id"`
	Result      interface{} `json:"result"`
}
