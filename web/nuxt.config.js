import colors from "vuetify/es5/util/colors";
const webpack = require("webpack");
export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: "IT.Reservation",
    htmlAttrs: {
      lang: "en",
    },
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      { hid: "description", name: "description", content: "" },
      { name: "format-detection", content: "telephone=no" },
    ],
    link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }],
    script: [
      {
        src: "https://cdn.lordicon.com/xdjxvujz.js",
      },
    ],
  },
  target: "server",
  generate: {
    fallback: "404.html",
  },
  // Global CSS: https://go.nuxtjs.dev/config-css
  css: ["@/assets/css/main.css"],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: ["@/plugins/vue-auth-image.js"],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    "@nuxt/postcss8",
    "nuxt-vite",
    "@nuxtjs/google-fonts",
    "@nuxtjs/vuetify",
  ],
  vite: {
    optimizeDeps: {
      include: ["cookie"],
    },
  },

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: ["@nuxtjs/axios", "@nuxtjs/apollo", "@nuxtjs/auth-next"],
  apollo: {
    clientConfigs: {
      default: {
        httpEndpoint: "http://localhost:3001/query",
      },
    },
    defaultOptions: {
      // See 'apollo' definition
      // For example: default query options
      $query: {
        loadingKey: "loading",
        fetchPolicy: "network-only",
      },
    },
  },
  sass: {
    indentedSyntax: true,
  },
  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
    // analyze: true,
    plugins: [
      new webpack.ProvidePlugin({
        // global modules
        $: "jquery",
        _: "lodash",
      }),
    ],
    // styleResources: {
    //   scss: "./assets/variables.scss",
    //   less: "./assets/*.less",
    //   // sass: ...,
    //   // scss: ...
    // },
  },
  googleFonts: {
    families: {
      // a simple name
      Kanit: [300],
    },
  },
  vuetify: {
    customVariables: ["~/assets/variables.scss"],
    theme: {
      dark: false,
      themes: {
        dark: {
          primary: colors.blue.darken2,
          accent: colors.grey.darken3,
          secondary: colors.amber.darken3,
          info: colors.teal.lighten1,
          warning: colors.amber.base,
          error: colors.deepOrange.accent4,
          success: colors.green.accent3,
        },
      },
    },
  },
  auth: {
    strategies: {
      google: {
        clientId:
          "860762400846-bs0ekdtkji2jd3ekj37atgpc0j6k3obe.apps.googleusercontent.com",
        codeChallengeMethod: "",
        responseType: "token id_token",
      },
    },
  },
};
