
fmt:
	go fmt pkg/ch00-search/*
	go fmt pkg/ch01-recursion/*
	go fmt pkg/ch02-sort/*
	go fmt pkg/ch03-hashtb/*
	go fmt pkg/ch04-graph/*
	go fmt pkg/ch05-greedy/*
	go fmt pkg/ch06-knn/*

vet:
	go vet pkg/ch00-search/*
	go vet pkg/ch01-recursion/*
	go vet pkg/ch02-sort/*
	go vet pkg/ch03-hashtb/*
	go vet pkg/ch04-graph/*
	go vet pkg/ch05-greedy/*
	go vet pkg/ch06-knn/*

test:
	go test pkg/ch00-search/*
	go test pkg/ch01-recursion/*
	go test pkg/ch02-sort/*
	go test pkg/ch03-hashtb/*
	go test pkg/ch04-graph/*
	go test pkg/ch05-greedy/*
	go test pkg/ch06-knn/*

clean:
	rm bin/*
