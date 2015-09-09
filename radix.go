package radix

type radixTree struct {
	root node
}

func stringSubsetPrefix(byt1, byt2 []byte) ([]byte, bool) {
	var biggerLen int
	if len(byt1) > len(byt2) {
		biggerLen = len(byt1)
	} else {
		biggerLen = len(byt2)
	}

	workByte := make([]byte, biggerLen)
	findSubset := false
	for i := 0; i < len(byt1) && i < len(byt2); i++ {
		if byt1[i] != byt2[i] {
			retByte := make([]byte, i+1)
			retByte = workByte[:i]
			return retByte, findSubset
		}
		findSubset = true
		workByte[i] = byt1[i]
	}
	return workByte, findSubset
}

func (t *radixTree) PrintTree() {
}

func (t *radixTree) Insert(searchKey []byte, value interface{}) {

}

func (t *radixTree) Lookup(searchKey []byte) (interface{}, bool) {
	return nil, false
}
