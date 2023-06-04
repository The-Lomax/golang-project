import os
import sys
from seleniumwire import webdriver
from selenium.webdriver import FirefoxOptions
from selenium.webdriver.common.by import By


def get_options():
    user_agent = 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36'

    options = FirefoxOptions()
    # options.binary_location = '/opt/headless-chromium'
    #options.add_argument("--headless")
    options.add_argument(f'user-agent={user_agent}')
    # options.add_argument("--disable-gpu")
    # options.add_argument("--no-sandbox")
    # options.add_argument('--disable-dev-shm-usage')
    # options.add_argument('--disable-gpu-sandbox')
    # options.add_argument("--single-process")

    return options


def main(chapter):
    URL = f"https://mangakakalot.com/chapter/icz277906593/chapter_{chapter}"
    browser = webdriver.Firefox(options=get_options())
    browser.get(URL)

    elem = browser.find_element(By.CLASS_NAME, 'container-chapter-reader')
    imgs = browser.execute_script('var urls = []; let childs = arguments[0].childNodes; for (let child of childs) {if (child.nodeName == "IMG") {urls.push(child);};}; return urls;', elem)
    
    for img in imgs:
        imgurl = img.get_attribute('src')
        urlsplit = imgurl.split('/')
        foldername = urlsplit[-2]
        filename = urlsplit[-1]
        if not(os.path.isdir(f"Downloads")):
            os.mkdir(f"Downloads")
        if not(os.path.isdir(f"Downloads/{foldername}")):
            os.mkdir(f"Downloads/{foldername}")

        open(f"Downloads/{foldername}/{filename}", 'wb').write(img.screenshot_as_png)

    browser.quit()



if __name__ == "__main__":
    pass
    # chapters = [
    #     29.5, 34.5, 36.5, 43.5, 51.5, 60.5, 68.5, 75.5, 77.5, 85.5, 93.5, 98.5
    # ]
    # for chapter in chapters:
    #     print(f"downloading chapter {chapter}")
    #     main(str(chapter))

    # dirlist = os.listdir("Downloads")
    # for dir in dirlist:
    #     print(f"renaming {dir}")
    #     if dir[0] == 'c':
    #         dirsplit = dir.split('_')
    #         if len(dirsplit) == 2:
    #             os.rename(f"Downloads/{dir}", f"Downloads/Chapter {dirsplit[1]}")
    #         if len(dirsplit) >= 3:
    #             os.rename(f"Downloads/{dir}", f"Downloads/Chapter {dirsplit[1]}.{dirsplit[2]}")
    #     elif dir[0] == 'v':
    #         dirsplit = dir.split('_')
    #         if len(dirsplit) == 3:
    #             os.rename(f"Downloads/{dir}", f"Downloads/Chapter {dirsplit[-1]}")
    #         if len(dirsplit) >= 4:
    #             os.rename(f"Downloads/{dir}", f"Downloads/Chapter {dirsplit[-2]}.{dirsplit[-1]}")
