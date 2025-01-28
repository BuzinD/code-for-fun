#include <Arduino.h>

#define LEFT_MOTORS_1 6
#define LEFT_MOTORS_2 5
#define RIGHT_MOTORS_1 3
#define RIGHT_MOTORS_2 4
#define MAIN_TIMER 0

uint32_t timers[1];
bool d = false;
// put function declarations here:
void moveForward();
void moveBackward();
void stop();
bool checkItsTime(byte timerId, uint32_t period);
void initTimer(byte timerId);

void setup()
{
  pinMode(LEFT_MOTORS_1, OUTPUT);
  pinMode(LEFT_MOTORS_2, OUTPUT);
  pinMode(RIGHT_MOTORS_1, OUTPUT);
  pinMode(RIGHT_MOTORS_2, OUTPUT);
  digitalWrite(LEFT_MOTORS_1, LOW);
  digitalWrite(LEFT_MOTORS_2, LOW);
  digitalWrite(RIGHT_MOTORS_1, LOW);
  digitalWrite(RIGHT_MOTORS_2, LOW);
  initTimer(MAIN_TIMER);
  moveForward();
}

void loop() {
  if (checkItsTime(MAIN_TIMER, 4000)) {
    if (!d) {
      moveBackward();
    } else {
      moveForward();
    }
    d = !d;
  }
}

void moveForward()
{
  stop();
  digitalWrite(LEFT_MOTORS_1, HIGH);
  digitalWrite(RIGHT_MOTORS_1, HIGH);
  digitalWrite(LEFT_MOTORS_2, LOW);
  digitalWrite(RIGHT_MOTORS_2, LOW);
}

void moveBackward()
{
  stop();
  digitalWrite(LEFT_MOTORS_1, LOW);
  digitalWrite(RIGHT_MOTORS_1, LOW);
  digitalWrite(LEFT_MOTORS_2, HIGH);
  digitalWrite(RIGHT_MOTORS_2, HIGH);
}

void stop()
{
  digitalWrite(LEFT_MOTORS_1, HIGH);
  digitalWrite(RIGHT_MOTORS_1, HIGH);
  digitalWrite(LEFT_MOTORS_2, HIGH);
  digitalWrite(RIGHT_MOTORS_2, HIGH);
}

bool checkItsTime(byte timerId, uint32_t period)
{
  if (millis() - timers[timerId] >= period) { // таймер на millis()
    timers[timerId] = millis();
    return true;
  }
  return false;
}

void initTimer(byte timerId) 
{
  timers[timerId] = millis();
}