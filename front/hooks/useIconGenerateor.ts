import { useEffect, useState } from "react"

const useIconGenerator = () => {
  const [icon, setIcon] = useState<Blob | null>(null)
  useEffect(() => {
    const cell = new Image()
    cell.onload = () => {
      const canvas = document.createElement('canvas')
      canvas.width = 512
      canvas.height = 512
      const ctx = canvas.getContext('2d')
      if(!ctx) return
      let r, g, b
      r = g = b = 0.93
      const s = 0.44
      const h = Math.random() * 6
      const f = h - Math.floor(h)
      if (h < 1) {
        g *= 1 - s * (1 - f)
        b *= 1 - s
      } else if (h < 2) {
        r *= 1 - s * f
        b *= 1 - s
      } else if (h < 3) {
        r *= 1 - s
        b *= 1 - s * (1 - f)
      } else if (h < 4) {
        r *= 1 - s * (1 - f)
        g *= 1 - s
      } else if (h < 5) {
        r *= 1 - s * (1 - f)
        g *= 1 - s
      } else {
        g *= 1 - s
        b *= 1 - s * f
      }
      ctx.fillStyle = `rgb(${r * 255},${g * 255},${b * 255})`
      ctx.fillRect(0, 0, 512, 512)
      ctx.drawImage(cell, 0, 0, 512, 512)
      canvas.toBlob((blob) => {
        if(blob){
          setIcon(blob)
        }
      }, 'image/png')
    }
    cell.src = '/img/icon_cell.png'
  }, [])
  return {
    icon,
  }
}

export default useIconGenerator
