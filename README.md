# domain-go

## Usage

```sh
go get -u github.com/byronhallett/domain-go
```

```Go
package main

import (
  "log"

  "github.com/davecgh/go-spew/spew"

  domain "github.com/byronhallett/domain-go"
  types "github.com/byronhallett/domain-go/types"
)

func main() {
  session, err := domain.NewSession(
    "client_**********************", "secret_************************",
  )
  if err != nil {
    log.Fatalln(err)
  }

  results, err := session.ResidentialSearch(types.SearchParameters{
    Locations: []types.SearchLocation{
      types.SearchLocation{
        Suburb:                    "Summer Hill",
        IncludeSurroundingSuburbs: true,
      },
    },
  })
  if err != nil {
    log.Fatalln(err)
  }
  spew.Dump(results.Listings)
}
```
