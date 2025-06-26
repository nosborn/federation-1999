#!/bin/bash
set -euo pipefail

cd "$(git rev-parse --show-toplevel)/web"
if [[ ! -d archive ]]; then
  mkdir archive
fi
cd archive

curl -fLOsS https://web.archive.org/web/20141223055443im_/http://www.ibgames.net/favicon.ico

curl -fOsS https://web.archive.org/web/19990423173726id_/http://www.ibgames.net/index.html
# gsed -i 's|images/buttons/start\.gif|https://web.archive.org/web/20000928105319im_/http://www.ibgames.net/&|' index.html
# gsed -i 's|images/fedpix/fedmain3\.gif|https://web.archive.org/web/20000425143056im_/http://www.ibgames.net/&|' index.html
# gsed -i 's|images/ibgames\.gif|https://web.archive.org/web/20000427033628id_/http://www.ibgames.net/&|' index.html
# gsed -i 's|images/whitespace\.gif|https://web.archive.org/web/20000429052830id_/http://www.ibgames.net/&|' index.html

curl -fsS -o login.html https://web.archive.org/web/19990202073011id_/http://www.ibgames.net/account/

curl -fOsS https://web.archive.org/web/19990423181320id_/http://www.ibgames.net/main.html
gsed -i 's|\.\./\.\./\(icons/rsacirated\.gif\)|https://web.archive.org/web/20000928015435im_/http://www.ibgames.net/\1|' main.html
gsed -i 's|/icons/ie4get_animated\.gif|https://web.archive.org/web/19980526143820im_/http://www.ibgames.net&|' main.html
gsed -i 's|/icons/now_anim_button\.gif|https://web.archive.org/web/19980526143841im_/http://www.ibgames.net&|' main.html
gsed -i 's|age/index\.html|https://web.archive.org/web/19990418023805id_/http://www.ibgames.net/&|' main.html
gsed -i 's|federation/index\.html|https://web.archive.org/web/19990423001538id_/http://www.ibgames.net/&|' main.html
gsed -i 's|help/index\.html|https://web.archive.org/web/19990423052928id_/http://www.ibgames.net/&|' main.html
gsed -i 's|help/secureserver\.html|https://web.archive.org/web/19990224163332id_/http://www.ibgames.net/&|' main.html
gsed -i 's|https://www.ibgames.net||' main.html
gsed -i 's|ibinfo/index\.html|https://web.archive.org/web/19990423104157id_/http://www.ibgames.net/&|' main.html
gsed -i 's|ibinfo/whatsnew\.html|https://web.archive.org/web/19990423173214id_/http://www.ibgames.net/&|' main.html
# gsed -i 's|images/aollogo\.gif|https://web.archive.org/web/20000427041921id_/http://www.ibgames.net/&|' main.html
# gsed -i 's|images/aolmemb\.gif|https://web.archive.org/web/20000928082352id_/http://www.ibgames.net/&|' main.html
# gsed -i 's|images/iblogosmall\.gif|https://web.archive.org/web/20000422035523id_/http://www.ibgames.net/&|' main.html
gsed -i 's|shop/index\.html|https://web.archive.org/web/19990423184905id_/http://www.ibgames.net/&|' main.html

mkdir -p account
cd account

curl -fsS -o signup.html https://web.archive.org/web/19990418015935id_/https://www.ibgames.net/account/signup
gsed -E -i 's/\<80\>/72/g' signup.html
gsed -i 's/gettheir/get their/' signup.html
gsed -i 's/thenplease/then please/' signup.html
gsed -i 's/uncheckthe/uncheck the/' signup.html
gsed -i 's|/help/index\.html|https://web.archive.org/web/19990423052928id_/http://www.ibgames.net&|' signup.html

cd ..

mkdir -p federation
cd federation

curl -fOsS https://web.archive.org/web/19990423001538id_/http://www.ibgames.net/federation/index.html
gsed -i 's|https://www.ibgames.net||' index.html

mkdir -p fedinfo
cd fedinfo

curl -fOsS https://web.archive.org/web/19990422213240id_/http://www.ibgames.net/federation/fedinfo/index.html
gsed -i 's|https://www.ibgames.net||' index.html

cd ..
cd ..

mkdir -p images
cd images

