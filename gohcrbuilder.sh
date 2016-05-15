#!/usr/bin/bash
go build main.go
#safe-check, we want to delete existing gohcr executable file and replace it with a new one.
rm gohcr 2> /dev/null # should there be errors, redirect to void.
mv main gohcr
