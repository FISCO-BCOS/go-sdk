# 权限

## 合约编译

* 合约文件编译成 abi bin

```bash
./solc-0.6.10 --bin --abi -o ./auth/contracts/abibin/ ./auth/contracts/sol/CommitteeManager.sol
```

* abi bin 生成对应的go 文件

```bash
./abigen --bin ./auth/contracts/abibin/Committee.bin --abi ./auth/contracts/abibin/Committee.abi --pkg auth --type Committee --out ./auth/Committee.go

./abigen --bin ./auth/contracts/abibin/CommitteeManager.bin --abi ./auth/contracts/abibin/CommitteeManager.abi --pkg auth --type CommitteeManager --out ./auth/CommitteeManager.go

./abigen --bin ./auth/contracts/abibin/ContractAuthPrecompiled.bin --abi ./auth/contracts/abibin/ContractAuthPrecompiled.abi --pkg auth --type ContractAuthPrecompiled --out ./auth/ContractAuthPrecompiled.go

./abigen --bin ./auth/contracts/abibin/ProposalManager.bin --abi ./auth/contracts/abibin/ProposalManager.abi --pkg auth --type ProposalManager --out ./auth/ProposalManager.go
```

### getCommitteeInfo

```bash
[group0]: /apps> getCommitteeInfo
---------------------------------------------------------------------------------------------
Committee address   : 0xa0974646d4462913a36c986ea260567cf471db1f
ProposalMgr address : 0x2568bd207f50455f1b933220d0aef11be8d096b2
---------------------------------------------------------------------------------------------
ParticipatesRate: 0% , WinRate: 0%
---------------------------------------------------------------------------------------------
Governor Address                                        | Weight
index0 : 0x357d2f663c8868b777eccc69a7bc8a9d7e4862ce     | 1

[group0]: /apps>
```

### UpdateGovernor

```bash
[group0]: /apps>  getLatestProposal
No proposal exists currently, try to propose one.

[group0]: /apps> getCommitteeInfo
---------------------------------------------------------------------------------------------
Committee address   : 0xa0974646d4462913a36c986ea260567cf471db1f
ProposalMgr address : 0x2568bd207f50455f1b933220d0aef11be8d096b2
---------------------------------------------------------------------------------------------
ParticipatesRate: 0% , WinRate: 0%
---------------------------------------------------------------------------------------------
Governor Address                                        | Weight
index0 : 0x357d2f663c8868b777eccc69a7bc8a9d7e4862ce     | 1

# 修改 UpdateGovernor 之后

[group0]: /apps> getCommitteeInfo
---------------------------------------------------------------------------------------------
Committee address   : 0xa0974646d4462913a36c986ea260567cf471db1f
ProposalMgr address : 0x2568bd207f50455f1b933220d0aef11be8d096b2
---------------------------------------------------------------------------------------------
ParticipatesRate: 0% , WinRate: 0%
---------------------------------------------------------------------------------------------
Governor Address                                        | Weight
index0 : 0x357d2f663c8868b777eccc69a7bc8a9d7e4862ce     | 1
index1 : 0xe2b91bb57b43239788740295db49301382d05021     | 1

[group0]: /apps>
```
