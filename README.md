# firstmock

go install github.com/golang/mock/mockgen@latest

mockgen -source=db.go -destination=mocks/db_mock.go -package=mocks

ref : https://www.liwenzhou.com/posts/Go/golang-unit-test-3/