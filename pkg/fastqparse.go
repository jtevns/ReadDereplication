package main

import(
	"os"
	"github.com/shenwei356/bio/seqio/fastx"
	"github.com/shenwei356/xopen"
)

func parseSingle(r1 string){
// use buffered out stream for output
    outfh, err := xopen.Wopen("-") // "-" for STDOUT
    checkError(err)
    defer outfh.Close()

// disable sequence validation could reduce time when reading large sequences
// seq.ValidateSeq = false

    reader, err := fastx.NewDefaultReader(r1)
    checkError(err)
    for {
	    record, err := reader.Read()
	    if err != nil {
		    if err == io.EOF {
			    break
		    }
		    checkError(err)
		    break
	    }

	// fmt is slow for output, because it's not buffered
	// fmt.Printf("%s", record.Format(0))

	    record.FormatToWriter(outfh, 0)
    }
}


func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}