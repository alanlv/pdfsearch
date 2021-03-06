// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package pdf_index

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PdfIndex struct {
	_tab flatbuffers.Table
}

func GetRootAsPdfIndex(buf []byte, offset flatbuffers.UOffsetT) *PdfIndex {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PdfIndex{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *PdfIndex) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PdfIndex) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PdfIndex) NumFiles() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PdfIndex) MutateNumFiles(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *PdfIndex) NumPages() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PdfIndex) MutateNumPages(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *PdfIndex) Index(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *PdfIndex) IndexLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *PdfIndex) Hipd(obj *HashIndexPathDoc, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *PdfIndex) HipdLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func PdfIndexStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func PdfIndexAddNumFiles(builder *flatbuffers.Builder, numFiles uint32) {
	builder.PrependUint32Slot(0, numFiles, 0)
}
func PdfIndexAddNumPages(builder *flatbuffers.Builder, numPages uint32) {
	builder.PrependUint32Slot(1, numPages, 0)
}
func PdfIndexAddIndex(builder *flatbuffers.Builder, index flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(index), 0)
}
func PdfIndexStartIndexVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func PdfIndexAddHipd(builder *flatbuffers.Builder, hipd flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(hipd), 0)
}
func PdfIndexStartHipdVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func PdfIndexEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
