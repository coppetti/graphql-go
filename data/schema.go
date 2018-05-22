package data

import (
	"github.com/graphql-go/graphql"
)

// Transaction Structure
type Transaction struct {
	Txi         []Input  `json:"inputs"`
	Txo         []Output `json:"outputs"`
	Hash        string   `json:"hash"`
	Ver         string   `json:"ver"`
	VinSz       string   `json:"vinsz"`
	VoutSz      string   `json:"voutsz"`
	Locktime    string   `json:"locktime"`
	Size        string   `json:"size"`
	Block       string   `json:"block"`
	Blocknumber string   `json:"blocknumber"`
	Time        string   `json:"timer"`
}

// Input structure
type Input struct {
	Hash      string `json:"hash"`
	N         string `json:"n"`
	ScriptSig string `json:"scriptsig"`
}

// Output Structure
type Output struct {
	Value        string `json:"value"`
	ScriptPubKey string `json:"scriptpubkey"`
	Address      string `json:"address"`
	Hash         string `json:"hash"`
	N            string `json:"n"`
}

func CreateTxiType() *graphql.Object {
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

func CreateTxoType() *graphql.Object {
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

func CreateTxType(txiType *graphql.Object, txoType *graphql.Object) (*graphql.Object, *graphql.Object, *graphql.Object) {
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
					if tx, ok := p.Source.(Transaction); ok {
						return tx.Txi, nil
					}
					return []interface{}{}, nil
				},
			},
			"outputs": &graphql.Field{
				Type: graphql.NewList(txoType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if tx, ok := p.Source.(Transaction); ok {
						return tx.Txo, nil
					}
					return []interface{}{}, nil
				},
			},
		},
	}), txiType, txoType
}

func CreateQuery(txType, txiType, txoType *graphql.Object) graphql.ObjectConfig {
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
	for _, v := range Inputs {
		if v.Hash == hash {
			return v, nil
		}
	}
	return nil, nil
}

func fetchOutputsByHash(hash string) (interface{}, error) {
	for _, v := range Outputs {
		if v.Hash == hash {
			return v, nil
		}
	}
	return nil, nil
}

func fetchTransactionByHash(hash string) (interface{}, error) {
	for _, v := range Transactions {
		if v.Hash == hash {
			return v, nil
		}
	}
	return nil, nil
}
