# DomiGo
DomiGo is a Go interface to access HCL Domino. Please be aware that at this point there are several methods that are not working properly yet (either due to wrong type implementations or stuff I just couldn't test properly due to my setup). So please feel free to participate.

```bash
go get -u "github.com/monstermichl/domigo"
```

## Usage
To get access to HCL Domino, just call *domigo.Initialize*. The function returns a NotesSession struct which can then be used as described in the [NotesSession documentation](https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESSESSION_CLASS.html).

**IMPORTANT:** The structs/struct-slices which are returned by some methods must be released by calling the *Release*-function (e.g. *defer session.Release()*).

**IMPORTANT:** If you're using a 32 bit installation of HCL Domino, make sure to set the environment variable *GOARCH=386*.

```go
package main

import (
    "fmt"
    "github.com/monstermichl/domigo"
)

func main() {
    // Initialize session.
    session, err := domigo.Initialize()
    defer session.Release()

    if err != nil {
        fmt.Println(err)
        return
    }

    // Get local database via session.
    db, err := session.GetDatabase("", "TestDatabase.nsf")
    defer db.Release()

    if err != nil {
        fmt.Println(err)
        return
    }

    // Get database title.
    title, err := db.Title()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Title", title)
}
```

## Test
See [README.md](https://github.com/monstermichl/domigo/tree/main/test/README.md)
