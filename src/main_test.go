package main

import "testing"
import "fmt"
import assert "github.com/stretchr/testify/assert"


func TestReplace(t *testing.T) {
	assert.Equal(t, "A", replaceAny("B","AB"), "WRONG")
}

func TestCountAndReplace(t *testing.T) {
	count,newString := countAndReplaceAny("B","AB")
	assert.Equal(t, 1, count, "Miscount")
	assert.Equal(t, "A", newString, "WRONG")
}

func TestLoopThroughDict(t *testing.T) {
	dict := [...]string{"A","B","C"}
	fmt.Println(dict)
	for key, letter := range dict {
		fmt.Println(key, letter)
	}
}
