package event

import (
	"bytes"

	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"

	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_NFT_ISSUE_DENOM = "MsgIssueDenom"
const MSG_NFT_ISSUE_DENOM_CREATED = "MsgIssueDenomCreated"
const MSG_NFT_ISSUE_DENOM_FAILED = "MsgIssueDenomFailed"

type MsgNFTIssueDenom struct {
	MsgBase

	DenomId   string `json:"denomId"`
	DenomName string `json:"denomName"`
	Schema    string `json:"schema"`
	Sender    string `json:"sender"`
}

func NewMsgNFTIssueDenom(
	msgCommonParams MsgCommonParams,
	params model.MsgNFTIssueDenomParams,
) *MsgNFTIssueDenom {
	return &MsgNFTIssueDenom{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_NFT_ISSUE_DENOM,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params.DenomId,
		params.DenomName,
		params.Schema,
		params.Sender,
	}
}

func (event *MsgNFTIssueDenom) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgNFTIssueDenom) String() string {
	return render.Render(event)
}

func DecodeMsgNFTIssueDenom(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgNFTIssueDenom
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
