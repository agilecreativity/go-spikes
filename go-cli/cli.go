package main

import (
    "os"
    "github.com/codegangsta/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "excel2csv"
    app.Usage = "convert excel file to csv"
    app.Action = func(c *cli.Context) {
        println("convert excel to csv")
    }
    app.Run(os.Args)
}
