var MODULE_REQUIRE
	/* built-in */
	, fs = require('fs')
	, path = require('path')
	, url = require('url')
	/* NPM */
	// , md = require('markdown-it')()
	/* in-package */
	, utils = require('./utils')
	;

// Code of supported VENDOR_CODES, the code MUST be lowercase.
var VENDOR_CODES = [ 'baidu', 'google' ];

// For fault tolerance, some common names are accepted as alias.
// By default, plural names will be accepted.
var OPTION_NAMES = [ 'base', 'protocol', 'host', 'hostname', 'port', [ 'pathname', 'pathnames', 'path', 'paths' ] ];
for (var i = 0; i < OPTION_NAMES.length; i++) {
	if (typeof OPTION_NAMES[i] == 'string') {
		OPTION_NAMES[i] = [ OPTION_NAMES[i], OPTION_NAMES[i] + 's' ];
	}
}

// Uniform vendor data.
function parseVendor(vendorCode, args) {
	var mod = require('./vendor/' + vendorCode);
	var options = {};
	if (mod.parse) {
		options = mod.parse(args);
	}
	else if (args instanceof Array && args.length == 1){
		options.token = args[0];
	}
	else if (args instanceof Object) {
		options = args;
	}
	else if (args) {
		options.token = args;
	}
	options.vendor = vendorCode;
	return options;
}

// Container of predefined vendors or vendor groups.
var _vendorGroups_ = {};

// Container to put context configs.
var _contexts_ = {};

var _book_;

function processContext(options) {
	// Just to simplify the related statements.
	var has = function(name) {
		return options.hasOwnProperty(name);
	};

	var each = function(name, action) {
		if (!options[name]) return;

		for (var i = 0; i < options[name].length; i++) {
			options[name][i] = action(options[name][i]);
		}
	}

	// Normalise vendor's code.
	var vendor = options.vendor.toLowerCase();
	delete options.vendor;

	// Ignore if vendor code is not supported.
	if (VENDOR_CODES.indexOf(vendor) < 0 && !_vendorGroups_[vendor]) {
		utils.warn('Vendor code "' + vendor + '" is not recognized.');
		return;
	}

	// ---------------------------
	// Uniform options' names and normalise the values.

	var total = 0;
	OPTION_NAMES.forEach(function(names) {
		var name = names[0];
		for (var i = 0, n = 0; i < names.length; i++) {
			if (!has(names[i])) continue;

			// The first name is the normal one we used.
			if (i > 1) {
				options[name] = options[names[i]];
				delete options[names[i]];
			}

			// For fault tolerance, we allow some common names and plural names used as option keynames.
			// However, the name and aliases should not coexist.
			// Because, in such cases, it is difficult to confirm that the locigal relation between the two options is 'AND' or 'OR'.
			if (++n > 1) {
				utils.error('Duplicated options "' + names.join('", "') + '" found. Please correct the book.json firstly.');
			}
		}

		// Use array form uniformly.
		// So, transform option value to an array if it is not.
		if (has(name) && !(options[name] instanceof Array)) {
			options[name] = [ options[name] ];
		}

		// Check option values.
		if (options[name]) {
			total++;
			var values = [];
			for (var i = 0; i < options[name].length; i++) {
				if (options[name][i]) {
					values.push(options[name][i]);
				}
				else {
					utils.warn('Empty value for option "' + name + '" is ignored.');
				}
			}
			if (values.length) {
				options[name] = values;
			}
			else {
				delete options[name];
			}

			if (total > 1 && options['base']) {
				utils.error('Option "base" conflicts with other options. Please correct the book.json firstly.');
			}
		}
	});

	// ---------------------------
	// Normalise options' values individually.

	each('protocol', function(protocol) {
		var re = /^(http|https)(:(\/\/)?)?$/;
		if (re.test(protocol)) {
			return RegExp.$1 + ':';
		}
		else {
			utils.error('Protocol "' + protocol + '" is invalid or not supported.');
		}
	});

	each('pathname', function(pathname) {
		return (pathname.charAt(0) != '/')  ? '/' + pathname : pathname;
	});

	// ---------------------------
	// Extract from base values.

	var optionsGroup;
	if (options['base']) {
		optionsGroup = [];
		var optionsBase;
		options['base'].forEach(function(base) {
			var opt = {};
			for (var keyname in options) {
				if (keyname != 'base') opt[keyname] = options[keyname];
			}

			// only protocol
			if (/^(http|https)(:(\/\/)?)?$/.test(base)) {
				opt.protocol = [ RegExp.$1 + ':' ];
			}

			// only hostname
			else if (/^[^:\/]+$/.test(base)) {
				opt.hostname = [ base ];
			}

			// only host
			else if (/^[^:\/]+:\d+$/.test(base)) {
				opt.host = [ base ];
			}

			// only port
			else if (/^\:?(\d+)$/.test(base)) {
				opt.port = [ RegExp.$1 ];
			}

			// url without protocol
			else if (/^\/\/([^\/:]+(:\d+)?)(.*)$/.test(base)) {
				opt.host = [ RegExp.$1 ];
				if (RegExp.$3) {
					opt.path = [ RegExp.$3 ];
				}
			}

			// only path
			else if (/^\//.test(base)) {
				opt.path = [ base ];
			}

			// real url
			else if (/^(http|https):\/\//.test(base)) {
				if (optionsBase) {
					optionsBase.base.push(base);
					return;
				}

				opt.base = [ base ];
			}

			// ...
			else {
				utils.error('Ambiguous "base" value: ' + base);
			}

			optionsGroup.push(opt);
		});
	}
	else {
		optionsGroup = [ options ];
	}

	// ---------------------------
	// Push to the container.

	var pushVendor = function(vendor, optionsGroup) {
		if (!_contexts_[vendor]) {
			_contexts_[vendor] = optionsGroup;
		}
		else {
			_contexts_[vendor] = _contexts_[vendor].concat(optionsGroup);
		}
	};

	if (_vendorGroups_[vendor]) {
		_vendorGroups_[vendor].forEach(function(vendorConfig) {
			// Copy vendor configs into analytics options.
			for (var i = 0; i < optionsGroup.length; i++) {
				for (var name in vendorConfig) {
					// The code of vendor is ignored.
					if (name != 'vendor') {
						optionsGroup[i][name] = vendorConfig[name];
					}
				}
			}
			pushVendor(vendorConfig.vendor, optionsGroup);
		});
	}
	else {
		pushVendor(vendor, optionsGroup);
	}
}

