package routers

import (
	"golibrary/handler"
	"html/template"

	"github.com/gin-gonic/gin"
	// "golibrary/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 自定义模板函数
	router.SetFuncMap(template.FuncMap{
		"noImg": func(str string) string { // 没有图片
			if str == "0" {
				return "/img/01.jpg"
			}
			return str
		},
		"stated": func(i int) string { // 是否借出
			if i == 0 {
				return "已借出"
			} else {
				return "未借出"
			}
		},
		"method": func(i int) string { // Record中是借还是还
			if i == 0 {
				return "借书"
			} else {
				return "还书"
			}
		},
		"level": func(i int) string { // 静态文件路径
			level := ""
			for n := 0; n < i; n++ {
				level += "../"
			}
			return level
		},
	})
	// 告诉gin框架模板文件引用的静态文件去哪里找
	router.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	router.LoadHTMLGlob("templates/**/*")

	// 路由组
	adminGroup := router.Group("admin")
	{
		// 首页
		adminGroup.GET("/index", handler.IndexHandler)
		// books 图书管理
		adminGroup.GET("/books", handler.BooksHander)
		// users 用户管理
		adminGroup.GET("/users", handler.UsersHander)
		// record 借阅归还记录
		adminGroup.GET("/record", handler.RecordsHander)
		// 查询框 模糊查询
		adminGroup.POST("/search", handler.SearchPOST)
	}
	// 嵌套路由组	books
	booksGroup := adminGroup.Group("books")
	{
		// add	增
		booksGroup.GET("/add", handler.AddBooks)       //页面
		booksGroup.POST("/add", handler.AddBooks_post) //数据
		// Delete	删
		booksGroup.POST("/delete", handler.DeleteBooks)
		// change	改
		booksGroup.GET("/bookDetails", handler.DetailsBooks) //书本详情页面
		booksGroup.POST("/bookDetails", handler.ChangeBooks) //修改书

		// query 	查书
	}

	// 嵌套路由组	users
	usersGroup := adminGroup.Group("users")
	{
		// add	增
		// booksGroup.GET("/add", handler.AddBooks)       //页面
		usersGroup.POST("/add", handler.AddUsersHandler) //数据
		// // Delete	删
		usersGroup.POST("/delete", handler.DeleteUsersHandler)
		// // change	改
		usersGroup.POST("/change", handler.ChangeUsersHandler) //修改用户信息

		// query 	查用户
	}
	// 嵌套路由组	record
	recordGroup := adminGroup.Group("record")
	{
		// add	增
		recordGroup.POST("/add", handler.AddRecordHandler) //数据
		// Delete	删
		recordGroup.POST("/delete", handler.DeleteRecordHandler)
		// // change	改
		// recordGroup.POST("/change", handler.ChangeRecordHandler) //修改用户信息

		// query 	借阅归还记录
	}
	return router
}
