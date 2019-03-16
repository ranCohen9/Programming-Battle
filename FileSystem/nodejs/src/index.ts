import * as fs from 'fs';
const FilesDirName = 'files';
class FsWorker {

    createDirectory() {
        if (!fs.existsSync(FilesDirName)) {
            fs.mkdir(FilesDirName, () => { });
        } else {
            let files = fs.readdirSync(FilesDirName);
            for (let index = 0; index < files.length; index++) {
                const fileName = files[index];
                fs.unlinkSync(`${FilesDirName}/${fileName}`);
            }
        }
    }

    createFile(fileNumber: number) {
        fs.open(`${FilesDirName}/file-${fileNumber}.txt`, 'w', () => { });
        //fs.openSync(`${FilesDirName}/file-${fileNumber}.txt`,'w');
    }

    createManyFiles() {
        this.createDirectory();
        let start = new Date();
        for (let index = 0; index < 1000; index++) {
            this.createFile(index);
        }
        let stop = new Date();
        //console.log(`start: ${start.getMilliseconds()} | stop: ${stop.getMilliseconds()}`);
        console.log(`Create files elapsed: ${stop.getTime() - start.getTime()}`);
    }

    writeToFile() {
        if (!fs.existsSync(FilesDirName)) {
            fs.mkdirSync(FilesDirName);
        }
        if (fs.existsSync('contentFile.txt')) {
            fs.unlinkSync(`contentFile.txt`);
        }
        let start = new Date();
        for (let index = 0; index < 1000; index++) {
            fs.appendFile(`contentFile.txt`, `index is ${index}\n`, (error) => {
            });
        }
        let stop=new Date();
        console.log(`Write to file elapsed: ${stop.getTime() - start.getTime()} ms`);
    }
}

let fsWorker = new FsWorker();
// fsWorker.createDirectory();
// fsWorker.createFile(0);
fsWorker.createManyFiles();
fsWorker.writeToFile();