package structure

import (
	"bytes"
	"fmt"
)

const CAPACITY = 64

type IntSet struct {
	words []uint64
}

func NewIntSet(size int) *IntSet {
	return &IntSet{words: make([]uint64, size/CAPACITY)}
}

func (s *IntSet) Clear() {
	for i := 0; i < len(s.words); i++ {
		s.words[i] = 0
	}
}

func (s *IntSet) Len() int {
	cnt := 0
	for _, bits := range s.words {
		if bits == 0 {
			continue
		}
		for i := 0; i < CAPACITY; i++ {
			cnt += int(bits >> i & 1)
		}
	}
	return cnt
}

func (s *IntSet) Add(v int) *IntSet {
	word, bit := v/CAPACITY, uint64(v)%CAPACITY
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] = s.words[word] | 1<<bit
	return s
}

func (s IntSet) Has(v int) bool {
	word, bit := v/CAPACITY, uint64(v)%CAPACITY
	if word >= len(s.words) {
		return false
	}
	return s.words[word] != 0 && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) unionWith(t *IntSet) *IntSet {
	for word, bit := range t.words {
		if word >= len(s.words) {
			s.words = append(s.words, bit)
		} else {
			s.words[word] = s.words[word] | bit
		}
	}
	return s
}

func (s *IntSet) intersectWith(t *IntSet) *IntSet {
	for word, bit := range t.words {
		if word >= len(s.words) {
			continue
		}
		s.words[word] = s.words[word] & bit
	}
	return s
}

func (s *IntSet) differenceWith(t *IntSet) *IntSet {
	for word, bit := range t.words {
		if word >= len(s.words) {
			s.words = append(s.words, bit)
		} else {
			s.words[word] = s.words[word] ^ bit
		}
	}
	return s
}

func (s *IntSet) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("{ ")
	for word, bit := range s.words {
		if bit == 0 {
			continue
		}
		for i := 0; i < CAPACITY; i++ {
			mask := 1 << uint64(i)
			if bit&uint64(mask) != 0 {
				_, err := fmt.Fprintf(&buffer, "%d ", word*CAPACITY+i)
				if err != nil {
					return ""
				}
			}
		}

	}
	buffer.WriteByte('}')
	return buffer.String()
}
