###############################
# Matsunaga Project 2022 - Makefile
#
# 便利コマンド集
###############################
BIN_DIR:=bin

ROOT_PACKAGE:=$(shell go list .)

DB_PATH:="root:root@/shanks?parseTime=true"

###############################
# Go Generals
###############################

## Go moduleの依存関係を最新のものにする。定期的に実行しておくと良い。
update_module:
	go mod tidy

###############################
# Install Development tools
#
# protoc-gen-go -> gRPCでprotoファイルからgoファイルを生成するためのツール
# goose -> DBのマイグレーションツール
###############################

## 開発に必要なツール群のインストール
install_tools: $(BIN_DIR)/goose $(BIN_DIR)/protoc-gen-go $(BIN_DIR)/make2help $(BIN_DIR)/table-to-go

$(BIN_DIR)/protoc-gen-go: go.sum
	@go build -o $@ google.golang.org/protobuf/cmd/protoc-gen-go

$(BIN_DIR)/goose: go.sum
	@go build -o $@ github.com/pressly/goose/v3/cmd/goose

$(BIN_DIR)/make2help: go.sum
	@go build -o $@ github.com/Songmu/make2help/cmd/make2help

$(BIN_DIR)/table-to-go: go.sum
	@go build -o $@ github.com/fraenky8/tables-to-go
###############################
# Make helps
#
###############################

## コマンド一覧の表示
help:
	$(BIN_DIR)/make2help $(MAKEFILE_LIST)
###############################
# DB Migrations
#
###############################

## 新たなマイグレーションファイルを作成。name=名前でマイグレーションの名前を指定する。
migrate_create:
	$(BIN_DIR)/goose -dir ./migrations create $(name) sql

## DBのマイグレーションを最新のバージョンにする。
migrate_up:
	$(BIN_DIR)/goose -dir ./migrations mysql $(DB_PATH) up

## DBのマイグレーションを1段階バージョンダウンする。
migrate_down:
	$(BIN_DIR)/goose -dir ./migrations mysql $(DB_PATH) down
###############################
# Docker containers
#
###############################

## ShanksのDBを起動する。
db_local:
	docker compose up -d db

###############################
# Auto generate DAO
#
###############################
## Shanksのテーブル構造をDAOのModelとして自動生成する。※現在は使用非推奨
dao_gen:
	$(BIN_DIR)/table-to-go -v \
		-t mysql \
		-h localhost \
		-d otsunaga \
		-u root \
		-p root \
		-pn model \
		-of ./internal/models/dao/model

###############################
# Go tests
#
###############################
## internal以下のテストを実行する。
test:
	go test ./internal/...