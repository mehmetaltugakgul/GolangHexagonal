mocks:
	mockgen -source internal/app/handler/auth.go -package mocks -destination internal/app/handler/mocks/auth_service_mock.go
	mockgen -source internal/app/handler/user.go -package mocks -destination internal/app/handler/mocks/user_service_mock.go
	mockgen -source internal/app/repository/user.go -package mocks -destination internal/app/service/mocks/user_repository_mock.go

test:
	go test -v -cover ./...