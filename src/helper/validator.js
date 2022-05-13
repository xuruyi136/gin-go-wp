const telephoneValidator = (value) => /^1[3|5|7|9]\d{9}$/.test(value);

export default {
    telephoneValidator,
};