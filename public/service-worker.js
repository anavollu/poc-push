self.addEventListener('push', (evt) => {
    const data = evt.data.json();
    self.registration.showNotification(data.title, {
        body: data.body,
        icon: data.icon,
    });
});

self.addEventListener("notificationclick", function (evt) {
    evt.notification.close();
    evt.waitUntil(
        clients.openWindow(evt.notification.body)
    );
});