package externalteam

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"strconv"
	"strings"
	"test-mem-hazmat/internal/helpers"
	models "test-mem-hazmat/internal/model"
	"time"
)

type ExternalTeam struct{}

func (e *ExternalTeam) UnpackProposta(target, proposta int) {
	start := time.Now()
	p := message.NewPrinter(language.Portuguese)

	if proposta == 1 {
		fmt.Println("========================================")
		fmt.Println("Executando teste de Unpack proposta1:")
		_, _ = p.Printf("==>INPUT: %v \n", target)

		for i := 0; i <= target; i++ {
			requestMockInBytes := helpers.GetJSONValue("proposta1")
			err := buildP1(requestMockInBytes)

			if err != nil {
				log.Fatal(err)
			}

		}

		end := time.Since(start)
		fmt.Println(fmt.Sprintf("Tempo de execução: %v", end))
		fmt.Println("========================================")
	} else if proposta == 2 {
		fmt.Println("========================================")
		fmt.Println("Executando teste de Unpack proposta2:")
		_, _ = p.Printf("==>INPUT: %v \n", target)

		for i := 0; i <= target; i++ {
			requestMockInBytes := helpers.GetJSONValue("proposta2")
			err := buildP2(requestMockInBytes)

			if err != nil {
				log.Fatal(err)
			}

		}

		end := time.Since(start)
		fmt.Println(fmt.Sprintf("Tempo de execução: %v ", end))
		fmt.Println("========================================")
	}

}

func buildP1(requestMockInBytes []byte) error {
	var p1 models.Proposta1
	err := json.Unmarshal(requestMockInBytes, &p1)
	if err != nil {
		return err
	}

	for attribute, listofValues := range p1.Attributes {
		if attribute == "hazmat" {
			_ = buildHazmatP1(listofValues)
		}
	}
	return nil
}

func buildP2(requestMockInBytes []byte) error {
	var p2 models.Proposta2
	err := json.Unmarshal(requestMockInBytes, &p2)
	if err != nil {
		return err
	}

	for attribute, listofValues := range p2.Attributes {
		if attribute == "hazmat" {
			_ = buildHazmatP2(listofValues)
		}
	}
	return nil
}

func buildHazmatP1(attValue []interface{}) []models.HazmatAttribute {
	var result []models.HazmatAttribute
	for _, hazmat := range attValue {
		r, ok := hazmat.(map[string]interface{})
		if !ok {
			log.Fatal("FAIL 1")
		}
		name := fmt.Sprintf("%v", r["name"])
		qty, err := strconv.ParseInt(fmt.Sprintf("%v", r["quantity"]), 10, 32)
		if err != nil {
			log.Fatal("FAIL 2")
		}
		values := r["hazmat_data"]
		if !ok {
			log.Fatal("FAIL 3")
		}
		hazmatData := buildHazmatData(name, values)
		h := models.HazmatAttribute{
			Name:       name,
			Quantity:   int(qty),
			HazmatData: hazmatData,
		}
		result = append(result, h)
	}
	return result
}

func buildHazmatP2(attValue []interface{}) []models.HazmatAttribute2 {
	var result []models.HazmatAttribute2
	for _, hazmat := range attValue {
		r, ok := hazmat.(map[string]interface{})
		if !ok {
			log.Fatal("FAIL 11")
		}
		name := fmt.Sprintf("%v", r["name"])
		qty, err := strconv.ParseInt(fmt.Sprintf("%v", r["quantity"]), 10, 32)
		if err != nil {
			log.Fatal("FAIL 22")
		}
		values := r["properties"]
		if !ok {
			log.Fatal("FAIL 33")
		}
		hazmatData := buildHazmatData2(name, values)
		h := models.HazmatAttribute2{
			Name:       name,
			Quantity:   int(qty),
			Properties: hazmatData,
		}
		result = append(result, h)
	}
	return result
}

func buildHazmatData(name string, values interface{}) interface{} {
	n := strings.ToLower(name)
	if n == "battery" {
		var bt []string
		var sf []string
		var sn []string
		vv, _ := values.(map[string][]interface{})
		for k, v := range vv {
			switch k {
			case "battery_type":
				for _, vv := range v {
					r := fmt.Sprintf("%v", vv)
					bt = append(bt, r)
				}
			case "sale_format":
				for _, vv := range v {
					r := fmt.Sprintf("%v", vv)
					sf = append(sf, r)
				}
			case "sale_number":
				for _, vv := range v {
					r := fmt.Sprintf("%v", vv)
					sn = append(sn, r)
				}
			}

		}

		return models.BatteryData{
			BatteryType: bt,
			SaleFormat:  sf,
			SaleNumber:  sn,
		}
	}
	return nil
}

func buildHazmatData2(name string, values interface{}) []models.TypeValue {
	n := strings.ToLower(name)
	v, _ := values.([]interface{})
	var result []models.TypeValue
	if n == "battery" {
		for _, vv := range v {
			r, k := vv.(map[string]interface{})
			if !k {
				print("not ok")
			}
			var tv models.TypeValue
			tv.Type = fmt.Sprintf("%v", r["type"])

			vl, _ := r["values"].([]interface{})
			for _, vlv := range vl {
				tv.Values = append(tv.Values, fmt.Sprintf("%v", vlv))
			}

			result = append(result, tv)

		}

	}
	return result
}
