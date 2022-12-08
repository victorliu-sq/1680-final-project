# Final Project Proposal

Jiaxin Liu

Zongjie Liu



In our final project, we are going to build a better snowcast based on gRPC, a modern open source high performance Remote Procedure Call framework that can run in any environment. After replacing the TCP package with gRPC framework, we no longer need to care about sending and receiving messages in terms of bytes.



At first, we will build a client and server that connects and exchanegs Snowcast's Hello/Welcome messages through the gRPC Request and Reply. After that, we can add the functionalities of SetStation and Welcome to that pair of Request and Reply, which can greatly simplify the configuration of protocols. Since gRPC uses HTTP/2 to multiplex multiple RPCs on the same TCP connection, we can still let the server to send bytes of mp3 files to client listener in a UDP socket. 



There is one thing in our mind. Actually, we also consider to use gRPC in the process of data streaming. Is it feasible to change the structure and send data of mp3 file to listener on TCP connection? 