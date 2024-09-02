export function initHelpers () {
  function copyToClipboard (textToCopy) {
    const successMessage = 'Text copied'
    const errorMessage = 'Error copying text'

    /* First, try the modern approach */
    if (navigator.clipboard && window.isSecureContext) {
      return navigator.clipboard
        .writeText(textToCopy)
        .then(() => {
          toaster.success(successMessage)
        })
        .catch((err) => {
          console.error(errorMessage, err)
          toaster.error(errorMessage)
        })
    }

    /* Fallback: use execCommand("copy") method */
    const textArea = document.createElement('textarea')
    textArea.value = textToCopy

    textArea.style.position = 'fixed'
    textArea.style.left = '-9999px'
    textArea.style.top = '-9999px'
    document.body.appendChild(textArea)
    textArea.focus()
    textArea.select()

    try {
      const successful = document.execCommand('copy')
      if (successful) {
        toaster.success(successMessage)
      } else {
        console.error('Fallback', errorMessage)
        toaster.error(errorMessage)
      }
    } catch (err) {
      console.error('Fallback', errorMessage, err)
      toaster.error(errorMessage)
    }

    document.body.removeChild(textArea)
  }

  function textareaAutoGrow (element) {
    if (!element?.style) return
    element.style.height = '1px'
    element.style.height = (element.scrollHeight + 2) + 'px'
  }

  function formatJson (inJSON) {
    if (typeof inJSON !== 'string' || inJSON === '') return inJSON
    try {
      return JSON.stringify(JSON.parse(inJSON), null, 2)
    } catch (e) {
      console.error(e)
      return inJSON
    }
  }

  window.copyToClipboard = copyToClipboard
  window.textareaAutoGrow = textareaAutoGrow
  window.formatJson = formatJson
}
