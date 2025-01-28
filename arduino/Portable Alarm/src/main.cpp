#include <Arduino.h>
#include <RCSwitch.h>

RCSwitch mySwitch = RCSwitch();

#define LED_PIN 13
#define MOTION_SENSOR_PIN 11
#define MOTION_SENSOR2_PIN 9
#define VIBRATION_SENSOR_PIN 3
#define ZUMMER_PIN A1

#define ALARM_PIN 7
#define LIGHT_PIN 8
#define ALARM_LONG 20000
#define RFC_IN_PIN 2
#define A 10790052
#define B 10790049
#define C 10790056
#define D 10790050

bool security = false;
bool panic = false;
unsigned long lightTimer = 1000;
unsigned long dangerousFindedTimer = 0;
unsigned long timer = 0;

void setup()
{
  //Serial.begin(9600);
  digitalWrite(LED_PIN, LOW);
  pinMode(ALARM_PIN, OUTPUT);
  pinMode(LIGHT_PIN, OUTPUT);
  digitalWrite(ALARM_PIN, LOW);
  dangerousFindedTimer = millis();
  timer = millis();
  mySwitch.enableReceive(digitalPinToInterrupt(RFC_IN_PIN));
}

bool hasMotion()
{
  return digitalRead(MOTION_SENSOR_PIN) || digitalRead(MOTION_SENSOR2_PIN);
}

bool hasVibration()
{
  return !digitalRead(VIBRATION_SENSOR_PIN);
}

bool timeToAlarm()
{
  if (((millis() - dangerousFindedTimer) <= 20000) && panic)
  {
    return true;
  }
  else
  {
    return false;
  }
}

bool hasDangerous()
{
  return hasMotion() || hasVibration();
}

void alarm()
{
  if (millis() > 7000)
  {
    if (!panic)
    {
      panic = true;
      dangerousFindedTimer = millis();
    }
    tone(ZUMMER_PIN, millis() % 1000, 500);
    if (millis() % 1000 < 10) {
      lightTimer = millis() + 500;
    }
  }
  digitalWrite(ALARM_PIN, HIGH);
}

void silent()
{
  timer = millis();
  panic = false;
  lightTimer = millis();
  digitalWrite(ALARM_PIN, LOW);
}

void securityOn()
{
  tone(ZUMMER_PIN, 1000, 500);
  security = true;
  panic = false;
  silent();
}

void securityOff()
{
  tone(ZUMMER_PIN, 1000, 500);
  security = false;
  panic = false;
  delay(1000);
  tone(ZUMMER_PIN, 1000, 500);
  lightTimer = 0;
  silent();
}

void lightHandler()
{
  if (lightTimer > millis()) {
    digitalWrite(LIGHT_PIN, HIGH);
  } else {
    digitalWrite(LIGHT_PIN, LOW);
  }
}

void loop()
{
  lightHandler();
  if (mySwitch.available())
  {
    long value = mySwitch.getReceivedValue();
    if (value == A)
    {
      securityOn();
    }
    if (value == B) {
      securityOff();
    }

    mySwitch.resetAvailable();
  }
  if (security)
  {
    if (millis() % 10000 < 100) {
      lightTimer = millis() + 50;
    }
    if (hasDangerous() || timeToAlarm())
    {
      alarm();
    }
    else
    {
      silent();
    }
  }
}
