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

var counter = ratecounter.New(50 * time.Millisecond, 5 * time.Second);

func doTask() {
    counter.Increment()
  
    // do some work
    
}

func printRate() {
    fmt.Println(counter.ValueBy(time.Second))
}
```