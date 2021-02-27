module.exports = {
  pwa: {
    name: 'MyNotes',
    themeColor: '#003840',
    msTileColor: '#0f1820',
    mobileWebAppCapable: 'yes',
    appleMobileWebAppStatusBarStyle: '#0f1820',
    start_url: '.',

    // configure the workbox plugin
    workboxPluginMode: 'InjectManifest',
     workboxOptions: {
      swSrc: `/service-worker.js`,
    }
  }
}