FROM alpine:latest

MAINTAINER Chris Robertson https://github.com/electronicsleep

RUN mkdir -p /usr/src/app

RUN apk update
RUN apk add bash

ADD chrisgr /usr/src/app
ADD public /usr/src/app/public

WORKDIR /usr/src/app
EXPOSE 8080

CMD ["./chrisgr"]
#CMD ["/bin/bash"]
