Command instructions to class of go Commands and functions:

go help 
go mod hello (Create an initial module. The purpose is to manage dependencies like npm with package-json.
	--> Create a file go.mod file with the next information:
	--
		module hello

		go 1.21.0
	--
go run hello.go
go build -o my-program.exe hello.go
	For windows: 	setx GOOS windows (setx GOOS <targetOperatingSystem> in local environment)
					setx GOARCH amd64 (setx GOARCH <Target Platform> in local environment)
					
					set GOOS=windows
					set GOARCH=amd64
					(command set temporal and only in the same console) 
					(command setx permanent and only in the same console)
	For linux:		   
					cd /usr/local/go/src
					sudo GOOS=linux GOARCH=arm ./make.bash --no-clean
					env GOOS=linux GOARCH=arm go build -o <target_executable> <source_file>
					https://gist.github.com/zenhob/f70bbb802e3e6729a1a6

	List of operating system and Target platform: https://gist.github.com/zfarbp/121a76d5a3fde562c3955a606a9d6fcc
	More detail setx: https://www.opentechguides.com/how-to/article/windows-11/222/path-environment-variable.html#:~:text=Use%20the%20setx%20command%20to,command%20to%20use%20is%20set.&text=c%3A%5C%3E%20setx%20TESTVAR%20someValue%20SUCCESS%3A%20Specified%20value%20was%20saved.

go test

	C:\projects\go\awesomeproject>go test 
	
go test -v
	C:\projects\go\awesomeproject>go test -v
	=== RUN   TestHello
	--- PASS: TestHello (0.00s)
	PASS
	ok      hello   0.048s
	
