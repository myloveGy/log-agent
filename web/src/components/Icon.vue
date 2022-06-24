<script lang="ts">
import {h, SetupContext} from 'vue'
import {ElIcon} from 'element-plus'

interface IconProps {
  name: string;
  color: string;
}

export default (props: IconProps, ctx: SetupContext) => {
  return h(ElIcon, {}, () => [
    h(props.name, ctx.attrs, ctx.slots),
  ])
}
</script>
