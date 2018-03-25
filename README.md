# :construction: Blockchain from Scratch in Go with Kubernetes

Blockchain implementation from scratch in Go with Kubernetes 

Implement a basic [IOU](https://en.wikipedia.org/wiki/IOU) system like https://www.splitwise.com/ which keeps track of [who owes who](http://feedback.splitwise.com/forums/162446-general/suggestions/3457936-change-who-owes-who-to-who-owes-whom) and then simplifies the debt. 

## Key terms in simple words

### Blockchain

It's like a public transaction record with multiple distributed copies, nothing central that is coordinating things.
Transactions are stored in groups called "blocks" and there are a bunch of these blocks in series (hence the word "chain") 
Collectively it's called a "ledger".

Blockchain has some key properties that make it ideal for decentralized money aka bitcoin aka cryptocurrency. 

1. Immutable: Things cannot be changed once stored in a block. (It is close to impossible to change something in a blockchain)

2. Decentralized: There is no central entity controlling and managing things, so no single point of failure. 
                  Every node in the network is empowered and together all the nodes form consensus, which makes
                  it extremely difficult to hack the majority of the nodes to temper the data. 
                  Think of this as direct democracy vs authoritarian rule.
                  
3. Distributed trust: We may not trust an individual node, but we trust a large group of nodes. An individual node can be
               manipulated/hacked but it's close to impossible to do so with so many nodes. Power in numbers.
               
4. Security: Blockchain uses [public key cryptography](https://medium.com/@vrypan/explaining-public-key-cryptography-to-non-geeks-f0994b3c2d5) and hashing to ensure the data security. 


### Hash

Hash is kinda like the summery of the data. Data of any size will have a unique fingerprint that changes even if 
the data changes by a single bit. This is used to ensure the data integrity. 

### Digital Signature

Digital signature is a [public key cryptography](https://en.wikipedia.org/wiki/Public-key_cryptography) concept. 
In super simple language, imagine Bob has a signature and they use a secrete pen (private key) to sign letters.
Other people can verify that it was signed by Bob using special magnifying glass (public key) which can verify Bob's signature.
Now this secrete pen (private key - only one copy- stays hidden with Bob) and the special magnifying glasses (public key - Bob hands them out 
to everyone - multiple copies) come in a pair. If the pen changes then the magnifying glasses also change.
Another thing that is special about a digital signature is it's different for different letters/content it signs.
So, the signature will be different for different data it is signing, however still using the same secrete pen (private key),
so it can be verified using the special magnifying glass (public key). Due to some complicated cyptography math and computation
limits it's close to impossible to forge the signature that can be verified using Bob's special magnifying glasses.

### Proof of Work (vs Proof of Stake)

To prevent fraudulent transactions, blockchain requires miners (computers that take part in helping calculate hashes)
to solve a crypto puzzle - a bunch of miners start solving the puzzle and whoever solves it first wins the transaction fees
for that block - (essentially calculate a hash that starts with X number of 0's) which takes time, so the bad person
can't insert a lot of bad blocks in the blockchain since each one takes some time to solve the puzzle (and computing power).
Solving this puzzle (having a hash with X number of 0's) is how a miner shows their "proof of work". Now this concept/algorithm 
has some issues since it favors miners with more computing power (and specialized equipments to calculate the hash faster), essentially
wasting a lot of power. A lot of miners get together in a "mining pool" and whenever one of them solves and wins the puzzle, they
share the prize with other miners in the pool, now on the surface this doesn't sound so bad since you're just sharing the prize,
but if a pool gets big enough, it can have 51% share in a blockchain and can do some serious damage through 51% attack - this also
defeats the purpose of being "decentralized" if one large pool can control all the transactions. 
    To solve these shortcomings of the Proof of Work algorithm, another algoritham was proposed called Proof of Stake. 
In Proof of Stake, a miner is selected at random (almost) and then they need to pay some money as a "security deposit" - which is higher
than the transaction fees you get paid if you solve the puzzle. You get your security deposit back after a certain time - if your 
transaction is not fraudulent. This solves the problem of not wasting so much electricity compared to Proof of Work since many miners
spend time and power solving the puzzle but only one answer is used rest of the miners' work is thrown away - waste of power. Another 
problem this solves is it's very very hard to get 51% stake in the network since it will require enormous "security deposit". 
There is aan [awesome video](https://www.youtube.com/watch?v=M3EFi_POhps) explaining both of these concepts.


