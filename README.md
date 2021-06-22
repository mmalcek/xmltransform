# xmltransform - XML to TEXT using GO templates
Transform XML to text file using GO "text/template"

use: *xmltransform -i inputdata.xml -o output.csv -t template.tmpl*

check verion: *xmltransform -v*

(If -o is not defined result is written to stdout)

Check template.tmpl and inputdata.xml to see how to format data

(more detailed info on https://golang.org/pkg/text/template/ )

## Lua custom functions
Aside of builtin functions you can use custom lua functions defined in ./lua/functions.lua file
- Input is always passed as json array of strings
- Output must be passed as string
- lua table array starts with 1
- Lua documentation http://www.lua.org/manual/5.1/

Minimal functions.lua example
```lua
json = require 'json'

function sum(incomingData) 
    dataTable = json.decode(incomingData)
    return tostring(tonumber(dataTable[1]) + tonumber(dataTable[2]))
end
```

### Call Lua function in template 
```
{{lua "mul" .val1 .val2}}
```
- "mul" - Lua function name

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
- add1f
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
- now
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