var MODULE_REQUIRE
	/* built-in */

	/* NPM */
	
	/* in-package */
	;

function print(msg) {
	console.log('[ analytics ] ' + msg);
};

module.exports = {
	log: function(msg, args) {
		print(msg);
	},

	warn: function(msg) {
		print('WARN');
		print(msg);
	},

	error: function(msg) {
		print('EXCEPTION FOUND')
		print(msg);
		print('We are so sorry that gitbook ceased forcely.');
		process.exit(1);
	}
};
