FROM alpine
ADD config.toml /config.toml
ADD user-web /user-web
WORKDIR /
ENTRYPOINT [ "/user-web" ]
