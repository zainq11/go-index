## insert

```shell
    insert(k, v)
      n = root
      n2, v = find(n, k)
      if item is not nil
        insert v into item.values
      else
        insert(n2, v)
        
    insert(n2, v)
      if size of n2 == order and n2 is a leaf node
        this means that there is an overflow
        split and move median up to root
```


```shell
isLeaf(n)


split(n)
  create 2 nodes from n and move median up
```

## find

```shell
find(n, k): *Node, *Item
  for each item in n.items
    if item.key == k
      return (n, item)
    if item.key > k
      if n is a leaf
        return (n, nil)
      return findNode(item.previous, k)
  return n
```


## overflow


## split