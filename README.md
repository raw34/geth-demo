# eth-demo

# 编译合约
```
solc --abi contracts/Store.sol -o contracts
solc --bin contracts/Store.sol -o contracts
abigen --bin=contracts/Store.bin --abi=contracts/Store.abi --pkg=contracts --out=contracts/Store.go
```
