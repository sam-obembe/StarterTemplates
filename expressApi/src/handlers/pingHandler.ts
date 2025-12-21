import { Request, Response } from "express";

function pingHandler(req: Request, res: Response) {
    var resVal: Demo = {
        message: 'pong'
    }
    res.json(resVal)
}

export { pingHandler }