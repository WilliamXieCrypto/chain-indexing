package event

import (
	"bytes"

	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"

	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_NFT_BURN_NFT = "MsgBurnNFT"
const MSG_NFT_BURN_NFT_CREATED = "MsgBurnNFTCreated"
const MSG_NFT_BURN_NFT_FAILED = "MsgBurnNFTFailed"

type MsgNFTBurnNFT struct {
	MsgBase

	DenomId string `json:"denomId"`
	TokenId string `json:"tokenId"`
	Sender  string `json:"sender"`
}

func NewMsgNFTBurnNFT(
	msgCommonParams MsgCommonParams,
	params model.MsgNFTBurnNFTParams,
) *MsgNFTBurnNFT {
	return &MsgNFTBurnNFT{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_NFT_BURN_NFT,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params.DenomId,
		params.TokenId,
		params.Sender,
	}
}

func (event *MsgNFTBurnNFT) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgNFTBurnNFT) String() string {
	return render.Render(event)
}

func DecodeMsgNFTBurnNFT(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgNFTBurnNFT
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
