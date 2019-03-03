FROM alpine
COPY ./dot-resolver /dot-resolver
ENTRYPOINT /dot-resolver
