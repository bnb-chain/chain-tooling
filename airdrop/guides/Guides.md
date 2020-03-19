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

You will pay on mainnet/testnet

```
0.0003 BNB * 5 = 0.0015 BNB
```

Note that we have two different binaries: `bairdrop` is used for mainnet and `tbairdrop` for testnet. 



## Install

We have a installer script (install.sh) that takes care of chain directory setup. This uses the following defaults:

* Home folder in ~/.bnbchaind
* Client executables stored in /usr/local/bin (i.e. bairdrop or tbairdrop)
```
# One-line install
sh <(wget -qO- https://raw.githubusercontent.com/onggunhao/node-binary/master/install.sh)
```
## Verify Transfer Transaction

To confirm that your transaction went through, you can use [mainet explorer]() or [testnet explorer]() to verify the airdrop result. 

 As you can see from the example output, this transaction is executed at block height 1412766 and you could read about the details. Double check with blockchain explorer if you interact with the network through a full-node.