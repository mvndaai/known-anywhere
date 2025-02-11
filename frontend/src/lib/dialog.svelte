<script>
  export let open = false;
  let dialog;

  // Function to handle keydown event
  const handleKeydown = (event) => {
    if (event.key === 'Escape' && dialog.open) {
      closeDialog();
    }
  };

  // Function to open the dialog
  const showModal = () => {
    dialog.showModal();
    window.addEventListener('keydown', handleKeydown);
    // Set focus to another element inside the dialog
    setTimeout(() => {
      const focusElement = dialog.querySelector('button:not(.btn-dialog-close)');
      if (focusElement) {
        focusElement.focus();
      }
    }, 0);
  };

  // Function to close the dialog
  const closeDialog = () => {
    dialog.close();
    window.removeEventListener('keydown', handleKeydown);
    open = false; // Reset the open variable
  };

  // Watch the `open` variable and open/close the dialog accordingly
  $: if (open) {
    showModal();
  } else if (dialog && dialog.open) {
    closeDialog();
  }
</script>

<dialog bind:this={dialog}>
  <button class='btn-dialog-close' onclick={closeDialog}>X</button>
  <slot>Missing dialog children</slot>
</dialog>

<style>
  dialog::backdrop {
    background-color: rgba(0,0,0,0.5);
  }
  .btn-dialog-close {
    position: absolute;
    top: 10px;
    right: 10px;
    background: none;
    border: none;
    color: black;
    font-size: 1.2em;
    cursor: pointer;
    user-select: none;
  }
  .btn-dialog-close:focus {
    outline: none;
  }
  .btn-dialog-close:hover {
    color: red;
  }
</style>
