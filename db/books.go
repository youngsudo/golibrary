package db

import (
	"fmt"
	"golibrary/getTime"
)

// 查询多条(全部)数据
func SelectAllBooks() (*[]Books, error) {

	var booksList = make([]Books, 0, 20) // 声明切片并返回切片
	var book Books
	sqlStr := `SELECT id,title,author,state,content,picture,tradingTime FROM tab_books;`

	rows, err := DB.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	// 3,一定要关闭连接rows
	defer rows.Close()
	// 4,循环取值
	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.State, &book.Content, &book.Picture, &book.TradingTime)
		if err != nil {
			return nil, err
		}
		// 追加切片元素
		booksList = append(booksList, book)
	}
	return &booksList, nil

}

// 单条数据查询	详情
func DetailsRowBook(id int) *Books {
	var book Books
	sqlStr := `Select id,title,author,state,content,picture,tradingTime from tab_books where id = ?;`
	rowObj := DB.QueryRow(sqlStr, id) //必须传指针,sql,参数
	rowObj.Scan(&book.ID, &book.Title, &book.Author, &book.State, &book.Content, &book.Picture, &book.TradingTime)
	return &book
}

// 查询多条数据（模糊查询）
func SelectLikeBooks(value *string) (*[]Books, error) {
	var booksList = make([]Books, 0, 20) // 声明切片并返回切片
	var book Books
	sqlStr := fmt.Sprintf("SELECT id,title,author,state,content,picture,tradingTime FROM tab_books WHERE title like '%%%s%%' or author like '%%%s%%' or content like '%%%s%%' ", *value, *value, *value)

	rows, err := DB.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	// 3,一定要关闭连接rows
	defer rows.Close()
	// 4,循环取值
	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.State, &book.Content, &book.Picture, &book.TradingTime)
		if err != nil {
			return nil, err
		}
		// 追加切片元素
		booksList = append(booksList, book)
	}
	return &booksList, nil
}

// 添加操作 exec	tab_book
func AddRowBook(title, author, content, picture *string) (err error) {
	tradingTime := getTime.GetTime()
	sqlStr := `INSERT INTO tab_books (title,author,content,picture,tradingTime) VALUES (?,?,?,?,?);`

	ret, err := DB.Exec(sqlStr, title, author, content, picture, tradingTime)
	if err != nil {
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		return
	}
	fmt.Printf("添加了%d行数据\n", n)
	return nil
}

// 更新操作 exec
func UpdateRowBook(b *Books) (err error) {

	// b:db.Books{ID:26, Title:"laor", Author:"1", State:1, Content:"222222222222", Picture:""}
	sqlStr := `update tab_books set title = ?,author= ?,state = ?,content= ? where id = ?;`
	ret, err := DB.Exec(sqlStr, b.Title, b.Author, b.State, b.Content, b.ID)
	if err != nil {
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		return
	}
	fmt.Printf("更新了%d行数据\n", n)
	return nil
}
