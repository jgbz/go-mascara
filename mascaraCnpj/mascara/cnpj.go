package mascara

import (
	"strings"
	"regexp"
)

const (
	ponto = "."
	barra = "/"
	sub = "-"
)

var reg = regexp.MustCompile(`^[0-9]*$`)

func MaskCnpj(this string) string{
	slice := strings.Split(this, "");
	if (len(slice) != 14) || reg.MatchString(this) == false{
		return "Invalido"
	}
	masked := strings.Join(slice[:2], "") + ponto + strings.Join(slice[2:5], "") + ponto + strings.Join(slice[5:8], "") + barra + strings.Join(slice[8:12], "") + sub + strings.Join(slice[12:], "");
	return masked
}

func UnmaskCnpj(this string) string{	
	this = strings.Replace(this, ponto, "", -1);
	this = strings.Replace(this, barra, "", -1);
	this = strings.Replace(this, sub, "", -1);
	if (len(this) != 14) || reg.MatchString(this) == false{
		return "cnpj invalido"
	}
	return this
}

//~ d:= regexp.MustCompile()
//~ d.MatchString(args)

