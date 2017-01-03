const path    = require('path');
const webpack = require('webpack');
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
  module: {
    loaders: [
      {
        test: /\.(css|scss)$/,
        loader: ExtractTextPlugin.extract('style-loader', 'css-loader!autoprefixer-loader!sass-loader!import-glob-loader')
      },
      {
        test: /\.(jpg|png|gif|svg)$/,
        loader: 'file-loader?name=[path][name].[ext]'
      }
    ],
    preloaders: [{test: /\.scss$/, loaders: 'import-glob-loader'}]
  },
  plugins: [
    new ExtractTextPlugin("[name].css"),
    new webpack.HotModuleReplacementPlugin(),
    new webpack.NoErrorsPlugin(),
    new webpack.optimize.CommonsChunkPlugin({
      name: 'application',
      chunk: 'application',
      minChunks: Infinity
    })
    // ,
    // new webpack.ProvidePlugin(
    //   {
    //     jQuery: 'jquery',
    //     $: 'jquery',
    //     'window.jQuery': 'jquery',
    //     'window.$': 'jquery',
    //     _: 'underscore',
    //     React: 'react',
    //     ReactDOM: 'react-dom',
    //     moment: 'moment'
    //   }
    // )
  ],
  resolve: {
    root: path.resolve(__dirname, 'assets', 'src'),
    extensions: ['', '.js', '.jsx', '.css', '.scss', '.sass'],
    alias: {
      // 'jquery': require.resolve('jquery'),
      // 'jquery-ui': 'jquery-ui/ui/widgets',
      // 'jquery-ui-css': 'jquery-ui/../../themes/base'
    }
  },
  devServer: {
    headers: { 'Access-Control-Allow-Origin': 'http://localhost:9000', 'Access-Control-Allow-Credentials': 'true' }
  }
}
