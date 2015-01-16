package main

import (
  "os"
  "path/filepath"
  "log"
  "bufio"
  "fmt"
  _"strings"
)

var logfiles []string

type message struct {
  timestamp, component, message string
}

type line struct {
  timestamp, component string
  message message
}

type logs struct {
  lines []line
}

func fileWalker(fp string, fi os.FileInfo, err error) error {
  if err != nil {
    log.Println(err) // can't walk here,
    return nil       // but continue walking elsewhere
  }
  if !!fi.IsDir() {
    return nil // not a file.  ignore.
  }
  matched, err := filepath.Match("*.*", fi.Name()) // look for all files
  if err != nil {
    log.Println(err) // malformed pattern
    return err       // this is fatal.
  }
  if matched {
    logfiles = append(logfiles, fp)
  }
  return nil
}

func fileScanner(fname string) {
  file, err := os.Open(fname)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
  fmt.Println(scanner.Text())
  //fmt.Println(strings.Split(scanner.Text(), " ")[2:])
  //inner_msg := strings.Split(scanner.Text(), " ")[2:]
  //fmt.Println(inner_msg[0:2])
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}

func main() {
  logpath := "loglab/logs/"
  filepath.Walk(logpath, fileWalker)
  log.Println("Files Found:", len(logfiles))
  for _, el := range logfiles{
    fileScanner(el)
  }

  line1 := line {timestamp:"TS",component:"CP", message:message{timestamp:"ITS", component:"ICP", message:"IMSG"}}
  line2 := line {timestamp:"TS1",component:"CP1", message:message{timestamp:"ITS1", component:"ICP1", message:"IMSG1"}}
  log_s := &logs {lines: []line{line1, line2}}
  for k, v := range(log_s.lines){
    fmt.Println(k, v)
  }

}


