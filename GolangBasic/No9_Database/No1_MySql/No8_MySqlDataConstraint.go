package No1_MySql

/*
	数据约束:
        1. 概念&分类:
            概念:
                约束是作用在表中列上的规则, 用于限制加入表中的数据
                约束的存在保证了数据库中数据的正确性, 有效性和完整性
            分类:
                非空约束: 保证列中所有数据不能有null值, 关键字-not null,
                唯一约束: 保证列中所有数据各不相同, 关键字-unique,
                主键约束: 主键是一行数据的唯一标识, 要求非空且唯一, 关键字-primary key,
                检查约束: 保证列中的值满足某一条件, 关键字-check,
                默认约束: 保存数据时, 未指定值采用默认值, 关键字-default,
                外键约束: 外键用来让两个表之间的数据之间建立连接, 保证数据的一致性和完整性, 关键字-foreign key
                注意: MySql不支持检查约束
        2. 外键约束:
            外键用来让两个表之间的数据之间建立连接, 保证数据的一致性和完整性, 关键字-foreign key
            语法:
                添加约束:
                    创建表时添加约束:
                        create table 表名(
                            列名 数据类型,
                            ...
                            constraint 外键名称 foreign key(外键列名) references 主表(主表列名)
                        )
                    建表后添加约束:
                        alter table 表名 add constraint 外键名称 foreign key(外键字段名称) references 主表名称(主表列名称);
                删除约束:
                    alter table 表名 drop foreign key 外键名称;
*/
