# Maintainer: Vlad Zagorodniy <vladzzag@gmail.com>
pkgname=update-notifier
pkgver=1.0
pkgrel=1
pkgdesc="PackageKit update notifier"
arch=('any')
url="https://github.com/zzag/update-notifier"
license=('MIT')
depends=(packagekit python-dbus)
source=(refresh-packagekit-cache.hook
        refresh-packagekit-cache.py
        update-notifier.service
        update-notifier.timer)
sha256sums=('1892a5d9897f4247209d8c933f64019005e41f5948d49af0c01fd013eab418f0'
            'e7c93906ac79b71c52331da4ba1c57ee7de8dc9861b18ab1430c4bf7d56abe8f'
            '303e644d411c6a6b7861d627f7761cd2082049e15fc4f30a6e1622ceaed5638f'
            '917de336b5fe7d19d3854364212df005842cb1bc8203df914e459e12006f91e0')

package() {
    install -Dm644 "$srcdir/refresh-packagekit-cache.hook" "$pkgdir/usr/share/libalpm/hooks/refresh-packagekit-cache.hook"
    install -Dm644 "$srcdir/refresh-packagekit-cache.py" "$pkgdir/usr/lib/update-notifier/refresh-packagekit-cache.py"

    install -Dm644 "$srcdir/update-notifier.service" "$pkgdir/usr/lib/systemd/system/update-notifier.service"
    install -Dm644 "$srcdir/update-notifier.timer" "$pkgdir/usr/lib/systemd/system/update-notifier.timer"
}

install=$pkgname.install
