package db

import (
	"database/sql"
	"fmt"
	"golibrary/cfg"

	_ "github.com/go-sql-driver/mysql" //init()
)

// 全局数据库对象
var DB *sql.DB

// 执行建表语句
func createTable(sql *string) error {
	_, err := DB.Exec(*sql)
	if err != nil {
		return err
	}
	return nil
}

// 初始化数据库	三张表
func InitDB(c *cfg.Config) (err error) {

	dbc := c.Connection // 结构体c MySQL部分
	// dsn := "root:123456@tcp(127.0.0.1:3306)/golibrary"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbc.User, dbc.Password, dbc.Host, dbc.Port, dbc.Database)

	DB, err = sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {                  // dsn格式不对会报错
		fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)
		return
	}

	err = DB.Ping() //尝试连接数据库
	if err != nil {
		fmt.Printf("Open %s failded,err:%v\n", dsn, err)
		return
	}

	// 设置数据库连接池最大连接数
	DB.SetMaxOpenConns(10)
	// 设置数据库最大空闲数
	DB.SetMaxIdleConns(5)

	// sql语句，如果没存在库表tab_books，则新建一个
	var sqlStr = `CREATE TABLE IF NOT EXISTS tab_books  (
		id int(0) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
		title varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '图书的标题',
		author varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '图书的作者',
		state tinyint(0) NULL DEFAULT 1 COMMENT '图书的状态,0为已借出',
		content text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '图书的内容',
		picture varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '图书的图片,相对静态文件路径',
		tradingTime varchar(25) NULL DEFAULT NULL COMMENT '交易时间(现在),时间字符串',
		PRIMARY KEY (id) USING BTREE
		) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;`

	// 执行建表语句
	err = createTable(&sqlStr)
	if err != nil {
		return err
	}

	// sql语句，如果没存在库表tab_user，则新建一个
	sqlStr = `CREATE TABLE IF NOT EXISTS tab_users  (
		id int(0) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
		username varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户名',
		password varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户密码',
		book_id int(0) NULL DEFAULT 0 COMMENT '借的书的id,0则表示没有借书',
		reputation int(0) NULL DEFAULT 100 COMMENT '用户信誉,默认100',
		PRIMARY KEY (id) USING BTREE
	  ) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;`

	// 执行建表语句
	err = createTable(&sqlStr)
	if err != nil {
		return err
	}

	// sql语句，如果没存在库表tab_record，则新建一个
	sqlStr = `CREATE TABLE IF NOT EXISTS tab_record  (
		id int(0) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
		user_id int(0) NULL DEFAULT NULL COMMENT '借书人的id',
		book_id int(0) NULL DEFAULT NULL COMMENT '借出或归还书的id',
		method tinyint(0) NULL DEFAULT NULL COMMENT '事件,0为借书,1为还书',
		tradingTime varchar(25) NULL DEFAULT NULL COMMENT '交易时间(现在),时间字符串',
		PRIMARY KEY (id) USING BTREE
	  ) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;`

	// 执行建表语句
	err = createTable(&sqlStr)
	if err != nil {
		return err
	}
	return
}
