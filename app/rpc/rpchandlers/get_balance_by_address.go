package rpchandlers

import (
	"github.com/ZorkNetwork/zorkmid/app/appmessage"
	"github.com/ZorkNetwork/zorkmid/app/rpc/rpccontext"
	"github.com/ZorkNetwork/zorkmid/domain/consensus/utils/txscript"
	"github.com/ZorkNetwork/zorkmid/infrastructure/network/netadapter/router"
	"github.com/ZorkNetwork/zorkmid/util"
	"github.com/pkg/errors"
)

// HandleGetBalanceByAddress handles the respectively named RPC command
func HandleGetBalanceByAddress(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := &appmessage.GetUTXOsByAddressesResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when kaspad is run without --utxoindex")
		return errorMessage, nil
	}

	getBalanceByAddressRequest := request.(*appmessage.GetBalanceByAddressRequestMessage)

	balance, err := getBalanceByAddress(context, getBalanceByAddressRequest.Address)
	if err != nil {
		rpcError := &appmessage.RPCError{}
		if !errors.As(err, &rpcError) {
			return nil, err
		}
		errorMessage := &appmessage.GetUTXOsByAddressesResponseMessage{}
		errorMessage.Error = rpcError
		return errorMessage, nil
	}

	response := appmessage.NewGetBalanceByAddressResponse(balance)
	return response, nil
}

func getBalanceByAddress(context *rpccontext.Context, addressString string) (uint64, error) {
	address, err := util.DecodeAddress(addressString, context.Config.ActiveNetParams.Prefix)
	if err != nil {
		return 0, appmessage.RPCErrorf("Couldn't decode address '%s': %s", addressString, err)
	}

	scriptPublicKey, err := txscript.PayToAddrScript(address)
	if err != nil {
		return 0, appmessage.RPCErrorf("Could not create a scriptPublicKey for address '%s': %s", addressString, err)
	}
	utxoOutpointEntryPairs, err := context.UTXOIndex.UTXOs(scriptPublicKey)
	if err != nil {
		return 0, err
	}

	balance := uint64(0)
	for _, utxoOutpointEntryPair := range utxoOutpointEntryPairs {
		balance += utxoOutpointEntryPair.Amount()
	}
	return balance, nil
}
