FROM scratch

ENV ${[( .EnvPrefix | ToUpper )]}_LOCAL_HOST 0.0.0.0
ENV ${[( .EnvPrefix | ToUpper )]}_LOCAL_PORT 8080
ENV ${[( .EnvPrefix | ToUpper )]}_LOG_LEVEL 0

EXPOSE $${[( .EnvPrefix | ToUpper )]}_LOCAL_PORT

COPY certs /etc/ssl/
COPY bin/linux-amd64/${[( .serviceName )]} /

CMD ["/${[( .serviceName )]}"]
