'use server';
import { promises as fs } from 'fs';
import path from 'path';

const filePath = path.join(process.cwd(), 'data', 'visits.json');

export async function AddVisitCount() {
  try {
    const fileContents = await fs.readFile(filePath, 'utf8');
    const visits = JSON.parse(fileContents);

    visits.home = (visits.home || 0) + 1;

    await fs.writeFile(filePath, JSON.stringify(visits, null, 2));

  } catch (error) {
    console.error('Error updating visit count:', error);
    throw new Error('Unable to update visit count');
  }
}

export async function loadVisitCount() {
  try {
    const fileContents = await fs.readFile(filePath, 'utf8');
    const visits = JSON.parse(fileContents);

    return visits.home;

  } catch (error) {
    console.error('Error reading file:', error);
    throw new Error('Unable to load visit count');
  }
}
