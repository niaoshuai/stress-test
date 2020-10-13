package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"time"

	logger "github.com/niaoshuai/stress-test/pkg/log"
	"gopkg.in/yaml.v2"
)

var (
	yamlPath   string
	perApiTime string
	apiToken   string
	logPath    string
)

type Stress struct {
	BaseAPI string      `yaml:"baseApi"`
	APIs    []StressAPI `yaml:"apis"`
}

type StressAPI struct {
	URL       string `yaml:"url"`
	APIName   string `yaml:"apiName"`
	ParamsStr string `yaml:"paramsStr"`
	Method    string `yaml:"method"`
}

func main() {
	// flag
	flag.StringVar(&yamlPath, "YAML_PATH", "", "性能测试 Yaml 路径")
	flag.StringVar(&perApiTime, "PER_API_TIME", "", "每个 API 接口")
	flag.StringVar(&apiToken, "TOKEN", "", " API 接口的 TOKEN")
	flag.StringVar(&logPath, "LOG_PATH", "stress-test.log", "log")
	flag.Parse()

	// check value
	if yamlPath == "" {
		log.Fatal("YAML_PATH not null")
	}

	if perApiTime == "" {
		log.Fatal("PER_API_TIME not null")
	}

	if apiToken == "" {
		log.Fatal("TOKEN not null")
	}

	// 初始化日志
	logger.InitLog(logPath)

	var stress Stress
	yamlS, readErr := ioutil.ReadFile(yamlPath)
	if readErr != nil {
		log.Fatal("FILE NOT EXIST")
	}
	// yaml解析的时候c.data如果没有被初始化，会自动为你做初始化
	err := yaml.Unmarshal(yamlS, &stress)
	if err != nil {
		log.Fatal("YAML 解析异常")
	}

	for _, api := range stress.APIs {

		fmt.Println("\r\n##########begin##############\r\n")
		fmt.Println(api.APIName + "\r\n")

		heyCmdExecute(api, stress.BaseAPI)
		time.Sleep(1 * time.Second)
		fmt.Println("\r\n##########end##############\r\n")
	}

}

func heyCmdExecute(api StressAPI, baseApi string) {
	paramData := fmt.Sprintf("params=%s&TOKEN=%s", api.ParamsStr, apiToken)
	urlData := fmt.Sprintf("%s%s", baseApi, api.URL)
	headerData := "Content-Type:application/x-www-form-urlencoded; charset=UTF-8"

	heyCmd := exec.Command("hey", "-m", "POST", "-H", headerData, "-d", paramData, "-z", "10s", urlData)
	var out1 bytes.Buffer
	heyCmd.Stdout = &out1
	err := heyCmd.Start()
	if err != nil {
		logger.Fatal(err)
	}
	err = heyCmd.Wait()
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(out1.String())
	logger.Info(out1.String())
}
