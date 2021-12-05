// on ready function
$(document).ready(function () {
  "use strict";

  // Preloader
  jQuery(window).on("load", function () {
    jQuery("#status").fadeOut();
    jQuery("#preloader").delay(350).fadeOut("slow");
  });

  //---------------- jQuery SlickNav / Onepage Mobile menu-----------//

  $(".mainmenu").slicknav({
    label: "",
    duration: 700,
    easingOpen: "easeOutBounce",
    prependTo: "#mobileMenu",
    closeOnClick: true,
  });

  //----------- jQuery MeanMenu / Multipage Mobile menu----------//
  $(".mobile-menu nav").meanmenu({
    meanScreenWidth: "992",
    meanMenuContainer: ".mobile-menu",
  });

  // Main Slider Animation

  (function ($) {
    //Function to animate slider captions
    function doAnimations(elems) {
      //Cache the animationend event in a variable
      var animEndEv = "webkitAnimationEnd animationend";

      elems.each(function () {
        var $this = $(this),
          $animationType = $this.data("animation");
        $this.addClass($animationType).one(animEndEv, function () {
          $this.removeClass($animationType);
        });
      });
    }

    //Variables on page load
    var $myCarousel = $("#carousel-example-generic"),
      $firstAnimatingElems = $myCarousel
        .find(".item:first")
        .find("[data-animation ^= 'animated']");

    //Initialize carousel
    $myCarousel.carousel();

    //Animate captions in first slide on page load
    doAnimations($firstAnimatingElems);

    //Pause carousel
    $myCarousel.carousel("pause");

    //Other slides to be animated on carousel slide event
    $myCarousel.on("click slide.bs.carousel", function (e) {
      var $animatingElems = $(e.relatedTarget).find(
        "[data-animation ^= 'animated']"
      );
      doAnimations($animatingElems);
    });
  })(jQuery);
});
