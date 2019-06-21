
package models

import (
	"gopkg.in/mgo.v2/bson"
)

/**
  Estructura que almacena la estructura y los datos necesarios para el dna
*/
type Dna struct {
	ID         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty" `
	N          int           `json:"N,omitempty" bson:"n,omitempty" `
	Size       int           `json:"Size,omitempty" bson:"size,omitempty"`
	Pattern    int           `json:"Pattern,omitempty" bson:"pattern,omitempty"`
	Sequence   []rune        `json:"Sequence,omitempty"  bson:"sequence,omitempty"`
	Mutant   bool          `json:"IsMutant,omitempty"  bson:"IsMutant,omitempty"`
	Response   Response      `json:"Response,omitempty"  bson:"response,omitempty"`
	Iterations int           `json:"iterations,omitempty"  bson:"iterations,omitempty"`
}

