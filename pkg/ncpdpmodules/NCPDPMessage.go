package ncpdpmodules

import (
	"github.com/transactrx/ncpdpBaseParser/pkg/orderedmap"
)

type NCPDPMessage struct {
	Groups []*Group
}

func (msg *NCPDPMessage) GetHeaderFieldAsString(fieldId string) *string {
	seg, ok := msg.Groups[0].Segments.GetAsValue("header")
	if !ok {
		return nil
	}

	value, ok := seg.(*orderedmap.OrderedMap[string]).GetAsValue(fieldId)
	if ok {
		return &value
	}
	return nil
}

func (msg *NCPDPMessage) GetResponsePatientFieldAsString(fieldId string) *string {
	return msg.GetFieldValueAsString(0, "29", fieldId)
}

func (msg *NCPDPMessage) GetResponseInsuranceFieldAsString(fieldId string) *string {
	return msg.GetFieldValueAsString(0, "25", fieldId)
}

func (msg *NCPDPMessage) GetResponseMessageFieldAsString(fieldId string) *string {
	return msg.GetFieldValueAsString(0, "20", fieldId)
}

func (msg *NCPDPMessage) GetRequestPatientFieldAsString(fieldId string) *string {
	return msg.GetFieldValueAsString(0, "01", fieldId)
}

func (msg *NCPDPMessage) GetRequestInsuranceFieldAsString(fieldId string) *string {
	return msg.GetFieldValueAsString(0, "04", fieldId)
}

func (msg *NCPDPMessage) GetFieldValueAsString(groupNumber int, segment string, fieldId string) *string {
	seg, ok := msg.Groups[groupNumber].Segments.GetAsValue(segment)
	if !ok {
		return nil
	}
	value, ok := seg.(*orderedmap.OrderedMap[any]).GetAsValue(fieldId)
	if ok {
		switch v := value.(type) {
		case []*orderedmap.OrderedMap[string]:
			return nil
		case string:
			return &v
		default:
			return nil
		}
	}

	return nil
}

func (msg *NCPDPMessage) GetResponseBenefitStageInfoSegment(statusIndex int) []*orderedmap.OrderedMap[string] {
	return msg.GetFieldValueAsGroup(statusIndex, "23", "MU")
}

func (msg *NCPDPMessage) GetResponseStatusAdditionalInfo(statusIndex int) []*orderedmap.OrderedMap[string] {
	return msg.GetFieldValueAsGroup(statusIndex, "21", "UF")
}

func (msg *NCPDPMessage) GetResponseStatusRejectCodes(statusIndex int) []*orderedmap.OrderedMap[string] {
	return msg.GetFieldValueAsGroup(statusIndex, "21", "FA")
}

func (msg *NCPDPMessage) GetRequestClinicalInfo(claimIndex int) []*orderedmap.OrderedMap[string] {
	return msg.GetFieldValueAsGroup(claimIndex, "13", "VE")
}

func (msg *NCPDPMessage) GetRequestCompoundInfoFieldAsString(claimIndex int, fieldId string) *string {
	return msg.GetFieldValueAsString(claimIndex, "10", fieldId)
}

func (msg *NCPDPMessage) GetRequestCompoundDtoInfo(claimIndex int) []*orderedmap.OrderedMap[string] {
	return msg.GetFieldValueAsGroup(claimIndex, "10", "EC")
}

func (msg *NCPDPMessage) GetRequestDURFieldAsString(claimIndex int) []*orderedmap.OrderedMap[string] {
	return msg.GetFieldValueAsGroup(claimIndex, "08", "7E")
}

func (msg *NCPDPMessage) GetFieldValueAsGroup(groupNumber int, segment string, fieldId string) []*orderedmap.OrderedMap[string] {
	seg, ok := msg.Groups[groupNumber].Segments.GetAsValue(segment)
	if !ok {
		return nil
	}
	value, ok := seg.(*orderedmap.OrderedMap[any]).GetAsValue(fieldId)
	if ok {
		switch v := value.(type) {
		case []*orderedmap.OrderedMap[string]:
			return v
		case *orderedmap.OrderedMap[string]:
			return nil
		default:
			return nil
		}
	}

	return nil
}

//func (msg *NCPDPMessage) GetFieldGroupCount(groupNumber int,segment string,fieldID string) int {
//	return msg.Groups[groupNumber].Segments.Get(segment).FieldGroups[fieldID].
//}

//
//func (msg *NCPDPMessage) GetFieldValue(groupNumber int, segmentId string, fieldGroupNumber int, fieldId string) string {
//	return ncpdpparser.GetFromMap(msg.Groups[strconv.Itoa(groupNumber)].Segments[segmentId].FieldGroups[fieldGroupNumber], fieldId)
//}
