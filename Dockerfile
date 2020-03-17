FROM golang:alpine as build

ENV BUILDPATH /warpin-message

RUN mkdir -p $BUILDPATH
WORKDIR $BUILDPATH
COPY . /warpin-message

# You'll need git for glide install to work properly
RUN apk add --no-cache curl
RUN go build -o warpin-message .
#RUN ls -lah $BUILDPATH
RUN apk add --no-cache tzdata
ENV TZ Asia/Jakarta
RUN date

#ENTRYPOINT ["nohup","./warpin-message"]
CMD ["./warpin-message"]
EXPOSE 8888