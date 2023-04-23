package No7_FileOperation

/*
	文件操作相关API:
		文件操作相关API位于os包下
		获取文件对象:
			1. func Create(path string) (file *File, err Error): 根据路径创建新的文件对象, 如果文件已存在则会清空文件数据,
																 默认权限是0666
			2. func NewFile(fd uintptr, path string) *File: 根据文件描述符创建文件对象, 文件描述符是一个非负整数, 本质上是一个索引值
				文件描述符为一个非负整数, 一般从3开始
			3. func Open(path string) (file *File, err Error): 根据指定路径以只读方式打开文件对象
			4. func OpenFile(path string, flag int, perm uint32) (file *File, err Error): 指定路径打开文件对象
				path ==> 目标文件路径
				flag ==> 文件读写方式
				perm ==> 操作权限, 通常新建文件默认权限为 0666, 新建文件夹默认权限为 0777
				flag可取值对应常量:
					os.O_WRONLY ==> 只写
					os.O_CREATE ==> 创建文件对象, 如果文件已存在会报错
					os.O_RDONLY ==> 只读
					os.O_RDWR ==> 可读可写
					os.O_TRUNC ==> 清空文件数据
					os.O_APPEND ==> 追加写入
			5. func (f *File) Readdir(n int) (fi []FileInfo, err error): 读取文件夹中指定 n 个文件对象的FileInfo, 返回一个切片
			6. func (f *File) Readdirnames(n int) (names []string, err error): 读取文件夹中指定 n 个文件对象的文件名, 返回一个切片
		关闭文件对象:
			func (file *File) Close() error: 关闭文件对象并返回一个错误对象
		获取文件信息:
			func Stat(name string) (fi FileInfo, err error): 获取指定文件对象信息, 返回一个 FileInfo 对象
				FileInfo ==> 用来描述一个文件对象, 具体参数如下:
					type FileInfo interface {
						Name() string // 文件的名字
						Size() int64 // 普通文件返回值表示其大小(字节), 其他文件的返回值含义各系统不同
						Mode() FileMode // 文件的模式位
						ModTime() time.Time // 文件的修改时间
						IsDir() bool // 等价于Mode().IsDir()
						Sys() interface{} // 底层数据来源（可以返回nil）
						}
				根据FileInfo判断文件是否为文件或文件夹 ==> fileInfo.Mode().IsDir / fileInfo.Mode().IsRegular, 返回值为bool类型
		删除文件对象:
			func Remove(path string) Error: 根据文件路径删除指定文件对象
			func RemoveAll(path string) error: 根据文件夹路径删除文件对象, 并删除文件夹中的内容;
												如果path指定的对象不存在, 会返回nil而不返回错误
		判断文件是否存在:
			1. func IsExist(err error) bool: 判断一个错误对象是否表示文件存在
			2. func IsNotExist(err error) bool: 判断一个错误对象是否表示文件不存在
		获取当前工作目录绝对路径 ==> func Getwd() (dir string, err error)
		修改当前工作目录路径为指定路径 ==> func Chdir(dir string) error
		创建文件夹:
			func Mkdir(name string, perm FileMode) error: 创建个单级文件夹对象
			func MkdirAll(path string, perm FileMode) error: 创建多级文件夹对象
		文件重命名 / 移动文件 ==> func Rename(oldpath, newpath string) error
*/
