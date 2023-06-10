# Hi dear! it is my Makefile
DATABASE_URL = $(MYSQL_ADDRS)
PROJECT_DIR = $(shell pwd)
APPS_DIR = $(PROJECT_DIR)/services

CMDS = $(shell find $(PROJECT_DIR)/services/*/cmd -mindepth 1 -maxdepth 1 -type d)
STATIC_TARGETS = $(addprefix static-,$(CMDS))

MIGRATIONS = $(shell find $(PROJECT_DIR)/services/*/migrations -mindepth 0 -maxdepth 0 -type d)
MIGRATION_TARGETS = $(addprefix migrate-,$(MIGRATIONS))

GOQLGEN = github.com/99designs/gqlgen
WIREGEN = github.com/google/wire/cmd/wire
GRAPHQL_LINTER = graphql-schema-linter

.PHONY: migrate
migrate: $(MIGRATION_TARGETS)
$(MIGRATION_TARGETS):
	$(eval FILE=$(subst migrate-,,$@))
	-migrate -source "file:$(FILE)" -database "mysql://$(DATABASE_URL)" $(move) $(step)

di:
	go run -mod=mod $(WIREGEN) gen $(APPS_DIR)/...

static: $(STATIC_TARGETS)
$(STATIC_TARGETS):
	$(eval FULLPATH=$(subst static-,,$@))
	$(eval CMD=$(shell basename ${FULLPATH}))
	$(eval TEMPATH=$(subst ${CMD},,${FULLPATH}))
	$(eval OUT=$(shell dirname ${TEMPATH})) 
	env GCO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(OUT)/out/$(CMD) $(FULLPATH)
