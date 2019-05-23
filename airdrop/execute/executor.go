package execute

import (
	"fmt"
	"github.com/binance-chain/chain-tooling/airdrop/plan"
	"github.com/binance-chain/go-sdk/client"
	"github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/types/msg"
	"log"
	"strings"
	"time"
)

type Executor struct {
	context *plan.ExecuteContext
}

func NewExecutor(context *plan.ExecuteContext) Executor {
	return Executor{context: context}
}

func (ex *Executor) Execute() error {

	var context = ex.context
	context.StartTime = time.Now()

	client, err := ex.getDexClient()
	if err != nil {
		return err
	}

	for _, task := range context.Tasks {
		time.Sleep(2 * time.Second)
		receivers := task.Receivers
		var transfers = make([]msg.Transfer, len(receivers))

		for index, receiver := range receivers {
			receiverAddr, err := types.AccAddressFromBech32(receiver)
			if err != nil {
				task.Exception = err
				break
			}
			transfers[index].ToAddr = receiverAddr
			transfers[index].Coins = types.Coins{types.Coin{task.Token, task.EachAmount}}
		}

		if task.Exception != nil {
			continue
		}

		log.Println(fmt.Sprintf("Trying to send %d(each) %s to %s", task.EachAmount, task.Token, strings.Join(task.Receivers, ",")))
		result, err := client.SendToken(transfers, true)

		if err == nil {
			task.TxHash = result.Hash
			log.Println(fmt.Sprintf("Complete with tx %s", result.Hash))
		} else {
			task.Exception = err
			log.Println(fmt.Sprintf("Failed with exception %s", err.Error()))
		}
	}
	context.CompleteTime = time.Now()

	return nil
}

func (ex *Executor) Validate() error {

	var context = ex.context
	client, err := ex.getDexClient()

	if err != nil {
		return err
	}

	for _, task := range context.Tasks {
		time.Sleep(1 * time.Second)
		if len(task.TxHash) > 0 {
			txResult, error := client.GetTx(task.TxHash)

			if error != nil {
				task.ValidException = error
			}

			if txResult != nil && len(txResult.Hash) > 0 {
				task.Affirmed = true
			}

		}
	}

	return nil
}

func (ex *Executor) getDexClient() (client.DexClient, error) {
	return client.NewDexClient(ex.context.Config.BaseUrl, ex.context.Config.Network, ex.context.KeyManager)
}
