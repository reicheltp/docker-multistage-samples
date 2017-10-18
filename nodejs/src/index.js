import express from 'express'
const app = express()

app.get('/', (req, res) => {
  res.setHeader('Content-Type', 'application/json')
  res.send(JSON.stringify({ message: 'Hallo Docker Berlin! 🐳' }))
})

let port = process.env.PORT || 80
app.listen(port, () => console.log(`Listening on port: ${port}`))
