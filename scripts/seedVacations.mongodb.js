const db = db.getSiblingDB("vacationsApi");

db.vacations.drop();

db.vacations.insertMany([
  {
    city: "Varenna",
    country: "Italy",
    food: "fish",
    wine: "",
    pointsOfInterest: [{ name: "", description: "" }],
  },
  {
    city: "Riomaggiore",
    country: "Italy",
    food: "anchovies and pesto",
    wine: "",
    pointsOfInterest: [{ name: "", description: "" }],
  },
  {
    city: "Vernazza",
    country: "Italy",
    food: "anchovies, pesto, gelato, focaccia",
    wine: "",
    pointsOfInterest: [{ name: "", description: "" }],
  },
  {
    city: "Corniglia",
    country: "Italy",
    food: "anchovies, pesto, gelato, focaccia",
    wine: "",
    pointsOfInterest: [{ name: "", description: "" }],
  },
  {
    city: "Firenze",
    country: "Italy",
    food: "Bistecca a'la Fiurentino",
    wine: "",
    pointsOfInterest: [{ name: "", description: "" }],
  },
  {
    city: "Montepulciano",
    country: "Italy",
    food: "",
    wine: "Montepulciano",
    pointsOfInterest: [{ name: "", description: "" }],
  },
  {
    city: "Montalcino",
    country: "Italy",
    food: "",
    wine: "Brunello di Montalcino",
    pointsOfInterest: [{ name: "", description: "" }],
  },
  {
    city: "Pula",
    country: "Croatia",
    food: "Konoba",
    wine: "Malvasia",
    pointsOfInterest: [
      {
        name: "Pula Colosseum",
        description: "Best preserved Roman colosseum from the venitian empire",
      },
    ],
  },
]);

print(
  "Seeded " +
    db.vacations.countDocuments() +
    " vacations into vacationsApi.vacations",
);
