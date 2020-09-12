package main

import (
	"errors"
	"fmt"
	"unsafe"
)

func (s *sdsD) Len() uint64 {
	switch s.flag {
	case 8:
		return uint64(s.s8.len)
	case 16:
		return uint64(s.s16.len)
	case 32:
		return uint64(s.s32.len)
	case 64:
		return s.s64.len
	}
	return 0
}

func (s *sdsD) Avail() uint64 {
	switch s.flag {
	case 8:
		return uint64(s.s8.cap - s.s8.len)
	case 16:
		return uint64(s.s16.cap - s.s16.len)
	case 32:
		return uint64(s.s32.cap - s.s32.len)
	case 64:
		return uint64(s.s64.cap - s.s64.len)
	}
}

func (s *sdsD) Print() {
	fmt.Printf("%s\n", s.byte2string())
	return
}

func (s *sdsD) byte2string() string {
	switch s.flag {
	case 8:
		return *(*string)(unsafe.Pointer(&s.s8.char))
	case 16:
		return *(*string)(unsafe.Pointer(&s.s16.char))
	case 32:
		return *(*string)(unsafe.Pointer(&s.s32.char))
	case 64:
		return *(*string)(unsafe.Pointer(&s.s64.char))
	}
	return ""
}

func ReqType(str string) uint64 {
	x := len(str)
	switch {
	case x <= 255:
		return 8
	case x <= 65535:
		return 16
	case x <= 4294967295:
		return 32
	case x > 4294967295:
		return 64
	}
	return 0
}

func (s *sdsD) Cap() uint64 {
	switch s.flag {
	case 8:
		return uint64(s.s8.cap)
	case 16:
		return uint64(s.s16.cap)
	case 32:
		return uint64(s.s32.cap)
	case 64:
		return s.s64.cap
	}
	return 0
}

func (s *sdsD) SetLen(length interface{}) error {
	switch s.flag {
	case 8:
		l, ok := length.(uint8)
		if !ok {
			return errors.New("length argument error")
		}
		if s.s8.cap < l {
			s.s8.cap = l
			s.s8.len = l
			buf := make([]byte, s.s8.cap)
			copy(buf, s.s8.char)
			s.s8.char = buf
		} else {
			s.s8.len = l
		}
	case 16:
		l, ok := length.(uint16)
		if !ok {
			return errors.New("length argument error")
		}
		if s.s16.cap < l {
			s.s16.cap = l
			s.s16.len = l
			buf := make([]byte, s.s16.cap)
			copy(buf, s.s16.char)
			s.s16.char = buf
		} else {
			s.s16.len = l
		}
	case 32:
		l, ok := length.(uint32)
		if !ok {
			return errors.New("length argument error")
		}
		if s.s32.cap < l {
			s.s32.cap = l
			s.s32.len = l
			buf := make([]byte, s.s32.cap)
			copy(buf, s.s32.char)
			s.s32.char = buf
		} else {
			s.s32.len = l
		}
	case 64:
		l, ok := length.(uint64)
		if !ok {
			return errors.New("length argument error")
		}
		if s.s64.cap < l {
			s.s64.cap = l
			s.s64.len = l
			buf := make([]byte, s.s64.cap)
			copy(buf, s.s64.char)
			s.s64.char = buf
		} else {
			s.s64.len = l
		}
	}
	return nil
}

func (s *sdsD) SetCap(length interface{}) error {
	switch s.flag {
	case 8:
		l, ok := length.(uint8)
		if !ok {
			return errors.New("length argument error")
		}
		s.s8.cap = l
		buf := make([]byte, s.s8.cap)
		copy(buf, s.s8.char)
		s.s8.char = buf
	case 16:
		l, ok := length.(uint16)
		if !ok {
			return errors.New("length argument error")
		}
		s.s16.cap = l
		buf := make([]byte, s.s16.cap)
		copy(buf, s.s16.char)
		s.s16.char = buf
	case 32:
		l, ok := length.(uint32)
		if !ok {
			return errors.New("length argument error")
		}
		s.s32.cap = l
		buf := make([]byte, s.s32.cap)
		copy(buf, s.s32.char)
		s.s32.char = buf
	case 64:
		l, ok := length.(uint64)
		if !ok {
			return errors.New("length argument error")
		}
		s.s64.cap = l
		buf := make([]byte, s.s64.cap)
		copy(buf, s.s64.char)
		s.s64.char = buf
	}
	return nil
}

func (s *sdsD) Reset() {
	switch s.flag {
	case 8:
		s.s8.char = s.s8.char[:10]
		s.s8.len = 0
		s.s8.cap = 10
	case 16:
		s.s16.char = s.s8.char[:10]
		s.s16.len = 0
		s.s16.cap = 10
	case 32:
		s.s32.char = s.s8.char[:10]
		s.s32.len = 0
		s.s32.cap = 10
	case 64:
		s.s64.char = s.s8.char[:10]
		s.s64.len = 0
		s.s64.cap = 10
	}
	return
}

func WriteByte(c byte) error {
	// TODO...
	return nil
}
func WriteByte(r rune) (int, error) {
	// TODO...
	return -1, nil
}

func (s *sdsD) WriteString(str string) (int, error) {
	// TODO...
	if s.Avail() < len(str) {
		return -1, errors.New("Overflow!")
	}
	tmp := (*stringStruct)(unsafe.Pointer(&str))
	ret := slice{array: unsafe.Pointer(tmp.str), len: tmp.len, cap: tmp.len}
	b := *(*[]byte)(unsafe.Pointer(&ret))
	var t []byte

	switch s.flag {
	case 8:
		t = s.s8.char
	case 16:
		t = s.s16.char
	case 32:
		t = s.s32.char
	case 64:
		t = s.s64.char
	}
	for i := 0; i < len(b); i++ {
		t[i+s.len] = b[i]
		// TODO change length
	}
	return len(b), nil
}

func (s *sdsD) Grow(size uint64) error {
	buf := make([]byte, size)
	switch s.flag {
	case 8:
		sum := uint8(size) + uint8(len(s.s8.char))
		if sum < uint8(size) || sum < uint8(len(s.s8.char)) {
			return errors.New("overflow...")
			// TODO extend to new char array
		}
		s.s8.char = append(s.s8.char, buf...)
	case 16:
		sum := uint16(size) + uint16(len(s.s16.char))
		if sum < uint16(size) || sum < uint16(len(s.s16.char)) {
			return errors.New("overflow...")
			// TODO extend to new char array
		}
		s.s16.char = append(s.s16.char, buf...)
	case 32:
		sum := uint32(size) + uint32(len(s.s32.char))
		if sum < uint32(size) || sum < uint32(len(s.s32.char)) {
			return errors.New("overflow...")
			// TODO extend to new char array
		}
		s.s32.char = append(s.s32.char, buf...)
	case 64:
		sum := uint64(size) + uint64(len(s.s64.char))
		if sum < uint64(size) || sum < uint64(len(s.s64.char)) {
			return errors.New("overflow...")
			// TODO extend to new char array
		}
		s.s64.char = append(s.s64.char, buf...)
	}
	return nil
}

func main() {
	s := NewSDS()
}
