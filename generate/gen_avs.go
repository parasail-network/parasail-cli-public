package generate

var AvsFiles = map[string]string{
	"main.js":      avsLogicJS,
	"package.json": avsLogicPackage,
	"Dockerfile":   avsDockerFile,
}

const avsLogicJS = `const express = require('express');
const app = express();

// Middleware to parse JSON request bodies
app.use(express.json());

async function checkSLA(id) {
    return true
}

// Handle POST requests to the /post endpoint
app.post('/check', async (req, res) => {
    const jsonData = req.body; // Automatically parsed JSON
    console.log('Received JSON:', jsonData);

    const result = await checkSLA(jsonData.paymentID)

    res.json({
        status: result ? 1 : 0,
    });
});

const PORT = process.env.PORT || 4003;
app.listen(PORT, () => {
    console.log("Server is running");
});`

const avsLogicPackage = `{
  "name": "avs",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "module": "ESModule",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "keywords": [],
  "author": "",
  "license": "ISC",
  "dependencies": {
    "express": "^4.21.1"
  },
  "devDependencies": {
    "@types/express": "^5.0.0"
  }
}
`

var avsDeployEnvs = `AVS_NAME=
DEPLOYER_PRIVATE_KEY=
OWNER_ADDRESS=
CHAIN_POINT=
BLOCK_TIME=30
REWARD_PER_BLOCK=
DEAL_PAYMENT_DEADLINE=10
COLLATERAL_FACTOR=120
SLASH_BURN_RATIO=50
MAX_CONSECUTIVE_FAILUES=7
TERMINATE_FEE_FACTOR
`

const avsDockerFile = `FROM node:latest

RUN npm install -g npm@10.5.0

WORKDIR /app

COPY package*.json ./

RUN npm install -g ts-node typescript
RUN npm install

COPY . .

ENTRYPOINT [ "node", "main.js" ]`
