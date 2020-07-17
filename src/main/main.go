package main

import (
	"database/sql"
	"fmt"
	"flag"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

var host= flag.String("host","127.0.0.1","Mysql host")
var port= flag.String("port","3306","Mysql port")
var username= flag.String("username","root","Username")
var password= flag.String("password","root","Password")
var adbname= flag.String("adbname","adbname","A database nmae")
var bdbname =flag.String("bdbname","bdbname","B database name")
var diffonly = flag.String("diffonly","false","Print different only")
func main (){
	flag.Parse();
	var sqlConnectA = *username+":"+*password+"@tcp("+*host+":"+*port+")/"+*adbname+"?charset=utf8&parseTime=True";
	descA,err := getTableDesc(sqlConnectA);
	if err!=nil{

	}
	var sqlConnectB = *username+":"+*password+"@tcp("+*host+":"+*port+")/"+*bdbname+"?charset=utf8&parseTime=True";
	descB,err := getTableDesc(sqlConnectB);
	if err!=nil{
		
	}
	compareList(descA,descB);

}

// getTableDesc: Get table description by using "show tables" and "desc `tablename`"
func getTableDesc(sqlConnect string)(map[string]map[string]string ,error){
	descMap := map[string]map[string]string{};
	db,err := sql.Open("mysql",sqlConnect)
    if err != nil{
		fmt.Printf("connect mysql fail ! [%s]",err);
		return descMap,err;
	}
	tables,err:= db.Query("show tables")   
	for tables.Next() {       
		var tableName string
		
		err = tables.Scan(&tableName) 
		if err!=nil{
			
		}
		tableDes,err := db.Query("desc "+tableName);
		if err != nil {
			
		}
		descM := map[string]string{};
		cursor := 0
		for tableDes.Next() {
			var _field string
			var _type string
			var _null string
			var _key string
			var _default sql.NullString
			var _extra string
			err = tableDes.Scan(&_field,&_type,&_null,&_key,&_default,&_extra)
			
			if err != nil {
				fmt.Println(err)
			}
			t := strconv.Itoa(cursor)
			descM[t] = _field+"  "+_type+"  "+_null+"  "+_key;
			cursor++;
		}
		
		descMap[tableName] = descM;
	}

	return descMap,nil;
}

// compareList： Compare tow table descriptions to check the diffent
func compareList(descA map[string]map[string]string ,descB map[string]map[string]string){
	for key := range descA {
		t1 := descA[key];
		t2 := descB[key];
		fmt.Printf(key+":\n");
		for key1 := range t1 {
			if t1[key1]==t2[key1] {
				if(*diffonly=="false"){
					fmt.Printf("	\x1b[%dm"+t1[key1]+"--------------"+t2[key1]+"---\x1b[0m\n", 32)
				}
			}else {
				fmt.Printf("	\x1b[%dm"+t1[key1]+"--------------"+t2[key1]+"---\x1b[0m\n", 31)
			}
		}
	}
}
