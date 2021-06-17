# xmltransform - XML to TEXT using GO templates
Transform XML to text file using GO "text/template"

use: *xmltransform -i inputdata.xml -o output.csv -t template.tmpl*

(If -o is not defined result is written to stdout)

Check template.tmpl and inputdata.xml to see how to format data

(more detailed info on https://golang.org/pkg/text/template/ )

## Available functions
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