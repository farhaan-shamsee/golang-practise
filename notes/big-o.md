14. The most common complexities you should know
For now, memorize this order:
O(1)
   ↓
O(log n)
   ↓
O(n)
   ↓
O(n log n)
   ↓
O(n²)
   ↓
O(2ⁿ)
   ↓
O(n!)

Generally, as you move downward, algorithms scale worse as n becomes large.
But don't worry about O(2ⁿ) and O(n!) yet.
For your upcoming string/array interview problems, these are the ones I want you to become extremely comfortable with first:
Complexity	Mental model
O(1)	Do a fixed amount of work
O(log n)	Keep eliminating a large portion of the search space
O(n)	Go through the input once
O(n log n)	Often sorting
O(n²)	Compare many/all pairs


Single traversal           → O(n)

Two sequential traversals  → O(n)

Nested n × n loops         → O(n²)

Nested n × constant loops  → O(n)

Extra array/map of n       → O(n) space

Fixed variables            → O(1) space


1. Traversal
       ↓
   Visit every element

2. Accumulator
       ↓
   sum / count / max / min

3. Early exit
       ↓
   Stop when answer is found

4. Two pointers
       ↓
   Process from both ends