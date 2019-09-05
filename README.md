# Golang Algorithms

## build

on Windows:

```cmd
set GOPROXY=https://goproxy.io
go test ./...
go test -cover ./...

cd search
go test -bench .

go build -o golang-algorithms.exe .
```

or Linux:

```shell
export GOPROXY=https://goproxy.io
go test ./...
go test -cover ./...

cd search
go test -bench .

go build -o golang-algorithms .
```

## 排序

| 算法             | 最坏情况       | 平均情况/期望  |  In Place     | Stable Sorting |
| ---------------- |:-------------:|:-------------:|:-------------:| --------------:|
| Bubble Sort      | O(n²)         | O(n²)         | Yes           | Yes            |
| Insertion Sort   | O(n²)         | O(n²)         | Yes           | Yes            |
| Merge Sort       | O(nlgn)       | O(nlgn)       | NO            | Yes            |
| Heap Sort        | O(nlgn)       | ---           | Yes           | No             |
| Quick Sort       | O(n²)         | O(nlgn) 期望   | Yes           | No             |
| Selection Sort   | O(n²)         | O(n²)         | Yes           | No             |
| Shell Sort       | O(nlg²n)      | O(nlgn)       | Yes           | No             |
| Counting Sort    | O(k+n)        | O(k+n)        | NO            | Yes            |
| Radix Sort       | O(d(k+n))     | O(d(k+n))     | NO            | Yes            |
| Bucket Sort      | O(n²)         | O(n) 平均      | NO            | Yes            |
