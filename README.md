# jpdocx [![Go Report Card](https://goreportcard.com/badge/github.com/sergey-chechaev/jpdocx)](https://goreportcard.com/report/github.com/sergey-chechaev/jpdocx)
Library for replacing special words in Microsoft word docx document. Library get stdin JSON format with key and value and takes two args path to docx file and path to result docx file

## Installation

If you haven't setup Go before, you need to first set a `GOPATH` (see [https://golang.org/doc/code.html#GOPATH](https://golang.org/doc/code.html#GOPATH)).

To fetch and build the code:

    $ go get github.com/sergey-chechaev/jpdocx/
    
Then compile it with the go tool:
    
    $ go install github.com/sergey-chechaev/jpdocx/
    
## Example Usage

You have test.docx documents with special words {{val_1}} and {{val_2}}:
```
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
Ut enim ad minim veniam, quis nostrud {{val_1}} ullamco laboris nisi ut aliquip ex ea commodo consequat. 
Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. 
Excepteur {{val_2}} occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum
```
Run:

    $ echo "[{\"key\": \"val_1\", \"value\": \"exercitation\"},{\"key\": \"val_2\", \"value\": \"sint\"}]" |\
    $ jpdocx "/path/to/you/test.docx" "/path/to/new/file/new_test.docx"

Output file new_test.docx:
```
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. 
Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. 
Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum
```


