FROM ubuntu
ADD config/config.yaml config/config.yaml 
ADD bin/amd64/httpserver /httpserver
ENTRYPOINT /httpserver
