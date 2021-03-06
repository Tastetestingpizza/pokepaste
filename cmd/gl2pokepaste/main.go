package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"os"
)

func main() {
	for mon := 1; mon <= 807; mon++ {
		for form := 0; ; form++ {
			code := (mon + form<<16) * 0x159a55e5 & 0xffffff
			url := fmt.Sprintf("http://n-3ds-pgl-contents.pokemon-gl.com/share/images/pokemon/300/%06x.png", code)

			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				break
			}

			src, err := png.Decode(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			dst := image.NewNRGBA(image.Rect(0, 0, 300, 300))

			draw.Draw(dst, image.Rect(0, 0, 150, 150), src, image.Pt(150, 150), draw.Src)
			draw.Draw(dst, image.Rect(0, 150, 150, 300), src, image.Pt(150, 0), draw.Src)
			draw.Draw(dst, image.Rect(150, 0, 300, 150), src, image.Pt(0, 150), draw.Src)
			draw.Draw(dst, image.Rect(150, 150, 300, 300), src, image.Pt(0, 0), draw.Src)

			name := fmt.Sprintf("%d-%d.png", mon, form)

			file, err := os.Create(name)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			err = png.Encode(file, dst)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("%s -> %s", url, name)
		}
	}
}
