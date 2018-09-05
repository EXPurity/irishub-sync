package document

import "gopkg.in/mgo.v2/bson"

const CollectionNmTxMsg = "tx_msg"

type TxMsg struct {
	Hash    string `bson:"hash"`
	Content string `bson:"content"`
}

func (m TxMsg) Name() string {
	return CollectionNmTxMsg
}

func (m TxMsg) PkKvPair() map[string]interface{} {
	return bson.M{"hash": m.Hash}
}
