# Go Sharding Partial Proof of Concept

1. This section combines two learning opportunities into one. This is an attempt at a re-write of the [Ethereum sharding partial proof of concept](https://github.com/ethereum/research/tree/master/sharding_fork_choice_poc), and improve personal knowledge in writing code in Go Lang.

## Document Reference
1.
## Concept References:
1. [Quadratic Residues](./implementation/beacon_chain_node.go) - the checkPow() func references quadratic residues as a requirement. Quadratic residues are defined as:
```In number theory, an integer q is called a quadratic residue modulo n if it is congruent to a perfect square modulo n; i.e., if there exists an integer x such that:  
x^2 &#8779; q (mod n)  
Otherwise, q is called a quadratic nonresidue modulo n.
```


#### Math notes:
1. &#8779; - let this symbol represent "is congruent"


#### Other resources
1. [HTML Math](https://www.w3.org/MarkUp/html3/maths.html)
2. [HTML Encoding Mathematic symbols](https://www.w3schools.com/charsets/ref_utf_math.asp)
