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
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet"
          integrity="sha384-1BmE4kWBq78iYhFldvKuhfTAU6auU8tT94WrHftjDbrCEXSU1oBoqyl2QvZ6jIW3" crossorigin="anonymous">
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js"
            integrity="sha384-ka7Sk0Gln4gmtz2MlQnikT1wXgYsOg+OMhuP+IlRH9sENBO0LRn5q+8nbTov4+1p"
            crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/chart.js@3.7.1/dist/chart.min.js"></script>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
    <script>
        function getFunctionValue() {

            let params = new Map()
            params["From"] = document.getElementById("x-from").value
            params["Step"] = document.getElementById("x-step").value
            params["End"] = document.getElementById("x-end").value
            //{{if .IsDerived}}
            //{{else}}
            //{{range $keyName, $keyDesc := .Keys}}
            params["params[{{$keyName}}]"] = document.getElementById("{{$keyName}}").value
            //{{end}}
            //{{end}}

            let myChart

            let url = '/function/{{.FunctionType}}/value?'
            for (let key in params) {
                url += key + "=" + params[key] + "&"
            }

            $.ajax({
                type: "get",
                url: url,
                success: function (response, status, xhr) {
                    let data=response
                    let dataSet = []
                    document.getElementById("executeExpression").innerText = data.expression
                    for (let x in data.values) {
                        let v = data.values[x]
                        dataSet.push({x: v.input.toString(), y: v.output})
                    }
                    let charItem = document.getElementById("dataLine").getContext('2d')
                    let chartStatus = Chart.getChart("dataLine");
                    if (chartStatus !== undefined) {
                        chartStatus.destroy();
                    }
                    if (myChart instanceof Chart) {
                        myChart.destroy();
                    }
                    myChart = new Chart(charItem, {
                        type: 'line',
                        data: {
                            datasets: [{
                                data: dataSet,
                                label: "Function data for {{.FunctionType}}",
                                borderWidth: 1,
                                tension: 0,
                                borderColor: 'rgb(75, 192, 192)'
                            }]
                        },
                        options: {
                            scales: {
                                y: {
                                    beginAtZero: true
                                }
                            },
                            parsing: {
                                xAxisKey: 'x',
                                yAxisKey: 'y',
                            }
                        }
                    });

                },
                error: function (jqXHR) {
                    document.getElementById("executeExpression").innerText = jqXHR.responseText
                }
            })
        }
    </script>
</head>
<body>

<div class="container">
    <div class="row align-items-start">
        <div class="col-1"></div>
        <div class="col-10">
            <div class="row align-items-start">
                <div class="col-4">
                    <h1>{{.FunctionType}}</h1>
                </div>
                <br>
                <hr>
                <br>
                <div class="row align-items-start">
                    <div class="col-3"><h3>Function name:</h3></div>
                    <div class="col-9"><h3>{{.FunctionName}}</h3></div>
                    <div class="col-3"></div>
                    <div class="col-9"><em>The function type is the type of function. But some function has same
                        type.</em>
                        <em>Specify the target function in metric-mocker with function's name instead of function's
                            type.</em></div>
                </div>
                <div class="col-12">

                </div>
            </div>

            <hr>
            <div class="row align-items-start">
                <div class="col-3">
                    <h3>Expression:</h3>
                </div>
                <div class="col-9">
                    <h3><code>{{.Expression}}</code></h3>
                </div>
            </div>

            <hr>
            <div class="row align-items-start">
                <div class="row align-items-start">
                    <div class="col-3"><h3>Keys:</h3></div>
                    <div class="col-8">
                        <em class="describe">Keys is the variable of function.</em>
                    </div>
                </div>
                <div class="row align-items-start">
                    <div class="col-12">
                        <div>
                            {{if .IsDerived}}
                            This function is derived function, from {{.FunctionName}}.
                            {{else if gt (len .Keys) 0 }}

                            <div style="margin-top: 1%; width: 100%">
                                <table class="table table-hover table-striped">
                                    <tr>
                                        <th scope="col">Key name</th>
                                        <th scope="col">Usage</th>
                                        <th scope="col">Type</th>
                                        <th scope="col">Default</th>
                                    </tr>
                                    {{range $keyName, $keyDesc := .Keys}}
                                    <tr style="height: 1%">
                                        <td scope="col"><code>{{$keyName}}</code></td>
                                        <td>{{$keyDesc.Mean}}</td>
                                        <td>{{$keyDesc.Type}}</td>
                                        <td>{{$keyDesc.Default}}</td>
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
                    </div>
                </div>
            </div>

            <hr>
            <div class="row align-items-start">
                <div class="col-3"><h3>Usage</h3></div>
                <div class="col-12">{{.Doc | unescaped}}</div>
            </div>
            <hr>
            <div class="row align-items-start">
                <div class="col-3"><h3>Mocker data:</h3></div>
                <div class="col-9"></div>
                <div class="col-12">
                    <div class="row align-items-start">
                        <div class="col-3">
                            <h5>X Range value</h5>
                        </div>
                        <div class="col-9"></div>

                        <div class="col-3">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="x-from-label">X-From</span>
                                <input id="x-from" type="text" class="form-control" aria-label="Sizing example input"
                                       aria-describedby="inputGroup-sizing-default" value="-10">
                            </div>
                        </div>

                        <div class="col-3">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="x-step-label">X-Step</span>
                                <input id="x-step" type="text" class="form-control" aria-label="Sizing example input"
                                       aria-describedby="inputGroup-sizing-default" value="1">
                            </div>
                        </div>

                        <div class="col-3">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="x-endlabel">X-End</span>
                                <input id="x-end" type="text" class="form-control" aria-label="Sizing example input"
                                       aria-describedby="inputGroup-sizing-default" value="10">
                            </div>
                        </div>
                    </div>
                </div>

                <div class="col-12">
                    <div class="row align-items-start">
                        <div class="col-3">
                            <h5>Function params:</h5>
                        </div>
                        <div class="col-9"></div>
                        {{if .IsDerived}}
                        {{else}}
                        {{range $keyName, $keyDesc := .Keys}}
                        <div class="col-3">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="{{$keyName}}-label">{{$keyName}}</span>
                                <input id="{{$keyName}}" type="text" class="form-control"
                                       aria-label="Sizing example input"
                                       aria-describedby="inputGroup-sizing-default" value="{{$keyDesc.Default}}">
                            </div>
                        </div>
                        {{end}}
                        {{end}}
                        <div class="col-12"></div>
                    </div>
                </div>

                <div class="col-9"></div>
                <div class="col-3">
                    <div class="d-grid gap-2">
                        <button type="button" class="btn btn-outline-primary" onclick="getFunctionValue()">Test</button>
                    </div>
                </div>
                <div class="col-12">
                    <div class="3"><h5> Execute expression:</h5></div>
                    <div class="9"><code id="executeExpression"></code></div>
                </div>
                <div class="col-12">
                    <canvas id="dataLine"></canvas>
                </div>
            </div>
        </div>
        <div class="col-1"></div>
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
