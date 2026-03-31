Place your focus.ico file here (16x16, 32x32, 48x48 multi-size ICO).

The build system (Makefile) will embed it into the .exe automatically via go-winres.

If you don't have an icon yet, the app will build and run fine without one —
just skip the `make icon` step and run `go build` directly.

Free icon tools:
- https://www.icoconverter.com  (PNG → ICO)
- https://www.favicon.cc        (draw custom)
