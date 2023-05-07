package No14_HttpPackage

/*
	http.Request ==> 代表表一个服务端接受到的或者客户端发送出去的HTTP请求
		type Request struct {
			Method string ==> 指定HTTP方法(GET、POST、PUT等), 如果是客户端, ""代表GET
			URL *url.URL ==> 服务端表示被请求的URI, 在客户端表示要访问的URL
			Proto string ==> 接收到的请求的协议版本, 例: "HTTP/1.0" / "HTTP/1.1"
			ProtoMajor int ==> 1
			ProtoMinor int ==> 0
			Header Header ==> 定义请求头
			Body io.ReadCloser ==> 请求的主体
			Close bool ==> 服务端指定是否在回复请求后关闭连接, 在客户端指定是否在发送请求后关闭连接
			Host string ==> 在服务端, Host指定URL会在其上寻找资源的主机, 格式可以是"host:port"; 在客户端, 请求的 Host 字段(可选地)用来重写请求的Host头
			Form url.Values ==> 是解析好的表单数据, 包括 URL 字段的 query 参数和 POST 或 PUT 的表单数据, 本字段只有在调用 ParseForm 后才有效
			PostForm url.Values ==> 是解析好的 POST 或 PUT 的表单数据, 只有在调用 ParseForm 后才有效
			// 其他字段等
		}
		相关函数:
			func NewRequest(method, urlStr string, body io.Reader) (*Request, error) ==> 使用指定的方法、网址和可选的主题创建并返回一个新的请求对象
			func ReadRequest(b *bufio.Reader) (req *Request, err error) ==> 从b读取并解析出一个HTTP请求(本函数主要用在服务端从下层获取请求)
		成员方法:
			func (r *Request) ProtoAtLeast(major, minor int) bool ==> 判断该请求使用的HTTP协议版本是否至少是major.minor
			func (r *Request) UserAgent() string ==> 返回请求中的客户端用户代理信息(请求的User-Agent头)
			func (r *Request) Referer() string ==> 返回请求中的访问来路信息(请求的Referer头)
			func (r *Request) AddCookie(c *Cookie) ==> 向请求中添加一个cookie
			func (r *Request) Cookies() []*Cookie ==> 解析并返回该请求的Cookie头设置的cookie
			func (r *Request) Cookie(name string) (*Cookie, error) ==> 返回请求中名为 name 的 cookie, 如果未找到该 cookie 会返回nil, ErrNoCookie
			func (r *Request) ParseForm() error ==> 解析URL中的查询字符串, 并将解析结果更新到r.Form字段
*/
