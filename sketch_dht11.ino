// Example testing sketch for various DHT humidity/temperature sensors
// Written by ladyada, public domain
 
#include "DHT.h"
#include "stdio.h"
 
#define DHTPIN 2     // what pin we're connected to
 
// Uncomment whatever type you're using!
#define DHTTYPE DHT11   // DHT 11 
//#define DHTTYPE DHT22   // DHT 22  (AM2302)
//#define DHTTYPE DHT21   // DHT 21 (AM2301)
 
// Connect pin 1 (on the left) of the sensor to +5V
// Connect pin 2 of the sensor to whatever your DHTPIN is
// Connect pin 4 (on the right) of the sensor to GROUND
// Connect a 10K resistor from pin 2 (data) to pin 1 (power) of the sensor
 
DHT dht(DHTPIN, DHTTYPE);

byte byteRead;
void setup() {
  Serial.begin(9600);  
  dht.begin();
}
 
void loop() {
  // Reading temperature or humidity takes about 250 milliseconds!
  // Sensor readings may also be up to 2 seconds 'old' (its a very slow sensor)
  if (Serial.available()) {
    byteRead = Serial.read();
    if (byteRead == 115) {
      float h = dht.readHumidity();
      float t = dht.readTemperature();
      if (isnan(t) || isnan(h)) {
          Serial.println("Failed to read from DHT");
      } else {
        String space = " ";
        String str = t + space + h;
        Serial.println(str);
      }
    }
  }
}
