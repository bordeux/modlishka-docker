FROM golang:alpine

MAINTAINER Krzysztof Bednarczyk <admin@bordeux.net>

ENV MODLISHKA_REPOSITORY="github.com/drk1wi/Modlishka"
ENV INSTALL_PACKAGES="git make gcc musl-dev"

COPY ./run-server.sh /bin/run-server.sh

RUN set -ex \
		&& chmod +x /bin/run-server.sh \
		&& apk add --no-cache ${INSTALL_PACKAGES}\
        && go get -u ${MODLISHKA_REPOSITORY} \
        && cd $GOPATH/src/github.com/drk1wi/Modlishka/ && make \
		&& apk del ${INSTALL_PACKAGES} && rm -rf /var/cache/apk/*
		
CMD ["run-server.sh"]

EXPOSE 80 443