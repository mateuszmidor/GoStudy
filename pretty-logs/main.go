package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// example stern line:
// user-auth-54f476d854-2k2bp primary-user-auth {"level":"debug","ts":1712240618.3682706,"caller":"internal/user.go:137","msg":"updating user's groups from IDP groups","user_id":101,"groups":[]}
func splitSternLine(input string) (pod, json string) {
	parts := strings.SplitN(input, "{", 2)
	if len(parts) == 2 {
		plain := strings.TrimSpace(parts[0])
		pod = strings.SplitN(plain, " ", 2)[0]
		json = "{" + strings.TrimSpace(parts[1])
	} else {
		pod = strings.TrimSpace(input)
		json = ""
	}
	return pod, json
}

func extract(key string, in map[string]interface{}) string {
	if value, ok := in[key]; ok {
		valueStr := fmt.Sprintf("%v", value)
		delete(in, key)
		return valueStr
	}
	return ""
}

func timestampToRFC3339(unixTS string) string {
	timestampFloat, err := strconv.ParseFloat(unixTS, 64)
	if err != nil {
		return "<no-date>"
	}

	timestamp := int64(timestampFloat)

	// Convert the Unix timestamp to time.Time
	t := time.Unix(timestamp, 0)

	// Format the time.Time value to RFC3339 string
	return t.Format(time.DateTime)
}

func red(s string) string {
	return "\033[31m" + s + "\033[0m"
}

func blue(s string) string {
	return "\033[1m\033[34m" + s + "\033[0m"
}
func green(s string) string {
	return "\033[32m" + s + "\033[0m"
}

func yellow(s string) string {
	return "\033[33m" + s + "\033[0m"
}

func cyan(s string) string {
	return "\033[36m" + s + "\033[0m"
}

func formatLevel(l string) string {
	switch strings.ToLower(l) {
	case "debug":
		return green("[DEBUG]")
	case "info":
		return green("[INFO] ")
	case "warn":
		return yellow("[WARN] ")
	case "error":
		return red("[ERROR]")
	case "fatal":
		return red("[FATAL]")
	case "audit":
		return yellow("[AUDIT]")
	default:
		return red(strings.ToUpper(l))
	}
}

func formatError(err string) string {
	end := strings.Index(err, "connectrpc.com/connect.NewUnaryHandler")
	if end > -1 {
		return err[0:end] + "<connect rpc middleware calls...>"
	}
	return err
}

func prettyPrint(plain string, items map[string]any) {
	extract("stacktrace", items) // added by loging lib, we dont want to see this
	verbose := extract("errorVerbose", items)
	out := ""

	if plain != "" {
		out = fmt.Sprintf("%s: ", plain)
	}

	level := extract("level", items)
	out += formatLevel(level)

	out += " "

	datetime := timestampToRFC3339(extract("ts", items))
	out += fmt.Sprintf("[%s]", datetime)

	out += " "

	caller := extract("caller", items)
	out += fmt.Sprintf("%s", caller)

	out += " "

	msg := extract("msg", items)
	out += yellow(msg)

	if len(items) > 0 {
		out += " | "

		pairs := []string{}
		for _, key := range sortedKeys(items) {
			value := items[key]
			valueStr := formatValue(value)
			pairs = append(pairs, fmt.Sprintf("%s:%s", blue(key), cyan(valueStr)))
		}
		if len(verbose) > 0 {
			pairs = append(pairs, cyan("errorVerbose:"+formatError(verbose)))
		}
		out += (strings.Join(pairs, ", "))
	}

	fmt.Println(out)
}

func formatValue(input any) string {
	result := ""
	switch val := input.(type) {
	case string:
		result = tryFormatList(val) // may happen to be weird-formatted list
	case []any:
		result = formatList(val) // format regular lists nicely
	default:
		result = fmt.Sprintf("%v", input)
	}
	if result == "" {
		result = "<empty>"
	}
	return result
}

// example input:  "names:\"AAD Oct 10 test\"  names:\"DTMX internal testers\""
// example output:  names:["AAD Oct 10 test","DTMX internal testers"]
func tryFormatList(input string) string {
	// check if this is actually this weird-formatted list where key repeats for every item making key:value pairs
	colonQuoteIndex := strings.Index(input, `:`)
	if colonQuoteIndex == -1 {
		return input
	}

	key := input[0 : colonQuoteIndex+1] // e.g. "names:"
	if strings.Count(input, key) < 2 {
		return input
	}

	// split items
	items := strings.Split(input, key)

	// remove prefix/suffix space and double quotes, and skip empty items
	sanitizedItems := []string{}
	for item := range items {
		sanitizedItem := strings.Trim(items[item], " \"")
		if sanitizedItem != "" {
			sanitizedItems = append(sanitizedItems, sanitizedItem)
		}
	}

	// format the list like: names:["AAD Oct 10 test","DTMX internal testers"]
	output := key + formatList(sanitizedItems)
	return output
}

func formatList[T any](input []T) string {
	formattedList, err := json.Marshal(input) // e.g. ["AAD Oct 10 test","DTMX internal testers"]
	if err != nil {
		fmt.Println("<pretty> - error formatting list: " + err.Error())
	}
	return string(formattedList)
}

const line = `{}
{"req":"names:\"AAD Oct 10 test\"  names:\"DTMX internal testers\"  names:\"AzureAD Integration Admins\"  names:\"DTMX test\"  names:\"AdamTestGroup\"", "color":"red"}
{"req":["andrzej","anna","arek"]}
{"req":[1,2,3]}
`

func main() {
	in := os.Stdin
	// in := strings.NewReader(line)
	reader := bufio.NewReader(in)

	for {
		// Read a line from stdin
		line, err := reader.ReadString('\n')
		if err != nil {
			// If an error occurs, break the loop
			fmt.Println("Pretty error; terminating: ", err)
			break
		}
		plain, jsonstring := splitSternLine(line)

		if jsonstring != "" {
			var jsonMap map[string]interface{}
			err = json.Unmarshal([]byte(jsonstring), &jsonMap)
			if err != nil {
				fmt.Println("###", line)
				continue
			}
			prettyPrint(plain, jsonMap)
		} else {
			fmt.Println(plain)
		}
	}
}

func sortedKeys(items map[string]interface{}) []string {
	var keys []string
	for k := range items {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
