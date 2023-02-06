const multer = require('multer');
const express = require('express');
const getDataCID = require("./getCID");
const app = express()
const port = 3000

//Configuration for Multer
const multerStorage = multer.diskStorage({
    destination: (req, file, cb) => {
      cb(null, "./files");
    },
    filename: (req, file, cb) => {
      const ext = file.mimetype.split("/")[1];
      cb(null, `new.${ext}`);
    },
  });
const upload = multer({ storage: multerStorage });


app.post("/upload", upload.single("demo_image"), (req, res) => {
    res.send(getDataCID("new.png"))
 });

app.listen(port, () => {
console.log(`Example app listening on port ${port}`)
})