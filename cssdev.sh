#!/bin/bash
node_modules/postcss-cli/bin/postcss --watch --use autoprefixer --output static/css/styles.build.css  static/css/styles.css
