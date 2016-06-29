/*
The MIT License (MIT)

Copyright (c) 2016 Jon Bonazza

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	pipeOutput *bool = flag.Bool("pipeOutput", false, "If true, output will be piped to STDOUT (default: false)")
	pipeErr    *bool = flag.Bool("pipeErr", false, "If true, error output will be piped to STDERR (default: false)")
	wait       *bool = flag.Bool("wait", false, "If true, baka will wait for app to complete before exiting (default: false)")
)

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}
	var cmd *exec.Cmd
	if flag.NArg() > 1 {
		cmd = exec.Command(flag.Arg(0), flag.Args()[1:]...)
	} else {
		cmd = exec.Command(flag.Arg(0))
	}

	cmd.Stdin = os.Stdin
	if *pipeOutput {
		cmd.Stdout = os.Stdout
	}
	if *pipeErr {
		cmd.Stderr = os.Stderr
	}
	var err error
	if *wait {
		err = cmd.Run()
	} else {
		err = cmd.Start()
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	if !*wait {
		fmt.Printf("[%d]\n", cmd.Process.Pid)
	}
}

func usage() {
	fmt.Println("Usage: baka <app> [args...]")
	flag.PrintDefaults()
}
