Framework Training Template GO Microservice
============================================
Overview
--------
This is a very simple micro service that can be built and packaged into a docker image. You should use this micro service as a template for building your own micro services. This microservice uses the Go.

How to build and run this microservice
--------------------------------------

1. Build the docker image

   ```
   $ docker build -t template-microservice-go .
   ```
2. Run the container

   ```
   $ docker run -dp 8080:8080 template-microservice-go
   ```
3. Test the micro service

   ```
   $ curl http://localhost:5000
   Success! The Framework Training GO template microservices is up and running!
   ```