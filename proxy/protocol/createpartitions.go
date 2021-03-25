package protocol

import "fmt"

type CreatePartitionsRequestFactory struct{}

func (f *CreatePartitionsRequestFactory) Produce(requestKeyVersion *RequestKeyVersion) (req ProtocolBody, err error) {
	switch requestKeyVersion.ApiVersion {
	case 0:
		return &CreatePartitionsRequestV0{}, nil
	case 1:
		return &CreatePartitionsRequestV1{}, nil
	case 2:
		return &CreatePartitionsRequestV2{}, nil
	case 3:
		return &CreatePartitionsRequestV3{}, nil
	default:
		return nil, fmt.Errorf("Not supported create partitions request %d", requestKeyVersion.ApiVersion)
	}
}

type CreatePartitionsRequestV0 struct {
	Topics []string
}

func (r *CreatePartitionsRequestV0) encode(pe packetEncoder) error {
	return nil
}

func (r *CreatePartitionsRequestV0) key() int16 {
	return 37
}

func (r *CreatePartitionsRequestV0) version() int16 {
	return 0
}

func (r *CreatePartitionsRequestV0) decode(pd packetDecoder) (err error) {
	// get length of topic array
	numTopics, err := pd.getInt32()
	if err != nil {
		return err
	}

	for i := int32(1); i <= numTopics; i++ {
		topicName, err := pd.getString()
		if err != nil {
			return err
		}

		r.Topics = append(r.Topics, topicName)

		// count - partition count
		_, err = pd.getInt32()
		if err != nil {
			return err
		}

		// get length of assignments array
		numAssignments, err := pd.getInt32()
		if err != nil {
			return err
		}

		for j := int32(1); j <= numAssignments; j++ {

			// get length of broker ids array
			numBrokers, err := pd.getInt32()
			if err != nil {
				return err
			}

			for j := int32(1); j <= numBrokers; j++ {
				// broker_ids
				_, err = pd.getInt32()
				if err != nil {
					return err
				}
			}
		}
	}

	// timeout_ms
	_, err = pd.getInt32()
	if err != nil {
		return err
	}

	// validate_only
	_, err = pd.getBool()
	if err != nil {
		return err
	}

	return err
}

func (r *CreatePartitionsRequestV0) GetTopics() []string {
	return r.Topics
}

type CreatePartitionsRequestV1 struct {
	Topics []string
}

func (r *CreatePartitionsRequestV1) encode(pe packetEncoder) error {
	return nil
}

func (r *CreatePartitionsRequestV1) key() int16 {
	return 37
}

func (r *CreatePartitionsRequestV1) version() int16 {
	return 1
}

func (r *CreatePartitionsRequestV1) decode(pd packetDecoder) (err error) {
	// get length of topic array
	numTopics, err := pd.getInt32()
	if err != nil {
		return err
	}

	for i := int32(1); i <= numTopics; i++ {
		topicName, err := pd.getString()
		if err != nil {
			return err
		}

		r.Topics = append(r.Topics, topicName)

		// count - partition count
		_, err = pd.getInt32()
		if err != nil {
			return err
		}

		// get length of assignments array
		numAssignments, err := pd.getInt32()
		if err != nil {
			return err
		}

		for j := int32(1); j <= numAssignments; j++ {

			// get length of broker ids array
			numBrokers, err := pd.getInt32()
			if err != nil {
				return err
			}

			for j := int32(1); j <= numBrokers; j++ {
				// broker_ids
				_, err = pd.getInt32()
				if err != nil {
					return err
				}
			}
		}
	}

	// timeout_ms
	_, err = pd.getInt32()
	if err != nil {
		return err
	}

	// validate_only
	_, err = pd.getBool()
	if err != nil {
		return err
	}

	return err
}

func (r *CreatePartitionsRequestV1) GetTopics() []string {
	return r.Topics
}

type CreatePartitionsRequestV2 struct {
	Topics []string
}

func (r *CreatePartitionsRequestV2) encode(pe packetEncoder) error {
	return nil
}

func (r *CreatePartitionsRequestV2) key() int16 {
	return 37
}

func (r *CreatePartitionsRequestV2) version() int16 {
	return 2
}

func (r *CreatePartitionsRequestV2) decode(pd packetDecoder) (err error) {
	// get length of topic array
	numTopics, err := pd.getInt32()
	if err != nil {
		return err
	}

	for i := int32(1); i <= numTopics; i++ {
		topicName, err := pd.getCompactString()
		if err != nil {
			return err
		}

		r.Topics = append(r.Topics, topicName)

		// count - partition count
		_, err = pd.getInt32()
		if err != nil {
			return err
		}

		// get length of assignments array
		numAssignments, err := pd.getInt32()
		if err != nil {
			return err
		}

		for j := int32(1); j <= numAssignments; j++ {

			// get length of broker ids array
			numBrokers, err := pd.getInt32()
			if err != nil {
				return err
			}

			for j := int32(1); j <= numBrokers; j++ {
				// broker_ids
				_, err = pd.getInt32()
				if err != nil {
					return err
				}
			}

			assignTf := TaggedFields{}
			err = assignTf.decode(pd)

			if err != nil {
				return err
			}
		}

		topTf := TaggedFields{}
		err = topTf.decode(pd)

		if err != nil {
			return err
		}
	}

	// timeout_ms
	_, err = pd.getInt32()
	if err != nil {
		return err
	}

	// validate_only
	_, err = pd.getBool()
	if err != nil {
		return err
	}

	reqTf := TaggedFields{}
	err = reqTf.decode(pd)

	if err != nil {
		return err
	}

	return err
}

func (r *CreatePartitionsRequestV2) GetTopics() []string {
	return r.Topics
}

type CreatePartitionsRequestV3 struct {
	Topics []string
}

func (r *CreatePartitionsRequestV3) encode(pe packetEncoder) error {
	return nil
}

func (r *CreatePartitionsRequestV3) key() int16 {
	return 37
}

func (r *CreatePartitionsRequestV3) version() int16 {
	return 3
}

func (r *CreatePartitionsRequestV3) decode(pd packetDecoder) (err error) {
	// get length of topic array
	numTopics, err := pd.getInt32()
	if err != nil {
		return err
	}

	for i := int32(1); i <= numTopics; i++ {
		topicName, err := pd.getCompactString()
		if err != nil {
			return err
		}

		r.Topics = append(r.Topics, topicName)

		// count - partition count
		_, err = pd.getInt32()
		if err != nil {
			return err
		}

		// get length of assignments array
		numAssignments, err := pd.getInt32()
		if err != nil {
			return err
		}

		for j := int32(1); j <= numAssignments; j++ {

			// get length of broker ids array
			numBrokers, err := pd.getInt32()
			if err != nil {
				return err
			}

			for j := int32(1); j <= numBrokers; j++ {
				// broker_ids
				_, err = pd.getInt32()
				if err != nil {
					return err
				}
			}

			assignTf := TaggedFields{}
			err = assignTf.decode(pd)

			if err != nil {
				return err
			}
		}

		topTf := TaggedFields{}
		err = topTf.decode(pd)

		if err != nil {
			return err
		}
	}

	// timeout_ms
	_, err = pd.getInt32()
	if err != nil {
		return err
	}

	// validate_only
	_, err = pd.getBool()
	if err != nil {
		return err
	}

	reqTf := TaggedFields{}
	err = reqTf.decode(pd)

	if err != nil {
		return err
	}

	return err
}

func (r *CreatePartitionsRequestV3) GetTopics() []string {
	return r.Topics
}
