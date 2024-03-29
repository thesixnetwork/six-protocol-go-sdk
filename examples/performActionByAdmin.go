package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/thesixnetwork/six-protocol-go-sdk/api"
	nftmngrtypes "github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func main() {
	nodeURL := "YOUR NODE URL"
	armor := `
-----BEGIN TENDERMINT PRIVATE KEY-----
YOUR KEY DETAIL
-----END TENDERMINT PRIVATE KEY-----
	`
	passphrase := "YOUR PASSPHARSE"
	chainID := "YOUR CHAIN ID"

	// Create a new API client
	gasPrice := "1.25usix" // default "1.25usix"
	clientOptions := &api.ClientOptions{
		BroadcastMode: "async", // default "block"
		GasPrices:     &gasPrice,
	}
	client, err := api.NewClient(
		nodeURL,
		armor,
		passphrase,
		chainID,
		clientOptions,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Perform Action By Admin
	refID := uuid.New()
	msg := &nftmngrtypes.MsgPerformActionByAdmin{
		Creator:       client.ConnectedAddress,
		NftSchemaCode: "six-protocol.develop_v220",
		TokenId:       "1",
		Action:        "test_read_nft",
		RefId:         refID.String(),
		Parameters:    []*nftmngrtypes.ActionParameter{},
	}

	txResponse, err := client.GenerateOrBroadcastTx(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("txResponse code", txResponse.Code)
	fmt.Println("txResponse hash", txResponse.TxHash)
}
