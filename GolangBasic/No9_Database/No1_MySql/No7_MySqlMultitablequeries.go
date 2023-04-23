package No1_MySql

/*
	多表查询:
        从多张表查询数据
        1. 连接查询:
            内连接:
                相当于查询两张表交集的数据
                查询语法:
                    隐式内连接:
                        select 字段 from 表名1, 表名2, ... where 条件;
                    显式内连接:
                        select 字段 from 表1 inner join 表2 on 条件;
                    注意: on是连接条件, where是连接后筛选条件
            外连接:
                左外连接: 相当于查询A表所有数据和交集部分数据
                    语法: select 字段列表 from 表1 left outer join 表2 on 条件;
                    注意: 对于右表中不存在的部分使用null填充
                右外连接: 相当于查询B表所有数据和交集部分数据
                    语法: select 字段列表 from 表1 right outer join 表2 on 条件;
                    注意:对于左表中不存在的部分使用null填充
            自连接:
                数据在同一张表中,通过起不同的别名来达到想要的查询数据结果
                语法:select * from 表名 as 别名1 inner join 表名 as 别名2 on 别名1.条件字段=别名2.条件字段;
        2. 子查询:
            指查询中嵌套查询, 称嵌套查询为子查询
            子查询根据查询结果不同, 作用不同:
                单行单列:
                    作为条件值, 可以使用算数运算符进行条件判断
                    select 字段 from 表 where 字段名 = (子查询语句);
                多行单列:
                    作为条件值, 使用in 等关键字进行条件判断
                    select 字段 from 表 where 字段名 in (子查询);
                多行多列:
                    作为虚拟表
                    select 字段 from (子查询) where 条件;
*/
