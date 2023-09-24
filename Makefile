run:
	swag init --parseDependency && CompileDaemon -command="./homelab-backend"

run-linux:
	/home/kaizer/go/bin/CompileDaemon -command="./homelab-backend"

swag:
	swag init --parseDependency