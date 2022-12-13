# create s3 bucket
aws s3 mb s3://sam-lambda-demo-the-lomax

# package CF template
sam package --s3-bucket sam-lambda-demo-the-lomax --template-file template.yaml --output-template-file gen/cloudformation.yaml

# deploy
sam deploy --template-file gen/cloudformation.yaml --stack-name hello-lambda-stack --capabilities CAPABILITY_IAM