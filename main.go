package main

import (
	"encoding/json"
	"log"
	"net/http"

	graphiql "github.com/mnmtanish/go-graphiql"

	data "github.com/coppetti/graphql-go/data"
	"github.com/graphql-go/graphql"
)

func serveGraphQL(s graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendError := func(err error) {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}

		req := &graphiql.Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			sendError(err)
			return
		}

		res := graphql.Do(graphql.Params{
			Schema:        s,
			RequestString: req.Query,
			// RequestString: r.URL.Query().Get("query"), // to use with ?query={...}
		})

		if err := json.NewEncoder(w).Encode(res); err != nil {
			sendError(err)
		}
	}
}

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(
			data.CreateQuery(
				data.CreateTxType(data.CreateTxiType(), data.CreateTxoType()),
			),
		),
	})
	if err != nil {
		log.Fatalf("failed to create schema, error: %v", err)
	}

	http.HandleFunc("/", graphiql.ServeGraphiQL)
	http.HandleFunc("/graphql", serveGraphQL(schema))

	log.Println("Server started at http://localhost:3000/graphql")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// Query example
// {
// 	transaction(hash:"f4184fc596403b9d638783cf57adfe4c75c605f6356fbc91338530e9831e9e16"){
// 	  inputs{
// 		hash
// 		n
// 		scriptsig
// 	  }
// 	  outputs{
// 		value
// 		address
// 		scriptpubkey
// 		hash
// 		n
// 	  }
// 	  block
// 	  blocknumber
// 	  time
// 	}
//   }
