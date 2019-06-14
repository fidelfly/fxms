FROM alpine
ADD html /html
ADD config.toml /config.toml
ADD auth-web /auth-web
WORKDIR /
ENTRYPOINT [ "/auth-web" ]
