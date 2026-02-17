from PIL import Image

img = Image.open('./pokedex/pikachu.jpg')
filtered_img = img.convert('L')

filtered_img.save("gray.png", "png")

filtered_img.show()
