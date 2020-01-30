package main

import (
	"github.com/aliereno/softwareproject-api/internal/handlers"
	"github.com/aliereno/softwareproject-api/internal/handlers/auth"
	controllersMain "github.com/aliereno/softwareproject-api/internal/handlers/controllers"
	log "github.com/aliereno/softwareproject-api/internal/logger"
	"github.com/aliereno/softwareproject-api/internal/orm"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var host, port string

func init() {
	host = "localhost"
	port = "7777"
}

func main() {
	orm, err := orm.Factory()
	controllers := controllersMain.Controller{ORM: orm}
	if err != nil {
		log.Panic(err)
	}
	endpoint := "http://" + host + ":" + port
	r := gin.Default()

	r.Use(cors.Default())

	// Static Files
	r.Static("/files", "./files/ringtones")

	r.GET("/ping", handlers.Ping())
	r.POST("/login", controllers.UserLogin())
	r.POST("/register", controllers.UserRegister())

	r.GET("/recent-activities", controllers.UserFetchRecentActivities())
	r.GET("/ringtones-featured", controllers.RingtonesFetchFeatued())
	r.GET("/ringtones", controllers.RingtonesFetchAll())
	r.GET("/ringtones/:id", controllers.RingtoneFetchDetail())
	r.GET("/categories", controllers.CategoriesFetchAll())
	r.GET("/categories/:id", controllers.CategoriesFetchDetail())

	api := r.Group("/user")
	api.Use(auth.LookUserTokenHandler())
	{
		api.GET("/settings/:id", controllers.UserFetchSetting())
		api.POST("/settings", controllers.UserUpdateSetting())
		// RINGTONE
		ringtone := api.Group("/ringtone")
		{
			ringtone.POST("/buy", controllers.UserPurchaseRingtone())
			ringtone.POST("/like", controllers.UserLikeRingtone())
			ringtone.POST("/dislike", controllers.UserDislikeRingtone())
			ringtone.POST("/comment", controllers.UserCommentRingtone())
		}
	}

	//admin := r.Group("/admin")
	//admin.Use(auth.LookAdminTokenHandler())
	//{
	//	// USER
	//	users := admin.Group("/users")
	//	{
	//		users.GET("/", controllers.UsersFetchAll())
	//		users.POST("/create", controllers.UsersFetchAll())
	//		users.POST("/delete/:id", controllers.UsersFetchAll())
	//		users.POST("/update/:id", controllers.UsersFetchAll())
	//	}
	//
	//	// RINGTONE
	//	ringtones := admin.Group("/ringtones")
	//	{
	//		ringtones.GET("/", controllers.RingtonesFetchAll())
	//		ringtones.GET("/:id", controllers.RingtoneFetchDetail())
	//		ringtones.POST("/create", controllers.RingtonesFetchAll())
	//		ringtones.POST("/delete/:id", controllers.RingtonesFetchAll())
	//		ringtones.POST("/update/:id", controllers.RingtonesFetchAll())
	//	}
	//}

	log.Info("Running @ " + endpoint)

	log.Fatal(r.Run(host + ":" + port))
}
