

## 合约编译

* 合约文件编译成 abi bin 
```shell script
./solc-0.6.10 --bin --abi -o ./auth/contracts/abibin/ ./auth/contracts/sol/CommitteeManager.sol
```

* abi bin 生成对应的go 文件

```shell script
./abigen --bin ./auth/contracts/abibin/Committee.bin --abi ./auth/contracts/abibin/Committee.abi --pkg auth --type Committee --out ./auth/Committee.go

./abigen --bin ./auth/contracts/abibin/CommitteeManager.bin --abi ./auth/contracts/abibin/CommitteeManager.abi --pkg auth --type CommitteeManager --out ./auth/CommitteeManager.go

./abigen --bin ./auth/contracts/abibin/ContractAuthPrecompiled.bin --abi ./auth/contracts/abibin/ContractAuthPrecompiled.abi --pkg auth --type ContractAuthPrecompiled --out ./auth/ContractAuthPrecompiled.go

./abigen --bin ./auth/contracts/abibin/ProposalManager.bin --abi ./auth/contracts/abibin/ProposalManager.abi --pkg auth --type ProposalManager --out ./auth/ProposalManager.go
```