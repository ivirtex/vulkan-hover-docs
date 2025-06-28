# VkBufferMemoryBarrier2(3) Manual Page

## Name

VkBufferMemoryBarrier2 - Structure specifying a buffer memory barrier



## [](#_c_specification)C Specification

The `VkBufferMemoryBarrier2` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkBufferMemoryBarrier2 {
    VkStructureType          sType;
    const void*              pNext;
    VkPipelineStageFlags2    srcStageMask;
    VkAccessFlags2           srcAccessMask;
    VkPipelineStageFlags2    dstStageMask;
    VkAccessFlags2           dstAccessMask;
    uint32_t                 srcQueueFamilyIndex;
    uint32_t                 dstQueueFamilyIndex;
    VkBuffer                 buffer;
    VkDeviceSize             offset;
    VkDeviceSize             size;
} VkBufferMemoryBarrier2;
```

or the equivalent

```c++
// Provided by VK_KHR_synchronization2
typedef VkBufferMemoryBarrier2 VkBufferMemoryBarrier2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcStageMask` is a [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2.html) mask of pipeline stages to be included in the [first synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes).
- `srcAccessMask` is a [VkAccessFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags2.html) mask of access flags to be included in the [first access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes).
- `dstStageMask` is a [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2.html) mask of pipeline stages to be included in the [second synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes).
- `dstAccessMask` is a [VkAccessFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags2.html) mask of access flags to be included in the [second access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes).
- `srcQueueFamilyIndex` is the source queue family for a [queue family ownership transfer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers).
- `dstQueueFamilyIndex` is the destination queue family for a [queue family ownership transfer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers).
- `buffer` is a handle to the buffer whose backing memory is affected by the barrier.
- `offset` is an offset in bytes into the backing memory for `buffer`; this is relative to the base offset as bound to the buffer (see [vkBindBufferMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindBufferMemory.html)).
- `size` is a size in bytes of the affected area of backing memory for `buffer`, or `VK_WHOLE_SIZE` to use the range from `offset` to the end of the buffer.

## [](#_description)Description

This structure defines a [memory dependency](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-memory) limited to a range of a buffer, and **can** define a [queue family ownership transfer operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers) for that range.

The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) and [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) described by this structure include only operations and memory accesses specified by the source stage mask and the source access mask.

The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) and [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) described by this structure include only operations and memory accesses specified by the destination stage mask and the destination access mask.

Both [access scopes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) are limited to only memory accesses to `buffer` in the range defined by `offset` and `size`.

If `buffer` was created with `VK_SHARING_MODE_EXCLUSIVE`, and `srcQueueFamilyIndex` is not equal to `dstQueueFamilyIndex`, this memory barrier defines a [queue family ownership transfer operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers). When executed on a queue in the family identified by `srcQueueFamilyIndex`, this barrier defines a [queue family release operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-release) for the specified buffer range, and if [VkDependencyInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfoKHR.html)::`dependencyFlags` did not include `VK_DEPENDENCY_QUEUE_FAMILY_OWNERSHIP_TRANSFER_USE_ALL_STAGES_BIT_KHR`, the second synchronization scope does not apply to this operation. When executed on a queue in the family identified by `dstQueueFamilyIndex`, this barrier defines a [queue family acquire operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-acquire) for the specified buffer range, and if [VkDependencyInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfoKHR.html)::`dependencyFlags` did not include `VK_DEPENDENCY_QUEUE_FAMILY_OWNERSHIP_TRANSFER_USE_ALL_STAGES_BIT_KHR`, the first synchronization scope does not apply to this operation.

A [queue family ownership transfer operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers) is also defined if the values are not equal, and either is one of the special queue family values reserved for external memory ownership transfers, as described in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers). A [queue family release operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-release) is defined when `dstQueueFamilyIndex` is one of those values, and a [queue family acquire operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-acquire) is defined when `srcQueueFamilyIndex` is one of those values.

Valid Usage

- [](#VUID-VkBufferMemoryBarrier2-srcStageMask-03929)VUID-VkBufferMemoryBarrier2-srcStageMask-03929  
  If the [`geometryShader`](#features-geometryShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcStageMask-03930)VUID-VkBufferMemoryBarrier2-srcStageMask-03930  
  If the [`tessellationShader`](#features-tessellationShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT` or `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcStageMask-03931)VUID-VkBufferMemoryBarrier2-srcStageMask-03931  
  If the [`conditionalRendering`](#features-conditionalRendering) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`
