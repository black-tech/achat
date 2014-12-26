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
				console.log("OVER:"+dataObj.password);
			}});

		console.log("Cookie:"+$.cookie('username'));
	});
});
