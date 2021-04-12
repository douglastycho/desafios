package main

import (
	"fmt"
	"net/http"
)

func helloworld(p http.ResponseWriter, req *http.Request) {
	const tpl = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>Imersão Full Cycle</title>

			<!-- Latest compiled and minified CSS -->
			<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

			<!-- Optional theme -->
			<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">

			<!-- Latest compiled and minified JavaScript -->
			<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
			<style>
				body {
					background-image: url(https://i.ytimg.com/vi/kI2MLEroC_0/maxresdefault.jpg);
					background-repeat: no-repeat;
					background-attachment: fixed;
					background-size: cover;
				}
				.background-white {
					background: white;
				}
				.container-template {
					padding: 40px 15px;
    				text-align: center;
				}
			</style>
		</head>
		<body>
			<div class="container background-white">
				<div class="container-template">
					<h1>Imersão Full Cycle</h1>
				</div>
				<div class="row-fluid">
					<div class="col-md-12">
						<p>Informações do desafio</p>
						<p>Antes de iniciar o desafio, acesse o repositório do projeto da imersão no Github, de uma estrela e faça um fork.</p>
						<p>Desenvolva uma aplicação utilizando Golang que disponibilizará um servidor web na porta 8000.</p>
						<p>Ao acessar deveremos ver uma página com o seguinte conteúdo:</p>
						<p>- Background com um tema de um filme, quadrinhos, seriado a sua escolha</p>
						<p>- De forma centralizada escreva: Imersão Full Cycle entre h1</p>
						<p>- Fique na liberdade para estilizar a página a seu gosto</p>
						<p>- Gere o arquivo executável da aplicação rodado: GOOS=linux go build main.go</p>
						<p>Gere uma imagem Docker que ao ser executada rode a aplicação. Faça o push da imagem no Docker Hub.</p>
						<p>OBS: Lembrando que para realizar o push no Docker Hub, sua imagem deve ser o seguinte padrão de nome: <seulogin>/nome-da-imagem.</p>
						<p>Informe a imagem no campo abaixo</p>
					</div>
				</div>
			</div>
		</body>
	</html>`
	fmt.Fprintln(p, tpl)
}

func main() {
	http.HandleFunc("/", helloworld)
	http.ListenAndServe(":8000", nil)
}
