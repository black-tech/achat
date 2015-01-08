<!DOCTYPE html>

<html ng-app>
  	<head>
    		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">

    		<title>{{.PageTitle}}</title>
  		<link rel="shortcut icon" href="{{.StaticHost}}/static/img/logo.ico">
		<link rel="stylesheet" href="{{.StaticHost}}/static/css/global.css">
		<link rel="stylesheet" href="{{.StaticHost}}/static/css/chat.css">
	</head>

  	<body>
  		<h2>{{.TestStr}}</h2>
		<div>
			<p>最多允许同时在线 {{.max_online_count}} 人</p>
			<p>在线人数: <span id="online_count">0</span>人</p>
			<p>等候人数: <span id="waitting_count">0</span>人</p>
			<p>在线列表: <span id="online_list"></span></p>
			<p>等候列表: <span id="waitting_list"></span></p>
		</div>
		<div id="log" style="height: 300px;overflow-y: scroll;border: 1px solid #CCC;">
		</div>
  		<div>
			<form>
				<p>
				  Message: <input id="msg" type="text" value="Hello, world!">
				  <button id="msgform" type="button">发送</button>
				</p>
			</form>
  		</div>

		<script type="text/javascript" src="{{.StaticHost}}/static/js/jquery-1.11.1.min.js"></script>
		<script type="text/javascript" src="{{.StaticHost}}/static/js/jquery.cookie.js"></script>
		<script type="text/javascript" src="{{.StaticHost}}/static/js/global.js"></script>
		<script type="text/javascript" src="{{.StaticHost}}/static/js/msg.js"></script>
		<script type="text/javascript" src="{{.StaticHost}}/static/js/chat.js"></script>
	</body>
</html>
