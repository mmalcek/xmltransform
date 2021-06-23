# xmltransform - XML to TEXT using GO templates
Transform XML to text file using GO "text/template"

## Relase duilds (windows, linux, mac)
https://github.com/mmalcek/xmltransform/releases

## Command line syntax

- Convert: **xmltransform.exe -i inputdata.xml -o output.csv -t template.tmpl**
- Check verion: **xmltransform.exe -v**

(If -o is not defined result is written to stdout)

## Template formating
- Basic iterate over lines (List all values for XML val1)
```
{{range .TOP_LEVEL.DATA_LINE}}{{.val1}}{{end}}
```
- Get XML tag (-VALUE = tag name)
```
{{index .Linked_Text "-VALUE"}}
```
- Use functions (count va1 + val2)
```
{{add .val1 .val2}} 
```
- If statement 
```
{{if gt (int $val1) (int $val2)}}Value1{{else}}Value2{{end}} is greater
```
Check template.tmpl and inputdata.xml for more advanced example

(more detailed info on https://golang.org/pkg/text/template/ )

## Lua custom functions
Aside of builtin functions you can write your own custom lua functions defined in ./lua/functions.lua file
- Input is always passed as json array of strings
- Output must be passed as string
- lua table array starts with 1
- Lua documentation http://www.lua.org/manual/5.1/

Minimal functions.lua example
```lua
json = require './lua/json'

function sum(incomingData) 
    dataTable = json.decode(incomingData)
    return tostring(tonumber(dataTable[1]) + tonumber(dataTable[2]))
end
```

### Call Lua function in template 
```
{{lua "sum" .val1 .val2}}
```
- "sum" - Lua function name

### Call builtin function
```
{{add .val1 .val2}}
```

## Available (builtin) functions
- add -> {{add .Value1 .Value2}}
- add1
- sub
- div
- mod
- mul
- randInt
- add1f - "...f" functions parse float but provide **decimal** operation using https://github.com/shopspring/decimal
- addf
- subf
- divf
- mulf
- round
- max
- min
- maxf
- minf
- dateFormat -> {{dateFormat .Value "oldFormat" "newFormat"}}
- now - {{now "02.01.2006"}} - GO format date (see notes below)
- b64enc
- b64dec
- b32enc
- b32dec
- uuid
- regexMatch
- upper
- lower
- trim
- trimAll
- trimSuffix
- trimPrefix
- atoi
- int64
- int
- float64

### dateFormat
dateFormat can parse date and time using GO time format
- https://programming.guide/go/format-parse-string-time-date-example.html