
var MakeDirWebpackPlugin = require('make-dir-webpack-plugin');

module.exports = {
  plugins: [
    new MakeDirWebpackPlugin({
      dirs: [
        { path: './dist/uploads' }
      ]
    })
  ]
}
