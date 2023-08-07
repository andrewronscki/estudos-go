package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
	Idade uint `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c)

	cachorroEmJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(cachorroEmJson))

	jsonCachorro := `{"nome":"Rex","raca":"Dalmata","idade":3}`

	var ca cachorro

	if erro := json.Unmarshal([]byte(jsonCachorro), &ca); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(ca)
}