package main

import (
	"fmt"
	"time"
	"strings"
	"strconv"
)

func main(){
	menu();
}

func menu(){	
	var data string
	fmt.Println("Insira uma data para verificar");
	fmt.Scan(&data);
	isData(data);
	menu();
}

func isData(data string){
	if(!strings.ContainsAny(data, "abcdefghijklmnopqrstuvwxyz")){
		if((len(data) == 8) || (len(data)  == 10)){
			if(strings.Contains(data, "-")){
				divido := strings.Split(data, "-");
				ano, _ := strconv.Atoi(divido[0])
				mes, _ := strconv.Atoi(divido[1])
				dia, _ := strconv.Atoi(divido[2])
				t := Date(ano, mes+1, 0);			
				switch {
					case len(divido[0]) == 4 && mes <=12 && dia <= t.Day():
						fmt.Println("data valida no formato banco");
					default:
						fmt.Println("data invalida");
				}
			}
			
			if(strings.Contains(data, "/")){
				divido := strings.Split(data, "/");
				ano, _ := strconv.Atoi(divido[2])
				mes, _ := strconv.Atoi(divido[1])
				dia, _ := strconv.Atoi(divido[0])
				t := Date(ano, mes+1, 0);			
				switch {
					case len(divido[2]) == 4 && mes <=12 && dia <= t.Day():
						fmt.Println("data valida no formato br");
					default:
						fmt.Println("data invalida");
				}
			}
		} else {
			fmt.Println("data invalida");
		}			
		
	} else {
		fmt.Println("Proibido caracteres")
	}
	
}

func Date(year, month, day int) time.Time {
        return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
