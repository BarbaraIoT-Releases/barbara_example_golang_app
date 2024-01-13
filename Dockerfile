FROM golang

WORKDIR /src

RUN apt update && \
    apt install -y  zip

COPY src /src
COPY build.sh /bin/native_build

RUN chmod +x /bin/native_build

CMD ["/bin/native_build"]

