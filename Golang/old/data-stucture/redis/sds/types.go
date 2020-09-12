package main

type SDSHdr8 struct {
	len  uint8
	cap  uint8
	char []byte
}

type SDSHdr16 struct {
	len  uint16
	cap  uint16
	char []byte
}

type SDSHdr32 struct {
	len  uint32
	cap  uint32
	char []byte
}

type SDSHdr64 struct {
	len  uint64
	cap  uint64
	char []byte
}

type sdsD struct {
	// s8  == 8
	// s16 == 16
	// s32 == 32
	// s64 == 64
	flag uint
	s8   *SDSHdr8
	s16  *SDSHdr16
	s32  *SDSHdr32
	s64  *SDSHdr64
}

type SDS interface {
	Len() uint64
	// sdslen(const sds s): 获取sds字符串长度。
	SetLen(length interface{})
	// sdssetlen(sds s, size_t newlen): 设置sds字符串长度。
	Cap() uint64
	// sdsalloc(const sds s): 获取sds字符串容量。
	SetCap(length interface{})
	// sdssetalloc(sds s, size_t newlen): 设置sds字符串容量。
	Avail() uint64
	// sdsavail(const sds s): 获取sds字符串空余空间（即alloc - len）。
	ReqType() uint64
	// sdsReqType(size_t string_size): 根据字符串数据长度计算所需要的header类型。
	NewSDS()
	// 返回一个新的 SDS 串
	Reset()
	// 回收重置 SDS 串
	Grow(size uint64)
	// 增长定量长度的串长
	byte2string() string
	Print()
	WriteByte(c byte) error
	WriteByte(r rune) (int, error)
	WriteString(s string) (int, error)
}

func NewSDS() *sdsD {
	return &sdsD{
		// default capacity == 10, maxCap == 2^8
		flag: 8,
		s8: &SDSHdr8{
			len:  0,
			cap:  10,
			char: make([]byte, 10),
		},
	}
}
