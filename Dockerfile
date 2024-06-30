FROM golang:1.22.4-bullseye as build
RUN mkdir -p /scheduler
WORKDIR /scheduler
# Copy dependencies list
# Build with optional lambda.norpc tag
COPY . .

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
RUN go build scheduler-api
# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2023
ENV LAMBDA=true
COPY --from=build /scheduler/scheduler-api ./scheduler-api
COPY --from=build /scheduler/.env ./.env

ENTRYPOINT [ "./scheduler-api" ]