import express from 'express';
const app = express();
const port = 3000;

app.get('/', (req, res) => {
    res.send('Hello World!');
});

app.get('/ping', (req, res) => {
    var resVal: Demo = {
        message: 'pong'
    }
    res.json(resVal)
})

app.listen(port, () => {
    return console.log(`Express is listening at http://localhost:${port}`);
});