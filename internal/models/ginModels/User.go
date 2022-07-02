package ginModels

type UserModel struct {
	UserID  string `json:"user_id"`
	IsAdmin bool   `json:"is_admin"`
}

func (u UserModel) VerifyAdminRole() bool {
	if u.IsAdmin {
		return true
	}
	return false
}