module.exports = {
    // Map of hooks
    hooks: {
		'init': function() {
			try {

			_book_ = this;

			// Reset the container.
			// This is useful under ``gitbook serve``.
			_contexts_ = {};
			_vendorGroups_ = {};

			// Obtain the plugin's related configurations.
			var PLUGIN_CONFIG = this.config.get('pluginsConfig.analytics');

			// Get predefined vendors.
			var vendors = PLUGIN_CONFIG['vendor'] || PLUGIN_CONFIG['vendors'];
			if (vendors) {
				for (var name in vendors) {
					var vendorList = vendors[name];
					if (vendorList.__proto__ == Object.prototype) {
						vendorList = [ vendorList ];
					}

					if (!(vendorList instanceof Array)) {
						utils.error('Predefined vendor (group) named expected: ' + name)
					}

					for (var i = 0; i < vendorList.length; i++) {
						if (vendorList[i] instanceof Array) {
							var vendorCode = vendorList[0];
							vendorList[i] = parseVendor(vendorCode, vendorCode)
						}
						else if (vendorList[i] instanceof Object) {
							vendorList[i] = parseVendor(vendorList[i].vendor, vendorList[i]);
						}
						else if (typeof vendorList[i] == 'string') {
							vendorList[i] = { vendor: vendorList[i] };
						}
						else {
							utils.error('Predefined vendor (group) named expected: ' + name);
						}
					}

					_vendorGroups_[name] = vendorList;
				}
			}

			for (var keyname in PLUGIN_CONFIG) {
				var value = PLUGIN_CONFIG[keyname];
				var options;

				if ([ 'vendor', 'vendors' ].indexOf(keyname) >= 0) {
					continue;
				}

				// If the keyname is a vendor code.
				else if (VENDOR_CODES.indexOf(keyname.toLowerCase()) >= 0) {
					var vendorCode = keyname.toLowerCase();
					options = parseVendor(vendorCode, value);
					processContext(options);
				}

				// Keyname "context" or "contexts" represents that the value is an array of context options.
				else if ([ 'context', 'contexts' ].indexOf(keyname) >= 0) {
					var contexts = (value instanceof Array) ? value : [ value ];
					contexts.forEach(function(context) {
						if (context.__proto__ != Object.prototype) {
							utils.error('Invalid context value ' + context);
						}
						processContext(context);
					});
				}

				// The others are regarded as "base" values.
				else {
					if (value instanceof Array) {
						var options = parseVendor(value[0], value.slice(1));
						options['base'] = keyname;
						processContext(options);
					}
					else if (value instanceof Object) {
						options = value;
						options.base = keyname;
						processContext(options);
					}
					else {
						utils.error('URL base expected: ' + keyname);
					}
				}
			}

			for (var vendor in _contexts_) {
				_contexts_[vendor].forEach(function(context) {
					utils.log('VENDOR: ' + vendor + ', CONTEXT: ' + JSON.stringify(context));
				});
			}

			var tracingJs = '';
			for (var vendor in _contexts_) {
				var render = require('./vendor/' + vendor);
				tracingJs += render(_contexts_[vendor]);
			}

			var jspath = path.join(this.output.root(), '_gitbook_plugin_analytics.js');
			fs.writeFileSync(jspath, tracingJs);

		} catch(ex) { console.log(ex.stack) }
		},

		'page': function(page) {
			var n = page.path.split(path.sep).length - 1;
			var src = '../'.repeat(n) + '_gitbook_plugin_analytics.js';
			page.content += '<script src="' + src + '"></' + 'script>';
			return page;
		}
	},

	// Map of new blocks
    blocks: {
	},

	// Map of new filters
	filters: {}
};
