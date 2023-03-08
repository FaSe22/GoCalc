FROM  golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get https://github.com/FaSe22/GoCalc/main
RUN cd /build && git clone https://github.com/FaSe22/GoCalc.git

RUN cd /build/GoCalc/main && go build

EXPOSE 8080

ENTRYPOINT ["build/GoCalc/main/main"]
