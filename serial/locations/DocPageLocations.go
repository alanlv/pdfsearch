// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package locations

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DocPageLocations struct {
	_tab flatbuffers.Table
}

func GetRootAsDocPageLocations(buf []byte, offset flatbuffers.UOffsetT) *DocPageLocations {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DocPageLocations{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *DocPageLocations) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DocPageLocations) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DocPageLocations) Locations(obj *TextLocation, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *DocPageLocations) LocationsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func DocPageLocationsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func DocPageLocationsAddLocations(builder *flatbuffers.Builder, locations flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(locations), 0)
}
func DocPageLocationsStartLocationsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DocPageLocationsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
