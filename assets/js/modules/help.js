// Callback used when the "Yes" button of a confirm modal is clicked.
function callAjax(event) {
  var $link = $(event.relatedTarget);

  $link.addClass("disabled");
  $link.find("[data-spinner]").removeClass("d-none");

  $.ajax({
    url: $link.attr("href")
  }).always(function () {
    $link.removeClass("disabled");
    $link.find("[data-spinner]").addClass("d-none");
  });
}

$(function () {
  var $linkContainer = $("[data-link=container]")

  // Ajax call on recalbox-support.sh
  $("[data-recalbox-support]").on("click", function () {
    var $this = $(this);

    $linkContainer.addClass("d-none");
    $this.addClass("disabled");
    $this.find("[data-spinner]").removeClass("d-none");

    $.ajax({
      url: $this.attr("href"),
      dataType: "json"
    }).done(function (data) {
      var $a = $linkContainer.find("a")
      $a.attr("href", data.url);
      $a.html("<strong> " + data.url + "</strong>")
      $linkContainer.removeClass("d-none");
    }).always(function () {
      $this.removeClass("disabled");
      $this.find("[data-spinner]").addClass("d-none");
    });

    return false;
  });

  // Get ES status
  var checkInterval = 3 * 1000;
  var $btnStart = $("[data-es=start]");
  var $btnStop = $("[data-es=stop]");
  var checkEsStatus = function () {
    $.ajax({
      url: $("[data-es-status-url]").data("es-status-url"),
      dataType: "json"
    }).done(function (data) {
      if (data.running) {
        $btnStart.addClass("d-none");
        $btnStop.removeClass("d-none");
      } else {
        $btnStart.removeClass("d-none");
        $btnStop.addClass("d-none");
      }
    }).always(function () {
      setTimeout(checkEsStatus, checkInterval);
    });
  }

  checkEsStatus();

  // Execute ES Actions
  $("[data-es]").on("click", function () {
    var $this = $(this);

    $this.addClass("disabled");
    $this.find("[data-spinner]").removeClass("d-none");

    $.ajax({
      url: $this.attr("href"),
      dataType: "json"
    }).always(function () {
      $this.removeClass("disabled");
      $this.find("[data-spinner]").addClass("d-none");
    });

    return false;
  });
});
