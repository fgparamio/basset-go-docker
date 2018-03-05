# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev
WORKDIR /app
# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/go/src/github.com/fgparamio/basset-go-docker
# Add the source code:
ADD . $SRC_DIR
# Build it:

RUN cd $SRC_DIR; go get -v github.com/gorilla/handlers
RUN cd $SRC_DIR; go get -v github.com/gorilla/mux
RUN cd $SRC_DIR; go build -o basset-back; cp basset-back /app/
ENTRYPOINT ["./basset-back"]
EXPOSE 4000