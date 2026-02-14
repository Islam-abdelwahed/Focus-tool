const fs = require('fs');
const http = require('http');
const path = require('path');
const os = require('os');


// =============================
// CONFIG
// =============================

const blockedDomains = [
    'facebook.com',
    'www.facebook.com',
    'instagram.com',
    'www.instagram.com',
    'x.com',
    'www.x.com',
    'tiktok.com',
    'www.tiktok.com',
    'web.whatsapp.com'
];

const hostsPath = 'C:\\Windows\\System32\\drivers\\etc\\hosts';
const backupPath = path.join(os.homedir(), 'hosts.focus.backup');
const endTimeFile = path.join(__dirname, 'focus.endtime.json');


// =============================
// RESTORE HOSTS
// =============================

function restoreHosts() {

    try {

        if (fs.existsSync(backupPath)) {

            const backup = fs.readFileSync(backupPath, 'utf8');
            fs.writeFileSync(hostsPath, backup);

            fs.unlinkSync(backupPath);

            console.log("✔ Focus stopped. Hosts restored.");
        }

        if (fs.existsSync(endTimeFile))
            fs.unlinkSync(endTimeFile);

    } catch (err) {

        console.log("Restore error:", err.message);

    }

}


// =============================
// START BLOCK SERVER
// =============================

function startServer(endTime) {

    const server = http.createServer((req, res) => {

        const host = req.headers.host;

        const htmlPath = path.join(__dirname, 'blocked.html');

        let html = fs.readFileSync(htmlPath, 'utf8');

        html = html.replace('{{END_TIME}}', endTime);

        res.writeHead(200, {
            'Content-Type': 'text/html'
        });

        res.end(html);

    });

    server.listen(80, '0.0.0.0', () => {

        console.log("Focus server running on port 80");

    });

}



// =============================
// BLOCK WEBSITES
// =============================

function block(minutes) {

    if (fs.existsSync(backupPath)) {

        console.log("⚠ Focus already active.");
        return;

    }

    try {

        const originalHosts = fs.readFileSync(hostsPath, 'utf8');

        fs.writeFileSync(backupPath, originalHosts);

        const entries = [

            '\n# Focus Block START',

            ...blockedDomains.flatMap(domain => [
                `127.0.0.1 ${domain}`,
                `127.0.0.1 www.${domain}`,
            ]),

            '# Focus Block END\n'

        ].join('\n');


        fs.writeFileSync(hostsPath, originalHosts + entries);


        const endTime = Date.now() + minutes * 60000;

        fs.writeFileSync(endTimeFile, JSON.stringify({
            endTime
        }));


        startServer(endTime);


        console.log(`🔥 Focus mode enabled for ${minutes} minutes.`);


        setTimeout(() => {

            restoreHosts();
            process.exit();

        }, minutes * 60000);


    } catch (err) {

        console.log("Block error:", err.message);
        restoreHosts();

    }

}


// =============================
// STATUS
// =============================

function status() {

    if (!fs.existsSync(endTimeFile)) {

        console.log("❌ No active focus session.");
        return;

    }

    const data = JSON.parse(
        fs.readFileSync(endTimeFile)
    );

    const remaining = data.endTime - Date.now();

    if (remaining <= 0) {

        console.log("⏰ Focus finished.");
        restoreHosts();
        return;

    }

    const minutes = Math.floor(remaining / 60000);
    const seconds = Math.floor((remaining % 60000) / 1000);

    console.log(`⏳ Remaining: ${minutes}m ${seconds}s`);

}


// =============================
// CLI HANDLER
// =============================

const args = process.argv.slice(2);


if (args.length === 0) {

    console.log(`
Usage:

focus <minutes>
focus stop
focus status

`);
    process.exit();

}


if (args[0] === "stop") {

    restoreHosts();
    process.exit();

}


if (args[0] === "status") {

    status();
    process.exit();

}


if (!isNaN(args[0])) {

    block(parseInt(args[0]));
    return;

}


console.log("Invalid command.");
