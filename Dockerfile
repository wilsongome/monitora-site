FROM golang


WORKDIR /go/src

CMD /bin/bash -c "echo 'Hello World'; sleep infinity"