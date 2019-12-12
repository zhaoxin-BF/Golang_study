<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    {{if .IsDisplay}}
        <em>{{.Content}}</em>
    {{else}}
        <em>{{.Content2}}</em>
    {{end}}


    {{range .Users}}
        {{.Uname}}{{.Upassword}}</br>
        {{$.len}}//是访问上一级目录
    {{end}}
</body>
</html>