import json
import boto3
import os

print("starting lambda handler")
tableName = os.getenv('TABLE_NAME')
regionName = os.getenv('REGION_NAME')
ddb = boto3.client('dynamodb', region_name=regionName)

def respond(err, res=None):
    ddb.put_item(TableName=tableName, Item={'name':{'S':'Jack'},'lastName':{'S':'Black'}, 'age':{'N':'31'}})
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
