# rate-counter

## Install


``` bash
go get github.com/enterprizesoftware/rate-counter
```

``` go
import "github.com/enterprizesoftware/rate-counter"
```

## Usage

``` go

var actions = ratecounter.New(50 * time.Millisecond, 5 * time.Second);

func doTask() {
    actions.Increment()
  
    // do some work
    
}

func printStats() {
    fmt.Println(actions.Total())
    fmt.Println(actions.RatePer(time.Second))
}
```