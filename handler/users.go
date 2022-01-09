package handler

import (
	"fmt"
	"golibrary/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// users 管理员管理用户的页面
func UsersHander(c *gin.Context) {

	// c.HTML(http.StatusOK, "admin/users.html", nil)

	usersAll, err := db.SelectAllUsers()
	if err != nil {
		fmt.Printf("UsersHander err:%v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err,
		})
	}
	fmt.Printf("users:%v\n", usersAll)
	if len(*usersAll) > 0 {
		c.HTML(http.StatusOK, "admin/users.html", gin.H{
			"num":  1, // http://localhost:8080/admin/users 路径的admin/后面有几级(level)
			"data": *usersAll,
		})
	} else {
		c.HTML(http.StatusOK, "admin/users.html", gin.H{
			"num":  1,
			"data": nil,
		})
	}
}

// 修改用户信息
func ChangeUsersHandler(c *gin.Context) {
	fmt.Printf("c:%#v,%T\n", c, c)
	// var u db.Users
	// if err := c.ShouldBind(&u); err == nil {
	idStr := c.PostForm("id")
	flag := c.PostForm("flag")
	value := c.PostForm("value")
	// string到int
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Printf("修改用户信息 err:%v\n", err)
	}
	// 更新book数据
	err = db.UpdateRowUser(idInt, &flag, &value)
	if err != nil {
		fmt.Printf("修改用户信息 数据库 err:%v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"result": 0})

}

// 删除用户信息
func DeleteUsersHandler(c *gin.Context) {
	idStr := c.PostForm("id")
	// string到int
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Printf("删除用户信息 err:%v\n", err)
	}
	// 删除用户数据
	tab := "tab_users"
	err = db.DeleteRow(&tab, idInt)
	if err != nil {
		fmt.Printf("删除用户信息 数据库 err:%v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"result": 0})
}

// 添加用户
func AddUsersHandler(c *gin.Context) {
	var newUser db.Users
	// name属性传值
	// id := c.PostForm("id")
	// username := c.PostForm("username")
	// password := c.PostForm("password")
	// fmt.Printf("id:%v,username:%v,password:%v\n", id, username, password)
	if err := c.ShouldBind(&newUser); err == nil { //{ID:1}
		fmt.Printf("newUser:%v\n", newUser)
		err := db.Adduser(&newUser)
		if err != nil {
			fmt.Printf("err:%v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		// 成功则重定向到原页面
		c.Redirect(http.StatusMovedPermanently, "/admin/users")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
