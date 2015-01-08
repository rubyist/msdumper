package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
	"text/template"
)

var tpl = `set term png large size 1024,768
set grid
set key bottom right
set title "{{.Title}}"
set ytics nomirror
set ylabel 'Memory'
set y2tics nomirror
set y2label 'Heap Objects'
plot  "{{.File}}" using 5 title "HeapSys"  with lines lw 2, "{{.File}}" using 4 title "HeapAlloc"  with lines lw 2, "{{.File}}" using 6 title "HeapIdle"  with lines lw 2, "{{.File}}" using 8 title "HeapObjects" with lines lw 2 lc rgb 'grey' axes x1y2
`

var output = flag.String("o", "graph.png", "output file")
var title = flag.String("t", "MemStats", "graph title")

type Data struct {
	File  string
	Title string
}

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatal("Usage: msgraph <file>")
	}

	var buf bytes.Buffer

	dataFile := flag.Arg(0)

	t := template.Must(template.New("output").Parse(tpl))

	t.Execute(&buf, Data{dataFile, *title})

	file, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("gnuplot")
	in, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	out, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(in, &buf)
	if err != nil {
		log.Fatal(err)
	}

	in.Close()

	_, err = io.Copy(file, out)
	if err != nil {
		log.Fatal(err)
	}

	if err = cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	log.Print("Done.")
}
