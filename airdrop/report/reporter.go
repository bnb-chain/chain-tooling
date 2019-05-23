package report

import (
	"bytes"
	"fmt"
	"github.com/binance-chain/chain-tooling/airdrop/plan"
	"github.com/landoop/tableprinter"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Report(context *plan.ExecuteContext) error {

	var conf = context.Config

	jobReport := fmt.Sprintf(_jobReportTemplate, conf.Env, conf.Token, conf.Amount, conf.ReceiversCount, conf.BatchSize, conf.ReportFile, context.Sender, context.StartTime.Format(time.RFC3339), context.CompleteTime.Format(time.RFC3339))

	executeReport, error := executeReport(context)
	if error != nil {
		log.Fatal(error)
	}

	reportFile := context.Config.ReportFile

	return output(jobReport, executeReport, reportFile)
}

const _jobReportTemplate = "env:%s\r\ntoken:%s\r\namount:%d\r\nreceivercount:%d\r\nbatchsize:%d\r\nreportfile:%s\r\nsender:%s\r\nstarttime:%s\r\ncompletetime:%s\r\n"

type executeTask struct {
	Token          string `header:"token"`
	EachAmount     string `header:"individual amount"`
	TxHash         string `header:"transaction hash"`
	Affirmed       string `header:"affirmed"`
	Exception      string `header:"execute exception"`
	ValidException string `header:"validate exception"`
	Receivers      string `header:"receiver list"`
}

func executeReport(context *plan.ExecuteContext) (string, error) {

	executeTaskList := make([]executeTask, len(context.Tasks))

	for index, task := range context.Tasks {
		executeTaskList[index].Token = task.Token
		executeTaskList[index].Receivers = strings.Join(task.Receivers, ",")
		executeTaskList[index].EachAmount = strconv.FormatInt(task.EachAmount, 10)
		executeTaskList[index].TxHash = task.TxHash
		executeTaskList[index].Affirmed = strconv.FormatBool(task.Affirmed)
		if task.Exception != nil {
			executeTaskList[index].Exception = task.Exception.Error()
		}
		if task.ValidException != nil {
			executeTaskList[index].ValidException = task.ValidException.Error()
		}
	}

	buf := new(bytes.Buffer)
	printer := tableprinter.New(buf)
	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	printer.NewLine = "\r\n"
	printer.Print(executeTaskList)
	return buf.String(), nil
}

func output(report1 string, report2 string, outputFile string) error {

	if len(report1) > 0 {
		log.Print(report1)
	}

	if len(report2) > 0 {
		log.Print("\n" + report2)
	}

	f, err := os.Create(outputFile)
	defer f.Close()

	if err != nil {
		return err
	}

	_, err2 := f.WriteString(report1)
	if err2 != nil {
		return err2
	}

	f.WriteString("\r\n")

	_, err3 := f.WriteString(report2)

	f.Sync()

	return err3
}
