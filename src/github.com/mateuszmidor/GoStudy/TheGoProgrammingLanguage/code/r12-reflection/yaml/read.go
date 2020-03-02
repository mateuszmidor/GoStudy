package yaml

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"text/scanner"
)

// Unmarshall reads yaml input into output
func Unmarshall(input []byte, output interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(input))

	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error in %s: %v", lex.scan.Position, x)
		}
	}()

	readYamlHeader(lex)
	read(lex, reflect.ValueOf(output).Elem(), 0)
	return nil
}

func readYamlHeader(lex *lexer) {
	lex.consumeRunes("---")
	lex.consumeEOL()
}

func read(lex *lexer, v reflect.Value, level int) {
	switch v.Kind() {

	// "Flag: true\n"
	case reflect.Bool:
		lex.consumeSpaces(level)
		_, strval := lex.readKeyValue()
		parseInto(strval, v)
		return // nothing more to do here

	// "Age: 32\n"
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8: // only unsigned ints
		lex.consumeSpaces(level)
		_, strval := lex.readKeyValue()
		parseInto(strval, v)
		return // nothing more to do here

	// "Name: "Shaike"\n"
	case reflect.String: // only quoted strings
		lex.consumeSpaces(level)
		_, strval := lex.readKeyValue()
		parseInto(strval, v)
		return // nothing more to do here

	// "Person:\n"
	// " Age: 32\n"
	// " Name: "Sheike"\n"
	// Note: assuming field order in YAML reflects field order in struct
	case reflect.Struct:
		lex.consumeSpaces(level)
		_ = lex.readObject()
		for i := 0; i < v.NumField(); i++ {
			read(lex, v.Field(i), level+1)
		}
		return // nothing more to do here

	// "FavouriteNumbers:\n"
	// "- 1\n"
	// "- 3\n"
	case reflect.Slice:
		_ = lex.readObject()
		for {
			// end of slice discovery
			if !lex.consumeSpaces(level) {
				break
			}
			var strval string
			strval, ok := lex.readListItem()
			if !ok {
				break
			}
			item := reflect.New(v.Type().Elem()).Elem()
			parseInto(strval, item)
			v.Set(reflect.Append(v, item))
		}
		return // nothing more to do here

	default:
		panic(fmt.Sprintf("Unexpected ouput type %v", v.Kind()))
	}
}

func parseInto(strval string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Bool:
		flag, err := strconv.ParseBool(strval)
		panicOnErr(err)
		v.SetBool(flag)

	case reflect.String:
		val, err := strconv.Unquote(strval)
		panicOnErr(err)
		v.SetString(val)

	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8: // only unsigned ints
		val, err := strconv.ParseUint(strval, 10, 64)
		panicOnErr(err)
		v.SetUint(val)

	default:
		panic(fmt.Sprintf("Unsupported type %v", v.Kind()))
	}
}

func panicOnErr(err error) {
	if err != nil {
		panic(fmt.Errorf("%v", err))
	}
}
