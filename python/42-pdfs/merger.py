import pypdf
import sys

if len(sys.argv) < 3:
    print('Missing arguments. At least two arguments must be given to merge pdfs.')
    exit(0)

files = sys.argv[1:]
merger = pypdf.PdfWriter()

for file in files:
    with open(file, 'rb') as pdf:
        reader = pypdf.PdfReader(pdf)
        for page in range(reader.get_num_pages()):
            merger.add_page(reader.get_page(page))


with open('merged.pdf', 'wb') as merged_pdf:
    merger.write(merged_pdf)
