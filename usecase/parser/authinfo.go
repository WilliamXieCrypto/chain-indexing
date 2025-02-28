package parser

import (
	"fmt"
	"strconv"

	cosmosapp_interface "github.com/WilliamXieCrypto/chain-indexing/appinterface/cosmosapp"
	"github.com/WilliamXieCrypto/chain-indexing/external/tmcosmosutils"
	"github.com/WilliamXieCrypto/chain-indexing/internal/base64"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/parser/utils"
)

func ParseSignerInfosToTransactionSigners(
	cosmosClient cosmosapp_interface.Client,
	signerInfos []utils.SignerInfo,
	accountAddressPrefix string,
	possibleSignerAddresses []string,
) ([]model.TransactionSigner, error) {
	var signers []model.TransactionSigner

	if len(signerInfos) <= 0 && len(possibleSignerAddresses) <= 0 {
		panic("error signer info not found")
	}

	for i, signer := range signerInfos {
		var transactionSignerInfo *model.TransactionSignerKeyInfo
		var address string

		sequence, parseErr := strconv.ParseUint(signer.Sequence, 10, 64)
		if parseErr != nil {
			return nil, fmt.Errorf("error parsing account sequence: %v", parseErr)
		}

		if signer.ModeInfo.MaybeSingle != nil {
			if signer.MaybePublicKey == nil {
				if len(possibleSignerAddresses) < i+1 {
					address = ""
				} else {
					address = possibleSignerAddresses[i]
					accountInfo, _ := cosmosClient.Account(address)
					if accountInfo != nil {
						transactionSignerInfo = &model.TransactionSignerKeyInfo{
							Type:       accountInfo.MaybePubkey.Type,
							IsMultiSig: false,
							Pubkeys:    []string{accountInfo.MaybePubkey.Key},
						}
					}
				}
			} else {
				transactionSignerInfo = &model.TransactionSignerKeyInfo{
					Type:       signer.MaybePublicKey.Type,
					IsMultiSig: false,
					Pubkeys:    []string{*signer.MaybePublicKey.MaybeKey},
				}

				parsedAddr, parseAddrErr := ParseTransactionSignerInfoToAddress(*transactionSignerInfo, accountAddressPrefix)
				if parseAddrErr != nil {
					return nil, fmt.Errorf("error parsing signer info to address: %v", parseAddrErr)
				}
				address = parsedAddr
			}
		} else {
			pubkeys := make([]string, 0, len(signer.MaybePublicKey.MaybePublicKeys))
			for _, pubkey := range signer.MaybePublicKey.MaybePublicKeys {
				pubkeys = append(pubkeys, pubkey.Key)
			}
			transactionSignerInfo = &model.TransactionSignerKeyInfo{
				Type:           signer.MaybePublicKey.Type,
				IsMultiSig:     true,
				Pubkeys:        pubkeys,
				MaybeThreshold: signer.MaybePublicKey.MaybeThreshold,
			}

			parsedAddr, parseAddrErr := ParseTransactionSignerInfoToAddress(*transactionSignerInfo, accountAddressPrefix)
			if parseAddrErr != nil {
				return nil, fmt.Errorf("error parsing signer info to address: %v", parseAddrErr)
			}
			address = parsedAddr
		}

		signers = append(signers, model.TransactionSigner{
			MaybeKeyInfo:    transactionSignerInfo,
			Address:         address,
			AccountSequence: sequence,
		})
	}

	return signers, nil
}

func ParseTransactionSignerInfoToAddress(
	signerInfo model.TransactionSignerKeyInfo,
	accountAddressPrefix string,
) (string, error) {
	var address string
	if signerInfo.IsMultiSig {
		addrPubKeys := make([][]byte, 0, len(signerInfo.Pubkeys))
		for _, pubKey := range signerInfo.Pubkeys {
			rawPubKey := base64.MustDecodeString(pubKey)
			addrPubKeys = append(addrPubKeys, rawPubKey)
		}
		var multiSigAddrErr error
		address, multiSigAddrErr = tmcosmosutils.MultiSigAddressFromPubKeys(
			accountAddressPrefix,
			addrPubKeys,
			*signerInfo.MaybeThreshold,
			false,
		)
		if multiSigAddrErr != nil {
			return "", fmt.Errorf("error converting public keys to multisig address: %v", multiSigAddrErr)
		}
	} else {
		var addrErr error
		pubKey := base64.MustDecodeString(signerInfo.Pubkeys[0])
		address, addrErr = tmcosmosutils.AccountAddressFromPubKey(accountAddressPrefix, pubKey)
		if addrErr != nil {
			return "", fmt.Errorf("error converting public key to address: %v", addrErr)
		}
	}
	return address, nil
}
