package db

import "fmt"

// 由于什么参数都不特殊,就可以拿出来
// 删除操作 exec	tab表名 id
func DeleteRow(tab *string, id int) (err error) {

	sqlStr := fmt.Sprintf("DELETE FROM %v where id = ?", *tab)

	ret, err := DB.Exec(sqlStr, id)
	if err != nil {
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		return
	}
	fmt.Printf("删除了%d行数据\n", n)
	return nil
}
