package one

//定义一个接口
type ReadWriter interface {
	Read(buf []byte) (n int, err error)
	Writer(buf []byte) (n int, err error)
}
