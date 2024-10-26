MIGRATE UP:
migrate -path ./migrations -database "mysql://root:@tcp(localhost:3306)/article_db" -verbose up 
MIGRATE DOWN: 
migrate -path ./migrations -database "mysql://root:@tcp(localhost:3306)/article_db" -verbose down
BUILD:
GO BUILD