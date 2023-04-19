# go-index
My naive attempt at implementing indexes. For starters, I am trying out a [B-Tree](https://www.cs.cornell.edu/courses/cs3110/2012sp/recitations/rec25-B-trees/rec25.html) 
implementation. This is still a very much WIP.

## Building a Btree

- Creating an m-way tree
  - A node
    - `m` children
    - `m - 1` values in ascending order
    - `isLeaf` method to determine if the node is a leaf
  
  
## Constraints for a btree and tasks

A B-Tree is a self-balancing search tree that is commonly used in databases and file systems. It differs from a typical m-way tree in that it has the following additional conditions:

1. Degree constraint: A B-Tree of order m has a degree constraint that requires each node to have at most m children.
2. Root constraint: The root of the B-Tree can have at least 2 children unless it's a leaf node.
3. Leaf node constraint: All leaf nodes of the B-Tree should appear at the same level, which ensures that the height of the tree is minimized, and the search and insertion operations can be performed efficiently.
4. Non-leaf node constraint: All non-leaf nodes except the root must have at least ⌈m/2⌉ (ceil) children. This constraint ensures that the nodes are well-distributed and the tree is balanced.
5. Key constraint: The keys in each node are sorted in non-decreasing order, and each key is unique within a node. (`Implement using a sorted set`)
6. Overflow constraint: When a node overflows, it must be split into two nodes, and the median key should be moved up to its parent node.
   - How do I know if the node overflows? More than size == m - 1 items in the node before the insertion.
   - Keep track of the median node so that it can be moved to the parent node.
7. Underflow constraint: When a node underflows, it should be merged with its sibling, and the median key should be moved down to the merged node.
    - This is a delete concern. Will handle it later.
These conditions ensure that a B-Tree remains balanced, and its height is minimized, which leads to efficient search and insertion operations.


