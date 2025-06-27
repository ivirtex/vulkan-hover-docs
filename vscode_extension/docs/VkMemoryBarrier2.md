# VkMemoryBarrier2(3) Manual Page

## Name

VkMemoryBarrier2 - Structure specifying a global memory barrier



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMemoryBarrier2` structure is defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkMemoryBarrier2 {
    VkStructureType          sType;
    const void*              pNext;
    VkPipelineStageFlags2    srcStageMask;
    VkAccessFlags2           srcAccessMask;
    VkPipelineStageFlags2    dstStageMask;
    VkAccessFlags2           dstAccessMask;
} VkMemoryBarrier2;
```

or the equivalent

``` c
// Provided by VK_KHR_synchronization2
typedef VkMemoryBarrier2 VkMemoryBarrier2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `srcStageMask` is a
  [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags2.html) mask of pipeline
  stages to be included in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
  target="_blank" rel="noopener">first synchronization scope</a>.

- `srcAccessMask` is a [VkAccessFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlags2.html) mask of
  access flags to be included in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
  target="_blank" rel="noopener">first access scope</a>.

- `dstStageMask` is a
  [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags2.html) mask of pipeline
  stages to be included in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
  target="_blank" rel="noopener">second synchronization scope</a>.

- `dstAccessMask` is a [VkAccessFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlags2.html) mask of
  access flags to be included in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
  target="_blank" rel="noopener">second access scope</a>.

## <a href="#_description" class="anchor"></a>Description

This structure defines a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-memory"
target="_blank" rel="noopener">memory dependency</a> affecting all
device memory.

The first <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
target="_blank" rel="noopener">access scope</a> described by this
structure include only operations and memory accesses specified by
`srcStageMask` and `srcAccessMask`.

The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
target="_blank" rel="noopener">access scope</a> described by this
structure include only operations and memory accesses specified by
`dstStageMask` and `dstAccessMask`.

Valid Usage

- <a href="#VUID-VkMemoryBarrier2-srcStageMask-03929"
  id="VUID-VkMemoryBarrier2-srcStageMask-03929"></a>
  VUID-VkMemoryBarrier2-srcStageMask-03929  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcStageMask-03930"
  id="VUID-VkMemoryBarrier2-srcStageMask-03930"></a>
  VUID-VkMemoryBarrier2-srcStageMask-03930  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcStageMask-03931"
  id="VUID-VkMemoryBarrier2-srcStageMask-03931"></a>
  VUID-VkMemoryBarrier2-srcStageMask-03931  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-VkMemoryBarrier2-srcStageMask-03932"
  id="VUID-VkMemoryBarrier2-srcStageMask-03932"></a>
  VUID-VkMemoryBarrier2-srcStageMask-03932  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-VkMemoryBarrier2-srcStageMask-03933"
  id="VUID-VkMemoryBarrier2-srcStageMask-03933"></a>
  VUID-VkMemoryBarrier2-srcStageMask-03933  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-VkMemoryBarrier2-srcStageMask-03934"
  id="VUID-VkMemoryBarrier2-srcStageMask-03934"></a>
  VUID-VkMemoryBarrier2-srcStageMask-03934  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`

- <a href="#VUID-VkMemoryBarrier2-srcStageMask-03935"
  id="VUID-VkMemoryBarrier2-srcStageMask-03935"></a>
  VUID-VkMemoryBarrier2-srcStageMask-03935  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`

- <a href="#VUID-VkMemoryBarrier2-srcStageMask-07316"
  id="VUID-VkMemoryBarrier2-srcStageMask-07316"></a>
  VUID-VkMemoryBarrier2-srcStageMask-07316  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-VkMemoryBarrier2-srcStageMask-04957"
  id="VUID-VkMemoryBarrier2-srcStageMask-04957"></a>
  VUID-VkMemoryBarrier2-srcStageMask-04957  
  If the [`subpassShading`](#features-subpassShading) feature is not
  enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`

