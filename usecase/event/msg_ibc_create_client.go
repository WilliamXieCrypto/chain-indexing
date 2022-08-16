package event

import (
	"bytes"

	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	ibc_model "github.com/WilliamXieCrypto/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_CREATE_CLIENT = "MsgCreateClient"
const MSG_IBC_CREATE_CLIENT_CREATED = "MsgCreateClientCreated"
const MSG_IBC_CREATE_CLIENT_FAILED = "MsgCreateClientFailed"

type MsgIBCCreateClient struct {
	MsgBase

	Params ibc_model.MsgCreateClientParams `json:"params"`
}

func NewMsgIBCCreateClient(msgCommonParams MsgCommonParams, params ibc_model.MsgCreateClientParams) *MsgIBCCreateClient {
	return &MsgIBCCreateClient{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_CREATE_CLIENT,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCCreateClient) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCCreateClient) String() string {
	return render.Render(event)
}

func DecodeMsgIBCCreateClient(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCCreateClient
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
