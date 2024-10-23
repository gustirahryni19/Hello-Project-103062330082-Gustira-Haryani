// Aset.cpp
#include "Aset.h"

Aset::Aset(string nama, string jenis, int jumlah, double harga)
    : nama(nama), jenis(jenis), jumlah(jumlah), harga(harga), tahun(0), next(nullptr) {}

void Aset::printInfo() {
    cout << "Nama: " << nama << endl;
    cout << "Tipe: " << jenis << endl;
    cout << "Harga: " << fixed << setprecision(2) << harga << endl;
    cout << "Jumlah: " << jumlah << endl;
}

Aset::~Aset() {}
