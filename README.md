# Syncop

Go library to handle error inside multiple go routine. Sometimes you need all go routine to run successfully without single failure, to make sure of that you need to handle error that happen inside the routines.

## Installation

```bash
go get github.com/faruqisan/syncop
```

## Usage

```go

package main

import (
    "github.com/faruqisan/syncop"
)

func main(){

    cop := syncop.New()

    cop.WgAdd(1)
    go func(){
        defer cop.WgDone()
        fooResult, err := fooServices()
        if err != nil {
            // if error happen we want the cop to handle the error
            cop.HandleError(err)
        }

        fooResult.DoSomething()
    }()

    cop.WgAdd(1)
    go func(){
        defer cop.WgDone()
        barResult, err := barServices()
        if err != nil {
            cop.HandleError(err)
        }

        barResult.DoOtherThing()
    }()

    err := cop.ListenForError()
    if err != nil {
        // do something with error
    }

}

```