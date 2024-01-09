package bitmap

import "errors"

const ByteSize = 8

type BitMap struct {
	Bit []byte
}

func NewBitMap(lens uint64) *BitMap {
	return &BitMap{
		Bit: make([]byte, lens/8+1),
	}
}

func (b *BitMap) Set(v uint64) error {
	if v/ByteSize > uint64(len(b.Bit)) {
		return errors.New("out of range")
	}
	index, offset := v/ByteSize, v%ByteSize
	b.Bit[index] |= 1 << byte(offset)
	return nil
}
func (b *BitMap) IsExist(pos uint64) bool {
	if pos/ByteSize > uint64(len(b.Bit)) {
		return false
	}
	index, offset := pos/ByteSize, pos%ByteSize
	return b.Bit[index]&(1<<offset) != 0
}
func (b *BitMap) Del(v uint64) error {
	if v/ByteSize > uint64(len(b.Bit)) {
		return errors.New("out of range")
	}
	index, offset := v/ByteSize, v%ByteSize
	// 1 左移多位后按位取反，再和原来的值与
	b.Bit[index] &= ^(1 << byte(offset))
	return nil
}
