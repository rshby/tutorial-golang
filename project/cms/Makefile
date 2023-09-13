migrateup:
	migrate -path db/migration -database 'mysql://root:@tcp(localhost:3306)/cms' -verbose up

migratedown:
	migrate -path db/migration -database 'mysql://root:@tcp(localhost:3306)/cms' -verbose down

buildapp:
	go build -o ./bin/cms.exe ./main.go

.PHONY:
	migrateup, migratedown, buildapp