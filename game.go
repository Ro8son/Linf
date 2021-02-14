package main

import(
	"time"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/img"
)

const g = 0.35

type GameState struct {
	objsys ObjectSystem
} 
var game GameState
func (s GameState) ObjectSystem() *ObjectSystem { return &s.objsys }

func (s *GameState) init() {
	s.objsys.init()
	now := time.Now()
	hour := now.Hour()
	bg := &Background{file: "./res/l_night.png", layer: 0}

	if hour <= 8 && hour >= 6{
		bg = &Background{file: "./res/l_sun_rise_yellow.png", layer: 0}
	}else if hour <= 18 && hour >= 9{
		bg = &Background{file: "./res/l_day.png", layer: 0}
	}else if hour <= 21 && hour >= 19{
		bg = &Background{file: "./res/l_sun_set.png", layer: 0}
	}
	s.objsys.addObject(bg, "bg")
	s.objsys.addObject(&Player{ x: 250, y:50, width: 50, height: 50, 
	layer: 1, textureFile: "./res/bird/bird3.png" }, "player")
	s.objsys.addObject(&Pipe{ x: winWidth, height: winHeight/2, 
	layer: 1, position: "bottom", textureFile: "./res/pipe.png" }, "pipe")
}

func (s *GameState) loop() {
	s.objsys.elements["player"].(*Player).update()
	s.objsys.elements["pipe"].(*Pipe).update()
}


type Player struct {
	x, y int32
	width, height int32
	layer int
	textureFile string
	accelX, accelY float32
	rotation float64
}

func (p *Player) X() *int32{ return &p.x }
func (p *Player) Y() *int32{ return &p.y }
func (p *Player) Layer() *int{ return &p.layer }

func (p *Player) update() {
	p.accelX = 0
	if p.accelY < -11 {
		p.accelY = -11
	}
	if p.accelY > 0 {
		p.textureFile = "./res/bird/bird3.png"
	} else {
		p.textureFile = "./res/bird/bird1.png"
	}
	p.accelY += g * float32(deltaTime) / 25
	p.x += int32(p.accelX)
	p.y += int32(p.accelY)
	p.rotation = float64(p.accelY * 5)
}

func (p *Player) draw(renderer *sdl.Renderer) {
	var texture *sdl.Texture

	texfile, _ := img.Load(p.textureFile)
	texture, _ = renderer.CreateTextureFromSurface(texfile)
	defer texfile.Free()
	defer texture.Destroy()

	// renderer.Copy(texture, nil, &sdl.Rect{p.x, p.y, p.width, p.height})
	renderer.CopyEx(texture, nil, &sdl.Rect{p.x-(p.width/2), p.y-(p.height/2), p.width, p.height}, p.rotation, &sdl.Point{p.width/2, p.height/2}, 0)
}

func (p *Player) handleInput(e sdl.Event) {
	switch t := e.(type) {
	case *sdl.KeyboardEvent:
		if t.Keysym.Sym == sdl.GetKeyFromName("Space") && t.State == 1 {
			p.accelY += -24
		}

	}
}

type Pipe struct {
	x, y int32
	height int32
	layer int
	textureFile string
	position string
}

func (p *Pipe) update() {
	p.x += -4
}

func (p *Pipe) X() *int32{ return &p.x }
func (p *Pipe) Y() *int32{ return &p.y }
func (p *Pipe) Layer() *int{ return &p.layer }

func (p *Pipe) draw(renderer *sdl.Renderer) {
	var texture *sdl.Texture

	texfile, _ := img.Load(p.textureFile)
	texture, _ = renderer.CreateTextureFromSurface(texfile)
	defer texfile.Free()
	defer texture.Destroy()

	// renderer.Copy(texture, nil, &sdl.Rect{p.x, p.y, p.width, p.height})
	var width, height int32 = 100, 450
	var rotation float64 = 0
	if p.position == "bottom" || p.position == "" {
		p.y = winHeight + height/2 - p.height
		rotation = 0
	} else {
		p.y = 0 - height/2 + p.height
		rotation = 180
	}
	renderer.CopyEx(texture, nil, &sdl.Rect{p.x-(width/2), p.y-(height/2), width, height}, rotation, &sdl.Point{width/2, height/2}, 0)

}

func (p *Pipe) handleInput(e sdl.Event) { }
