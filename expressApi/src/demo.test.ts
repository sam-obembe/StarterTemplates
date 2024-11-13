import { getDemo } from "./demorunner";

describe('testing runner', function () {
    test('runner function works', function () {
        const d = getDemo("hello world")
        expect(d.message).toBe("hello world")
    })
})