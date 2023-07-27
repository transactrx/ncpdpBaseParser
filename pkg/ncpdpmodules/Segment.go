package ncpdpmodules

import "github.com/transactrx/ncpdpBaseParser/pkg/orderedmap"

type Segment struct {
	FieldGroups *orderedmap.OrderedMap[interface{}]
	Header      bool
}
