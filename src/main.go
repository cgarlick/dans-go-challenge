package main

import "strings"


func replaceAny(toHunt, input string) string{
	return strings.Replace(input,toHunt,"",-1)
}

func countAndReplaceAny(toHunt, input string) (count int, newString string){
	return strings.Count(input, toHunt) , replaceAny(toHunt, input)
}
