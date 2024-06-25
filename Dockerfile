FROM golang:1.22.4-bullseye as build
RUN mkdir -p /scheduler
WORKDIR /scheduler
# Copy dependencies list
# Build with optional lambda.norpc tag
COPY . .
RUN ls -al
#RUN go get -d ./...
RUN go get -t -x
RUN go build  -o main main.go
# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /scheduler/main ./main
ENTRYPOINT [ "./main" ]