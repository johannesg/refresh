/*

Refresh is a command line tool that build and restart web application
when you change Go and other source files.  Refresh will watch for
file events like create, modifiy or delete and it will build and
restart the application.

Installation

The refresh command can be installed using this command:

  go get github.com/baijum/refresh

Usage

Change to the project directory:

  cd /path/to/myapp

Start Refresh:

  refresh

Refresh will watch for file events like create, modifiy or delete and
it will build and restart the application.  If `go build` returns an
error, it will log it in the `tmp` folder.

You can use the `-c` options if you want to specify a config file.  By
default, `.refresh.conf` is used:

  refresh -c .refresh.conf

Here is a sample config file with the default settings:

  root:              .
  tmp_path:          ./tmp
  build_name:        runner-build
  build_log:         runner-build-errors.log
  valid_ext:         .go, .tpl, .tmpl, .html
  build_delay:       600
  colors:            1
  log_color_main:    cyan
  log_color_build:   yellow
  log_color_runner:  green
  log_color_watcher: magenta
  log_color_app:
  exclude_dir:

*/
package main
