package ncpdpparser

import (
	"errors"
	"fmt"
	"github.com/transactrx/ncpdpBaseParser/pkg/orderedmap"
)

func DetermineTransactionType(data []byte) (int, error) {

	//check if response object
	headerInfo := string(data[0:4])
	if headerInfo == "D0B1" || headerInfo == "DXB1" {
		return B1ResponseType, nil
	}
	if headerInfo == "D0B2" || headerInfo == "DXB2" {
		return B2ResponseType, nil
	}
	if headerInfo == "D0B3" || headerInfo == "DXB3" {
		return B3ResponseType, nil
	}
	if headerInfo == "D0S1" || headerInfo == "DXS1" {
		return S1ResponseType, nil
	}
	if headerInfo == "D0S2" || headerInfo == "DXS2" {
		return S2ResponseType, nil
	}
	if headerInfo == "D0E1" || headerInfo == "DXE1" {
		return E1ResponseType, nil
	}
	if headerInfo == "D0N1" || headerInfo == "DXN1" {
		return N1ResponseType, nil
	}
	if headerInfo == "D0Q1" || headerInfo == "DXQ1" {
		return Q1ResponseType, nil
	}
	if headerInfo == "D0Q2" || headerInfo == "DXQ2" {
		return Q2ResponseType, nil
	}

	headerInfo = string(data[6:10])
	if headerInfo == "D0B1" || headerInfo == "DXB1" {
		return B1RequestType, nil
	}
	if headerInfo == "D0B2" || headerInfo == "DXB2" {
		return B2RequestType, nil
	}
	if headerInfo == "D0B3" || headerInfo == "DXB3" {
		return B3RequestType, nil
	}
	if headerInfo == "D0S1" || headerInfo == "DXS1" {
		return S1RequestType, nil
	}
	if headerInfo == "D0S2" || headerInfo == "DXS2" {
		return S2RequestType, nil
	}

	if headerInfo == "D0Q1" || headerInfo == "DXQ1" {
		return Q1RequestType, nil
	}
	if headerInfo == "D0Q2" || headerInfo == "DXQ2" {
		return Q2RequestType, nil
	}

	fmt.Printf("Unable to parse transactions. NCPDP message is invalid or unsupported -> %s", string(data))
	return 0, NCPDPMessageInvalidOrUnsupported
}

func newRequestHeader(data []byte) (*orderedmap.OrderedMap[string], error) {

	//021684D0B1DATAUNVAIL4011669742144     20230705
	if len(data) != 56 {
		return nil, errors.New(fmt.Sprintf("header cannot be parsed, it doesnt have the right length %d", len(data)))
	}

	seg := orderedmap.NewOrderedMap[string]()
	bin := data[0:6]
	versionRel := data[6:8]
	transactionCode := data[8:10]
	pcn := data[10:20]
	transactionCount := data[20:21]
	serviceProviderIdQualifier := data[21:23]
	serviceProviderId := data[23:38]
	dos := data[38:46]
	softwareCertId := data[46:]

	seg.Put("bin", string(bin))
	seg.Put("versionRel", string(versionRel))
	seg.Put("transactionCode", string(transactionCode))
	seg.Put("pcn", string(pcn))
	seg.Put("transactionCount", string(transactionCount))
	seg.Put("serviceProviderIdQualifier", string(serviceProviderIdQualifier))
	seg.Put("serviceProviderId", string(serviceProviderId))
	seg.Put("dos", string(dos))
	seg.Put("softwareCertId", string(softwareCertId))

	return seg, nil
}

func newResponseHeader(data []byte) (*orderedmap.OrderedMap[string], error) {

	//D0B11A011861586273     20230629
	if len(data) != 31 {
		return nil, errors.New(fmt.Sprintf("header cannot be parsed, it doesnt have the right length %d", len(data)))
	}

	seg := orderedmap.NewOrderedMap[string]()
	versionRel := data[0:2]
	transactionCode := data[2:4]
	transactionCount := data[4:5]
	transactionResponseStatus := data[5:6]
	serviceProviderIdQualifier := data[6:8]
	serviceProviderId := data[8:23]
	dos := data[23:31]

	seg.Put("versionRel", string(versionRel))
	seg.Put("transactionCode", string(transactionCode))
	seg.Put("transactionCount", string(transactionCount))
	seg.Put("transactionResponseStatus", string(transactionResponseStatus))
	seg.Put("serviceProviderIdQualifier", string(serviceProviderIdQualifier))
	seg.Put("serviceProviderId", string(serviceProviderId))
	seg.Put("dos", string(dos))

	return seg, nil
}

func parseHeader(data []byte) (*orderedmap.OrderedMap[string], error) {
	messageType, err := DetermineTransactionType(data)

	if err != nil {
		return nil, err
	}

	switch messageType {
	case B1RequestType, B2RequestType, B3RequestType, N1RequestType, S1RequestType, S2RequestType, E1RequestType, Q1RequestType, Q2RequestType:
		return newRequestHeader(data)
	case B1ResponseType, B2ResponseType, B3ResponseType, N1ResponseType, S1ResponseType, S2ResponseType, E1ResponseType, Q1ResponseType, Q2ResponseType:
		return newResponseHeader(data)
	default:
		return nil, errors.New("unable to determine transaction type")

	}

}
