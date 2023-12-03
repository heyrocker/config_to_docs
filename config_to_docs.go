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

var start_delimiter = "{{ "
var end_delimiter = " }}"

var template = `
title: {{ name }}
---
# {{ name }}
{{ description }}
<img src="{{ third_party_settings.type_tray.type_icon }}">
`

var data = `
uuid: 3cc9b304-6eca-4b3c-b9d8-cb07789c91e2
langcode: en
bool_value: true
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
preview_mode: 1
`

func main() {
	configData := make(map[string]interface{})

	err := yaml.Unmarshal([]byte(data), &configData)
	if err != nil {
		panic(err)
	}
	for key, value := range configData {
		replaceVal(value, 1, key)
	}

	fmt.Printf(template)
}

// Print a specific value for a key, depending on what type it is.
// If it is a standard type just print the value, if it is a map
// (indicating a new branch in the tree) then call a function to
// handle that case.
func replaceVal(i interface{}, depth int, original_key string) {
	val_type := reflect.TypeOf(i).Kind()
	if val_type == reflect.Int || val_type == reflect.String || val_type == reflect.Bool {
		// Replace any matching key in the template with the value here
		template = strings.Replace(template, start_delimiter+original_key+end_delimiter, fmt.Sprintf("%v", i), -1)
	} else if val_type == reflect.Slice {
		// Deal with slices then replace
		replaceSlice(i.([]interface{}), depth+1, original_key)
	} else if val_type == reflect.Map {
		// Deal with a new branch
		replaceMap(i.(map[interface{}]interface{}), depth+1, original_key)
	}
}

// Print an entire child branch, and generate unique keys for each value.
// If you encounter another child branch, this code will recurse via
// replaceVal().
func replaceMap(m map[interface{}]interface{}, depth int, original_key string) {
	var combined_key string

	for k, v := range m {
		combined_key = original_key + "." + k.(string)
		replaceVal(v, depth+1, combined_key)
	}
}

func replaceSlice(slc []interface{}, depth int, original_key string) {
	for _, v := range slc {
		replaceVal(v, depth+1, original_key)
	}
}
