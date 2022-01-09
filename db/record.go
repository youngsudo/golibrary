package db

import (
	"fmt"
	"golibrary/getTime"
)

// 查询所有用户信息
func SelectAllrecord() (*[]RecordHTML, error) {
	var recordsList = make([]RecordHTML, 0, 20) //传一个切片
	var record RecordHTML
	sqlStr := `SELECT r.id,r.user_id,r.book_id,r.method,r.tradingTime,u.username,u.reputation,b.title
	FROM  tab_record as r
	LEFT JOIN tab_books b
	ON r.book_id = b.id
	LEFT JOIN tab_users u
	ON r.user_id = u.id;`

	rows, err := DB.Query(sqlStr)
	if err != nil {
		fmt.Printf("exec %s Select faild,err:%#v\n", sqlStr, err)
		return nil, err
	}
	// 3,一定要关闭连接rows
	defer rows.Close()
	// 4,循环取值
	for rows.Next() {
		err = rows.Scan(&record.ID, &record.User_id, &record.Book_id, &record.Method, &record.TradingTime, &record.Username, &record.Reputation, &record.Title)
		if err != nil {
			fmt.Printf("scan faild err:%v\n", err)
			return nil, err
		}
		fmt.Printf("record:%v\n", record)
		recordsList = append(recordsList, record)
	}
	return &recordsList, nil
}

// 添加一条
func AddRecordDB(newRecord *Record) (err error) {
	// 获取现在时间日期
	TradingTime := getTime.GetTime()
	fmt.Println(newRecord.ID, newRecord.User_id, newRecord.Book_id, newRecord.Method, TradingTime)
	sqlStr := `INSERT INTO tab_record (id,user_id,book_id,method,tradingTime) VALUES (?,?,?,?,?);`

	ret, err := DB.Exec(sqlStr, newRecord.ID, newRecord.User_id, newRecord.Book_id, newRecord.Method, TradingTime)
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

//
func SelectLikerecord(value *string) (*[]RecordHTML, error) {

	var recordList = make([]RecordHTML, 0, 20)
	var record RecordHTML //传一个切片

	sqlStr := fmt.Sprintf(`SELECT r.id,r.user_id,r.book_id,r.method,r.tradingTime,u.username,u.reputation,b.title
	FROM  tab_record as r
	LEFT JOIN tab_books b
	ON r.book_id = b.id
	LEFT JOIN tab_users u
	ON r.user_id = u.id 
	where r.id = %v;`, *value)
	rows, err := DB.Query(sqlStr)
	if err != nil {
		fmt.Printf("exec %s SelectLikeUsers faild,err:%#v\n", sqlStr, err)
		return nil, err
	}
	// 3,一定要关闭连接rows
	defer rows.Close()
	// 4,循环取值
	for rows.Next() {
		err = rows.Scan(&record.ID, &record.User_id, &record.Book_id, &record.Method, &record.TradingTime, &record.Username, &record.Reputation, &record.Title)
		if err != nil {
			fmt.Printf("scan faild err:%v\n", err)
			return nil, err
		}

		recordList = append(recordList, record)
	}
	return &recordList, nil
}
