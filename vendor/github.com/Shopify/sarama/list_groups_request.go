package sarama

type ListGroupsRequest struct {
}

func (r *ListGroupsRequest) encode(pe packetEncoder) error {
	return nil
}

func (r *ListGroupsRequest) decode(pd packetDecoder, version int16) (err error) {
	return nil
}

func (r *ListGroupsRequest) key() int16 {
	return 16
}

func (r *ListGroupsRequest) version() int16 {
	return 2
}

func (r *ListGroupsRequest) requiredVersion() KafkaVersion {
	return V2_3_0_0
}
