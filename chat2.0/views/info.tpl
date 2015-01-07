<!DOCTYPE html>

<html ng-app>
  	<head>
    		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">

    		<title>{{.PageTitle}}</title>
  		<link rel="shortcut icon" href="{{.StaticHost}}/static/img/logo.ico">
		<link rel="stylesheet" href="{{.StaticHost}}/static/css/global.css">
		<link rel="stylesheet" href="{{.StaticHost}}/static/css/info.css">
	</head>

  	<body>
  		<header class="hero-unit">
			<div class="container">
				<div class="row">
					<div class="hero-text">
				  		<div class="hero-text">
							<h1>Welcome to {{.PageTitle}}!</h1>
							<p class="description">
							<br />
							<br />
							ViewCount {{.ViewCount}}
							</p>
				  		</div>
					</div>
					<div class=""></div>
				</div>
				<div>
					<p>共 {{.PageCount}} 个页面</p>
				</div>
				<div>
					{{range $key, $val := .PageMap}}
					<div>
						<p>{{$key}} : {{$val}}</p>
					</div>
					{{end}}
				</div>
			</div>
			

		</header>
		<script type="text/javascript" src="{{.StaticHost}}/static/js/global.js"></script>
		<script type="text/javascript" src="{{.StaticHost}}/static/js/info.js"></script>
		<script type="text/javascript" src="{{.StaticHost}}/static/js/jquery-1.11.1.min.js"></script>
	</body>
</html>
