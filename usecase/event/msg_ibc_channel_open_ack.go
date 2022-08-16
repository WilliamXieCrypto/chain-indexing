package event

import (
	"bytes"

	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	ibc_model "github.com/WilliamXieCrypto/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_CHANNEL_OPEN_ACK = "MsgChannelOpenAck"
const MSG_IBC_CHANNEL_OPEN_ACK_CREATED = "MsgChannelOpenAckCreated"
const MSG_IBC_CHANNEL_OPEN_ACK_FAILED = "MsgChannelOpenAckFailed"

type MsgIBCChannelOpenAck struct {
	MsgBase

	Params ibc_model.MsgChannelOpenAckParams `json:"params"`
}

func NewMsgIBCChannelOpenAck(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgChannelOpenAckParams,
) *MsgIBCChannelOpenAck {
	return &MsgIBCChannelOpenAck{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_CHANNEL_OPEN_ACK,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCChannelOpenAck) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCChannelOpenAck) String() string {
	return render.Render(event)
}

func DecodeMsgIBCChannelOpenAck(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCChannelOpenAck
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
