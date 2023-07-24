package ncpdpparser

import "errors"

var NCPDPMessageInvalidOrUnsupported = errors.New("the NCPDP message is invalid or unsupported.  Currently support only exist for b1 and b2 request and response messages")
