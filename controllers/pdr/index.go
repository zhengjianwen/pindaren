package pdr

import (
	"github.com/zhengjianwen/utils/log"
	"github.com/zhengjianwen/pdrgo/controllers"
)

type Index struct {
	controllers.Base
}

func (pdr *Index)Get()  {
	pdr.TplName = "user/index.html"
	pdr.Data["name"] = "首页"

	log.Debugf("[index] %#v",pdr)

}
