package bd

// ------------------------------------- //
//             IMPORTACIONES             //
// ------------------------------------- //
import(

	"context"
	"log" // para grabar nombre - texto en el log de ejecución	
	// IMPORTAMOS PAQUETES INSTALADOS //
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ------------------------------------- //
//               VARIABLES               //
// ------------------------------------- //
// variable de conexíón
var MongoC = conectarBd() 
// funcion del paquete de Mongo que permite setear la URL de la BD
var clientOptions = options.Client().ApplyURI( "mongodb+srv://mongoPoncho:alfonso123@tweetpost.plrsj.mongodb.net/test" )  

// ------------------------------------- //
//             CONEXIÓN BD               //
// ------------------------------------- //

// devuelve un objeto de tipo mongo client
func conectarBd() *mongo.Client {

	// hacemos la conexión a la base de datos
	// toma la conexión que tiene la variable clientOptions que tiene la URL de la base de datos
	client, err := mongo.Connect( context.TODO(), clientOptions )
}




/*
en go no existen lo que son las variables globales, el contexto, es un espacio en memoria donde podré ir compartiendo
cosas  



*/