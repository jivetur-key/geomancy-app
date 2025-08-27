# geomancy-app

A command-line geomancy engine for generating and interpreting traditional shield readings.

## API Usage

This project uses the Google Gemini API for generating interpretations. To run this application, you will need your own API key.

1. Obtain an API key from Google AI Studio.
2. Set an evironment variable before running the application.
3. Set the GEMINI_API_KEY environment variable.

export GEMINI_API_KEY="YOUR_API_KEY"

Note: Your use of the Gemini API is subject to the Google Generative AI API Terms of Service.

# Installation

Clone the repo and build:

```sh
git clone https://github.com/jivetur-key/geomancy-app.git
cd geomancer-app
make

# Usage
./geomancer
./geomancer -planet sun

