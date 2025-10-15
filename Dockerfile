FROM debian:stable-slim

# COPY src dst
COPY dockerfiles /bin/dockerfiles

ENV PORT=8010

# start the server
CMD ["/bin/dockerfiles"]
