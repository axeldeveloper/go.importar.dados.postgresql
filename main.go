package main
 
import (
	"fmt"
	pg "github.com/axeldeveloper/go.importar.dados.postgresql/postgres"

)

func main() {
	fmt.Printf("iniciando \n");
	pg.Run();
}