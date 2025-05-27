import { writable, get } from "svelte/store";

const defaultVolume = 0.1;
const audioVolumeStore = writable(defaultVolume);

function playAudio(name: string) {

    try {
        const basePath = '/sounds/';
        
        const audio = new Audio(basePath + name + ".mp3");
        
        audio.volume = get(audioVolumeStore);
        
        navigator.userActivation.isActive && audio.play();
    } catch (error) {
        console.error("Error playing audio:", error);
    }

}


function volumeChange(volume: number) {
    audioVolumeStore.set(volume);
}

export {playAudio, volumeChange};