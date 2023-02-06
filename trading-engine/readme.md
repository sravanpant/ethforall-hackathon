
# Trading Engine

This folder contains whole order book matching algorithm which will be executed by the oracles. All the additional files for the oracles is added.

Geth was used for many things in this repo.




## Files available

Description of the important functions available in the smart contract

```bash
  main.go - This file executes the orderbook matching algorithms exposing the functions through a local REST API.
  Also it calls the smart contract to prove the order.
```

```bash
  logs.go - This file listens to different logs from the smart contracts and acts based on that.
  For addask and addbid, it calls the main.go file.
  For orderProved event, it creates a deal with the respected SP.
```

```bash
  eth.go - This file imitates an interactionw with the smart contract.
```



