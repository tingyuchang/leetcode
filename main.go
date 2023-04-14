package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"
)

var TEMPLATEFILE string = "./0_Daily_Prac/template.go"

func main() {
	template, err := readTemplate()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createDailyPracticeFile(template)

	if err != nil {
		fmt.Println(err)
	}
}

func readTemplate() ([]string, error) {
	f, err := os.Open(TEMPLATEFILE)

	if err != nil {
		return nil, err

	}
	defer f.Close()

	r := bufio.NewReader(f)

	data := make([]string, 0)

	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return nil, err
		}

		data = append(data, line)
	}

	for i := 0; i < len(data); i++ {
		if data[i] == "// COPY FROM HERE\n" {
			data = data[i+1:]
			break
		}
	}

	return data, nil
}

func createDailyPracticeFile(template []string) error {
	today := time.Now().Format("20060102")
	dirName := "./0_Daily_Prac/" + today
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(dirName + "/" + today + ".go")

	if err != nil {
		return err
	}
	defer f.Close()

	// insert package name to template
	packageName := "package " + "_" + string([]byte(today)[1:]) + "\n\n"
	template = append([]string{packageName}, template...)
	sum := 0
	for i := 0; i < len(template); i++ {
		n, err := io.WriteString(f, template[i])
		if err != nil {
			return err
		}
		sum += n
	}

	fmt.Printf("Create Daily practice success with %d data\n", sum)
	defer cleanUp(err)
	return nil
}

func cleanUp(err error) {
	if err != nil {
		fmt.Println("need remove file")
	}
}

func myLog(input interface{}) {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	LstFlags := log.Ldate | log.Lshortfile
	iLog := log.New(f, "LNum ", LstFlags)
	iLog.Println(input)
	iLog.SetFlags(log.Lshortfile | log.LstdFlags)
}

func printSubsequence(s string) {

	res := make([]string, 0)

	combinationOfSubSequence(s, 0, &res, make([]byte, 0))

	fmt.Println(res)

}

func combinationOfSubSequence(s string, currentIndex int, res *[]string, current []byte) {
	if currentIndex > len(s) {
		return
	}

	if len(current) > 0 {
		*res = append(*res, string(current))
	}

	for i := currentIndex; i < len(s); i++ {
		current = append(current, s[i])
		combinationOfSubSequence(s, i+1, res, current)
		current = current[:len(current)-1]
	}
}