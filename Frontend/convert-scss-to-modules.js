const fs = require('fs');
const path = require('path');

// Define the source directory for pages styles
const stylesPagesDir = path.join(__dirname, 'src', 'styles', 'pages');

console.log(`Processing files in: ${stylesPagesDir}`);

// Function to convert dash-case class names to camelCase
function toCamelCase(str) {
  return str.replace(/([-_][a-z])/g, group => 
    group.toUpperCase()
      .replace('-', '')
      .replace('_', '')
  );
}

// Function to convert a SCSS file to a CSS Module
function convertToModule(file) {
  if (file.endsWith('.module.scss')) {
    console.log(`Skipping ${file} - already a module`);
    return;
  }

  const sourcePath = path.join(stylesPagesDir, file);
  const outputPath = path.join(stylesPagesDir, file.replace('.scss', '.module.scss'));
  
  try {
    console.log(`Reading file: ${sourcePath}`);
    let content = fs.readFileSync(sourcePath, 'utf8');
    
    // Replace @use imports to include proper namespace
    content = content.replace(/@use ['"]\.\.\/variables\.scss['"] as \*/g, "@use '../../styles/variables.scss' as vars");
    content = content.replace(/@use ['"]\.\.\/mixins\.scss['"] as \*/g, "@use '../../styles/mixins.scss' as mixins");
    
    // Add vars. prefix to all SCSS variables
    content = content.replace(/\$([a-zA-Z0-9_-]+)/g, 'vars.$1');
    
    // Replace @include with mixins namespace
    content = content.replace(/@include ([a-zA-Z0-9_-]+)/g, '@include mixins.$1');
    
    // Convert class names from dash-case to camelCase
    const classRegex = /\.([a-z0-9-_]+)(\s*\{|\s*,|\s*\.|\s*>|\s*\+|\s*~|\s*:|\s*\[)/g;
    const matches = [...content.matchAll(classRegex)];
    
    // Start from the end to avoid messing up positions
    for (let i = matches.length - 1; i >= 0; i--) {
      const match = matches[i];
      const className = match[1];
      const punctuation = match[2];
      
      // Skip if already camelCase or if it's a pseudo-class
      if (className.includes('-')) {
        const camelClassName = toCamelCase(className);
        const start = match.index;
        const end = start + match[0].length;
        
        content = 
          content.substring(0, start) + 
          '.' + camelClassName + punctuation + 
          content.substring(end);
      }
    }
    
    fs.writeFileSync(outputPath, content, 'utf8');
    console.log(`Converted ${file} to ${file.replace('.scss', '.module.scss')}`);
  } catch (error) {
    console.error(`Error processing ${file}:`, error);
  }
}

// Main function to process all SCSS files
function main() {
  try {
    if (!fs.existsSync(stylesPagesDir)) {
      console.error(`Directory not found: ${stylesPagesDir}`);
      return;
    }
    
    const files = fs.readdirSync(stylesPagesDir)
      .filter(file => file.endsWith('.scss') && !file.endsWith('.module.scss'));
    
    console.log(`Found ${files.length} SCSS files to convert: ${files.join(', ')}`);
    
    // Process each file
    for (const file of files) {
      convertToModule(file);
    }
    
    console.log('Conversion complete!');
  } catch (error) {
    console.error('Error:', error);
  }
}

main(); 