package main

import (
	"cyntra/user/rpc/types/user"
	"fmt"

	"github.com/bojand/ghz/runner"
	"github.com/jhump/protoreflect/desc"
	"google.golang.org/protobuf/proto"
)

func dataFunc(mtd *desc.MethodDescriptor, cd *runner.CallData) []byte {
	msg := &user.VerifyTokenRequest{
		AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQzNTU2NTUsIlVzZXJJZCI6IjcwMTE1NGE1LTNiM2MtNGFlNi04ODY5LWQ3MDBkNTIzZGMxMSJ9.UOHncmvm8VYcO0JDKRWjHzgATwsjedNXPoG0Drh5CzU",
	}

	binData, _ := proto.Marshal(msg)

	return binData
}

func main() {
	report, err := runner.Run(
		"user.User.verifyToken",
		"127.0.0.1:8080",
		runner.WithProtoFile("../user.proto", []string{}),
		runner.WithInsecure(true),
		runner.WithBinaryDataFunc(dataFunc),
		runner.WithTotalRequests(1000),
		runner.WithConcurrency(8000),
		runner.WithAsync(true),
	)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("RPS: ", report.Rps)
		fmt.Println("Count: ", report.Count)
		fmt.Println("Slowest: ", report.Slowest)
		fmt.Println("Fastest: ", report.Fastest)
		fmt.Println("Average: ", report.Average)
		fmt.Println("Total: ", report.Total)
		fmt.Println("Concurrency: ", report.Options.Concurrency)
		fmt.Println("EndReason: ", report.EndReason)

	}

}
