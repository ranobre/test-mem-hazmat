package gopacking

import (
	"encoding/json"
	"fmt"
	bs "github.com/inhies/go-bytesize"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	models "test-mem-hazmat/internal/model"
	"time"
)

type Gopacking struct {
}

func (g *Gopacking) Candidates(target, proposta int) {
	start := time.Now()
	p := message.NewPrinter(language.Portuguese)
	var jsonResponse [][]byte
	if proposta == 1 {
		fmt.Println("========================================")
		fmt.Println("Executando teste de Pack proposta1:")
		_, _ = p.Printf("==>INPUT: %v \n", target)
		var p1 []models.Proposta1
		var tamanho int64

		for i := 0; i < target; i++ {
			r := g.buildProposta1(i)
			p1 = append(p1, r)
			j, err := json.Marshal(r)
			jsonResponse = append(jsonResponse, j)
			tamanho += int64(len(j))
			if err != nil {
				log.Fatal(err)
			}

		}
		end := time.Since(start)
		fmt.Println(fmt.Sprintf("Tempo de execução: %v", end))
		formater := bs.ByteSize(tamanho)
		txt := formater.Format("JSONList TotalByteSize: %.2f ", "MB", true)
		fmt.Println(fmt.Sprintf("%v", txt))
		fmt.Println("========================================")
	} else if proposta == 2 {
		fmt.Println("========================================")
		fmt.Println("Executando teste de Pack proposta2:")
		_, _ = p.Printf("==>INPUT: %v \n", target)
		var p1 []models.Proposta2
		var tamanho int64
		for i := 0; i < target; i++ {
			r := g.buildProposta2(i)
			p1 = append(p1, r)
			j, err := json.Marshal(r)
			jsonResponse = append(jsonResponse, j)

			tamanho += int64(len(j))
			if err != nil {
				log.Fatal(err)
			}
		}
		end := time.Since(start)
		fmt.Println(fmt.Sprintf("Tempo de execução: %v ", end))

		formater := bs.ByteSize(tamanho)
		txt := formater.Format("JSONList TotalByteSize: %.2f ", "MB", true)

		fmt.Println(fmt.Sprintf("%v", txt))

		fmt.Println("========================================")
	}
}

func (g Gopacking) buildProposta1(idx int) models.Proposta1 {
	return models.Proposta1{
		Attributes: map[string][]interface{}{"hazmat": {buildB1(idx), buildA1(idx)}},
	}
}

func (g Gopacking) buildProposta2(idx int) models.Proposta2 {
	return models.Proposta2{
		Attributes: map[string][]interface{}{"hazmat": {buildB2(idx), buildA2(idx)}},
	}
}

func buildB1(qty int) models.HazmatAttribute {
	return models.HazmatAttribute{
		Name:     "battery",
		Quantity: qty,
		HazmatData: models.BatteryData{
			BatteryType: []string{"lithium_ion", "lithium_metal"},
			SaleFormat:  []string{"contained_in_equipment", "packed_with_equipment"},
			SaleNumber:  []string{fmt.Sprintf("%v", qty*1), fmt.Sprintf("%v", qty*2)},
		},
	}
}

func buildB2(qty int) models.HazmatAttribute2 {
	return models.HazmatAttribute2{
		Name:     "battery",
		Quantity: qty,
		Properties: []models.TypeValue{
			{Type: "battery_type", Values: []string{"lithium_ion", "lithium_metal"}},
			{Type: "sale_format", Values: []string{"contained_in_equipment", "packed_with_equipment"}},
			{Type: "sale_number", Values: []string{fmt.Sprintf("%v", qty*1), fmt.Sprintf("%v", qty*2)}},
		},
	}
}

func buildA1(qty int) models.HazmatAttribute {
	return models.HazmatAttribute{
		Name:     "aerosol",
		Quantity: qty,
		HazmatData: models.AerosolData{
			Volume:      []string{"100ml"},
			PackageType: []string{"glass"},
		},
	}
}

func buildA2(qty int) models.HazmatAttribute2 {
	return models.HazmatAttribute2{
		Name:     "battery",
		Quantity: qty,
		Properties: []models.TypeValue{
			{Type: "volume", Values: []string{"lithium_ion", "lithium_metal"}},
			{Type: "package_type", Values: []string{"contained_in_equipment", "packed_with_equipment"}},
		},
	}
}
