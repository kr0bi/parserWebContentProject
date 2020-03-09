# parserWebContentProject
A simple parser to a HTML table from http://users.dimi.uniud.it/~alberto.policriti/home/?q=node/42 

# How to compile it
cd main

go build html.go

go build gopl.io/ch1/fetch

./fetch http://users.dimi.uniud.it/~alberto.policriti/home/?q=node/42 | ./html
