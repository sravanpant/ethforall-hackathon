// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract OrderBook{

    struct Order {
        string orderId;
        address bidder;
        address asker;
        uint256 amountFixed;
        uint256 sizeFixed;
    }

    struct Bid{
        address bidder;
        uint256 bidAmount;
        uint256 bidSize;
    }

    struct Ask{
        address asker;
        uint256 askAmount;
        uint256 askSize;
    }

    Order[] public orders;

    event NewBidEvent(address bidder, uint256 bidAmount, uint256 bidSize);
    event NewAskEvent(address asker, uint256 askAmount, uint256 askSize);
    event orderCreated(address bidder, address asker, uint256 amount, uint256 size);

    function addBid(uint256 _bidAmount, uint256 _bidSize) external{
        emit NewBidEvent(msg.sender, _bidAmount, _bidSize);
    }

    function addAsk(uint256 _askAmount, uint256 _askSize) external{
        emit NewAskEvent(msg.sender, _askAmount, _askSize);
    }

    function addOrder(string memory _orderId, address _bidder, address _asker, uint256 _amountFixed, uint256 _sizeFixed) external{
        orders.push(Order(_orderId, _bidder, _asker, _amountFixed, _sizeFixed));
    }

}