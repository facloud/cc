.PHONY: clean test

cc-pm: main.go
	go build -o cc-pm .

test:
	ginkgo -r -p

clean:
	rm -f cc-pm
