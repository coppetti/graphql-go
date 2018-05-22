package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/coppetti/graphql-go/data"
	graphiql "github.com/mnmtanish/go-graphiql"

	"github.com/graphql-go/graphql"
)

// Transaction data structure
type Transaction data.Transaction

// Input data structure
type Input data.Input

// Output data structure
type Output data.Output

var tx = data.Transactions
var txi = data.Inputs
var txo = data.Outputs

func createTxiType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Input",
		Fields: graphql.Fields{
			"hash": &graphql.Field{
				Type: graphql.String,
			},
			"n": &graphql.Field{
				Type: graphql.String,
			},
			"scriptsig": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func createTxoType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Output",
		Fields: graphql.Fields{
			"hash": &graphql.Field{
				Type: graphql.String,
			},
			"n": &graphql.Field{
				Type: graphql.String,
			},
			"scriptpubkey": &graphql.Field{
				Type: graphql.String,
			},
			"value": &graphql.Field{
				Type: graphql.String,
			},
			"address": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func createTxType(txiType *graphql.Object, txoType *graphql.Object) (*graphql.Object, *graphql.Object, *graphql.Object) {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Transaction",
		Fields: graphql.Fields{
			"hash": &graphql.Field{
				Type: graphql.String,
			},
			"ver": &graphql.Field{
				Type: graphql.String,
			},
			"block": &graphql.Field{
				Type: graphql.String,
			},
			"blocknumber": &graphql.Field{
				Type: graphql.String,
			},
			"time": &graphql.Field{
				Type: graphql.String,
			},
			"inputs": &graphql.Field{
				Type: graphql.NewList(txiType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if tx, ok := p.Source.(data.Transaction); ok {
						return tx.Txi, nil
					}
					return []interface{}{}, nil
				},
			},
			"outputs": &graphql.Field{
				Type: graphql.NewList(txoType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if tx, ok := p.Source.(data.Transaction); ok {
						return tx.Txo, nil
					}
					return []interface{}{}, nil
				},
			},
		},
	}), txiType, txoType
}

func createQuery(txType, txiType, txoType *graphql.Object) graphql.ObjectConfig {
	return graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"input": &graphql.Field{
				Type: txiType,
				Args: graphql.FieldConfigArgument{
					"hash": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					hash := params.Args["hash"].(string)
					return fetchInputsByHash(hash)
				},
			},
			"output": &graphql.Field{
				Type: txoType,
				Args: graphql.FieldConfigArgument{
					"hash": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					hash := params.Args["hash"].(string)
					return fetchOutputsByHash(hash)
				},
			},
			"transaction": &graphql.Field{
				Type: txType,
				Args: graphql.FieldConfigArgument{
					"hash": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					hash := params.Args["hash"].(string)
					return fetchTransactionByHash(hash)
				},
			},
		},
	}
}

func fetchInputsByHash(hash string) (interface{}, error) {
	for _, v := range txi {
		if v.Hash == hash {
			return v, nil
		}
	}
	return nil, nil
}

func fetchOutputsByHash(hash string) (interface{}, error) {
	for _, v := range txo {
		if v.Hash == hash {
			return v, nil
		}
	}
	return nil, nil
}

func fetchTransactionByHash(hash string) (interface{}, error) {
	for _, v := range tx {
		if v.Hash == hash {
			return v, nil
		}
	}
	return nil, nil
}

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

	txType, txiType, txoType := createTxType(createTxiType(), createTxoType())

	rootQuery := graphql.NewObject(createQuery(txType, txiType, txoType))

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(
			createQuery(
				createTxType(createTxiType(), createTxoType()),
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
