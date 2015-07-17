package main

import (
	"fmt"
	"io"
	"strings"
)

type BOMItem struct {
	Quantity     int
	Value        string
	Footprint    string
	References   []string
	CustomFields []string
}

type BillOfMaterials struct {
	Headers       []string
	Items         []BOMItem
	customHeaders []string
}

func NewBOM(e KicadExport) (*BillOfMaterials, error) {
	res := &BillOfMaterials{}
	err := res.generateHeaders(e)
	if err != nil {
		return nil, err
	}

	err = res.generateItemList(e)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *BillOfMaterials) appendField(fieldName string) {
	for _, f := range b.Headers {
		if f == fieldName {
			return
		}
	}
	b.customHeaders = append(b.customHeaders, fieldName)
	b.Headers = append(b.Headers, fieldName)
}

func (b *BillOfMaterials) generateHeaders(e KicadExport) error {
	b.Headers = []string{"Quantity", "References", "Value", "Footprint"}
	for _, comp := range e.Components.List {
		for _, f := range comp.Fields.List {
			b.appendField(f.Name)
		}
	}
	return nil
}

func (b *BillOfMaterials) makeUniqueIdentifier(comp Component) string {
	ident := fmt.Sprintf("Value=%s_Footprint=%s", comp.Value, comp.Footprint)
	//generate mapping of field
	mapping := make(map[string]int, len(comp.Fields.List))
	for i, f := range comp.Fields.List {
		mapping[f.Name] = i
	}

	for _, h := range b.customHeaders {
		fValue := "empty"
		if fIdx, ok := mapping[h]; ok == true {
			fValue = comp.Fields.List[fIdx].Value
		}
		ident = ident + "_" + h + "=" + fValue
	}

	return ident
}

func (b *BillOfMaterials) generateItemList(e KicadExport) error {
	mapping := make(map[string]int, len(e.Components.List))

	for _, comp := range e.Components.List {
		//create the unique identifier
		ident := b.makeUniqueIdentifier(comp)
		idx, ok := mapping[ident]
		if ok == false {
			//in that case we create a new entry
			mapping[ident] = len(b.Items)
			//	b.Headers = []string{"Quantity", "References", "Value", "Footprint"}
			newItem := BOMItem{
				Quantity:     1,
				References:   []string{comp.Reference},
				Value:        comp.Value,
				Footprint:    comp.Footprint,
				CustomFields: make([]string, 0, len(b.customHeaders)),
			}
			fieldMapping := make(map[string]int, len(comp.Fields.List))
			for i, f := range comp.Fields.List {
				fieldMapping[f.Name] = i
			}
			for _, h := range b.customHeaders {
				fValue := ""
				if fIdx, ok := fieldMapping[h]; ok == true {
					fValue = comp.Fields.List[fIdx].Value
				}
				newItem.CustomFields = append(newItem.CustomFields, fValue)
			}

			b.Items = append(b.Items, newItem)
			continue
		}

		b.Items[idx].Quantity = b.Items[idx].Quantity + 1
		b.Items[idx].References = append(b.Items[idx].References, comp.Reference)
	}

	return nil
}

func (b *BillOfMaterials) writeItem(w io.Writer, i BOMItem) error {
	res := make([]string, 0, len(i.CustomFields)+4)
	res = append(res, fmt.Sprintf("%d", i.Quantity))
	res = append(res, fmt.Sprintf(`"%s"`, strings.Join(i.References, ",")))
	res = append(res, i.Value)
	res = append(res, i.Footprint)
	res = append(res, i.CustomFields...)
	_, err := fmt.Fprintln(w, strings.Join(res, ","))

	return err
}

func (b *BillOfMaterials) WriteCSV(w io.Writer) error {

	_, err := fmt.Fprintln(w, strings.Join(b.Headers, ","))
	if err != nil {
		return err
	}

	for _, i := range b.Items {
		err = b.writeItem(w, i)
		if err != nil {
			return err
		}
	}
	return nil
}
