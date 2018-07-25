package models

type City struct {
	ID 			int64 	`json:"id"`
	Name 		string	`json:"name"` // 名称
	ParentID 	int64 	`json:"parent_id"` // 父级id
	Level 		int64 	`json:"level"` // 级别 0 省 1是市 2 县 3 乡
}

func (c *City) Validate() string {
	if c.Name == ""{
		return "名称不能为空"
	}
	if c.ParentID == 0{
		c.Level = 0
	}else {

	}

	return ""
} 