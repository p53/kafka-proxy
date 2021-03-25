package protocol

import "fmt"

type DeleteTopicsRequestFactory struct{}

func (f *DeleteTopicsRequestFactory) Produce(requestKeyVersion *RequestKeyVersion) (req ProtocolBody, err error) {
	switch requestKeyVersion.ApiVersion {
	case 0:
		return &DeleteTopicsRequestV0{}, nil
	case 1:
		return &DeleteTopicsRequestV1{}, nil
	case 2:
		return &DeleteTopicsRequestV2{}, nil
	case 3:
		return &DeleteTopicsRequestV3{}, nil
	case 4:
		return &DeleteTopicsRequestV4{}, nil
	case 5:
		return &DeleteTopicsRequestV5{}, nil
	default:
		return nil, fmt.Errorf("Not supported topic delete request %d", requestKeyVersion.ApiVersion)
	}
}

type DeleteTopicsRequestV0 struct {
	Topics []string
}

func (r *DeleteTopicsRequestV0) encode(pe packetEncoder) error {
	return nil
}

func (r *DeleteTopicsRequestV0) key() int16 {
	return 20
}

func (r *DeleteTopicsRequestV0) version() int16 {
	return 0
}

func (r *DeleteTopicsRequestV0) decode(pd packetDecoder) (err error) {
	r.Topics, err = pd.getStringArray()
	if err != nil {
		return err
	}

	// timeout_ms
	_, err = pd.getInt32()
	if err != nil {
		return err
	}

	return err
}

func (r *DeleteTopicsRequestV0) GetTopics() []string {
	return r.Topics
}

type DeleteTopicsRequestV1 struct {
	Topics []string
}

func (r *DeleteTopicsRequestV1) encode(pe packetEncoder) error {
	return nil
}

func (r *DeleteTopicsRequestV1) key() int16 {
	return 20
}

func (r *DeleteTopicsRequestV1) version() int16 {
	return 1
}

func (r *DeleteTopicsRequestV1) decode(pd packetDecoder) (err error) {
	r.Topics, err = pd.getStringArray()
	if err != nil {
		return err
	}

	// timeout_ms
	_, err = pd.getInt32()
	if err != nil {
		return err
	}

	return err
}

func (r *DeleteTopicsRequestV1) GetTopics() []string {
	return r.Topics
}

type DeleteTopicsRequestV2 struct {
	Topics []string
}

func (r *DeleteTopicsRequestV2) encode(pe packetEncoder) error {
	return nil
}

func (r *DeleteTopicsRequestV2) key() int16 {
	return 20
}

func (r *DeleteTopicsRequestV2) version() int16 {
	return 2
}

func (r *DeleteTopicsRequestV2) decode(pd packetDecoder) (err error) {
	r.Topics, err = pd.getStringArray()
	if err != nil {
		return err
	}

	// timeout_ms
	_, err = pd.getInt32()
	if err != nil {
		return err
	}

	return err
}

func (r *DeleteTopicsRequestV2) GetTopics() []string {
	return r.Topics
}

type DeleteTopicsRequestV3 struct {
	Topics []string
}

func (r *DeleteTopicsRequestV3) encode(pe packetEncoder) error {
	return nil
}

func (r *DeleteTopicsRequestV3) key() int16 {
	return 20
}

func (r *DeleteTopicsRequestV3) version() int16 {
	return 3
}

func (r *DeleteTopicsRequestV3) decode(pd packetDecoder) (err error) {
	r.Topics, err = pd.getStringArray()
	if err != nil {
		return err
	}

	// timeout_ms
	_, err = pd.getInt32()
	if err != nil {
		return err
	}

	return err
}

func (r *DeleteTopicsRequestV3) GetTopics() []string {
	return r.Topics
}

type DeleteTopicsRequestV4 struct {
	Topics []string
}

func (r *DeleteTopicsRequestV4) encode(pe packetEncoder) error {
	return nil
}

func (r *DeleteTopicsRequestV4) key() int16 {
	return 20
}

func (r *DeleteTopicsRequestV4) version() int16 {
	return 4
}

func (r *DeleteTopicsRequestV4) decode(pd packetDecoder) (err error) {
	r.Topics, err = pd.getCompactStringArray()
	if err != nil {
		return err
	}

	// timeout_ms
	_, err = pd.getInt32()
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

func (r *DeleteTopicsRequestV4) GetTopics() []string {
	return r.Topics
}

type DeleteTopicsRequestV5 struct {
	Topics []string
}

func (r *DeleteTopicsRequestV5) encode(pe packetEncoder) error {
	return nil
}

func (r *DeleteTopicsRequestV5) key() int16 {
	return 20
}

func (r *DeleteTopicsRequestV5) version() int16 {
	return 5
}

func (r *DeleteTopicsRequestV5) decode(pd packetDecoder) (err error) {
	r.Topics, err = pd.getCompactStringArray()
	if err != nil {
		return err
	}

	// timeout_ms
	_, err = pd.getInt32()
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

func (r *DeleteTopicsRequestV5) GetTopics() []string {
	return r.Topics
}
