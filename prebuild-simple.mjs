import path from "path";
import fs from "fs";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

function findSqlFiles(dir, fileList = []) {
  const files = fs.readdirSync(dir);
  
  for (const file of files) {
    const filePath = path.join(dir, file);
    const stat = fs.statSync(filePath);
    
    if (stat.isDirectory()) {
      findSqlFiles(filePath, fileList);
    } else if (file.endsWith('.sql')) {
      fileList.push(filePath);
    }
  }
  
  return fileList;
}

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
  const serviceDir = path.join(__dirname, "internal/service");
  const files = findSqlFiles(serviceDir);
  
  let outFileContent = `-- This file is auto-generated. DO NOT EDIT.\n\n`;
  for (const fileName of files) {
    let content = fs.readFileSync(fileName, "utf-8");
    content = prefixQueriesWithFileName(content, fileName);
    outFileContent += `${content}\n\n`;
  }

  const outputDir = path.join(__dirname, "internal/database/dbgen");
  if (!fs.existsSync(outputDir)) {
    fs.mkdirSync(outputDir, { recursive: true });
  }
  
  const outputFile = path.join(outputDir, "queries.gen.sql");
  fs.writeFileSync(outputFile, outFileContent);
  console.log("SQLC prebuild completed successfully");
} catch (error) {
  console.error("Error creating SQLC prebuild:", error);
  process.exit(1);
}
