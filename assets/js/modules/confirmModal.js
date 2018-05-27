$(function () {
  $("#confirmModal").on("show.bs.modal", function (event) {
    var $link = $(event.relatedTarget);
    var $modal = $(this);
    var content = $link.data("message");
    var callback = $link.data("callback");
    var $yesBtn = $modal.find("[data-href]");

    $yesBtn.attr("href", $link.attr("href"));
    $yesBtn.off("click");
    $modal.find(".modal-body").html(
      content.replace("%s", "<strong>" + $link.data("message-arg") + "</strong>")
    );

    if ("" !== callback) {
      $yesBtn.on("click", function () {
        $modal.modal("hide");

        window[callback](event);

        return false;
      });
    }
  });
});
