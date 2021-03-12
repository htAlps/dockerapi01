//  ·───────·───────────────────·───────────────────·───────────────────·───────────────────·───────────────────·───────────────────·───────────────────·
//  λλ 70-httpServers-2.go  ▻21➢03➣10τ14➛00➝13› 
//  λ MainV1
    package main


    import( "fmt"; "os";
            "net/http";
    )

    // C == const, H == Header

    const CH0 = `<h1>   alpsec.dev - pre </h1>    `
    const CH1 = `<textarea cols="30" rows="8">   `
    const CHG = `&#13;`
    const CHH = `  ╔═════════════════════════╗    &#13;`
    const CHI = `  ║  Welcome to AlpSec API  ║    &#13;`
    const CHJ = `  ║  Questions or comments  ║    &#13;`
    const CHK = `  ║     rafa@alpsec.dev     ║    &#13;`
    const CHL = `  ╚═════════════════════════╝    &#13;`
    const CHM = `   alpsec.dev (utf-8 rocks!)     &#13;`
    const CHN = `&#13;`
    const CHO = `</textarea>`
    const CHP = `<br/> `

    // print constants
    const PL0TnFmt          string  = "\n»%s\n\n"                   // for titles
    const PL4T8V8T8V8nFmt   string  = "    %-7T %-7v %-7T %-7v \n"

    // local print & pointer-to helpers
    var pr  = fmt.Printf;
    var spr = fmt.Sprintf;

    func IfErrPanic(ee error) {         // to reduce clutter; but better log than panic
        if ee != nil { panic(ee) }
    }

    func defHandler(ww http.ResponseWriter, req *http.Request) {

        switch req.URL.Path {
            case "/":
                msg := CH0 + CH1 + CHG + CHH + CHI + CHJ + CHK + CHL + CHM + CHN + CHO + CHP
                fmt.Fprintf(ww, msg)
            case "/hello":
                msg := "<h1>hello Gandalf, welcome to Isengard</h1>"
                fmt.Fprintf(ww, msg)
            case "/name":
                msg := "<h1>What is your name?</h1>"
                fmt.Fprintf(ww, msg)
            case "/headers":
                for name, hh := range req.Header {
                    for _, vv := range hh {
                        fmt.Fprintf(ww, name, vv, "\n")
                    }
                }
            default:
                http.NotFound(ww, req)
        }

        for name, hh := range req.Header {
            for _, vv := range hh {
                pr(PL4T8V8T8V8nFmt, name, name, vv, vv)
            }
        }
    }

    func headerHandler(ww http.ResponseWriter, req *http.Request) {

    }

    func main() {

        pr(PL0TnFmt, "instantiating mux & registering handlers on server routes")
        mux := &http.ServeMux{}
        mux.HandleFunc("/", defHandler)                 // http.HandleFunc("/", defHandler)
        mux.HandleFunc("/header", headerHandler)        // http.HandleFunc("/headers", headerHandler)

        port := os.Getenv("muxport"); pr("%s\n", port);
        if port == "" {
            port = "8080"
        }
        pr("listening & serving on port: %s\n", port)
        err := http.ListenAndServe(":" + port, mux)     // err := http.ListenAndServe(":" + port, nil)
        IfErrPanic(err)
    }
//  end