- <a href="#VUID-VkMemoryBarrier2-srcStageMask-04995"
  id="VUID-VkMemoryBarrier2-srcStageMask-04995"></a>
  VUID-VkMemoryBarrier2-srcStageMask-04995  
  If the [`invocationMask`](#features-invocationMask) feature is not
  enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`

- <a href="#VUID-VkMemoryBarrier2-srcStageMask-07946"
  id="VUID-VkMemoryBarrier2-srcStageMask-07946"></a>
  VUID-VkMemoryBarrier2-srcStageMask-07946  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`

<!-- -->

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03900"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03900"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03900  
  If `srcAccessMask` includes `VK_ACCESS_2_INDIRECT_COMMAND_READ_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03901"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03901"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03901  
  If `srcAccessMask` includes `VK_ACCESS_2_INDEX_READ_BIT`,
  `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_INDEX_INPUT_BIT`,
  `VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03902"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03902"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03902  
  If `srcAccessMask` includes `VK_ACCESS_2_VERTEX_ATTRIBUTE_READ_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VERTEX_ATTRIBUTE_INPUT_BIT`,
  `VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03903"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03903"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03903  
  If `srcAccessMask` includes `VK_ACCESS_2_INPUT_ATTACHMENT_READ_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`,
  `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03904"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03904"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03904  
  If `srcAccessMask` includes `VK_ACCESS_2_UNIFORM_READ_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03905"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03905"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03905  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_SAMPLED_READ_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03906"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03906"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03906  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_STORAGE_READ_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03907"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03907"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03907  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-07454"
  id="VUID-VkMemoryBarrier2-srcAccessMask-07454"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-07454  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_READ_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03909"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03909"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03909  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_WRITE_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03910"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03910"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03910  
  If `srcAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT`
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03911"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03911"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03911  
  If `srcAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT`
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03912"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03912"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03912  
  If `srcAccessMask` includes
  `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT`, `srcStageMask`
  **must** include `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT`,
  `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03913"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03913"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03913  
  If `srcAccessMask` includes
  `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT`, `srcStageMask`
  **must** include `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT`,
  `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03914"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03914"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03914  
  If `srcAccessMask` includes `VK_ACCESS_2_TRANSFER_READ_BIT`,
  `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_COPY_BIT`,
  `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_RESOLVE_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03915"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03915"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03915  
  If `srcAccessMask` includes `VK_ACCESS_2_TRANSFER_WRITE_BIT`,
  `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_COPY_BIT`,
  `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_RESOLVE_BIT`,
  `VK_PIPELINE_STAGE_2_CLEAR_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, or
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03916"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03916"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03916  
  If `srcAccessMask` includes `VK_ACCESS_2_HOST_READ_BIT`,
  `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_HOST_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03917"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03917"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03917  
  If `srcAccessMask` includes `VK_ACCESS_2_HOST_WRITE_BIT`,
  `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_HOST_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03918"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03918"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03918  
  If `srcAccessMask` includes
  `VK_ACCESS_2_CONDITIONAL_RENDERING_READ_BIT_EXT`, `srcStageMask`
  **must** include `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03919"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03919"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03919  
  If `srcAccessMask` includes
  `VK_ACCESS_2_FRAGMENT_DENSITY_MAP_READ_BIT_EXT`, `srcStageMask`
  **must** include
  `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03920"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03920"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03920  
  If `srcAccessMask` includes
  `VK_ACCESS_2_TRANSFORM_FEEDBACK_WRITE_BIT_EXT`, `srcStageMask`
  **must** include `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-04747"
  id="VUID-VkMemoryBarrier2-srcAccessMask-04747"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-04747  
  If `srcAccessMask` includes
  `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT`, `srcStageMask`
  **must** include `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT`,
  `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03922"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03922"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03922  
  If `srcAccessMask` includes
  `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT`, `srcStageMask`
  **must** include `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03923"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03923"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03923  
  If `srcAccessMask` includes
  `VK_ACCESS_2_SHADING_RATE_IMAGE_READ_BIT_NV`, `srcStageMask` **must**
  include `VK_PIPELINE_STAGE_2_SHADING_RATE_IMAGE_BIT_NV`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-04994"
  id="VUID-VkMemoryBarrier2-srcAccessMask-04994"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-04994  
  If `srcAccessMask` includes
  `VK_ACCESS_2_INVOCATION_MASK_READ_BIT_HUAWEI`, `srcStageMask` **must**
  include `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03924"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03924"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03924  
  If `srcAccessMask` includes
  `VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_NV`, `srcStageMask` **must**
  include `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03925"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03925"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03925  
  If `srcAccessMask` includes
  `VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_NV`, `srcStageMask` **must**
  include `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03926"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03926"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03926  
  If `srcAccessMask` includes
  `VK_ACCESS_2_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT`
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03927"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03927"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03927  
  If `srcAccessMask` includes
  `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `srcStageMask`
  **must** include
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-03928"
  id="VUID-VkMemoryBarrier2-srcAccessMask-03928"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-03928  
  If `srcAccessMask` includes
  `VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`, `srcStageMask`
  **must** include
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-06256"
  id="VUID-VkMemoryBarrier2-srcAccessMask-06256"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-06256  
  If the [`rayQuery`](#features-rayQuery) feature is not enabled and
  `srcAccessMask` includes
  `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `srcStageMask`
  **must** not include any of the `VK_PIPELINE_STAGE_*_SHADER_BIT`
  stages except `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-07272"
  id="VUID-VkMemoryBarrier2-srcAccessMask-07272"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-07272  
  If `srcAccessMask` includes
  `VK_ACCESS_2_SHADER_BINDING_TABLE_READ_BIT_KHR`, `srcStageMask`
  **must** include `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT` or
  `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-04858"
  id="VUID-VkMemoryBarrier2-srcAccessMask-04858"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-04858  
  If `srcAccessMask` includes `VK_ACCESS_2_VIDEO_DECODE_READ_BIT_KHR`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-04859"
  id="VUID-VkMemoryBarrier2-srcAccessMask-04859"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-04859  
  If `srcAccessMask` includes `VK_ACCESS_2_VIDEO_DECODE_WRITE_BIT_KHR`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-04860"
  id="VUID-VkMemoryBarrier2-srcAccessMask-04860"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-04860  
  If `srcAccessMask` includes `VK_ACCESS_2_VIDEO_ENCODE_READ_BIT_KHR`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-04861"
  id="VUID-VkMemoryBarrier2-srcAccessMask-04861"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-04861  
  If `srcAccessMask` includes `VK_ACCESS_2_VIDEO_ENCODE_WRITE_BIT_KHR`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-07455"
  id="VUID-VkMemoryBarrier2-srcAccessMask-07455"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-07455  
  If `srcAccessMask` includes `VK_ACCESS_2_OPTICAL_FLOW_READ_BIT_NV`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-07456"
  id="VUID-VkMemoryBarrier2-srcAccessMask-07456"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-07456  
  If `srcAccessMask` includes `VK_ACCESS_2_OPTICAL_FLOW_WRITE_BIT_NV`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-07457"
  id="VUID-VkMemoryBarrier2-srcAccessMask-07457"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-07457  
  If `srcAccessMask` includes `VK_ACCESS_2_MICROMAP_WRITE_BIT_EXT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-07458"
  id="VUID-VkMemoryBarrier2-srcAccessMask-07458"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-07458  
  If `srcAccessMask` includes `VK_ACCESS_2_MICROMAP_READ_BIT_EXT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` or
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-08118"
  id="VUID-VkMemoryBarrier2-srcAccessMask-08118"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-08118  
  If `srcAccessMask` includes
  `VK_ACCESS_2_DESCRIPTOR_BUFFER_READ_BIT_EXT`, `srcStageMask` **must**
  include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

<!-- -->

- <a href="#VUID-VkMemoryBarrier2-dstStageMask-03929"
  id="VUID-VkMemoryBarrier2-dstStageMask-03929"></a>
  VUID-VkMemoryBarrier2-dstStageMask-03929  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstStageMask-03930"
  id="VUID-VkMemoryBarrier2-dstStageMask-03930"></a>
  VUID-VkMemoryBarrier2-dstStageMask-03930  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstStageMask-03931"
  id="VUID-VkMemoryBarrier2-dstStageMask-03931"></a>
  VUID-VkMemoryBarrier2-dstStageMask-03931  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-VkMemoryBarrier2-dstStageMask-03932"
  id="VUID-VkMemoryBarrier2-dstStageMask-03932"></a>
  VUID-VkMemoryBarrier2-dstStageMask-03932  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-VkMemoryBarrier2-dstStageMask-03933"
  id="VUID-VkMemoryBarrier2-dstStageMask-03933"></a>
  VUID-VkMemoryBarrier2-dstStageMask-03933  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-VkMemoryBarrier2-dstStageMask-03934"
  id="VUID-VkMemoryBarrier2-dstStageMask-03934"></a>
  VUID-VkMemoryBarrier2-dstStageMask-03934  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`

- <a href="#VUID-VkMemoryBarrier2-dstStageMask-03935"
  id="VUID-VkMemoryBarrier2-dstStageMask-03935"></a>
  VUID-VkMemoryBarrier2-dstStageMask-03935  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`

- <a href="#VUID-VkMemoryBarrier2-dstStageMask-07316"
  id="VUID-VkMemoryBarrier2-dstStageMask-07316"></a>
  VUID-VkMemoryBarrier2-dstStageMask-07316  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-VkMemoryBarrier2-dstStageMask-04957"
  id="VUID-VkMemoryBarrier2-dstStageMask-04957"></a>
  VUID-VkMemoryBarrier2-dstStageMask-04957  
  If the [`subpassShading`](#features-subpassShading) feature is not
  enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`

- <a href="#VUID-VkMemoryBarrier2-dstStageMask-04995"
  id="VUID-VkMemoryBarrier2-dstStageMask-04995"></a>
  VUID-VkMemoryBarrier2-dstStageMask-04995  
  If the [`invocationMask`](#features-invocationMask) feature is not
  enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`

- <a href="#VUID-VkMemoryBarrier2-dstStageMask-07946"
  id="VUID-VkMemoryBarrier2-dstStageMask-07946"></a>
  VUID-VkMemoryBarrier2-dstStageMask-07946  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`

<!-- -->

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03900"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03900"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03900  
  If `dstAccessMask` includes `VK_ACCESS_2_INDIRECT_COMMAND_READ_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03901"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03901"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03901  
  If `dstAccessMask` includes `VK_ACCESS_2_INDEX_READ_BIT`,
  `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_INDEX_INPUT_BIT`,
  `VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03902"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03902"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03902  
  If `dstAccessMask` includes `VK_ACCESS_2_VERTEX_ATTRIBUTE_READ_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VERTEX_ATTRIBUTE_INPUT_BIT`,
  `VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03903"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03903"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03903  
  If `dstAccessMask` includes `VK_ACCESS_2_INPUT_ATTACHMENT_READ_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`,
  `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03904"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03904"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03904  
  If `dstAccessMask` includes `VK_ACCESS_2_UNIFORM_READ_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03905"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03905"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03905  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_SAMPLED_READ_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03906"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03906"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03906  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_STORAGE_READ_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03907"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03907"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03907  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-07454"
  id="VUID-VkMemoryBarrier2-dstAccessMask-07454"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-07454  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_READ_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03909"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03909"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03909  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_WRITE_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03910"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03910"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03910  
  If `dstAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT`
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03911"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03911"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03911  
  If `dstAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT`
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03912"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03912"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03912  
  If `dstAccessMask` includes
  `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT`, `dstStageMask`
  **must** include `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT`,
  `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03913"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03913"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03913  
  If `dstAccessMask` includes
  `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT`, `dstStageMask`
  **must** include `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT`,
  `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03914"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03914"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03914  
  If `dstAccessMask` includes `VK_ACCESS_2_TRANSFER_READ_BIT`,
  `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_COPY_BIT`,
  `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_RESOLVE_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03915"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03915"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03915  
  If `dstAccessMask` includes `VK_ACCESS_2_TRANSFER_WRITE_BIT`,
  `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_COPY_BIT`,
  `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_RESOLVE_BIT`,
  `VK_PIPELINE_STAGE_2_CLEAR_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, or
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03916"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03916"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03916  
  If `dstAccessMask` includes `VK_ACCESS_2_HOST_READ_BIT`,
  `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_HOST_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03917"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03917"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03917  
  If `dstAccessMask` includes `VK_ACCESS_2_HOST_WRITE_BIT`,
  `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_HOST_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03918"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03918"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03918  
  If `dstAccessMask` includes
  `VK_ACCESS_2_CONDITIONAL_RENDERING_READ_BIT_EXT`, `dstStageMask`
  **must** include `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03919"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03919"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03919  
  If `dstAccessMask` includes
  `VK_ACCESS_2_FRAGMENT_DENSITY_MAP_READ_BIT_EXT`, `dstStageMask`
  **must** include
  `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03920"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03920"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03920  
  If `dstAccessMask` includes
  `VK_ACCESS_2_TRANSFORM_FEEDBACK_WRITE_BIT_EXT`, `dstStageMask`
  **must** include `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-04747"
  id="VUID-VkMemoryBarrier2-dstAccessMask-04747"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-04747  
  If `dstAccessMask` includes
  `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT`, `dstStageMask`
  **must** include `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT`,
  `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03922"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03922"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03922  
  If `dstAccessMask` includes
  `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT`, `dstStageMask`
  **must** include `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03923"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03923"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03923  
  If `dstAccessMask` includes
  `VK_ACCESS_2_SHADING_RATE_IMAGE_READ_BIT_NV`, `dstStageMask` **must**
  include `VK_PIPELINE_STAGE_2_SHADING_RATE_IMAGE_BIT_NV`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-04994"
  id="VUID-VkMemoryBarrier2-dstAccessMask-04994"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-04994  
  If `dstAccessMask` includes
  `VK_ACCESS_2_INVOCATION_MASK_READ_BIT_HUAWEI`, `dstStageMask` **must**
  include `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03924"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03924"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03924  
  If `dstAccessMask` includes
  `VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_NV`, `dstStageMask` **must**
  include `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03925"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03925"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03925  
  If `dstAccessMask` includes
  `VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_NV`, `dstStageMask` **must**
  include `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03926"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03926"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03926  
  If `dstAccessMask` includes
  `VK_ACCESS_2_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT`
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03927"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03927"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03927  
  If `dstAccessMask` includes
  `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `dstStageMask`
  **must** include
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-03928"
  id="VUID-VkMemoryBarrier2-dstAccessMask-03928"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-03928  
  If `dstAccessMask` includes
  `VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`, `dstStageMask`
  **must** include
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-06256"
  id="VUID-VkMemoryBarrier2-dstAccessMask-06256"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-06256  
  If the [`rayQuery`](#features-rayQuery) feature is not enabled and
  `dstAccessMask` includes
  `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `dstStageMask`
  **must** not include any of the `VK_PIPELINE_STAGE_*_SHADER_BIT`
  stages except `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-07272"
  id="VUID-VkMemoryBarrier2-dstAccessMask-07272"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-07272  
  If `dstAccessMask` includes
  `VK_ACCESS_2_SHADER_BINDING_TABLE_READ_BIT_KHR`, `dstStageMask`
  **must** include `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT` or
  `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-04858"
  id="VUID-VkMemoryBarrier2-dstAccessMask-04858"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-04858  
  If `dstAccessMask` includes `VK_ACCESS_2_VIDEO_DECODE_READ_BIT_KHR`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-04859"
  id="VUID-VkMemoryBarrier2-dstAccessMask-04859"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-04859  
  If `dstAccessMask` includes `VK_ACCESS_2_VIDEO_DECODE_WRITE_BIT_KHR`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-04860"
  id="VUID-VkMemoryBarrier2-dstAccessMask-04860"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-04860  
  If `dstAccessMask` includes `VK_ACCESS_2_VIDEO_ENCODE_READ_BIT_KHR`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-04861"
  id="VUID-VkMemoryBarrier2-dstAccessMask-04861"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-04861  
  If `dstAccessMask` includes `VK_ACCESS_2_VIDEO_ENCODE_WRITE_BIT_KHR`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-07455"
  id="VUID-VkMemoryBarrier2-dstAccessMask-07455"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-07455  
  If `dstAccessMask` includes `VK_ACCESS_2_OPTICAL_FLOW_READ_BIT_NV`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-07456"
  id="VUID-VkMemoryBarrier2-dstAccessMask-07456"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-07456  
  If `dstAccessMask` includes `VK_ACCESS_2_OPTICAL_FLOW_WRITE_BIT_NV`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-07457"
  id="VUID-VkMemoryBarrier2-dstAccessMask-07457"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-07457  
  If `dstAccessMask` includes `VK_ACCESS_2_MICROMAP_WRITE_BIT_EXT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-07458"
  id="VUID-VkMemoryBarrier2-dstAccessMask-07458"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-07458  
  If `dstAccessMask` includes `VK_ACCESS_2_MICROMAP_READ_BIT_EXT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` or
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-08118"
  id="VUID-VkMemoryBarrier2-dstAccessMask-08118"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-08118  
  If `dstAccessMask` includes
  `VK_ACCESS_2_DESCRIPTOR_BUFFER_READ_BIT_EXT`, `dstStageMask` **must**
  include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryBarrier2-sType-sType"
  id="VUID-VkMemoryBarrier2-sType-sType"></a>
  VUID-VkMemoryBarrier2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_BARRIER_2`

- <a href="#VUID-VkMemoryBarrier2-srcStageMask-parameter"
  id="VUID-VkMemoryBarrier2-srcStageMask-parameter"></a>
  VUID-VkMemoryBarrier2-srcStageMask-parameter  
  `srcStageMask` **must** be a valid combination of
  [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html) values

- <a href="#VUID-VkMemoryBarrier2-srcAccessMask-parameter"
  id="VUID-VkMemoryBarrier2-srcAccessMask-parameter"></a>
  VUID-VkMemoryBarrier2-srcAccessMask-parameter  
  `srcAccessMask` **must** be a valid combination of
  [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2.html) values

- <a href="#VUID-VkMemoryBarrier2-dstStageMask-parameter"
  id="VUID-VkMemoryBarrier2-dstStageMask-parameter"></a>
  VUID-VkMemoryBarrier2-dstStageMask-parameter  
  `dstStageMask` **must** be a valid combination of
  [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html) values

- <a href="#VUID-VkMemoryBarrier2-dstAccessMask-parameter"
  id="VUID-VkMemoryBarrier2-dstAccessMask-parameter"></a>
  VUID-VkMemoryBarrier2-dstAccessMask-parameter  
  `dstAccessMask` **must** be a valid combination of
  [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkAccessFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlags2.html),
[VkDependencyInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyInfo.html),
[VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryBarrier2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
