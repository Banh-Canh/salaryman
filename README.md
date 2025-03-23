# Salaryman

Build your resume with HTML/CSS and JSON Data

It's a fork from the original project: [Resumme Builder](https://github.com/lucasnevespereira/resumme-builder) by [lucasnevespereira](https://github.com/lucasnevespereira).

## Table of Contents

<img align=left src="public/gopher-salaryman.png" width="170vw" />

- [Introduction](#introduction)
- [Requirements](#requirements)
- [Usage](#local-usage)
- [Templates](#templates)
- [Image](#image)
- [Languages](#languages)
- [Date Formats](#date-formats)
- [License](#license)

## Introduction

Salaryman is a tool that allows you to generate a resume using HTML/CSS templates and JSON data.
It follows the [JSON Resume](https://jsonresume.org/) standard for structuring resume data.

## Requirements

The following package is required:

`google-chrome`, it is used for html to pdf conversion.

## Usage

See [Documentations](docs/salaryman.md).

### Local Usage

Create a JSON file with your resume data. You can specify the file path using the `file` flag.

<i>You can see a json file example in [examples/example.json](examples/example.json)</i>

Run the following command to generate the resume in PDF and HTML formats:

```shell
go run *.go local -f="data/resume.json" -o="cv.pdf"
go run *.go local -f="data/resume.json" -o="cv.html --html"
```

Using Nix:

```shell
nix-shell
salaryman local -f data/resume.json -o cv.pdf
```

The generated resume files (`resume.pdf` and `resume.html`) will be saved in the `output` directory.

### API Usage

To use Salaryman as an API, follow these steps:

Start the server by running the following command:

```
export SALARYMAN_OUTPUTDIR=/tmp
go run . server
```

Using Nix:

```shell
nix-shell
salaryman server
```

Send a POST request to `http://localhost:9000/pdf` with the JSON resume data in the request body.
You can use the example JSON data provided in [examples/example.json](examples/example.json).

In server mode, the html and pdf files will be outputed to the path you set in `SALARYMAN_OUTPUTDIR`.

By default it will output to `./output`

```bash
curl -X POST http://localhost:9000/pdf -H "Content-Type: application/json" -d @examples/example.json
```

The server will generate the resume in PDF format and return it as a response.

e.g example json data request in [examples/example.json](examples/example.json)

## Templates

Salaryman provides the following templates for generating resumes:

- Classic
- Basic
- Simple
- Oldman
- Stackoveoverflow
- Modern

To use a specific template, specify the template name in the JSON resume data:

```json
  "template": "classic"
```

## Image

If you want to include an image in your resume, provide the image URL in the JSON resume data:

```json
 "image": "https://i.imgur.com/tHA5l7T.jpg"
```

Upload your image to a service like [imgur](https://imgur.com/) and copy the direct link.

## Languages

Salaryman supports multiple languages for your resume, allowing you to create your resume in a language that suits
your needs. The default language is English (en), but you can choose to use other supported languages as well.

Currently, the following languages are supported:

- English (en_US)
- French (fr_FR)

This will automatically translate labels such as "Education," "Experiences," and other sections based on the chosen
language.

To set the language for your resume, include the following field in the JSON resume data:
e.g [examples/example.resume.json](examples/example.json)

```json
"lang": "fr_FR"
```

## Date Formats

Salaryman supports the following date formats:

- `2006-01-02` (e.g., "2024-07-09")
- `2006-01` (e.g., "2024-07")
- `January 2 2006` (e.g., "July 9 2024")
- `January 2006` (e.g., "July 2024")
- `2006` (e.g., "2024")

Example of date fields in JSON resume data:

```json
{
  "education": [
    {
      "institution": "University of Example",
      "area": "Computer Science",
      "studyType": "Bachelor",
      "startDate": "2015-09-01",
      "endDate": "2019-06-30"
    }
  ],
  "work": [
    {
      "company": "Example Corp",
      "position": "Software Engineer",
      "startDate": "2020-01",
      "endDate": "2024-07"
    }
  ]
}
```

Feel free to contribute additional templates and features to enhance the Salaryman project!

## License

This project is under [MIT License](LICENSE)
