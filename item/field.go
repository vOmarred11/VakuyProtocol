package item

type ItemHash map[string]int

func (t ItemHash) Hash() ItemHash {
	return map[string]int{}
}
