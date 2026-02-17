from PIL import Image, ImageFilter

img = Image.open('./pokedex/pikachu.jpg')
filtered_img = img.filter(ImageFilter.FIND_EDGES)

print(img.format) # JPEG
print(img.size) # (640, 640)

print(dir(img))

filtered_img.save("edges.png", "png")
