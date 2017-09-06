const path = require('path');
const webpack = require('webpack');
const autoprefixer = require('autoprefixer');

module.exports = {
    entry: {
        app: './src/index.js'
    },

    output: {
        path: path.join(__dirname, '/build'),
        filename: '[name].js'
    },

    module: {
        rules: [
            {
                test: /\.jsx?$/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        //不使用.babelrc 文件
                        babelrc: false,

                        presets: [
                            ['es2015', { 'modules': false}],
                            'react'
                        ],

                        plugins: [
                            'react-require',
                            'transform-object-rest-spread'
                        ]
                    }
                },
                exclude: /node_modules/
            },

            {
                test: /\.css$/,
                use: [
                    { loader: 'style-loader' },
                    {
                        loader: 'css-loader'
                    },
                    {
                        loader: 'postcss-loader',
                        options: {
                            plugins: function() {
                                return [
                                    autoprefixer
                                ]
                            }
                        }
                    }
                ]
            },

            // 当图片文件大于 10kb 复制到指定目录 小于 10kb 转为 base64 编码
            {
                test: /\.(png|jpg|jpeg|gif)$/,
                use: [
                    {
                        loader: 'url-loader',
                        options: {
                            limit: 10000,
                            name: './images/[name].[ext]'
                        }
                    }
                ]
            }
        ]
    }
}

