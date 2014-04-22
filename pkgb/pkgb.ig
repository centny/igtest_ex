M pkgb_begin testing packate B begin
//
Y $PKGA_KEY_V
//
HG $URL_ADD_PKGB key=$PKGA_KEY_V val=333 #data=a_pkb
M pkgb_add $(/a_pkb)
Y $(/a_pkb/code)==0
//
HG $URL_GET_PKGB key=$PKGA_KEY_V #data=g_pkb
M pkgb_get $(/g_pkb)
Y $(/g_pkb/code)==0
Y $(/g_pkb/data)==333
//
