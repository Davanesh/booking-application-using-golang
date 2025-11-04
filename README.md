# 🎟️ GoTicket — The Simplest CLI Ticket Booking App  

Welcome to **GoTicket**, a mini GoLang project that lets you book tickets right from your terminal 💻  
It’s simple, fun, and powered by Go routines for realistic async ticket delivery!

---

## 🚀 Features

- 🧍‍♂️ User-friendly CLI interface  
- 💸 Ticket variants — **Standard**, **Premium**, and **Box**  
- ⚡ Real-time ticket sending simulation using **goroutines** and **WaitGroups**  
- 🕒 Tracks booking time and generates unique booking IDs  
- 🧮 Auto-calculates total price per user  
- ✅ Handles invalid inputs and limits ticket availability  

---

## 🧠 Tech Stack

- **Language:** Go (Golang)  
- **Concepts Used:** Structs, Functions, Goroutines, WaitGroups, Switch Statements, Random ID Generation  

---

## 🧰 Project Structure

```
GoTicket/
│
├── go.mod          # Module definition (auto-created with `go mod init`)
├── main.go         # Handles user input and booking flow
├── user.go         # Contains UserData struct and ticket creation logic
└── README.md       # You’re here!
```

---

## 🏃‍♂️ How to Run the Project

### 1️⃣ Clone the Repo
```bash
git clone https://github.com/<your-username>/GoTicket.git
cd GoTicket
```

### 2️⃣ Run the App
```bash
go run .
```

### 3️⃣ Example Output
```
🎟️  Welcome to GoTicket — The Simplest Booking App Ever!
---------------------------------------------------------
We’ve got a total of 50 tickets available right now!

👉 Enter your first name: Davanesh
👉 Enter your last name: Saminathan
📧 Enter your email ID: davaneshsaminathan335@example.com
🎟️ Choose ticket type (1: standard, 2: premium, 3: box):
2
🎫 How many tickets do you want to book? 2

📨 Sending ticket...
---------------------------------------------------------
📩 Sending ticket:
2 ticket(s) for Dava Nesh
To email: dava@example.com
---------------------------------------------------------
✅ Booking confirmed for Dava Nesh (Premium, ₹1998)
```

---

## 📦 Example User Struct
```go
type UserData struct {
	ID            string
	FirstName     string
	LastName      string
	Email         string
	NumOfTickets  uint
	TicketType    string
	TotalPrice    uint32
	BookedAt      time.Time
}
```

---

## 🌟 Future Enhancements
- Add payment simulation  
- Store user data in JSON or DB  
- Add cancellation & refund system  
- Web version with Go Fiber or React frontend  

---

## 🧑‍💻 Author
**Davanesh S**  
Full Stack Developer | Cloud Enthusiast | GoLang Explorer 🚀  
🔗 [LinkedIn](https://www.linkedin.com/in/davanesh-saminathan/)
