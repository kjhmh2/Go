package main

import (
    "github.com/spf13/pflag" 
    "fmt"
    "io"
    "os"
    "os/exec"
    "bufio"
)

type selgpArgs struct {
	startPage int                   // start page [-s10]
	endPage int                     // end page   [-e10]
	inputFile string                // input file
	pageLen int                     // page length
	pageType string                 // row number [-l10] or line break [-f]
	printDest string                // print destination
}

var args selgpArgs

func parseArgs(args *selgpArgs) {
	pflag.Usage = func() {
		fmt.Println("----------------------------Help----------------------------")
		fmt.Println("--usage: selpg -s[startPage] -e[endPage] [optional parameters] [filename]")
		fmt.Println("[optional parameters] -l: number of lines per page [default: 72].")
		fmt.Println("[optional parameters] -f: line break.")
		fmt.Println("[optional parameters] -l and -f are mutual exclusion.")
		fmt.Println("[optional parameters] -d: the destination of output.")
		fmt.Println("[filename] input file [default: input from the console].")
	}
	pflag.IntVarP(&args.startPage, "start page", "s", 0, "Start page of file")
	pflag.IntVarP(&args.endPage, "end page", "e", 0, "End page of file")
	pflag.IntVarP(&args.pageLen, "page length", "l", 72, "lines in one page")
	pflag.StringVarP(&args.pageType, "page type", "f", "l", "flag splits page")
	pflag.StringVarP(&args.printDest, "print destination", "d", "", "name of printer")
	pflag.Lookup("page type").NoOptDefVal = "f"
	pflag.Parse()
	otherArgs := pflag.Args()
	if (len(otherArgs) > 0) {
		args.inputFile = otherArgs[0]
	} else {
		args.inputFile = ""
	}
}

func check(args * selgpArgs) {
	if (args.startPage <= 0 || args.endPage <= 0) {
		fmt.Println("[Error] page number should be positive")
		os.Exit(1)
	} else if (args.startPage > args.endPage) {
		fmt.Println("[Error] end page should be greater than start page")
		os.Exit(2)
	} else if (args.pageLen <= 0) {
		fmt.Println("[Error] page length should be positive")
		os.Exit(3)
	} else if (args.pageType == "f" && args.pageLen != 72) {
		fmt.Println("[Error] -l and -f are mutual exclusion")
		os.Exit(4)
	}
	//fmt.Printf("start page: %d\n end page: %d\n input file: %s\n page len: %d\n page type: %s\n print destination: %s\n", args.startPage, args.endPage, args.inputFile, args.pageLen, args.pageType, args.printDest)
}

func run(arg selgpArgs) {
	var fin *os.File 
	var err error
	
	if (arg.inputFile == "") {
		fin = os.Stdin
	} else {
		fin, err = os.Open(arg.inputFile)
		if err != nil {
			fmt.Println("[Error] could not open input file")
			os.Exit(5)
		}
		defer fin.Close()
	}
	var fout io.WriteCloser
	cmd := &exec.Cmd{}

	if arg.printDest == "" {
		fout = os.Stdout
	} else {
		cmd = exec.Command("cat")
		cmd.Stdout, err = os.OpenFile(arg.printDest, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("[Error] could not open output file")
			os.Exit(6)
		}
		fout, err = cmd.StdinPipe()
		if err != nil {
			fmt.Println("[Error] could not open pipe")
			os.Exit(7)
		}
		cmd.Start()
		defer fout.Close()
	}

	pageNum := 1
	lineNum := 1
	bufFin := bufio.NewReader(fin)

	if (arg.pageType == "l") {
		for {
			line, err := bufFin.ReadString('\n')
			if (err != nil) {
				break
			}
			if ((pageNum >= arg.startPage) && (pageNum <= arg.endPage)) {
				_, err := fout.Write([]byte(line))
				if err != nil {
					fmt.Println("[Error] write error")
					os.Exit(8)
				}
		 	}
			lineNum ++
			if (lineNum > arg.pageLen) {
				pageNum ++
				lineNum = 1
			}
		}  
	} else {
		for {
			page, err := bufFin.ReadString('\f')
			if err != nil {
				break
			}
			if ((pageNum >= arg.startPage) && (pageNum <= arg.endPage)) {
				_, err := fout.Write([]byte(page))
				if err != nil {
					fmt.Println("[Error] write error")
					os.Exit(8)
				}
			}
			pageNum ++
		}
	}
	
	if (pageNum < arg.startPage) {
		fmt.Println("[Error] start page is greater than total pages")
	} else if (pageNum < arg.endPage) {
		fmt.Println("[Error] end page is greater than total pages")
	} 
	
}

func main() {
	parseArgs(&args)
	check(&args)
	run(args)
}