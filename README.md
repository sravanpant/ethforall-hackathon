
# DataSuite FEVM Hackathon

We planned out the project to be a whole data suite as the name suggests. It consists of an order book-based mechanism where we match orders from clients and SPs and store the data. As implementing the order book on solidity is very inefficient, we created an oracle mechanism where the order matching happens off-chain in the oracles. The workflow is the client comes to the website and fills up a form with the data and the amount he is ready to pay. During the process, data is transferred to one of the oracles, and the oracles will create the CID of the data. And then all these things will be packed up and addbid function is called in the smart contract. Similarly, the SP owner will come and verify himself as the owner of the SP and fills up a form with the amount. As the form is filled, the data is passed to the smart contract. The oracles track the logs of the contract and match order when possible and ask the smart contracts to prove it. Once the smart contract gets enough number of votes, it asks the oracle to pass on the data to the SP acting as a lotus lite node. Once the deal is published, the SP owner can come and get the bounty from the contract. After this, the data gets listed in the marketplace if the data owner wants where anyone can come and buy the data. The bridge will help to pass the requests from other chains about particular data listed in the marketplace.

There are multiple parts of the project :- 


## Parts of the project

#### smart-contract folder

This particular folder contains the hardhat toolkit using which we developed smart contract

#### trading-engine folder

In this folder, we have the trading engine implemenation in Golang along with the other important parts of the project.

#### storage-provider folder

This folder contains all the lotus related files needed for creating a lite node or a miner in Hyperspace testnet.

#### frontend folder

This folder contains frontend part of the project implemented in NextJS.

## Whole workflow

We were only able to complete the orderbook part of the whole ecosystem for this hackathon. We have plans in future to make the marketplace along with the data bridge.

[Click to check the workflow](https://drive.google.com/file/d/15ewSLzcvHGoMkC1Vspu4SjUalrAeW9ZS/view)

Youtube link :- https://youtu.be/Bx3oI2cGf8Q
