rem Make sure module is built for 32 bit architecture as Domino COM is only
rem available as 32 bit...
set GOARCH=386

rem Build program.
go build main.go
