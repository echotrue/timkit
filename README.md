# Timkit
A time toolkit for Golang reference PHP's library `Carbon` 。 `Timkit` reference [carbon](https://github.com/uniplaces/carbon) .But `TimKit` is Thread-Safety

## Getting started

Install TimeKit:
```go
go get github.com/echotrue/timkit
```

## Usage
Some simple example to get started，or refer to the `example` directory
```go
package main

import (
    "fmt"
    "github.com/echotrue/timkit"
    "log"
    "time"
)

func main() {
    fmt.Println(timkit.Now())
    fmt.Println(timkit.Now().AddDays(12))
    fmt.Println(timkit.Now().SubDays(12))
    fmt.Println(timkit.Now().StartOfWeek())
    fmt.Println(timkit.Now().EndOfWeek())
    fmt.Println(timkit.Now().IsWeekday())
    
    l, err := time.LoadLocation("Local")
    if err != nil {
        log.Fatal(err)
    }
    ntk := timkit.NewTimeKit(time.Date(2021, 1, 2, 15, 4, 5, 0, l))
    fmt.Println(timkit.Now().DiffInDays(ntk, true))
    
    fmt.Println(timkit.Now().DiffInDaysFiltered(ntk, func(kit *timkit.TimeKit) bool {
        return kit.IsWeekday()
    }, true))
}

```

## Benchmark
```shell script
goos: windows
goarch: amd64
pkg: github.com/echotrue/timkit
BenchmarkTimeKit_AddCenturies
BenchmarkTimeKit_AddCenturies-12    	19989604	        59.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkTimeKit_StartOfWeek-12    	13043790	        92.5 ns/op	       0 B/op	       0 allocs/op
PASS
```

## Licence

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.