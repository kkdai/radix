package radix

type radixTree struct {
}

func stringSubset(byt1, byt2 []byte) ([]byte, bool) {
	var retByte []byte
	findSubset := false
	for i = 0; i < len(byt1)-1 && i < len(byt2)-1; i++ {
		if byt1[i] != byt2[i] {
			return retByte, findSubset
		}
		findSubset = true
		retByte[i] = byt1[i]
	}
	return retByte, findSubset
}

func (t *radixTree) PrintTree() {
}

func (t *radixTree) Insert(searchKey []byte, value interface{}) {

}

func (t *radixTree) Lookup(searchKey []byte) (interface{}, bool) {
	return nil, false
}
