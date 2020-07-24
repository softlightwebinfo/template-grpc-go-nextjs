export default (server, {addService}) => {
    server.get("/api/multiply/:a/:b", (req, res) => {
        addService.runService('Multiply', {
            a: req.params.a,
            b: req.params.b
        }, (e, resp) => {
            if (e) {
                return res.status(500).json({error: e.toString(), data: {}})
            }
            return res.json(resp);
        });
    });
}
