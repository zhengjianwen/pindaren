package models

func Initdb()  {
	init_class()
}

func init_class()  {
	bean := new(Classify)
	n,err := Orm.Count(bean)
	if err != nil{
		return
	}
	if n == 0{
		var names = []string{"体验报告","购物晒单"," 求助","爆料"}
		for _,name := range names{
			bean.Name = name
			Orm.InsertOne(bean)
		}
	}
}
