package data

// {
// 	"hash":"f4184fc596403b9d638783cf57adfe4c75c605f6356fbc91338530e9831e9e16",
// 	"ver":1,
// 	"vin_sz":1,
// 	"vout_sz":2,
// 	"lock_time":0,
// 	"size":275,
// 	"in":[
// 	  {
// 		"prev_out":{
// 		  "hash":"0437cd7f8525ceed2324359c2d0ba26006d92d856a9c20fa0241106ee5a597c9",
// 		  "n":0
// 		},
// 		"scriptSig":"304402204e45e16932b8af514961a1d3a1a25fdf3f4f7732e9d624c6c61548ab5fb8cd410220181522ec8eca07de4860a4acdd12909d831cc56cbbac4622082221a8768d1d0901"
// 	  }
// 	],
// 	"out":[
// 	  {
// 		"value":"10.00000000",
// 		"scriptPubKey":"04ae1a62fe09c5f51b13905f07f06b99a2f7159b2225f374cd378d71302fa28414e7aab37397f554a7df5f142c21c1b7303b8a0626f1baded5c72a704f7e6cd84c OP_CHECKSIG",
// 		"address":"1Q2TWHE3GMdB6BZKafqwxXtWAWgFt5Jvm3",
// 		"next_in":{
// 		  "hash":"ea44e97271691990157559d0bdd9959e02790c34db6c006d779e82fa5aee708e",
// 		  "n":0
// 		}
// 	  },
// 	  {
// 		"value":"40.00000000",
// 		"scriptPubKey":"0411db93e1dcdb8a016b49840f8c53bc1eb68a382e97b1482ecad7b148a6909a5cb2e0eaddfb84ccf9744464f82e160bfa9b8b64f9d4c03f999b8643f656b412a3 OP_CHECKSIG",
// 		"address":"12cbQLTFMXRnSzktFkuoG3eHoMeFtpTu3S"
// 	  }
// 	],
// 	"block": "00000000d1145790a8694403d4063f323d499e655c83426834d4ce2f8dd4a2ee",
// 	"blocknumber": 170,
// 	"time": "2009-01-12 04:30:25"
// }

var Inputs []Input = []Input{
	Input{
		Hash:      "0437cd7f8525ceed2324359c2d0ba26006d92d856a9c20fa0241106ee5a597c9",
		N:         "0",
		ScriptSig: "304402204e45e16932b8af514961a1d3a1a25fdf3f4f7732e9d624c6c61548ab5fb8cd410220181522ec8eca07de4860a4acdd12909d831cc56cbbac4622082221a8768d1d0901",
	},
}
var Outputs []Output = []Output{
	Output{
		Value:        "10.0000",
		ScriptPubKey: "04ae1a62fe09c5f51b13905f07f06b99a2f7159b2225f374cd378d71302fa28414e7aab37397f554a7df5f142c21c1b7303b8a0626f1baded5c72a704f7e6cd84c OP_CHECKSIG",
		Address:      "1Q2TWHE3GMdB6BZKafqwxXtWAWgFt5Jvm3",
		Hash:         "ea44e97271691990157559d0bdd9959e02790c34db6c006d779e82fa5aee708e",
		N:            "0",
	},
	{
		Value:        "40.0000",
		ScriptPubKey: "0411db93e1dcdb8a016b49840f8c53bc1eb68a382e97b1482ecad7b148a6909a5cb2e0eaddfb84ccf9744464f82e160bfa9b8b64f9d4c03f999b8643f656b412a3 OP_CHECKSIG",
		Address:      "12cbQLTFMXRnSzktFkuoG3eHoMeFtpTu3S",
	},
}

var Transactions []Transaction = []Transaction{
	Transaction{
		Hash:        "f4184fc596403b9d638783cf57adfe4c75c605f6356fbc91338530e9831e9e16",
		Ver:         "1",
		VinSz:       "1",
		VoutSz:      "2",
		Locktime:    "0",
		Size:        "275",
		Txi:         Inputs,
		Txo:         Outputs,
		Block:       "00000000d1145790a8694403d4063f323d499e655c83426834d4ce2f8dd4a2ee",
		Blocknumber: "170",
		Time:        "2009-01-12 04:30:25",
	},
}
