// Copyright (c) 2020 Mohamad All rights reserved
//
// Created by: Mohamad
// Created on: Sep 2020
// This file contains the JS functions for index.html

/*
* This function 
*/
function calculate () {
  // This function calculates the volume of a sphere

  // Recieve input from user
  const radius = parseFloat(document.getElementById('radius').value)

  // Calculate volume
  const volume = (4 / 3) * Math.PI * Math.pow(radius, 3)

  // Output volume to the user
  document.getElementById('answer').innerHTML = "The volume of the sphere is: " + volume.toFixed(2) + " cm³"
}
