pragma solidity >=0.5.1;

contract HelloWorld{

    string public Myame = "kzhang";

    function changeName(string memory newName) public{
        Myame = newName;
    }
}
