var ws = new WebSocket("ws://127.0.0.1:8080/ws_chat");
ws.onopen = function(e){
    console.log("onopen");
    console.dir(e);
};
ws.onmessage = function(e){
    console.log("onmessage");
    console.dir(e);
    $('#log').append('<p>'+e.data+'<p>');
    $('#log').get(0).scrollTop = $('#log').get(0).scrollHeight;
};
ws.onclose = function(e){
    console.log("onclose");
    console.dir(e);
};
ws.onerror = function(e){
    console.log("onerror");
    console.dir(e);
};

$(function(){
    $('#msgform').submit(function(){
        ws.send($('#msg').val()+"\n");
        $('#log').append('<p style="color:red;">My > '+$('#msg').val()+'<p>');
        $('#log').get(0).scrollTop = $('#log').get(0).scrollHeight;
        $('#msg').val('');
        return false;
    });
});