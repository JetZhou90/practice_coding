import request, base64

url = 'http://localhost:9090/receiver'

# headers = {
#     'Connection': 'Keep-Alive',
#     'Accept': 'text/html, application/xhtml+xml, */*',
#     'Accept-Language': 'en-US,en;q=0.8,zh-Hans-CN;q=0.5,zh-Hans;q=0.3',
#     'User-Agent': 'Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; rv:11.0) like Gecko'
#     }

with open('test.jpg', 'rb') as f:
    image_data= f.read()
image_data = base64.b64encode(image_data).decode()
res={'image':image_data}
print(image_data)
response = request.post(url, data=res)
