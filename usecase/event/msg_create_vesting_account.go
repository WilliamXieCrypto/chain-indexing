package event

import (
	"bytes"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"

	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_CREATE_VESTING_ACCOUNT = "MsgCreateVestingAccount"
const MSG_CREATE_VESTING_ACCOUNT_CREATED = "MsgCreateVestingAccountCreated"
const MSG_CREATE_VESTING_ACCOUNT_FAILED = "MsgCreateVestingAccountFailed"

type MsgCreateVestingAccount struct {
	MsgBase

	Params model.MsgCreateVestingAccountParams `json:"params"`
}

func NewMsgCreateVestingAccount(
	msgCommonParams MsgCommonParams,
	params model.MsgCreateVestingAccountParams,
) *MsgCreateVestingAccount {
	return &MsgCreateVestingAccount{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_CREATE_VESTING_ACCOUNT,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgCreateVestingAccount) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgCreateVestingAccount) String() string {
	return render.Render(event)
}

func DecodeMsgCreateVestingAccount(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgCreateVestingAccount
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
