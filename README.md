# Binance Chain Tool Box

# Airdrop
 __Disclaimer: Airdrop is very error prone and can cause big financial loss. No guarantee is provided to use this tool. Please do test it on testnet and use it carefully.__

## Installation

If you don't want to build from the source code, just clone this repository there is already a runnable file - **airdrop** you can make up the configurations and directly run against the runnable file.

```bash
git clone https://github.com/binance-chain/chain-tooling.git
cd chain-tooling/airdrop
    
// do something to the configuration file
    
./airdrop
```

If you want build from the source code or make some changes to the code ,you can firstly get the code using `go get`, change the code and then build the tool use `go install`

```bash
go get github.com/binance-chain/chain-tooling
cd $GOPATH/src/github.com/binance-chain/chain-tooling/airdrop

// make some changes to the code

go install

```

## Configuration
There is a configuration file called **airdrop.conf**, you can make you configurations there in the file.

- **env** : could be testnet or prod
- **token** : the token name you wish to deliver
- **amount** : the amount you want to deliver, notice the amount here should be real amount multiplied by 10^8 , and it should always be an integer with no fractional part
- **batchsize** : the airdrop task would be divided in to several batches, the batchsize is the number of addresses in one batch. We suggest 500~1000 would be a good choice as that a small batchsize would result in more batches which would waste time while a large batchsize would cause failure delivery
- **mnemonic** : paste the mnemonic of your sender account here 
- **receivers** : paste the addresses you want to deliver the tokens here, addresses should be separated by comma

## Run

run is easy with 

```bash
    ./airdrop
```

## Execution report

After execute the airdrop task, there should be a report file to record the task.

# Token APP

> :no_entry: [DEPRECATED] Active at https://github.com/binance-chain/chain-tooling/token-app 

:point_right:Please follow the latest guideline on how to manage your BEP2 tokens and submit listing proposals: 

* [English Version](https://community.binance.org/topic/2487)
* [Chinese Version](https://community.binance.org/topic/2488/)

 __Disclaimer: When you try to issue asset on Binance Chain, you should be extra careful with your command. It is recommended to use a hardware wallet to sign your transactions. Please do test it on testnet and use it carefully.__

## Install Git LFS

~~Git Large File Storage (LFS) replaces large files such as audio samples, videos, datasets, and graphics with text pointers inside Git, while storing the file contents on a remote server like GitHub.com or GitHub Enterprise.~~

~~Please go to https://git-lfs.github.com/ and install git lfs.~~

~~Download Binary with Git LFS:~~
~~git lfs clone https://github.com/binance-chain/chain-tooling.git~~

## Download the zip
~~Only version for MacOS is released now.~~ 

~~1. Download the installer and unzip the file~~

~~2. Copy the app to your Application folder and double-click on the icon~~

## How to Use

~~Please read this [guide](./token-app/binance-chain-gui.pdf) carefully.~~


