import './index.css';

const senderTextInput = document.getElementById(
  'sender-text'
) as HTMLInputElement;
const sendButton = document.getElementById('send-button');

sendButton?.addEventListener('click', () => {
  const senderText = senderTextInput?.value;
  if (senderText) {
    console.log(`Sending message: ${senderText}`);
  }
});
