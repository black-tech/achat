<!DOCTYPE html>

<html>
  	<head>
    	<title>Beego</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    	<style type="text/css" src="/static/css/home.css"></style>
	</head>

  	<body>
          <div>
            <p>最多允许同时在线 {{.max_online_count}} 人</p>
            <p>在线人数: <span id="online_count">{{.online_count}}</span></p>
            <p>等候人数: <span id="waitting_count">{{.waitting_count}}</span></p>
            <p>在线列表: <span id="online_list">{{.online_list}}</span></p>
            <p>等候列表: <span id="waitting_list">{{.waitting_list}}</span></p>
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
		<script type="text/javascript" src="/static/js/jquery-1.11.1.min.js"></script>
		<script type="text/javascript" src="/static/js/jquery.cookie.js"></script>
    <script type="text/javascript" src="/static/js/chat.js"></script>
		
		<script type="text/javascript">

          </script>  
	</body>
</html>
