// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "erc721a/contracts/extensions/ERC721AQueryable.sol";

contract SampleERC721 is ERC721AQueryable {

    constructor(string memory tokenName, string memory tokenSymbol) ERC721A(tokenName, tokenSymbol) {}

    function mint() public {
        _safeMint(msg.sender, 1);
    }

    function mint(uint256 size) public {
        _safeMint(msg.sender,size);
    }
}