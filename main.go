package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/coppetti/graphql-go/data"

	"github.com/graphql-go/graphql"
)

type Song data.Song
type Artist data.Artist
type Album data.Album

var songs = data.Songs
var artists = data.Artists
var albums = data.Albums

func Filter(songs []data.Song, f func(data.Song) bool) []data.Song {
	vsf := make([]data.Song, 0)
	for _, v := range songs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func main() {
	songType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Song",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"album": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"duration": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	artistType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Artist",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	albumType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Album",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"artist": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"year": &graphql.Field{
				Type: graphql.String,
			},
			"genre": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"songs": &graphql.Field{
				Type: graphql.NewList(songType),
				Args: graphql.FieldConfigArgument{
					"album": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					album := params.Args["album"].(string)
					filtered := Filter(songs, func(v data.Song) bool {
						return strings.Contains(v.Album, album)
					})
					return filtered, nil
				},
			},
			"album": &graphql.Field{
				Type: albumType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["id"].(string)
					for _, album := range albums {
						if album.ID == id {
							return album, nil
						}
					}
					return nil, nil
				},
			},
			"artist": &graphql.Field{
				Type: graphql.NewList(artistType),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name := params.Args["name"].(string)
					for _, artist := range artists {
						if artist.Name == name {
							return artist, nil
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
