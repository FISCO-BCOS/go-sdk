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

**单播** 指的是节点从监听同一 Topic 的多个订阅者中随机抽取一个订阅者转发消息

-   运行 AMOP 消息发布者：

    ```shell
    # go run examples/amop/unicast_pub/publisher.go [endpoint] [topic]
    > go run examples/amop/unicast_pub/publisher.go 127.0.0.1:20200 hello
    
      2020/08/10 22:43:31 publish message: hello, FISCO BCOS, I am unique broadcast publisher! 0 
      2020/08/10 22:43:33 PushTopicDataRandom failed, err: sendMessage failed, err: error code 100, remote peer unavailable
      2020/08/10 22:43:33 publish message: hello, FISCO BCOS, I am unique broadcast publisher! 1 
      2020/08/10 22:43:35 PushTopicDataRandom failed, err: sendMessage failed, err: error code 100, remote peer unavailable
      2020/08/10 22:43:35 publish message: hello, FISCO BCOS, I am unique broadcast publisher! 2 
      2020/08/10 22:43:37 PushTopicDataRandom failed, err: sendMessage failed, err: error code 100, remote peer unavailable
      2020/08/10 22:43:37 publish message: hello, FISCO BCOS, I am unique broadcast publisher! 3 
      2020/08/10 22:43:39 publish message: hello, FISCO BCOS, I am unique broadcast publisher! 4
    ```

-   启动 AMOP 消息订阅者：

    ```shell
    # go run examples/amop/sub/subscriber.go [endpoint] [topic]
    > go run examples/amop/sub/subscriber.go 127.0.0.1:20201 hello
      
      Subscriber success
      2020/08/10 22:51:57 received: hello, FISCO BCOS, I am unique broadcast publisher! 3
      2020/08/10 22:51:59 received: hello, FISCO BCOS, I am unique broadcast publisher! 4
    ```

> 注意：控制台输出 *PushTopicDataRandom failed, err: sendMessage failed, err: error code 100, remote peer unavailable* 表示网络中没有对应的 topic 消息订阅者

## 多播案例

**多播** 指的时节点向监听同一个 Topic 的所有订阅者转发消息。只要网络正常，即使没有监听 Topic 的订阅者，消息发布者也会收到消息发送成功的网络回包

-   运行 AMOP 消息发布者：

    ```shell
    # go run examples/amop/multicast_pub/publisher.go [endpoint] [topic]
    > go run examples/amop/multicast_pub/publisher.go 127.0.0.1:20200 hello
    
      2020/08/10 22:54:06 publish message: hello, FISCO BCOS, I am multi broadcast publisher! 0 
      2020/08/10 22:54:08 publish message: hello, FISCO BCOS, I am multi broadcast publisher! 1 
      2020/08/10 22:54:10 publish message: hello, FISCO BCOS, I am multi broadcast publisher! 2 
      2020/08/10 22:54:12 publish message: hello, FISCO BCOS, I am multi broadcast publisher! 3 
      2020/08/10 22:54:14 publish message: hello, FISCO BCOS, I am multi broadcast publisher! 4 
      2020/08/10 22:54:16 publish message: hello, FISCO BCOS, I am multi broadcast publisher! 5
    ```

-   启动 AMOP 消息订阅者：

    ```shell
    # go run examples/amop/sub/subscriber.go [endpoint] [topic]
    > go run examples/amop/sub/subscriber.go 127.0.0.1:20201 hello
    
      Subscriber success
      2020/08/10 22:54:12 received: hello, FISCO BCOS, I am multi broadcast publisher! 3
      2020/08/10 22:54:14 received: hello, FISCO BCOS, I am multi broadcast publisher! 4
      2020/08/10 22:54:16 received: hello, FISCO BCOS, I am multi broadcast publisher! 5
    ```
