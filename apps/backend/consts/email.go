package consts

const (
	EmailVerificationSubject = "Dotsync - Email Verification"
	EmailVerificationBody    = `
		<h1>Email Verification</h1>
		<p>Please verify your email address by clicking the link below:</p>
		<p><a href="%s">Verify Email</a></p>
		<p>If the link above does not work, please copy and paste this URL into your browser.</p>
		<p>Verification Link: %s</p>
		<p>This link will expire in 30 minutes.</p>
		<p>If you did not register for Dotsync, please ignore this email.</p>
	`
)
