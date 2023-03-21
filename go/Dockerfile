FROM golang:latest

# Set the working directory to the project root
WORKDIR /app

# Copy the source code to the working directory
COPY . .

# Install dependencies
# RUN go get -d -v

# add go to path
ENV PATH /go/bin:$PATH


# install postgres 

RUN  apt-get update && apt-get install -y postgresql-client

# install postgis
RUN apt-get update && apt-get install -y postgis

# install gdal
RUN apt-get update && apt-get install -y gdal-bin



# Expose the port that the app listens on
EXPOSE 8080

# build the app
RUN go build -o app .