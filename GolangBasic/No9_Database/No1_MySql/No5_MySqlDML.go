package No1_MySql

/*
	DML-添加数据:
        1. 给指定列添加数据:
            insert into 表名(列名1, 列名2...) values(值1, 值2, ...);
        2. 给全部列添加数据:
            insert into 表名 values(值1, 值2, ...);
        3. 批量添加数据:
            方式一: insert into 表名(列名1, 列名2, ...) values(值1, 值2, ...), (值1, 值2, ...), ...;
            方式二: insert into 表名 values(值1, 值2, ...), (值1, 值2, ...), ...;
    DML-修改数据:
        1. 指定条件修改表数据:
            update 表名 set 列名1=值1, 列名2=值2, ... where 条件;
            注意: 如果不加条件, 会将所有数据都修改
        2. 不指定条件修改数据:
            update 表名 set 字段1=值1,字段2=值2,....;
    DML-删除数据:
        1. 删除数据:
            delete from 表名 where 条件;
            注意: 如果不加条件, 会将所有数据都删除
        消除重复行: select distinct 字段 from 表名;
*/
