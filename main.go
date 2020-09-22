package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("/tmp/pod-tzzh-aws.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	for {
		message := ReadMessage()
		if message.Op == "describe" {
			response := &DescribeResponse{
				Format: "json",
				Namespaces: []Namespace{
					{Name: "pod.tzzh.dynamodb",
						Vars: []Var{
							{Name: "batch-get-item"},
							{Name: "batch-write-item"},
							{Name: "describe-table"},
							{Name: "get-item"},
							{Name: "list-tables"},
						},
					},
				},
			}
			WriteDescribeResponse(response)

		} else if message.Op == "invoke" {

			switch message.Var {
			case "pod.tzzh.dynamodb/batch-get-item":
				res, err := BatchGetItem(message)
				if err != nil {
					WriteErrorResponse(message, err)
				} else {
					WriteInvokeResponse(message, res)
				}
			case "pod.tzzh.dynamodb/batch-write-item":
				res, err := BatchWriteItem(message)
				if err != nil {
					WriteErrorResponse(message, err)
				} else {
					WriteInvokeResponse(message, res)
				}
			case "pod.tzzh.dynamodb/describe-table":
				res, err := DescribeTable(message)
				if err != nil {
					WriteErrorResponse(message, err)
				} else {
					WriteInvokeResponse(message, res)
				}
			case "pod.tzzh.dynamodb/get-item":
				res, err := GetItem(message)
				if err != nil {
					WriteErrorResponse(message, err)
				} else {
					WriteInvokeResponse(message, res)
				}
			case "pod.tzzh.dynamodb/list-tables":
				res, err := ListTables(message)
				if err != nil {
					WriteErrorResponse(message, err)
				} else {
					WriteInvokeResponse(message, res)
				}
			}
		}
	}
}
