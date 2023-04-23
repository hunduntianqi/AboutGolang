package No2_MongoDB

/*
	MongoDB:
        非关系型数据库, 数据以键值对方式基于磁盘存储
        数据类型单一: 值为JSON文档
        数据库结构: 库->集合->文档
    MongoDB常用命令:
        1. 进入MongoDB命令行: mongo
        2. 查看所有库: show dbs
        3. 切换到指定库: use 库名
        4. 查看当前库中的所有集合: show collections
        5. 查看当前库中的文档: db.集合名.find().pretty()
        6. 统计集合中文档数量: db.集合名.count()
        7. 删除集合: db.集合名.drop()
        8. 删除当前库: db.dropDatabase()
        注意:
            MongoDB无需提前建库建集合, 直接操作即可, 会自动建库建集合
        9. 插入文档:
            a. insert(): db.COLLECTION_NAME.insert(document)
                若插入的数据主键已经存在，则会抛 org.springframework.dao.DuplicateKeyException 异常,
                提示主键重复，不保存当前数据
            b. db.collection.insertOne(document):
                用于向集合插入一个新文档
            c. db.collection.insertMany([document1, document2...]): 用于向集合插入一个多个文档
*/
