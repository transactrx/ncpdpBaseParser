package ncpdpparser

import (
	"errors"
	"fmt"
	"github.com/transactrx/ncpdpBaseParser/pkg/ncpdpmodules"
	"github.com/transactrx/ncpdpBaseParser/pkg/orderedmap"
	"strings"
	"sync"
)

var groupSeparator = "\u001D"
var fieldSeparator = "\u001C"
var segmentSeparator = "\u001E"

var GROUP_FIELDS = initializeGroupedFields()

func initializeGroupedFields() map[string][]string {
	grpFields := make(map[string][]string)
	//COMPOUNDS
	grpFields["EC"] = []string{"RE", "TE", "ED", "EE", "UE"}

	//RESPONSE MESSAGES

	//APPROVED MESSAGES
	grpFields["5F"] = []string{"6F"} //RESPONSE STATUS SEGMENT - APPROVED MESSAGE CODE

	grpFields["UF"] = []string{"UH", "FQ", "UG"} //RESPONSE STATUS SEGMENT - ADDITIONAL MESSAGE INFORMATION
	//REJECT CODES
	grpFields["FA"] = []string{"FB", "4F"} //RESPONSE STATUS SEGMENT -ADDITIONAL MESSAGE INFORMATION

	grpFields["2G"] = []string{"2H"} //COMPOUND SEGMENT

	grpFields["9F"] = []string{"AR", "AS", "AT", "AU"} //RESPONSE CLAIM SEGMENT

	grpFields["J2"] = []string{"J3", "J4", "J5"} //RESPONSE PRICING SEGMENT

	grpFields["MU"] = []string{"MV", "MW"} //COORDINATION OF BENEFITS/OTHER PAYMENTS SEGMENT

	grpFields["J6"] = []string{"E4", "FS", "FT", "FU", "FV", "FW", "FX ", "FY ", "NS"} //RESPONSE DUR/PPS SEGMENT

	grpFields["NT"] = []string{"5C", "6C", "7C", "MH", "NU", "MJ", "UV", "UB", "UW", "UX", "UY"} //RESPONSE COORDINATION OF BENEFITS/OTHER PAYERS SEGMENT

	return grpFields
}

///	Npi,Name,Address

var mu sync.Mutex

func New(ncpdp string) (*ncpdpmodules.NCPDPMessage, error) {
	if ncpdp == "" {
		return nil, errors.New("NCPDP string is empty")

	}

	msg := ncpdpmodules.NCPDPMessage{}
	msg.Groups = make([]*ncpdpmodules.Group, 0)
	groups := strings.Split(ncpdp, groupSeparator)

	for i, groupString := range groups {
		msg.Groups = append(msg.Groups, parseGroup(groupString, i))
	}

	return &msg, nil

}

func parseGroup(groupString string, groupNumber int) *ncpdpmodules.Group {

	resultGroup := ncpdpmodules.Group{}
	segs := orderedmap.NewOrderedMap[any]()

	resultGroup.Segments = segs

	if groupString[0:1] == segmentSeparator {
		groupString = groupString[1:]
	}
	segments := strings.Split(groupString, segmentSeparator)

	for i, val := range segments {
		if i == 0 && groupNumber == 0 {
			//header
			continue
		}
		newSegment := parseSegment(val)
		segmentID := newSegment.Get("AM")
		if segmentID != nil {

			resultGroup.Segments.Put(fmt.Sprintf("%s", *segmentID), newSegment)
		}

	}
	return &resultGroup
}

func parseSegment(s string) *orderedmap.OrderedMap[any] {
	segment := orderedmap.NewOrderedMap[any]()

	if s[0:1] == fieldSeparator {
		s = s[1:]
	}
	fieldData := strings.Split(s, fieldSeparator)

	var currentGroupedFieldIds []string
	currentGroupMemberPos := -1
	var groupFieldId *string
	for _, fieldDatum := range fieldData {

		fieldId := fieldDatum[0:2]
		fieldValue := fieldDatum[2:]
		//IS THIS FIELD THE COUNTER FOR A REPEATING GROUP?
		val, ok := GROUP_FIELDS[fieldId]
		if ok {

			currentGroupedFieldIds = val
			if segment.Get(fieldId) == nil {
				groupedFieldArr := make([]*orderedmap.OrderedMap[string], 1)
				groupFieldId = &fieldId
				groupedFieldArr[0] = orderedmap.NewOrderedMap[string]()
				currentGroupMemberPos = -1
				segment.Put(fieldId, groupedFieldArr)
			}

		} else {
			foundPos := ContainsInArray(currentGroupedFieldIds, fieldId)

			if groupFieldId != nil && segment.Get(*groupFieldId) != nil && foundPos > -1 {

				var groupValArrObj = *(segment.Get(*groupFieldId))
				groupValArray, _ := groupValArrObj.([]*orderedmap.OrderedMap[string])
				curIndex := len(groupValArray) - 1

				if foundPos > currentGroupMemberPos {
					currentGroupMemberPos = foundPos
					groupValArray[curIndex].Put(fieldId, fieldValue)
				} else {
					currentGroupMemberPos = foundPos

					groupValArray = append(groupValArray, orderedmap.NewOrderedMap[string]())
					curIndex = len(groupValArray) - 1

					segment.Put(*groupFieldId, groupValArray)

					groupValArray[curIndex].Put(fieldId, fieldValue)

				}

			} else {

				currentGroupedFieldIds = nil
				segment.Put(fieldId, fieldValue)
			}

		}

	}

	return segment

}
