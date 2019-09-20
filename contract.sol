pragma solidity ^0.5.0;

contract Ownable {
  address payable private _owner;

  event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

  modifier onlyOwner() {
    require(msg.sender == _owner, "Forbidden");
    _;
  }

  constructor() public {
    _owner = msg.sender;
  }

  function owner() public view returns (address payable) {
    return _owner;
  }

  function transferOwnership(address payable newOwner) public onlyOwner {
    require(newOwner != address(0), "Non-zero address required.");
    emit OwnershipTransferred(_owner, newOwner);
    _owner = newOwner;
  }
}

contract merkleRootHashOracle is Ownable {
    
    mapping(bytes32 => bytes32[])RootHashContents;

    event merkleRootHashReceived(uint256 date, bytes32 indexed hash);

    function registerMerkleRootHash(bytes32 hash, bytes32[] memory contents) public onlyOwner {
        RootHashContents[hash] = contents;
        emit merkleRootHashReceived(now,hash);
    }
    
    function getRootHashContents(bytes32 hash) public view returns (bytes32[] memory) {
        return RootHashContents[hash];
    }
}
