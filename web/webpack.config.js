module.exports = {
  mode: "development",
  entry: "./fixture-import.scss",
  module: {
    rules: [
      {
        test: /\.s[ac]ss$/i,
        use: [
          "css-loader",
          {
            loader: "sass-loader",
            options: {
              webpackImporter: false,
              sassOptions: {
                includePaths: ["node_modules"],
              },
              implementation: require("dart-sass"),
            },
          },
        ],
      },
    ],
  },
};
