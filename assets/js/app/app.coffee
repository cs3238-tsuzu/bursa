# Hack to get bootstrap working
window.jQuery = jQuery = require 'jquery'

React      = require "react"
bootstrap  = require "bootstrap"
dispatcher = require "./dispatcher/dispatcher.coffee"
Nav        = require './cjsx/nav.cjsx'