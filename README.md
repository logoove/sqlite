# sqlite

一个不使用cgo的database/sql标准sqlite3驱动,原驱动`modernc.org/sqlite`,但是速度不行,在github上边没有,就从gitlab搬迁过来,同时修改了驱动名称为sqlite3,这样不用任何修改,就能在goframe框架中使用.

### 使用方法
`_ github.com/logoove/sqlite`

支持goframe的orm框架,也支持原生'sqlx'
`"github.com/jmoiron/sqlx"`
原生使用例子
~~~
var db *sqlx.DB
	type User struct{
		Id int64 `db:"id"`
		Name string `db:"name"`
	}
	db, _ = sqlx.Open("sqlite3","./db.db")
db.Exec(`CREATE TABLE users (
  "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  "name" char(15) NOT NULL
);`)
	db.Exec(`insert into users(name) values(?)`,"李白")
	db.Exec(`insert into users(name) values(?)`,"白居易")
	rows:=[]User{}
	db.Select(&rows,"SELECT id,name FROM users ORDER BY id")
	fmt.Println(rows)
~~~

goframe1.16中只需在boot 下面go文件里面加上 `_ github.com/logoove/sqlite`即可在orm使用,不需要任何其他修改
goframe2.0中,在manifeat/config/config.yaml配置
~~~
# 数据库连接配置
database:
  logger:
    path:    "./temp/logs/sql"
    level:   "all"
    stdout:  true
    ctxKeys: ["RequestId"]
  default:
    type: "sqlite"
    link: "./resource/db.db" #数据库路径根据自己的填写
    debug:  true
~~~
,在internel/cmd/cmd.go 中加入_ "github.com/logoove/sqlite"即可,已经将驱动改成sqlite3,所以能够直接在goframe中使用.
插入数据
~~~
id, _ := g.Model("user").Data(g.Map{"name": "john", "age": 1}).InsertAndGetId()
~~~
### 更新日志
2022-3-22 v1.15.3 新增win amd64编译,解决内存泄漏问题.
2021-11-3 v1.13.0 新增更多系统编译,支持sqlite 3.36.0