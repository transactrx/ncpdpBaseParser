package ncpdpmodules

import (
	"github.com/transactrx/ncpdpBaseParser/pkg/orderedmap"
	"strconv"
	"strings"
)

type Group struct {
	Segments *orderedmap.OrderedMap[any]
}

//func (msg *Group) GetCurrencyFieldValue(segmentId, fieldID string, returnNull bool) (float64, error) {
//	fields := msg.GetField(segmentId, fieldID)
//	if len(fields) == 0 {
//		return float64(0.0), nil
//	}
//	val, err := parseNCPDPCurrencyString(fields, returnNull)
//
//	return *val, err
//
//}
//
//func (msg *Group) GetFieldValue(segmentId, fieldID string) string {
//	fields := msg.GetField(segmentId, fieldID)
//	if len(fields) == 0 {
//		return ""
//	}
//	return fields
//
//}
//func (msg *Group) GetField(segmentId, fieldID string) string {
//	segment := msg.Segments[segmentId]
//	return segment.FieldGroups[fieldID]
//
//}

func parseNCPDPCurrencyString(value string, returnNull bool) (*float64, error) {

	if value == "" {
		if returnNull {
			return nil, nil
		} else {
			t := float64(0.0)
			return &t, nil
		}

	}

	lastChar := string(value[len(value)-1])
	v := ""
	negative := false

	switch {

	case lastChar == "}" || lastChar == "{":
		negative = lastChar == "}"
		v = "0"
		break
	case lastChar == "A" || lastChar == "J":
		negative = lastChar == "J"
		v = "1"
		break
	case lastChar == "B" || lastChar == "K":
		negative = lastChar == "K"
		v = "2"
		break
	case lastChar == "C" || lastChar == "L":
		negative = lastChar == "L"
		v = "3"
		break
	case lastChar == "D" || lastChar == "M":
		negative = lastChar == "M"
		v = "4"
		break

	case lastChar == "N" || lastChar == "E":
		negative = lastChar == "N"
		v = "5"
		break

	case lastChar == "F" || lastChar == "O":
		negative = lastChar == "O"
		v = "6"
		break
	case lastChar == "G" || lastChar == "P":
		negative = lastChar == "P"
		v = "7"
		break
	case lastChar == "H" || lastChar == "Q":
		negative = lastChar == "Q"
		v = "8"
		break
	case lastChar == "I" || lastChar == "R":
		negative = lastChar == "R"
		v = "9"
		break
	}
	value = strings.Replace(value, lastChar, v, 1)

	tResult, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, err
	}
	result := float64(tResult / 100)
	if negative {
		result = float64(-tResult / 100)
	}

	return &result, nil
}
