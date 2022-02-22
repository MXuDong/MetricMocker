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
    <script>
        function getFunctionValue() {
            let httpRequest = new XMLHttpRequest();//第一步：建立所需的对象
            httpRequest.open('GET', '/function/{{.FunctionName}}/value', true);//第二步：打开连接  将请求参数写在url中  ps:"./Ptest.php?name=test&nameone=testone"
            httpRequest.send();//第三步：发送请求  将请求参数写在URL中
            /**
             * 获取数据后的处理程序
             */
            httpRequest.onreadystatechange = function () {
                if (httpRequest.status === 200) {
                    let json = httpRequest.responseText;//获取到json字符串，还需解析
                    console.log(json);
                } else {
                    console.log(httpRequest.responseText)
                }
            };
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
                    <h1>{{.FunctionName}}</h1>
                </div>
                <br>
                <hr>
                <br>
                <div class="row align-items-start">
                    <div class="col-3">
                        <h3>Function type:</h3>
                    </div>
                    <div class="col-8">
                        <h3>{{.FunctionType}}</h3>
                    </div>
                    <div class="col-3"></div>
                    <div class="col-8"><em>The function type is the type of function. But some function has same
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
                <div class="col-8">
                    <h3><code>{{.Expression}}</code></h3>
                </div>
            </div>

            <hr>
            <div class="row align-items-start">
                <div class="row align-items-start">
                    <div class="col-3">
                        <h3>Keys:</h3>
                    </div>
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
                <div class="col-3">
                    <h3>Usage</h3>
                </div>
                <div class="col-12">{{.Doc | unescaped}}</div>
            </div>
            <hr>


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
