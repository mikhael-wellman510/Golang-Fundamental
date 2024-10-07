const data = [
  { line: "upstream", speed: 36, classification: "bus" },
  { line: "upstream", speed: 58, classification: "car" },
  { line: "upstream", speed: 14, classification: "bus" },
  { line: "downstream", speed: 47, classification: "bus" },
  { line: "upstream", speed: 84, classification: "car" },
  { line: "downstream", speed: 78, classification: "car" },
  { line: "upstream", speed: 40, classification: "car" },
  { line: "upstream", speed: 64, classification: "bus" },
  { line: "upstream", speed: 64, classification: "car" },
  { line: "downstream", speed: 17, classification: "car" },
  { line: "downstream", speed: 55, classification: "bus" },
  { line: "downstream", speed: 26, classification: "car" },
  { line: "upstream", speed: 55, classification: "car" },
  { line: "downstream", speed: 18, classification: "bus" },
  { line: "upstream", speed: 59, classification: "2-axle truck" },
  { line: "upstream", speed: 75, classification: "car" },
  { line: "upstream", speed: 72, classification: "2-axle truck" },
  { line: "downstream", speed: 60, classification: "2-axle truck" },
  { line: "downstream", speed: 74, classification: "car" },
  { line: "upstream", speed: 47, classification: "car" },
  { line: "downstream", speed: 89, classification: "car" },
  { line: "downstream", speed: 76, classification: "bus" },
  { line: "upstream", speed: 43, classification: "car" },
  { line: "downstream", speed: 20, classification: "car" },
  { line: "upstream", speed: 44, classification: "2-axle truck" },
  { line: "upstream", speed: 97, classification: "car" },
  { line: "downstream", speed: 69, classification: "bus" },
  { line: "downstream", speed: 87, classification: "car" },
  { line: "upstream", speed: 90, classification: "bus" },
  { line: "upstream", speed: 42, classification: "car" },
  { line: "upstream", speed: 39, classification: "2-axle truck" },
  { line: "upstream", speed: 49, classification: "2-axle truck" },
  { line: "upstream", speed: 22, classification: "car" },
  { line: "downstream", speed: 99, classification: "car" },
  { line: "downstream", speed: 38, classification: "car" },
  { line: "downstream", speed: 59, classification: "bus" },
  { line: "downstream", speed: 36, classification: "2-axle truck" },
  { line: "downstream", speed: 73, classification: "2-axle truck" },
  { line: "downstream", speed: 45, classification: "bus" },
  { line: "upstream", speed: 10, classification: "2-axle truck" },
  { line: "downstream", speed: 17, classification: "car" },
  { line: "downstream", speed: 90, classification: "bus" },
  { line: "upstream", speed: 79, classification: "2-axle truck" },
  { line: "upstream", speed: 59, classification: "bus" },
  { line: "upstream", speed: 78, classification: "car" },
  { line: "downstream", speed: 43, classification: "2-axle truck" },
  { line: "upstream", speed: 47, classification: "car" },
  { line: "upstream", speed: 77, classification: "2-axle truck" },
  { line: "downstream", speed: 98, classification: "2-axle truck" },
  { line: "upstream", speed: 14, classification: "bus" },
];
class SmallVehicles {
  apply(vehicle) {
    return vehicle.classification === "car";
  }
}

class SlowVehicle {
  apply(vehicle) {
    return vehicle.speed > 70;
  }
}

class Line {
  apply(vehicle) {
    return vehicle.line === "upstream";
  }
}

const line = new Line();
const slow = new SlowVehicle();
let result = data;
data.forEach((val) => {
  result = result.filter((item) => line.apply(item) && slow.apply(item));
});

console.log(result);
