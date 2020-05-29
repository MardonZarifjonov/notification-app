# Notification app

- For the developement I used golang framework called gin-gonic
- For database I used PostgreSQL

## Installation

1. Clone the repository
2. Install neccessary dependencies ( gin-gonic, postgresql-driver)
3. Setup the .env file for database connection
4. Run the application

## Key points explanation

Using postman you can send POST requests to send notifications<br/>;
In form-data you need to specify the method by which you want to send the notification<br/>;
In order to do so specify one of the key or both of them (send_by_email and send_by_sms) to 1 <br/>;

## Screenshots of execution reslults

- Sending by email<br/>

![Alt text](screenshots/res1.png?raw=true)
<br/>
<br/>
![Alt text](screenshots/res1_1.png?raw=true)
