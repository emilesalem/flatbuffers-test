// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package b

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type B struct {
	_tab flatbuffers.Table
}

func GetRootAsB(buf []byte, offset flatbuffers.UOffsetT) *B {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &B{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *B) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *B) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *B) Field1() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func BStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func BAddField1(builder *flatbuffers.Builder, field1 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(field1), 0)
}
func BEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}