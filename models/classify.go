package models

// 分类信息表
type Classify struct {
	ID  		uint64 		`json:"id" orm:"rel(fk);auto;unique;size(14)"`
	Name 		string 		`json:"name"`
}
