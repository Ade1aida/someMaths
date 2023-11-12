package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/Ade1aida/someMaths/pkg/calculations"
	"github.com/sirupsen/logrus"
)

func main() {

	var n int64
	var logr = logrus.New()
	logr.Formatter.(*logrus.TextFormatter).DisableTimestamp = true
	logrus.SetLevel(logrus.InfoLevel)
	logr.Out = os.Stdout

	arg1 := os.Args[2]

	n, err := strconv.ParseInt(arg1, 10, 64)
	if err != nil {
		panic(err)
	}

	logs := flag.Bool("log", false, "флаг, указывающий на использование логирования")
	flag.Parse()

	result := calculations.Factorial(n, *logs)
	logr.WithFields(logrus.Fields{
		"Результат ": result,
	}).Info()

}
