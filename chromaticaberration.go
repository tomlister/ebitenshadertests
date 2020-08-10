// +build ignore
package main

var Time float
var Cursor vec2

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	center := vec2(320, 240)
	amount := normalize(center-Cursor).x / 100
	clr := vec3(0, 0, 0)
	clr.r = texture2At(vec2(texCoord.x+amount, texCoord.y)).r
	clr.g = texture2At(texCoord).g
	clr.b = texture2At(vec2(texCoord.x-amount, texCoord.y)).b
	return vec4(clr, 1.0)
}
