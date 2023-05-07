package No14_HttpPackage

/*
	net / http 包: 提供了HTTP客户端和服务端的实现
		客户端 ==> 使用 Get、Head、Post 和 PostForm 函数发出 HTTP/ HTTPS 请求:
			发送Get请求 ==> resp, err := http.Get(url String) ==> 参数为字符串形式的 url, 返回值为 resp *Response, err error
			发送POST请求 ==> resp, err := http.Post(url, contentType string, body io.Reader)
				==> 返回值为: resp *Response, err error
			PostFrom ==> resp, err := http.PostForm(url string, data url.Values) ==> 返回值为: resp *Response, err error
			所有请求在使用完毕后必须关闭回复主体 ==> resp.Body.Close()
			获取请求到的数据内容 ==> data, err := io.ReadAll(resp.Body), data为一个字节切片
		服务端 ==> type Server ==> Server类型定义了运行HTTP服务端的参数:
			type Server struct {
				Addr string // 设置服务端监听地址与端口
				Handler Handler // 设置处理器对象, 如为nil会调用http.DefaultServeMux
				ReadTimeout time.Duration // 设置读取客户端请求数据的超时时间
				WriteTimeout time.Duration // 设置回复客户端数据操作的超时时间
				MaxHeaderBytes int // 请求的头域最大长度, 如为0则用DefaultMaxHeaderBytes
				TLSConfig *tls.Config // 可选配置, 用于ListenAndServeTLS方法
				TLSNextProto map[string]func(*Server, *tls.Conn, Handler) // 可选配置, 指定一个函数来在一个NPN型协议升级出现时
																			接管TLS连接的所有权
				ConnState func(net.Conn, ConnState) // 可选配置,指定一个回调函数, 该函数会在一个与客户端的连接改变状态时被调用
				ErrorLog *log.Logger // 可选配置,指定一个日志记录器, 用于记录接收连接时的错误和处理器不正常的行为
			}
			创建服务端对象:
				1. 使用包内置对象 ==> http.ListenAndServe("监听地址与端口", nil):
					nil 表示使用采用包变量DefaultServeMux作为处理器
				2. 自定义创建 http.Server对象
					server := &http.Server{} ==> 创建服务端对象
					server.ListenAndServe() ==> 启动服务端
		type Header ==> 代表HTTP头域的键值对
			type Header map[string][]string
			成员方法:
				func (h Header) Get(key string) string ==> 返回键对应的第一个值, 如果键不存在会返回""
				func (h Header) Set(key, value string) ==> 添加键值对到Header对象, 如键已存在则会用新的切片取代旧切片
				func (h Header) Add(key, value string) ==> 添加键值对到Header对象, 如键已存在则会将新的值附加到旧值切片后面
				func (h Header) Del(key string) ==> 删除键值对
				func (h Header) Write(w io.Writer) error ==> 以有线格式将头域写入w
				func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error ==>
					以有线格式将头域写入w, 当 exclude 不为 nil 时, 如果 h 的键在 exclude 中存在且其对应值为真, 该键值对就不会被写入w
		type Cookie ==> 代表一个出现在 HTTP 回复的头域中 Set-Cookie 头的值里或者 HTTP 请求的头域中 Cookie 头的值里的 HTTP cookie
			type Cookie struct {
				Name       string
				Value      string
				Path       string
				Domain     string
				Expires    time.Time
				RawExpires string
				// MaxAge=0表示未设置Max-Age属性
				// MaxAge<0表示立刻删除该cookie，等价于"Max-Age: 0"
				// MaxAge>0表示存在Max-Age属性，单位是秒
				MaxAge   int
				Secure   bool
				HttpOnly bool
				Raw      string
				Unparsed []string // 未解析的“属性-值”对的原始文本
			}
			func (c *Cookie) String() string ==> 返回该cookie的序列化结果, 如果只设置了Name和Value字段, 序列化结果可用于
					HTTP请求的Cookie头或者HTTP回复的Set-Cookie头; 如果设置了其他字段, 序列化结果只能用于HTTP回复的Set-Cookie头
		type CookieJar ==> 管理cookie的存储和在HTTP请求中的使用
			type CookieJar interface {
				SetCookies(u *url.URL, cookies []*Cookie) // SetCookies管理从u的回复中收到的cookie
				Cookies(u *url.URL) []*Cookie // Cookies返回发送请求到u时应使用的cookie
			}
		type Handler ==> 实现了Handler接口的对象可以注册到HTTP服务端，为特定的路径及其子树提供服务
			type Handler interface {
				ServeHTTP(ResponseWriter, *Request)
			}
			成员方法:
				func NotFoundHandler() Handler ==> 返回一个简单的请求处理器, 该处理器会对每个请求都回复"404 page not found"
				func RedirectHandler(url string, code int) Handler ==> 返回一个请求处理器, 该处理器会对每个请求都使用状态码code重定向到网址url
				func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler ==> 返回一个采用指定时间限制的请求处理器
					如果某一次调用耗时超过了时间限制, 该处理器会回复请求状态码503 Service Unavailable, 并将msg作为回复的主体
				func StripPrefix(prefix string, h Handler) Handler ==> 返回一个处理器，该处理器会将请求的 URL 中给定前缀 prefix
					去除后再交由h处理, 会向 URL 中没有给定前缀的请求回复404 page not found
		type HandlerFunc func(ResponseWriter, *Request) ==> 通过类型转换让我们可以将普通的函数作为HTTP处理器使用
			func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) ==> 该方法会调用f(w, r)
		func Head(url string) (resp *Response, err error) ==> 向指定的URL发出一个HEAD请求
		func Get(url string) (resp *Response, err error) ==> 向指定的URL发出一个GET请求
		func Post(url string, bodyType string, body io.Reader) (resp *Response, err error) ==> Post向指定的URL发出一个POST请求
			bodyType ==> POST数据的类型
			body ==> POST数据, 作为请求的主体
		func PostForm(url string, data url.Values) (resp *Response, err error) ==> 向指定的URL发出一个POST请求
			url.Values类型的data会被编码为请求的主体
		func Serve(l net.Listener, handler Handler) error ==> 接收监听器收到的每一个连接, 并为每一个连接创建一个新的服务
			该服务会读取请求，然后调用 handler 回复请求, handler参数一般会设为nil, 此时会使用DefaultServeMux
		func ListenAndServe(addr string, handler Handler) error ==> 监听 TCP 地址 addr, 使用 handler 参数调用 Serve 函数处理接收到的连接
*/

func main() {
}
