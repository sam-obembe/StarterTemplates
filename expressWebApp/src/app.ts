import express from 'express';

const app = express();
const port = 3000;

app.use(express.static('public'))
app.set('view engine', 'ejs')
app.get('/', (req, res) => {
    res.send('Hello World!');
});

app.get('/ping', (req, res) => {
    var resVal: Demo = {
        message: 'pong'
    }
    res.json(resVal)
})


app.get('/demo', function (req, res) {
    let users = [
        { name: 'tobi', email: 'tobi@learnboost.com' },
        { name: 'loki', email: 'loki@learnboost.com' },
        { name: 'jane', email: 'jane@learnboost.com' }
    ];
    res.render('demo', {
        users: users,
        title: "Demo",
        header: "Some users"
    });
});

app.get("/404", function (req, res) {
    res.render('404', { title: "Not Found" });
})

app.get("/500", function (req, res) {
    res.render('500', { title: 'Internal server error' })
})

app.use((req, res, next) => { res.status(404).redirect("/404") });
app.use((req, res, next) => { res.status(500).redirect("/500") });

app.listen(port, () => {
    return console.log(`Express is listening at http://localhost:${port}`);
});