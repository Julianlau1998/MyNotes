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
      swSrc: 'src/service-worker.js'
    }
  }
}