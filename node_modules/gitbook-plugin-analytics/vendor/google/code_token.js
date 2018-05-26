var GA_TOKEN;
TRACERS_JSON.forEach(function(T) {
	var L = window.location;
	var M = true
		&& (!T.host || T.host.indexOf(L.host) >= 0)
		&& (!T.hostname || T.hostname.indexOf(L.hostname) >= 0)
		&& (!T.protocol || T.protocol.indexOf(L.protocol) >= 0)
		&& (!T.port || T.port.indexOf(L.port) >= 0)
		;
	if (T.base && M) {
		M = false;
		for (var i = 0; i < T.base.length; i++) {
			M = M || L.href.startsWith(T.base[i]);
		}
	}
	if (T.path && M) {
		M = false;
		for (var i = 0; i < T.path.length; i++) {
			M = M || L.pathname.startsWith(T.path[i]);
		}
	}
	if (M) {
		GA_TOKEN = T.token;
	}
});
