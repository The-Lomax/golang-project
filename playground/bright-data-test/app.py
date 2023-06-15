import asyncio
import time
import datetime
from playwright.async_api import async_playwright
from python_internal.captcha import get_number_from_image

bdcid = 'hl_c261c78a'
bdzone = 'scrapebrowser'
bdpw = 'zti755qje8aa'

auth = f'brd-customer-{bdcid}-zone-{bdzone}:{bdpw}'
browser_url = f'wss://{auth}@zproxy.lum-superproxy.io:9222'

async def intercept_network_request(request):
    # Print some info about the request
    print("URL:", request.url)
    print("Method:", request.method)
    print("Headers:", request.headers)

    if request.url().endsWith('.png') or request.url().endsWith('.jpg'):
        await request.abort()
        return
    
    # NOTE: You can also await request.abort() to abort the requst1
    await request.continue_()

def getTime(starttime: int) -> int:
    return int(datetime.datetime.now().timestamp()) - starttime

async def main(awb: str):
    starttime = int(datetime.datetime.now().timestamp())
    awbnumber = awb.split('-')[1]
    url = 'http://www.airchinacargo.com/en/search_order.php'
    async with async_playwright() as pw:
        print(getTime(starttime),':', 'connecting');
        browser = await pw.chromium.connect_over_cdp(browser_url)
        print(getTime(starttime),':', 'connected');
        page = await browser.new_page()
        print(getTime(starttime),':', 'goto')
        await page.goto(url, timeout=120000)
        print(getTime(starttime),':', 'page loaded, filling awb')
        awbinput = page.locator('input[id="orders0"]')
        print(getTime(starttime),':', 'found awb input')
        await awbinput.fill(awbnumber)
        print(getTime(starttime),':', 'filled awb input, getting captcha')
        selector = page.locator('img[id="yz"]')
        img = await selector.screenshot()
        print(getTime(starttime),':', 'attempting solve captcha')
        captcha_code = get_number_from_image(img)
        print(getTime(starttime),':', 'captcha: ', captcha_code)
        print(getTime(starttime),':', 'getting captcha input')
        captchainput = page.locator('input[id="usercheckcode"]')
        inner = await captchainput.inner_html()
        print(getTime(starttime),':', 'found captcha input', inner)
        await captchainput.fill(captcha_code)
        print(getTime(starttime),':', 'filled captcha input')
        print(getTime(starttime),':', 'getting search button')
        searchbtn = page.locator('input[id="search"]')
        print(getTime(starttime),':', 'found search button')
        print(getTime(starttime),':', 'clicking search button')
        await searchbtn.click()
        print(getTime(starttime),':', 'clicked search button')
        time.sleep(5)
        print(getTime(starttime),':', 'getting result')
        res = await page.evaluate('() => document.querySelector("body").innerHTML')
        print(getTime(starttime),':', 'got result')
        print(getTime(starttime),':', 'writing result to file')
        with open('result.test', 'w') as f:
            f.write(res)
        print(getTime(starttime),':', 'closing browser')
        await browser.close()


if __name__ == '__main__':
    print('starting request')
    test_awb = "999-69787373"
    asyncio.run(main(test_awb))
    print('complete')
