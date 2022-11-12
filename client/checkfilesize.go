package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

func main(){
	const recordfilename string = "./recordprocess.txt"
	recordfile, err1:= os.OpenFile(recordfilename, os.O_WRONLY|os.O_CREATE,0600)
	if err1 != nil {
		panic(any(err1))
	}
	defer recordfile.Close()

	recordwriter := bufio.NewWriter(recordfile)
	i := 0
	recordwriter.WriteString(string(i)+"\n")
	recordwriter.Flush()

	starttime := time.Now()
	for {
		fi, err := os.Stat("video.ts")
		if err == nil {
			curtime := time.Since(starttime)
			filesize := fi.Size()

			var slice []string
			slice = append(slice, strconv.FormatInt(int64(curtime),10))
			slice = append(slice, strconv.FormatInt(filesize,10))
			outputstr := strings.Join(slice, ",")
			outputstr += "\n"		// 换行符

			recordwriter.WriteString(outputstr)
			recordwriter.Flush()
		}else{
			starttime = time.Now()
			i ++
			recordwriter.WriteString(string(i)+"\n")
			recordwriter.Flush()
			continue
		}
	}
}
