package No11_GolangReflection

/*
	Golang反射 ==> 指在运行期间探知对象的类型信息和内存结构, 更新变量,调用他们的方法
	反射的使用场景:
		1. 函数的参数类型是 interface{}, 需要在运行时对原始数据类型进行判断, 针对不同的类型采取不同的处理方式, 如: json.Marshal(v interface{})
		2. 在运行时根据某些条件动态决定调用哪个函数, 比如根据配置文件执行相应的算子函数
	反射的弊端:
		1. 代码难以阅读, 难以维护
		2. 编译期间不能发现类型错误, 覆盖测试难度很大, 有些bug需要到线上运行很长时间才能发现, 可能会造成严重后果
		3. 性能很差, 通常比正常代码慢一到两个数量级, 在对性能要求很高, 或大量反复调用的代码块里建议不要使用反射
	StructField ==> 描述结构体中的一个字段的信息
		type StructField struct {
				Name string // 字段的名字
				PkgPath string // 非导出字段的包路径, 对导出字段该字段为""
				Type Type // 字段的类型
				Tag StructTag // 字段的标签
				Offset uintptr // 字段在结构体中的字节偏移量
				Index []int // 用于Type.FieldByIndex时的索引切片
				Anonymous bool // 是否匿名字段
		}
	StructTag ==> 结构体字段的标签, 一般来说, 标签字符串是空格分隔的一连串 key:"value" 对
		func (tag StructTag) Get(key string) string ==> 获取标签字符串中指定标签的值
		例:
			type Person struct {
				UserId   int    `db:"user_id"`
				Username string `db:"usernmae"`
				Sex      string `db:"sex"`
				Email    string `db:"email"`
			}
			如上字段对应的字符串标签分别为: db:"user_id", db:"usernmae", db:"sex", db:"email"
	Method ==> 代表一个方法
		type Method struct {
			Name string // 方法名
			PkgPath string // 非导出字段的包路径, 对导出字段该字段为""
			Type Type // 方法类型
			Func Value // 方法的值
			Index int // 用于Type.Method的索引
		}
	反射中的两种基本数据类型:
		type Type interface{} 接口:
			获取 reflect.Type 对象 ==> func TypeOf(i interface{}) Type
			reflect.Type接口定义方法:
				1. Kind() Kind ==> 获取变量具体类型
				2. Name() string ==> 获取变量在自身包内的类型名
				3. PkgPath() string ==> 返回变量对应类型的包路径, 即明确指定包的import路径
				4. String() string ==> 返回变量对应类型的字符串表示
				5. Size() uintptr ==> 返回保存对应变量需要的字节数
				6. Implements(u Type) bool ==> 判断对应变量代表的类型是否实现了类型 'u'
				7. Implements(u Type) bool ==> 判断该变量是否可以直接赋值给类型 'u' 对应的变量
				8. ConvertibleTo(u Type) bool ==> 判断该变量是否可以直接转换为类型 'u' 对应的变量
				9. Bits() int ==> 返回该类型的字位数(仅适用于数字类型数据)
				10. Len() int ==> 返回容器类型的长度(数据个数, 仅适用于容器类型数据)
				11. Elem() Type ==> 获取对应变量中存储的指针数据类型, 可使用数据类型为: Array, Chan, Map, Point, Slice
				12. Key() Type ==> 获取一个 Map 类型变量对应 key 的数据类型
				13. ChanDir() ChanDir ==> 获取一个 Chan 类型变量的方向(chan, <-chan, chan<-)
				14. NumField() int ==> 获取一个结构体类型的字段数, 匿名字段算作一个字段
				15. Field(i int) StructField ==> 获取指定结构体中第 i 个字段的数据类型
				16. FieldByIndex(index []int) StructField ==> 返回索引序列指定的嵌套字段的数据类型
				17. FieldByName(name string) (StructField, bool) ==> 获取结构体中字段名为 'name' 的字段数据类型, 会查找匿名字段及其嵌套字段
				18. FieldByNameFunc(match func(string) bool) (StructField, bool) ==> 获取结构体中第一个满足函数 'match' 的字段
				19. IsVariadic() bool ==> 判断函数类型最后一个参数是否为可变参数(...type)
				20. NumIn() int ==> 获取函数类型需要的参数个数
				21. In(i int) Type ==> 获取函数类型第 'i' 个参数需要的数据类型
				22. NumOut() int ==> 获取函数类型需要的返回值个数
				23. Out(i int) Type ==> 获取函数类型第 'i' 个返回值的数据类型
				24. NumMethod() int ==> 获取对应结构体类型成员方法的个数
				25. Method(i int) Method ==> 获取对应类型方法集中的第 'i' 个方法
				26. MethodByName(string) (Method, bool) ==> 根据方法名, 获取对应类型方法集中的方法
		type Value struct{} 类型:
			type Value struct {
				typ *rtype ==> 变量存储数据的数据类型
				ptr unsafe.Pointer ==> 变量中指向原始数据的指针
			}
			获取 reflect.Value 对象 ==> func ValueOf(i interface{}) Value
			reflect.Value成员方法:
				1. func (v Value) IsNil() bool ==> 判断 Value 代表的变量值是否为 nil
				2. func (v Value) Kind() Kind ==> 返回 Value代表的变量存储数据的类型, 如果是类型零值, 则返回 Invalid
				3. func (v Value) Type() Type ==> 将 reflect.Value 转换为 reflect.Type
				4. func (v Value) Convert(t Type) Value ==> 将 Value 的数据类型转换为 't', 并返回一个 reflect.Value 对象
				5. func (v Value) Cap() int ==> 获取 Value 代表的容器变量的容量
				6. func (v Value) Len() int ==> 获取 Value 代表的容器变量的数据长度(个数)
				7. func (v Value) Index(i int) Value ==> 获取 Value 代表的容器变量的第 'i' 个数据
				8. func (v Value) MapIndex(key Value) Value ==> 根据 reflect.Value 的 key, 获取对应值的 reflect.Value 对象
				9. func (v Value) MapKeys() []Value ==> 将Value代表的Map的所有 key 以 []Value 的形式返回
				10. func (v Value) NumField() int ==> 获取 Value 代表的结构体的字段数量
				11. func (v Value) Field(i int) Value ==> 按照索引获取 Value 代表的结构体字段的 reflect.Value 对象
				12. func (v Value) FieldByName(name string) Value ==> 按照字段名获取 Value 代表的结构体字段的 reflect.Value 对象
				13. func (v Value) FieldByNameFunc(match func(string) bool) Value ==> 获取结构体中第一个满足函数 'match' 的字段
				14. func (v Value) Recv() (x Value, ok bool) ==> 从 chan 中获取一个值, 以 reflect.Value 形式
				15. func (v Value) Send(x Value) ==> 向 chan 中存入一个 Value类型的值, 该值数据类型要与 chan 的数据类型一致
				16. func (v Value) Close() ==> 关闭 Value 代表的 chan
				17. func (v Value) Call(in []Value) []Value ==> 根据输入参数 in 调用 v 代表的函数, 返回值为[]Value类型
				18. func (v Value) CallSlice(in []Value) []Value ==> 根据输入参数 in 调用 v 代表的可变参数函数, 返回值为[]Value类型
				19. func (v Value) NumMethod() int ==> 获取 Value 代表的结构体方法集的方法数目
				20. func (v Value) Method(i int) Value ==> 获取对应类型方法集中的第 'i' 个方法, 以 Value 形式
				21. func (v Value) MethodByName(name string) Value ==> 根据方法名, 获取对应类型方法集中的方法
				22. func (v Value) Interface() (i interface{}) ==> 将 Value 类型变量转换为原始数据类型
		反射的其他方法:
			func PtrTo(t Type) Type ==> 获取 reflect.Type 类型变量 't' 对应的指针的 reflect.Type 类型
			func SliceOf(t Type) Type ==> 获取 reflect.Type 类型变量 't' 对应的切片的 reflect.Type 类型
			func MapOf(key, elem Type) Type ==> 获取一个键为 reflect.Type 为 key, 值 reflect.Type 为 elem 的映射类型
			func ChanOf(dir ChanDir, t Type) Type ==> 获取元素 reflect.Type 为t、方向为dir的通道类型
			func New(typ Type) Value ==> 将 reflect.Type 类型变量 'typ' 转换为 reflect.Value 类型
			func (v Value) Type() Type ==> 将 reflect.Value 类型变量 'v' 转换为 reflect.Type 类型
*/
