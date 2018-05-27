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

  $("[data-ajax]").on("click", function () {
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
});
