package main

// #cgo CFLAGS: -DDOT_ONLY=1
// #cgo pkg-config: libgvc libcgraph libcdt libpng
// #cgo LDFLAGS: -lgvc -lcgraph -lcdt -lpng
// #include <gvc.h>
// #include <gvcext.h>
// #include <time.h>
// #ifdef HAVE_UNISTD_H
// #include <unistd.h>
// #endif
// lt_symlist_t lt_preloaded_symbols[] = { { 0, 0 } };
import "C"
import "os"

func main() {
  gvc := C.gvContext()
  //C.GvExitOnUsage = C.int(1);
  argc := C.int(len(os.Args))
  argv := make([]*C.char, argc)
  for i, arg := range os.Args {
    argv[i] = C.CString(arg)
  }
  C.gvParseArgs(gvc, argc, (**C.char)(&argv[0]))

  var prev *C.Agraph_t
  g := C.gvPluginsGraph(gvc)
  if g != nil {
    C.gvLayoutJobs(gvc, g)
    C.gvRenderJobs(gvc, g)
  } else {
    for {
      g := C.gvNextInputGraph(gvc)
      if g == nil {
        break
      }

      if prev != nil {
        C.gvFreeLayout(gvc, prev)
        C.agclose(prev)
      }
      C.gvLayoutJobs(gvc, g)
      C.gvRenderJobs(gvc, g)
      C.gvFinalize(gvc)
      prev = g
    }
  }

  C.gvFreeContext(gvc)
}
