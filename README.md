## Benchmark

**测试环境**
* CPU:    2.5 GHz Intel Core i7
* Memory: 16 GB 2133 MHz LPDDR3
* OS:     MacBook Pro (13-inch, 2017, Two Thunderbolt 3 ports)
* Go:     1.11

测试代码client是通过protobuf编解码和server通讯的。
请求发送给server, server解码、更新两个字段、编码再发送给client，所以整个测试会包含客户端的编解码和服务器端的编解码。
消息的内容大约为581 byte, 在传输的过程中会增加少许的头信息，所以完整的消息大小在600字节左右。

测试用的proto文件如下：

```proto
syntax = "proto2";

package main;

option optimize_for = SPEED;


message BenchmarkMessage {
  required string field1 = 1;
  optional string field9 = 9;
  optional string field18 = 18;
  optional bool field80 = 80 [default=false];
  optional bool field81 = 81 [default=true];
  required int32 field2 = 2;
  required int32 field3 = 3;
  optional int32 field280 = 280;
  optional int32 field6 = 6 [default=0];
  optional int64 field22 = 22;
  optional string field4 = 4;
  repeated fixed64 field5 = 5;
  optional bool field59 = 59 [default=false];
  optional string field7 = 7;
  optional int32 field16 = 16;
  optional int32 field130 = 130 [default=0];
  optional bool field12 = 12 [default=true];
  optional bool field17 = 17 [default=true];
  optional bool field13 = 13 [default=true];
  optional bool field14 = 14 [default=true];
  optional int32 field104 = 104 [default=0];
  optional int32 field100 = 100 [default=0];
  optional int32 field101 = 101 [default=0];
  optional string field102 = 102;
  optional string field103 = 103;
  optional int32 field29 = 29 [default=0];
  optional bool field30 = 30 [default=false];
  optional int32 field60 = 60 [default=-1];
  optional int32 field271 = 271 [default=-1];
  optional int32 field272 = 272 [default=-1];
  optional int32 field150 = 150;
  optional int32 field23 = 23 [default=0];
  optional bool field24 = 24 [default=false];
  optional int32 field25 = 25 [default=0];
  optional bool field78 = 78;
  optional int32 field67 = 67 [default=0];
  optional int32 field68 = 68;
  optional int32 field128 = 128 [default=0];
  optional string field129 = 129 [default="xxxxxxxxxxxxxxxxxxxxx"];
  optional int32 field131 = 131 [default=0];
}
```

测试的并发client是 100, 1000,2000 and 5000。总请求数一百万。

**测试结果**

### teleport

#### 一个服务器和一个客户端，在同一台机器上

并发client|平均值(ms)|中位数(ms)|最大值(ms)|最小值(ms)|p99|吞吐率(TPS)
-------------|-------------|-------------|-------------|-------------|-------------|-------------
100|1|1|30|0|8|50002
500|11|10|67|0|29|44942
1000|23|22|145|0|57|42867


可以看出平均值和中位数值相差不大，说明没有太多的离谱的延迟。

随着并发数的增大，服务器延迟也越长，这是正常的。


### gRPC
[gRPC](https://github.com/grpc/grpc-go) 是Google开发的一个RPC框架，支持多种编程语言。

我对gRPC和rpcx进行了相同的测试，得到了相应的测试结果。结果显示rpcx的性能要远远好于gRPC。
gRPC的优势之一就是随着并发数的增大，吞吐率比较稳定，而rpcx随着并发数的增加性能有所下降，但总体吞吐率还是要高于gRPC的。


#### 一个服务器和一个客户端，在同一台机器上

并发client|平均值(ms)|中位数(ms)|最大值(ms)|最小值(ms)|p99|吞吐率(TPS)
-------------|-------------|-------------|-------------|-------------|-------------|-------------
100|3|2|55|0|14|33089
500|15|14|154|0|77|32069
1000|34|31|233|0|104|28769

### rpcx

[rpcx](https://github.com/smallnest/rpcx) 是一款号称最快的RPC框架。

#### 一个服务器和一个客户端，在同一台机器上

并发client|平均值(ms)|中位数(ms)|最大值(ms)|最小值(ms)|p99|吞吐率(TPS)
-------------|-------------|-------------|-------------|-------------|-------------|-------------
100|2|1|85|0|13|45442
500|11|10|108|0|63|42025
1000|26|23|263|0|144|36837

<table>
<tr><th>Environment</th><th>Throughputs</th><th>Mean Latency</th><th>P99 Latency</th></tr>
<tr>
<td width="30%"><img src="https://github.com/henrylee2cn/rpc-benchmark/raw/master/result/env.png"></td>
<td width="30%"><img src="https://github.com/henrylee2cn/rpc-benchmark/raw/master/result/throughput.png"></td>
<td width="30%"><img src="https://github.com/henrylee2cn/rpc-benchmark/raw/master/result/mean_latency.png"></td>
<td width="30%"><img src="https://github.com/henrylee2cn/rpc-benchmark/raw/master/result/p99_latency.png"></td>
</tr>
</table>
