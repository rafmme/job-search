VERSION ?= 0.0.1  
NAME ?= "Job Search Tool"  
AUTHOR ?= "Timileyin Farayola"   
  
.PHONY: build run fp p

build:  
	go build

run:
	go run main.go

fp:
	git push -f

p:
	git push


DEFAULT: build
