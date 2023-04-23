package No1_MySql

/*
	数据库设计:
        1. 软件的研发步骤:
            需求分析-设计-编码-测试-安装部署
        2. 数据库设计概念:
            指建立数据库中的表结构以及表与表之间关联关系的过程
            确认有哪些表, 表中有哪些字段, 表和表之间有什么关系等
        3. 数据库设计步骤:
            3.1 需求分析(数据是什么, 数据具有哪些属性, 数据与属性的特点)
            3.2 逻辑分析(通过ER图对数据库进行逻辑建模, 不需要考虑所选用的数据库管理系统)
            3.3 物理设计(根据数据库自身的特点把逻辑设计转换为物理设计)
            3.4 维护设计(1. 对新的需求进行建表, 2. 表优化)
    表关系:
        1. 一对一:
            如用户和用户详情
            一对一关系多用于表拆分, 将一个实体中经常使用的字段放在一张表, 不经常使用的字段放在另一张表, 用于提升查询性能
            实现方式: 在任意一方加入外键, 关联另一方的主键, 并设置外键为唯一(unique)
        2. 一对多(多对一):
            如部门和员工--一个部门对应多个员工, 一个员工对应一个部门
            实现方式:
                在多的一方建立外键, 指向一的一方主键
        3. 多对多:
            如商品和订单--一个商品包含多个订单, 一个订单包含多个商品
            实现方式:
                建立第三张中间表, 中间表至少包含两个外键, 分别关联两方的主键
*/
