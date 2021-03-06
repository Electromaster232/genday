package genday

import (
	"bytes"
	"fmt"
)

type Listing struct {
	Timeslot Timeslot
	Name string
}

func NewListing(timeslot Timeslot, name string) *Listing {
	return &Listing{
		Timeslot: timeslot,
		Name: name,
	}
}

func (l *Listing) ZoneCorrectedSlot(timezone int) int {
	ts := int(l.Timeslot)
	ts -= (timezone - 1) * 2

	if ts > 48 {
		ts -= 48
	} else if ts <= 0 {
		ts += 48
	}

	return ts
}

func (l *Listing) ToBytes(c *Curday) []byte {
	var out bytes.Buffer

	// {str(timeslot)}\x001\x0034\x000\x000\x00{name}\x00

	out.WriteString(fmt.Sprintf("%d", l.ZoneCorrectedSlot(c.Timezone)))
	out.WriteByte(0x00)
	out.WriteString("1")
	out.WriteByte(0x00)
	out.WriteString("34")
	out.WriteByte(0x00)
	out.WriteString("0")
	out.WriteByte(0x00)
	out.WriteString("0")
	out.WriteByte(0x00)
	out.WriteString(l.Name)
	out.WriteByte(0x00)

	return out.Bytes()
}