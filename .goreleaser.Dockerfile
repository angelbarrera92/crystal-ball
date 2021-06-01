FROM alpine:3.13.5

COPY crystal-ball /bin/
RUN mkdir -p /orakuru/etc
WORKDIR /orakuru

CMD ["/bin/crystal-ball"]
