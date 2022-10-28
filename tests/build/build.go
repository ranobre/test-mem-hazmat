package build

import (
	"encoding/json"
	"fmt"
	"log"
	"test-mem-hazmat/internal/helpers"
	models "test-mem-hazmat/internal/model"
	"time"
)

func TestBuildProposta1(target int) []models.Proposta1 {
	start := time.Now()

	p1Result := make([]models.Proposta1, target)
	for i := 1; i <= target; i++ {
		c := helpers.GetJSONValue("proposta1")
		var p models.Proposta1
		err := json.Unmarshal(c, &p)
		if err != nil {
			log.Fatal(err)
		}

		p1Result = append(p1Result, p)
	}

	end := time.Since(start)

	fmt.Println("============")
	fmt.Printf("Tempo: %v s \n", end)
	fmt.Println("============")
	return p1Result
}

func TestBuildProposta2(target int) []models.Proposta2 {
	start := time.Now()

	p2Result := make([]models.Proposta2, target)
	for i := 1; i <= target; i++ {
		c := helpers.GetJSONValue("proposta2")
		var p models.Proposta2
		err := json.Unmarshal(c, &p)
		if err != nil {
			log.Fatal(err)
		}

		p2Result = append(p2Result, p)
	}

	end := time.Since(start)

	fmt.Println("============")
	fmt.Printf("Tempo: %v s \n", end)
	fmt.Println("============")
	return p2Result
}

//func TestValidateProposta2(attributes []models.Proposta2) {
//	for _, proposta2 := range attributes {
//		for k, v := range proposta2.Attributes {
//			if k == "hazmat" {
//				for _, ha := range v {
//					var at models.Attribute2
//					at, ok := ha.(models.HazmatAttribute2)
//				}
//			}
//		}
//
//	}
//}
