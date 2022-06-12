user_svc_build:
	docker build -f ./user/rpc/Dockerfile -t cyntra-user-svc:latest .

user_svc_load_test: 
	cd ./user/rpc/test; go run ./loadTest.go