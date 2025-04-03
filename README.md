# sqlite

一个不使用cgo的database/sql标准sqlite3驱动,原驱动`modernc.org/sqlite`,增加对goframe2+框架的支持(最新goframe已经有不使用cgo驱动了).

### 常规使用方法
~~~
	"github.com/jmoiron/sqlx"
	_ "github.com/logoove/sqlite"
    
  var db *sqlx.DB

type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

func main() {
	db, _ = sqlx.Open("sqlite", "./db.db")
	db.Exec(`CREATE TABLE users (
  "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  "name" char(15) NOT NULL
);`)
	db.Exec(`insert into users(name) values(?)`, "李白")
	db.Exec(`insert into users(name) values(?)`, "白居易")
	var rows []User
	db.Select(&rows, "SELECT id,name FROM users ORDER BY id")
	fmt.Println(rows)
}  
~~~

### 更新日志
2025-04-03 v1.37.0 ,同步到官方1.37.0版本,主要升级有升级到 SQLite 3.49.0,内存消耗问题,事物相关问题.
2023-10-09 v1.16.1 修复对goframe支持,更新到gitlab.com/cznic/sqlite最新1.26同步.对于goframe2以下版本可能不能使用.

2022-3-22 v1.15.3 新增win amd64编译,解决内存泄漏问题.

2021-11-3 v1.13.0 新增更多系统编译,支持sqlite 3.36.0


### goframe2.0+使用方法
在manifeat/config/config.yaml配置
~~~
_ "github.com/logoove/go/sqlite" //加在internal\cmd\cmd.go文件中

# 数据库连接配置
database:
  default:
    type: "sqlite"
    link: "./resource/db.db" #数据库路径根据自己的填写
    debug:  true
~~~

### goframe2.0+ 非配置
~~~
"github.com/gogf/gf/v2/database/gdb"
_ "github.com/logoove/go/sqlite" //这是专门用于goframe2.0以上的驱动, 单文件写在main.go中, mvc架构放在internal\cmd\cmd.go

db, _ := gdb.New(gdb.ConfigNode{
		Link: "sqlite::@file(xx/db.db)", //支持相对路径和绝对路径
	})
db.SetDebug(true) //可以显示sql语句
list, _ := db.Model("zanzhu").Where("1=1").Order("id desc").Limit(page, pagesize).All() //简单查询, 可以不用建立数据表结构体

~~~