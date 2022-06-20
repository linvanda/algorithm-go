"".main STEXT size=78 args=0x0 locals=0x20 funcid=0x0
	0x0000 00000 (main.go:5)	TEXT	"".main(SB), ABIInternal, $32-0
	0x0000 00000 (main.go:5)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:5)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:5)	PCDATA	$0, $-2
	0x000d 00013 (main.go:5)	JLS	71
	0x000f 00015 (main.go:5)	PCDATA	$0, $-1
	0x000f 00015 (main.go:5)	SUBQ	$32, SP
	0x0013 00019 (main.go:5)	MOVQ	BP, 24(SP)
	0x0018 00024 (main.go:5)	LEAQ	24(SP), BP
	0x001d 00029 (main.go:5)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:5)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:6)	MOVSD	$f64.4008000000000000(SB), X0
	0x0025 00037 (main.go:6)	MOVSD	X0, (SP)
	0x002a 00042 (main.go:6)	MOVSD	$f64.4010000000000000(SB), X0
	0x0032 00050 (main.go:6)	MOVSD	X0, 8(SP)
	0x0038 00056 (main.go:6)	PCDATA	$1, $0
	0x0038 00056 (main.go:6)	CALL	math.Pow(SB)
	0x003d 00061 (main.go:7)	MOVQ	24(SP), BP
	0x0042 00066 (main.go:7)	ADDQ	$32, SP
	0x0046 00070 (main.go:7)	RET
	0x0047 00071 (main.go:7)	NOP
	0x0047 00071 (main.go:5)	PCDATA	$1, $-1
	0x0047 00071 (main.go:5)	PCDATA	$0, $-2
	0x0047 00071 (main.go:5)	CALL	runtime.morestack_noctxt(SB)
	0x004c 00076 (main.go:5)	PCDATA	$0, $-1
	0x004c 00076 (main.go:5)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 38 48  eH..%....H;a.v8H
	0x0010 83 ec 20 48 89 6c 24 18 48 8d 6c 24 18 f2 0f 10  .. H.l$.H.l$....
	0x0020 05 00 00 00 00 f2 0f 11 04 24 f2 0f 10 05 00 00  .........$......
	0x0030 00 00 f2 0f 11 44 24 08 e8 00 00 00 00 48 8b 6c  .....D$......H.l
	0x0040 24 18 48 83 c4 20 c3 e8 00 00 00 00 eb b2        $.H.. ........
	rel 5+4 t=17 TLS+0
	rel 33+4 t=16 $f64.4008000000000000+0
	rel 46+4 t=16 $f64.4010000000000000+0
	rel 57+4 t=8 math.Pow+0
	rel 72+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 math..inittask+0
type..importpath.math. SRODATA dupok size=7
	0x0000 00 00 04 6d 61 74 68                             ...math
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
$f64.4008000000000000 SRODATA size=8
	0x0000 00 00 00 00 00 00 08 40                          .......@
$f64.4010000000000000 SRODATA size=8
	0x0000 00 00 00 00 00 00 10 40                          .......@
