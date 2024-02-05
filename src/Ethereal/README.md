
# The Ethereal Quest

## Story
Welcome to Ethereal. Here, your quest is to forge the Seed Blade, a powerful sword assembled from 10 mystical parts. With each part, your blade becomes stronger, symbolizing your growing cryptographic mastery.

Your challenge: to mint Gems with this blade. 
Mint 10 gems, and you demonstrate your skill and balance. 
Mint 20, and you push the boundaries of your abilities. But beware, each gem brings both power and peril. Will your Seed Blade be a tool of wisdom or lead to your downfall? The path is yours to choose.

## Environment Setup

### Starting an Ethereum Node
Launch an Ethereum node:
```sh
anvil
```

### Deploying the Contract
Navigate to the contract directory and deploy the contract:
```sh
# Enter the contract directory
cd mint-contracts

# Deploy the Mint contract
forge create Mint --private-key=your_key

# Return to the previous directory
cd -
```

## Starting the Challenge

### Initial Setup
Change to the client directory:
```sh
cd mint-client
```

## Challenge Steps

### 1. The Initiation of the Ethereal Quest
Begin your journey by registering:
```sh
go run . -contract your_contract_address -action register
```

### 2. The Forging of Ethereum Gems
Mint Ethereum Gems, each one bringing you closer to your goal:
```sh
go run . -contract your_contract_address -action mint
```

### 3. The Path of Redemption
Should you falter, replay and learn from your missteps:
```sh
go run . -contract your_contract_address -action replay
```

### 4. Unveil Your Treasures

Verifying the Escape0 Flag: If you've successfully minted 10 gems, the "escape0" flag should now be set, symbolizing a key achievement in your quest. This flag reflects your adeptness in navigating the challenges and evading the watchful eyes of the Watchdog.

Verifying the Escape1 Flag: For the bold adventurers who pushed their limits to mint 20 gems, the "escape1" flag would be their coveted prize. Achieving this flag marks you as one of the few who have reached the pinnacle of daring and strategy in the quest.
To check your achievements and claim your flags, invoke the following incantation:
```sh
go run . -contract your_contract_address -action query
```

## Command Options
Use these options for specific actions:
```bash
go run .
# Options:
# -action [register, mint, replay, query]
# -contract [contract_address]
# -rpc [RPC_endpoint]
# -account [account_address]
# -privateKey [private_key]
```
