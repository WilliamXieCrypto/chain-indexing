package event

import (
	"bytes"

	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	ibc_model "github.com/WilliamXieCrypto/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_CHANNEL_OPEN_INIT = "MsgChannelOpenInit"
const MSG_IBC_CHANNEL_OPEN_INIT_CREATED = "MsgChannelOpenInitCreated"
const MSG_IBC_CHANNEL_OPEN_INIT_FAILED = "MsgChannelOpenInitFailed"

type MsgIBCChannelOpenInit struct {
	MsgBase

	Params ibc_model.MsgChannelOpenInitParams `json:"params"`
}

func NewMsgIBCChannelOpenInit(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgChannelOpenInitParams,
) *MsgIBCChannelOpenInit {
	return &MsgIBCChannelOpenInit{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_CHANNEL_OPEN_INIT,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCChannelOpenInit) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCChannelOpenInit) String() string {
	return render.Render(event)
}

func DecodeMsgIBCChannelOpenInit(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCChannelOpenInit
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
