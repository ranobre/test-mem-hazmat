package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
)

func GetJSONValue(proposta string) []byte {

	content, err := ioutil.ReadFile(fmt.Sprintf("/Users/ranobre/MeLi/PyMe/ProjetosTeste/test-mem-hazmat/internal/json/%s.json", proposta))
	if err != nil {
		log.Fatal(err)
	}
	return content

}
