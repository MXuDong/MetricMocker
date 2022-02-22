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
        .baseLayout {
            width: 80%;
            margin-left: 10%;
            margin-top: 2%;
            font-size: 150%;
        }

        .keyName {

        }

        /*The describes.*/
        .describe {
            color: gray;
            margin-left: 3%;
            font-size: 75%;
            width: 75%;
        }

        .tableBellowLines{
            border-bottom:1px dashed gray;
        }
        /*Table cell style.*/
        .tableCell {
            padding-left: 10%;
            border-left: 0;
            border-right: 0;
            width: 20%;
            font-size: 70%;
        }
        /*Doc of function style.*/
        .functionDoc {
            margin-top: 1%;
            margin-left: 2%;
            font-size: 75%;
        }

    </style>
</head>
<body>
<a href="/function" >Functions</a>
<div class="baseLayout">
    <h2> {{.FunctionName}} </h2>

    <p>Function type: <strong>{{.FunctionType}}</strong></p>

    <em class="describe">The function type is the type of function. But some function has same type.</em>
    <br>
    <em class="describe">Specify the target function in metric-mocker with function's name instead of function's type.</em>

    <hr>
    Expression:
    <br>
    <div style="margin-left: 2%;color: midnightblue">
        <code style="font-size: 150%">
            {{.Expression}}
        </code>
    </div>
    <em class="describe">The derived function has no param, so the derived function can't specify any param.</em>
    <hr>
    <div style="width: 100%">
        <h4>Keys:</h4>
        <em class="describe">Keys is the variable of function.</em><br>

        {{if .IsDerived}}
        This function is derived function, from {{.FunctionName}}.
        {{else if gt (len .Keys) 0 }}

        <div style="margin-top: 1%; width: 100%">
            <table>
                <tr>
                    <th class="tableCell" style="font-size: 80%;">Key name</th>
                    <th class="tableCell" style="font-size: 80%;">Usage</th>
                    <th class="tableCell" style="font-size: 80%;">Type</th>
                    <th class="tableCell" style="font-size: 80%;">Default</th>
                </tr>
                {{range $keyName, $keyDesc := .Keys}}
                <tr style="height: 1%">
                    <td class="tableBellowLines"><code>{{$keyName}}</code></td>
                    <td class="tableCell tableBellowLines">{{$keyDesc.Mean}}</td>
                    <td class="tableCell tableBellowLines">{{$keyDesc.Type}}</td>
                    <td class="tableCell tableBellowLines">{{$keyDesc.Default}}</td>
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
    <h4>Function usage:</h4>
    <div class="functionDoc">
        {{.Doc | unescaped}}
    </div>
</div>
</body>
</html>
`

func UnEscapedHtml(value string) interface{} {
	return template.HTML(value)
}

func GetHtml(funcInterface funcs.BaseFuncInterface) string {
	temp, err := template.New("func_template").
		Funcs(template.FuncMap{"unescaped": UnEscapedHtml}).
		Parse(templateVar)

	if err != nil {
		return err.Error()
	}

	byteBuffer := bytes.NewBuffer(nil)
	if err := temp.Execute(byteBuffer, GetFunctionDescribe(funcInterface)); err != nil {
		return err.Error()
	}

	return byteBuffer.String()
}
