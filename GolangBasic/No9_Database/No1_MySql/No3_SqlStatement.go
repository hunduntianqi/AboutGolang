package No1_MySql

/*
	SQL简介:
        结构化查询语言, 是一门用来操作关系型数据库的编程语言
        定义了操作所有关系型数据库的统一标准
        对于同一个需求, 每一种数据库操作的方式可能会存在一些不一样的地方, 被称为"方言"
    SQL通用语法:
        1. SQL语句可以单行或多行书写, 以分号结尾
        2. MySQL数据库的SQL语句不区分大小写, 关键字建议使用大写
        3. 注释:
            单行注释: -- 注释内容 或 #注释内容(MySQL特有)
            多行注释: /* 注释  */
/*	SQL分类:
	1. DDL(Data Definition Language): 数据定义语言, 用来定义数据库对象-数据库, 表, 列等
	2. DML(Data Manipulation Language): 数据操作语言, 用来对数据库中表的数据进行增删改
	3. DQL(Data Query Language): 数据查询语言, 用来查询数据库中表的记录(数据)
	4. DCL(Data Control Language): 数据控制语言, 用来定义数据库的访问权限和安全级别, 及创建用户
*/
