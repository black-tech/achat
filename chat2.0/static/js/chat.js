var addr = "ws://127.0.0.1:8080/chat_websocket";
// var addr = "ws://fx.ckeyer.com/chat_websocket";
var ws = new WebSocket(addr);
var connExi = false;
var msg ;
 ws.onopen = function(e){
      console.log("WS:onopen");
      connExi = true;
      console.dir(e);
      msg = new  Message('conn','achat client');
      ws.send(msg.toBase64String());

      ws.onmessage = function(e){
            console.log("onmessage");
            console.dir(e);
            var msg = new Message();
            msg.decodeJSON(e.data);
            console.log(msg.toString());
            controlMsg(msg)
      };
      ws.onclose = function(e){
            console.log("onclose");
            console.dir(e);
      };
      ws.onerror = function(e){
            console.log("onerror");
            console.dir(e);
      };
};
$(function(){
      $('#msgform').click(function(){
            msg  = new Message("msg");
            msg.Data = $('#msg').val();
            ws.send(msg.toBase64String());
            // ws.send('{"Code":"msg","Data":"'+$('#msg').val()+'"}\n');
            $('#log').append('<p style="color:red;">My > '+$('#msg').val()+'<p>');
            $('#log').get(0).scrollTop = $('#log').get(0).scrollHeight;
            $('#msg').val('');
            return false;
      });
});

function controlMsg (msg) {

      if (msg.Code == "msg") {
            $('#log').append('<p style="color:red;">Someone > '+msg.Data+'<p>');
            $('#log').get(0).scrollTop = $('#log').get(0).scrollHeight;
      } else if (msg.Code == "online_user_count"){
            $('#online_count').html(msg.Data);
      }
      else if (msg.Code == "waitting_user_count"){
            $('#waitting_count').html(msg.Data);
      } 
      else if (msg.Code == "online_user_list"){
            $('#online_list').html(msg.Data);
      }
      else if (msg.Code == "waitting_user_list"){
            $('#waitting_list').html(msg.Data);
      }
      else if (msg.Code == "wait"){
            $('#waitting_count').html(msg.Data);
            $('#msgform').attr('disabled','disabled');
            $('#msgform').html('排队等候中');
      }
      else if (msg.Code == "OK"){
            $('#msgform').removeAttr('disabled');
            $('#msgform').html('发送');
      }
      else if (msg.Code == ""){};
}
