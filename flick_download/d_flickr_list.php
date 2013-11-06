<?php

require_once( "phpFlickr.php" );

$o = new phpFlickr( '6791ccf468e1c2276c1ba1e0c41683a4' );
$d = $o->photos_search( array(
                                ##'tags' => '' ,
                                'camera' => 'htc one',
                                'extras' => 'url_o',
                                'page' => 1 ,
                                'per_page' => 5 
                        )   
        );  
# print_r( $d  );

$totalfile=$d['total'];
function get_by_page ($page , $per_page) {
	$o = new phpFlickr( '6791ccf468e1c2276c1ba1e0c41683a4' );
	$d = $o->photos_search( array(
					'tags' => '' ,
					'camera' => 'htc one',
                    'content_type' => 1 ,
                    'sort' => 'date-posted-asc' ,
					'extras' => 'url_l,url_o',
					'page' => $page ,
					'per_page' => $per_page
				)   
		);
    for($loop=0 ; $loop <=499 ; $loop++ ){
            $photodown = $d['photo'][$loop]['url_l'] ;
        if (!empty($photodown)){
            print_r( "Page: $page , Count: $loop = ".$photodown."\n" );
        }
    }
} 

$page =1 ;
$per_page = 500;
while ($totalfile >= 0){
	echo "page : ".$page." \n";
	get_by_page($page,$per_page);
	$page++;
	$totalfile -= $per_page ;
}


?>
