FROM golang:1.22.4-bullseye as build
RUN mkdir -p /scheduler
WORKDIR /scheduler
# Copy dependencies list
# Build with optional lambda.norpc tag
COPY . .
RUN ls -al
#RUN go get -d ./...

#ARG DB_USERNAME
#ARG DB_HOST
#ARG DB_PASSWORD
#ARG DB_NAME

#ENV user=$DB_USERNAME
#ENV pass=$DB_PASSWORD
#ENV host=$DB_HOST
#ENV port=3500
    #		os.Getenv("port"),
    #		os.Getenv("db_name"))

RUN go get -t -x
RUN go build  -o main main.go
# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2023

COPY --from=build /scheduler/main ./main
COPY --from=build /scheduler/.env ./main
EXPOSE 3500

ENTRYPOINT [ "./main" ]