package No1_MySql

/*
	DDL-操作数据库:
        1. 查询所有数据库:
            show databases;
        2. 创建数据库:
            create database 数据库名称 charset=字符集;
            判断数据库是否存在, 如果不存在, 创建数据库:
                create database if not exists 数据库名称 charset=字符集;
        3. 删除数据库:
            drop database 数据库名称;
            判断数据库是否存在后删除:
                drop database if exists 数据库名;
        4. 使用数据库:
            use 数据库名;
        5. 查看当前使用的数据库:
            select database();
        6. 查看创建数据库的语句:
            show create database 数据库名;
    DDL-操作数据表:
        1. 查询表:
            1.1 查询当前数据库下的所有表:
                show tables;
            1.2 查询表结构:
                desc 表名;
            1.3 查看表的创建语句:
                show create table 数据表名;
        2. 创建表:
            字段以,分隔, 最后一个字段后不需要加逗号
            字段后面必须先写数据类型, 再写约束, 数据类型必须有, 约束可有可无
            创建表的语法如下:
                create table 表名(
                    字段1 数据类型 约束,
                    字段2 数据类型 约束,
                    .....
                    字段n 数据类型 约束
                )
                注意: 最后一行末尾不能加逗号
        3. 删除表:
            3.1 直接删除:
                drop table 表名;
            3.2 判断后删除:
                drop table if exists 表名;
        4. 修改表:
            4.1 修改表名:
                alter table 表名 rename to 新表名;
            4.2 添加字段:
                alter table 表名 add 列名 数据类型 约束;
            4.3 修改字段数据类型:
                alter table 表名 modify 列名 新数据类型 约束;
            4.4 修改字段名和数据类型:
                alter table 表名 change 列名 新列名 数据类型 约束;
            4.5 删除字段:
                alter table 表名 drop 列名;
*/
