package ncpdpparser

import (
	"errors"
	"fmt"
	"github.com/transactrx/ncpdpBaseParser/pkg/ncpdpmodules"
)

func DetermineTransactionType(data []byte) (int, error) {

	//check if response object
	headerInfo := string(data[0:4])
	if headerInfo == "D0B1" {
		return B1ResponseType, nil
	}
	if headerInfo == "D0B2" {
		return B2ResponseType, nil
	}
	if headerInfo == "D0B3" {
		return B3ResponseType, nil
	}
	if headerInfo == "D0S1" {
		return S1ResponseType, nil
	}
	if headerInfo == "D0S2" {
		return S2ResponseType, nil
	}
	if headerInfo == "D0E1" {
		return E1ResponseType, nil
	}
	if headerInfo == "D0N1" {
		return N1ResponseType, nil
	}

	headerInfo = string(data[6:10])
	if headerInfo == "D0B1" {
		return B1RequestType, nil
	}
	if headerInfo == "D0B2" {
		return B2RequestType, nil
	}
	if headerInfo == "D0B3" {
		return B3RequestType, nil
	}
	if headerInfo == "D0S1" {
		return S1RequestType, nil
	}
	if headerInfo == "D0S2" {
		return S1RequestType, nil
	}

	fmt.Printf("Unable to parse transactions. NCPDP message is invalid or unsupported -> %s", string(data))
	return 0, NCPDPMessageInvalidOrUnsupported
}

func newSegmentHeader(seg *ncpdpmodules.Segment, data []byte) error {

	bin := data[0:6]
	versionRel := data[6:8]
	transactionCode := data[8:10]
	pcn := data[10:20]
	transactionCount := data[20:21]
	serviceProviderIdQualifier := data[21:23]
	serviceProviderId := data[23:38]
	dos := data[38:46]
	softwareCertId := data[46:]

	field1, err := newHeaderField(bin, "bin", 6)
	if err != nil {
		return nil, err
	}
	seg.addField(field1)
	field2, err := newHeaderField(versionRel, "versionRel", 2)
	if err != nil {
		return nil, err
	}
	seg.addField(field2)
	field3, err := newHeaderField(transactionCode, "transactionCode", 2)
	if err != nil {
		return nil, err
	}
	seg.addField(field3)
	field4, err := newHeaderField(pcn, "pcn", 10)
	if err != nil {
		return nil, err
	}
	seg.addField(field4)
	field5, err := newHeaderField(transactionCount, "transactionCount", 1)
	if err != nil {
		return nil, err
	}
	seg.addField(field5)
	field6, err := newHeaderField(serviceProviderIdQualifier, "serviceProviderIdQualifier", 2)
	if err != nil {
		return nil, err
	}
	seg.addField(field6)
	field7, err := newHeaderField(serviceProviderId, "serviceProviderId", 15)
	if err != nil {
		return nil, err
	}
	seg.addField(field7)
	field8, err := newHeaderField(dos, "dos", 8)
	if err != nil {
		return nil, err
	}
	seg.addField(field8)
	field9, err := newHeaderField(softwareCertId, "softwareCertId", 10)
	if err != nil {
		return nil, err
	}
	seg.addField(field9)

	return &seg, nil
}

func newResponseSegmentHeader(data []byte) (*segment, error) {
	seg := segment{header: true}
	seg.id = "header"
	seg.fieldsMap = make(map[string]interface{})

	versionRel := data[0:2]
	transactionCode := data[2:4]
	transactionCount := data[4:5]
	transactionResponseStatus := data[5:6]
	serviceProviderIdQualifier := data[6:8]
	serviceProviderId := data[8:23]
	dos := data[23:31]

	field1, err := newHeaderField(versionRel, "versionRel", 2)
	if err != nil {
		return nil, err
	}
	seg.addField(field1)
	field2, err := newHeaderField(transactionCode, "transactionCode", 2)
	if err != nil {
		return nil, err
	}
	seg.addField(field2)

	field3, err := newHeaderField(transactionCount, "transactionCount", 1)
	if err != nil {
		return nil, err
	}
	seg.addField(field3)

	field4, err := newHeaderField(transactionResponseStatus, "transactionResponseStatus", 1)
	if err != nil {
		return nil, err
	}
	seg.addField(field4)

	field5, err := newHeaderField(serviceProviderIdQualifier, "serviceProviderIdQualifier", 2)
	if err != nil {
		return nil, err
	}
	seg.addField(field5)
	field6, err := newHeaderField(serviceProviderId, "serviceProviderId", 15)
	if err != nil {
		return nil, err
	}
	seg.addField(field6)
	field7, err := newHeaderField(dos, "dos", 8)
	if err != nil {
		return nil, err
	}
	seg.addField(field7)

	return &seg, nil
}

func parseHeader(data []byte) (*ncpdpmodules.Segment, error) {
	messageType, err := DetermineTransactionType(data)

	if err != nil {
		return nil, err
	}

	switch messageType {
	case B1RequestType, B2RequestType, B3RequestType, N1RequestType, S1RequestType, S2RequestType, E1RequestType:
		{
			seg := ncpdpmodules.Segment{Header: true}

		}
	case B1ResponseType, B2ResponseType, B3ResponseType, N1ResponseType, S1ResponseType, S2ResponseType, E1ResponseType:
		{

		}
	default:
		return nil, errors.New("unable to determine transaction type")

	}

}
