package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)
/*
func readerPost(str string) int {
	numbers := 0
	inputFile := str

	file, err := os.Open(inputFile) // 打开
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer file.Close() // 关闭

	line := bufio.NewReader(file)
	for {
		content, _, err := line.ReadLine()
		if err == io.EOF {
			return numbers
		}
		numbers++
		fmt.Println(string(content))
	}
}
*/
func check(judge bool, str string) string {
	if judge {
		return str
	}
	return " "
}

type Node struct {
	apiName, apiUrl string
}

func addWay(str, tem string) Node {
	judge1 := strings.Contains(str, "openapi/")
	judge2 := strings.Contains(str, "databases")
	judge3 := strings.Contains(str, "search")
	judge4 := strings.Contains(str, "features")
	judge5 := strings.Contains(str, "keys")
	judge6 := strings.Contains(str, "system")
	judge7 := strings.Contains(str, "compare")
	judge8 := strings.Contains(str, "detect")
	judge9 := strings.Contains(str, "qualit")

	if judge1{
		tem += "app_id"
	} else {
		return Node{"",""}
	}

	if judge2{
		tem += " databases"
		if judge4{
			tem += " db_id" + " features"
			if judge5{
				tem += " keys"
			}
		} else if !judge3 && !judge6 && !judge7 && !judge8 && !judge9{
			var tmp = ""
			for i := len(str)-2; tmp != "databases"; i -- {
				if str[i] == '/'{
					if tmp != ""{
						tem += " db_id"
						break
					}
					tmp = ""
				}
				tmp = string(str[i]) + tmp
			}
		}
		if judge3{ tem += " search" }

	}
	if judge6 {tem += "system"}
	if judge7 {tem += "compare"}
	if judge8 {tem += "detect"}
	if judge9 {tem += "qualit"}
	var tmp, tmp1 string
	tmp = ""
	tmp1 = ""
	//var s, f, z bool
	//s = false; f = false; z = false
	for i := 0; i < len(str); i ++ {
		if str[i] == '/' || str[i] == '?' {
			//fmt.Println(tmp)
			if tmp == "openapi" || tmp == "databases" || tmp == "keys" {
				//fmt.Println("+++++++++++++++++++++++++++++++++")
				if tmp == "openapi" {
					cnt := 0
					j := i + 1
					tmp1 += " "
					for ; cnt < 3 && j < len(str); j++ {
						if cnt == 2 && str[j] != '/' {
							tmp1 += string(str[j])
						}
						if str[j] == '/' {
							cnt++
						}

					}
					i = j - 1
					//s = true
				} else if tmp == "databases" {
					var tpl = ""
					j := i + 1
					for ; j < len(str); j++ {
						if str[j] == '"' || str[j] == '/' {
							break
						}
						tpl += string(str[j])
					}
					if len(tpl) > 12 {
						//fmt.Println(tpl)
						tmp1 += " " + tpl
						//f = true
					}
				} else if tmp == "keys" {
					j := i + 1
					tmp1 += " "
					for ; j < len(str) && str[j] != '"'; j++ {
						tmp1 += string(str[j])
					}
					//z = true
				}
			}
			tmp = ""
			continue
		}
		tmp += string(str[i])
	}
	var count =  strings.Count(tmp1, " ")
	if count < 3 {
		if count == 0 {tmp1 += "      "}else if count == 1{
			tmp1 += "    "
		}else {tmp1 += "  "}

	}
	return Node{tem,tmp1}
}

var mmp map[string]string

func Init() {
	mmp = make(map[string]string)
	mmp["POST app_id databases search"] = "Cross Databases Face Search"
	mmp["POST app_id compare"] = "Compare Faces in 2 Images"
	mmp["POST app_id detect"] = "Detect Faces in Images"
	mmp["POST app_id quality"] = "Face Quality Check"
	mmp["GET app_id databases"] = "List Feature Databases"
	mmp["POST app_id databases"] = "Create Feature Database"
	mmp["GET app_id databases db_id"] = "Get Feature Database Info"
	mmp["DELETE app_id databases db_id"] = "Delete Feature Database"
	mmp["GET app_id databases db_id features"] = "Batch Get Features by Ids"
	mmp["POST app_id databases db_id features"] = "Batch Add Faces to Database"
	mmp["DELETE app_id databases db_id features"] = "Batch Delete Features by Ids"
	mmp["GET app_id databases db_id features keys"] = "Batch Get Features by Keys"
	mmp["DELETE app_id databases db_id features keys"] = "Batch Delete Features by Keys"
	mmp["POST app_id databases search"] = "Cross Databases Face Search"
	mmp["GET system"] = "Get System Info"
}
var cliName = flag.String("inputPath", "", "Input Your inputFilePath")
var cliName1 = flag.String("outputPath", "", "Input your outputFilePath")
func main() {
	Init()
	flag.Parse()
	//fmt.Println(*cliName)
	inputFile := *cliName
	var inp = *cliName
	var oup = *cliName1
	if inp == "" || oup == ""{
		fmt.Printf("请输入:-inputPath=? -outputPath=?")
		return
	}
	outputFile, outputError := os.OpenFile(*cliName1,os.O_WRONLY | os.O_CREATE, 0666)
	if nil != outputError {
		fmt.Println("creat error")
		return
	}
	defer outputFile.Close()

	// 2.逐行读取
	file, err := os.Open(inputFile)  // 打开
	if err != nil { fmt.Println(err); return }
	defer file.Close()  // 关闭

	line := bufio.NewReader(file)
	outputWrite := bufio.NewWriter(outputFile)

	var str, res, tmp string
	var t Node
	var sign bool
	for {
		sign = false
		content, err := line.ReadString('\n')
		if err == io.EOF{ break }
		str = ""; res = ""; tmp = ""; cnt := 0; count := 0
		for i:= 0; i < len(string(content)); i ++ {
			if cnt == 15 {
				str += string(content[i])
				if string(content[i]) == "\""{ count ++ }
				if count == 2 {
					if !sign{
						res += str + " " + tmp //结尾加的东西
						outputWrite.WriteString(res + "\n")
					}
					break
				}
				continue
			}
			if string(content[i]) == " " {
				cnt ++
				if str == "\"GET" || str == "\"POST" || str == "\"DELETE" {
					tmp = str[1:len(str)] + " "    //只会出现一次
				}
				if cnt == 3 || cnt == 4 || cnt == 5 {
					str = ""; continue
				} else if str == "HTTP/1.1\"" || str == "http/2.0\""{
					res += "\"" + " "
				}else {
					if len(str) > 15 && "http" == str[0:4] { // 只会出现一次
						t = addWay(str, tmp)
						if t.apiName == "" && t.apiUrl == ""{
							sign = true
							break
						}
						if mmp[t.apiName] != "" {
							tmp = "\""  + mmp[t.apiName] + "\""
						}else {tmp = " "}
						//if t.apiName == "GET system" {fmt.Println(tmp)}
						tmp += t.apiUrl

					}
					res += str
					if cnt  != 14 {res += " "}
				}
				str = "";continue
			}
			str += string(content[i])
		}
	}
	outputWrite.Flush()
	//var ans = readerPost(*cliName1)
	//fmt.Printf("POST请求次数%d\n", ans)
}

