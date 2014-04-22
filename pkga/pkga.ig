M pkga_begin testing packate A begin
//
//
HG $URL_ADD_PKGA key=abc val=123jj #data=a_pka
M pkga_add $(/a_pka)
Y $(/a_pka/code)==0
//
HG $URL_DEL_PKGA key=abc #data=d_pka
M pkga_del $(/d_pka)
Y $(/d_pka/code)==0
//
HG $URL_ADD_PKGA key=pkgb val=123jj #data=a_pka
M pkga_add $(/a_pka)
Y $(/a_pka/code)==0
//
$PKGA_KEY_V=pkgb
