{\rtf1\ansi\ansicpg1252\cocoartf2580
\cocoatextscaling0\cocoaplatform0{\fonttbl\f0\fswiss\fcharset0 Helvetica;}
{\colortbl;\red255\green255\blue255;}
{\*\expandedcolortbl;;}
\paperw11900\paperh16840\margl1440\margr1440\vieww11520\viewh8400\viewkind0
\pard\tx566\tx1133\tx1700\tx2267\tx2834\tx3401\tx3968\tx4535\tx5102\tx5669\tx6236\tx6803\pardirnatural\partightenfactor0

\f0\fs24 \cf0 \
package main\
import (\
	"github.com/01-edu/z01"\
)\
func Rot14(s string) string \{\
	SLICE := []string\{\}\
	for _,x := range s\{	\
		\
			if x >= 'a' && x <= 'z' \{\
				x = x+14\
				if x>122\{\
				x-= 122\
				x=96+x\
				\}\
			\}\
			if x >= 'A' && x <= 'Z'\{\
				x = x+14\
				\
				if x>90\{\
				x-= 90\
				x=64+x\
				\}\
			\}\
		xs := string(x)\
		SLICE = append(SLICE, xs)\
	\}\
	strdemerde := ""\
	for boucle := 0; boucle < len(SLICE); boucle++\{\
		\
		strdemerde += SLICE[boucle]	\
	\}\
	//fmt.Print(SLICE)\
	return strdemerde\
\}\
	\
\
\
\
\
func main() \{\
	\
	result := Rot14("Hello! How are You?")\
	for _, r := range result \{\
		z01.PrintRune(r)\
	\}\
	z01.PrintRune('\\n')\
	\
\}\
/*\
if x >= 'a' && x <= 'z' || x >= 'A' && x <= 'Z'\{\
			x = x+14\
			if x != 'a' ||x != 'z' ||x != 'A' ||x != 'Z' \{\
			\
			\}\
			z01.PrintRune(rune(x))\
		\}\
\
\
for _, r := range result \{\
		z01.PrintRune(r)\
	\}\
	z01.PrintRune('\\n')\
*/\
`}