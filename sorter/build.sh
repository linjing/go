set -x
export GOPATH=$PWD

go build algorithms/qsort/qsort.go
go build algorithms/bubblesort/bubblesort.go
go test algorithms/qsort
go test algorithms/bubblesort

go build sorter
