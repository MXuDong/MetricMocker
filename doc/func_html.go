package doc

import (
	"bytes"
	"html/template"
	"mmocker/pkg/funcs"
)

var templateVar = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.FunctionName}}</title>
    <style>
        .baseLayout{
            width: 80%;
            margin-left: 10%;
            margin-top: 1%;
            font-size: 150%;
        }
        .keyName{

        }
        .keyUsage{
            margin-left: 1%;
            padding-left: 10%;
            font-size: 70%;
        }
        .functionDoc{
			margin-top: 1%;
            margin-left: 2%;
			font-size: 75%;
        }
    </style>
</head>
<body>
<div class="baseLayout">
    <p>Function name:
        <strong>{{.FunctionName}}</strong>
    </p>
    <hr>
    Expression:
    <div style="display: inline-block;margin-left: 2%;color: midnightblue">
        <code style="font-size: 150%">
           {{.Expression}} 
        </code>
    </div>
<hr>
{{if gt (len .Keys) 0 }}
    <div style="width: 100%">
        Keys:<em style="color: gray; font-size: 50%;margin-left: 3%">Keys is the variable of function.</em>

        <div style="margin-top: 1%; width: 100%">
            <table>
                <tr>
                    <th style="font-size: 80%; width: 30%">Key name</th>
                    <th style="font-size: 80%;padding-left: 1%;width: 70%">Usage</th>
                </tr>
     			{{range $keyName, $keyDesc := .Keys}}
                <tr style="height: 1%">
                    <td><code>{{$keyName}}</code></td>
                    <td class="keyUsage">{{$keyDesc.Mean}}</td>
                </tr>
				{{end}}
            </table>
        </div>
    </div>
{{else}}
<div>
This function has no param.
</div>
{{end}}
    <hr>
    Function usage:
    <div class="functionDoc">
        {{.Doc}}
    </div>
</div>
</body>
</html>
`

func GetHtml(funcInterface funcs.BaseFuncInterface) string {
	temp, err := template.New("func_template").Parse(templateVar)
	if err != nil {
		return err.Error()
	}

	byteBuffer := bytes.NewBuffer(nil)
	if err := temp.Execute(byteBuffer, GetFunctionDescribe(funcInterface)); err != nil {
		return err.Error()
	}

	return byteBuffer.String()
}
