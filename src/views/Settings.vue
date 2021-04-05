<template>
    <div id="app">
        <h1 class="header">
            My<span style="color:rgb(215, 0, 0);">N</span>otes
        </h1>
        <hr class="whiteLine">
        <br>
        <div>
            <hr class="splitLine"> 
            <h3 @click="deleteData()" class="settingWrapper">
                <img src="../assets/trash.png" alt="trash icon" id="trash">
                <span id="setting">{{ $t("text.settings.delete") }}</span>
            </h3>

            <hr class="splitLine">

            <h3 class="settingWrapper">
                <img src="../assets/translation.png" alt="trash icon" id="trash">
                <select
                    v-model="$i18n.locale"
                    id="language"
                    v-bind="selectedLanguage"
                    >
                    <option
                        v-for="(lang, i) in langs" :key="`Lang${i}`"
                        :value="lang"
                    >
                        {{ lang }}
                    </option>
                </select>
            </h3>

            <hr class="splitLine">
        </div>
        <button
            @click="$router.push('/')"
            class="saveButton"
            id="home"
        >
                {{ $t("text.settings.home") }}
        </button>
    </div>
</template>

<script>
export default {
    data () {
        return {
            langs: ['German', 'English'],
            selectedLanguage: '',
        }
    },
    methods: {
        deleteData () {
            const vm = this
            this.$dialog.confirm(this.$t('text.deleteAlert'))
                .then (function () {
                    localStorage.clear()
                    vm.$router.push('/')
                })
                .catch (function () {
                })
        }
    },
    mounted () {

    },
    created () {
        document.getElementById('body').style.overflow = 'hidden'
        // this.$store.state.transitionName = 'swipe-right'
    },
    beforeRouteLeave (to, from, next) {
        localStorage.setItem('language', this.$i18n.locale)
        next()
    }
}
</script>

<style scoped>
    .splitLine{
      border-top: 3px solid rgb(134, 134, 134);
      width: 100%;
      opacity: 0.5;
      margin-bottom: 1.3rem;
    }
    #trash {
        width: 2.2rem;
        margin-bottom: 0.5rem;
        margin-right: 0.5rem;
    }
    .settingWrapper {
        text-align: left;
        margin-left: 1rem;
        cursor: pointer;
    }
    #setting {
        margin-left: 1rem;
    }
    #home {
        margin-bottom: 1.5rem;
    }
    #language {
        margin-left: 1rem;
    }
</style>