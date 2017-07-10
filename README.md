# WEEK 4 - TSM Dojo - Dan's problem
## Golang install
### Mac’s install using homebrew
1. brew install golang
2. configure the GOPATH & GOBIN
 * $ vi ~/.bash_profile
 * export GOPATH=/Users/thomas.dittmer/code/golang
 * export GOBIN=/Users/thomas.dittmer/code/golang/bin
 * :wq!
 * $ source ~/.bash_profile

### Windows get a MAC then follow the install setup above
### Atom is the text editor of choice for this week
1. Install atom which is on the company preferred software list
2. Because of an SSL issues we need to run two commands
 * apm config set strict-ssl false
 * apm install go-plus
3. Now your IDE is setup
4. Test the installation get the sample project from here:
 * run $ go install src/main.go
 * $GOPATH/src/main

### Go tutorial
https://tour.golang.org/welcome/1

## Dan’s problem of the week!

Count the number of letters in a given string and alphabet

### Rules
1. case insensitive
2. non alphabet letters are ignored
3. return is a array of strings formatted "lowercase letter:count" ordered by most frequent letter
4. support alphabet are Welsh and English :

### English
a|b|c|d|e|f|g|h|i|j|k|l|m|n|o|p|q|r|s|t|u|v|w|x|y|z

### Welsh
a|b|c|ch|d|dd|e|f|ff|g|ng|h|i|j|l|ll|m|n|o|p|ph|r|rh|s|t|th|u|w|y

### Input: "English", "Llanfairpwllgwyngyllgogerychwyrndrobwyll-llantysiliogogogoch"

[
"l:11"
"g:7"
"o:6"
"y:6"
"r:4"
"n:4"
"w:4"
"a:3"
"i:3"
"c:2"
"h:2"
"d:1"
"e:1"
"f:1"
"b:1"
"p:1"
"s:1"
"t":1"
]

### Input : "Welsh", "Llanfairpwllgwyngyllgogerychwyrndrobwyll-llantysiliogogogoch"

[
"g:6"
"o:6"
"y:6"
"ll:5"
"r:4"
"w:4"
"a:3"
"i:3"
"n:3"
"ch:2"
"ng:1"
"l:1"
"d:1"
"e:1"
"f:1"
"b:1"
"p:1"
"s:1"
"t:1"
]
