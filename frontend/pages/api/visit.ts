'use server';
// import fs from 'fs';
// import path from 'path';

// export default function handler(req, res) {
//   const filePath = path.join(process.cwd(), 'data', 'visits.json');
//   const fileContents = fs.readFileSync(filePath, 'utf8');
//   const visits = JSON.parse(fileContents);

//   visits.home = (visits.home || 0) + 1;

//   fs.writeFileSync(filePath, JSON.stringify(visits, null, 2));

//   res.status(200).json({ visit_count: visits.home });
// }
