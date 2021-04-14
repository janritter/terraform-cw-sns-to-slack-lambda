output "lambda_arn" {
  value = aws_lambda_function.lambda.arn
}

output "lambda_function_name" {
  value       = aws_lambda_function.lambda.function_name
  description = "Name of the function"
}

output "iam_role_arn" {
  value       = aws_iam_role.lambda_exec_role.arn
  description = "ARN of the IAM role of the lambda function"
}

output "iam_role_name" {
  value       = aws_iam_role.lambda_exec_role.name
  description = "Name of the IAM role of the lambda function"
}
