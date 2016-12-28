// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs shared.go

package shm

type sharedBlock struct {
	Next      uint64
	Prev      uint64
	DoneRead  uint64
	DoneWrite uint64
	Size      uint64
	Flags     [24]uint8
}

type sharedMem struct {
	BlockCount uint64
	BlockSize  uint64
	ReadStart  uint64
	ReadEnd    uint64
	WriteStart uint64
	WriteEnd   uint64
	SemSignal  [16]byte
	SemAvail   [16]byte
	Flags      [48]uint8
}

const (
	sharedHeaderSize = 0x80
	blockHeaderSize  = 0x40
	blockFlagsSize   = len(sharedBlock{}.Flags)
	headerFlagsSize  = len(sharedMem{}.Flags)
)