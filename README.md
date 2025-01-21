# Email Open Tracker in Golang

This project is a Golang-based email tracking system that embeds unique tracking pixels for each email recipient. By including a 1-pixel image in the email body with a unique tracking ID, you can determine if and when the recipient opens the email.

## Features

- **Email Sending**: Sends emails with embedded tracking pixel.
- **Email Tracking**: Detects when the email is opened by the recipient.
- **Multiple Recipients**: Supports multiple "To" and "CC" email addresses.
- **SMTP Integration**: Uses the Golang `smtp` package for sending emails.


## Prerequisites

- Go 1.18 or higher
- An SMTP server (e.g., Gmail, SendGrid, Mailgun, or your own SMTP server)
- Basic knowledge of Golang and email protocols

## Setup

### Step 1: Clone the Repository

```bash
git clone https://github.com/yourusername/email-tracker.git
cd email-tracker
```
### Step 2: Run the following command to install any required Go dependencies:

```bash
go get -u
```
### Step 3: Create a config.json file in the project root directory with the following structure:

```bash
{
  "smtp_server": "smtp.example.com",
  "smtp_port": "587",
  "smtp_username": "your-email@example.com",
  "smtp_password": "your-email-password"
}
```
## Endpoints

### `POST /send`
#### Description:
Sends an email with a tracking pixel embedded in the body. The request body should contain the email recipients, subject, and body content. The response will include a `trackerId` for each recipient, which can be used to track how many times the email was opened.

#### Request Body:
```json
{
  "to": ["recipient1@example.com", "recipient2@example.com"],
  "cc": ["cc1@example.com"],
  "subject": "Test Email with Open Tracker",
  "body": "Hello, this is a test email. Please open it to test the tracker. <img src='http://your-tracker-url.com/pixel' width='1' height='1'/>"
  "is_html":true
}
```
### `GET /track-count/{trackerId}`
#### Description:
Retrieves the number of times an email with a specific trackerId has been opened. This endpoint provides insights into how many times the email was accessed by a recipient.

```
GET http://your-tracker-url.com/track-count/12345
```

