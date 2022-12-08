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



# Initialize

server

```shell
./snowcast_server 8888 ./mp3/Beethoven-SymphonyNo5.mp3 ./mp3/DukeEllington-Caravan.mp3 ./mp3/FX-Impact193.mp3
```



client1

```shell
./snowcast_control localhost 8888 5000
```

listener1

```shell
./snowcast_listener 5000 | pv > /dev/null
```



# Add another pair of client&listener

client2

```shell
./snowcast_control localhost 8888 6000
```

listener2

```shell
./snowcast_listener 6000 | pv > /dev/null
```



# Operation

```shell
# client2
2
# client1
0
# server
p
# client1
1
# server
p
p out
```



# Extra credit

All of them are on the side of server

Test of adding a file (which will start a Daeome to iterate chunks in a file)

```shell
# not enough chars
a 
# filename does not exist in folder
a ./mp3/NonExistedFile.mp3
# filename exists in the server
a ./mp3/Beethoven-SymphonyNo5.mp3
# valid filename
p
a ./mp3/ManchurianCandidates-Breakin.mp3
p
```



Test of removing a file (which will stop a Daeome to iterate chunks in a file)

```shell
# not enough chars
r 
# stationIdx is out of range
r 3
# stationIdx exists in the server
r 0
```

