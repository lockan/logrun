# Logrun

Logrun is a Go logging library that wraps logrus (github.com/sirupsen/logrus package), but addes a few custom changes. 

* Logrun uses the runtime package to determine the file and line where the log was located.
* Logrun outputs log messages using the WithFields options and includes 
	* The calling file
	* The line number
	* The log level (info, warn, debug, etc)
	* The log message

This is useful for debugging program errors from log output since it is easy to see exactly where in the program the log came from, similar to what you might see in a stack trace.
