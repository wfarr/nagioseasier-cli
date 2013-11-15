# Nagios Easier

CLI tool for working with the [module](https://github.com/wfarr/nagioseasier-module).

Gives you nice tabular output!

```
# /usr/local/bin/nagioseasier status analytics1
+--------------------+----------+--------------------------------------------------------------------------+
| Service            | Status   | Details                                                                  |
+--------------------+----------+--------------------------------------------------------------------------+
| analytics1/reports | CRITICAL | FAILURE: a really long error condition with lots of stuff and things you |
|                    |          | probably care about, wrapped neatly <3                                   |
+--------------------+----------+--------------------------------------------------------------------------+
```
