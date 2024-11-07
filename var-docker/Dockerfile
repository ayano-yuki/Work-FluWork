FROM gcc:latest
WORKDIR /app
COPY . /app
RUN apt-get update && apt-get install -y make
CMD ["make"]
