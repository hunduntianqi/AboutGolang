package No14_HttpPackage

/*
	http协议通信:
        http请求报文组成:
            1.请求行-定义请求方式, 请求资源路径, 协议版本: GET /index.2022-1-12-html/1.1 \r\n（\r\n实现换行）
            2.请求头-设置网页的相关属性-格式为: 多对key:value, 末尾为\r\n换行
            3.空行-\r\n，换行
            4.请求体(GET方式没有请求体)-想要给服务器传递的数据
        http应答报文组成:
            1.应答行-HTTP/1.1 200 ok(HTTP协议版本, 状态码, 状态描述)+ \r\n
            2.应答头-设置应答数据的相关属性-格式为: 多对key:value+\r\n
            3.空行--\r\n, 换行
            4.应答体-服务器发送给浏览器的数据-各种格式(HTML, css, js, 图片视频等)
        常见http请求方法:
            1.GET: 请求指定的页面信息, 并返回实体主体, 不会修改服务器数据
            2.HEAD: 类似于GET请求, 只不过返回的响应中没有具体内容, 用于获取报头
            3.POST: 向指定资源提交数据进行处理请求(例如提交表单或这上传文件)数据被包含在请求体中, POST请求可能会导致
                    新的资源建立或已有资源的修改
            4.PUT: 从客户端向服务器传送的数据取代指定的文档的内容
            5.DELETE: 请求服务器删除指定的页面
            6.CONNECT: HTTP/1.1协议中预留给能够将链接改为管道方式的代理服务器
            7.OPTIONS: 允许客户端查看服务器的性能
            8.TRACE: 回显服务器收到的请求, 主要用于测试和诊断
            9.PATCH: 是对PUT方法的补充, 用来对已知资源进行局部更新
        请求参数:
            Get请求:
                1. 参数在URL中, 在网页中可见
                2. 参数在网址的'?'后面, 以键值对形式存在: 参数名=参数值
                3. 多个参数之间使用'&'符号连接: 参数1=参数1的值&参数2=参数2的值&...&参数N=参数N的值
            POST请求:
                1. 请求参数存在于请求体中, 在网页中不可见
                2. 通过浏览器开发者工具可以观察请求参数
        请求头:
            爬虫的关键技术之一，服务器通常通过请求头来判定请求是否合法
            请求头列表重要参数：
            1. User-Agent: 代表发起访问的是什么浏览器, 如果不写, 基本会被判定为爬虫,
                           直接拒绝(“httpbin.org/get”可以查看请求内容的网站)
            2. Cookie: 记录了登录信息, 或者上次请求服务端设置的信息, 是常用的反爬判定点
            3. Referer: 表示这次请求是从哪里点过来的, 有的网站不允许直接访问某个网页, 必须是从其他网页点过来才行,
                        这时候需要设置一个Referer值, 模拟是从别的网页点过来的情况
        响应状态码:
            浏览器访问网页时, 服务器向浏览器回复的用于表示请求状态的代码
            常见http status code:
                1.200-请求成功
                2.301-资源（网页等）被永久转移到其他URL
                3.404-请求的资源或网页不存在
                4.500-内部服务器错误
            HTTP状态码有三个数字组成, 第一个数字定义了状态码的类型, 后两个数字表示具体状态, HTTP状态码共分为5种类型:
                1**-信息, 服务器收到请求,需要请求者继续执行操作
                2**-成功, 操作被成功接收并处理
                3**-重定向, 需要进一步的操作以完成请求
                4**-客户端错误, 请求包含语法错误或无法完成请求
                5**-服务器错误, 服务器在处理请求的过程中发生了错误
    http通信流程:
        1. 发起请求:
            a. 请求地址和请求方式, 例如: GET: URL
            b. 请求头 ==> 用于描述请求和发送者的一些信息
            c. 请求参数: 要请求的具体内容信息
        2. 服务器响应:
            a. 服务器返回响应状态码, 判断通信是否成功
            b. 响应头 ==> 描述响应内容的一些信息
            c. 响应内容: 服务器返回的具体数据, 可以是文字, 图片, 超链接, 音视频等
        通信实例:
            客户端请求:
                GET /hello.txt HTTP/1.1
                User-Agent：curl/7.16.3 libcurl/7.16.3 OpenSSL/0.9.71 zlib/1.2.3
                Host:www.example.com
                Accept-Language:en,mi
            服务器响应:
                HTTP/1.1 200 ok
                Date：Mon，27 Jul 2009  12:28:53 GMT
                Server:Apache
                Last-Modified:Wed, 22 Jul 2009 19:15:56 GMT
                ETag:"34aa387-d-1568eb00"
                Accept-Ranges:bytes
                Content-Length:51
                Vary:Accept-Encoding
                Content-Type:text/plain
*/
