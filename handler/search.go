package handler

import (
	"fmt"
	"golibrary/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	0 全部
	1 书本	tab_books
	2 用户	tab_users
	3 记录	tab_record

*/
func SearchPOST(c *gin.Context) {

	t := c.PostForm("search_select")
	v := c.PostForm("search_input")
	// 当t不为nul,v 为空字符串时,由于是模糊查询,查询出来的结果是全部结果
	fmt.Printf("t:%v,v:%v\n", t, v)
	switch t {
	case "0":
		panic("Select error search.go")
	case "1":
		bookLike, err := db.SelectLikeBooks(&v) // 表名与关键字
		// fmt.Println("select like tab_books")
		if err != nil {
			fmt.Printf("select like tab_books faild,err:%v\n", err)
		}
		if len(*bookLike) > 0 {
			c.HTML(http.StatusOK, "admin/books.html", gin.H{
				"num":  1,
				"data": *bookLike,
			})
		} else {
			c.HTML(http.StatusOK, "admin/books.html", gin.H{
				"num":  1,
				"data": nil,
			})
		}
	case "2":
		// fmt.Println("select like tab_users")
		usersLike, err := db.SelectLikeUsers(&v) // 表名与关键字
		if err != nil {
			fmt.Printf("select like tab_users faild,err:%v\n", err)
		}
		if len(*usersLike) > 0 {
			c.HTML(http.StatusOK, "admin/users.html", gin.H{
				"num":  1,
				"data": *usersLike,
			})
			// } else {
			c.HTML(http.StatusOK, "admin/users.html", gin.H{
				"num":  1,
				"data": nil,
			})
		}
	case "3":
		recordLike, err := db.SelectLikerecord(&v) // 表名与关键字
		if err != nil {
			fmt.Printf("select like tab_record faild,err:%v\n", err)
		}
		if len(*recordLike) > 0 {
			c.HTML(http.StatusOK, "admin/record.html", gin.H{
				"num":  1,
				"data": *recordLike,
			})
		} else {
			c.HTML(http.StatusOK, "admin/record.html", gin.H{
				"num":  1,
				"data": nil,
			})
		}
	default:
		panic("Select error search.go")
	}

}
