const express = require('express');
const http = require('http');
const {json} = require("express");
const {mergeObjects,merge,checkValue} = require("./utils");
const cors = require('cors');

const app = express();
app.use(cors());

var validUrls = {"urls": {"file_server":"127.0.0.1"}}
var liveUrls = {"urls":{}}


app.get("/api/server/nodes", (req, res) => {
    res.json(liveUrls);
});

app.get("/api/server", (req, res) => {
    let responseData = {
        success : true,
        message : "Hello SCTF",
    }
    res.json(responseData);
});

app.get('/api/server/check', (req, res) => {
    const { hostname, port } = req.query;
    const urlToCheck = `http://${hostname}:${port}`;

    if (!hostname || !port) {
        let responseData = {
            success : false,
            message : 'Hostname and port are mandatory.'
        }
        return res.json(responseData);
    }

    if (!checkValue(validUrls,hostname)){
        let responseData = {
            success : false,
            message : 'Add the domain to the whitelist before proceeding.'
        }
        return res.json(responseData);
    }

    http.get(urlToCheck, (response) => {
        if (response.statusCode !== 200) {
            res.sendStatus(response.statusCode);
            return
        }
        mergeObjects(liveUrls,validUrls,hostname)
        console.log(`Added ${hostname} to liveUrls array.`);

        let resData = '';
        response.on('data', (chunk) => {
            resData += chunk;
        });

        response.on('end', () => {
            let responseData = {
                success : true,
                message : resData
            }
            res.json(responseData)
        });

    }).on('error', (error) => {
        let responseData = {
            success : true,
            message : error
        }
        res.json(responseData)
    });
});

app.get("/api/server/import", (req, res) => {
    const keys = Object.keys(req.query);

    keys.forEach(function(key) {
        merge(validUrls,key,req.query[key])
    });

    console.log(validUrls)
    let responseData = {
        success : true,
        message : "Successfully added the address to the whitelist"
    }
    res.json(responseData)
});

app.listen(3000, () => {
    console.log('Server is running on port 3000');
});
