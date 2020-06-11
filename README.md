
NAME
----
Marc-Util-Go - A marc record parser in Go-Lang 

SYNOPSIS
--------

`marcUtil -f` <filename.mrc.>

DESCRIPTION
-----------
Caution: This is Alpha software!

The current goals of `marc-util-go` is to explore and learn about the Marc
record specification and Go-Lang specific software design considerations.


OPTIONS
-------

`-f`
  <a  marc file of parse>

COMPILE AND RUN
---------------

```bash
git clone https://gihub.com/dcslagel/marc-util-go
cd marc-util-go 
make clean
make
make test
```

This will output:

```
go test cmd/marcUtil/*.go -v
=== RUN   TestLeader
24
--- PASS: TestLeader (0.00s)
PASS
ok      command-line-arguments  0.005s
```

PROJECT ROADMAP AND PLANNING
----------------------------

marc-util-go:s's project road-map is managed in github milestones at:
- https://github.com/dcslagel/marc-util-go/milestones

To request a feature or report a bug create an issue at:
- https://github.com/dcslagel/marc-util-go/issues

BUGS
----

- Functionality is severly limited to reading the leader fields from a marc file
  containing only a leader data line.  Additional functionality will be added
  in future iterations.

- Report bugs by creating an issue at:
  - https://github.com/dcslagel/marc-util-go/issues

COPYRIGHT
------

Copyright (c) 2020 DC Slagel and contributers
