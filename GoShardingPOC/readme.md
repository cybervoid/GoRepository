# Go Sharding Partial Proof of Concept

1. This section combines two learning opportunities into one. This is an attempt at a re-write of the [Ethereum sharding partial proof of concept](https://github.com/ethereum/research/tree/master/sharding_fork_choice_poc), and improve personal knowledge in writing code in Go Lang.

## Document Reference
1. [implementation/beacon_chain_node.go](./main/implementation/beacon_chain_node.go) corresponds to: [sharding_fork_choice_poc/beacon_chain_node.py](https://github.com/ethereum/research/blob/master/sharding_fork_choice_poc/beacon_chain_node.py)
2. [implementation/distribution.go](./main/implementation/distribution.go) corresponds to: [sharding_fork_choice_poc/distributions.py](https://github.com/ethereum/research/blob/master/sharding_fork_choice_poc/distributions.py)
3. [implementation/networksim.go](./main/implementation/networksim.go) corresponds to: [sharding_fork_choice_poc/networksim.py](https://github.com/ethereum/research/blob/master/sharding_fork_choice_poc/networksim.py)
## Concept References:
1. [Quadratic Residues](https://en.wikipedia.org/wiki/Quadratic_residue) are referenced in [Beacon Chain Node](./main/implementation/beacon_chain_node.go) - the checkPow() func references quadratic residues as a requirement. Quadratic residues are defined as:  
`In number theory, an integer q is called a quadratic residue modulo n if it is congruent to a perfect square modulo n; i.e., if there exists an integer x such that:`  
x^2 &#8779; q (mod n)  
`Otherwise, q is called a quadratic nonresidue modulo n.`
  1. [Congruence](https://github.com/cybervoid/Cryptography/blob/master/CryptographyCsharp/CryptographyCsharp/Concepts/Congruence.cs) and [Equivalence](https://github.com/cybervoid/Cryptography/blob/master/CryptographyCsharp/CryptographyCsharp/Concepts/Equivalence.cs) were previously studied in greater detail when learning more about [Cryptography](https://github.com/cybervoid/Cryptography/) in Q3 and Q4 of 2017.

#### Math notes:
1. &#8779; - let this symbol represent "is congruent"


#### Other resources
1. [HTML Math](https://www.w3.org/MarkUp/html3/maths.html)
2. [HTML Encoding Mathematic symbols](https://www.w3schools.com/charsets/ref_utf_math.asp)
