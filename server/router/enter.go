package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/video"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Video   video.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
