" 自動縮排。
 set ai
" "
" 使用自動縮排以後，在貼上剪貼簿的資料時排版可能會亂掉，這時可以手動切換至貼上模式
" :set paste 再進行貼上動作。
"
" " 啟用暗色背景模式。
 set background=dark
"
" " 啟用行游標提示。
 set cursorline
"
" " 文字編碼加入utf8。
 set enc=utf8
"
" " 標記關鍵字。
 set hls
"
" " 只在 Normal 以及 Visual 模式使用滑鼠，也就是取消 Insert 模式的滑鼠，
 set mouse=nv
"
" " 顯示行號。
" set number
"
" " 搜尋不分大小寫。
 set ic
"
" " 自訂縮排(Tab)位元數。
 set tabstop=4
 set shiftwidth=4
 set expandtab
 retab

 syntax on
 set foldmarker={{{,}}}
 set foldmethod=marker
 set foldlevel=1 " 預設全部關閉
 let php_folding = 3 " 這個很重要
 set foldnestmax=3 

