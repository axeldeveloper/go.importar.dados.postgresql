package postgres

import (
	"fmt"
	"log"
	_ "github.com/lib/pq"
	models "github.com/axeldeveloper/go.importar.dados.postgresql/models"
	"os"
	"github.com/360EntSecGroup-Skylar/excelize"
    "time"
    "strconv"

)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func CreateCSV(){

	fmtData := "2 Jan 2006 15:04:05"
	fmt.Printf("Begin %s \n",  time.Now().Format(fmtData) )
	members, err := models.BuscaTodos()	
	if err != nil {
		log.Fatal("Read failed:", err.Error())
	}
	fmt.Printf("Read %d rows successfully.\n", len(members))

	//f, e := os.Create("../log1.log")
	f, e := os.Create("./test.csv")
	if e != nil {
        panic(e)
	}	
	exin := 1
	for _, value := range members {
		exin++	
		q := ` %d; %s; %s; %s; %s; %s; %s; %s; %t; %s; %d; %d; %d; %d; %d`
		sq := fmt.Sprintf(q,
			value.ID.Int64, 
			value.Nome.String ,
			value.Apelido.String ,
			value.Cpf.String ,
			value.Pai.String,
			value.Mae.String, 
			value.Nascimento.String, 
			value.Sexo.String,
			value.Status,
			value.Date_joined,
			value.Estabelecimento_id.Int64,
			value.Religiao_id.Int64,
			value.Tratamento_id.Int64,
			value.Usuario_id.Int64,
			value.Vinculo_id.Int64,
			
		)
		_, err := f.WriteString( sq + "\n")
		check(err)
		time.Now()
		fmt.Printf("End %s \n",  time.Now().Format(fmtData) )
		//fmt.Println(n2)
	}
	
	
	defer f.Close()
	// Issue a `Sync` to flush writes to stable storage.
	f.Sync()

	// check error here...
    //exec.Command("/bin/sh", "-c", "echo "+e.Error()+" >> ../log1.log").Run()
	
}

func ReadPessoa(){	

	logger, e := os.Create("./pessoa.log")
	check(e)
	
	fmtData := "2 Jan 2006 15:04:05"
	fmt.Printf("Begin %s \n",  time.Now().Format(fmtData) )
	
	_, lerr := logger.WriteString( "Begin " + time.Now().Format(fmtData) + "\n")
	if lerr != nil {
		log.Fatal("Read failed:", lerr.Error())
	}


	members, err := models.BuscaTodos()	
	if err != nil {
		log.Fatal("Read failed:", err.Error())
	}
	fmt.Printf("Read %d rows successfully.\n", len(members))	
	
	_, lerr = logger.WriteString( "Total register :" + strconv.Itoa( len(members)) + "\n")
	if lerr != nil {
		check(lerr)
	}
	
	fe := excelize.NewFile()
    // Create a new sheet.
    index := fe.NewSheet("Pessoa")
    // Set value of a cell.
	fe.SetCellValue("Pessoa", "A1", "ID")
	fe.SetCellValue("Pessoa", "B1", "Nome")
	fe.SetCellValue("Pessoa", "C1", "Apelido")
	fe.SetCellValue("Pessoa", "D1", "Cpf")
	fe.SetCellValue("Pessoa", "E1", "Pai")
	fe.SetCellValue("Pessoa", "F1", "Mae")
	fe.SetCellValue("Pessoa", "G1", "Nascimento")
	fe.SetCellValue("Pessoa", "H1", "Sexo")	
	fe.SetCellValue("Pessoa", "I1", "Status")	
	fe.SetCellValue("Pessoa", "J1", "Date_joined")
	fe.SetCellValue("Pessoa", "K1", "Estabelecimento_id")
	fe.SetCellValue("Pessoa", "L1", "Religiao_id")
	fe.SetCellValue("Pessoa", "M1", "Tratamento_id")
	fe.SetCellValue("Pessoa", "N1", "Usuario_id")	
	fe.SetCellValue("Pessoa", "O1", "Vinculo_id")
    fe.SetCellValue("Endereco", "A1", 100)
    // Set active sheet of the workbook.
    fe.SetActiveSheet(index)
    
	exin := 1
	for _, value := range members {
		exin++
		fe.SetCellValue("Pessoa", "A" + strconv.Itoa(exin), strconv.FormatInt(value.ID.Int64, 10))
		fe.SetCellValue("Pessoa", "B" + strconv.Itoa(exin), value.Nome.String)
		fe.SetCellValue("Pessoa", "C" + strconv.Itoa(exin), value.Apelido.String)
		fe.SetCellValue("Pessoa", "D" + strconv.Itoa(exin), value.Cpf.String)
		fe.SetCellValue("Pessoa", "E" + strconv.Itoa(exin), value.Pai.String)
		fe.SetCellValue("Pessoa", "F" + strconv.Itoa(exin), value.Mae.String)
		fe.SetCellValue("Pessoa", "G" + strconv.Itoa(exin), value.Nascimento.String)
		fe.SetCellValue("Pessoa", "H" + strconv.Itoa(exin), value.Sexo.String) 	
		fe.SetCellValue("Pessoa", "I" + strconv.Itoa(exin), value.Status) 
		fe.SetCellValue("Pessoa", "J" + strconv.Itoa(exin), value.Date_joined.String() )
		fe.SetCellValue("Pessoa", "K" + strconv.Itoa(exin), strconv.FormatInt(value.Estabelecimento_id.Int64, 10) ) 
		fe.SetCellValue("Pessoa", "L" + strconv.Itoa(exin), strconv.FormatInt(value.Religiao_id.Int64, 10) ) 
		fe.SetCellValue("Pessoa", "M" + strconv.Itoa(exin), strconv.FormatInt(value.Tratamento_id.Int64, 10) ) 
		fe.SetCellValue("Pessoa", "N" + strconv.Itoa(exin), strconv.FormatInt(value.Usuario_id.Int64, 10) ) 
		fe.SetCellValue("Pessoa", "O" + strconv.Itoa(exin), strconv.FormatInt(value.Vinculo_id.Int64, 10) ) 
	
		//n2, err := f.WriteString( sq + "\n")
		//check(err)
		//time.Now()
		//fmt.Println(n2)

	}
	
	_, lerr = logger.WriteString( "Salvando excel.\n")
	if lerr != nil {
		check(lerr)
	}
	// Save xlsx file by the given path.
    if err := fe.SaveAs("./Pessoa.xlsx"); err != nil {
		check(lerr)
    }

	
	fmt.Printf("End %s \n",  time.Now().Format(fmtData) )
	_, lerr = logger.WriteString( "End " + time.Now().Format(fmtData) + "\n")
	if lerr != nil {
		check(lerr)
	}


	defer logger.Close()
	// Issue a `Sync` to flush writes to stable storage.
	logger.Sync()

	// check error here...
    //exec.Command("/bin/sh", "-c", "echo "+e.Error()+" >> ../log1.log").Run()
}


func Run() {
	
	ReadPessoa();
	// ReadMinicipios();
}