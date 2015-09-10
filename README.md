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
        	rTree := radixTree{}
        	rTree.Insert("test", 1)
        	rTree.Insert("team", 2)
        
            rTree.PrintTree()
        }
```

Inspired By
=============


- [armon/go-radix](https://github.com/tuxychandru/pubsub)


License
---------------

This package is licensed under MIT license. See LICENSE for details.
