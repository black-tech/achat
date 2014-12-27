<!DOCTYPE html>

<html>
  	<head>
    	<title>Beego</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    	<style type="text/css" src="/static/css/home.css"></style>
	</head>

  	<body>
          <div>
            <p>在线人数: <span id="online_count">{{.online_count}}</span></p>
            <p>最多允许同时在线 {{.max_online_count}} 人</p>
            <p>{{.testData}}</p>
          </div>
  		<div id="log" style="height: 300px;overflow-y: scroll;border: 1px solid #CCC;">
  		</div>
  		<div>
  		  <form>
  		      <p>
  		          Message: <input id="message" type="text" value="Hello, world!">
  		          <button onclick="send();" type="button">Send Message</button>

  		      </p>
  		  </form>
  		</div>
		<script type="text/javascript" src="/static/js/jquery-1.11.1.min.js"></script>
		<script type="text/javascript" src="/static/js/jquery.cookie.js"></script>
		
		<script type="text/javascript">
                  console.log("Cookie:"+$.cookie('username'));
          </script>  
	</body>
</html>
