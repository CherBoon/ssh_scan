FROM ruby:alpine
MAINTAINER Jonathan Claudius
ENV PROJECT=github.com/mozilla/ssh_scan

RUN mkdir /app
ADD . /app
WORKDIR /app
CMD ["chmod +x /app/gosshscan"]

# required for ssh-keyscan
RUN apk --update add openssh-client

ENV GEM_HOME /usr/local/bundle/ruby/$RUBY_VERSION
ENV PATH="/app:${PATH}"
RUN apk --update add --virtual build-dependencies build-base && \
    bundle install && \
    apk del build-dependencies build-base && \
    rm -rf /var/cache/apk/*
EXPOSE 13337
# CMD /app/bin/ssh_scan
ENTRYPOINT ["./gosshscan"]