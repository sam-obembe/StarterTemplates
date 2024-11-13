import express from 'express';
const app = express();
const port = process.env.PORT;
app.use(express.static("public"));
app.use((req, res, next) => {
    console.log(req.url);
    next();
});
app.get('/', (req, res) => {
    res.send('Hello World!');
});
app.get('/swagger', (req, res) => {
    res.redirect("/swagger.html");
});
app.get('/ping', (req, res) => {
    var resVal = {
        message: 'pong'
    };
    res.json(resVal);
});
app.get("/404", function (req, res) {
    res.json({ message: "not found" });
});
app.get("/500", function (req, res) {
    res.json({ message: "internal error" });
});
app.use((req, res, next) => {
    res.status(200).json({ data: Object.assign({}, req.body) });
    next();
});
app.use((req, res, next) => { res.status(404).redirect("/404"); });
app.use((req, res, next) => { res.status(500).redirect("/500"); });
app.listen(port, () => {
    return console.log(`Express is listening at http://localhost:${port}`);
});
//# sourceMappingURL=app.js.map