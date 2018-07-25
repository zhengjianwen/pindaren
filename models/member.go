package models

// 权限表
type Member struct {
	Id  uint64 `json:"id"`
	UserId uint64 `json:"user_id"` // 用户ID
	GroupId uint64 `json:"group_id"` // 用户组ID
	IsAdmin uint64 `json:"is_admin"` // 是否是管理员

}