- [](#VUID-VkBufferMemoryBarrier2-srcStageMask-03932)VUID-VkBufferMemoryBarrier2-srcStageMask-03932  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- [](#VUID-VkBufferMemoryBarrier2-srcStageMask-03933)VUID-VkBufferMemoryBarrier2-srcStageMask-03933  
  If the [`transformFeedback`](#features-transformFeedback) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`
- [](#VUID-VkBufferMemoryBarrier2-srcStageMask-03934)VUID-VkBufferMemoryBarrier2-srcStageMask-03934  
  If the [`meshShader`](#features-meshShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`
- [](#VUID-VkBufferMemoryBarrier2-srcStageMask-03935)VUID-VkBufferMemoryBarrier2-srcStageMask-03935  
  If the [`taskShader`](#features-taskShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`
- [](#VUID-VkBufferMemoryBarrier2-srcStageMask-07316)VUID-VkBufferMemoryBarrier2-srcStageMask-07316  
  If neither of the [`shadingRateImage`](#features-shadingRateImage) or the [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate) features are enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-srcStageMask-04957)VUID-VkBufferMemoryBarrier2-srcStageMask-04957  
  If the [`subpassShading`](#features-subpassShading) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`
- [](#VUID-VkBufferMemoryBarrier2-srcStageMask-04995)VUID-VkBufferMemoryBarrier2-srcStageMask-04995  
  If the [`invocationMask`](#features-invocationMask) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`
- [](#VUID-VkBufferMemoryBarrier2-srcStageMask-07946)VUID-VkBufferMemoryBarrier2-srcStageMask-07946  
  If neither the [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) extension or the [`rayTracingPipeline`](#features-rayTracingPipeline) feature are enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-srcStageMask-10751)VUID-VkBufferMemoryBarrier2-srcStageMask-10751  
  If the [`accelerationStructure`](#features-accelerationStructure) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-srcStageMask-10752)VUID-VkBufferMemoryBarrier2-srcStageMask-10752  
  If the [`rayTracingMaintenance1`](#features-rayTracingMaintenance1) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-srcStageMask-10753)VUID-VkBufferMemoryBarrier2-srcStageMask-10753  
  If the [`micromap`](#features-micromap) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`

<!--THE END-->

- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03900)VUID-VkBufferMemoryBarrier2-srcAccessMask-03900  
  If `srcAccessMask` includes `VK_ACCESS_2_INDIRECT_COMMAND_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03901)VUID-VkBufferMemoryBarrier2-srcAccessMask-03901  
  If `srcAccessMask` includes `VK_ACCESS_2_INDEX_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_INDEX_INPUT_BIT`, `VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03902)VUID-VkBufferMemoryBarrier2-srcAccessMask-03902  
  If `srcAccessMask` includes `VK_ACCESS_2_VERTEX_ATTRIBUTE_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_VERTEX_ATTRIBUTE_INPUT_BIT`, `VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03903)VUID-VkBufferMemoryBarrier2-srcAccessMask-03903  
  If `srcAccessMask` includes `VK_ACCESS_2_INPUT_ATTACHMENT_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`, `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03904)VUID-VkBufferMemoryBarrier2-srcAccessMask-03904  
  If `srcAccessMask` includes `VK_ACCESS_2_UNIFORM_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03905)VUID-VkBufferMemoryBarrier2-srcAccessMask-03905  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_SAMPLED_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03906)VUID-VkBufferMemoryBarrier2-srcAccessMask-03906  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_STORAGE_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03907)VUID-VkBufferMemoryBarrier2-srcAccessMask-03907  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-07454)VUID-VkBufferMemoryBarrier2-srcAccessMask-07454  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03909)VUID-VkBufferMemoryBarrier2-srcAccessMask-03909  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_WRITE_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03910)VUID-VkBufferMemoryBarrier2-srcAccessMask-03910  
  If `srcAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT` `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03911)VUID-VkBufferMemoryBarrier2-srcAccessMask-03911  
  If `srcAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT` `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03912)VUID-VkBufferMemoryBarrier2-srcAccessMask-03912  
  If `srcAccessMask` includes `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT`, `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03913)VUID-VkBufferMemoryBarrier2-srcAccessMask-03913  
  If `srcAccessMask` includes `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT`, `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03914)VUID-VkBufferMemoryBarrier2-srcAccessMask-03914  
  If `srcAccessMask` includes `VK_ACCESS_2_TRANSFER_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_COPY_BIT`, `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_RESOLVE_BIT`, `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_CONVERT_COOPERATIVE_VECTOR_MATRIX_BIT_NV`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03915)VUID-VkBufferMemoryBarrier2-srcAccessMask-03915  
  If `srcAccessMask` includes `VK_ACCESS_2_TRANSFER_WRITE_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_COPY_BIT`, `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_RESOLVE_BIT`, `VK_PIPELINE_STAGE_2_CLEAR_BIT`, `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_CONVERT_COOPERATIVE_VECTOR_MATRIX_BIT_NV`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03916)VUID-VkBufferMemoryBarrier2-srcAccessMask-03916  
  If `srcAccessMask` includes `VK_ACCESS_2_HOST_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_HOST_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03917)VUID-VkBufferMemoryBarrier2-srcAccessMask-03917  
  If `srcAccessMask` includes `VK_ACCESS_2_HOST_WRITE_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_HOST_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03918)VUID-VkBufferMemoryBarrier2-srcAccessMask-03918  
  If `srcAccessMask` includes `VK_ACCESS_2_CONDITIONAL_RENDERING_READ_BIT_EXT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03919)VUID-VkBufferMemoryBarrier2-srcAccessMask-03919  
  If `srcAccessMask` includes `VK_ACCESS_2_FRAGMENT_DENSITY_MAP_READ_BIT_EXT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03920)VUID-VkBufferMemoryBarrier2-srcAccessMask-03920  
  If `srcAccessMask` includes `VK_ACCESS_2_TRANSFORM_FEEDBACK_WRITE_BIT_EXT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-04747)VUID-VkBufferMemoryBarrier2-srcAccessMask-04747  
  If `srcAccessMask` includes `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT`, `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03922)VUID-VkBufferMemoryBarrier2-srcAccessMask-03922  
  If `srcAccessMask` includes `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03923)VUID-VkBufferMemoryBarrier2-srcAccessMask-03923  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADING_RATE_IMAGE_READ_BIT_NV`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_SHADING_RATE_IMAGE_BIT_NV`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-04994)VUID-VkBufferMemoryBarrier2-srcAccessMask-04994  
  If `srcAccessMask` includes `VK_ACCESS_2_INVOCATION_MASK_READ_BIT_HUAWEI`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03924)VUID-VkBufferMemoryBarrier2-srcAccessMask-03924  
  If `srcAccessMask` includes `VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_NV`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03925)VUID-VkBufferMemoryBarrier2-srcAccessMask-03925  
  If `srcAccessMask` includes `VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_NV`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03926)VUID-VkBufferMemoryBarrier2-srcAccessMask-03926  
  If `srcAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT` `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03927)VUID-VkBufferMemoryBarrier2-srcAccessMask-03927  
  If `srcAccessMask` includes `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-03928)VUID-VkBufferMemoryBarrier2-srcAccessMask-03928  
  If `srcAccessMask` includes `VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-06256)VUID-VkBufferMemoryBarrier2-srcAccessMask-06256  
  If the [`rayQuery`](#features-rayQuery) feature is not enabled and `srcAccessMask` includes `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `srcStageMask` **must** not include any of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages except `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-07272)VUID-VkBufferMemoryBarrier2-srcAccessMask-07272  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_BINDING_TABLE_READ_BIT_KHR`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT` or `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-04858)VUID-VkBufferMemoryBarrier2-srcAccessMask-04858  
  If `srcAccessMask` includes `VK_ACCESS_2_VIDEO_DECODE_READ_BIT_KHR`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-04859)VUID-VkBufferMemoryBarrier2-srcAccessMask-04859  
  If `srcAccessMask` includes `VK_ACCESS_2_VIDEO_DECODE_WRITE_BIT_KHR`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-04860)VUID-VkBufferMemoryBarrier2-srcAccessMask-04860  
  If `srcAccessMask` includes `VK_ACCESS_2_VIDEO_ENCODE_READ_BIT_KHR`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-04861)VUID-VkBufferMemoryBarrier2-srcAccessMask-04861  
  If `srcAccessMask` includes `VK_ACCESS_2_VIDEO_ENCODE_WRITE_BIT_KHR`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-07455)VUID-VkBufferMemoryBarrier2-srcAccessMask-07455  
  If `srcAccessMask` includes `VK_ACCESS_2_OPTICAL_FLOW_READ_BIT_NV`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-07456)VUID-VkBufferMemoryBarrier2-srcAccessMask-07456  
  If `srcAccessMask` includes `VK_ACCESS_2_OPTICAL_FLOW_WRITE_BIT_NV`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-07457)VUID-VkBufferMemoryBarrier2-srcAccessMask-07457  
  If `srcAccessMask` includes `VK_ACCESS_2_MICROMAP_WRITE_BIT_EXT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-07458)VUID-VkBufferMemoryBarrier2-srcAccessMask-07458  
  If `srcAccessMask` includes `VK_ACCESS_2_MICROMAP_READ_BIT_EXT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` or `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-08118)VUID-VkBufferMemoryBarrier2-srcAccessMask-08118  
  If `srcAccessMask` includes `VK_ACCESS_2_DESCRIPTOR_BUFFER_READ_BIT_EXT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-10670)VUID-VkBufferMemoryBarrier2-srcAccessMask-10670  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_TILE_ATTACHMENT_READ_BIT_QCOM`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT` or `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT`
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-10671)VUID-VkBufferMemoryBarrier2-srcAccessMask-10671  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_TILE_ATTACHMENT_WRITE_BIT_QCOM`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT` or `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT`

<!--THE END-->

- [](#VUID-VkBufferMemoryBarrier2-dstStageMask-03929)VUID-VkBufferMemoryBarrier2-dstStageMask-03929  
  If the [`geometryShader`](#features-geometryShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstStageMask-03930)VUID-VkBufferMemoryBarrier2-dstStageMask-03930  
  If the [`tessellationShader`](#features-tessellationShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT` or `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstStageMask-03931)VUID-VkBufferMemoryBarrier2-dstStageMask-03931  
  If the [`conditionalRendering`](#features-conditionalRendering) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`
