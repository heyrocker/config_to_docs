package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"reflect"
	"strings"
)

var data = `
uuid: 3cc9b304-6eca-4b3c-b9d8-cb07789c91e2
langcode: en
status: true
dependencies:
  module:
    - menu_ui
    - type_tray
third_party_settings:
  type_tray:
    type_category: content
    type_thumbnail: ""
    type_icon: /themes/custom/wildrose/images/icons/typetray/icon-event.svg
    type_description: ""
    existing_nodes_link_text: ""
    type_weight: 0
name: Event
type: event
description: "Events open to the public, or official meetings they should be informed about."
help: ""
new_revision: true
preview_mode: 1
display_submitted: false
`

func main() {
	configData := make(map[string]interface{})

	err := yaml.Unmarshal([]byte(data), &configData)
	if err != nil {
		panic(err)
	}
	for key, value := range configData {
		fmt.Printf("Key:%s ", key)
		printVal(value, 1)
	}
}

func printVal(i interface{}, depth int) {
	typ := reflect.TypeOf(i).Kind()
	if typ == reflect.Int || typ == reflect.String || typ == reflect.Bool {
		fmt.Printf("%s%v\n", strings.Repeat(" ", depth), i)
	} else if typ == reflect.Slice {
		fmt.Printf("\n")
		printSlice(i.([]interface{}), depth+1)
	} else if typ == reflect.Map {
		fmt.Printf("\n")
		printMap(i.(map[interface{}]interface{}), depth+1)
	}
}

func printMap(m map[interface{}]interface{}, depth int) {
	for k, v := range m {
		fmt.Printf("%sKey:%s", strings.Repeat(" ", depth), k.(string))
		printVal(v, depth+1)
	}
}

func printSlice(slc []interface{}, depth int) {
	for _, v := range slc {
		printVal(v, depth+1)
	}
}
