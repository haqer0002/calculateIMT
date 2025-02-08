# IMT Calculator

Simple IMT (Body Mass Index) calculator written in Go that takes height and weight as input and outputs the BMI and its
corresponding category.

## Features
- User-friendly interface for entering height and weight.
- BMI calculation and result classification:
- < 16: Severe body weight deficiency
- < 18.5: Body weight deficiency
- < 25: Normal weight
- < 30: Excess weight
- < 35: First degree of obesity
- < 40: Second degree of obesity
- = 40: Third degree of obesity
- Ability to repeat the calculation.

## Installation

Clone the repository:

```bash
git clone https://github.com/haqer0002/calculateIMT.git
cd calculateIMT
```

### Docker Setup
To run the application with Docker, use the provided `Dockerfile`:
1. Build the Docker image with:

```bash
docker build -t imt-calculator .
```

2. Run the container:
```bash
docker run -it imt-calculator
```

## Usage

You can also run the application directly:
```bash
go run main.go
```

## License
MIT License. See LICENSE for details.