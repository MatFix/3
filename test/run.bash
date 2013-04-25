#! /bin/bash

files=*.go
files=$(echo $files | sed s/doc.go//g)

for f in $files; do
	go build $f || exit 1;
done;

for f in $files; do
	a=$(echo $f | sed s/.go//g)
	./$a -http="" -o $a.out -f -s || exit 1;
	rm $a
done;
