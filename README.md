benchcmd
===========

benchcmd is a tool for benchmark command and show average and stdev.

# Usage

```
Usage of benchcmd:
  -n int
          number of times to run (default 10)
  -summary
          only output summary result
```

# Example Output

```
benchcmd -n 10 'go version'
$ run "go version"
15.711435ms
19.42685ms
19.967876ms
24.896109ms
19.557516ms
21.698995ms
16.499163ms
22.28795ms
23.461732ms
18.743192ms
count:  10 times executed
avg:    20.225081ms
stdev:  2.634089ms
```

# Lisence

Copyright (c) 2015 Keiji Yoshimi

This software is released under the MIT License. http://opensource.org/licenses/mit-license.php
