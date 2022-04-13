#!/usr/bin/env bash

mockgen -source internal/domain/user/repository/repository.go -destination internal/domain/user/repository/mocks/repository_mock.go
mockgen -source internal/domain/user/publisher/publisher.go -destination internal/domain/user/publisher/mocks/publisher_mock.go
mockgen -source internal/domain/user/service/service.go -destination internal/domain/user/service/mocks/service_mock.go
