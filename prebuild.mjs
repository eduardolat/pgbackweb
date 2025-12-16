import path from "path";
import fs from "fs";
import { glob } from "glob";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const sourceGlobs = ["./internal/service/**/*.sql"];
const outputFilePath = "./internal/database/dbgen/queries.gen.sql";
const rootDir = __dirname;

function prefixQueriesWithFileName(content, fileName) {
  const lines = content.split("\n");
  const modifiedLines = [];

  for (const line of lines) {
    if (line.startsWith("-- name: ")) {
      modifiedLines.push(`-- file: ${fileName}`);
    }
    modifiedLines.push(line);
  }

  return modifiedLines.join("\n");
}

try {
  const files = [];
  for (const sourceGlob of sourceGlobs) {
    const foundFiles = await glob(path.join(rootDir, sourceGlob));
    files.push(...foundFiles);
  }

  let outFileContent = `-- This file is auto-generated. DO NOT EDIT.\n\n`;
  for (const fileName of files) {
    let content = fs.readFileSync(fileName, "utf-8");
    content = prefixQueriesWithFileName(content, fileName);
    outFileContent += `${content}\n\n`;
  }

  const outputDir = path.dirname(path.join(rootDir, outputFilePath));
  if (!fs.existsSync(outputDir)) {
    fs.mkdirSync(outputDir, { recursive: true });
  }
  
  fs.writeFileSync(path.join(rootDir, outputFilePath), outFileContent);
  console.log("SQLC prebuild completed successfully");
} catch (error) {
  console.error("Error creating SQLC prebuild:", error);
  process.exit(1);
}
