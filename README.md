Radix Tree
==============

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/radix/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/pubsub?status.svg)](https://godoc.org/github.com/kkdai/radix)  [![Build Status](https://travis-ci.org/kkdai/radix.svg?branch=master)](https://travis-ci.org/kkdai/radix)


![image](https://upload.wikimedia.org/wikipedia/commons/thumb/a/ae/Patricia_trie.svg/400px-Patricia_trie.svg.png)

What is Radix Tree
=============
a radix tree (also radix trie or compact prefix tree) is a data structure that represents a space-optimized trie in which each node that is the only child is merged with its parent.   (sited from [here](https://en.wikipedia.org/wiki/Radix_tree))


Installation and Usage
=============


Install
---------------
        go get github.com/kkdai/radix


Usage
---------------

```go

package main

import (
	"fmt"

	. "github.com/kkdai/radix"
)

func main() {
	rTree := NewRadixTree()
	rTree.Insert("test", 1)
	rTree.Insert("team", 2)
	rTree.Insert("trobot", 3)
	rTree.Insert("apple", 4)
	rTree.Insert("app", 5)
	rTree.Insert("tesla", 6)

	rTree.PrintTree()

// [1]Node has 2 edges
// [1]Edge[t]
// 	[2]Node has 2 edges
// 	[2]Edge[e]
// 		[3]Node has 2 edges
// 		[3]Edge[s]
// 			[4]Node has 2 edges
// 			[4]Edge[t]
// 				[5]Leaf key:test value:1
// 			[4]Edge[la]
// 				[5]Leaf key:tesla value:6
// 		[3]Edge[am]
// 			[4]Leaf key:team value:2
// 	[2]Edge[robot]
// 		[3]Leaf key:trobot value:3
// [1]Edge[app]
// 	[2]Node has 2 edges
// 	[2]Edge[le]
// 		[3]Leaf key:apple value:4
// 	[2]Edge[]
// 		[3]Leaf key:app value:5


	ret, find := rTree.Lookup("app")
	fmt.Println(ret, find)
//5 true
}

```

Inspired By
=============

- [Radix Tree: Wiki](https://en.wikipedia.org/wiki/Radix_tree)
- [armon/go-radix](https://github.com/armon/go-radix)


License
---------------

This package is licensed under MIT license. See LICENSE for details.
