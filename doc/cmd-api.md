

### getCommitteeInfo

```shell script
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

```shell script
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