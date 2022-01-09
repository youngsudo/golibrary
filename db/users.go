package db

import "fmt"

// 查询所有用户信息
func SelectAllUsers() (*[]Users, error) {
	var usersList = make([]Users, 0, 20) //传一个切片
	var user Users
	sqlStr := `SELECT id,username,password,book_id,reputation FROM tab_users;`

	rows, err := DB.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	// 3,一定要关闭连接rows
	defer rows.Close()
	// 4,循环取值
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Book_id, &user.Reputation)
		if err != nil {
			return nil, err
		}
		usersList = append(usersList, user)
	}
	return &usersList, nil
}

//修改用户信息
func UpdateRowUser(id int, flag, value *string) error {
	// 如果用户输入的值为空则将该数据设为null
	var val = value
	if *val == "" {
		*val = "null"
	}
	sqlStr := fmt.Sprintf("update tab_users set %v = %v where id = ?;", *flag, *val)
	ret, err := DB.Exec(sqlStr, id)

	if err != nil {
		return err
	}
	n, err := ret.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("更新了%d行数据\n", n)
	return nil
}

// 添加用户信息
func Adduser(newUser *Users) (err error) {
	// fmt.Printf("newUser:%#v\n", newUser)
	// fmt.Printf("newUser:%#v\n", newUser.Username)
	sqlStr := `INSERT INTO tab_users (id,username,password) VALUES (?,?,?);`

	ret, err := DB.Exec(sqlStr, newUser.ID, newUser.Username, newUser.Password)
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

// 查询多条数据（模糊查询）
func SelectLikeUsers(value *string) (*[]Users, error) {

	sqlStr := fmt.Sprintf("SELECT id,username,password,book_id,reputation FROM tab_users WHERE id = '%s' or username like '%%%s%%' or reputation like '%%%s%%' ", *value, *value, *value)
	var usersList = make([]Users, 0, 20)
	var user Users //传一个切片

	rows, err := DB.Query(sqlStr)
	if err != nil {
		fmt.Printf("exec %s SelectLikeUsers faild,err:%#v\n", sqlStr, err)
		return nil, err
	}
	// 3,一定要关闭连接rows
	defer rows.Close()
	// 4,循环取值
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Book_id, &user.Reputation)
		if err != nil {
			fmt.Printf("scan faild err:%v\n", err)
			return nil, err
		}
		usersList = append(usersList, user)
	}
	return &usersList, nil
}
