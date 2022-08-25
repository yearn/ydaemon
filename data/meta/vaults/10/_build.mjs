import fs from "fs-extra";
import path from "path";
import { fileURLToPath } from "url";

const __dirname = path.dirname(fileURLToPath(import.meta.url));

let all = [];

for (const name of fs.readdirSync(__dirname)) {
  const file = path.join(__dirname, name);
  const stat = fs.lstatSync(file);
  if (stat.isFile()) {
    const parse = path.parse(file);
    if (parse.ext === ".json") {
      const source = JSON.parse(fs.readFileSync(file));
      source.address = parse.name;
      all.push(source);
    }
  }
}

fs.writeFileSync(path.join(__dirname, "all"), JSON.stringify(all));
