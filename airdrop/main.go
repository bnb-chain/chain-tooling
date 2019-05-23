package main

import (
	_ "github.com/binance-chain/chain-tooling/airdrop/config"
	"github.com/binance-chain/chain-tooling/airdrop/execute"
	"github.com/binance-chain/chain-tooling/airdrop/plan"
	"github.com/binance-chain/chain-tooling/airdrop/report"
	"log"
)

func main() {

	log.Println("Start to init context ... ")
	var planMaker = plan.PlanMaker{}
	error := planMaker.InitializeContext()
	if error != nil {
		log.Fatal(error)
		panic(error)
	}

	log.Println("Start to make execution plan ...")
	error2 := planMaker.MakeExecutePlan()
	if error2 != nil {
		log.Fatal(error2)
		panic(error)
	}

	log.Println("Start to execute the plan ...")
	executor := execute.NewExecutor(planMaker.Context)
	error3 := executor.Execute()
	if error3 != nil {
		log.Fatal(error3)
	}

	log.Println("Start to validate the execute result ...")
	error4 := executor.Validate()
	if error4 != nil {
		log.Fatal(error4)
	}

	log.Println("Start to generate execution report...")
	report.Report(planMaker.Context)
}
