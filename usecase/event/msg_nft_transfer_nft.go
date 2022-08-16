package event

import (
	"bytes"

	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"

	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_NFT_TRANSFER_NFT = "MsgTransferNFT"
const MSG_NFT_TRANSFER_NFT_CREATED = "MsgTransferNFTCreated"
const MSG_NFT_TRANSFER_NFT_FAILED = "MsgTransferNFTFailed"

type MsgNFTTransferNFT struct {
	MsgBase

	DenomId   string `json:"denomId"`
	TokenId   string `json:"tokenId"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
}

func NewMsgNFTTransferNFT(
	msgCommonParams MsgCommonParams,
	params model.MsgNFTTransferNFTParams,
) *MsgNFTTransferNFT {
	return &MsgNFTTransferNFT{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_NFT_TRANSFER_NFT,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params.DenomId,
		params.TokenId,
		params.Sender,
		params.Recipient,
	}
}

func (event *MsgNFTTransferNFT) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgNFTTransferNFT) String() string {
	return render.Render(event)
}

func DecodeMsgNFTTransferNFT(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgNFTTransferNFT
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
