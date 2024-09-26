let SYSTEM_INFO = uni.getSystemInfoSync()
export const getStatusBarHeight = () => SYSTEM_INFO.statusBarHeight || 0
export const getTitbarHeight =()=>{
    let {top,height} = uni.getMenuButtonBoundingClientRect()
    return  height + (top - SYSTEM_INFO.statusBarHeight) * 2
}

export const getNavBarheight = () => getStatusBarHeight() + getTitbarHeight()