# VkImageMemoryBarrier2(3) Manual Page

## Name

VkImageMemoryBarrier2 - Structure specifying an image memory barrier



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageMemoryBarrier2` structure is defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkImageMemoryBarrier2 {
    VkStructureType            sType;
    const void*                pNext;
    VkPipelineStageFlags2      srcStageMask;
    VkAccessFlags2             srcAccessMask;
    VkPipelineStageFlags2      dstStageMask;
    VkAccessFlags2             dstAccessMask;
    VkImageLayout              oldLayout;
    VkImageLayout              newLayout;
    uint32_t                   srcQueueFamilyIndex;
    uint32_t                   dstQueueFamilyIndex;
    VkImage                    image;
    VkImageSubresourceRange    subresourceRange;
} VkImageMemoryBarrier2;
```

or the equivalent

``` c
// Provided by VK_KHR_synchronization2
typedef VkImageMemoryBarrier2 VkImageMemoryBarrier2KHR;
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

- `oldLayout` is the old layout in an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-image-layout-transitions"
  target="_blank" rel="noopener">image layout transition</a>.

- `newLayout` is the new layout in an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-image-layout-transitions"
  target="_blank" rel="noopener">image layout transition</a>.

- `srcQueueFamilyIndex` is the source queue family for a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers"
  target="_blank" rel="noopener">queue family ownership transfer</a>.

- `dstQueueFamilyIndex` is the destination queue family for a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers"
  target="_blank" rel="noopener">queue family ownership transfer</a>.

- `image` is a handle to the image affected by this barrier.

- `subresourceRange` describes the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views"
  target="_blank" rel="noopener">image subresource range</a> within
  `image` that is affected by this barrier.

## <a href="#_description" class="anchor"></a>Description

This structure defines a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-memory"
target="_blank" rel="noopener">memory dependency</a> limited to an image
subresource range, and **can** define a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers"
target="_blank" rel="noopener">queue family ownership transfer
operation</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-image-layout-transitions"
target="_blank" rel="noopener">image layout transition</a> for that
subresource range.

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

Both <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
target="_blank" rel="noopener">access scopes</a> are limited to only
memory accesses to `image` in the subresource range defined by
`subresourceRange`.

If `image` was created with `VK_SHARING_MODE_EXCLUSIVE`, and
`srcQueueFamilyIndex` is not equal to `dstQueueFamilyIndex`, this memory
barrier defines a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers"
target="_blank" rel="noopener">queue family ownership transfer
operation</a>. When executed on a queue in the family identified by
`srcQueueFamilyIndex`, this barrier defines a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers-release"
target="_blank" rel="noopener">queue family release operation</a> for
the specified image subresource range, and the second synchronization
scope does not apply to this operation. When executed on a queue in the
family identified by `dstQueueFamilyIndex`, this barrier defines a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers-acquire"
target="_blank" rel="noopener">queue family acquire operation</a> for
the specified image subresource range, and the first synchronization,
the first synchronization scope does not apply to this operation.

A <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers"
target="_blank" rel="noopener">queue family ownership transfer
operation</a> is also defined if the values are not equal, and either is
one of the special queue family values reserved for external memory
ownership transfers, as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers"
class="bare" target="_blank"
rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers</a>.
A <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers-release"
target="_blank" rel="noopener">queue family release operation</a> is
defined when `dstQueueFamilyIndex` is one of those values, and a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers-acquire"
target="_blank" rel="noopener">queue family acquire operation</a> is
defined when `srcQueueFamilyIndex` is one of those values.

If `oldLayout` is not equal to `newLayout`, then the memory barrier
defines an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-image-layout-transitions"
target="_blank" rel="noopener">image layout transition</a> for the
specified image subresource range. If this memory barrier defines a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers"
target="_blank" rel="noopener">queue family ownership transfer
operation</a>, the layout transition is only executed once between the
queues.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>When the old and new layout are equal, the layout values are ignored
- data is preserved no matter what values are specified, or what layout
the image is currently in.</p></td>
</tr>
</tbody>
</table>

If `image` has a multi-planar format and the image is *disjoint*, then
including `VK_IMAGE_ASPECT_COLOR_BIT` in the `aspectMask` member of
`subresourceRange` is equivalent to including
`VK_IMAGE_ASPECT_PLANE_0_BIT`, `VK_IMAGE_ASPECT_PLANE_1_BIT`, and (for
three-plane formats only) `VK_IMAGE_ASPECT_PLANE_2_BIT`.

Valid Usage

- <a href="#VUID-VkImageMemoryBarrier2-srcStageMask-03929"
  id="VUID-VkImageMemoryBarrier2-srcStageMask-03929"></a>
  VUID-VkImageMemoryBarrier2-srcStageMask-03929  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcStageMask-03930"
  id="VUID-VkImageMemoryBarrier2-srcStageMask-03930"></a>
  VUID-VkImageMemoryBarrier2-srcStageMask-03930  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcStageMask-03931"
  id="VUID-VkImageMemoryBarrier2-srcStageMask-03931"></a>
  VUID-VkImageMemoryBarrier2-srcStageMask-03931  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-VkImageMemoryBarrier2-srcStageMask-03932"
  id="VUID-VkImageMemoryBarrier2-srcStageMask-03932"></a>
  VUID-VkImageMemoryBarrier2-srcStageMask-03932  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-VkImageMemoryBarrier2-srcStageMask-03933"
  id="VUID-VkImageMemoryBarrier2-srcStageMask-03933"></a>
  VUID-VkImageMemoryBarrier2-srcStageMask-03933  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-VkImageMemoryBarrier2-srcStageMask-03934"
  id="VUID-VkImageMemoryBarrier2-srcStageMask-03934"></a>
  VUID-VkImageMemoryBarrier2-srcStageMask-03934  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`

