const miniCss = require('mini-css-extract-plugin');
module.exports = {
    // mode: "development",
    entry: './static/ts/index.ts',
    output: {
        filename: './static/ts/bundle.js',
        path: '/usr/app'
    },
    module: {
        rules: [
            {
                test: /\.css$/,
                use: [
                    miniCss.loader,
                    'css-loader',
                ],
            },
            {
                test:/\.(s*)css$/,
                use: [
                    miniCss.loader,
                    'css-loader',
                    'sass-loader',
                ]
            },
            {
                test: /\.tsx?$/,
                use: 'ts-loader',
                exclude: /node_modules/
            },
            {
                test: /\.html$/,
                use: 'raw-loader'
            }
        ]
    },
    plugins: [
        new miniCss({
            filename: './static/css/style.css',
        }),
    ],
    resolve: {
        extensions: ['', '.ts', '.tsx', '.js', '.es6', '.jsx']
    }
};