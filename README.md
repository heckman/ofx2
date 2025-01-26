# ofx-txns

This is a bash script for extracting data from OFX files.

I wrote it to generate CSV files of the transactions in the OFX files
for my various bank and credit-card accounts.

It depends on the utility _ofx2xml_,
the source for which can be found in the `src/ofx2xml` directory.

The _ofx2xml_ utility remarshals OFX files to version-2 OFX files,
which are vavid XML and therefore easier to parse.
The utility also interprets
`<INTU.BID>` tags as `<FI><FID>` tag pairs.

## Usage

To perform its primary function run `ofx-txns csv -w <ofx-file>`.
This invocation will create a csv file
of the transactions found in the ofx-file,
the name of which will be generated from
the financial instutution,
the account number,
and the date range of the transactions.

The script includes other options,
and a few other tangentally-related functions.
For complete usage information,
run `ofx-txns --help`,
or examine [the source code](./ofx-txns),
where the usage appears near the top.

## Installation

I wrote this for _MacOS_,
but it should work on _Unix_-like systems.

The script _ofx-txns_ is in the root of this repository.
There is a second script _ofx2csv_
which calls _ofx-txns_ with some preset options.

In addition to the included _ofx2xml_ utility,
the script requires both _jq_ and _xq_ to be installed.
They are both avaiable via _homebrew_.

The _ofx2xml_ utility is written in _go_.
I needs to be compiled, and the resulting binary
should be put on the system's path.

## Configuration

Some configuration is required to customize the script
for use with OFX files from specific financial institutions.

In the _ofx-txns_ script, immediately following the usage information,
a varable `_fidir_` is defined.
It maps `INTU.BID` values to the names of financial institutions.
Add the financial institutions you use to this variable.
Values for `INTU.BID` for institutions in North America
can be found [from intuit.com](https://ofx-prod-filist.intuit.com/qm2400/data/fidir.txt).

## Development

This project makes use of the _direnv_ utility
to add local directories to the environment's PATH.
It is usefull for putting _ofx2xml_ on the path
before it has been installed.

The `dev-bin` directory contains scripts to
build and remove the compiled _ofx2xml_ binary.

## Rationale

To track my finances,
I used to download my transactions as CSV files
and then import them into a spreadsheet program.

The problem with this approach is that there is
no consistency between the format of the CSV files
from the various financial institutions.

The _ofx-txns_ script attempts to address this problem.

Most financial institutions
will provide transaction history in OFX format,
although in my experience it's usually an older non-XML version of OFX.
Furthermore, in my experience,
they don't include a `<FI>` tag to identify the financial institution,
but instead use a (not-to-spec) `<INTU.BID>` tag.
