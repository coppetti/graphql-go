FROM golang:onbuild

RUN go get github.com/mnmtanish/go-graphiql
RUN go get github.com/graphql-go/graphql

WORKDIR /go/src/github.com/coppetti/graphql-go/graphql-go

RUN git clone https://github.com/coppetti/graphql-go.git .

RUN go build -o main .

CMD ["./main"]

EXPOSE 3000


# Running
# docker run -p 3000:3000 <image id>

# Sample Query
# {
# 	transaction(hash:"f4184fc596403b9d638783cf57adfe4c75c605f6356fbc91338530e9831e9e16"){
# 	  inputs{
# 		hash
# 		n
# 		scriptsig
# 	  }
# 	  outputs{
# 		value
# 		address
# 		scriptpubkey
# 		hash
# 		n
# 	  }
# 	  block
# 	  blocknumber
# 	  time
# 	}
#   }