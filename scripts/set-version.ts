import { execSync } from "child_process";
import fs from "fs";
import path from "path";

/**
 * A script to update version strings in specified files based on the latest Git tag.
 * Used in CI/CD pipelines to ensure version consistency.
 */

// =================== CONFIGURATION ===================
const filesToUpdate: FileConfig[] = [
  {
    path: "./internal/config/version.go",
    searchText: 'const Version = "v0.0.0-dev"',
    replacement: 'const Version = "%"',
  },
];
// =====================================================

interface FileConfig {
  /** The relative path to the file to modify. */
  path: string;
  /** The exact string to find in the file (e.g., the development version line). */
  searchText: string;
  /** The template for the replacement. Use '%' as the version placeholder. */
  replacement: string;
}

function getLatestGitTag(): string {
  try {
    const tag = execSync("git describe --tags --abbrev=0").toString().trim();
    return tag.startsWith("v") ? tag : `v${tag}`;
  } catch (error) {
    console.error(
      "Error: Could not get Git tag. Make sure you are in a Git repository with at least one tag.",
    );
    process.exit(1);
  }
}

function run() {
  const version = getLatestGitTag();
  console.log(`Updating files to version: ${version}`);

  let updatedCount = 0;
  let warnings = 0;

  for (const file of filesToUpdate) {
    const filePath = path.resolve(process.cwd(), file.path);

    if (!fs.existsSync(filePath)) {
      console.warn(`Warning: File not found, skipping: ${file.path}`);
      warnings++;
      continue;
    }

    try {
      const originalContent = fs.readFileSync(filePath, "utf-8");

      if (!originalContent.includes(file.searchText)) {
        console.warn(
          `Warning: Text not found in ${file.path}: "${file.searchText}"`,
        );
        warnings++;
        continue;
      }

      const finalReplacement = file.replacement.replace("%", version);
      const updatedContent = originalContent.replace(
        file.searchText,
        finalReplacement,
      );

      fs.writeFileSync(filePath, updatedContent, "utf-8");
      updatedCount++;
    } catch (error) {
      console.error(`Error processing file ${file.path}:`, error);
      process.exit(1);
    }
  }

  console.log(`\nDone. Updated ${updatedCount} file(s)`);
  if (warnings > 0) {
    console.log(`Encountered ${warnings} warning(s).`);
  }
}

run();
