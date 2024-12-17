package main

import (
	"fmt"
	"math"
	"math/big"
)

type HyperLogLog struct {
	alpha        float64
	numIndexBits int
	numBuckets   int
	register     []int32
}

func NewHyperLogLog(numIndexBits int) *HyperLogLog {
	return &HyperLogLog{
		alpha:        0.79402,
		numIndexBits: numIndexBits,
		numBuckets:   int(math.Pow(2.0, float64(numIndexBits))),
		register:     make([]int32, int(math.Pow(2.0, float64(numIndexBits)))),
	}
}

func (h *HyperLogLog) findMsbIndex(hashValue int32) int {
	if hashValue == 0 {
		return -1
	}
	index := 0
	for hashValue > 1 {
		hashValue >>= 1
		index++
	}
	return index
}

/*
this function will find the first `b` bits starting from the MSB.
We do MSB, because if we don't it is very highly likely that the index will be
0 for most cases.
*/
func (h *HyperLogLog) getFirstBBits(hashValue int32) (int32, int) {
	i64Hash, msbIndex := int32(hashValue), h.findMsbIndex(hashValue)
	return i64Hash >> (int32(msbIndex) + 1 - int32(h.numIndexBits)), msbIndex
}

func (h *HyperLogLog) getOtherBits(hashValue int32, msbIndex int) int32 {
	mask := big.NewInt(0)
	for i := 0; i < (msbIndex + 1 - h.numIndexBits); i++ {
		mask.SetBit(mask, i, 1)
	}
	return int32(mask.Uint64()) & hashValue
}

func (h *HyperLogLog) getMsbRelativeToIndex(hashValue, msbIndex, maskedValue int32) int32 {
	mIndex := h.findMsbIndex(maskedValue)
	return msbIndex - int32(h.numIndexBits) - int32(mIndex) + 1
}

func (h *HyperLogLog) Ingest(hashValue int32) {
	registerIndex, msbIndex := h.getFirstBBits(hashValue)
	value := h.getOtherBits(hashValue, msbIndex)
	h.register[registerIndex] = max(h.register[registerIndex], h.getMsbRelativeToIndex(hashValue, int32(msbIndex), value))
}

func (h *HyperLogLog) Cardinality() float64 {
	z := 0.0
	for i := 0; i < len(h.register); i++ {
		z += math.Pow(2, -float64(h.register[i]))
	}
	z = math.Pow(z, -1)
	return h.alpha * float64(h.numBuckets) * float64(h.numBuckets) * z
}

func (h *HyperLogLog) Print() {
	fmt.Println("alpha:", h.alpha)
	fmt.Println("indexBits:", h.numIndexBits)
	fmt.Println("numBuckets:", h.numBuckets)
	fmt.Println("registers:", h.register)
}
