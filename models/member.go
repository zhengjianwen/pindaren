package models

// 权限表
type Member struct {
	ID  uint64 `json:"id" orm:"rel(fk);auto;unique;size(14)"`
	UserID uint64 `json:"user_id"` // 用户ID
	GroupID uint64 `json:"group_id"` // 用户组ID
	IsAdmin uint64 `json:"is_admin"` // 是否是管理员

}

