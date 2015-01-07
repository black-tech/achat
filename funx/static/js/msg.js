function Message (code) {
	this.Code = '';
	this.Data = '';
	this.Desc = '';
	this.Supp = new Array();
}
function Message (code) {
	this.Code = code;
	this.Data = '';
	this.Desc = '';
	this.Supp = new Array();
}
function Message (code,data) {
	this.Code = code;
	this.Data = data;
	this.Desc = '';
	this.Supp = new Array();
}
function Message (code,data,desc) {
	this.Code = code;
	this.Data = data;
	this.Desc = desc;
	this.Supp = new Array();
}
Message.prototype.decodeJSON = function(JsonStr) {
            var jsonObj=eval("("+ JsonStr +")");
            if(null != jsonObj.Code)  {
                  this.Code=decode64(jsonObj.Code);
            }
            if(null != jsonObj.Data){
                  this.Data=decode64(jsonObj.Data);
            }
            if (null != jsonObj.Desc){
                  this.Desc=decode64(jsonObj.Desc);
            }
            this.Supp = new Array();
            for(var i=0;i<jsonObj.Supp.length;i++){
                  this.Supp[i]=decode64(jsonObj.Supp[i]);
            }
}
Message.prototype.encodeJSON = function(){
	 if(null != this.Code)  {
                  this.Code=encode64(this.Code);
            }
            if(null != this.Data){
                  this.Data=encode64(this.Data);
            }
            if (null != this.Desc){
                  this.Desc=encode64(this.Desc);
            }
            for(var i=0;i<this.Supp.length;i++){
                  this.Supp[i]=encode64(this.Supp[i]);
            }
}
Message.prototype.toString = function() {
	this.encodeJSON();
	var s ='{"Code":"';
	if(null != this.Code)  {
		s+=this.Code+'"';
            }
            if(null != this.Data){
            	s+=',"Data":"'+this.Data+'"';
            }
            if (null != this.Desc){
            	s+=',"Desc":"'+this.Desc+'"';
            }
            if (this.Supp.length>0) {
            	s+=',"Supp":['
	           for(var i=0;i<this.Supp.length;i++){
	            	if(0==i){
	            		s+='"'+this.Supp[i]+'"';
	            	}else{
	            		s+=',"'+this.Supp[i]+'"';
	            	}
	            }
	            s+=']';
	}
	s+='}';
	return s;
}

var testStr = '{"Code":"Y29kZXN0cmluZw==","Data":"aGVyb2hlcm9oZXJv","Desc":"ZGVzY2Rlc2NkZXNj","Supp":["dGVzdDExMQ==","dGVzdDIyMjI=","dGVzdDMzMzM="]}';

test_jsonToStr=function() {
	var v= new Message();
      	console.log(testStr);
	v.decodeJSON(testStr);
	// v.Code = "codestringjs";
	// v.Data = "datajs";
	// var arr = new Array("123123","135135");
	// v.Supp = arr;
	// console.log(v.Code);
	// console.log(v.Data);
	// console.log(v.Desc);
	// for(var i=0;i<v.Supp.length;i++)
	// {
	// 	console.log(v.Supp[i]);
	// }
	console.log(v.toString());
}

var keyStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";
function encode64(input) {
    input = utf16to8(input);
    var output = "";
    var chr1, chr2, chr3 = "";
    var enc1, enc2, enc3, enc4 = "";
    var i = 0;
    do {
        chr1 = input.charCodeAt(i++);
        chr2 = input.charCodeAt(i++);
        chr3 = input.charCodeAt(i++);
        enc1 = chr1 >> 2;
        enc2 = ((chr1 & 3) << 4) | (chr2 >> 4);
        enc3 = ((chr2 & 15) << 2) | (chr3 >> 6);
        enc4 = chr3 & 63;
        if (isNaN(chr2)) {
            enc3 = enc4 = 64;
        } else if (isNaN(chr3)) {
            enc4 = 64;
        }
        output = output + keyStr.charAt(enc1) + keyStr.charAt(enc2)
+ keyStr.charAt(enc3) + keyStr.charAt(enc4);
        chr1 = chr2 = chr3 = "";
        enc1 = enc2 = enc3 = enc4 = "";
    } while (i < input.length);
    return output;
}
//将Base64编码字符串转换成Ansi编码的字符串
function decode64(input) {
    var output = "";
    var chr1, chr2, chr3 = "";
    var enc1, enc2, enc3, enc4 = "";
    var i = 0;
    if (input.length % 4 != 0) {
        return "";
    }
    var base64test = /[^A-Za-z0-9\+\/\=]/g;
    if (base64test.exec(input)) {
        return "";
    }
    do {
        enc1 = keyStr.indexOf(input.charAt(i++));
        enc2 = keyStr.indexOf(input.charAt(i++));
        enc3 = keyStr.indexOf(input.charAt(i++));
        enc4 = keyStr.indexOf(input.charAt(i++));
        chr1 = (enc1 << 2) | (enc2 >> 4);
        chr2 = ((enc2 & 15) << 4) | (enc3 >> 2);
        chr3 = ((enc3 & 3) << 6) | enc4;
        output = output + String.fromCharCode(chr1);
        if (enc3 != 64) {
            output += String.fromCharCode(chr2);
        }
        if (enc4 != 64) {
            output += String.fromCharCode(chr3);
        }
        chr1 = chr2 = chr3 = "";
        enc1 = enc2 = enc3 = enc4 = "";
    } while (i < input.length);
    return utf8to16(output);
}
function utf16to8(str) {
    var out, i, len, c;
    out = "";
    len = str.length;
    for (i = 0; i < len; i++) {
        c = str.charCodeAt(i);
        if ((c >= 0x0001) && (c <= 0x007F)) {
            out += str.charAt(i);
        } else if (c > 0x07FF) {
            out += String.fromCharCode(0xE0 | ((c >> 12) & 0x0F));
            out += String.fromCharCode(0x80 | ((c >> 6) & 0x3F));
            out += String.fromCharCode(0x80 | ((c >> 0) & 0x3F));
        } else {
            out += String.fromCharCode(0xC0 | ((c >> 6) & 0x1F));
            out += String.fromCharCode(0x80 | ((c >> 0) & 0x3F));
        }
    }
    return out;
}
function utf8to16(str) {
    var out, i, len, c;
    var char2, char3;
    out = "";
    len = str.length;
    i = 0;
    while (i < len) {
        c = str.charCodeAt(i++);
        switch (c >> 4) {
            case 0: case 1: case 2: case 3: case 4: case 5: case 6: case 7:
                // 0xxxxxxx
                out += str.charAt(i - 1);
                break;
            case 12: case 13:
                // 110x xxxx 10xx xxxx
                char2 = str.charCodeAt(i++);
                out += String.fromCharCode(((c & 0x1F) << 6) | (char2 & 0x3F));
                break;
            case 14:
                // 1110 xxxx 10xx xxxx 10xx xxxx
                char2 = str.charCodeAt(i++);
                char3 = str.charCodeAt(i++);
                out += String.fromCharCode(((c & 0x0F) << 12) |
((char2 & 0x3F) << 6) |
((char3 & 0x3F) << 0));
                break;
        }
    }
    return out;
}
