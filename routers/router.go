package routers

import (
	"github.com/MarxGo/h5upload/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:ToIndex")
	uploadNS := beego.NewNamespace("/upload",

		beego.NSRouter("/to", &controllers.UploadController{}, "get:ToUpload"),
		beego.NSRouter("/checkFileExist", &controllers.UploadController{}, "post:CheckFileExist"),
		beego.NSRouter("/checkFileBlockExist", &controllers.UploadController{}, "post:CheckFileBlockExist"),
		beego.NSRouter("/getBlockSizeAndWorkerNum", &controllers.UploadController{}, "post:GetBlockSizeAndWorkerNum"),
		beego.NSRouter("/receiveFile", &controllers.UploadController{}, "post:ReceiveFile"),
		beego.NSRouter("/empty", &controllers.UploadController{}, "post:Empty"),
	)

	beego.AddNamespace(uploadNS)
}
