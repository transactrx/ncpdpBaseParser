package main

import (
	"github.com/transactrx/ncpdpBaseParser/pkg/ncpdpparser"
	"testing"
)

func TestCreateNcpdpObject(t *testing.T) {

	request := "61009751B19999      1011669688842     20111013          \u001E\u001CAM01\u001CCX01\u001CCY1045\u001CC419310202\u001CC52\u001CCADOROTHY\u001CCBPARRISH\u001E\u001CAM04\u001CC20055168301\u001CCCDOROTHY\u001CCDPARRISH\u001CC90\u001CC1PDPIND\u001CC3001\u001CC61\u001D\u001E\u001CAM07\u001CEM1\u001CD20723583\u001CE103\u001CD700006496300\u001CE70000001000\u001CD30\u001CD530\u001CD61\u001CD80\u001CDE20111013\u001CDJ3\u001CDK00\u001CC81\u001CDT0\u001E\u001CAM03\u001CEZ01\u001CDB1457347619\u001CDRLEINDECKER\u001E\u001CAM11\u001CD91850{\u001CDC30{\u001CE3200{\u001CDQ2050{\u001CDU2080{\u001CDN07\u001E\u001CAM08\u001C7E1\u001CE5MA\u001C8E11"
	ncpdp, err := ncpdpparser.New(request)
	if err != nil {
		t.Fatalf("Error parsing NCPDP request: %v", err)
	}

	if ncpdp == nil {
		t.Fatalf("NCPDP object is nil")
	}

	if ncpdp.GetFieldValue("01", "CA") == "" {
		t.Errorf("Unable to find field value for 01-CA. Expected Patient First Name")
	}

	if ncpdp.GetFieldValue("01", "CB") == "" {
		t.Errorf("Unable to find field value for 01-CB. Expected Patient Last Name")
	}

}
