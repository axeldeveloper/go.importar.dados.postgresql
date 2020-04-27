package postgres

import (
	"fmt"
	"log"
	_ "github.com/lib/pq"
	models "github.com/axeldeveloper/go.importar.dados.postgresql/models"
	//"github.com/360EntSecGroup-Skylar/excelize"
    "time"
    //"reflect"
   // "strconv"

)

type ExcelData interface {
    CreateMap(arr []string) map[string]interface{}
    ChangeTime(source string) time.Time
}

type ExcelStrcut struct {
    temp  [][]string
    Model interface{}
    Info  []map[string]string
}

/*
func (excel *ExcelStrcut)ReadExcel(file string) *ExcelStrcut{

    xlsx, err := excelize.OpenFile(file)
    if err != nil {
        fmt.Println("read excel:",err)
    }

    rows := xlsx.GetRows("Page1")
    excel.temp = rows

    return excel

}

func (excel *ExcelStrcut)CreateMap() *ExcelStrcut{

    for _,v:=range excel.temp{

        var info = make(map[string]string)
        for i:=0;i<reflect.ValueOf(excel.Model).NumField();i++{

            obj:=reflect.TypeOf(excel.Model).Field(i)
            info[obj.Name] = v[i]

        }
        excel.Info=append(excel.Info, info)

    }
    return excel
}


func (excel *ExcelStrcut)ChangeTime(source string) time.Time{
    ChangeAfter,err:=time.Parse("2006-01-02", source)
    if err!=nil {
        log.Fatalf("pau:%s",err)
    }
    return ChangeAfter
}


func (excel *ExcelStrcut)SaveDb(temp *models.Pessoa) *ExcelStrcut{


    for i:=1 ;i<len(excel.Info);i++{

        t:=reflect.ValueOf(temp).Elem()
        for k,v:=range excel.Info[i]{

            //fmt.Println(t.FieldByName(k).t.FieldByName(k).Kind())
            //fmt.Println("key:%v---val:%v",t.FieldByName(k),t.FieldByName(k).Kind())

            switch t.FieldByName(k).Kind(){
            case reflect.String:
                t.FieldByName(k).Set(reflect.ValueOf(v))
            case reflect.Float64:
                tempV,err:= strconv.ParseFloat(v,64)
                if err != nil{
                    log.Printf("string to float64 err：%v",err)
                }

                t.FieldByName(k).Set(reflect.ValueOf(tempV))
            case reflect.Uint64:
                reflect.ValueOf(v)
                tempV, err := strconv.ParseUint(v, 0, 64)
                if err != nil{
                    log.Printf("string to uint64 err：%v",err)
                }
                t.FieldByName(k).Set(reflect.ValueOf(tempV))

            case reflect.Struct:
                tempV,err:=time.Parse("2006-01-02", v)
                if err!=nil {
                    log.Fatalf("string to time err:%v",err)
                }
                t.FieldByName(k).Set(reflect.ValueOf(tempV))
            default:
                fmt.Println("type err")

            }


        }
		
		
		err:=DB.Create(&temp).Error
        if err != nil{
            log.Fatalf("save temp table err:%v",err)
        }
		fmt.Printf("导入临时表成功")
		

    }
    return excel
}

*/

func ReadPesso(){	
	members, err := models.BuscaTodos()	
	if err != nil {
		log.Fatal("Read failed:", err.Error())
	}
	fmt.Printf("Read %d rows successfully.\n", len(members))

	
	//e:=ExcelStrcut{}
    //temp := models.Pessoa{}
    //e.Model=temp
    //e.ReadExcel("../test.xlsx").CreateMap().SaveDb(&temp)

	
	for _, value := range members {
		fmt.Println(value.Nome.String)
	}
	
	
}


func Rum() {
	
	ReadPesso();
	// ReadMinicipios();
}