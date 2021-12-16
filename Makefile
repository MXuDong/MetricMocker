# Make file

# Build docker image
.PHONE: docker
docker-build-%: Dockerfile
	@docker build -f Dockerfile -t mxudong/metric-mocker:v$* .

output:
	@echo 'mkdir output'

metric-mocker-%.tar: output docker-build-%
	@docker save mxudong/metric-mocker:v$* > output/$@