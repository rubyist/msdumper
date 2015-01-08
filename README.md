## msdumper

msdumper just periodically dumps `runtime.MemStats` to a file.

`go get github.com/rubyist/msdumper`

```go
import "github.com/rubyist/msdumper"

if err := msdumper.Start("memstats.dat", time.Second); err != nil {
  log.Fatal(err)
}

// Some time later, if you like ...

msdumper.Stop()
```

## msgraph

msgraph will use gnuplot to generate a graph from the data file.

`go install github.com/rubyist/msdumper/msgraph`

```
$ msgraph memstats.dat
```

This will output `graph.png` by default.

![graph](https://cloud.githubusercontent.com/assets/143/5670568/211ad202-974f-11e4-9c47-0920ab33c42f.png)

You can specify an output file and graph title:

```
$ msgraph -h
Usage of msgraph:
  -o="graph.png": output file
  -t="MemStats": graph titlep
```
