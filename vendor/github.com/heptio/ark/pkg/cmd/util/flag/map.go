/*
Copyright 2017 Heptio Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package flag

import (
	"fmt"
	"strings"
)

// Map is a Cobra-compatible wrapper for defining a flag containing
// map data (i.e. a collection of key-value pairs).
type Map struct {
	data              map[string]string
	entryDelimiter    string
	keyValueDelimiter string
}

// NewMap returns a Map using the default delimiters ("=" between keys and
// values, and "," between map entries, e.g. k1=v1,k2=v2)
func NewMap() Map {
	m := Map{
		data: make(map[string]string),
	}

	return m.WithEntryDelimiter(",").WithKeyValueDelimiter("=")
}

// WithEntryDelimiter sets the delimiter to be used between map
// entries.
//
// For example, in "k1=v1&k2=v2", the entry delimiter is "&"
func (m Map) WithEntryDelimiter(delimiter string) Map {
	m.entryDelimiter = delimiter
	return m
}

// WithKeyValueDelimiter sets the delimiter to be used between
// keys and values.
//
// For example, in "k1=v1&k2=v2", the key-value delimiter is "="
func (m Map) WithKeyValueDelimiter(delimiter string) Map {
	m.keyValueDelimiter = delimiter
	return m
}

// String returns a string representation of the Map flag.
func (m *Map) String() string {
	var a []string
	for k, v := range m.data {
		a = append(a, fmt.Sprintf("%s%s%s", k, m.keyValueDelimiter, v))
	}
	return strings.Join(a, m.entryDelimiter)
}

// Set parses the provided string according to the delimiters and
// assigns the result to the Map receiver. It returns an error if
// the string is not parseable.
func (m *Map) Set(s string) error {
	for _, part := range strings.Split(s, m.entryDelimiter) {
		kvs := strings.SplitN(part, m.keyValueDelimiter, 2)
		if len(kvs) != 2 {
			return fmt.Errorf("error parsing %q", part)
		}
		m.data[kvs[0]] = kvs[1]
	}
	return nil
}

// Type returns a string representation of the
// Map type.
func (m *Map) Type() string {
	return "mapStringString"
}

// Data returns the underlying golang map storing
// the flag data.
func (m *Map) Data() map[string]string {
	return m.data
}
