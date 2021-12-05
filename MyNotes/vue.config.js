module.exports = {
  pwa: {
    name: 'MyNotes',
    themeColor: '#003840',
    msTileColor: '#0f1820',
    mobileWebAppCapable: 'yes',
    appleMobileWebAppStatusBarStyle: 'default',
    start_url: '.',

    // configure the workbox plugin
    workboxPluginMode: 'InjectManifest',
     workboxOptions: {
      swSrc: 'src/service-worker.js',
      exclude: [/\.map$/, /_redirects/],
    }
  },
  devServer: {
    host: '0.0.0.0',
    disableHostCheck: true,
    compress: true,
    proxy: {
      '^/api/*': {
        target: 'http://192.168.178.58:8081/',
        //target: 'http://localhost:1324/',
        changeOrigin: true,
        ws: true,
      }
    }
  }
}