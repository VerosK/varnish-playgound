# varnish-playground

This docker-compose can be used to evaluate Varnish VCL features
without breaking production environment.

Directory `containers` contains Docker image with working and
broken server. Run `make` to build the local version of image.

Change `defaults.vcl` to modify varnish behaviour.

Run `docker-compose up` to watch backend outputs.

## Ports

There are two backends `working` and `broken` both listening on 
port 8983.  

Varnish listens on port 8983 (exported to localhost) and
tries to distribute requests to both backends.


