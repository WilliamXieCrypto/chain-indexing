package event

import (
	"bytes"

	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	ibc_model "github.com/WilliamXieCrypto/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_TIMEOUT_ON_CLOSE = "MsgTimeoutOnClose"
const MSG_IBC_TIMEOUT_ON_CLOSE_CREATED = "MsgTimeoutOnCloseCreated"
const MSG_IBC_TIMEOUT_ON_CLOSE_FAILED = "MsgTimeoutOnCloseFailed"

type MsgIBCTimeoutOnClose struct {
	MsgBase

	Params ibc_model.MsgTimeoutOnCloseParams `json:"params"`
}

func NewMsgIBCTimeoutOnClose(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgTimeoutOnCloseParams,
) *MsgIBCTimeoutOnClose {
	return &MsgIBCTimeoutOnClose{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_TIMEOUT_ON_CLOSE,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCTimeoutOnClose) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCTimeoutOnClose) String() string {
	return render.Render(event)
}

func DecodeMsgIBCTimeoutOnClose(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCTimeoutOnClose
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}

const MSG_ALREADY_RELAYED_IBC_TIMEOUT_ON_CLOSE = "MsgAlreadyRelayedTimeoutOnClose"
const MSG_ALREADY_RELAYED_IBC_TIMEOUT_ON_CLOSE_CREATED = "MsgAlreadyRelayedTimeoutOnCloseCreated"
const MSG_ALREADY_RELAYED_IBC_TIMEOUT_ON_CLOSE_FAILED = "MsgAlreadyRelayedTimeoutOnCloseFailed"

type MsgAlreadyRelayedIBCTimeoutOnClose struct {
	MsgBase

	Params ibc_model.MsgTimeoutOnCloseParams `json:"params"`
}

func NewMsgAlreadyRelayedIBCTimeoutOnClose(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgTimeoutOnCloseParams,
) *MsgAlreadyRelayedIBCTimeoutOnClose {
	return &MsgAlreadyRelayedIBCTimeoutOnClose{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_ALREADY_RELAYED_IBC_TIMEOUT_ON_CLOSE,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgAlreadyRelayedIBCTimeoutOnClose) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgAlreadyRelayedIBCTimeoutOnClose) String() string {
	return render.Render(event)
}

func DecodeMsgAlreadyRelayedIBCTimeoutOnClose(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgAlreadyRelayedIBCTimeoutOnClose
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
