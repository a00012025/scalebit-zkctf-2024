import hashlib
import json
import requests
from pwn import *
from web3 import Web3
import time
import os

RPC_URL = "http://47.76.89.7:8545/"
FAUCET = "http://47.76.89.7:8080/api/claim"
CHALLENGE_IP = "47.76.89.7"
CHALLENGE_PORT = 20011
CHAINID = 19875
AIRDROP = False

provider = Web3(Web3.HTTPProvider(RPC_URL))

# $ cast wallet new
private_key = "0x2abc57ce76f7b2117b8ff48d02617bbaf1a0c42f25540f7926cdffff0f621884"
player_address = "0x490c28B6E6880b30642753F03eD7c25Ef180E16e"


def wait_for_tx(provider, tx_hash):
    # eth.wait_for_transaction_receipt is available,
    # but manually implemented for in-progress message
    while True:
        try:
            receipt = provider.eth.get_transaction_receipt(tx_hash)
            print("Receipt:", receipt)
            print("Balance after tx:", provider.eth.get_balance(player_address))
            return receipt
        except Exception as e:
            if "not found" in str(e):
                print("Waiting for the transaction...")
            else:
                raise e
        time.sleep(3)


def build_param(provider):
    return {
        "gas": 5000000,
        "gasPrice": Web3.to_wei(100, "gwei"),
        "nonce": provider.eth.get_transaction_count(player_address),
    }


def transact_and_wait(provider, transaction, private_key):
    print("Sending:", transaction)
    signed_tx = provider.eth.account.sign_transaction(transaction, private_key)
    tx_hash = provider.eth.send_raw_transaction(signed_tx.rawTransaction)
    return wait_for_tx(provider, tx_hash)


if AIRDROP:
    print("[*] Get ether from faucet")
    tx_hash = requests.post(
        FAUCET, data={"address": player_address}).text.split(': ')[1]
    wait_for_tx(provider, tx_hash)

print("[*] Generate abi from source")
if not os.path.exists("out"):
    os.system(f"solc verifier.sol --abi -o out")

print("[*] Deploy challenge contract")
io = remote(CHALLENGE_IP, CHALLENGE_PORT)
io.sendlineafter(b"input your choice:", b"1")
io.recvuntil(b"deployer account: ")
deployer_address = io.recvline().strip().decode()
io.recvuntil(b"token: ")
token = io.recvline().strip()
print(f"{deployer_address=}")
print(f"{token=}")
io.close()

transact_and_wait(provider, {
    "chainId": CHAINID,
    "to": deployer_address,
    "value": provider.to_wei(0.01, "ether"),
    "gas": 5000000,
    "gasPrice": Web3.to_wei(100, "gwei"),
    "nonce": provider.eth.get_transaction_count(player_address)
}, private_key)

io = remote(CHALLENGE_IP, CHALLENGE_PORT)
io.sendlineafter(b"input your choice:", b"2")
io.sendlineafter(b"input your token:", token)
io.recvuntil(b"contract address: ")
contract_address = io.recvline().strip().decode()
print(f"{contract_address=}")

with open(f"out/Roundabout.abi") as f:
    abi = json.load(f)
    challenge = provider.eth.contract(contract_address, abi=abi)

print("[*] Compile circult")
os.system(f"circom Roundabout.circom --r1cs --wasm")

print("[*] Generate proof")
with open("input.json", 'w') as f:
    json.dump({"a": 19830713, "b": 2}, f)
os.system(f"node Roundabout_js/generate_witness.js Roundabout_js/Roundabout.wasm input.json witness.wtns")
os.system(f"snarkjs g16p Roundabout_groth16.zkey witness.wtns proof.json public.json")
os.system("snarkjs generatecall > calldata.txt")

print("[*] Submit proof")
with open("calldata.txt") as f:
    calldata = f.read().strip().replace("][", "], [")
pa, pb, pc, public_inputs = eval(calldata)
pa = [int(x, 0) for x in pa]
pb = [[int(x, 0) for x in line] for line in pb]
pc = [int(x, 0) for x in pc]
public_inputs = [int(x, 0) for x in public_inputs]


function = challenge.functions.verify(pa, pb, pc, public_inputs)
transaction = function.build_transaction(build_param(provider))
transact_and_wait(provider, transaction, private_key)

print("[*] Get flag")
io = remote(CHALLENGE_IP, CHALLENGE_PORT)
io.sendlineafter(b"input your choice:", b"3")
io.sendlineafter(b"input your token:", token)
io.interactive()
