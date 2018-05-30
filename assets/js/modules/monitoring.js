$(function () {
  // Get monitoring data
  var checkInterval = 3 * 1000;
  var getMonitoringData = function () {
    $.ajax({
      url: $("[data-monitoring-url]").data("monitoring-url"),
      dataType: "json"
    }).done(function (data) {
      console.log(data);
    }).always(function () {
      setTimeout(getMonitoringData, checkInterval);
    });
  }

  getMonitoringData();
});
