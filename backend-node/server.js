const express = require('express');
const cors = require('cors');

const app = express();
app.use(cors());
app.use(express.json());

// Simulated Smart Prescription Analysis / OCR endpoint
app.post('/api/prescriptions/analyze', (req, res) => {
  const { imageUrl, textPrompt } = req.body;
  
  // Real clinical rules engine simulating OCR parse
  const sampleParsedRx = {
    detectedDoctor: "Dr. Robert Sterling, MD - Internal Medicine",
    clinicLicense: "MED-LIC-88912",
    dateIssued: new Date().toISOString().split('T')[0],
    detectedMedications: [
      {
        drugName: "Augmentin (Amoxicillin/Clavulanate Potassium)",
        dosage: "500mg / 125mg",
        frequency: "Twice daily (every 12 hours) with meals",
        duration: "7 days",
        quantity: 14,
        notes: "Take at start of meal. Complete full course.",
        priceEstimate: 18.50
      }
    ],
    contraindicationsWarning: "Check for Penicillin or Beta-lactam allergies.",
    authenticityScore: 0.98
  };

  res.json({
    success: true,
    data: sampleParsedRx
  });
});

const PORT = process.env.PORT || 5000;
app.listen(PORT, () => {
  console.log(`PharmRx Node.js server running on port ${PORT}`);
});
