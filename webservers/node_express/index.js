const express = require('express')
const app = express()
const port = 3004

function getRandomInt(max) {
  return Math.floor(Math.random() * max);
}

function getRandomTemp(min, max){
  let random = Math.floor(Math.random()*max) - min;
  return random===0 ? getRandomTemp(min, max) : random;
}

Date.prototype.addDays = function(days) {
  let date = new Date(this.valueOf());
  date.setDate(date.getDate() + days);
  return date;
}

let summaries = ["Freezing","Bracing","Chilly","Cool","Mild","Warm","Balmy","Hot","Sweltering","Scorching"]
let forecast = class {
  constructor (day) {
    this.date = new Date().addDays(day)
    this.celsius = getRandomTemp(-50,120)
    this.fahrenheit = 32 + (this.celsius / 0.5556)
    this.summary = summaries[getRandomInt(summaries.length)]
  }
}

app.get('/', (req, res) => {
  let temps = [];
  for (let i=0;i<5;i++){
    temps[i] = new forecast(i)
  }
  res.send(JSON.stringify(temps))
})

app.listen(port, () => {
  console.log(`Node running on port ${port}`)
})