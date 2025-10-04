#!/usr/bin/env perl -s
#
# $Id: txt2cc.pl,v 1.1 1998/12/09 09:43:36 nick Exp $

use strict;

my $process = "";

my $msg_text = "";
my $trim     = 0;

my $expand = 0;

my $in_text = 0;

&begin_file();

LINE: while (<STDIN>) {
    chop;

    if (/^\#/) {
        s/^\#//;

        if (/^message/) {
            if ( $in_text == 1 ) {
                $in_text = 0;
                &end_message;
            }
        }
        elsif (/^process (.*)/) {
            if ( $in_text == 1 ) {
                $in_text = 0;
                &end_message;
            }

            $process = $1;
        }
        elsif (/^trim/) {
            $trim = 1;
        }
        else {
            next LINE;
        }
    }
    else {
        if ( $in_text == 0 ) {
            &begin_message;
            $in_text = 1;
        }

        $msg_text .= "$_\\n";
    }
}

&end_file();

###############################################################################

sub begin_file {
    print <<EOF;
/****** GENERATED FILE - DO NOT EDIT! ******/

EOF

    print <<EOF;
#include <fed.hh>
#include <typedefs.hh>

#include "messages.hh"

extern const message_t messages[] = {
   {0,"<<NULL MESSAGE>>"},
EOF
}

sub begin_message {
    $expand = 0;
}

sub end_message {
    &vet_message;

    $msg_text =~ s/[\{\}]//g;

    if ( $process eq "L" ) {
        if ( $msg_text ne "" && $msg_text ne "\\n" ) {
            $msg_text =~ s/\"/\\\"/g;

            my @msg_text = split( /\\n/, $msg_text );

            &write_text( $msg_text[0] . "\\n" );
            &write_text($msg_text);
        }

        $msg_text = "";
    }
    elsif ( $process eq "O" ) {
        if ( $msg_text ne "" && $msg_text ne "\\n" ) {
            $msg_text =~ s/\"/\\\"/g;

            my @msg_text = split( /\\n/, $msg_text );

            &write_text( $msg_text[0] . "\\n" );
            &write_text( $msg_text[1] . "\\n" );
        }

        $msg_text = "";
    }
    else {
        ## Remove "\n".
        #substr($msg_text,-2) = '' unless length($msg_text) < 2;

        if ($trim) {

            # Remove "\n".
            substr( $msg_text, -2 ) = '' unless length($msg_text) < 2;
        }
    }

    if ( $msg_text ne "" ) {
        $msg_text =~ s/\"/\\\"/g;

        my $full = $msg_text;
        &write_text($full);

        $msg_text = "";
    }

    $trim = 0;
}

sub end_file {
    &end_message();

    print "};\n";
}

sub vet_message {
    my $vet_text = $msg_text;
    $vet_text =~ s/%%//g;

    if ( $vet_text =~ /[^{]%/ ) {
        print STDERR "\n$msg_text\n\n";
        print STDERR "**** Unbracketed conversion specification ****\n\n";

        #die "**** Unbracketed conversion specification ****\n\n";
    }

    while ( $vet_text =~ /\{([^\}]*)\}/ ) {
        my $spec = $1;

        if ( $spec =~ /^%[1-9]\$/ ) {
            $expand = 1;
        }
        elsif ( $spec =~ /^%[-\'\+0]*\d*(\.\d+)?[hl]?[cdfisuX]$/ ) {
            $expand = 1;
        }
        else {
            print STDERR "\n$msg_text\n\n";
            die "**** Bad conversion specification: {$spec} ****\n\n";
        }

        $vet_text =~ s/\{//;
        $vet_text =~ s/\}//;
    }
}

sub write_text {
    my ($text) = @_;

    print "   {${expand},\"${text}\"},\n";
}

###############################################################################
