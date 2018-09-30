/**
 * Based on http://snowsyn.net/2016/09/11/creating-shared-libraries-in-go/
 */

#include <iostream>
#include "libimgutil.h"

using namespace std;

int main(int argc, char **argv) {
    GoUint w, h;

    if (argc < 2) {
        cerr << "missing image filename argument" << endl;
        return 1;
    }

    char* path = argv[1];
    char* err = ImgutilGetImageSize(path, &w, &h);

    if (err) {
        cerr << "error: " << err << endl;
        return 1;
    }

    cout << path << " " << w << "x" << h << endl;
    return 0;
}