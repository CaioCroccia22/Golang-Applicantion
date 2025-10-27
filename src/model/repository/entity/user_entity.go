// Conversão do entity no banco
// Inserindo fora do domain para fins de organização pois era algo especifico do mongo db
package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	Id       primitive.ObjectID `bson:"_id, omitempty"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Age      int8               `bson:"age"`
}
