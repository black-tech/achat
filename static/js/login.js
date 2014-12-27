$(document).ready(function(){
	$("#b_login").click(function(){
		console.log($("#t_password").val());


		htmlobj=$.ajax({url:"/login",
			type:"POST",
			data:{code:'login',
				username:$("#t_username").val(),
				password:$("#t_password").val()},
			success:function(msg){ 
				var dataObj=eval("("+ msg +")");
				console.log("MSG:"+msg);
				console.log("OVER:"+dataObj.code);
				if (dataObj.code != "OK") {
					$("#login_status").html("<p>用户名或密码有误</p>");
					console.log("Login Error");
				}else{
					$("#login_status").html("<p>登陆成功</p>");
					location.href = "/chat_home"
				};
			}});

		console.log("Cookie:"+$.cookie('username'));
	});
});
