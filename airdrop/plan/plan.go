package plan

import (
	"github.com/binance-chain/chain-tooling/airdrop/config"
	"github.com/binance-chain/go-sdk/keys"
	"log"
	"time"
)

type ExecuteContext struct {
	Config *config.Conf

	KeyManager keys.KeyManager

	EachAmount int64
	Sender     string

	Tasks []*ExecuteTask

	StartTime    time.Time
	CompleteTime time.Time
}

type ExecuteTask struct {
	Token          string
	Receivers      []string
	EachAmount     int64
	TxHash         string
	Affirmed       bool
	Exception      error
	ValidException error
}

type PlanMaker struct {
	Context *ExecuteContext
}

func (pm *PlanMaker) InitializeContext() error {

	var context = ExecuteContext{}
	context.Config = config.RawConf

	km, error := keys.NewMnemonicKeyManager(context.Config.Mnemonic)

	if error != nil {
		log.Fatal(error)
		return error
	}

	context.KeyManager = km

	context.Sender = km.GetAddr().String()
	context.EachAmount = context.Config.Amount / int64(context.Config.ReceiversCount)

	pm.Context = &context
	return nil
}

func (pm *PlanMaker) MakeExecutePlan() error {
	var context = pm.Context

	batchSize := context.Config.BatchSize
	taskCount := ((context.Config.ReceiversCount - 1) / batchSize) + 1

	context.Tasks = make([]*ExecuteTask, taskCount)

	for index, task := range context.Tasks {
		task = &ExecuteTask{}

		task.Token = context.Config.Token
		task.EachAmount = context.EachAmount

		var start = index * batchSize
		var end = (index + 1) * batchSize

		if end >= context.Config.ReceiversCount {
			end = context.Config.ReceiversCount
		}

		task.Receivers = context.Config.Receivers[start:end]
		context.Tasks[index] = task
	}

	return nil
}
