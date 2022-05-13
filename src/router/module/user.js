const userRoutes = [
    {
        path: '/login',
        name: 'login',
        component: () => import('@/components/Login'),
    },
    {
        path: '/register',
        name: 'register',
        component: () => import('@/components/Register'),
    },
    {
        path: '/profile',
        name: 'profile',
        meta: {
            auth: true,
        },
        component: () => import('@/components/Profile'),
    }
]
export default userRoutes;