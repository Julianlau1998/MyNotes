module.exports = {
  pwa: {
    name: 'MyNotes',
    themeColor: '#0f1820',
    msTileColor: '#0f1820',
    appleMobileWebAppCapable: 'yes',
    appleMobileWebAppStatusBarStyle: '0f1820',

    // configure the workbox plugin
    workboxPluginMode: 'InjectManifest',
     workboxOptions: {
      swSrc: './public/service-worker.js',

      manifestOptions: {
        start_url: '.',
        name: 'MyNotes',
        themeColor: '#0f1820',
        display: 'standalone',
      }
    }
  }
}