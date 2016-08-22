// automatically generated by the FlatBuffers compiler, do not modify

package mongo_serialization

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Item struct {
	_tab flatbuffers.Table
}

func GetRootAsItem(buf []byte, offset flatbuffers.UOffsetT) *Item {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Item{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Item) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Item) SeriesId(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Item) SeriesIdLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Item) SeriesIdBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Item) TimestampNanos() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Item) MutateTimestampNanos(n int64) bool {
	return rcv._tab.MutateInt64Slot(6, n)
}

func (rcv *Item) ValueType() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Item) MutateValueType(n int8) bool {
	return rcv._tab.MutateInt8Slot(8, n)
}

func (rcv *Item) LongValue() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Item) MutateLongValue(n int64) bool {
	return rcv._tab.MutateInt64Slot(10, n)
}

func (rcv *Item) DoubleValue() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Item) MutateDoubleValue(n float64) bool {
	return rcv._tab.MutateFloat64Slot(12, n)
}

func ItemStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func ItemAddSeriesId(builder *flatbuffers.Builder, seriesId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(seriesId), 0)
}
func ItemStartSeriesIdVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func ItemAddTimestampNanos(builder *flatbuffers.Builder, timestampNanos int64) {
	builder.PrependInt64Slot(1, timestampNanos, 0)
}
func ItemAddValueType(builder *flatbuffers.Builder, valueType int8) {
	builder.PrependInt8Slot(2, valueType, 0)
}
func ItemAddLongValue(builder *flatbuffers.Builder, longValue int64) {
	builder.PrependInt64Slot(3, longValue, 0)
}
func ItemAddDoubleValue(builder *flatbuffers.Builder, doubleValue float64) {
	builder.PrependFloat64Slot(4, doubleValue, 0.0)
}
func ItemEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}