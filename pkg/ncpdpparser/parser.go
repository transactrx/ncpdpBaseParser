package ncpdpparser

import (
	"errors"
	"github.com/transactrx/ncpdpBaseParser/pkg/ncpdpmodules"
	"strings"
	"sync"
)

var segmentSeparator = "\u001CAM"
var fieldSeparator = "\u001C"

var mu sync.Mutex

func New(ncpdp string) (*ncpdpmodules.NCPDPMessage, error) {
	if ncpdp == "" {
		return nil, errors.New("NCPDP string is empty")

	}
	msg := ncpdpmodules.NCPDPMessage{}
	msg.Segments = make(map[string]ncpdpmodules.Segment)
	segments := strings.Split(ncpdp, segmentSeparator)

	for i, val := range segments {
		if i == 0 {
			continue
		}
		msg.Segments[val[:2]] = parseSegment(val[2:])
	}
	return &msg, nil
}

func parseSegment(s string) ncpdpmodules.Segment {
	segment := ncpdpmodules.Segment{}
	segment.Fields = make(map[string][]string)
	fields := strings.Split(s, fieldSeparator)
	for i, val := range fields {

		if i == 0 {
			continue
		}
		//if already exists, append
		fieldVal := val[2:]
		fieldId := val[:2]
		if _, ok := segment.Fields[fieldId]; ok {
			segment.Fields[val[:2]] = append(segment.Fields[fieldId], fieldVal)
			continue
		}
		segment.Fields[fieldId] = []string{fieldVal}
	}
	return segment
}