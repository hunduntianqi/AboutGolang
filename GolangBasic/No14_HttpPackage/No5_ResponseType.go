package No14_HttpPackage

/*
	http.Response ==> 代表一个HTTP请求的响应对象
		type Response struct {
			Status string ==> 响应状态, 例: "200 OK"
			StatusCode int ==> 响应状态码, 例: 200
			Proto string ==> http协议版本, 例: "HTTP/1.0"
			ProtoMajor int ==> 例如1
			ProtoMinor int ==> 例如0
			Header Header ==> 头域的键值对
			Body io.ReadCloser ==> 代表回复的主体
			// 其他字段
		}
		func ReadResponse(r *bufio.Reader, req *Request) (*Response, error) ==> 从r读取并返回一个HTTP 回复, req参数是可选的, 指定该回复对应的请求
		成员方法:
			func (*Response) ProtoAtLeast ==> 判断该请求使用的HTTP协议版本是否至少是major.minor
			func (r *Response) Cookies() []*Cookie ==> 解析并返回该回复中的 Set-Cookie 头设置的 cookie
			func (r *Response) Write(w io.Writer) error ==> 以有线格式将回复写入w(用于将回复写入下层TCPConn等)

*/
