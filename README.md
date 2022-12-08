# snowcast-victorliu-sq
snowcast-victorliu-sq created by GitHub Classroom



# Deisgn

## Server

The server will start a goroutine in background called RevControlConn() at first to accept conn from clients and start a RevControlCmd() for the new socket. Then CLI from server can be handled in the main function.  In addition, the server will periodically accept conn from clients, start a RevControlCmd() for the new socket to receive and handle command msg from control. In terms of UDP socket connected with listener, the server will start a Daemon to send chunks of files for each station and broadcast chunks of the corresponding file to all of connected listeners 



| function         | description                                                  |
| ---------------- | ------------------------------------------------------------ |
| main()           | handle the server CLI                                        |
| ScanServerCLI()  | receive CLI from server and sent it to channel               |
| RevControlConn() | accept conn from clients and start a RevControlCmd() for the new socket |
| RevControlCmd()  | receive command msg from control and handle those msgs       |
| DaemsonUDP()     | for each station, broadcast chunks of the corresponding file to all of connected listeners |



## Client Control

The client control will implement a ScanClientCLI() to receive CLI from client and sent it to channel so CLI in channel can be handle CLI in the main function. After connecting to the server, the client control will receive reply msg from server and handle those msgs in RevReplyMsg().



| function        | description                                                  |
| --------------- | ------------------------------------------------------------ |
| main            | handle CLI from the control client                           |
| ScanClientCLI() | receive CLI from client and sent it to channel               |
| RevReplyMsg()   | after connecting to the server, receive reply msg from server and handle those msgs |



## Client Listener

The client Listener will only create a listener and receive chunks of file from server





# Extra Task

I have added support for adding and removing stations while the server is running through the command line interface which is worth 10 extra scores. Please uncomment from line 68 to line 71 in pkg/controlPkg/connServer.go at first because these lines will fail some test cases in grading server.



## Add a file to server

Test of adding a file (which will start a Daeome to iterate chunks in a file)

those are the input CLI

```shell
# not enough chars -> nothing happens
a 
# filename does not exist in folder -> output "file to add does not exist in folder"
a ./mp3/NonExistedFile.mp3
# filename exists in the server -> output "file to add has existed in server"
a ./mp3/Beethoven-SymphonyNo5.mp3
# valid filename -> new file will get into the server and start a Daemon to send UDP data to listener
p
a ./mp3/ManchurianCandidates-Breakin.mp3
p
```



## Remove a file from server 

Test of removing a file (which will stop a Daeome to iterate chunks in a file)

```shell
# not enough chars -> nothing will happen
r 
# stationIdx is out of range -> output "The stationIdx does not exist in server"
r 3
# stationIdx exists in the server -> remove the target stationIdx and all stationIdx of files whose stationIdx > target station will minus 1
r 0
```
