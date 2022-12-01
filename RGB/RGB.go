package RGB

import "fmt"

func RGB_Text(r, g, b int, s string) string {
	color := fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r, g, b, s)
	return color
}

func RGB_BG(r, g, b int, s string) string {
	color := fmt.Sprintf("\x1b[38;2;255;255;255;48;2;%d;%d;%dm%s\x1b[0m", r, g, b, s)
	return color
}

func RGB_TextnBG(r, g, b, r_bg, g_bg, b_bg int, s string) string {
	color := fmt.Sprintf("\x1b[38;2;%d;%d;%d;48;2;%d;%d;%dm%s\x1b[0m", r, g, b, r_bg, g_bg, b_bg, s)
	return color
}

func Hex_Text(h, s string) string {
	if len(h) == 6 {
		r := base16(h[0])*16 + base16(h[1])
		g := base16(h[2])*16 + base16(h[3])
		b := base16(h[4])*16 + base16(h[5])
		return RGB_Text(r, g, b, s)
	}
	return s
}

func Hex_BG(h, s string) string {
	if len(h) == 6 {
		r := base16(h[0])*16 + base16(h[1])
		g := base16(h[2])*16 + base16(h[3])
		b := base16(h[4])*16 + base16(h[5])
		return RGB_BG(r, g, b, s)
	}
	return s
}

func Hex_TextnBG(h1, h2, s string) string {
	if len(h1) == 6 && len(h2) == 6 {
		r := base16(h1[0])*16 + base16(h1[1])
		g := base16(h1[2])*16 + base16(h1[3])
		b := base16(h1[4])*16 + base16(h1[5])
		r_bg := base16(h2[0])*16 + base16(h2[1])
		g_bg := base16(h2[2])*16 + base16(h2[3])
		b_bg := base16(h2[4])*16 + base16(h2[5])
		return RGB_TextnBG(r, g, b, r_bg, g_bg, b_bg, s)
	}
	return s
}

func base16(s byte) int {
	if s >= 65 && s <= 70 {
		return (int(s - 55))
	} else if s >= 97 && s <= 122 {
		return (int(s - 87))
	} else if s >= 48 && s <= 57 {
		return (int(s - 48))
	} else {
		return -1
	}
}
