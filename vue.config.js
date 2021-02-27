module.exports = {
  // ...other vue-cli plugin options...
  pwa: {
    name: 'MyNotes',
    themeColor: '#0f1820',
    msTileColor: '#0f1820',
    appleMobileWebAppCapable: 'yes',
    appleMobileWebAppStatusBarStyle: '0f1820',

    // configure the workbox plugin
    workboxPluginMode: 'InjectManifest',
    workboxOptions: {
      // swSrc is required in InjectManifest mode.
      swSrc: 'dev/sw.js',
      // ...other Workbox options...
    },
    manifestOptions: {
      start_url: '.',
      name: 'MyNotes',
      themeColor: '#0f1820',
      display: 'standalone',
    }
  }
}