go get github.com/siniliako/gocommands/blob/main/modrunbuildtest/hello.go
	This last command is not possible because you need to use the package name to download the library.
	
go mod tidy // Remove unnecessary dependencies or clean the mod.go at the same level.
go: warning: "all" matched no packages

PS C:\projects\go\goCommands\getinstallfmt> go get github.com/google/uuid
	The dependency file configuration "mod.go" add a new line "require github.com/google/uuid v1.3.1 indirect"
	
	The last command downloaded the complete package uuid from the space of github.
	Url with the files with package "uuid": https://github.com/google/uuid

go mod tidy // Remove uuid because is unnecessary in go.mod.

go install 

go fmt notformatted.go // It isn't function, I didn't find how is possible to put an example

go vet hello.go

go mod init // It was used in the last project
go mod tidy // It was used in the last project

go get github.com/google/uuid
go mod download

go clean



