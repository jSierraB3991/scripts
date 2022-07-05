from PIL import Image
import os


downloadFolder = "/home/lelouch/Descargas/"
imageFolder="/home/lelouch/Imágenes/"
videoFolder="/home/lelouch/Vídeos/"
documentFolder="/home/lelouch/Documentos/"

class bcolors:
    HEADER = '\033[95m'
    OKBLUE = '\033[94m'
    OKCYAN = '\033[96m'
    OKGREEN = '\033[92m'
    WARNING = '\033[93m'
    FAIL = '\033[91m'
    ENDC = '\033[0m'
    BOLD = '\033[1m'
    UNDERLINE = '\033[4m'

def moveFiles():
    for filename in os.listdir(downloadFolder):
        name, extension = os.path.splitext(downloadFolder + filename)

        if extension != "":
            if extension in [ ".jpg", ".jpeg", ".png" ]:
                picture = Image.open(downloadFolder + filename)
                newName = imageFolder + "compress_" + filename
                picture.save(newName, optimize=True, quality=60)
                print(bcolors.OKGREEN + "Image save in ", newName + bcolors.ENDC)
                os.remove(downloadFolder + filename)
                print(bcolors.OKGREEN + "remove images ", downloadFolder + filename + bcolors.ENDC)
            elif extension in [ ".mp4", ".mkv" ]:
                os.rename(downloadFolder+filename, videoFolder+filename)
                print("Video move ", filename, " to ", videoFolder+filename)
            elif extension in [ ".pdf" ]:
                os.rename(downloadFolder+filename, documentFolder+filename)
                print("Document move ", filename, " to ", documentFolder+filename)
            else:
                print(bcolors.WARNING + "no move name: ", name, "extension: ", extension + bcolors.ENDC)


if "__main__" == __name__:
    moveFiles()