curl -fLOsS https://web.archive.org/web/20000427041921id_/http://www.ibgames.net/images/aollogo.gif
curl -fLOsS https://web.archive.org/web/20000928082352id_/http://www.ibgames.net/images/aolmemb.gif
curl -fLOsS https://web.archive.org/web/20000507060520id_/http://www.ibgames.net/images/blueball.gif
curl -fLOsS https://web.archive.org/web/20000929120538id_/http://www.ibgames.net/images/bullthumb.jpg
curl -fLOsS https://web.archive.org/web/20000929071839id_/http://www.ibgames.net/images/computerthumb.jpg
curl -fLOsS https://web.archive.org/web/20000427033628id_/http://www.ibgames.net/images/ibgames.gif
curl -fLOsS https://web.archive.org/web/20000422035523id_/http://www.ibgames.net/images/iblogosmall.gif
curl -fLOsS https://web.archive.org/web/20000429052830id_/http://www.ibgames.net/images/whitespace.gif

mkdir -p buttons
cd buttons

curl -fLOsS https://web.archive.org/web/19990418015935im_/http://www.ibgames.net/images/buttons/click.gif
curl -fLOsS https://web.archive.org/web/19990418015935im_/http://www.ibgames.net/images/buttons/signup.gif
curl -fLOsS https://web.archive.org/web/19990418015935im_/http://www.ibgames.net/images/buttons/small05.gif
curl -fLOsS https://web.archive.org/web/19990418015935im_/http://www.ibgames.net/images/buttons/small19.gif
curl -fLOsS https://web.archive.org/web/20000928105319im_/http://www.ibgames.net/images/buttons/start.gif

cd ..

mkdir -p fedpix
cd fedpix

curl -fLOsS https://web.archive.org/web/19990423001538id_/http://www.ibgames.net/images/fedpix/fedlogo.gif
curl -fLOsS https://web.archive.org/web/20000425143056im_/http://www.ibgames.net/images/fedpix/fedmain3.gif

cd ..
cd ..

mkdir -p stuff
cd stuff

curl -fOsS https://web.archive.org/web/19990202104000id_/http://www.ibgames.net/stuff/computer.html
# gsed -i 's|\.\./\(images/iblogosmall\.gif\)|https://web.archive.org/web/20000422035523id_/http://www.ibgames.net/&|' computer.html
# gsed -i 's|\.\./\(images/whitespace\.gif\)|https://web.archive.org/web/20000429052830id_/http://www.ibgames.net/\1|' computer.html
gsed -i 's|https://www.ibgames.net||' computer.html

curl -fOsS https://web.archive.org/web/19990202150148id_/http://www.ibgames.net/stuff/index.html
# gsed -i 's|\.\./\(images/blueball\.gif\)|https://web.archive.org/web/20000507060520id_/http://www.ibgames.net/&|' index.html
# gsed -i 's|\.\./\(images/iblogosmall\.gif\)|https://web.archive.org/web/20000422035523id_/http://www.ibgames.net/&|' index.html
# gsed -i 's|\.\./\(images/whitespace\.gif\)|https://web.archive.org/web/20000429052830id_/http://www.ibgames.net/\1|' index.html
gsed -i 's|https://www.ibgames.net||' index.html

curl -fOsS https://web.archive.org/web/19990202162036id_/http://www.ibgames.net/stuff/officeview.html
# gsed -i 's|\.\./\(images/iblogosmall\.gif\)|https://web.archive.org/web/20000422035523id_/http://www.ibgames.net/&|' officeview.html
# gsed -i 's|\.\./\(images/whitespace\.gif\)|https://web.archive.org/web/20000429052830id_/http://www.ibgames.net/\1|' officeview.html
gsed -i 's|https://www.ibgames.net||' officeview.html

curl -fOsS https://web.archive.org/web/19990202174102id_/http://www.ibgames.net/stuff/photos.html
# gsed -i 's|\.\./\(images/bullthumb\.gif\)|https://web.archive.org/web/20000929120538id_/http://www.ibgames.net/&|' photos.html
# gsed -i 's|\.\./\(images/computerthumb\.gif\)|https://web.archive.org/web/20000929071839id_/http://www.ibgames.net/&|' photos.html
# gsed -i 's|\.\./\(images/iblogosmall\.gif\)|https://web.archive.org/web/20000422035523id_/http://www.ibgames.net/&|' photos.html
# gsed -i 's|\.\./\(images/whitespace\.gif\)|https://web.archive.org/web/20000429052830id_/http://www.ibgames.net/\1|' photos.html
gsed -i 's|bull\.html|https://web.archive.org/web/19990202100129id_/http://www.ibgames.net/stuff/&|' photos.html
gsed -i 's|https://www.ibgames.net||' photos.html

cd ..
