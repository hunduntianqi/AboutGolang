package No1_MySql

/*
	安装第三方包:
		1. go get github.com/go-sql-driver/mysql ==> mysql驱动
		2. go get github.com/jmoiron/sqlx ==> 基于mysql驱动的封装
	导包:
		_ "github.com/go-sql-driver/mysql"  // 这个包在编写代码时不会用到, 但是按照此格式必须导入, 否则会找不到mysql驱动
		"github.com/jmoiron/sqlx"
	连接MySql:
		database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
		例: database, err := sqlx.Open("mysql", "root:XXXX@tcp(127.0.0.1:3306)/test")
	执行insert语句:
		1. 数据库连接对象调用Exec函数, 执行insert语句:
			database.Exec(update语句)
			例:
				1. database.Exec("insert into 表名(列名1, 列名2...) values (?, ?, ?...)", value1, value2, value3...)
				2. database.Exec("insert into 表名 values (?, ?, ?...)", value1. value2, value3...)
				解释: (?, ?, ?...)相当于占位符, 避免SQL注入, 使用后面的 value1, value2, value3...来传入对应数据
		Exec(insert语句): 该函数会返回一个 Result 对象,
			1. Result对象调用 LastInsertId() 方法可获取最后写入数据的主键值
			2. Result对象调用 RowsAffected() 方法可获取受操作影响的行数
	执行select语句:
		1. 定义结构体类型, 字段结构与要读取数据表字段结构一致并且以大写字母开头, 并通过`db:"数据表字段"`指定映射关系
		2. 创建结构体类型切片, 存储读取的数据
		3. 数据库连接对象调用Select()方法执行select语句
			database.Select(&Slice, select语句, 参数...)
	执行update语句:
		数据库连接对象调用Exec函数, 执行update语句
			database.Exec(update语句)
			例:
				1. database.Exec("update 表名 set 列名1=?, 列名2=?, ... where 条件", 参数列表)
				2. database.Exec("update 表名 set 字段1=?,字段2=?,....", 参数列表)
				解释: ? 相当于占位符
		Exec(update语句): 该函数会返回一个 Result 对象,
			1. Result对象调用 LastInsertId() 方法可获取最后写入数据的主键值
			2. Result对象调用 RowsAffected() 方法可获取受操作影响的行数
	执行delete语句:
		数据库连接对象调用Exec函数, 执行delete语句
			database.Exec(delete语句)
			例:
				1. database.Exec("delete from 表名 where 条件", 参数列表)
				解释: ? 相当于占位符
				注意: 删除操作不过不加条件约束, 会将所有数据全部删除
		Exec(delete语句): 该函数会返回一个 Result 对象,
			1. Result对象调用 LastInsertId() 方法可获取最后写入数据的主键值
			2. Result对象调用 RowsAffected() 方法可获取受操作影响的行数
	事务操作:
		database.Begin() ==> 开启事务, 并返回一个 *Tx 对象
			begin, err := database.Begin()
		*Tx.Commit() ==> 提交事务
			err:= begin.Commit()
		*Tx.Rollback() ==> 回滚事务
			err:= begin.Rollback()
		事务操作流程:
			1. 在执行增删改查前开启事务 ==> database.Begin()
			2. 在出现异常的步骤处回滚事务 ==> *Tx.Rollback()
			3. 所有操作执行完毕后无异常再提交事务 ==> *Tx.Commit()
*/
