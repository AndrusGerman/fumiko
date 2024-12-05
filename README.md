# Fumiko (Multi social chatbot LLM) (Whatsapp/Telegram/Discord) 

A simple tool to set up an LLM service for multiple social networks.
Built with clean code patterns and hexagonal architecture for easier scaling.


## How to init (All)

`go run .\cmd\fumiko\main.go -discord -whatsapp -telegram`

You can remove the 'WhatsApp, Telegram or Discord' flags to deactivate the social networks that you are not going to use


## Set a base LLM Context
create new `.env` file and add this line
```env
BASE_LLMCONTEXT="Your name is Fumiko and..."
```


## How to login in whatsapp
`go run .\cmd\whatsapplogin\main.go -whatsapp`

## How to add token for discord and telegram
create new `.env` file and add this line
```env
DISCORD_TOKEN="Example token for discord"
TELEGRAM_TOKEN="Example token for telegram"
```
