package No1_MySql

/*
	索引:
        索引在MySQL中也叫作"键", 是一个特殊的文件, 保存着数据表里所有记录的位置信息
        索引的作用: 加快数据查询速度
        应用场景:当数据库中的数据量很大时, 查找数据会变得很慢, 我们就可以通过索引来提高数据库查询效率
        查询表中已有的索引: show index from 表名;
        开启时间监测: set profiling=1;
        给字段创建索引: alter table 表名 add index(字段);
        开启时间监测后, 查看每条SQL语句的执行时间: show profiles;
*/
