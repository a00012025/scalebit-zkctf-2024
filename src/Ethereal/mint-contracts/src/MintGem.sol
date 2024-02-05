pragma solidity ^0.8.15;

import "forge-std/console.sol";
import "./KZGVerifier.sol";

contract Mint is Verifier {
    struct DepositInfo {
        bool registered;
        uint256 nonce;
        Pairing.G1Point commitment;
        Pairing.G1Point soulBox;
        uint256[] collectedY;
        bool escaped0;
        bool escaped1;
    }

    // function is_sovled(DepositInfo info) returns (bool) {
    //     if(DepositInfo.escaped0 && DepositInfo escaped1)
    //     {return true;}
    //     else
    //     {return false;}
    // }
    

    function getCollectedY(address user) public view returns (uint256[] memory) {
        return deposits[user].collectedY;
    }

    mapping(address => DepositInfo) public deposits;
    event Deposit(Pairing.G1Point indexed soulbox, Pairing.G1Point commitment,uint256 timestamp);
    event Minted(address indexed user, uint256 nonce);
    event Pwned(uint256 indexed privateKey);

    function register(
        Pairing.G1Point calldata _commitment, 
        Pairing.G1Point calldata _proof, 
        Pairing.G1Point calldata _soulBox
        ) external payable {
        address sender = msg.sender;
        require(!deposits[sender].registered, "User registered");
        require(verifySoulBox(_commitment, _proof, _soulBox), "Invalid  soul box proof");
        deposits[sender] = DepositInfo({
            registered: true, 
            nonce: 1, 
            commitment: _commitment,
            soulBox: _soulBox,
            collectedY: new uint256[](0),
            escaped0: false,
            escaped1: false
        });
        emit Deposit(_soulBox, _commitment, block.timestamp);
    }

    function getNonce(address _user) public view returns (uint256) {
        return deposits[_user].nonce;
    }

    function mint(
        Pairing.G1Point calldata _proof, 
        uint256 _value
    ) external {
        address sender = msg.sender;
        uint256 nonce = getNonce(sender);
        require(deposits[sender].registered, "User registered");
        require(verify(deposits[sender].commitment, _proof, nonce, _value), "Invalid mint proof");
        deposits[sender].collectedY.push(_value);

        uint256[] memory xArray = new uint256[](nonce);
        for (uint256 i = 0; i < nonce; i++) {
            xArray[i] = i+1;
        }
        uint256[] memory wArray = getBarycentricWeights(xArray);
        // recover private_key = f(0) by lagrange interpolation
        uint256 soul = evaluateBarycentricPolynomial(0, deposits[sender].collectedY, wArray);
        Pairing.G1Point memory soulBox = Pairing.mulScalar(SRS_G1_0, soul);
        require(soulBox.X != deposits[sender].soulBox.X, "Soul dispersed");
        
        if(nonce == 10){
            deposits[sender].escaped0 = true;
        }
        if(nonce == 20){
            deposits[sender].escaped1 = true;
        }
        
        emit Minted(sender, deposits[sender].nonce);
        deposits[sender].nonce += 1;
    }

    function replay() external {
        address sender = msg.sender;
        require(deposits[sender].registered, "User not registered");
        delete(deposits[sender]);
        delete(deposits[sender].collectedY);
    }
}