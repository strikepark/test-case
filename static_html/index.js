$(document).ready(function() {
    var apiUrl = 'planadotest.heroku.com/api/orders';

    $.get(apiUrl, function(data) {
        var orders = JSON.parse(data);
        var $app = $('#app');

        $.each(orders, function(orderIndex, order) {
            $app.append('<ul class="order-' + orderIndex + '" style="margin: 20px;"></ul>')

            var $ul = $('.order' + orderIndex);
            $.each(order, function(key, value) {
                $ul.append('<b>' + key + ':</b> ' + value);
            }
        });
    })
});