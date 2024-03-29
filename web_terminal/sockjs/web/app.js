if (!window.location.origin) { // Some browsers (mainly IE) do not have this property, so we need to build it manually...
  window.location.origin = window.location.protocol + '//' + window.location.hostname + (window.location.port ? (':' + window.location.port) : '');
}


var sock = new SockJS(window.location.origin+'/echo')

sock.onopen = function() {
	// console.log('connection open');
	document.getElementById("status").innerHTML = "connected";
	document.getElementById("send").disabled=false;
};

sock.onmessage = function(e) {
	document.getElementById("output").value += e.data +"\n";
};

sock.onclose = function() {
	// console.log('connection closed');
	document.getElementById("status").innerHTML = "disconnected";
	document.getElementById("send").disabled=true;
};

var sock2 = new SockJS('https://mydomain.com/my_prefix');
sock2.onopen = function() {
	console.log('open');
	sock.send('test');
};

sock2.onmessage = function(e) {
	console.log('message', e.data);
	sock.close();
};

sock2.onclose = function() {
	console.log('close');
};
