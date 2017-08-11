# kicad-golang-bom-generator

A [KiCad](http://www.kicad-pcb.org/) BOM tool generator. It simply
generates a CSV qith grouped entries. An entry is grouped if its
Value, Footprint and any Custom Field are the same.

## Installation

1. Executable installation:
```bash
go get -u github.com/atuleu/kicad-golang-bom-generator
```

2. Within KiCad
  1. Go to BOM Generator
  2. Click Add Plugin
  3. Select  `$GOPATH/bin/kicad-golang-bom-generator`
  4. Edit the command line to be `<path_to_GOBIN>/kicad-golang-bom-generator "%I" "%O.csv"`

## Example Output

In this example, some of the components have a custom `Part Number` in EESCHEMA :

```csv
Quantity,References,Value,Footprint,Part Number
1,"P3",USB,,MC000991
1,"P1",UC_PGM,,188275-6
1,"P4",PWR,,1725656
1,"P5",RS485,,188275-4
1,"P2",TTL_IN,,188275-4
1,"U1",ADM3078E,Housings_SOIC:SOIC-8_3.9x4.9mm_Pitch1.27mm,
13,"C1,C7,C8,C12,C11,C13,C15,C16,C14,C17,C19,C20,C22",100nF,Capacitors_SMD:C_0402,GRM155R61H104KE19D
3,"F1,F2,F3",0.35R, 500mA,Resistors_SMD:R_0402,BLM15AX601SN1D
2,"C2,C3",1uF 35V,Capacitors_SMD:C_0603,GRM188R61H105KAALD
1,"R1",100k,Resistors_SMD:R_0402,MCWR04X1003FTL
2,"D2,D1",30V, 500mA,,PMEG3005EB,115
1,"U2",MIC5239-3.3,Housings_SSOP:MSOP-8_3x3mm_Pitch0.65mm,
1,"C4",22uF 6.3V,Capacitors_Tantalum_SMD:TantalC_SizeR_EIA-2012,TAJP226M006RNJ
1,"U3",FT2232HQ,Housings_DFN_QFN:QFN-64-1EP_9x9mm_Pitch0.5mm,FT2232HQ
1,"Y1",12MHz,,ABM8-12.000MHZ-B2-T
2,"C10,C9",18pF,Capacitors_SMD:C_0402,04025A180FAT2A
2,"FB1,FB2",0.34R, 500mA,Resistors_SMD:R_0402,BLM15AX601SN1D
2,"C6,C5",4.7uF,Capacitors_SMD:C_0603,GRM188R61A475KE15D
1,"R2",12k,Resistors_SMD:R_0402,RC0402FR-0712KL
1,"R3",1k,Resistors_SMD:R_0402,RC0402JR-071KL
1,"C18",3.3uF,Capacitors_SMD:C_0402,GRM188R61A335KE15D
1,"U4",dsPIC33FJ128MC802-I/MM,Housings_DFN_QFN:QFN-24-1EP_5x5mm_Pitch0.65mm,
1,"SW1",RST,,EVPAA202K
2,"R9,R6",10k,Resistors_SMD:R_0402,RC0402FR-0710KL
1,"C21",10uF,Capacitors_SMD:C_0603,GRM188R61A106ME69D
4,"Q1,Q2,Q3,Q4",2SK3541T2L,,2SK3541T2L
1,"D3",GREEN,,KPHHS-1005MGCK
4,"R4,R5,R7,R8",100,Resistors_SMD:R_0402,RC0402JR-07100RL
2,"D4,D5",YELLOW,,KPHHS-1005SYCK
1,"D6",RED,,KPHHS-1005SURCK
```


## LICENSE

This project is licensed under [GPL version 3](https://www.gnu.org/licenses/gpl-3.0.en.html).
