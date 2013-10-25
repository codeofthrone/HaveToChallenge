#!/usr/bin/perl -w

use strict;
use warnings;
use CGI;
use CGI::Carp qw(fatalsToBrowser);
use File::Basename;
use URI::Escape;
use File::Copy;


our $picdir;

if(!$ARGV[0]){
	print "Please input picture directory!\n";
}else{

	$picdir=$ARGV[0];
	#print $picdir;
	&OpenDir;

}


sub OpenDir {
    opendir ( DIR, $picdir ) or die $!;

    while (my @file = grep {$_ ne '.' && $_ ne '..'} readdir(DIR)){
		foreach my $i (@file) {
			my $OriginalName = $i;
			#print  $i."\n";
			$i =~ s/\.[^.]+$//;
			#print  $i."\n";
			my @filename = split ("_",$i);
			my $array_size = $#filename+1;
			my $i ;
			my $newname =$filename[0];
			for ( $i=1 ; $i<$array_size ; $i++){
				$newname = $newname."_".&SubTagName($filename[$i]); 
				#print "New Name ==> $newname , Sub Tag Name ==> " . &SubTagName($filename[$i]) . "\n";
			}
			$newname=$newname.".jpg";
	    	#	print "array value:$newname\n" ; 
			#print "array size = $array_size \n";
			print "old name => $picdir$OriginalName newname => $picdir$newname  \n";
			system ("mv $picdir$OriginalName $picdir$newname");
		}   
    }
    closedir (DIR);
}

sub SubTagName{
    my($tagid ) = @_;
	
	#print "Tag ID : $tagid \n";

    open(open_file,"mapping_tagid_tagname.txt") or die "open file error" ;
    while(<open_file>){
        chomp ;
		if ($_ =~ $tagid ){
			my $r = $_;
			my @tagname = split ("\t",$r);
        	return $tagname[1];
		}
    }
    close(open_file);
}


#print &SubTagName(108);
