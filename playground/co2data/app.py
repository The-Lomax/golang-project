import requests
import json

url = "https://cargo2zero.p.rapidapi.com/co2calculation"

headers = {
	"X-RapidAPI-Key": "6cf90677femsh92c1376b11be7cap10099ajsnc7400ea4ecad",
	"X-RapidAPI-Host": "cargo2zero.p.rapidapi.com"
}

data = {}
with open("awbs.json", "r") as f:
    data = json.load(f)

awbs = data['awbs']

co2_data = []

if len(awbs) > 0:
    count = 0
    for awb in awbs:
        count += 1
        querystring = {"awb": awb}
        resp = requests.get(url, headers=headers, params=querystring)
        print(f"[{count}/{len(awbs)}] --> awb: {awb} --> {resp.status_code}")
        awb_co2 = resp.json()
        co2_data.append({awb: awb_co2})

with open("co2.json", "w") as f:
    f.write(json.dumps(co2_data))