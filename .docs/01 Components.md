This project is designed to make it easy to manage multiple Johnny Decimal projects on a filesystem. Johnny Decimal is a file organization system described at [Johnny Decimal](https://johnnydecimal.com).

Below is a breakdown of the different components.

## Johnny Decimal Number

A number in the form PRO.AC.ID where PRO is a 3-digit Project number, AC is a 2 digit Area and Category number, and ID is a 2 digit ID number. Normally commands are scoped to a single project in which case a JD number with just AC.ID is valid.

## Project

The top level object in a JD system is the project. Each project is represented by an arbitrary 3-digit number between 101-999 and a title in the JD index. On the file system it is represented by a directory.

## Index

The index keeps track of project numbers are taken and the location on the filesystem of the project's directory. For each project it keeps track of each area, category, and id number to enforce the constraints of the JD system. For example, when a new project is created the index will be searched for the next free number available.

## Area

Directory within a project. Labeled with 1 of 10 ranges (10-19, 20-29, ... , 90-99) and a user assigned name. Can be created, deleted, renamed, renumbered.

## Category

A directory within an Area. Labeled with 1 of 10 numbers within the range of it's parent Area. Ex: if the category folder is within area '20-29 Area Name' then it can be labled 21 to 29. Can be created, deleted, renamed, renumbered.

## ID Folder

A directory within a Category. Labeled with the full Johnny Decimal Number for this ID. The number is in the form AC.ID where AC is the Area and Category number and the ID is a double digit number between 01-99 inclusive. Can be created, deleted, renamed, and renumbered.
