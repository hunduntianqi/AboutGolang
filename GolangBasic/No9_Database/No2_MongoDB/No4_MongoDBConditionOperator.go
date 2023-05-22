package No2_MongoDB

/*
   MongoDB条件操作符:
       等于: {<keyName>:<value>} ==> db.collectionName.find("by":"菜鸟教程").pretty() ==> where by="菜鸟教程"
       小于: {<keyName>:{$lt:<value>}} ==> db.collectionName.find("likes":{$lt:<50>}).pretty() ==> where likes < 50
       小于或等于: {<keyName>:{$lte:<value>}} ==> db.collectionName.find("likes":{$lte:<50>}).pretty() ==> where likes <= 50
       大于: {<keyName>:{$gt:<value>}} ==> db.collectionName.find("likes":{$gt:<50>}).pretty() ==> where likes > 50
       大于或等于: {<keyName>:{$gte:<value>}} ==> db.collectionName.find("likes":{$gte:<50>}).pretty() ==> where likes >= 50
       不等于: {<keyName>:{$ne:<value>}} ==> db.collectionName.find("likes":{$ne:<50>}).pretty() ==> where likes != 50
       AND: {keyName1:value1, keyName2: value2, ...} ==> db.collectionName.find({keyName1:value1, keyName2: value2}).pretty()
               ==> where keyName1=value1 AND keyName2=value2
       OR: {$or[{keyName1:value1, keyName2: value2, ...}]} ==> db.collectionName.find({$or:[{keyName1:value1, keyName2: value2}]}).pretty()
               ==> where keyName1=value1 OR keyName2=value2
   $type操作符:
       $type操作符是基于BSON类型来检索集合中匹配的数据类型, 并返回结果
           语法格式: {<keyName>: {$type: <dataType>}} ==> 查找指定键名为指定数据类型的文档
       例: db.collectionName.find({"title" : {$type : 'string'}}) / db.collectionName.find({"title" : {$type : 2}})
               ==> 用于查找集合中 title 值为字符串类型的文档
*/
