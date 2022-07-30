SQL_MIGRATE_CLI		= sql-migrate

export UID_GID=$(shell id -u):$(shell id -g)

help:		## Show this help.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

lint:
	staticcheck ./...
	
