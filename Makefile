VERSION ?= 0.0.1  
NAME ?= "Job Search Tool"  
AUTHOR ?= "Timileyin Farayola"   
  
.PHONY: build run start fp p

build:  
	go build -o ./job-search

run:
	go run main.go

start:
	make && ./job-search 2>&1 | tee job_app_logs.txt

fp:
	git push -f

p:
	git push


DEFAULT: build
