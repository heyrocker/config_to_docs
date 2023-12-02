// This code is written to read a YAML file, iterate its keys, check
// for those keys in a template, then replace any matches that are found
// with the values. YAML does not force uniquness in keys outside of a
// specific branch in the tree. For instance the following is legal.
//
// key:
//   key2: hello
// another_key:
//   key2: world
//
// The template would distinguish between these two by referring to
// them as
//
// key.key2
// another_key.key2
//
// We generate these combined keys along the way as we iterate the yaml.

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
		printVal(value, 1, key)
	}
}

// Print a specific value for a key, depending on what type it is.
// If it is a standard type just print the value, if it is a map
// (indicating a new branch in the tree) then call a function to
// handle that case.
func printVal(i interface{}, depth int, original_key string) {
	val_type := reflect.TypeOf(i).Kind()
	if val_type == reflect.Int || val_type == reflect.String || val_type == reflect.Bool {
		fmt.Printf("%s%v\n", strings.Repeat(" ", depth), i)
	} else if val_type == reflect.Slice {
		fmt.Printf("\n")
		printSlice(i.([]interface{}), depth+1, original_key)
	} else if val_type == reflect.Map {
		fmt.Printf("\n")
		printMap(i.(map[interface{}]interface{}), depth+1, original_key)
	}
}

// Print an entire child branch, and generate unique keys for each value.
// If you encounter another child branch, this code will recurse via
// printVal().
func printMap(m map[interface{}]interface{}, depth int, original_key string) {
	var combined_key string

	for k, v := range m {
		combined_key = original_key + "." + k.(string)
		fmt.Printf("%sKey:%s", strings.Repeat(" ", depth), combined_key)
		printVal(v, depth+1, combined_key)
	}
}

func printSlice(slc []interface{}, depth int, original_key string) {
	for _, v := range slc {
		printVal(v, depth+1, original_key)
	}
}
