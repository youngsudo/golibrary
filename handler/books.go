package handler

import (
	"fmt"
	"golibrary/db"
	"log"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type Book_id struct {
	ID int `sql:"id" json:"id" form:"id"`
}

func BooksHander(c *gin.Context) {

	bookAll, err := db.SelectAllBooks()
	if err != nil {
		panic("tab_book select err")
	}
	// fmt.Printf("books:%#v\n", booksList) //切片数组
	// for i, v := range booksList {
	// 	fmt.Println(i, v)
	// }
	if len(*bookAll) > 0 {
		c.HTML(http.StatusOK, "admin/books.html", gin.H{
			"num":  1,
			"data": *bookAll,
		})
	} else {
		c.HTML(http.StatusOK, "admin/books.html", gin.H{
			"num":  1,
			"data": nil,
		})
	}
	// fmt.Printf("%v\n", getTime.GetTime())	时间
}

// 添加书get
func AddBooks(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/addBooks.html", gin.H{
		"num":  2,
		"data": nil,
	})
}

// 添加书 写入数据库
func AddBooks_post(c *gin.Context) {
	// 从请求中读取文件
	file, err := c.FormFile("image") //和从请求中获取携带的参数一样的	image 文件名name
	title := c.PostForm("title")
	author := c.PostForm("author")
	content := c.PostForm("content")
	// fmt.Println(title, content, author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Println(file.Filename)
	dst := path.Join("./static", "img", file.Filename) //通过path标准库	path指的是当前项目的绝对路径
	// 上传文件到指定的目录（服务端本地）
	c.SaveUploadedFile(file, dst)

	picture := fmt.Sprintf("/img/%s", file.Filename)
	err = db.AddRowBook(&title, &author, &content, &picture)

	if err != nil {
		fmt.Printf("添加图书err:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"data": "添加失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "添加成功",
	})
}

// 删除书	要传表名,因为删除可以同一个方法
func DeleteBooks(c *gin.Context) {

	var _id Book_id

	if err := c.ShouldBind(&_id); err == nil { //{ID:1}

		tab_name := "tab_books"
		err := db.DeleteRow(&tab_name, _id.ID)
		if err != nil {
			fmt.Printf("err:%v\n", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": 0,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// 查书 详情
func DetailsBooks(c *gin.Context) {
	var _id Book_id
	// var b *db.Books
	if err := c.ShouldBind(&_id); err == nil { //{ID:1}

		book := db.DetailsRowBook(_id.ID)

		c.HTML(http.StatusOK, "admin/bookDetails.html", gin.H{
			"num":  2,
			"data": *book,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

// 修改书 详情
func ChangeBooks(c *gin.Context) {
	var b db.Books
	if err := c.ShouldBind(&b); err == nil { //{ID:1}
		// 更新book数据
		err = db.UpdateRowBook(&b)
		if err != nil {
			fmt.Printf("err:%v\n", err)
		}
		c.JSON(http.StatusOK, gin.H{"result": 0})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
