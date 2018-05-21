package data

// Transaction Structure
type Transaction struct {
	Hash        string   `json:"hash"`
	Ver         string   `json:"ver"`
	VinSz       string   `json:"vinsz"`
	VoutSz      string   `json:"voutsz"`
	Locktime    string   `json:"locktime"`
	Size        string   `json:"size"`
	Block       string   `json:"block"`
	Blocknumber string   `json:"blocknumber"`
	Time        string   `json:"timer"`
	Inputs      []Input  `json:"inputs"`
	Outputs     []Output `json:"outputs"`
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
