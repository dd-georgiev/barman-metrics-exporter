This projects aims to provide the same metrics as [https://github.com/marcinhlybin/prometheus-barman-exporter](https://github.com/marcinhlybin/prometheus-barman-exporter). 

The reason for creating it is because I had troubles running the [prometheus-barman-exporter](https://github.com/marcinhlybin/prometheus-barman-exporter) on newer versions of Linux with newer version of Barman. 

The project is written in [Golang](https://go.dev/) and compiled as static binary. It is tested on Linux with [glibc](https://www.gnu.org/software/libc/) and [amd64](https://en.wikipedia.org/wiki/X86-64). Cross compilation for Musl based Linux(i.e. alpine) or FreeBSD should be fairly simple. Fell free to raise an issue if you need such binary. 

The metrics which are currently exported are the same as the ones from [prometheus-barman-exporter](https://github.com/marcinhlybin/prometheus-barman-exporter).  
You can find a list and description of them in [docs/spec/metrics.md](docs/spec/metrics.md). The list include planned metrics for the future

More information about installing and operating the exporter can be found in [docs/operations/READE.md](docs/operations/overview.md)


Development documentation can be found in [docs/dev/overview.md](/docs/dev/overview.md)
