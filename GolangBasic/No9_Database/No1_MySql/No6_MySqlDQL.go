package No1_MySql

/*
	DQL-基础查询:
        1. 查询所有数据:
            select * from 表名;
        2. 查询对应列数据:
            select 字段1, 字段2, ... from 表名;
        3. 去除重复数据:
            select distinct 查询字段 from 表名;
        4. 查询时为字段起别名:
            select 字段 as 别名 from 表名;
    DQL-条件查询(where):
        1. 条件查询语法:
            select 字段列表 from 表名 where 条件列表;
        2. 条件:
            >: 大于
                select 字段 from 表名 where 字段 > 数值;
            <: 小于
                select 字段 from 表名 where 字段 < 数值;
            >=: 大于或等于
                select 字段 from 表名 where 字段 >= 数值;
            <=: 小于或等于
                select 字段 from 表名 where 字段 <= 数值;
            =: 等于
                select 字段 from 表名 where 字段 = 数值;
            <> 或 !=: 不等于
                select 字段 from 表名 where 字段 != 数值;
                select 字段 from 表名 where 字段 <> 数值;
            between...and...: 在某个范围之内
                select 字段 from 表名 where 条件字段 between 下限数值 and 上限数值;
            in(...): 多选一
                select 字段 from 表名 where 条件字段 in (条件1, 条件2, 条件3...);
            like 占位符: 模糊查询, _-单个任意字符, %-多个任意字符
                例:
                    1. 查询姓马的学员信息:
                        select * from stu where name like '马%';
                    2. 查询第二个字是花的学员信息:
                        select * from stu where name like '_花%';
                    3. 查询名字中包含德的学员信息:
                        select * from stu where name like '%德%';
            is null: 是null
                select 字段 from 表名 where 字段 is null;
            is not null: 不是null
                select 字段 from 表名 where 字段 is not null;
            and 或 &&: 并且
                select 字段 from 表名 where 条件1 and 条件2 and ...;
            or 或 ||: 或者
                select 字段 from 表名 where 条件1 or 条件2 or ...;
            not 或 !: 非, 不是
    DQL-排序查询(order by):
        1. 排序查询语法:
            select 字段 from 表名 order by 排序字段名1 排序方式, 排序字段名2 排序方式(升序-asc/降序desc), ...;
        2. 排序方式:
            ASC: 升序排列(默认值)
            DESC: 降序排列
        注意: 如果有多个排序条件, 当前面的条件值一样时, 才会根据第二条件进行排序
    DQL-分组查询(group by):
        1. 聚合函数:
            将一列数据作为一个整体, 进行纵向计算
        2. 聚合函数分类:
            count(列名): 统计数量(一般选用不为null的列)
            max(列名): 最大值
            min(列名): 最小值
            sum(列名): 求和
            avg(列名): 平均值
        3. 聚合函数语法:
            select 聚合函数1(列名), 聚合函数2(列名),... from 表;
            注意: null值不参与所有聚合函数的运算
        4. 分组查询语法:
            select 字段列表 from 表名 where 分组前条件 group by 分组字段名 having 分组后条件;
            注意: 分组之后查询的字段为聚合函数和分组字段, 查询其他字段无任何意义
            where和having的区别:
                4.1 执行时机不一样: where是分组之前进行限定, 不满足where条件则不参与分组, 而having是分组之后对结果进行过滤
                4.2 可判断的条件不一样: where不能对聚合函数进行判断, having可以
            执行顺序: where > 聚合函数 > having
    DQL-分页查询:
        1. 分页查询语法:
            a. 不带条件查询: select 字段 from 表名 limit 起始索引, 查询数量;
            b. 带条件查询: select 字段 from 表名 where 条件1,条件2,..... limit 起始索引, 查询数量
            起始索引: 从0开始
            计算公式: 起始索引等于 = (当前页码 - 1) * 每页显示数据数量
            tips:
                分页查询limit是MySql数据库的方言
                Oracle数据库分页查询使用rownumber
                SQL Server分页查询使用top
*/
