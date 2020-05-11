from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
import youtube_dl
from os import listdir, rename, mkdir
from os.path import isfile, join
import os
import zipfile
import sys
import shutil
from webdriver_manager.chrome import ChromeDriverManager

URL = "https://www.youtube.com/results?search_query="

class YoutubeVideoDownloader():   
    linkXpath = "(//a[@id='video-title'])[1]"

    # Setups the chromedriver
    def setup(self):
        options = Options()
        options.add_argument("--headless")
        self.browser = webdriver.Chrome(ChromeDriverManager().install(), chrome_options=options)

    # Cleanups the chromedriver 
    def teardown(self):
        self.browser.close()
         
    # Zips a given directory  
    def zipdir(self, path, ziph):
        for root, _, files in os.walk(path):
            for file in files:
                ziph.write(os.path.join(root, file))
    
    # Zips a given folder         
    def zip_folder(self, folder):
        zipf = zipfile.ZipFile(folder+".zip", 'w', zipfile.ZIP_DEFLATED)
        self.zipdir(folder, zipf)
        zipf.close()
    
    # Downloads a list of songs
    def download_files(self, fileUrls, userFolder):
        ydl_opts = {
            'cachedir': False,
            'format': 'bestaudio/best',
            'postprocessors': [{
                'key': 'FFmpegExtractAudio',
                'preferredcodec': 'mp3',
                'preferredquality': '192',
            }],
            'outtmpl': './' + userFolder + '/%(title)s.%(ext)s',
        }
        with youtube_dl.YoutubeDL(ydl_opts) as ydl:
            ydl.download(fileUrls)
        
    # Main driver for downloading
    def execute_test(self, songs, userFolder):
        # Navigate to url
        song_urls = list()
        
        for song in songs:
            tempUrl = URL + song
        
            self.browser.get(tempUrl)
        
            # click on first
            link = self.browser.find_element(By.XPATH, self.linkXpath)
            link.click()
        
            song_urls.append(self.browser.current_url)
            
        self.teardown()
            
        self.download_files(song_urls, userFolder)
        
        self.zip_folder(userFolder)

        shutil.rmtree('./'+userFolder)
        

def main(songs, user):
    basics = YoutubeVideoDownloader()
    basics.setup()
    basics.execute_test(songs, user)

length = len(sys.argv)

if length < 3:
    print("You did not provide enough arguments")
    exit(-1)

songs = sys.argv[1].split(',')
user = sys.argv[2]

main(songs, user)

