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

| 算法             | 最坏情况       | 平均情况/期望  | Stable Sorting |
| ---------------- |:-------------:|:-------------:| --------------:|
| Bubble Sort      | O(n²)         | O(n²)         | Yes           |
| Insertion Sort   | O(n²)         | O(n²)         | Yes           |
| Merge Sort       | O(nlgn)       | O(nlgn)       | Yes           |
| Heap Sort        | O(nlgn)       | ---           | No            |
| Quick Sort       | O(n²)         | O(nlgn) 期望  | No            |
| Selection Sort   | O(n²)         | O(n²)         | No            |
| Counting Sort    | O(k+n)        | O(k+n)        | Yes           |
| Radix Sort       | O(d(k+n))     | O(d(k+n))     | Yes           |
| Bucket Sort      | O(n²)         | O(n) 平均情况  | Yes           |
