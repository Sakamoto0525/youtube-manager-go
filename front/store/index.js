import { createRequestClient } from '~/store/request-client';

export const state = () => ({
  items: [],
  item: [],
  meta: {},
})

export const actions = {
  /**
   * Video一覧を取得
   */
  async fetchPopularVideos({ commit }, payload) {
    const client = createRequestClient(this.$axios)
    const res = await client.get(payload.uri, payload.params)
    commit('mutatePopularVideos', res)
  },
  /**
   * Video詳細を取得
   */
  async findVideo({ commit }, payload) {
    const client = createRequestClient(this.$axios)
    const res = await client.get(payload.uri)
    const params = {
      ...res.video_list
    }
    commit('mutateVideo', params)
  }
}

export const mutations = {
  mutatePopularVideos(state, payload) {
    state.items = payload.items ? state.items.concat(payload.items) : []
    state.meta = payload
  },
  mutateVideo(state, payload) {
    console.log('きてるよ！')
    const params = (payload.items && payload.items.length > 0) ? payload.items[0] : {}
    state.item = params
  }
}

export const getters = {
  getPopularVideos(state) {
    return state.items
  },
  getMeta(state) {
    return state.meta
  },
  getVideo(state) {
    return state.item
  }
}