- [](#VUID-VkBufferMemoryBarrier2-dstStageMask-03932)VUID-VkBufferMemoryBarrier2-dstStageMask-03932  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- [](#VUID-VkBufferMemoryBarrier2-dstStageMask-03933)VUID-VkBufferMemoryBarrier2-dstStageMask-03933  
  If the [`transformFeedback`](#features-transformFeedback) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`
- [](#VUID-VkBufferMemoryBarrier2-dstStageMask-03934)VUID-VkBufferMemoryBarrier2-dstStageMask-03934  
  If the [`meshShader`](#features-meshShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`
- [](#VUID-VkBufferMemoryBarrier2-dstStageMask-03935)VUID-VkBufferMemoryBarrier2-dstStageMask-03935  
  If the [`taskShader`](#features-taskShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`
- [](#VUID-VkBufferMemoryBarrier2-dstStageMask-07316)VUID-VkBufferMemoryBarrier2-dstStageMask-07316  
  If neither of the [`shadingRateImage`](#features-shadingRateImage) or the [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate) features are enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-dstStageMask-04957)VUID-VkBufferMemoryBarrier2-dstStageMask-04957  
  If the [`subpassShading`](#features-subpassShading) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`
- [](#VUID-VkBufferMemoryBarrier2-dstStageMask-04995)VUID-VkBufferMemoryBarrier2-dstStageMask-04995  
  If the [`invocationMask`](#features-invocationMask) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`
- [](#VUID-VkBufferMemoryBarrier2-dstStageMask-07946)VUID-VkBufferMemoryBarrier2-dstStageMask-07946  
  If neither the [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) extension or the [`rayTracingPipeline`](#features-rayTracingPipeline) feature are enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-dstStageMask-10751)VUID-VkBufferMemoryBarrier2-dstStageMask-10751  
  If the [`accelerationStructure`](#features-accelerationStructure) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-dstStageMask-10752)VUID-VkBufferMemoryBarrier2-dstStageMask-10752  
  If the [`rayTracingMaintenance1`](#features-rayTracingMaintenance1) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-dstStageMask-10753)VUID-VkBufferMemoryBarrier2-dstStageMask-10753  
  If the [`micromap`](#features-micromap) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`

<!--THE END-->

- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03900)VUID-VkBufferMemoryBarrier2-dstAccessMask-03900  
  If `dstAccessMask` includes `VK_ACCESS_2_INDIRECT_COMMAND_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03901)VUID-VkBufferMemoryBarrier2-dstAccessMask-03901  
  If `dstAccessMask` includes `VK_ACCESS_2_INDEX_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_INDEX_INPUT_BIT`, `VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03902)VUID-VkBufferMemoryBarrier2-dstAccessMask-03902  
  If `dstAccessMask` includes `VK_ACCESS_2_VERTEX_ATTRIBUTE_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_VERTEX_ATTRIBUTE_INPUT_BIT`, `VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03903)VUID-VkBufferMemoryBarrier2-dstAccessMask-03903  
  If `dstAccessMask` includes `VK_ACCESS_2_INPUT_ATTACHMENT_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`, `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03904)VUID-VkBufferMemoryBarrier2-dstAccessMask-03904  
  If `dstAccessMask` includes `VK_ACCESS_2_UNIFORM_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03905)VUID-VkBufferMemoryBarrier2-dstAccessMask-03905  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_SAMPLED_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03906)VUID-VkBufferMemoryBarrier2-dstAccessMask-03906  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_STORAGE_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03907)VUID-VkBufferMemoryBarrier2-dstAccessMask-03907  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-07454)VUID-VkBufferMemoryBarrier2-dstAccessMask-07454  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03909)VUID-VkBufferMemoryBarrier2-dstAccessMask-03909  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_WRITE_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03910)VUID-VkBufferMemoryBarrier2-dstAccessMask-03910  
  If `dstAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT` `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03911)VUID-VkBufferMemoryBarrier2-dstAccessMask-03911  
  If `dstAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT` `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03912)VUID-VkBufferMemoryBarrier2-dstAccessMask-03912  
  If `dstAccessMask` includes `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT`, `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03913)VUID-VkBufferMemoryBarrier2-dstAccessMask-03913  
  If `dstAccessMask` includes `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT`, `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03914)VUID-VkBufferMemoryBarrier2-dstAccessMask-03914  
  If `dstAccessMask` includes `VK_ACCESS_2_TRANSFER_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_COPY_BIT`, `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_RESOLVE_BIT`, `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_CONVERT_COOPERATIVE_VECTOR_MATRIX_BIT_NV`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03915)VUID-VkBufferMemoryBarrier2-dstAccessMask-03915  
  If `dstAccessMask` includes `VK_ACCESS_2_TRANSFER_WRITE_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_COPY_BIT`, `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_RESOLVE_BIT`, `VK_PIPELINE_STAGE_2_CLEAR_BIT`, `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_CONVERT_COOPERATIVE_VECTOR_MATRIX_BIT_NV`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03916)VUID-VkBufferMemoryBarrier2-dstAccessMask-03916  
  If `dstAccessMask` includes `VK_ACCESS_2_HOST_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_HOST_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03917)VUID-VkBufferMemoryBarrier2-dstAccessMask-03917  
  If `dstAccessMask` includes `VK_ACCESS_2_HOST_WRITE_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_HOST_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03918)VUID-VkBufferMemoryBarrier2-dstAccessMask-03918  
  If `dstAccessMask` includes `VK_ACCESS_2_CONDITIONAL_RENDERING_READ_BIT_EXT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03919)VUID-VkBufferMemoryBarrier2-dstAccessMask-03919  
  If `dstAccessMask` includes `VK_ACCESS_2_FRAGMENT_DENSITY_MAP_READ_BIT_EXT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03920)VUID-VkBufferMemoryBarrier2-dstAccessMask-03920  
  If `dstAccessMask` includes `VK_ACCESS_2_TRANSFORM_FEEDBACK_WRITE_BIT_EXT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-04747)VUID-VkBufferMemoryBarrier2-dstAccessMask-04747  
  If `dstAccessMask` includes `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT`, `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03922)VUID-VkBufferMemoryBarrier2-dstAccessMask-03922  
  If `dstAccessMask` includes `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03923)VUID-VkBufferMemoryBarrier2-dstAccessMask-03923  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADING_RATE_IMAGE_READ_BIT_NV`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_SHADING_RATE_IMAGE_BIT_NV`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-04994)VUID-VkBufferMemoryBarrier2-dstAccessMask-04994  
  If `dstAccessMask` includes `VK_ACCESS_2_INVOCATION_MASK_READ_BIT_HUAWEI`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03924)VUID-VkBufferMemoryBarrier2-dstAccessMask-03924  
  If `dstAccessMask` includes `VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_NV`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03925)VUID-VkBufferMemoryBarrier2-dstAccessMask-03925  
  If `dstAccessMask` includes `VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_NV`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03926)VUID-VkBufferMemoryBarrier2-dstAccessMask-03926  
  If `dstAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT` `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03927)VUID-VkBufferMemoryBarrier2-dstAccessMask-03927  
  If `dstAccessMask` includes `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-03928)VUID-VkBufferMemoryBarrier2-dstAccessMask-03928  
  If `dstAccessMask` includes `VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-06256)VUID-VkBufferMemoryBarrier2-dstAccessMask-06256  
  If the [`rayQuery`](#features-rayQuery) feature is not enabled and `dstAccessMask` includes `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `dstStageMask` **must** not include any of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages except `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-07272)VUID-VkBufferMemoryBarrier2-dstAccessMask-07272  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_BINDING_TABLE_READ_BIT_KHR`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT` or `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-04858)VUID-VkBufferMemoryBarrier2-dstAccessMask-04858  
  If `dstAccessMask` includes `VK_ACCESS_2_VIDEO_DECODE_READ_BIT_KHR`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-04859)VUID-VkBufferMemoryBarrier2-dstAccessMask-04859  
  If `dstAccessMask` includes `VK_ACCESS_2_VIDEO_DECODE_WRITE_BIT_KHR`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-04860)VUID-VkBufferMemoryBarrier2-dstAccessMask-04860  
  If `dstAccessMask` includes `VK_ACCESS_2_VIDEO_ENCODE_READ_BIT_KHR`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-04861)VUID-VkBufferMemoryBarrier2-dstAccessMask-04861  
  If `dstAccessMask` includes `VK_ACCESS_2_VIDEO_ENCODE_WRITE_BIT_KHR`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-07455)VUID-VkBufferMemoryBarrier2-dstAccessMask-07455  
  If `dstAccessMask` includes `VK_ACCESS_2_OPTICAL_FLOW_READ_BIT_NV`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-07456)VUID-VkBufferMemoryBarrier2-dstAccessMask-07456  
  If `dstAccessMask` includes `VK_ACCESS_2_OPTICAL_FLOW_WRITE_BIT_NV`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-07457)VUID-VkBufferMemoryBarrier2-dstAccessMask-07457  
  If `dstAccessMask` includes `VK_ACCESS_2_MICROMAP_WRITE_BIT_EXT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-07458)VUID-VkBufferMemoryBarrier2-dstAccessMask-07458  
  If `dstAccessMask` includes `VK_ACCESS_2_MICROMAP_READ_BIT_EXT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` or `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-08118)VUID-VkBufferMemoryBarrier2-dstAccessMask-08118  
  If `dstAccessMask` includes `VK_ACCESS_2_DESCRIPTOR_BUFFER_READ_BIT_EXT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-10670)VUID-VkBufferMemoryBarrier2-dstAccessMask-10670  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_TILE_ATTACHMENT_READ_BIT_QCOM`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT` or `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT`
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-10671)VUID-VkBufferMemoryBarrier2-dstAccessMask-10671  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_TILE_ATTACHMENT_WRITE_BIT_QCOM`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT` or `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT`

<!--THE END-->

- [](#VUID-VkBufferMemoryBarrier2-offset-01187)VUID-VkBufferMemoryBarrier2-offset-01187  
  `offset` **must** be less than the size of `buffer`
- [](#VUID-VkBufferMemoryBarrier2-size-01188)VUID-VkBufferMemoryBarrier2-size-01188  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be greater than `0`
- [](#VUID-VkBufferMemoryBarrier2-size-01189)VUID-VkBufferMemoryBarrier2-size-01189  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be less than or equal to than the size of `buffer` minus `offset`
- [](#VUID-VkBufferMemoryBarrier2-buffer-01931)VUID-VkBufferMemoryBarrier2-buffer-01931  
  If `buffer` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkBufferMemoryBarrier2-buffer-09095)VUID-VkBufferMemoryBarrier2-buffer-09095  
  If `buffer` was created with a sharing mode of `VK_SHARING_MODE_EXCLUSIVE`, and `srcQueueFamilyIndex` and `dstQueueFamilyIndex` are not equal, `srcQueueFamilyIndex` **must** be `VK_QUEUE_FAMILY_EXTERNAL`, `VK_QUEUE_FAMILY_FOREIGN_EXT`, or a valid queue family
- [](#VUID-VkBufferMemoryBarrier2-buffer-09096)VUID-VkBufferMemoryBarrier2-buffer-09096  
  If `buffer` was created with a sharing mode of `VK_SHARING_MODE_EXCLUSIVE`, and `srcQueueFamilyIndex` and `dstQueueFamilyIndex` are not equal, `dstQueueFamilyIndex` **must** be `VK_QUEUE_FAMILY_EXTERNAL`, `VK_QUEUE_FAMILY_FOREIGN_EXT`, or a valid queue family
- [](#VUID-VkBufferMemoryBarrier2-None-09097)VUID-VkBufferMemoryBarrier2-None-09097  
  If the [VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html) extension is not enabled, and the value of [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`apiVersion` used to create the [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) is not greater than or equal to Version 1.1, `srcQueueFamilyIndex` **must** not be `VK_QUEUE_FAMILY_EXTERNAL`
- [](#VUID-VkBufferMemoryBarrier2-None-09098)VUID-VkBufferMemoryBarrier2-None-09098  
  If the [VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html) extension is not enabled, and the value of [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`apiVersion` used to create the [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) is not greater than or equal to Version 1.1, `dstQueueFamilyIndex` **must** not be `VK_QUEUE_FAMILY_EXTERNAL`
- [](#VUID-VkBufferMemoryBarrier2-srcQueueFamilyIndex-09099)VUID-VkBufferMemoryBarrier2-srcQueueFamilyIndex-09099  
  If the [VK\_EXT\_queue\_family\_foreign](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_queue_family_foreign.html) extension is not enabled `srcQueueFamilyIndex` **must** not be `VK_QUEUE_FAMILY_FOREIGN_EXT`
- [](#VUID-VkBufferMemoryBarrier2-dstQueueFamilyIndex-09100)VUID-VkBufferMemoryBarrier2-dstQueueFamilyIndex-09100  
  If the [VK\_EXT\_queue\_family\_foreign](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_queue_family_foreign.html) extension is not enabled `dstQueueFamilyIndex` **must** not be `VK_QUEUE_FAMILY_FOREIGN_EXT`
- [](#VUID-VkBufferMemoryBarrier2-srcStageMask-03851)VUID-VkBufferMemoryBarrier2-srcStageMask-03851  
  If either `srcStageMask` or `dstStageMask` includes `VK_PIPELINE_STAGE_2_HOST_BIT`, `srcQueueFamilyIndex` and `dstQueueFamilyIndex` **must** be equal

Valid Usage (Implicit)

- [](#VUID-VkBufferMemoryBarrier2-sType-sType)VUID-VkBufferMemoryBarrier2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER_2`
- [](#VUID-VkBufferMemoryBarrier2-pNext-pNext)VUID-VkBufferMemoryBarrier2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkExternalMemoryAcquireUnmodifiedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryAcquireUnmodifiedEXT.html) or [VkMemoryBarrierAccessFlags3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrierAccessFlags3KHR.html)
- [](#VUID-VkBufferMemoryBarrier2-sType-unique)VUID-VkBufferMemoryBarrier2-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkBufferMemoryBarrier2-srcStageMask-parameter)VUID-VkBufferMemoryBarrier2-srcStageMask-parameter  
  `srcStageMask` **must** be a valid combination of [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html) values
- [](#VUID-VkBufferMemoryBarrier2-srcAccessMask-parameter)VUID-VkBufferMemoryBarrier2-srcAccessMask-parameter  
  `srcAccessMask` **must** be a valid combination of [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html) values
- [](#VUID-VkBufferMemoryBarrier2-dstStageMask-parameter)VUID-VkBufferMemoryBarrier2-dstStageMask-parameter  
  `dstStageMask` **must** be a valid combination of [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html) values
- [](#VUID-VkBufferMemoryBarrier2-dstAccessMask-parameter)VUID-VkBufferMemoryBarrier2-dstAccessMask-parameter  
  `dstAccessMask` **must** be a valid combination of [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html) values
- [](#VUID-VkBufferMemoryBarrier2-buffer-parameter)VUID-VkBufferMemoryBarrier2-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle

## [](#_see_also)See Also

[VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkAccessFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags2.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfo.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferMemoryBarrier2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0