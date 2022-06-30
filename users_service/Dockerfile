FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on

RUN cd /build && git clone https://github.com/SibinStojanovicRA168-2018/user_service.git

RUN cd /build/grpc_test && go build

EXPOSE 9092

ENTRYPOINT [ "/build/grpc_test/main" ]