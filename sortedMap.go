package main

import "sort"

// SortedMap is a sorted string->string map
type SortedMap struct {
	theMap  map[string]string
	theKeys []string
}

func (sm *SortedMap) Len() int {
	return len(sm.theMap)
}

func (sm *SortedMap) Less(i, j int) bool {
	return sm.theMap[sm.theKeys[i]] < sm.theMap[sm.theKeys[j]]
}

func (sm *SortedMap) Swap(i, j int) {
	sm.theKeys[i], sm.theKeys[j] = sm.theKeys[j], sm.theKeys[i]
}

// SortedKeys return the sorted keys to iterate from
func SortedKeys(m map[string]string) []string {
	sm := new(SortedMap)
	sm.theMap = m
	sm.theKeys = make([]string, len(m))
	i := 0
	for key := range m {
		sm.theKeys[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.theKeys
}
