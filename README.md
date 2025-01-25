# ofx

This is a bash script for extracting data from OFX files.

I wrote it to generate CSV files of the transactions in the OFX files
for my various bank and credit-card accounts.

It depends on the utility, _ofx2xml_,
the source for which can be found in the `src/ofx2xml` directory.

The _ofx2xml_ utility remarshals OFX files to version-2 OFX files,
which are vavid XML and easier to parse.
Furthermore, the utility interprets
`<INTU.BID>` tags as a `<FI><FID>` tag pair.

## Usage

For usage information, run `ofx --help`,
or examine [the source code](./ofx),
where the usage appears near the top.

## Installation

I wrote this for _MacOS_,
but it should work on _Unix_-like systems.

The script _ofx_ is in the root of this repository.
There is a second script _ofx2csv_
which calls _ofx_ with some preset options.

In addition to the included _ofx2xml_ utility,
the script requires both _jq_ and _xq_ to be installed.
They are both avaiable via _homebrew_.

The _ofx2xml_ utility is written in _go_.
I needs to be compiled, and the resulting binary
should be put on the system's path.

## Configuration

Some configuration is required to customize the script
for use with OFX files from your financial institutions.

In the _ofx_ script, immediately following the usage information,
a varable `_fidir_` is set containing the definition of a json-object
which maps `INTU.BID` values to names of financial institutions.
You will need to add your own entries to this object.

## Development

This project makes use of the _direnv_ utility
add local directories to the environment's PATH.

The `dev-bin` directory contains scripts to
build and remove the compiled _ofx2xml_ binary.
It also contains a script that calls _go_ to
run the source-code of _ofx2xml_ directly.

## Rationale

To track my finances,
I used to download my transactions as CSV files
and then import them into a spreadsheet program.

The problem with this approach is that there is
no consistency between the format of the CSV files
from the various financial institutions.

The _ofx_ script attempts to address this problem.

The banks that I use consistenly
(mostly) support the OFX specification.
The OFX specification doesn not include an `<INTU.BID>` tag,
but it is consistently included by my financial institutions,
none of which include the standard's <FI> tag.
