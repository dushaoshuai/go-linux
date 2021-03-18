# Manual
ecgo -- "echo" written in Go, display a line of text.  
"-e" and "--e" are equivalent, "-h" and "-help" are equivalent.  
"--" stops option parsing.

### SYNOPSIS
```
  echo [-neE] [string]...   
  echo --help  
  echo --version
```
  
### DESCRIPTION
```
  -n        do not output the trailing newline  
  -e        enable interpretation of backslash escapes  
  -E        disable interpretation of backslash escapes (default)  
  --help    display help information and exit  
  --version output version information and exit
```
`-e` enables interpretation of following backslash escapes:
```
  \a    alert
  \b    backspace
  \f    form feed
  \n    newline
  \r    carriage return
  \t    horizontal tab
  \v    vertical tab
```


### BUG
1. Don't recognize "\\\\", "\c", "\e", "\0NNN", "\xHH" as "echo" does.