- <a href="#VUID-VkImageMemoryBarrier2-srcStageMask-03935"
  id="VUID-VkImageMemoryBarrier2-srcStageMask-03935"></a>
  VUID-VkImageMemoryBarrier2-srcStageMask-03935  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`

- <a href="#VUID-VkImageMemoryBarrier2-srcStageMask-07316"
  id="VUID-VkImageMemoryBarrier2-srcStageMask-07316"></a>
  VUID-VkImageMemoryBarrier2-srcStageMask-07316  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-srcStageMask-04957"
  id="VUID-VkImageMemoryBarrier2-srcStageMask-04957"></a>
  VUID-VkImageMemoryBarrier2-srcStageMask-04957  
  If the [`subpassShading`](#features-subpassShading) feature is not
  enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`

- <a href="#VUID-VkImageMemoryBarrier2-srcStageMask-04995"
  id="VUID-VkImageMemoryBarrier2-srcStageMask-04995"></a>
  VUID-VkImageMemoryBarrier2-srcStageMask-04995  
  If the [`invocationMask`](#features-invocationMask) feature is not
  enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`

- <a href="#VUID-VkImageMemoryBarrier2-srcStageMask-07946"
  id="VUID-VkImageMemoryBarrier2-srcStageMask-07946"></a>
  VUID-VkImageMemoryBarrier2-srcStageMask-07946  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`

<!-- -->

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03900"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03900"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03900  
  If `srcAccessMask` includes `VK_ACCESS_2_INDIRECT_COMMAND_READ_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03901"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03901"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03901  
  If `srcAccessMask` includes `VK_ACCESS_2_INDEX_READ_BIT`,
  `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_INDEX_INPUT_BIT`,
  `VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03902"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03902"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03902  
  If `srcAccessMask` includes `VK_ACCESS_2_VERTEX_ATTRIBUTE_READ_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VERTEX_ATTRIBUTE_INPUT_BIT`,
  `VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03903"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03903"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03903  
  If `srcAccessMask` includes `VK_ACCESS_2_INPUT_ATTACHMENT_READ_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`,
  `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03904"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03904"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03904  
  If `srcAccessMask` includes `VK_ACCESS_2_UNIFORM_READ_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03905"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03905"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03905  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_SAMPLED_READ_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03906"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03906"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03906  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_STORAGE_READ_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03907"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03907"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03907  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-07454"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-07454"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-07454  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_READ_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03909"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03909"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03909  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_WRITE_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03910"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03910"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03910  
  If `srcAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT`
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03911"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03911"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03911  
  If `srcAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT`
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03912"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03912"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03912  
  If `srcAccessMask` includes
  `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT`, `srcStageMask`
  **must** include `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT`,
  `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03913"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03913"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03913  
  If `srcAccessMask` includes
  `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT`, `srcStageMask`
  **must** include `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT`,
  `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03914"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03914"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03914  
  If `srcAccessMask` includes `VK_ACCESS_2_TRANSFER_READ_BIT`,
  `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_COPY_BIT`,
  `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_RESOLVE_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03915"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03915"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03915  
  If `srcAccessMask` includes `VK_ACCESS_2_TRANSFER_WRITE_BIT`,
  `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_COPY_BIT`,
  `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_RESOLVE_BIT`,
  `VK_PIPELINE_STAGE_2_CLEAR_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, or
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03916"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03916"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03916  
  If `srcAccessMask` includes `VK_ACCESS_2_HOST_READ_BIT`,
  `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_HOST_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03917"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03917"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03917  
  If `srcAccessMask` includes `VK_ACCESS_2_HOST_WRITE_BIT`,
  `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_HOST_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03918"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03918"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03918  
  If `srcAccessMask` includes
  `VK_ACCESS_2_CONDITIONAL_RENDERING_READ_BIT_EXT`, `srcStageMask`
  **must** include `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03919"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03919"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03919  
  If `srcAccessMask` includes
  `VK_ACCESS_2_FRAGMENT_DENSITY_MAP_READ_BIT_EXT`, `srcStageMask`
  **must** include
  `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03920"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03920"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03920  
  If `srcAccessMask` includes
  `VK_ACCESS_2_TRANSFORM_FEEDBACK_WRITE_BIT_EXT`, `srcStageMask`
  **must** include `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-04747"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-04747"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-04747  
  If `srcAccessMask` includes
  `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT`, `srcStageMask`
  **must** include `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT`,
  `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03922"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03922"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03922  
  If `srcAccessMask` includes
  `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT`, `srcStageMask`
  **must** include `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03923"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03923"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03923  
  If `srcAccessMask` includes
  `VK_ACCESS_2_SHADING_RATE_IMAGE_READ_BIT_NV`, `srcStageMask` **must**
  include `VK_PIPELINE_STAGE_2_SHADING_RATE_IMAGE_BIT_NV`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-04994"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-04994"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-04994  
  If `srcAccessMask` includes
  `VK_ACCESS_2_INVOCATION_MASK_READ_BIT_HUAWEI`, `srcStageMask` **must**
  include `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03924"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03924"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03924  
  If `srcAccessMask` includes
  `VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_NV`, `srcStageMask` **must**
  include `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03925"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03925"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03925  
  If `srcAccessMask` includes
  `VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_NV`, `srcStageMask` **must**
  include `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03926"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03926"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03926  
  If `srcAccessMask` includes
  `VK_ACCESS_2_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT`
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03927"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03927"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03927  
  If `srcAccessMask` includes
  `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `srcStageMask`
  **must** include
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-03928"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-03928"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-03928  
  If `srcAccessMask` includes
  `VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`, `srcStageMask`
  **must** include
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-06256"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-06256"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-06256  
  If the [`rayQuery`](#features-rayQuery) feature is not enabled and
  `srcAccessMask` includes
  `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `srcStageMask`
  **must** not include any of the `VK_PIPELINE_STAGE_*_SHADER_BIT`
  stages except `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-07272"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-07272"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-07272  
  If `srcAccessMask` includes
  `VK_ACCESS_2_SHADER_BINDING_TABLE_READ_BIT_KHR`, `srcStageMask`
  **must** include `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT` or
  `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-04858"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-04858"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-04858  
  If `srcAccessMask` includes `VK_ACCESS_2_VIDEO_DECODE_READ_BIT_KHR`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-04859"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-04859"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-04859  
  If `srcAccessMask` includes `VK_ACCESS_2_VIDEO_DECODE_WRITE_BIT_KHR`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-04860"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-04860"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-04860  
  If `srcAccessMask` includes `VK_ACCESS_2_VIDEO_ENCODE_READ_BIT_KHR`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-04861"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-04861"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-04861  
  If `srcAccessMask` includes `VK_ACCESS_2_VIDEO_ENCODE_WRITE_BIT_KHR`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-07455"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-07455"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-07455  
  If `srcAccessMask` includes `VK_ACCESS_2_OPTICAL_FLOW_READ_BIT_NV`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-07456"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-07456"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-07456  
  If `srcAccessMask` includes `VK_ACCESS_2_OPTICAL_FLOW_WRITE_BIT_NV`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-07457"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-07457"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-07457  
  If `srcAccessMask` includes `VK_ACCESS_2_MICROMAP_WRITE_BIT_EXT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-07458"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-07458"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-07458  
  If `srcAccessMask` includes `VK_ACCESS_2_MICROMAP_READ_BIT_EXT`,
  `srcStageMask` **must** include
  `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` or
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-08118"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-08118"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-08118  
  If `srcAccessMask` includes
  `VK_ACCESS_2_DESCRIPTOR_BUFFER_READ_BIT_EXT`, `srcStageMask` **must**
  include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

<!-- -->

- <a href="#VUID-VkImageMemoryBarrier2-dstStageMask-03929"
  id="VUID-VkImageMemoryBarrier2-dstStageMask-03929"></a>
  VUID-VkImageMemoryBarrier2-dstStageMask-03929  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstStageMask-03930"
  id="VUID-VkImageMemoryBarrier2-dstStageMask-03930"></a>
  VUID-VkImageMemoryBarrier2-dstStageMask-03930  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstStageMask-03931"
  id="VUID-VkImageMemoryBarrier2-dstStageMask-03931"></a>
  VUID-VkImageMemoryBarrier2-dstStageMask-03931  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-VkImageMemoryBarrier2-dstStageMask-03932"
  id="VUID-VkImageMemoryBarrier2-dstStageMask-03932"></a>
  VUID-VkImageMemoryBarrier2-dstStageMask-03932  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-VkImageMemoryBarrier2-dstStageMask-03933"
  id="VUID-VkImageMemoryBarrier2-dstStageMask-03933"></a>
  VUID-VkImageMemoryBarrier2-dstStageMask-03933  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-VkImageMemoryBarrier2-dstStageMask-03934"
  id="VUID-VkImageMemoryBarrier2-dstStageMask-03934"></a>
  VUID-VkImageMemoryBarrier2-dstStageMask-03934  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`

- <a href="#VUID-VkImageMemoryBarrier2-dstStageMask-03935"
  id="VUID-VkImageMemoryBarrier2-dstStageMask-03935"></a>
  VUID-VkImageMemoryBarrier2-dstStageMask-03935  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`

- <a href="#VUID-VkImageMemoryBarrier2-dstStageMask-07316"
  id="VUID-VkImageMemoryBarrier2-dstStageMask-07316"></a>
  VUID-VkImageMemoryBarrier2-dstStageMask-07316  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-dstStageMask-04957"
  id="VUID-VkImageMemoryBarrier2-dstStageMask-04957"></a>
  VUID-VkImageMemoryBarrier2-dstStageMask-04957  
  If the [`subpassShading`](#features-subpassShading) feature is not
  enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`

- <a href="#VUID-VkImageMemoryBarrier2-dstStageMask-04995"
  id="VUID-VkImageMemoryBarrier2-dstStageMask-04995"></a>
  VUID-VkImageMemoryBarrier2-dstStageMask-04995  
  If the [`invocationMask`](#features-invocationMask) feature is not
  enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`

- <a href="#VUID-VkImageMemoryBarrier2-dstStageMask-07946"
  id="VUID-VkImageMemoryBarrier2-dstStageMask-07946"></a>
  VUID-VkImageMemoryBarrier2-dstStageMask-07946  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`

<!-- -->

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03900"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03900"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03900  
  If `dstAccessMask` includes `VK_ACCESS_2_INDIRECT_COMMAND_READ_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03901"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03901"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03901  
  If `dstAccessMask` includes `VK_ACCESS_2_INDEX_READ_BIT`,
  `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_INDEX_INPUT_BIT`,
  `VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03902"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03902"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03902  
  If `dstAccessMask` includes `VK_ACCESS_2_VERTEX_ATTRIBUTE_READ_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VERTEX_ATTRIBUTE_INPUT_BIT`,
  `VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03903"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03903"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03903  
  If `dstAccessMask` includes `VK_ACCESS_2_INPUT_ATTACHMENT_READ_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`,
  `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03904"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03904"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03904  
  If `dstAccessMask` includes `VK_ACCESS_2_UNIFORM_READ_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03905"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03905"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03905  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_SAMPLED_READ_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03906"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03906"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03906  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_STORAGE_READ_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03907"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03907"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03907  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-07454"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-07454"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-07454  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_READ_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03909"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03909"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03909  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_WRITE_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03910"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03910"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03910  
  If `dstAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT`
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03911"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03911"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03911  
  If `dstAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT`
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03912"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03912"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03912  
  If `dstAccessMask` includes
  `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT`, `dstStageMask`
  **must** include `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT`,
  `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03913"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03913"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03913  
  If `dstAccessMask` includes
  `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT`, `dstStageMask`
  **must** include `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT`,
  `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03914"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03914"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03914  
  If `dstAccessMask` includes `VK_ACCESS_2_TRANSFER_READ_BIT`,
  `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_COPY_BIT`,
  `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_RESOLVE_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03915"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03915"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03915  
  If `dstAccessMask` includes `VK_ACCESS_2_TRANSFER_WRITE_BIT`,
  `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_COPY_BIT`,
  `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_RESOLVE_BIT`,
  `VK_PIPELINE_STAGE_2_CLEAR_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, or
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03916"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03916"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03916  
  If `dstAccessMask` includes `VK_ACCESS_2_HOST_READ_BIT`,
  `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_HOST_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03917"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03917"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03917  
  If `dstAccessMask` includes `VK_ACCESS_2_HOST_WRITE_BIT`,
  `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_HOST_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03918"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03918"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03918  
  If `dstAccessMask` includes
  `VK_ACCESS_2_CONDITIONAL_RENDERING_READ_BIT_EXT`, `dstStageMask`
  **must** include `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03919"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03919"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03919  
  If `dstAccessMask` includes
  `VK_ACCESS_2_FRAGMENT_DENSITY_MAP_READ_BIT_EXT`, `dstStageMask`
  **must** include
  `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03920"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03920"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03920  
  If `dstAccessMask` includes
  `VK_ACCESS_2_TRANSFORM_FEEDBACK_WRITE_BIT_EXT`, `dstStageMask`
  **must** include `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-04747"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-04747"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-04747  
  If `dstAccessMask` includes
  `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT`, `dstStageMask`
  **must** include `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT`,
  `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03922"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03922"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03922  
  If `dstAccessMask` includes
  `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT`, `dstStageMask`
  **must** include `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03923"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03923"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03923  
  If `dstAccessMask` includes
  `VK_ACCESS_2_SHADING_RATE_IMAGE_READ_BIT_NV`, `dstStageMask` **must**
  include `VK_PIPELINE_STAGE_2_SHADING_RATE_IMAGE_BIT_NV`,
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-04994"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-04994"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-04994  
  If `dstAccessMask` includes
  `VK_ACCESS_2_INVOCATION_MASK_READ_BIT_HUAWEI`, `dstStageMask` **must**
  include `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03924"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03924"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03924  
  If `dstAccessMask` includes
  `VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_NV`, `dstStageMask` **must**
  include `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03925"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03925"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03925  
  If `dstAccessMask` includes
  `VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_NV`, `dstStageMask` **must**
  include `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03926"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03926"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03926  
  If `dstAccessMask` includes
  `VK_ACCESS_2_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT`
  `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03927"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03927"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03927  
  If `dstAccessMask` includes
  `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `dstStageMask`
  **must** include
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-03928"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-03928"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-03928  
  If `dstAccessMask` includes
  `VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`, `dstStageMask`
  **must** include
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`,
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` or
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-06256"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-06256"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-06256  
  If the [`rayQuery`](#features-rayQuery) feature is not enabled and
  `dstAccessMask` includes
  `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `dstStageMask`
  **must** not include any of the `VK_PIPELINE_STAGE_*_SHADER_BIT`
  stages except `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-07272"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-07272"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-07272  
  If `dstAccessMask` includes
  `VK_ACCESS_2_SHADER_BINDING_TABLE_READ_BIT_KHR`, `dstStageMask`
  **must** include `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT` or
  `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-04858"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-04858"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-04858  
  If `dstAccessMask` includes `VK_ACCESS_2_VIDEO_DECODE_READ_BIT_KHR`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-04859"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-04859"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-04859  
  If `dstAccessMask` includes `VK_ACCESS_2_VIDEO_DECODE_WRITE_BIT_KHR`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-04860"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-04860"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-04860  
  If `dstAccessMask` includes `VK_ACCESS_2_VIDEO_ENCODE_READ_BIT_KHR`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-04861"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-04861"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-04861  
  If `dstAccessMask` includes `VK_ACCESS_2_VIDEO_ENCODE_WRITE_BIT_KHR`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-07455"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-07455"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-07455  
  If `dstAccessMask` includes `VK_ACCESS_2_OPTICAL_FLOW_READ_BIT_NV`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-07456"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-07456"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-07456  
  If `dstAccessMask` includes `VK_ACCESS_2_OPTICAL_FLOW_WRITE_BIT_NV`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-07457"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-07457"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-07457  
  If `dstAccessMask` includes `VK_ACCESS_2_MICROMAP_WRITE_BIT_EXT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-07458"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-07458"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-07458  
  If `dstAccessMask` includes `VK_ACCESS_2_MICROMAP_READ_BIT_EXT`,
  `dstStageMask` **must** include
  `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` or
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-08118"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-08118"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-08118  
  If `dstAccessMask` includes
  `VK_ACCESS_2_DESCRIPTOR_BUFFER_READ_BIT_EXT`, `dstStageMask` **must**
  include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`,
  `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of
  `VK_PIPELINE_STAGE_*_SHADER_BIT` stages

<!-- -->

- <a href="#VUID-VkImageMemoryBarrier2-oldLayout-01208"
  id="VUID-VkImageMemoryBarrier2-oldLayout-01208"></a>
  VUID-VkImageMemoryBarrier2-oldLayout-01208  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL` then `image` **must** have
  been created with `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-oldLayout-01209"
  id="VUID-VkImageMemoryBarrier2-oldLayout-01209"></a>
  VUID-VkImageMemoryBarrier2-oldLayout-01209  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` then `image`
  **must** have been created with
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-oldLayout-01210"
  id="VUID-VkImageMemoryBarrier2-oldLayout-01210"></a>
  VUID-VkImageMemoryBarrier2-oldLayout-01210  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL` then `image`
  **must** have been created with
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-oldLayout-01211"
  id="VUID-VkImageMemoryBarrier2-oldLayout-01211"></a>
  VUID-VkImageMemoryBarrier2-oldLayout-01211  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL` then `image` **must** have
  been created with `VK_IMAGE_USAGE_SAMPLED_BIT` or
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-oldLayout-01212"
  id="VUID-VkImageMemoryBarrier2-oldLayout-01212"></a>
  VUID-VkImageMemoryBarrier2-oldLayout-01212  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL`
  then `image` **must** have been created with
  `VK_IMAGE_USAGE_TRANSFER_SRC_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-oldLayout-01213"
  id="VUID-VkImageMemoryBarrier2-oldLayout-01213"></a>
  VUID-VkImageMemoryBarrier2-oldLayout-01213  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`
  then `image` **must** have been created with
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-oldLayout-01197"
  id="VUID-VkImageMemoryBarrier2-oldLayout-01197"></a>
  VUID-VkImageMemoryBarrier2-oldLayout-01197  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), `oldLayout`
  **must** be `VK_IMAGE_LAYOUT_UNDEFINED` or the current layout of the
  image subresources affected by the barrier

- <a href="#VUID-VkImageMemoryBarrier2-newLayout-01198"
  id="VUID-VkImageMemoryBarrier2-newLayout-01198"></a>
  VUID-VkImageMemoryBarrier2-newLayout-01198  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), `newLayout`
  **must** not be `VK_IMAGE_LAYOUT_UNDEFINED` or
  `VK_IMAGE_LAYOUT_PREINITIALIZED`

- <a href="#VUID-VkImageMemoryBarrier2-oldLayout-01658"
  id="VUID-VkImageMemoryBarrier2-oldLayout-01658"></a>
  VUID-VkImageMemoryBarrier2-oldLayout-01658  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL` then
  `image` **must** have been created with
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-oldLayout-01659"
  id="VUID-VkImageMemoryBarrier2-oldLayout-01659"></a>
  VUID-VkImageMemoryBarrier2-oldLayout-01659  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` then
  `image` **must** have been created with
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04065"
  id="VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04065"></a>
  VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04065  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL` then `image` **must** have
  been created with at least one of
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`,
  `VK_IMAGE_USAGE_SAMPLED_BIT`, or `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04066"
  id="VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04066"></a>
  VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04066  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` then `image` **must** have
  been created with `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT` set

- <a href="#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04067"
  id="VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04067"></a>
  VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04067  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL` then `image` **must** have
  been created with at least one of
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`,
  `VK_IMAGE_USAGE_SAMPLED_BIT`, or `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04068"
  id="VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04068"></a>
  VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04068  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` then `image` **must**
  have been created with `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
  set

- <a href="#VUID-VkImageMemoryBarrier2-synchronization2-07793"
  id="VUID-VkImageMemoryBarrier2-synchronization2-07793"></a>
  VUID-VkImageMemoryBarrier2-synchronization2-07793  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `oldLayout` **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or
  `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-synchronization2-07794"
  id="VUID-VkImageMemoryBarrier2-synchronization2-07794"></a>
  VUID-VkImageMemoryBarrier2-synchronization2-07794  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `newLayout` **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or
  `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-03938"
  id="VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-03938"></a>
  VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-03938  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL`,
  `image` **must** have been created with
  `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-03939"
  id="VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-03939"></a>
  VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-03939  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL`,
  `image` **must** have been created with at least one of
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`,
  `VK_IMAGE_USAGE_SAMPLED_BIT`, or `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-oldLayout-02088"
  id="VUID-VkImageMemoryBarrier2-oldLayout-02088"></a>
  VUID-VkImageMemoryBarrier2-oldLayout-02088  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_FRAGMENT_SHADING_RATE_ATTACHMENT_OPTIMAL_KHR` then
  `image` **must** have been created with
  `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR` set

- <a href="#VUID-VkImageMemoryBarrier2-image-09117"
  id="VUID-VkImageMemoryBarrier2-image-09117"></a>
  VUID-VkImageMemoryBarrier2-image-09117  
  If `image` was created with a sharing mode of
  `VK_SHARING_MODE_EXCLUSIVE`, and `srcQueueFamilyIndex` and
  `dstQueueFamilyIndex` are not equal, `srcQueueFamilyIndex` **must** be
  `VK_QUEUE_FAMILY_EXTERNAL`, `VK_QUEUE_FAMILY_FOREIGN_EXT`, or a valid
  queue family

- <a href="#VUID-VkImageMemoryBarrier2-image-09118"
  id="VUID-VkImageMemoryBarrier2-image-09118"></a>
  VUID-VkImageMemoryBarrier2-image-09118  
  If `image` was created with a sharing mode of
  `VK_SHARING_MODE_EXCLUSIVE`, and `srcQueueFamilyIndex` and
  `dstQueueFamilyIndex` are not equal, `dstQueueFamilyIndex` **must** be
  `VK_QUEUE_FAMILY_EXTERNAL`, `VK_QUEUE_FAMILY_FOREIGN_EXT`, or a valid
  queue family

- <a href="#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04070"
  id="VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04070"></a>
  VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04070  
  If `srcQueueFamilyIndex` is not equal to `dstQueueFamilyIndex`, at
  least one of `srcQueueFamilyIndex` or `dstQueueFamilyIndex` **must**
  not be `VK_QUEUE_FAMILY_EXTERNAL` or `VK_QUEUE_FAMILY_FOREIGN_EXT`

- <a href="#VUID-VkImageMemoryBarrier2-None-09119"
  id="VUID-VkImageMemoryBarrier2-None-09119"></a>
  VUID-VkImageMemoryBarrier2-None-09119  
  If the [VK_KHR_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory.html) extension
  is not enabled, and the value of
  [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`apiVersion` used to
  create the [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) is not greater than or equal
  to Version 1.1, `srcQueueFamilyIndex` **must** not be
  `VK_QUEUE_FAMILY_EXTERNAL`

- <a href="#VUID-VkImageMemoryBarrier2-None-09120"
  id="VUID-VkImageMemoryBarrier2-None-09120"></a>
  VUID-VkImageMemoryBarrier2-None-09120  
  If the [VK_KHR_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory.html) extension
  is not enabled, and the value of
  [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`apiVersion` used to
  create the [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) is not greater than or equal
  to Version 1.1, `dstQueueFamilyIndex` **must** not be
  `VK_QUEUE_FAMILY_EXTERNAL`

- <a href="#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-09121"
  id="VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-09121"></a>
  VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-09121  
  If the [VK_EXT_queue_family_foreign](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_queue_family_foreign.html)
  extension is not enabled `srcQueueFamilyIndex` **must** not be
  `VK_QUEUE_FAMILY_FOREIGN_EXT`

- <a href="#VUID-VkImageMemoryBarrier2-dstQueueFamilyIndex-09122"
  id="VUID-VkImageMemoryBarrier2-dstQueueFamilyIndex-09122"></a>
  VUID-VkImageMemoryBarrier2-dstQueueFamilyIndex-09122  
  If the [VK_EXT_queue_family_foreign](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_queue_family_foreign.html)
  extension is not enabled `dstQueueFamilyIndex` **must** not be
  `VK_QUEUE_FAMILY_FOREIGN_EXT`

- <a href="#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07120"
  id="VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07120"></a>
  VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07120  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_DECODE_SRC_KHR`
  then `image` **must** have been created with
  `VK_IMAGE_USAGE_VIDEO_DECODE_SRC_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07121"
  id="VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07121"></a>
  VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07121  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_DECODE_DST_KHR`
  then `image` **must** have been created with
  `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07122"
  id="VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07122"></a>
  VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07122  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_DECODE_DPB_KHR`
  then `image` **must** have been created with
  `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07123"
  id="VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07123"></a>
  VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07123  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_ENCODE_SRC_KHR`
  then `image` **must** have been created with
  `VK_IMAGE_USAGE_VIDEO_ENCODE_SRC_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07124"
  id="VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07124"></a>
  VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07124  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_ENCODE_DST_KHR`
  then `image` **must** have been created with
  `VK_IMAGE_USAGE_VIDEO_ENCODE_DST_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07125"
  id="VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07125"></a>
  VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07125  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_ENCODE_DPB_KHR`
  then `image` **must** have been created with
  `VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07006"
  id="VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07006"></a>
  VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07006  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT` then `image`
  **must** have been created with either the
  `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT` usage bits, and the
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_SAMPLED_BIT`
  usage bits, and the `VK_IMAGE_USAGE_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`
  usage bit

- <a href="#VUID-VkImageMemoryBarrier2-attachmentFeedbackLoopLayout-07313"
  id="VUID-VkImageMemoryBarrier2-attachmentFeedbackLoopLayout-07313"></a>
  VUID-VkImageMemoryBarrier2-attachmentFeedbackLoopLayout-07313  
  If the
  [`attachmentFeedbackLoopLayout`](#features-attachmentFeedbackLoopLayout)
  feature is not enabled, `newLayout` **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`

- <a href="#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-09550"
  id="VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-09550"></a>
  VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-09550  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR` then `image` **must** have
  been created with either `VK_IMAGE_USAGE_STORAGE_BIT`, or with both
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` and either of
  `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-dynamicRenderingLocalRead-09551"
  id="VUID-VkImageMemoryBarrier2-dynamicRenderingLocalRead-09551"></a>
  VUID-VkImageMemoryBarrier2-dynamicRenderingLocalRead-09551  
  If the
  [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead)
  feature is not enabled, `oldLayout` **must** not be
  `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR`

- <a href="#VUID-VkImageMemoryBarrier2-dynamicRenderingLocalRead-09552"
  id="VUID-VkImageMemoryBarrier2-dynamicRenderingLocalRead-09552"></a>
  VUID-VkImageMemoryBarrier2-dynamicRenderingLocalRead-09552  
  If the
  [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead)
  feature is not enabled, `newLayout` **must** not be
  `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR`

<!-- -->

- <a href="#VUID-VkImageMemoryBarrier2-subresourceRange-01486"
  id="VUID-VkImageMemoryBarrier2-subresourceRange-01486"></a>
  VUID-VkImageMemoryBarrier2-subresourceRange-01486  
  `subresourceRange.baseMipLevel` **must** be less than the `mipLevels`
  specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image`
  was created

- <a href="#VUID-VkImageMemoryBarrier2-subresourceRange-01724"
  id="VUID-VkImageMemoryBarrier2-subresourceRange-01724"></a>
  VUID-VkImageMemoryBarrier2-subresourceRange-01724  
  If `subresourceRange.levelCount` is not `VK_REMAINING_MIP_LEVELS`,
  `subresourceRange.baseMipLevel` + `subresourceRange.levelCount`
  **must** be less than or equal to the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a href="#VUID-VkImageMemoryBarrier2-subresourceRange-01488"
  id="VUID-VkImageMemoryBarrier2-subresourceRange-01488"></a>
  VUID-VkImageMemoryBarrier2-subresourceRange-01488  
  `subresourceRange.baseArrayLayer` **must** be less than the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `image` was created

- <a href="#VUID-VkImageMemoryBarrier2-subresourceRange-01725"
  id="VUID-VkImageMemoryBarrier2-subresourceRange-01725"></a>
  VUID-VkImageMemoryBarrier2-subresourceRange-01725  
  If `subresourceRange.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `subresourceRange.baseArrayLayer` + `subresourceRange.layerCount`
  **must** be less than or equal to the `arrayLayers` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a href="#VUID-VkImageMemoryBarrier2-image-01932"
  id="VUID-VkImageMemoryBarrier2-image-01932"></a>
  VUID-VkImageMemoryBarrier2-image-01932  
  If `image` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-VkImageMemoryBarrier2-image-09241"
  id="VUID-VkImageMemoryBarrier2-image-09241"></a>
  VUID-VkImageMemoryBarrier2-image-09241  
  If `image` has a color format that is single-plane, then the
  `aspectMask` member of `subresourceRange` **must** be
  `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-image-09242"
  id="VUID-VkImageMemoryBarrier2-image-09242"></a>
  VUID-VkImageMemoryBarrier2-image-09242  
  If `image` has a color format and is not *disjoint*, then the
  `aspectMask` member of `subresourceRange` **must** be
  `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-image-01672"
  id="VUID-VkImageMemoryBarrier2-image-01672"></a>
  VUID-VkImageMemoryBarrier2-image-01672  
  If `image` has a multi-planar format and the image is *disjoint*, then
  the `aspectMask` member of `subresourceRange` **must** include at
  least one [multi-planar aspect mask](#formats-planes-image-aspect) bit
  or `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-image-03320"
  id="VUID-VkImageMemoryBarrier2-image-03320"></a>
  VUID-VkImageMemoryBarrier2-image-03320  
  If `image` has a depth/stencil format with both depth and stencil and
  the
  [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts)
  feature is not enabled, then the `aspectMask` member of
  `subresourceRange` **must** include both `VK_IMAGE_ASPECT_DEPTH_BIT`
  and `VK_IMAGE_ASPECT_STENCIL_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-image-03319"
  id="VUID-VkImageMemoryBarrier2-image-03319"></a>
  VUID-VkImageMemoryBarrier2-image-03319  
  If `image` has a depth/stencil format with both depth and stencil and
  the
  [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts)
  feature is enabled, then the `aspectMask` member of `subresourceRange`
  **must** include either or both `VK_IMAGE_ASPECT_DEPTH_BIT` and
  `VK_IMAGE_ASPECT_STENCIL_BIT`

- <a href="#VUID-VkImageMemoryBarrier2-aspectMask-08702"
  id="VUID-VkImageMemoryBarrier2-aspectMask-08702"></a>
  VUID-VkImageMemoryBarrier2-aspectMask-08702  
  If the `aspectMask` member of `subresourceRange` includes
  `VK_IMAGE_ASPECT_DEPTH_BIT`, `oldLayout` and `newLayout` **must** not
  be one of `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkImageMemoryBarrier2-aspectMask-08703"
  id="VUID-VkImageMemoryBarrier2-aspectMask-08703"></a>
  VUID-VkImageMemoryBarrier2-aspectMask-08703  
  If the `aspectMask` member of `subresourceRange` includes
  `VK_IMAGE_ASPECT_STENCIL_BIT`, `oldLayout` and `newLayout` **must**
  not be one of `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkImageMemoryBarrier2-subresourceRange-09601"
  id="VUID-VkImageMemoryBarrier2-subresourceRange-09601"></a>
  VUID-VkImageMemoryBarrier2-subresourceRange-09601  
  `subresourceRange.aspectMask` **must** be valid for the `format` the
  `image` was created with

- <a href="#VUID-VkImageMemoryBarrier2-srcStageMask-03854"
  id="VUID-VkImageMemoryBarrier2-srcStageMask-03854"></a>
  VUID-VkImageMemoryBarrier2-srcStageMask-03854  
  If either `srcStageMask` or `dstStageMask` includes
  `VK_PIPELINE_STAGE_2_HOST_BIT`, `srcQueueFamilyIndex` and
  `dstQueueFamilyIndex` **must** be equal

- <a href="#VUID-VkImageMemoryBarrier2-srcStageMask-03855"
  id="VUID-VkImageMemoryBarrier2-srcStageMask-03855"></a>
  VUID-VkImageMemoryBarrier2-srcStageMask-03855  
  If `srcStageMask` includes `VK_PIPELINE_STAGE_2_HOST_BIT`, and
  `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers"
  target="_blank" rel="noopener">queue family ownership transfer</a> or
  `oldLayout` and `newLayout` define an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-image-layout-transitions"
  target="_blank" rel="noopener">image layout transition</a>,
  `oldLayout` **must** be one of `VK_IMAGE_LAYOUT_PREINITIALIZED`,
  `VK_IMAGE_LAYOUT_UNDEFINED`, or `VK_IMAGE_LAYOUT_GENERAL`

Valid Usage (Implicit)

- <a href="#VUID-VkImageMemoryBarrier2-sType-sType"
  id="VUID-VkImageMemoryBarrier2-sType-sType"></a>
  VUID-VkImageMemoryBarrier2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER_2`

- <a href="#VUID-VkImageMemoryBarrier2-pNext-pNext"
  id="VUID-VkImageMemoryBarrier2-pNext-pNext"></a>
  VUID-VkImageMemoryBarrier2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkExternalMemoryAcquireUnmodifiedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryAcquireUnmodifiedEXT.html)
  or [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationsInfoEXT.html)

- <a href="#VUID-VkImageMemoryBarrier2-sType-unique"
  id="VUID-VkImageMemoryBarrier2-sType-unique"></a>
  VUID-VkImageMemoryBarrier2-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkImageMemoryBarrier2-srcStageMask-parameter"
  id="VUID-VkImageMemoryBarrier2-srcStageMask-parameter"></a>
  VUID-VkImageMemoryBarrier2-srcStageMask-parameter  
  `srcStageMask` **must** be a valid combination of
  [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html) values

- <a href="#VUID-VkImageMemoryBarrier2-srcAccessMask-parameter"
  id="VUID-VkImageMemoryBarrier2-srcAccessMask-parameter"></a>
  VUID-VkImageMemoryBarrier2-srcAccessMask-parameter  
  `srcAccessMask` **must** be a valid combination of
  [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2.html) values

- <a href="#VUID-VkImageMemoryBarrier2-dstStageMask-parameter"
  id="VUID-VkImageMemoryBarrier2-dstStageMask-parameter"></a>
  VUID-VkImageMemoryBarrier2-dstStageMask-parameter  
  `dstStageMask` **must** be a valid combination of
  [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html) values

- <a href="#VUID-VkImageMemoryBarrier2-dstAccessMask-parameter"
  id="VUID-VkImageMemoryBarrier2-dstAccessMask-parameter"></a>
  VUID-VkImageMemoryBarrier2-dstAccessMask-parameter  
  `dstAccessMask` **must** be a valid combination of
  [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2.html) values

- <a href="#VUID-VkImageMemoryBarrier2-oldLayout-parameter"
  id="VUID-VkImageMemoryBarrier2-oldLayout-parameter"></a>
  VUID-VkImageMemoryBarrier2-oldLayout-parameter  
  `oldLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html)
  value

- <a href="#VUID-VkImageMemoryBarrier2-newLayout-parameter"
  id="VUID-VkImageMemoryBarrier2-newLayout-parameter"></a>
  VUID-VkImageMemoryBarrier2-newLayout-parameter  
  `newLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html)
  value

- <a href="#VUID-VkImageMemoryBarrier2-image-parameter"
  id="VUID-VkImageMemoryBarrier2-image-parameter"></a>
  VUID-VkImageMemoryBarrier2-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkImageMemoryBarrier2-subresourceRange-parameter"
  id="VUID-VkImageMemoryBarrier2-subresourceRange-parameter"></a>
  VUID-VkImageMemoryBarrier2-subresourceRange-parameter  
  `subresourceRange` **must** be a valid
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkAccessFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlags2.html),
[VkDependencyInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyInfo.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html),
[VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageMemoryBarrier2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
