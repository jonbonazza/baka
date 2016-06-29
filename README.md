# baka
Simple Go app that allows for spawning of a a detached process as the same user. This is useful, for example, in ubuntu-based linux distros when starting applications that utilize the user's enviornment variables since ubuntu shortcuts are run as a special user.

## Usage
```
Usage: baka <app> [args...]
  -pipeErr
        If true, error output will be piped to STDERR (default: false)
  -pipeOutput
        If true, output will be piped to STDOUT (default: false)
  -wait
        If true, baka will wait for app to complete before exiting (default: false)
```        
