package ncpdpmodules

type NCPDPMessage struct {
	Segments map[string]Segment
}

func (msg *NCPDPMessage) GetFieldValue(segmentId, fieldID string) string {
	fields := msg.GetField(segmentId, fieldID)
	if len(fields) == 0 {
		return ""
	}
	return fields[0]

}
func (msg *NCPDPMessage) GetField(segmentId, fieldID string) []string {
	segment := msg.Segments[segmentId]
	return segment.Fields[fieldID]

}
