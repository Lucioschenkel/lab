from pypdf import PdfReader, PdfWriter

stamp = PdfReader("wtr.pdf").pages[0]
writer = PdfWriter(clone_from="merged.pdf")

for page in writer.pages:
    page.merge_page(stamp, over=False)  # here set to False for watermarking

writer.write("watermarked.pdf")
