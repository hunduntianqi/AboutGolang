package No2_MongoDB

/*
   MongoDB其他操作方法:
       limit()方法 ==> 指定要读取数据的文档条数:
           接受一个数字参数，该参数指定从MongoDB中读取的文档条数
           基本语法: db.COLLECTION_NAME.find().limit(NUMBER)
           如果没有指定number, 则默认显示所有文档数据
       skip()方法 ==> 跳过指定数量的数据:
           db.COLLECTION_NAME.find().limit(NUMBER).skip(NUMBER)
           skip()默认参数为0
       sort()方法 ==> 通过指定键名, 对数据进行排序:
           db.COLLECTION_NAME.find().sort({KEY:1 / -1})
               1 ==> 升序排序; -1 ==> 降序排序
       createIndex()方法 ==> 指定字段创建索引:
           db.collectionName.createIndex({"keyName":1 / -1}, options)
               1 ==> 按升序创建索引; -1 ==> 按降序创建索引
               options: 可选参数
           单个字段创建索引 ==> db.collectionName.createIndex({"keyName":1 / -1})
           多个字段创建索引(复合索引) ==> db.collectionName.createIndex({"keyName1":1 / -1, "keyName2":1 / -1,...})
           查看集合索引 ==> db.collectionName.getIndexes()
           查看集合索引大小 ==> db.collection.totalIndexSize()
           删除集合所有索引 ==> db.collection.dropIndexes()
           删除集合指定索引 ==> db.collection.dropIndexes("indexName")
*/
