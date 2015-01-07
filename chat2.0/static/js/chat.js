var addr = "ws://127.0.0.1:8080/chat_websocket";
// var addr = "ws://fx.ckeyer.com/chat_websocket";
// var ws = new WebSocket(addr);// ws.onopen = function(e){
//       console.log("WS:onopen");
//       console.dir(e);
//       ws.send('{"Code":"cookie","Data":"'+$.cookie('achat')+'"}\n');
//       ws.onmessage = function(e){
//             console.log("onmessage");
//             console.dir(e);
//             var dataObj=eval("("+ e.data +")");
//             console.log("OVER:"+dataObj);
//             controlMsg(dataObj)
//             // $('#log').append('<p>'+dataObj.Data+'<p>');
//             // $('#log').get(0).scrollTop = $('#log').get(0).scrollHeight;
//       };
//       ws.onclose = function(e){
//             console.log("onclose");
//             console.dir(e);
//       };
//       ws.onerror = function(e){
//             console.log("onerror");
//             console.dir(e);
//       };
// };

$(function(){
      $('#msgform').click(function(){
            test_jsonToStr();
            // var s = new Message("code");
            // console.log(s.Code);
            // console.log(s.Supp.length);
            // ws.send('{"Code":"msg","Data":"'+$('#msg').val()+'"}\n');
            // $('#log').append('<p style="color:red;">My > '+$('#msg').val()+'<p>');
            // $('#log').get(0).scrollTop = $('#log').get(0).scrollHeight;
            // $('#msg').val('');
            return false;
      });
});

function controlMsg (dataObj) {
      if (dataObj.Code == "msg") {
            $('#log').append('<p>'+dataObj.Data+'<p>');
      } else if (dataObj.Code == "online_user_count"){
            $('#online_count').html(dataObj.Data);
      }
      else if (dataObj.Code == "waitting_user_count"){
            $('#waitting_count').html(dataObj.Data);
      } 
      else if (dataObj.Code == "online_user_list"){
            $('#online_list').html(dataObj.Data);
      }
      else if (dataObj.Code == "waitting_user_list"){
            $('#waitting_list').html(dataObj.Data);
      }
      else if (dataObj.Code == "wait"){
            $('#waitting_count').html(dataObj.Data);
            $('#msgform').attr('disabled','disabled');
            $('#msgform').html('排队等候中');
      }
      else if (dataObj.Code == "OK"){
            $('#msgform').removeAttr('disabled');
            $('#msgform').html('发送');
      }
      else if (dataObj.Code == ""){};
}
