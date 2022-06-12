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
		AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQ0NDgyOTQsIlVzZXJJZCI6ImU5OTFhNjYxLWU3NmUtNDE2OS04NTFiLWQyZTliMTZiNmNmZiJ9.kUM0Gsc5G4M299AJ-i3Fm8RG4OKysqRG0_lr66BmhOs",
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
		runner.WithTotalRequests(10000),
		runner.WithConcurrency(1000),
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

// ghz --insecure \                                                                                                                                                                                              ✹ ✚main
//   --async \
//   -n 5000 \
//   -c 1000 \
//   --proto ./user/rpc/user.proto \
//   --call user.User.verifyToken \
//   -d '{"access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQ0NDgyOTQsIlVzZXJJZCI6ImU5OTFhNjYxLWU3NmUtNDE2OS04NTFiLWQyZTliMTZiNmNmZiJ9.kUM0Gsc5G4M299AJ-i3Fm8RG4OKysqRG0_lr66BmhOs"}' \
//   0.0.0.0:8080
