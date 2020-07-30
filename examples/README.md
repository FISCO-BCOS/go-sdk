# AMOP 使用案例

AMOP（Advanced Messages Onchain Protocol）即链上信使协议，旨在为联盟链提供一个安全高效的消息信道，联盟链中的各个机构，只要部署了区块链节点，无论是共识节点还是观察节点，均可使用AMOP进行通讯，AMOP有如下优势：

-   实时：AMOP消息不依赖区块链交易和共识，消息在节点间实时传输，延时在毫秒级。
-   可靠：AMOP消息传输时，自动寻找区块链网络中所有可行的链路进行通讯，只要收发双方至少有一个链路可用，消息就保证可达。
-   高效：AMOP消息结构简洁、处理逻辑高效，仅需少量cpu占用，能充分利用网络带宽。
-   安全：AMOP的所有通讯链路使用SSL加密，加密算法可配置,支持身份认证机制。
-   易用：使用AMOP时，无需在SDK做任何额外配置。

进一步了解 AMOP，请参考：[链上信使协议](https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/manual/amop_protocol.html)

**初始化**：

-   搭建单群组四节点区块链网络，可参考：[安装](https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/installation.html)。

## 单播案例

**单播** 就是节点从监听同一 Topic 的多个订阅者中随机抽取一个订阅者转发消息

-   启动 AMOP 服务端：

    ```shell
    go run examples/server/amop_server.go
    ```

    控制台输出 "subscribe success" 表示服务启动成功

-   修改 config.toml 配置文件中 NodeURL 配置项，从 127.0.0.1:20200 修改为 127.0.0.1:20201，因为 FISCO BCOS 2.5.0 不支持 AMOP 连接同一节点（FISCO BCOS 2.6.0 开始支持连接同一节点）

-   运行 AMOP 客户端：

    ```shell
    go run examples/singleclient/unicast_client.go
    ```

    服务端控制台输出 "hello, unique broadcast!"，客户端输出 "PushTopicDataRandom success" 表示运行成功

## 多播案例

**多播** 就是节点向监听同一个 Topic 的所有订阅者转发消息。只要网络正常，即使没有监听 Topic 的订阅者，消息发布者也会收到消息发送成功的网络回包

-   启动 AMOP 服务端：同单播案例

-   修改配置文件：同单播案例

-   运行 AMOP 客户端：

    ```shell
    go run examples/singleclient/unicast_client.go
    ```

    服务端控制台输出 "hello, multi broadcasts!"，客户端输出 "PushTopicDataToALL success" 表示运行成功
