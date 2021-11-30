package main

import (
	"fmt"
	"sort"
)

type Player struct {
	Name string
	Age  int
	Goal int
	Pass float64
}

type Players []Player

func (s Players) Len() int           { return len(s) }
func (s Players) Less(i, j int) bool { return s[i].Goal > s[j].Goal }
func (s Players) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	s := []Player{
		{"나통키", 13, 45, 78.4}, {"오맹태", 16, 24, 67.4},
		{"오동도", 18, 54, 50.8}, {"황금산", 16, 36, 89.7}}

	sort.Sort(Players(s))
	fmt.Println(s)
}
