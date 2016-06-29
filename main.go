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
