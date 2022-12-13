import json
import requests

def lambda_handler(event, context):
    endpoint = f"https://api.dev.cargoai.co/solutions/track?awb={event['awb']}"
    headers = {
        "x-api-key": "04qQ3kba9plqnXdaUSytaumoRb640D33HzJYFJpe"
    }

    resp = requests.get(endpoint, headers=headers).json()
    return json.dumps(resp)



# no data in awb
{
    "awb": "000-00000184"
}

# data present but not delivered
{
    "awb": "615-35926236"
}

# delivered but not to destination
{
    "awb": "615-35926236"
}

# delivered to destination
{
    "awb": "180-34817042"
}