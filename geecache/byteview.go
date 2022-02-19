package geecache

// 做缓存用 实现Value()接口
type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) ByteSilce() []byte {
	return cloneBytes(v.b)

}

func (v ByteView) String() string {
	return string(v.b)
}

// 只读，所以需要克隆
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
