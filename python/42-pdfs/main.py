import pypdf

# read file in binary mode
with open('dummy.pdf', 'rb') as file:
    reader = pypdf.PdfReader(file)
    page = reader.get_page(0)
    with open('tilt.pdf', 'wb') as new_file:
        page.rotate(90)
        writer = pypdf.PdfWriter()
        writer.add_page(page)
        writer.write(new_file)
