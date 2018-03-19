# :construction: Blockchain from Scratch in Go with Kubernetes

Blockchain implementation from scratch in Go with Kubernetes 

Implement a [IOU](https://en.wikipedia.org/wiki/IOU) system like https://www.splitwise.com/ which keeps track of who ows who and then simplifies the debt. 

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

### Proof of Work

