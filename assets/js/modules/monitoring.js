$(function () {
  // Get monitoring data
  var checkInterval = 3 * 1000;
  var $memoryAvailable = $("[data-memory=available]");
  var $memoryUsed = $("[data-memory=used]");
  var $memoryTotal = $("[data-memory=total]");
  var $temperature = $("[data-temperature]");
  var getMonitoringData = function () {
    $.ajax({
      url: $("[data-monitoring-url]").data("monitoring-url"),
      dataType: "json"
    }).done(function (data) {
      // CPU update
      for (var i in data.cpu) {
        var cpu = data.cpu[i];

        $("[data-cpu=" + cpu.Number + "]").text(cpu.Value + "%").css("width", cpu.Value + "%");
      }

      // Memory update
      $memoryAvailable.text(" " + data.memory.available);
      $memoryUsed.text(" " + data.memory.usedPercent.toFixed(2) + "%").css("width", data.memory.usedPercent + "%");
      $memoryTotal.text(" " + data.memory.total);

      // Temperature update
      if (0 < $temperature.length) {
        $temperature.removeClass("orange green");
        $temperature.removeClass(function (index, className) {
          return (className.match (/\c\d+/g) || []).join(' ');
        });
        $temperature.addClass("c" + data.temperature.CurrentPercent + " " + data.temperature.Color);
        $temperature.find("span").attr("title", data.temperature.Current + "°c / " + data.temperature.Max + "°c");
        $temperature.find("span").text(data.temperature.Current + "°c");
      }
    }).always(function () {
      setTimeout(getMonitoringData, checkInterval);
    });
  }

  getMonitoringData();
});
