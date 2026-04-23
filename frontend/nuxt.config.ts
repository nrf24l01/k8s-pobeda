export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  experimental: {
    appManifest: false
  },
  css: ['~/assets/css/main.css'],
  runtimeConfig: {
    public: {
      apiBaseUrl: process.env.NUXT_PUBLIC_API_BASE_URL || '/api'
    }
  },
  app: {
    head: {
      htmlAttrs: {
        lang: 'ru'
      },
      title: 'КУБЫПОБЕДА.РФ',
      meta: [
        {
          name: 'description',
          content: 'Очередной год победы Kubernetes.'
        }
      ],
      link: [
        { rel: 'icon', type: 'image/svg+xml', href: '/kubernetes-logo.svg' },
        { rel: 'icon', type: 'image/png', sizes: '32x32', href: '/favicon-32x32.png' },
        { rel: 'icon', type: 'image/png', sizes: '16x16', href: '/favicon-16x16.png' },
        { rel: 'icon', sizes: 'any', href: '/favicon.ico' },
        { rel: 'apple-touch-icon', sizes: '180x180', href: '/apple-touch-icon.png' }
      ]
    }
  }
})
