package No2_MongoDB

/*
   MongoDB连接:
       标准URL连接语法:
           mongodb://[username:password@]host1[:port1][,host2[:port2],...[,hostN[:portN]]][/[database][?options]]
               mongodb:// ==> 固定格式, 必须指定
               username:password@ ==> 可选项, 登录数据库的用户名和密码
               host1 ==> 必须指定的一个数据库所在服务器地址
               port ==> 指定连接端口, 默认为27017
               /database ==> 指定要连接的数据库, 如果指定username:password@, 连接并验证登录指定数据库, 若不指定, 默认打开 test 数据库
               ?options ==> 连接选项, 如果不使用'/database', 则前面需要加上'/', 所有连接选项都是键值对name=value, 键值对之间通过'&'或';'
           连接实例:
               mongodb://localhost ==> 使用默认端口连接到本地服务器
               mongodb://username:password@host:port/ ==> 使用用户名和密码连接到默认数据库(test)
               mongodb://username:password@host:port/databaseName ==> 使用用户名和密码连接到指定数据库
   MongoDB数据库操作:
       创建数据库: use databaseName
           如果指定数据库不存在, 则会创建一个新的数据库; 数据库已存在, 则会切换到指定数据库
       查看所有数据库: show dbs
           注意: 当新创建的数据库中无任何数据时, 则数据库不会被显示出来, 需要在数据库中插入一定数据后才会被显示
       查看当前正在使用的数据库: db
       删除数据库: db.dropDatabase() ==> 删除当前正在使用的数据库
       退出数据库: exit
   MongoDB集合操作:
       创建集合: db.createCollection(name, options)
           name: 指定集合名称
           options: 可选参数, 指定有关内存大小及索引的选项
               capped ==> 为布尔类型, 如果设置为true, 则会创建一个固定大小的集合, 当达到最大值时, 会自动覆盖最早的文档
                           当指定该参数为true时, 必须同步指定size参数
               size ==> 为数值类型, 为固定集合指定一个最大字节数
               max ==> 数值类型, 指定固定集合中包含文档的最大数量
           注意:
               1. 在插入文档时, MongoDB 首先检查固定集合的 size 字段, 然后检查 max 字段
               2. 在插入文档时, 如果指定集合不存在, 会自动创建一个集合
       集合创建实例:
           db.createCollection("collectionName") ==> 不指定可选参数创建集合
           db.createCollection("collectionName", { capped : true, size : 6142800, max : 10000})
               ==> 指定可选参数创建集合
       查看数据库已有集合: show collections / show tables
       删除集合: db.collection.drop() ==> 如果成功删除集合, 则drop()方法返回true, 否则发挥false
   MongoDB文档操作:
       插入文档:
           db.collection.insert() ==> 若插入的文档数据主键已经存在, 则会抛 org.springframework.dao.DuplicateKeyException 异常,
                                      提示主键重复，不保存当前数据
           db.collection.insertOne() ==> 用于向集合中插入一个新文档, 语法如下:
               db.collection.insertOne(<document>, {writeConcern: <document>})
                   document ==> 要写入的文档
                   writeConcern ==> 写入策略, 默认为1, 即要求确认是写操作, 0则是不要求
           db.collection.insertMany() ==>  用于向集合插入一个或多个文档, 语法如下:
               db.collection.insertMany([<document 1>, <document 2>, ... ],{writeConcern:<document>, ordered:<boolean>})
                   document ==> 要写入的文档
                   writeConcern ==> 写入策略, 默认为1, 即要求确认是写操作, 0则是不要求
                   ordered ==> 指定是否按顺序写入, 默认 true, 按顺序写入
       删除文档:
           db.collectionName.remove(<query>, {justOne:<boolean>, writeConcern:<document>})
               query: 可选参数, 删除文档的条件, 类似 SQL 语句 where 后面的条件参数
               justOne: 可选参数, 如果设为true或1, 则只删除一条符合条件的文档, 如果不设置该参数, 或使用默认值false, 则删除所有匹配条件的文档
               writeConcern: 可选参数, 指定抛出异常的级别
           删除集合中的所有文档: db.collectionName.remove()
           删除集合中的一条文档:db.collectionName.deleteOne(<query>)
           删除集合中的多条文档:db.collectionName.deleteMany(<query>)
       更新文档:
           update() 方法用于更新已存在的文档, 语法格式如下:
               db.collectionName.update(<query>, <update>, {upsert:<boolean>, multi:<boolean>, writeConcern<document>})
                   query: 查询条件, 类似 SQL 语句 where 后面的条件参数
                   update: 要更新数据的对象和一些更新的操作符, 可以理解为 sql update 查询内 set 后面的参数
                   upsert: 可选参数, 如果不存在update的记录, 是否插入objNew, true为插入, 默认是false, 不插入
                   multi: 可选参数, 默认是false, 只更新找到的第一条记录, 如果这个参数为true, 就把按条件查出来多条记录全部更新
                   writeConcern: 可选参数, 指定抛出异常的级别
               例:
                   db.collectionName.update({"keyName":"oldValue"}, {$set:{"keyName":"newValue"}})
                       ==> 将找到的第一条 keyName 为 oldVale 的文档的 keyName 对应值修改为 newValue
                   db.collectionName.update({"keyName":"oldValue"}, {$set:{"keyName":"newValue"}}, {multi:true})
                       ==> 将找到的所有 keyName 为 oldVale 的文档的 keyName 对应值修改为 newValue
       查询文档:
           find() 方法以非结构化的方式来显示所有文档, 语法格式如下:
               db.collectionName.find(query, projection):
                   query: 查询条件, 是可选参数, 类似 SQL 语句 where 后面的条件参数
                   projection: 可选参数, 使用投影操作符指定返回的键, 查询时返回文档中所有键值, 只需省略该参数即可(默认省略)
               db.collectionName.find(query, projection).pretty(): 以易读方式来读取数据
*/
