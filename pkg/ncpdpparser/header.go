package ncpdpparser

import (
	"fmt"
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
