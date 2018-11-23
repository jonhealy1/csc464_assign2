## goByzGen

to compile: go build

to run: ./goByzGen


## Sample Output:

---------------------------

How many Generals? (One will be the Commander): 3

How many Traitors? Or is the Commander a Traitor (enter 0): 1

Does the Commander order ATTACK(enter 1), RETREAT(0)?: 1

Lieutenant 1 received: [1 0] UNDECIDED

Lieutenant 2 received: [0 1] TRAITOR

---------------------------

How many Generals? (One will be the Commander): 5

How many Traitors? Or is the Commander a Traitor (enter 0): 0

Lieutenant 1 received: [1 0 1 0] UNDECIDED

Lieutenant 2 received: [0 1 1 0] UNDECIDED

Lieutenant 3 received: [1 1 0 0] UNDECIDED

Lieutenant 4 received: [0 1 0 1] UNDECIDED

----------------------------


## Introduction:

The Byzantine Generals Problem was introduced by Leslie Lamport and is a classic problem in the field of distributed systems regarding consensus. In this problem there are a number of generals who are surrounding a city and they must decide on a plan of action amongst themselves by using a series of messages to communicate. 

There is one general who is designated as the Commander and this general begins by sending each other general a message to Attack or Retreat. The remaing generals or Lieutenants then communicate with each other in order to come to a consensus on a unified plan of attack (or retreat).

Problems arise when either the Commander or a number of the Lieutenants are Traitors. According to Lamport, in a system with 3m + 1 Generals, a maximum of m Traitors can exist before serious complications occur. Additionally, if the Commander is a Traitor and instructs some Lieutenants to Attack and others to Retreat there can be no consensus in a system with an even number of Lieutenants (this last sentence in my take on it). 

The two conditions as presented by Lamport are:

1. All loyal generals decide upon the same plan of action. He further stipulates that loyal generals should be able to come up with a reasonable plan.

2. A small number of generals cannot cause the loyal generals to come up with a bad plan. 

What is implied here is not just consensus but consensus on the right plan or the plan as presented by the Commander. With a large number of Traitors there can be consensus but the agreement reached will violate the Commander's orders and probably lead to total disaster.

Note: I am not asking for m, the level of recursion, in the console as in my understanding m is the number of traitors.

**Example 1:** Let's look at a system with 10 generals. According to Lamport this system should be able to handle at most 3 traitors gracefully. 3m + 1 = 10 ; m = 3

How many Generals? (One will be the Commander):
10
How many Traitors? Or is the Commander a Traitor (enter 0):
3
Does the Commander order ATTACK(enter 1), RETREAT(0)?:
1

Lieutenant 1 ATTACK
Lieutenant 2 ATTACK
Lieutenant 3 ATTACK
Lieutenant 4 ATTACK
Lieutenant 5 ATTACK
Lieutenant 6 ATTACK
Lieutenant 7 TRAITOR
Lieutenant 8 TRAITOR
Lieutenant 9 TRAITOR


**Example 2:** Let's add 1 Traitor to the above example. Havoc. m = 4; 3m + 1 = 13

How many Generals? (One will be the Commander):
10
How many Traitors? Or is the Commander a Traitor (enter 0):
4
Does the Commander order ATTACK(enter 1), RETREAT(0)?:
1

Lieutenant 1 UNDECIDED
Lieutenant 2 UNDECIDED
Lieutenant 3 UNDECIDED
Lieutenant 4 UNDECIDED
Lieutenant 5 UNDECIDED
Lieutenant 6 TRAITOR
Lieutenant 7 TRAITOR
Lieutenant 8 TRAITOR
Lieutenant 9 TRAITOR


**Example 3:** Adding one more Traitor... m = 5; 3m + 1 = 16

How many Generals? (One will be the Commander):
10
How many Traitors? Or is the Commander a Traitor (enter 0):
5
Does the Commander order ATTACK(enter 1), RETREAT(0)?:
1

Lieutenant 1 RETREAT
Lieutenant 2 RETREAT
Lieutenant 3 RETREAT
Lieutenant 4 RETREAT
Lieutenant 5 TRAITOR
Lieutenant 6 TRAITOR
Lieutenant 7 TRAITOR
Lieutenant 8 TRAITOR
Lieutenant 9 TRAITOR


## Implementation:

Lamport introduced a recursive algorithm to deal with the Byzantine Generals Problem without using signed messages and that algorithm is generally followed here in the code that accompanies this report. This is Lamport's algorithm:

**Algorithm OM(0)**

1. The general sends his value to every lieutenant.
2. Each lieutenant uses the value he receives from the general.

**Algorithm OM(m), m > 0**

1. The general sends his value to each lieutenant.
2. For each i, let vi be the value lieutenant i receives from the general. Lieutenant i acts as the general in Algorithm OM(m-1) to send the value vi to each of the n-2 other lieutenants.
3. For each i, and each j ≠ i, let vi be the value lieutenant i received from lieutenant j in step 2(using Algorithm (m-1)). Lieutenant i uses the value majority (v1, v2, ... vn)

In my implementation, the vote from the Commander is appended before starting the recursion. Each Lieutenant is represented as a separate goroutine. A Mutex is used to Lock and Unlock critical sections and finally a WaitGroup is used to ensure that all goroutines finish before evaluating the final results. Votes themselves are being appended to a map representing a separate slice for each Lieutenant. 

To understand this problem I drew numerous diagrams. After cross referencing what I came up with compared to what was presented in Lamport's paper I felt confident enough to start writing some code. To test my implementation I looked at the mathematics that Lamport proves in his paper and ran numerous cases using that math to ensure that my program outputs the expected results. An example of this approach is presented above in examples 1, 2, and 3. 

Note: Further testing, late at night on the due date for this assignment, with larger numbers of Traitors (> 5) did not produce the results I expected. I am not sure if this is because of a bug in my code or is a result of me not understanding the problem completely. I watched some slides online and the author said that 2m + 1 loyal nodes are needed to handle m Traitors not 3m + 1?

