export const validate = (data) => {
    return !!data.match(/^\d{10}$/);
};
