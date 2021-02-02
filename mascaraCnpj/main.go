package main

import (
	"fmt"
	"mascaraCnpj/mascara"
)

func main(){
	menu();
}

func menu(){
	var cnpj string
	//~ fmt.Println("informe um cnpj para mascarar");
	//~ fmt.Scan(&cnpj);	
	//~ fmt.Println(mascara.MaskCnpj(cnpj));
	fmt.Println("informe um cnpj para desmacarar");
	fmt.Scan(&cnpj);
	fmt.Println(mascara.UnmaskCnpj(cnpj));
	menu();
}
