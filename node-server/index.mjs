import axios from "axios";
import process from "process";
import fs from "fs";
let filePath =''
async function getUrl() {
    let file = fs.readFileSync(filePath);
    let upload_url = "";
    try {
        let { data } = await axios.post(
            "https://api.assemblyai.com/v2/upload",
            file,
            {
                headers: {
                    authorization: "",
                    "Transer-Encoding": "chunked",
                },
            }
        );
        upload_url = data.upload_url;
    } catch (error) {
        console.log("err during getting url");
        process.exit();
    }
    getTranscript(upload_url);
}

async function getTranscript(url) {
    let id = "";
    try {
        let { data } = await axios.post(
            "https://api.assemblyai.com/v2/transcript",
            {
                audio_url: url,
            },
            {
                headers: {
                    authorization: "",
                    "Transer-Encoding": "chunked",
                },
            }
        );
        id = data.id;
    } catch (error) {
        console.log("there was an err");
        console.log(error);
        process.exit();
    }
    getText(id);
}

function getText(id) {
    setTimeout(async () => {
        try {
            let {data} = await axios.get(
                `https://api.assemblyai.com/v2/transcript/${id}`,
                {
                    headers:{
                        authorization: "",
                        "Transer-Encoding": "chunked",
                    }
                }
            );
            if (data['status']=='completed') {
                console.log(data['text'])
                process.exit()
            }
            getText(id)
        } catch (error) {
            console.log('error during getting text Format')
            process.exit() 
        }
    }, 10000);
}

getUrl();
