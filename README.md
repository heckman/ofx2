# ofx2

**_ofx2_** is a bash script for extracting data from _OFX_ files.

I wrote it to generate _CSV_ files from the transactions in the _OFX_ files
that I download for my bank accounts and credit-cards.

It depends on the utility _**ofx2xml**_,
the source for which can be found in the `src/ofx2xml` directory.

The _**ofx2xml**_ utility remarshals _OFX_ files,
ensuring they are at least version-2 _OFX_ files,
which are valid XML and thereby easier to parse.
The utility also interprets `<INTU.BID>` tags as `<FI><FID>` tag pairs,
as my banks tend to include the former, but not the latter, in their _OFX_ files.

## Usage

To perform its primary function, run `ofx2 csv -w <ofx-file>`.
This invocation will create a _CSV_ file of the transactions found in the _OFX_ file,
the name of which will be generated from
the financial instutution,
the account number,
and the date range of the transactions.
I've inlcuded a helper-script **_ofx2csv_** which calls **_ofx2_** in this way.

The script includes other options,
and a few other tangentally-related operations.
For complete usage information,
run `ofx2 --help`.

## Installation

I wrote this for _MacOS_,
but it should work on _posix_-like systems.

The script _**ofx2**_
and its helper script _**ofx2csv**_
are both in the root of this repository.

The _**ofx2xml**_ utility is written in _go_.
I needs to be compiled.

The scripts require _**jq**_ and **_xq_** to be installed.
To generate hashes of transactions, _**mlr**_ is also required.
(It's not needed if hashing is never used.)
All three utilities are avaiable via _homebrew_.

The script `dev-bin/install`
builds _**ofx2xml**_, and copies the compiled binary,
along with _**ofx2**_ ann _**ofx2csv**_,
to `$HOME/bin`.
(At least on my machine. YMMV. Don't use it blindly.)

## Configuration

Some configuration is required to customize the script
for use with OFX files from specific financial institutions.

In the _ofx-txns_ script, immediately following the usage information,
you'll find the definition of a varable `jq_fidir`.
It contains a json object
that maps `INTU.BID` values
to the names of financial institutions.
Values for `INTU.BID` for institutions in North America
can be found from [intuit.com's list](https://ofx-prod-filist.intuit.com/qm2400/data/fidir.txt).
Add the ones you use to this variable.

## Development

This project makes use of the _direnv_ utility
to add local directories to the environment's PATH.

I find it useful but its use is not necessary.

_**ofx2xml**_ is the first thing I've written in _Go_,
and I leant heavilly on AI to do it,
as I really don't know it at all.
Any tips there would be particularely welcome.

## Motivation

In order to track my finances,
I used to download my transactions as CSV files
and then import them into a spreadsheet.

The problem with this approach is that there is
no consistency between the format of the CSV files
from my various financial institutions.

The _ofx2_ script attempts to address this problem
by using the standardized _OFX_ format to generate
a consistent format for the CSV files.
