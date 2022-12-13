import json

print("starting lambda handler")

def respond(err, res=None):
    return {
        "statusCode": "400" if err else "200",
        "body": json.dumps({
            "operation": "lambda",
            "result": "Failure" if err else "Success",
            "status": 400 if err else 200,
            "response": "400 Bad Request" if err else "200 OK",
            "message": "bad request" if err else "Hello from lambda deployed using SAM!"
        }),
        "headers": {
            "Content-Type": "application/json"
        }
    }

def lambda_handler(event, context):
    print("event: " + json.dumps(event))
    return respond(None)
