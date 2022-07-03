from PIL import Image
import os


downloadFolder = "/home/lelouch/Descargas/"
imageFolder="/home/lelouch/Imágenes/"
videoFolder="/home/lelouch/Vídeos/"
documentFolder="/home/lelouch/Documentos/"

def moveFiles():
    for filename in os.listdir(downloadFolder):
        name, extension = os.path.splitext(downloadFolder + filename)

        if extension in [ ".jpg", ".jpeg", ".png" ]:
            picture = Image.open(downloadFolder + filename)
            newName = imageFolder + "compress_" + filename
            picture.save(newName, optimize=True, quality=60)
            print("Image save in ", newName)
            os.remove(downloadFolder + filename)
            print("remove images ", downloadFolder + filename)
        elif extension in [ ".mp4", ".mkv" ]:
            os.rename(downloadFolder+filename, videoFolder+filename)
            print("Video move ", filename, " to ", videoFolder+filename)
        elif extension in [ ".pdf" ]:
            os.rename(downloadFolder+filename, documentFolder+filename)
            print("Document move ", filename, " to ", documentFolder+filename)
        else:
            print("no move name: ", name, "extension: ", extension)


if "__main__" == __name__:
    moveFiles()
