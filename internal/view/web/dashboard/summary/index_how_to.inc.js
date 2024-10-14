window.alpineSummaryHowToSlider = function () {
  return {
    slidesQty: 4,
    currentSlide: 1,

    get hasNextSlide () {
      return this.currentSlide < this.slidesQty
    },

    get hasPrevSlide () {
      return this.currentSlide > 1
    },

    nextSlide () {
      if (this.hasNextSlide) this.currentSlide++
    },

    prevSlide () {
      if (this.hasPrevSlide) this.currentSlide--
    }
  }
}
