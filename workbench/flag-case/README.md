
- https://pkg.go.dev/flag

```
-flag
--flag   // double dashes are also permitted
-flag=x
-flag x  // non-boolean flags only
```
Flag parsing stops just before the first non-flag argument ("-" is a non-flag argument) or after the terminator "--".

If a Value has an IsBoolFlag() bool method returning true, the command-line parser makes -name equivalent to -name=true rather than using the next command-line argument.
