# sqlite

一个不使用cgo的database/sql标准sqlite3驱动,原驱动`modernc.org/sqlite`,但是速度不行,在github上边没有,就从gitlab搬迁过来,同时修改了驱动名称为sqlite3,这样不用任何修改,就能在goframe框架中使用.

### 使用方法
`_ github.com/logoove/sqlite`

支持goframe的orm框架,也支持原生'sqlx'
`"github.com/jmoiron/sqlx"`
原生使用例子
~~~
var db *sqlx.DB
db, _ = sqlx.Open("sqlite3","./users.db")
rows:=[]Users{}
db.Select(&rows,"SELECT id,name FROM users ORDER BY id")

~~~
### 更新日志
2021-11-3 v1.13.0 新增更多系统编译,支持sqlite 3.36.0