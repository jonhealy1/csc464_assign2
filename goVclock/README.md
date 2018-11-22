# goVclock

to compile: go build

to run: ./goVclock2

The purpose of a vector clock is to create a partial ordering of events in a distributed system. This enables a system to detect causality violations. Each of N processes keeps a vector with N elements and increments those elements according to these rules:

1. When sending a message, a process increments its own element and then sends a copy of its vector to the process it is sending a message to.

2. When a process receives a message, it increments its own element and then compares the values in its vector against the vector that was sent to it, choosing to use the higher of the two values for each element.

In this way, order is maintained in a distributed system without a reliable universal clock that spans multiple processes.

![alt text](https://i.postimg.cc/J7TFRDF8/go-Vclock2.png)

Sample Output:

How many iterations?
1

msg from C: 0 0 1  to B: 0 0 0

B updated to: 0 1 1

msg from B: 0 2 1 to A: 0 0 0

A updated to: 1 2 1

msg from A: 2 2 1 to B: 0 2 1

B updated to: 2 3 1

msg from A: 3 2 1 to C: 0 0 1

C updated to: 3 2 2

msg from B: 2 4 1 to C: 3 2 2

C updated to: 3 4 3

final value for A: {3 2 1}

final value for B: {2 4 1}

final value for C: {3 4 3}

A sent 2 messages and received 1 for a total of 3

B sent 2 messages and received 2 for a total of 4

C sent 1 messages and received 2 for a total of 3
