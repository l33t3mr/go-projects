package main

import (
	"github.com/gin-gonic/gin"
	"github.com/l33t3mr/go-projects/shortclips/controller"
	"github.com/l33t3mr/go-projects/shortclips/model"
)

var clips []model.Clip

func main() {

	r := gin.Default()
	db, err := model.ConnectToDB()
	if err != nil {
		panic(err)
	}
	// do migrations
	model.MigrateModels(db, []model.Migrator{model.Clip{}, model.Director{}})
	r.GET(controller.GetClipRoute, controller.GetClip)
	r.GET(controller.GetAllCLIPRoute, controller.GetAllClips)
	r.POST(controller.PostClipRoute, controller.PostClip)

	r.Run()

}
