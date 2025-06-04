import fs from 'fs';
import path from 'path';

// Function to recursively find all scss files
function findScssFiles(dir, fileList = []) {
  const files = fs.readdirSync(dir);
  
  files.forEach(file => {
    const filePath = path.join(dir, file);
    const stat = fs.statSync(filePath);
    
    if (stat.isDirectory()) {
      findScssFiles(filePath, fileList);
    } else if (path.extname(file) === '.scss') {
      fileList.push(filePath);
    }
  });
  
  return fileList;
}

// Function to update scss files
function updateScssFile(filePath) {
  let content = fs.readFileSync(filePath, 'utf8');
  
  // Replace @import with @use
  content = content.replace(/@import ['"](.+)['"];/g, '@use \'$1\' as *;');
  
  fs.writeFileSync(filePath, content);
  console.log(`Updated: ${filePath}`);
}

// Main function
function main() {
  const stylesDir = path.join(process.cwd(), 'src', 'styles');
  const scssFiles = findScssFiles(stylesDir);
  
  console.log(`Found ${scssFiles.length} SCSS files to update.`);
  
  scssFiles.forEach(file => {
    updateScssFile(file);
  });
  
  console.log('SCSS files update completed.');
}

main(); 