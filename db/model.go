package db

// 定义数据库的访问模型
// 图书结构体
type Books struct {
	// 结构体成员tag，包含3种，一种是数据库，一种是json，一种是form表单
	ID          int    `sql:"id" json:"id" form:"id"`                            // 唯一标识的
	Title       string `sql:"title" json:"title" form:"title"`                   // 图书的标题
	Author      string `sql:"author" json:"author" form:"author"`                // 图书的作者
	State       int    `sql:"state" json:"state" form:"state"`                   // 图书的状态，为true表示已完成，为false表示未完成
	Content     string `sql:"content" json:"content" form:"content"`             // 图书的内容
	Picture     string `sql:"picture" json:"picture" form:"picture"`             // 图书的内容
	TradingTime string `sql:"tradingTime" json:"tradingTime" form:"tradingTime"` //交易时间 时间字符串
}

//  用户结构体
type Users struct {
	ID         int    `sql:"id" json:"id" form:"id"`
	Username   string `sql:"username" json:"username" form:"username"`
	Password   string `sql:"password" json:"password" form:"password"`
	Book_id    int    `sql:"book_id" json:"book_id" form:"book_id"`          // 借的书的id,null则表示没有借书
	Reputation int    `sql:"reputation" json:"reputation" form:"reputation"` // 用户信誉,默认100
}

// 借阅记录结构体
type Record struct {
	ID          int    `sql:"id" json:"id" form:"id"`
	User_id     int    `sql:"user_id" json:"user_id" form:"user_id"`
	Book_id     int    `sql:"book_id" json:"book_id" form:"book_id"`
	Method      int    `sql:"method" json:"method" form:"method"`                // 借阅事件,0为借书,1为还书
	TradingTime string `sql:"tradingTime" json:"tradingTime" form:"tradingTime"` //交易时间 时间字符串
}

// 借阅记录结构体展示
type RecordHTML struct {
	Record
	Username   string `sql:"username" json:"username" form:"username"`
	Reputation int    `sql:"reputation" json:"reputation" form:"reputation"` // 用户信誉,默认100
	Title      string `sql:"title" json:"title" form:"title"`                // 图书的标题
}
