var fs = require('fs');
var path = require('path');

var CODE = fs.readFileSync(path.join(__dirname, 'code.js'), 'utf8');
var CODE_TOKEN = fs.readFileSync(path.join(__dirname, 'code_token.js'), 'utf8');

var mod = function(tracers) {
	var jscode;
	if (tracers.length == 1 && Object.keys(tracers[0]).length == 1 && tracers[0].token) {
		jscode = CODE.replace(/BAIDU_TONGJI_TOKEN/g, "'" + tracers[0].token + "'");
	}
	else {
		jscode = CODE_TOKEN.replace('TRACERS_JSON', JSON.stringify(tracers)) + CODE;
	}
	return jscode;
};

module.exports = mod;
