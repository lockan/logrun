# Logrun

Logrun is a Go logging library that wraps logrus (http://github.com/sirupsen/logrus) package, but addes a few custom changes. 

* Logrun uses the runtime package to determine the file and line of the call to the log function. 
* All log messages use the 'WithFields' options to include:
	* The calling file
	* The line number
	* The log level (info, warn, debug, etc)
	* The log message

This can useful for debugging program errors from log output since it is easy to see exactly where in the program the log came from, similar to what you might see in a stack trace.

### Dependencies: ###
http://github.com/sirupsen/logrus
