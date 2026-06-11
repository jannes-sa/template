# 3D Renewable Energy Animation

An interactive Three.js scene visualizing how a renewable energy park works:
solar panels track the sun across a full day/night cycle, wind turbines spin,
and glowing energy particles flow from the generators into a battery storage
unit and onward to a city whose windows light up at night.

## How to run

Everything is a single self-contained file (`index.html`) — Three.js is loaded
from a CDN via an import map, so no build step or `npm install` is needed.

Because it uses ES modules, serve it over HTTP rather than opening the file
directly:

```bash
# from the repo root
cd 3d
python3 -m http.server 8000
# then open http://localhost:8000
```

Any static server works (`npx serve`, `php -S`, etc.).

## What's in the scene

- **Animated sun cycle** — the sun arcs across a gradient-shader sky; lighting,
  fog, and sky colors shift between day and dusk.
- **Solar farm** — a 4×6 grid of PV modules on single-axis trackers that tilt
  to follow the sun. Solar energy flow dims at night.
- **Wind turbines** — four turbines with independently spinning rotors.
- **Energy flow** — additive-blended particles travel along bezier curves:
  solar (amber) and wind (cyan) into the battery, then green flow to the city.
- **City** — procedural low-poly buildings whose windows glow brighter at
  night, drawing on stored energy.
- **Live stats HUD** — simulated solar/wind output (MW) and cumulative CO₂
  avoided.
- **Controls** — drag to orbit, scroll to zoom (OrbitControls).

## Background research: SUN Energy

The scene's solar-first theme is inspired by [SUN Energy](https://sunenergy.id/en)
(PT Surya Utama Nuansa), Indonesia's leading solar energy developer:

- Founded in **2016**, focused on commercial & industrial (C&I) solar in
  Indonesia, with operating assets in Australia, Thailand, and Vietnam and
  offices in Jakarta, Semarang, Surabaya, and Singapore.
- Portfolio of **280+ MWp** across more than 35 industries; named
  "Indonesia's Leading Solar Project Developer" by the National Energy Council.
- Installed the largest industrial rooftop solar system in Indonesia at
  **18.6 MWp** (the solar-output figure in the HUD nods to this).
- Offers Sustainability-as-a-Service: solar system development, energy
  management, and sustainable financing (including low/zero-upfront-cost
  models) to help clients cut emissions and improve ESG performance.
- Notable clients include Unilever, Djarum, Sido Muncul, Fonterra,
  Astra Honda Motor, and Astra Daihatsu Motor.
- Certified ISO 9001 (quality), ISO 14001 (environmental), and ISO 45001
  (occupational health & safety).
