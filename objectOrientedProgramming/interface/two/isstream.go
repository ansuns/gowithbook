package two

//定义一个接口
type IStream interface {
	Writer(buf []byte) (n int, err error)
	Read(buf []byte) (n int, err error)
}
