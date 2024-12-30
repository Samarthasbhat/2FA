# 2 Factor Authentication
This project implements two-factor authentication (2FA) in Go, providing an extra layer of security for user accounts. The application uses QR codes for setting up the 2FA and verifies the one-time passwords (OTPs) for each login attempt.

## Features

- User registration and login.
- Two-factor authentication (2FA) using Time-based One-Time Passwords (TOTP).
- QR code generation for setting up 2FA with an authenticator app (e.g., Google Authenticator).
- Secure session management for authenticated users.

## Directory Structure

```plaintext
Samarthasbhat-2FA/
├── main.go               # Entry point for the Go application
├── go.sum                # Go dependencies checksum file
├── users/                # User management logic
│   └── users.go          # User authentication and management logic
├── go.mod                # Go modules for managing dependencies
├── README.md             # Project documentation (this file)
└── templates/            # HTML templates for rendering pages
    ├── dashboard.html    # User dashboard after successful login
    ├── validate.html     # OTP input page for 2FA validation
    ├── index.html        # Homepage for the application
    ├── qrcode.html       # Displays QR code for 2FA setup
    ├── login.html        # Login page
    └── error.html        # Error page for failed actions
```

### Clone the repository

```bash
git clone https://github.com/your-username/Samarthasbhat-2FA.git
cd Samarthasbhat-2FA
```
