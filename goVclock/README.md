## goVclock

to compile: go build

to run: ./goVclock


## Introduction:

The purpose of a vector clock is to create an ordering of events in a distributed system. This enables a system to detect causality violations. Each of N processes keeps a vector with N elements and increments those elements according to these rules:

1. When a process does works it implements the clock value of its node in its own vector.

2. When sending a message, a process increments its own element and then sends a copy of its vector to the process it is sending a message to.

3. When a process receives a message, it increments its own element and then compares the values in its vector against the vector that was sent to it, choosing to use the higher of the two values for each element.

In this way, order is maintained in a distributed system without a reliable universal clock that spans multiple processes.

![alt text](https://i.postimg.cc/J7TFRDF8/go-Vclock2.png)


## Sample Output:

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


## Implementation:

This program presents a very simple simulation to show how vector clocks work. There are only 3 processes used in this implemenation which are labelled A, B, and C. To represent these processes I used a separate goroutine for each. To communicate and send information between A, B, and C, I used go channels. There are 3 channels in total, one for A and C, one for A and B, and finally one for B and C. I used a WaitGroup to ensure that A, B, and C would be finished before displaying the final results. 

I wrote more functions than I needed to because it made it easier to look at the code and visualize where messages were coming from and going to. I did collapse all of the functions at one point but it became harder for me to understand what was going on. For example, I have a receiveAB function and a receiveAC function instead of just a general receive function. ReceiveAB simply means that B is receiving a message from A. 

For testing, I drew numerous diagrams and compared them to what was outputed after running this code. I did not come up with a better method for testing. This is the output for this program after 100 iterations:

final value for A: {300 398 298}

final value for B: {299 400 298}

final value for C: {300 400 300}

A sent 200 messages and received 100 for a total of 300

B sent 200 messages and received 200 for a total of 400

C sent 100 messages and received 200 for a total of 300

In this implementation, one process (C) always has the correct number of total messages in its vector however this is not something that would always hold true and is only dependent on the order in which messages are sent. If the last two messages sent were from B to A and then B to C, no process would have the correct number of total messages. All three processes do have close to the same vector clock values and this ensures that there is relative ordering in this system. 