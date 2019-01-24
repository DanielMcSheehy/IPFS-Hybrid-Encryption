# IPFS-Hybrid-Encryption
Demo showcasing IPFS Hybrid file encryption and decryption. Will soon be replaced with an open source library. 

## What this currently does: 
The included test generates keys, and tests strings and files that will be encrypted, uploaded to ipfs, retreived, and then finally unencrypted.
It will verifiy the files havent been tampered with. 

Soon I will replace this demo with a fully working library that will easily allow ipfs and encryption integration.

## Requirements: 
- Go

## Installation
- Install go-ipfs (https://github.com/ipfs/go-ipfs).

- Start IPFS daeman with `ipfs daemon`.

- Clone this repo to desired directory with `git clone https://github.com/DanielMcSheehy/IPFS-Hybrid-Encryption.git`

- To run tests run `go test -v`
