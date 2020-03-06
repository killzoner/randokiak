const node = require('rollup-plugin-node-resolve');
const commonjs = require('rollup-plugin-commonjs');

module.exports = {
  plugins: [
    node({
      mainFields: ['browser', 'es2015', 'module', 'jsnext:main', 'main'],
    }),
    commonjs({
      namedExports: {
        /**
         * Explicit export of hidden parts from protobufjs/minimal
         * shitty exports
         */
        'protobufjs/minimal': [ 
          'Writer', 
          'BufferWriter', 
          'Reader', 
          'BufferReader',
          'util',
          'rpc',
          'roots'
        ]
      }
    })
  ],
};
