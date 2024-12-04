# Fumiko (Multi social chatbot LLM) (Whatsapp/Telegram/Discord) 

A simple tool to set up an LLM service for multiple social networks.
Built with clean code patterns and hexagonal architecture for easier scaling.

## How to whatsapplogin
`go run .\cmd\whatsapplogin\main.go`


## How to init

`go run .\cmd\fumiko\main.go`


## Set a base LLM Context

create new `.env` file and add this line
```env
BASE_LLMCONTEXT="Your name is Fumiko and..."
```