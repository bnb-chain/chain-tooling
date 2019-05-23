package config

import (
	"errors"
	"github.com/binance-chain/go-sdk/common/types"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type Conf struct {
	Env            string
	BaseUrl        string
	Token          string
	Amount         int64
	Receivers      []string
	ReceiversCount int
	Sender         string
	Mnemonic       string
	BatchSize      int
	BatchInterval  int
	ReportFile     string
	Network        types.ChainNetwork
}

var RawConf *Conf

func init() {

	rawConf, error := parseConfig("airdrop.conf")
	if error != nil {
		log.Fatal(error)
		panic(error)
	}
	error2 := validateConfig(rawConf)

	if error2 != nil {
		log.Fatal(error2)
		panic(error2)
	}

	logConf(rawConf)

	RawConf = rawConf
}

func logConf(conf *Conf) {
	log.Println("env:" + conf.Env)
	log.Println("token:" + conf.Token)
	log.Println("amount:" + strconv.FormatInt(conf.Amount, 10))
	log.Println("batch size:" + strconv.Itoa(conf.BatchSize))
	log.Println("batch interval (s):" + strconv.Itoa(conf.BatchInterval))
	log.Println("receivers count:" + strconv.Itoa(len(conf.Receivers)))
}

func parseConfig(configFile string) (*Conf, error) {
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	configs := strings.Split(string(content), "\n")
	result := Conf{}
	for _, config := range configs {

		kv := strings.Split(config, "=")
		if strings.HasPrefix(config, "#") || len(kv) < 2 {
			continue
		}

		key := kv[0]
		value := kv[1]
		switch key {
		case "env":
			result.Env = value
		case "token":
			result.Token = value
		case "amount":
			result.Amount, _ = strconv.ParseInt(value, 10, 64)
		case "mnemonic":
			result.Mnemonic = value
		case "receivers":
			if len(value) > 0 {
				result.Receivers = strings.Split(value, ",")
				result.ReceiversCount = len(result.Receivers)
			}
		case "batchsize":
			result.BatchSize, _ = strconv.Atoi(value)
		case "batchinterval":
			result.BatchInterval, _ = strconv.Atoi(value)
		case "reportfile":
			result.ReportFile = value
		}
	}
	return &result, nil
}

func validateConfig(conf *Conf) error {

	conf.Env = strings.ToLower(conf.Env)
	if conf.Env == "testnet" {
		conf.BaseUrl = "testnet-dex.binance.org"
		conf.Network = types.TestNetwork
		types.Network = types.TestNetwork
	} else if conf.Env == "prod" {
		conf.BaseUrl = "dex.binance.org"
		conf.Network = types.ProdNetwork
	} else {
		return errors.New("env must be testnet or prod ")
	}

	if strings.TrimSpace(conf.Token) == "" {
		return errors.New("token must be specified ")
	}

	if conf.Amount <= 0 {
		return errors.New("amount must be greater than zero ")
	}

	if conf.ReceiversCount == 0 {
		return errors.New("there must be at least one receiver ")
	}

	if conf.BatchSize <= 0 || conf.BatchSize > 1000 {
		return errors.New("batchsize must be greater than 0 and less than 1001")
	}

	if len(conf.ReportFile) == 0 {
		conf.ReportFile = "report." + strconv.FormatInt(time.Now().UnixNano()/int64(1000000000), 10)
	}

	if conf.BatchInterval <= 0 {
		conf.BatchInterval = 5
	}

	return nil
}
