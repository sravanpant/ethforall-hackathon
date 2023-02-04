// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import { MinerAPI } from "@zondax/filecoin-solidity/contracts/v0.8/MinerAPI.sol";
import { MinerTypes } from "@zondax/filecoin-solidity/contracts/v0.8/types/MinerTypes.sol";

contract OrderBook{

    struct Order {
        bytes orderId;
        string dataCID;
        address bidder;
        address asker;
        uint256 amountFixed;
        uint256 sizeFixed;
        address oracleAddress;
        uint256 confirmations;
        bool confirmed;
        bool bountyClaimed;
    }

    struct Bid{
        uint256 bidId;
        address bidder;
        uint256 bidAmount;
        uint256 bidSize;
        address oracleAddress;
        string dataCID;
    }

    struct Ask{
        uint256 askId;
        address asker;
        uint256 askAmount;
        uint256 askSize;
    }

    
    mapping (uint256 => Bid) bids;
    mapping (uint256 => Ask) asks;
    mapping (bytes => Order) orders;
    address owner;
    address[] oracles;
    string[] oraclesIP;
    uint256 bidId = 0;
    uint256 askId = 0;

    modifier onlyOwner(){
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    modifier onlyOracle(){
        uint256 count = 0;
        for(uint256 i=0; i<oracles.length; i++){
            if(msg.sender == oracles[i]){
                count++;
                break;
            }
        }
        require(count==1, "Oracles can only call this function");
        _;
    }

    event NewBidEvent(uint256 bidId, address bidder, uint256 bidAmount, uint256 bidSize);
    event NewAskEvent(uint256 askId, address asker, uint256 askAmount, uint256 askSize);
    event orderCreatedEvent(bytes orderId, uint256 bidId, address bidder, uint256 askId, address asker, address oracleAddress, string dataCID);

    constructor(address[] memory _oracles, string [] memory _oraclesIP){
        owner = msg.sender;
        oracles = _oracles;
        oraclesIP = _oraclesIP;
    }

    function registerOracle(address _oracle, string memory _oracleIP) external onlyOwner{
        oracles.push(_oracle);
        oraclesIP.push(_oracleIP);
    }


    function addBid(address _oracleAddress, uint256 _bidAmount, uint256 _bidSize, string memory _dataCID) external payable {
        require(msg.value >= _bidAmount);
        bids[bidId] = Bid(bidId, msg.sender, _bidAmount, _bidSize, _oracleAddress, _dataCID);
        bidId++;
        emit NewBidEvent(bidId-1, msg.sender, _bidAmount, _bidSize);
    }

    function addAsk(bytes memory _target, uint256 _askAmount, uint256 _askSize) external{
        require(keccak256(MinerAPI.getOwner(_target).owner) == keccak256(abi.encodePacked(msg.sender)));
        asks[askId] = Ask(askId, msg.sender, _askAmount, _askSize);
        askId++;
        emit NewAskEvent(askId-1, msg.sender, _askAmount, _askSize);
    }

    function proveOrder(bytes memory _orderId, uint256 _bidId, uint256 _askId, uint256 _amountFixed, uint256 _sizeFixed) external onlyOracle{
        if(orders[_orderId].amountFixed != 0){
            orders[_orderId].confirmations++;
            if(orders[_orderId].confirmations > oracles.length/2 && orders[_orderId].confirmed==false){
                orders[_orderId].confirmed = true;
                emit orderCreatedEvent(_orderId, _bidId, bids[_bidId].bidder, askId, asks[_askId].asker, bids[_bidId].oracleAddress, bids[_bidId].dataCID);
            }
        }else{
            orders[_orderId] = Order(_orderId, bids[_bidId].dataCID, bids[_bidId].bidder, asks[_askId].asker, _amountFixed, _sizeFixed, bids[_bidId].oracleAddress, 1, false, false);
        }
    }

    function askBounty(bytes memory _orderId) external{
        require(orders[_orderId].bountyClaimed == false);
        orders[_orderId].bountyClaimed = true;
        
    }

}