export const genericSlider = {
  name: 'genericSlider',
  fn: (slidesQty = 0) => ({
    currentSlide: slidesQty > 0 ? 1 : 0,
    get hasNextSlide () {
      return this.currentSlide < slidesQty
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
  })
}
