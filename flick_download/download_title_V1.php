#!/usr/bin/php


<?php

require_once( "phpFlickr.php" );

$filepath = $argv[1];

function photo_list ($tags , $page, $total_download){
    
    print("Page ".$page." start \n\n");
    $o = new phpFlickr( 'b05f1fec64df5011d8cc8133fb216a59' );
    $d = $o->photos_search( array(
                'text' => $tags ,
                'content_type' => 1 ,
                'sort' => 'date-posted-asc' ,
                'extras' => 'url_o,url_l,url_b,url_c',
                'page' => $page ,
                'per_page' =>  $total_download
                )   
            );
    for($index = 0 ; $index <= $total_download ; $index++){
        $dir = str_replace(" ","_",$tags);
        print("Count ".$index."\n");
        if(!empty( $d['photo'][$index]['url_o'] )){
            print_r( $d['photo'][$index]['url_o'] ."\n");
            $downfile =  $d['photo'][$index]['url_o'] ;
            system( "wget $downfile -P $dir");
        }elseif(!empty( $d['photo'][$index]['url_l'] )){
            print_r( $d['photo'][$index]['url_l'] ."\n");
            $downfile =  $d['photo'][$index]['url_l'] ;
            $path = `pwd`;
            system( "wget  $downfile -P $dir");
        #system( "wget $downfile ");
        }elseif(!empty( $d['photo'][$index]['url_b'] )){
            print_r( $d['photo'][$index]['url_b'] ."\n");
            $downfile =  $d['photo'][$index]['url_b'] ;
            system( "wget $downfile ");
        }elseif(!empty( $d['photo'][$index]['url_c'] )){
            print_r( $d['photo'][$index]['url_c'] ."\n");
            $downfile =  $d['photo'][$index]['url_c'] ;
            system( "wget $downfile ");

        }else{
            print ("orignal and large size is not available \n");
        }
    }

}


function total_photo ($tags  ){
    $o = new phpFlickr( 'b05f1fec64df5011d8cc8133fb216a59' );
    $d = $o->photos_search( array(
        'text' => $tags ,
        'content_type' => 1 ,
        'sort' => 'date-posted-asc' ,
        'extras' => 'url_o,url_l,url_b,url_c',
        'page' => 1 ,
        'per_page' => 500 
        )   
    );
    print_r ( "total page :".$d['pages']." " );
    $total_page = $d['pages'] ;
    print_r ( "total photo :".$d['total']."\n" );
    $total_photo = $d['total'] ;
    $dir = str_replace(" ","_",$tags);
    system (" mkdir $dir");
    for($page_index = 1 ; $page_index <= $total_page ; $page_index++){
        $per_page = $total_photo - (500 * $page_index);
        if($per_page >= 500){
            $per_page = 500;
        }   
        photo_list ($tags , $page_index, $per_page) ;
    }
}


$handle = fopen($filepath, "r");
if ($handle) {
    while (($line = fgets($handle)) !== false) {
        // process the line read.
        total_photo ($line  ) ;
    }
} else {
    // error opening the file.
    print "error openfile";
}




?>
