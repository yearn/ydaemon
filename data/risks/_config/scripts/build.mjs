import fs from "fs-extra";
import path from "path";
import { execSync } from "child_process";

const IndexName = "index";
const DataDirectory = "data/risks/";
const SchemaDirectory = "data/risks/_config/schema/";
const OutDirectory = "build/risks";

const CustomBuildScript = "_build.mjs";

const TrimmedExtensions = [".json"];
const ExcludedExtensions = [".mjs"];

function build(directory) {
  const map = { files: [], directories: [] };
  const customBuildScript = path.join(directory, CustomBuildScript);
  if (fs.existsSync(customBuildScript)) {
    try {
      const stdout = execSync(`node ${customBuildScript}`);
      process.stdout.write(stdout.toString());
    } catch (error) {
      process.stdout.write(error.stdout.toString());
      process.stderr.write(error.stderr.toString());
      process.exit(error.status || 1);
    }
  }
  for (let name of fs.readdirSync(directory)) {
    if (name.startsWith(".") || name === IndexName) continue;
    const file = path.join(directory, name);
    const stat = fs.lstatSync(file);
    if (stat.isFile()) {
      const parse = path.parse(file);
      if (ExcludedExtensions.includes(parse.ext)) {
        fs.removeSync(file);
        continue;
      }
      if (TrimmedExtensions.includes(parse.ext)) {
        name = parse.name;
        fs.renameSync(file, path.join(directory, name));
      }
      map.files.push(name);
    } else if (stat.isDirectory()) {

      build(file);
      map.directories.push(name);
    }
  }
  fs.writeFileSync(path.join(directory, IndexName), JSON.stringify(map));
}

const cwd = process.cwd();
if (!fs.existsSync(path.join(cwd, ".git"))) {
  console.error("Error: script should be run in the root of the repo.");
  process.exit(1);
}

try {
  if (fs.existsSync(OutDirectory)) {
    fs.removeSync(OutDirectory);
  }
  fs.mkdirSync(OutDirectory, {recursive: true});
  fs.copySync(path.join(DataDirectory, 'networks'), path.join(OutDirectory, 'networks'));
  fs.copySync(SchemaDirectory, path.join(OutDirectory, 'schema'));
  build(OutDirectory);
  console.log("Ok: build artifact generated!");
} catch (error) {
  console.error(error);
  process.exit(1);
}
