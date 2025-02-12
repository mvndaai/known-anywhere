const defaultLang = 'en';
const avaliableLangs = [defaultLang];
const langsFile = {}

const loadLang = async (l) => {
    try {
        const module = await import(`./translations/${l}.js`);
        return module.default;
    } catch (e) {
        console.error("missing lang", l, e)
    }
};

const getLang = async (l) => {
    const f = langsFile[l];
    if (f) {
        return f;
    }
    if (!avaliableLangs.includes(l)) {
        return
    }
    langsFile[l] = await loadLang(l)
    return langsFile[l]
}

const getString = async (lang, keys, count, args) => {
    const f = await getLang(lang);
    if (!f) {
        return
    }
    const keyList = Array.isArray(keys) ? keys : keys?.split('.');
    let v = keyList?.reduce((obj, key) => obj?.[key], f);
    if (v && typeof v === 'object') {
        v = count === 1 ? v?.single : v?.multi;
        v = formatString(v, args)
    }
    return v
}

const formatString = (str, args) => {
    for( var arg in args ) {
        str = str.replace("{" + arg + "}", args[arg]);
    }
    return str;
};

const i18n = async (keys, count, args) => {
    // Todo get Accept-Language header
    const langs = []
    if (!langs.includes(defaultLang)){
        // Add fallback to default language
        langs.push(defaultLang)
    }

    for (const lang of langs) {
        const v = await getString(lang, keys, count, args)
        if (v) {
            return v
        }
    }
    console.error("missing i18n", keys, count)
    return null;
};



export { i18n };