package main

import (
	"encoding/json"
	"net/http"

	"github.com/coppetti/graphql-go/data"

	"github.com/graphql-go/graphql"
)

type Transaction data.Transaction
type Input data.Input
type Output data.Output

var tx = data.Transactions
var txi = data.Inputs
var txo = data.Outputs

func Filter(tx []data.Transaction, f func(data.Transaction) bool) []data.Transaction {
	vsf := make([]data.Transaction, 0)
	for _, v := range tx {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func main() {

	txiType := graphql.NewObject(graphql.ObjectConfig{
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

	txoType := graphql.NewObject(graphql.ObjectConfig{
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

	txType := graphql.NewObject(graphql.ObjectConfig{
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
				Type: txiType,
			},
			"outputs": &graphql.Field{
				Type: txoType,
			},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
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
					for _, t := range txi {
						if t.Hash == hash {
							return t, nil
						}
					}
					return nil, nil
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
					for _, t := range txo {
						if t.Hash == hash {
							return t, nil
						}
					}
					return nil, nil
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
					for _, t := range tx {
						if t.Hash == hash {
							return t, nil
						}
					}
					return nil, nil
				},
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":12345", nil)
}
