# Guides

**Disclaimer: Airdrop is very error prone and can cause big financial loss. Please do test it on testnet and use it carefully.**

Binance Chain airdrop tool helps you send tokens to multiple addresses with the help of [Multi-send transaction](https://docs.binance.org/transfer.html#multi-send). The benefit of doing so is that you can get a 20% discount of transfer fees. The total transaction fee is 0.0003 BNB per token per address. For example, if you send 3 BNB  to 3 different addresses.

```
[
   {
      "to":"bnb1g5p04snezgpky203fq6da9qyjsy2k9kzr5yuhl",
      "amount":"100000000:BNB"
   },
   {
      "to":"bnb1l86xty0m55ryct9pnypz6chvtsmpyewmhrqwxw",
      "amount":"100000000:BNB"
   },
   {
      "to":"bnb1l86xty0maxdgst9pnypz6chvtsmpydkjflfioe",
      "amount":"100000000:BNB"
   }
]
```

You will pay on mainnet/testnet `0.0003 BNB * 3 = 0.0009 BNB`, rather than 0.000375BNB*3

## Install

We have a installer script (install.sh) that takes care of chain directory setup. This uses the following defaults:

* Home folder in ~/.bairdrop
* Client executables stored in /usr/local/bin (i.e. bairdrop or tbairdrop)
```
# One-line install
sh <(wget -qO- )
```
Note that we have two different binaries: `bairdrop` is used for mainnet and `tbairdrop` for testnet.

To confirm the installament is successful:
```
$ bairdrop
BlockChain Airdrop

Usage:
  bairdrop [command]

Available Commands:
  keys        Add or view local private keys
  run         Run airdrop
  help        Help about any command

Flags:
  -h, --help            help for bairdrop
      --home string     directory for config and data (default "/Users/fletcher/.tbairdrop")
  -o, --output string   Output format (text|json) (default "text")

Use "tbairdrop [command] --help" for more information about a command.
```

### Prepare Distribution File
First of all, we need to prepare the CSV file of transfer details, including target accounts and the amount for the individual. It has its specific format which looks like the below:

![transfer_details.csv](./assets/transfer_example.png?raw=true "example")

>**Note that there is no Header in this CSV. From left to right are `target address`,`amount`,`asset`. For `amount`, the amout needs to be boosted by 1e8 for decimal part. For instance,1st row means transfer tbnb1rtzy6szuyzcj4amfn6uarvne8a5epxrdc28nhr 0.1BNB.**

Save it as CSV format

![save](./assets/transfer_save.png?raw=true "save")

The airdrop tool processes the transfer in batches if the number of reeceiver accounts exceeds 300. It creates every 300 rows as a batch, with an interval of 2s between each batch.

### Usage
There are two subcommands:
* keys: for manage accounts as source of airdrop
* run: execute airdrop transactions

#### Keys
Before distributing, you have to configure the source account for this distribution . This subcommand works the sam as [keys](https://docs.binance.org/keys.html) subcommand of `bnbcli`. You can add a new key or import a key from seed phrase. Please note that the default home folder is `~/.bairdrop`

#### Run
Then we can start distributing with subcommand `run`.

**Parameters**
* file:  File of airdrop details
* from: Name or address of private key with which to sign multi-send transaction. You can see the names with `bairdrop keys list`
* chain-id:  Chain ID of blockchain
* memo:  Memo to send along with transaction
* node: fullnode host:port,  default "tcp://localhost:26657"
* home: directory for config and data (default "/Users/fletcher/.tbairdrop")


### Example

Now, let us do a demo with `tbairdrop`. First of all, make sure the you have the right account locally. In this example, we use the account with name of `fromAcc`.

```
  $ tbairdrop keys list
  NAME:	TYPE:	ADDRESS:						PUBKEY:
  fromAcc	local	tbnb1m38ds8d69kwd8a4uaz5fm3hmvh94wk5gfeszxn	bnbp1addwnpepqdqlls9gxnqujgkdpty6nluxtc6cuurqe7fhe8jmp87exwzq7s5vkfxcvxk
```
Secondly, you need to prepare a CSV file, eg . `transfer_details.csv`

At last, you can execute the `airdrop` with the transfer file specified.
```
$ tbairdrop run --file transfer_details.csv --chain-id Binance-Chain-Nile  --node data-seed-pre-0-s1.binance.org:80 --from fromAcc
```

Following is the log of this execution, you can get `txhash`of this airdrop transaction.

```
1.  ==>Start to run with file: transfer_details.csv
2.  ==>Start batch 1(from tbnb1rtzy6szuyzcj4amfn6uarvne8a5epxrdc28nhr to tbnb1tyrc4usqp52ne60y2qnta4jk997e79tzcmvlcm)
3.  Password to sign with 'fromAcc':
4.  ==>Transaction hash: BDB452AE09AB9961FD77109DB6DB36559C64415D37F7CEDEF7027FEC43D1130B, sending...
5.  ==>Sending completed, committed at block 69726953 (tx hash: BDB452AE09AB9961FD77109DB6DB36559C64415D37F7CEDEF7027FEC43D1130B)

6.  ==>Start batch 2(from tbnb1vze2xyajsl3dpumkewuz0jdschstnyr6wtyctz to tbnb15xemc2fk9cvxewa87d0js3ypq7aawrku24l7px)
7.  ==>Transaction hash: C66DFFCF6DE1146468CE0947AB5E554E7FF1C87942969F20167692228F8B37A1, sending...
8.  ==>Sending completed, committed at block 69726961 (tx hash: C66DFFCF6DE1146468CE0947AB5E554E7FF1C87942969F20167692228F8B37A1)

9.  ==>Start batch 3(from tbnb17jql0796cjuxzxwmjfdm779gd2306lpexemx27 to tbnb1qgtdt7y062mk66vgu37e0lgamngwgj5asce74f)
10. ==>Transaction hash: F474F3ABF5E770EE1291821C07234236403028767BD912011F5CAB42228D8725, sending...
11. ==>Sending completed, committed at block 69726969 (tx hash: F474F3ABF5E770EE1291821C07234236403028767BD912011F5CAB42228D8725)

12. ==>Start batch 4(from tbnb1zgczyzlxyk34rwqdzrpxez35k52qy7uyl7lp8q to tbnb1zgczyzlxyk34rwqdzrpxez35k52qy7uyl7lp8q)
13. ==>Transaction hash: DD1C8F9B6A7C0F526CFA41A6D221D28AC26DD196D5AC109B7F72655B0E726E63, sending...
14. ==>Sending completed, committed at block 69726979 (tx hash: DD1C8F9B6A7C0F526CFA41A6D221D28AC26DD196D5AC109B7F72655B0E726E63)`
```
## Verify Transfer Transaction

To confirm that your transaction went through, you can use [mainet explorer](https://explorer.binance.org/) or [testnet explorer](https://testnet-explorer.binance.org/) for searching airdrop transacions by `txhash` to verify the airdrop result.

### Errors
It is worth noting that if for any reason it fails during a batch execution, you need to double confirm that whether this transfer is on chain based on tx hash the console prints. And you should remove the records of the successful transfer from the CSV to prevent repeated transfer, and execute it again.

In the previous  example, 
+ if it crashes after the console prints line: 8 or 9, it means batch1 and batch2 are transferred successfully and  batch3 is not executed yet. We should remove the related rows from the CSV, and re-run this file with the command.
+  if it crashes after the console prints line 10, it means batch1 and batch2 are transferred successfully. But it is not sure if batch3 is successfully executed, we should take a look with the transaction hash *F474F3ABF5E770EE1291821C07234236403028767BD912011F5CAB42228D8725* offline. For instance, you can check it on the explorer `https://testnet-explorer.binance.org/tx/F474F3ABF5E770EE1291821C07234236403028767BD912011F5CAB42228D8725`.

