resource "null_resource" "build_lambda" {
  provisioner "local-exec" {
    command = "cd ${path.module} && make build"
  }

  triggers = {
    always_run = timestamp()
  }
}

data "archive_file" "lambda" {
  depends_on  = [null_resource.build_lambda]
  type        = "zip"
  source_dir  = "${path.module}/bin/"
  output_path = "${path.module}/cloudwatch-sns-to-slack.zip"
}

resource "aws_lambda_function" "lambda" {
  function_name    = "cloudwatch-sns-to-slack"
  handler          = "cloudwatch-sns-to-slack"
  runtime          = "go1.x"
  filename         = "${path.module}/cloudwatch-sns-to-slack.zip"
  source_code_hash = data.archive_file.lambda.output_base64sha256
  role             = aws_iam_role.lambda_exec_role.arn
  timeout          = 30

  environment {
    variables = {
      WEBHOOK_URL = var.slack_webhook_url
    }
  }

  tags = var.tags
}

resource "aws_cloudwatch_log_group" "lambda" {
  name              = "/aws/lambda/${aws_lambda_function.lambda.function_name}"
  retention_in_days = 14

  tags = var.tags
}

resource "aws_iam_role_policy_attachment" "policy_attachment" {
  role       = aws_iam_role.lambda_exec_role.name
  policy_arn = aws_iam_policy.lambda_execution.arn
}

resource "aws_iam_role" "lambda_exec_role" {
  name               = "cloudwatch-sns-to-slack"
  assume_role_policy = data.aws_iam_policy_document.instance-assume-role-policy.json

  tags = var.tags
}

data "aws_iam_policy_document" "instance-assume-role-policy" {
  statement {
    actions = [
      "sts:AssumeRole",
    ]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_policy" "lambda_execution" {
  policy = data.aws_iam_policy_document.lambda_logging.json
}

data "aws_iam_policy_document" "lambda_logging" {
  statement {
    effect = "Allow"

    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    resources = ["${aws_cloudwatch_log_group.lambda.arn}:*"]
  }
}
