//SPDX-License-Identifier: MIT

pragma solidity >=0.6.0 <0.9.0;

abstract contract ERC20 {
    string public constant name = "";
    string public constant symbol = "";
    uint8 public constant decimals = 0;

    function totalSupply() public virtual returns (uint);
    function balanceOf(address tokenOwner) public virtual returns (uint balance);
    function allowance(address tokenOwner, address spender) public virtual returns (uint remaining);
    function transfer(address to, uint tokens) public virtual returns (bool success);
    function approve(address spender, uint tokens) public virtual returns (bool success);
    function transferFrom(address from, address to, uint tokens) public virtual returns (bool success);

    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}
