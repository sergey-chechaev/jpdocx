# docx_template_replace
Replase special words in docx document

## Installation

If you haven't setup Go before, you need to first set a `GOPATH` (see [https://golang.org/doc/code.html#GOPATH](https://golang.org/doc/code.html#GOPATH)).

To fetch and build the code:

    $ go get github.com/sergey-chechaev/docx_template_replace/
    
## Example Usage

You have test.docx documents with special words {{val_1}} and {{val_2}}:
```
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud {{val_1}} ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur {{val_2}} occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum
```
Run:
```
echo "[{\"key\": \"val_1\", \"value\": \"exercitation\"},{\"key\": \"val_2\", \"value\": \"sint\"}]" | jpdocx "/path/to/you/test.docx" "/path/to/new/file/new_test.docx"
```


