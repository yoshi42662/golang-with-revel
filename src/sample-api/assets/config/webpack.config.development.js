/**
 * Require basic plugins
 */
const path    = require('path');
const webpack = require('webpack');

/**
 * Require webpack plugins
 */
const ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
  devtool: 'inline-source-map',
  cache: true,
  context: path.join(__dirname + '/../../assets/src'),
  entry: {
    application: [
      'webpack-dev-server/client?http://localhost:8080',
      'webpack/hot/dev-server',
      './js/application.js'
    ]
  },
  scripts: 'webpack-dev-server --hot --inline --port 8080 --host 0.0.0.0 --progress --profile --colors',
  output: {
    path: path.join(__dirname, '../../public'),
    filename: '[name].js',
    publicPath: 'http://localhost:8080/assets/'
  },
  plugins: [
    new ExtractTextPlugin("[name].css"),
    new webpack.HotModuleReplacementPlugin()
  ],
  devServer: {
    headers: { 'Access-Control-Allow-Origin': 'http://localhost:9000', 'Access-Control-Allow-Credentials': 'true' }
  }
}
