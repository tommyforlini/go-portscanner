# go-portscanner

Demo Sync and Async Port Scanning

## Running

```bash
go run .
```

> Sample result

```bash
2020/06/10 23:25:46 Scanning Host: localhost | Protocol: tcp


2020/06/10 23:25:46 Start:       execute scanSync
2020/06/10 23:26:00 sync slice size 1
2020/06/10 23:26:00 sync results [{80 tcp true}]
2020/06/10 23:26:00 End:         execute scanSync took 13.792467535s
2020/06/10 23:26:00 =======

2020/06/10 23:26:00 Start:       execute scanASync
2020/06/10 23:26:00 async slice size 1
2020/06/10 23:26:00 async results [{80 tcp true}]
2020/06/10 23:26:00 End:         execute scanASync took 113.515364ms
2020/06/10 23:26:00 =======

2020/06/10 23:26:00 Start:       execute scanASyncNoChannel
2020/06/10 23:26:00 scanASyncNoChannel slice size 1
2020/06/10 23:26:00 scanASyncNoChannel results [{80 tcp true}]
2020/06/10 23:26:00 End:         execute scanASyncNoChannel took 124.287618ms
2020/06/10 23:26:00 =======
```
