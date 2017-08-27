package example
import ("github.com/MartinSteffen/tiny/absynt"  // that should do it.
)
// what is uncler is if it is taken from the net. Probably no


var F = &absynt.ID{I:"s"}   // factor
var T = &absynt.FACTOR{F}   // term
var SE = &absynt.TERM{T}    // simple expr
var E = &absynt.SIMPLEEXPR{SE}  // expr 
var SR = &absynt.READ{I:"x"}    // read stmt
var SL1 = SR             // stmt (list) 
var SL2 = SL1            // stmt (list) 
var S = &absynt.IF{E,SL1,SL2}     // stmt 
var AO =  &absynt.PLUS{}           // add op "+"
var F2 = &absynt.NUMBER{N:42}      // factor
var T2 = &absynt.FACTOR{F2}      // term
var SE2 =  &absynt.TERM{T2}       // simple expression
var AO4 = &absynt.MINUS{}         // add op "-"
var F5  = &absynt.ID{"f5"}
var T5 = &absynt.FACTOR{F5}
var SE5 = &absynt.TERM{T5}      
var SE4 =  &absynt.ADDEXPR{AO4,SE5,T2}    // simple expression
var E3  = &absynt.SIMPLEEXPR{SE4} // expression
var F3  = &absynt.EXPR{E3}      // factor
var T3 = &absynt.FACTOR{F3}           // term
var SE3 = &absynt.ADDEXPR{AO4,SE2,T3}  // simple expression
var S2 = &absynt.REPEAT{S,E3}     // stmt 


//repeat (if "s" (read "x") (read "x"))
// until  ("f5" - 42)
