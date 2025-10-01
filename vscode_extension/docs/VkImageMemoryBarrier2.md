# VkImageMemoryBarrier2(3) Manual Page

## Name

VkImageMemoryBarrier2 - Structure specifying an image memory barrier



## [](#_c_specification)C Specification

The `VkImageMemoryBarrier2` structure is defined as:

```c++
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

```c++
// Provided by VK_KHR_synchronization2
typedef VkImageMemoryBarrier2 VkImageMemoryBarrier2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcStageMask` is a [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2.html) mask of pipeline stages to be included in the [first synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes).
- `srcAccessMask` is a [VkAccessFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags2.html) mask of access flags to be included in the [first access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes).
- `dstStageMask` is a [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2.html) mask of pipeline stages to be included in the [second synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes).
- `dstAccessMask` is a [VkAccessFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags2.html) mask of access flags to be included in the [second access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes).
- `oldLayout` is the old layout in an [image layout transition](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-image-layout-transitions).
- `newLayout` is the new layout in an [image layout transition](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-image-layout-transitions).
- `srcQueueFamilyIndex` is the source queue family for a [queue family ownership transfer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers).
- `dstQueueFamilyIndex` is the destination queue family for a [queue family ownership transfer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers).
- `image` is a handle to the image affected by this barrier.
- `subresourceRange` describes the [image subresource range](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views) within `image` that is affected by this barrier.

## [](#_description)Description

This structure defines a [memory dependency](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-memory) limited to an image subresource range, and **can** define a [queue family ownership transfer operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers) and [image layout transition](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-image-layout-transitions) for that subresource range.

The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) and [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) described by this structure include only operations and memory accesses specified by the source stage mask and the source access mask.

The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) and [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) described by this structure include only operations and memory accesses specified by the destination stage mask and the destination access mask.

Both [access scopes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) are limited to only memory accesses to `image` in the subresource range defined by `subresourceRange`.

If `image` was created with `VK_SHARING_MODE_EXCLUSIVE`, and `srcQueueFamilyIndex` is not equal to `dstQueueFamilyIndex`, this memory barrier defines a [queue family ownership transfer operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers). When executed on a queue in the family identified by `srcQueueFamilyIndex`, this barrier defines a [queue family release operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-release) for the specified image subresource range, and if [VkDependencyInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfoKHR.html)::`dependencyFlags` did not include `VK_DEPENDENCY_QUEUE_FAMILY_OWNERSHIP_TRANSFER_USE_ALL_STAGES_BIT_KHR`, the second synchronization scope does not apply to this operation. When executed on a queue in the family identified by `dstQueueFamilyIndex`, this barrier defines a [queue family acquire operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-acquire) for the specified image subresource range, and if [VkDependencyInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfoKHR.html)::`dependencyFlags` did not include `VK_DEPENDENCY_QUEUE_FAMILY_OWNERSHIP_TRANSFER_USE_ALL_STAGES_BIT_KHR`, the first synchronization scope does not apply to this operation.

A [queue family ownership transfer operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers) is also defined if the values are not equal, and either is one of the special queue family values reserved for external memory ownership transfers, as described in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers). A [queue family release operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-release) is defined when `dstQueueFamilyIndex` is one of those values, and a [queue family acquire operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-acquire) is defined when `srcQueueFamilyIndex` is one of those values.

If `oldLayout` is not equal to `newLayout`, then the memory barrier defines an [image layout transition](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-image-layout-transitions) for the specified image subresource range. If this memory barrier defines a [queue family ownership transfer operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers), the layout transition is only executed once between the queues.

Note

When the old and new layout are equal, the layout values are ignored - data is preserved no matter what values are specified, or what layout the image is currently in.

If `image` is a 3D image created with `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` and the [`maintenance9`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance9) feature is enabled, the `baseArrayLayer` and `layerCount` members of `subresourceRange` specify the subset of slices of the 3D image affected by the memory barrier, including the layout transition. Any slices of a 3D image not included in `subresourceRange` are not affected by the memory barrier and remain in their existing layout.

Note

Enabling the [`maintenance9`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance9) feature modifies the behavior of image barriers targeting 3D images created with `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT`. Previously, a `layerCount` equal to 1 would cover the entire 3D image, but this has a different meaning when the `maintenance9` feature is enabled. Linking this behavioral change solely to the `maintenance9` feature caused an unintended break in forward-compatibility. Validation layers are expected to flag a warning for the scenario where the `maintenance9` feature is not enabled, and the application uses `layerCount` equal to 1 on this kind of 3D image. `layerCount` can be set to `VK_REMAINING_ARRAY_LAYERS` instead, which has the same semantics with or without the extension. This validation check should make it feasible for software to avoid any breaking changes should the `maintenance9` feature be enabled in the future, either explicitly by application or by a layer outside the control of the application.

If `image` has a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar) and the image is *disjoint*, then including `VK_IMAGE_ASPECT_COLOR_BIT` in the `aspectMask` member of `subresourceRange` is equivalent to including `VK_IMAGE_ASPECT_PLANE_0_BIT`, `VK_IMAGE_ASPECT_PLANE_1_BIT`, and (for three-plane formats only) `VK_IMAGE_ASPECT_PLANE_2_BIT`.

Valid Usage

- [](#VUID-VkImageMemoryBarrier2-srcStageMask-03929)VUID-VkImageMemoryBarrier2-srcStageMask-03929  
  If the [`geometryShader`](#features-geometryShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcStageMask-03930)VUID-VkImageMemoryBarrier2-srcStageMask-03930  
  If the [`tessellationShader`](#features-tessellationShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT` or `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcStageMask-03931)VUID-VkImageMemoryBarrier2-srcStageMask-03931  
  If the [`conditionalRendering`](#features-conditionalRendering) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`
- [](#VUID-VkImageMemoryBarrier2-srcStageMask-03932)VUID-VkImageMemoryBarrier2-srcStageMask-03932  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- [](#VUID-VkImageMemoryBarrier2-srcStageMask-03933)VUID-VkImageMemoryBarrier2-srcStageMask-03933  
  If the [`transformFeedback`](#features-transformFeedback) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`
- [](#VUID-VkImageMemoryBarrier2-srcStageMask-03934)VUID-VkImageMemoryBarrier2-srcStageMask-03934  
  If the [`meshShader`](#features-meshShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`
- [](#VUID-VkImageMemoryBarrier2-srcStageMask-03935)VUID-VkImageMemoryBarrier2-srcStageMask-03935  
  If the [`taskShader`](#features-taskShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`
- [](#VUID-VkImageMemoryBarrier2-srcStageMask-07316)VUID-VkImageMemoryBarrier2-srcStageMask-07316  
  If neither of the [`shadingRateImage`](#features-shadingRateImage) or the [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate) features are enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcStageMask-04957)VUID-VkImageMemoryBarrier2-srcStageMask-04957  
  If the [`subpassShading`](#features-subpassShading) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`
- [](#VUID-VkImageMemoryBarrier2-srcStageMask-04995)VUID-VkImageMemoryBarrier2-srcStageMask-04995  
  If the [`invocationMask`](#features-invocationMask) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`
- [](#VUID-VkImageMemoryBarrier2-srcStageMask-07946)VUID-VkImageMemoryBarrier2-srcStageMask-07946  
  If neither the [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) extension or the [`rayTracingPipeline`](#features-rayTracingPipeline) feature are enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcStageMask-10751)VUID-VkImageMemoryBarrier2-srcStageMask-10751  
  If the [`accelerationStructure`](#features-accelerationStructure) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcStageMask-10752)VUID-VkImageMemoryBarrier2-srcStageMask-10752  
  If the [`rayTracingMaintenance1`](#features-rayTracingMaintenance1) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcStageMask-10753)VUID-VkImageMemoryBarrier2-srcStageMask-10753  
  If the [`micromap`](#features-micromap) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`

<!--THE END-->

- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03900)VUID-VkImageMemoryBarrier2-srcAccessMask-03900  
  If `srcAccessMask` includes `VK_ACCESS_2_INDIRECT_COMMAND_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_COPY_INDIRECT_BIT_KHR`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03901)VUID-VkImageMemoryBarrier2-srcAccessMask-03901  
  If `srcAccessMask` includes `VK_ACCESS_2_INDEX_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_INDEX_INPUT_BIT`, `VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03902)VUID-VkImageMemoryBarrier2-srcAccessMask-03902  
  If `srcAccessMask` includes `VK_ACCESS_2_VERTEX_ATTRIBUTE_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_VERTEX_ATTRIBUTE_INPUT_BIT`, `VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03903)VUID-VkImageMemoryBarrier2-srcAccessMask-03903  
  If `srcAccessMask` includes `VK_ACCESS_2_INPUT_ATTACHMENT_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`, `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03904)VUID-VkImageMemoryBarrier2-srcAccessMask-03904  
  If `srcAccessMask` includes `VK_ACCESS_2_UNIFORM_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03905)VUID-VkImageMemoryBarrier2-srcAccessMask-03905  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_SAMPLED_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03906)VUID-VkImageMemoryBarrier2-srcAccessMask-03906  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_STORAGE_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03907)VUID-VkImageMemoryBarrier2-srcAccessMask-03907  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-07454)VUID-VkImageMemoryBarrier2-srcAccessMask-07454  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03909)VUID-VkImageMemoryBarrier2-srcAccessMask-03909  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_WRITE_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03910)VUID-VkImageMemoryBarrier2-srcAccessMask-03910  
  If `srcAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT` `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03911)VUID-VkImageMemoryBarrier2-srcAccessMask-03911  
  If `srcAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT` `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03912)VUID-VkImageMemoryBarrier2-srcAccessMask-03912  
  If `srcAccessMask` includes `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT`, `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03913)VUID-VkImageMemoryBarrier2-srcAccessMask-03913  
  If `srcAccessMask` includes `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT`, `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03914)VUID-VkImageMemoryBarrier2-srcAccessMask-03914  
  If `srcAccessMask` includes `VK_ACCESS_2_TRANSFER_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_COPY_BIT`, `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_RESOLVE_BIT`, `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_CONVERT_COOPERATIVE_VECTOR_MATRIX_BIT_NV`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03915)VUID-VkImageMemoryBarrier2-srcAccessMask-03915  
  If `srcAccessMask` includes `VK_ACCESS_2_TRANSFER_WRITE_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_COPY_BIT`, `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_RESOLVE_BIT`, `VK_PIPELINE_STAGE_2_CLEAR_BIT`, `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_CONVERT_COOPERATIVE_VECTOR_MATRIX_BIT_NV`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03916)VUID-VkImageMemoryBarrier2-srcAccessMask-03916  
  If `srcAccessMask` includes `VK_ACCESS_2_HOST_READ_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_HOST_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03917)VUID-VkImageMemoryBarrier2-srcAccessMask-03917  
  If `srcAccessMask` includes `VK_ACCESS_2_HOST_WRITE_BIT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_HOST_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03918)VUID-VkImageMemoryBarrier2-srcAccessMask-03918  
  If `srcAccessMask` includes `VK_ACCESS_2_CONDITIONAL_RENDERING_READ_BIT_EXT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03919)VUID-VkImageMemoryBarrier2-srcAccessMask-03919  
  If `srcAccessMask` includes `VK_ACCESS_2_FRAGMENT_DENSITY_MAP_READ_BIT_EXT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03920)VUID-VkImageMemoryBarrier2-srcAccessMask-03920  
  If `srcAccessMask` includes `VK_ACCESS_2_TRANSFORM_FEEDBACK_WRITE_BIT_EXT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-04747)VUID-VkImageMemoryBarrier2-srcAccessMask-04747  
  If `srcAccessMask` includes `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT`, `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03922)VUID-VkImageMemoryBarrier2-srcAccessMask-03922  
  If `srcAccessMask` includes `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03923)VUID-VkImageMemoryBarrier2-srcAccessMask-03923  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADING_RATE_IMAGE_READ_BIT_NV`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_SHADING_RATE_IMAGE_BIT_NV`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-04994)VUID-VkImageMemoryBarrier2-srcAccessMask-04994  
  If `srcAccessMask` includes `VK_ACCESS_2_INVOCATION_MASK_READ_BIT_HUAWEI`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03924)VUID-VkImageMemoryBarrier2-srcAccessMask-03924  
  If `srcAccessMask` includes `VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_NV`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03925)VUID-VkImageMemoryBarrier2-srcAccessMask-03925  
  If `srcAccessMask` includes `VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_NV`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03926)VUID-VkImageMemoryBarrier2-srcAccessMask-03926  
  If `srcAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT` `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03927)VUID-VkImageMemoryBarrier2-srcAccessMask-03927  
  If `srcAccessMask` includes `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-03928)VUID-VkImageMemoryBarrier2-srcAccessMask-03928  
  If `srcAccessMask` includes `VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-06256)VUID-VkImageMemoryBarrier2-srcAccessMask-06256  
  If the [`rayQuery`](#features-rayQuery) feature is not enabled and `srcAccessMask` includes `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `srcStageMask` **must** not include any of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages except `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-07272)VUID-VkImageMemoryBarrier2-srcAccessMask-07272  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_BINDING_TABLE_READ_BIT_KHR`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT` or `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-04858)VUID-VkImageMemoryBarrier2-srcAccessMask-04858  
  If `srcAccessMask` includes `VK_ACCESS_2_VIDEO_DECODE_READ_BIT_KHR`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-04859)VUID-VkImageMemoryBarrier2-srcAccessMask-04859  
  If `srcAccessMask` includes `VK_ACCESS_2_VIDEO_DECODE_WRITE_BIT_KHR`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-04860)VUID-VkImageMemoryBarrier2-srcAccessMask-04860  
  If `srcAccessMask` includes `VK_ACCESS_2_VIDEO_ENCODE_READ_BIT_KHR`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-04861)VUID-VkImageMemoryBarrier2-srcAccessMask-04861  
  If `srcAccessMask` includes `VK_ACCESS_2_VIDEO_ENCODE_WRITE_BIT_KHR`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-07455)VUID-VkImageMemoryBarrier2-srcAccessMask-07455  
  If `srcAccessMask` includes `VK_ACCESS_2_OPTICAL_FLOW_READ_BIT_NV`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-07456)VUID-VkImageMemoryBarrier2-srcAccessMask-07456  
  If `srcAccessMask` includes `VK_ACCESS_2_OPTICAL_FLOW_WRITE_BIT_NV`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-07457)VUID-VkImageMemoryBarrier2-srcAccessMask-07457  
  If `srcAccessMask` includes `VK_ACCESS_2_MICROMAP_WRITE_BIT_EXT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-07458)VUID-VkImageMemoryBarrier2-srcAccessMask-07458  
  If `srcAccessMask` includes `VK_ACCESS_2_MICROMAP_READ_BIT_EXT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` or `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-08118)VUID-VkImageMemoryBarrier2-srcAccessMask-08118  
  If `srcAccessMask` includes `VK_ACCESS_2_DESCRIPTOR_BUFFER_READ_BIT_EXT`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-10670)VUID-VkImageMemoryBarrier2-srcAccessMask-10670  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_TILE_ATTACHMENT_READ_BIT_QCOM`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT` or `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-10671)VUID-VkImageMemoryBarrier2-srcAccessMask-10671  
  If `srcAccessMask` includes `VK_ACCESS_2_SHADER_TILE_ATTACHMENT_WRITE_BIT_QCOM`, `srcStageMask` **must** include `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT` or `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT`

<!--THE END-->

- [](#VUID-VkImageMemoryBarrier2-dstStageMask-03929)VUID-VkImageMemoryBarrier2-dstStageMask-03929  
  If the [`geometryShader`](#features-geometryShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstStageMask-03930)VUID-VkImageMemoryBarrier2-dstStageMask-03930  
  If the [`tessellationShader`](#features-tessellationShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT` or `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstStageMask-03931)VUID-VkImageMemoryBarrier2-dstStageMask-03931  
  If the [`conditionalRendering`](#features-conditionalRendering) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`
- [](#VUID-VkImageMemoryBarrier2-dstStageMask-03932)VUID-VkImageMemoryBarrier2-dstStageMask-03932  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- [](#VUID-VkImageMemoryBarrier2-dstStageMask-03933)VUID-VkImageMemoryBarrier2-dstStageMask-03933  
  If the [`transformFeedback`](#features-transformFeedback) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`
- [](#VUID-VkImageMemoryBarrier2-dstStageMask-03934)VUID-VkImageMemoryBarrier2-dstStageMask-03934  
  If the [`meshShader`](#features-meshShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`
- [](#VUID-VkImageMemoryBarrier2-dstStageMask-03935)VUID-VkImageMemoryBarrier2-dstStageMask-03935  
  If the [`taskShader`](#features-taskShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`
- [](#VUID-VkImageMemoryBarrier2-dstStageMask-07316)VUID-VkImageMemoryBarrier2-dstStageMask-07316  
  If neither of the [`shadingRateImage`](#features-shadingRateImage) or the [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate) features are enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-dstStageMask-04957)VUID-VkImageMemoryBarrier2-dstStageMask-04957  
  If the [`subpassShading`](#features-subpassShading) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`
- [](#VUID-VkImageMemoryBarrier2-dstStageMask-04995)VUID-VkImageMemoryBarrier2-dstStageMask-04995  
  If the [`invocationMask`](#features-invocationMask) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`
- [](#VUID-VkImageMemoryBarrier2-dstStageMask-07946)VUID-VkImageMemoryBarrier2-dstStageMask-07946  
  If neither the [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) extension or the [`rayTracingPipeline`](#features-rayTracingPipeline) feature are enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-dstStageMask-10751)VUID-VkImageMemoryBarrier2-dstStageMask-10751  
  If the [`accelerationStructure`](#features-accelerationStructure) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-dstStageMask-10752)VUID-VkImageMemoryBarrier2-dstStageMask-10752  
  If the [`rayTracingMaintenance1`](#features-rayTracingMaintenance1) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-dstStageMask-10753)VUID-VkImageMemoryBarrier2-dstStageMask-10753  
  If the [`micromap`](#features-micromap) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`

<!--THE END-->

- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03900)VUID-VkImageMemoryBarrier2-dstAccessMask-03900  
  If `dstAccessMask` includes `VK_ACCESS_2_INDIRECT_COMMAND_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_COPY_INDIRECT_BIT_KHR`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03901)VUID-VkImageMemoryBarrier2-dstAccessMask-03901  
  If `dstAccessMask` includes `VK_ACCESS_2_INDEX_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_INDEX_INPUT_BIT`, `VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03902)VUID-VkImageMemoryBarrier2-dstAccessMask-03902  
  If `dstAccessMask` includes `VK_ACCESS_2_VERTEX_ATTRIBUTE_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_VERTEX_ATTRIBUTE_INPUT_BIT`, `VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03903)VUID-VkImageMemoryBarrier2-dstAccessMask-03903  
  If `dstAccessMask` includes `VK_ACCESS_2_INPUT_ATTACHMENT_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`, `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03904)VUID-VkImageMemoryBarrier2-dstAccessMask-03904  
  If `dstAccessMask` includes `VK_ACCESS_2_UNIFORM_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03905)VUID-VkImageMemoryBarrier2-dstAccessMask-03905  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_SAMPLED_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03906)VUID-VkImageMemoryBarrier2-dstAccessMask-03906  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_STORAGE_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03907)VUID-VkImageMemoryBarrier2-dstAccessMask-03907  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-07454)VUID-VkImageMemoryBarrier2-dstAccessMask-07454  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03909)VUID-VkImageMemoryBarrier2-dstAccessMask-03909  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_WRITE_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03910)VUID-VkImageMemoryBarrier2-dstAccessMask-03910  
  If `dstAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT` `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03911)VUID-VkImageMemoryBarrier2-dstAccessMask-03911  
  If `dstAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT` `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03912)VUID-VkImageMemoryBarrier2-dstAccessMask-03912  
  If `dstAccessMask` includes `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT`, `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03913)VUID-VkImageMemoryBarrier2-dstAccessMask-03913  
  If `dstAccessMask` includes `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT`, `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03914)VUID-VkImageMemoryBarrier2-dstAccessMask-03914  
  If `dstAccessMask` includes `VK_ACCESS_2_TRANSFER_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_COPY_BIT`, `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_RESOLVE_BIT`, `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_CONVERT_COOPERATIVE_VECTOR_MATRIX_BIT_NV`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03915)VUID-VkImageMemoryBarrier2-dstAccessMask-03915  
  If `dstAccessMask` includes `VK_ACCESS_2_TRANSFER_WRITE_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_COPY_BIT`, `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_RESOLVE_BIT`, `VK_PIPELINE_STAGE_2_CLEAR_BIT`, `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_CONVERT_COOPERATIVE_VECTOR_MATRIX_BIT_NV`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03916)VUID-VkImageMemoryBarrier2-dstAccessMask-03916  
  If `dstAccessMask` includes `VK_ACCESS_2_HOST_READ_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_HOST_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03917)VUID-VkImageMemoryBarrier2-dstAccessMask-03917  
  If `dstAccessMask` includes `VK_ACCESS_2_HOST_WRITE_BIT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_HOST_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03918)VUID-VkImageMemoryBarrier2-dstAccessMask-03918  
  If `dstAccessMask` includes `VK_ACCESS_2_CONDITIONAL_RENDERING_READ_BIT_EXT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03919)VUID-VkImageMemoryBarrier2-dstAccessMask-03919  
  If `dstAccessMask` includes `VK_ACCESS_2_FRAGMENT_DENSITY_MAP_READ_BIT_EXT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03920)VUID-VkImageMemoryBarrier2-dstAccessMask-03920  
  If `dstAccessMask` includes `VK_ACCESS_2_TRANSFORM_FEEDBACK_WRITE_BIT_EXT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-04747)VUID-VkImageMemoryBarrier2-dstAccessMask-04747  
  If `dstAccessMask` includes `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT`, `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03922)VUID-VkImageMemoryBarrier2-dstAccessMask-03922  
  If `dstAccessMask` includes `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03923)VUID-VkImageMemoryBarrier2-dstAccessMask-03923  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADING_RATE_IMAGE_READ_BIT_NV`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_SHADING_RATE_IMAGE_BIT_NV`, `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-04994)VUID-VkImageMemoryBarrier2-dstAccessMask-04994  
  If `dstAccessMask` includes `VK_ACCESS_2_INVOCATION_MASK_READ_BIT_HUAWEI`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03924)VUID-VkImageMemoryBarrier2-dstAccessMask-03924  
  If `dstAccessMask` includes `VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_NV`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03925)VUID-VkImageMemoryBarrier2-dstAccessMask-03925  
  If `dstAccessMask` includes `VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_NV`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03926)VUID-VkImageMemoryBarrier2-dstAccessMask-03926  
  If `dstAccessMask` includes `VK_ACCESS_2_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT` `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03927)VUID-VkImageMemoryBarrier2-dstAccessMask-03927  
  If `dstAccessMask` includes `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-03928)VUID-VkImageMemoryBarrier2-dstAccessMask-03928  
  If `dstAccessMask` includes `VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` or `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-06256)VUID-VkImageMemoryBarrier2-dstAccessMask-06256  
  If the [`rayQuery`](#features-rayQuery) feature is not enabled and `dstAccessMask` includes `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `dstStageMask` **must** not include any of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages except `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-07272)VUID-VkImageMemoryBarrier2-dstAccessMask-07272  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_BINDING_TABLE_READ_BIT_KHR`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT` or `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-04858)VUID-VkImageMemoryBarrier2-dstAccessMask-04858  
  If `dstAccessMask` includes `VK_ACCESS_2_VIDEO_DECODE_READ_BIT_KHR`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-04859)VUID-VkImageMemoryBarrier2-dstAccessMask-04859  
  If `dstAccessMask` includes `VK_ACCESS_2_VIDEO_DECODE_WRITE_BIT_KHR`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-04860)VUID-VkImageMemoryBarrier2-dstAccessMask-04860  
  If `dstAccessMask` includes `VK_ACCESS_2_VIDEO_ENCODE_READ_BIT_KHR`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-04861)VUID-VkImageMemoryBarrier2-dstAccessMask-04861  
  If `dstAccessMask` includes `VK_ACCESS_2_VIDEO_ENCODE_WRITE_BIT_KHR`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-07455)VUID-VkImageMemoryBarrier2-dstAccessMask-07455  
  If `dstAccessMask` includes `VK_ACCESS_2_OPTICAL_FLOW_READ_BIT_NV`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-07456)VUID-VkImageMemoryBarrier2-dstAccessMask-07456  
  If `dstAccessMask` includes `VK_ACCESS_2_OPTICAL_FLOW_WRITE_BIT_NV`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-07457)VUID-VkImageMemoryBarrier2-dstAccessMask-07457  
  If `dstAccessMask` includes `VK_ACCESS_2_MICROMAP_WRITE_BIT_EXT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-07458)VUID-VkImageMemoryBarrier2-dstAccessMask-07458  
  If `dstAccessMask` includes `VK_ACCESS_2_MICROMAP_READ_BIT_EXT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` or `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-08118)VUID-VkImageMemoryBarrier2-dstAccessMask-08118  
  If `dstAccessMask` includes `VK_ACCESS_2_DESCRIPTOR_BUFFER_READ_BIT_EXT`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT`, `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`, or one of `VK_PIPELINE_STAGE_*_SHADER_BIT` stages
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-10670)VUID-VkImageMemoryBarrier2-dstAccessMask-10670  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_TILE_ATTACHMENT_READ_BIT_QCOM`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT` or `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT`
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-10671)VUID-VkImageMemoryBarrier2-dstAccessMask-10671  
  If `dstAccessMask` includes `VK_ACCESS_2_SHADER_TILE_ATTACHMENT_WRITE_BIT_QCOM`, `dstStageMask` **must** include `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT` or `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT`

<!--THE END-->

- [](#VUID-VkImageMemoryBarrier2-oldLayout-01208)VUID-VkImageMemoryBarrier2-oldLayout-01208  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier2-oldLayout-01209)VUID-VkImageMemoryBarrier2-oldLayout-01209  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier2-oldLayout-01210)VUID-VkImageMemoryBarrier2-oldLayout-01210  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier2-oldLayout-01211)VUID-VkImageMemoryBarrier2-oldLayout-01211  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_SAMPLED_BIT` or `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier2-oldLayout-01212)VUID-VkImageMemoryBarrier2-oldLayout-01212  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_TRANSFER_SRC_BIT`
- [](#VUID-VkImageMemoryBarrier2-oldLayout-01213)VUID-VkImageMemoryBarrier2-oldLayout-01213  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_TRANSFER_DST_BIT`
- [](#VUID-VkImageMemoryBarrier2-oldLayout-01197)VUID-VkImageMemoryBarrier2-oldLayout-01197  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), `oldLayout` **must** be `VK_IMAGE_LAYOUT_UNDEFINED` or the current layout of the image subresources affected by the barrier
- [](#VUID-VkImageMemoryBarrier2-oldLayout-10767)VUID-VkImageMemoryBarrier2-oldLayout-10767  
  If the [zeroInitializeDeviceMemory](#features-zeroInitializeDeviceMemory) feature is not enabled, `oldLayout` **must** not be `VK_IMAGE_LAYOUT_ZERO_INITIALIZED_EXT`
- [](#VUID-VkImageMemoryBarrier2-oldLayout-10768)VUID-VkImageMemoryBarrier2-oldLayout-10768  
  If `oldLayout` is `VK_IMAGE_LAYOUT_ZERO_INITIALIZED_EXT`, then all subresources **must** be included in the barrier
- [](#VUID-VkImageMemoryBarrier2-newLayout-01198)VUID-VkImageMemoryBarrier2-newLayout-01198  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), `newLayout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED` or `VK_IMAGE_LAYOUT_ZERO_INITIALIZED_EXT` or `VK_IMAGE_LAYOUT_PREINITIALIZED`
- [](#VUID-VkImageMemoryBarrier2-oldLayout-01658)VUID-VkImageMemoryBarrier2-oldLayout-01658  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier2-oldLayout-01659)VUID-VkImageMemoryBarrier2-oldLayout-01659  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04065)VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04065  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL` then `image` **must** have been created with at least one of `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`, `VK_IMAGE_USAGE_SAMPLED_BIT`, or `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04066)VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04066  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT` set
- [](#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04067)VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04067  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL` then `image` **must** have been created with at least one of `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`, `VK_IMAGE_USAGE_SAMPLED_BIT`, or `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04068)VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-04068  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT` set
- [](#VUID-VkImageMemoryBarrier2-synchronization2-07793)VUID-VkImageMemoryBarrier2-synchronization2-07793  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `oldLayout` **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`
- [](#VUID-VkImageMemoryBarrier2-synchronization2-07794)VUID-VkImageMemoryBarrier2-synchronization2-07794  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `newLayout` **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-03938)VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-03938  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL`, `image` **must** have been created with `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-03939)VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-03939  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL`, `image` **must** have been created with at least one of `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`, `VK_IMAGE_USAGE_SAMPLED_BIT`, or `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier2-oldLayout-02088)VUID-VkImageMemoryBarrier2-oldLayout-02088  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_FRAGMENT_SHADING_RATE_ATTACHMENT_OPTIMAL_KHR` then `image` **must** have been created with `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR` set
- [](#VUID-VkImageMemoryBarrier2-image-09117)VUID-VkImageMemoryBarrier2-image-09117  
  If `image` was created with a sharing mode of `VK_SHARING_MODE_EXCLUSIVE`, and `srcQueueFamilyIndex` and `dstQueueFamilyIndex` are not equal, `srcQueueFamilyIndex` **must** be `VK_QUEUE_FAMILY_EXTERNAL`, `VK_QUEUE_FAMILY_FOREIGN_EXT`, or a valid queue family
- [](#VUID-VkImageMemoryBarrier2-image-09118)VUID-VkImageMemoryBarrier2-image-09118  
  If `image` was created with a sharing mode of `VK_SHARING_MODE_EXCLUSIVE`, and `srcQueueFamilyIndex` and `dstQueueFamilyIndex` are not equal, `dstQueueFamilyIndex` **must** be `VK_QUEUE_FAMILY_EXTERNAL`, `VK_QUEUE_FAMILY_FOREIGN_EXT`, or a valid queue family
- [](#VUID-VkImageMemoryBarrier2-None-09119)VUID-VkImageMemoryBarrier2-None-09119  
  If the [VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html) extension is not enabled, and the value of [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`apiVersion` used to create the [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) is not greater than or equal to Version 1.1, `srcQueueFamilyIndex` **must** not be `VK_QUEUE_FAMILY_EXTERNAL`
- [](#VUID-VkImageMemoryBarrier2-None-09120)VUID-VkImageMemoryBarrier2-None-09120  
  If the [VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html) extension is not enabled, and the value of [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`apiVersion` used to create the [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) is not greater than or equal to Version 1.1, `dstQueueFamilyIndex` **must** not be `VK_QUEUE_FAMILY_EXTERNAL`
- [](#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-09121)VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-09121  
  If the [VK\_EXT\_queue\_family\_foreign](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_queue_family_foreign.html) extension is not enabled `srcQueueFamilyIndex` **must** not be `VK_QUEUE_FAMILY_FOREIGN_EXT`
- [](#VUID-VkImageMemoryBarrier2-dstQueueFamilyIndex-09122)VUID-VkImageMemoryBarrier2-dstQueueFamilyIndex-09122  
  If the [VK\_EXT\_queue\_family\_foreign](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_queue_family_foreign.html) extension is not enabled `dstQueueFamilyIndex` **must** not be `VK_QUEUE_FAMILY_FOREIGN_EXT`
- [](#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07120)VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07120  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_DECODE_SRC_KHR` then `image` **must** have been created with `VK_IMAGE_USAGE_VIDEO_DECODE_SRC_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07121)VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07121  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_DECODE_DST_KHR` then `image` **must** have been created with `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07122)VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07122  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_DECODE_DPB_KHR` then `image` **must** have been created with `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07123)VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07123  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_ENCODE_SRC_KHR` then `image` **must** have been created with `VK_IMAGE_USAGE_VIDEO_ENCODE_SRC_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07124)VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07124  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_ENCODE_DST_KHR` then `image` **must** have been created with `VK_IMAGE_USAGE_VIDEO_ENCODE_DST_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07125)VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07125  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_ENCODE_DPB_KHR` then `image` **must** have been created with `VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-10287)VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-10287  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_ENCODE_QUANTIZATION_MAP_KHR` then `image` **must** have been created with `VK_IMAGE_USAGE_VIDEO_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR` or `VK_IMAGE_USAGE_VIDEO_ENCODE_EMPHASIS_MAP_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07006)VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-07006  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT` then `image` **must** have been created with either the `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT` usage bits, and the `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_SAMPLED_BIT` usage bits, and the `VK_IMAGE_USAGE_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT` usage bit
- [](#VUID-VkImageMemoryBarrier2-attachmentFeedbackLoopLayout-07313)VUID-VkImageMemoryBarrier2-attachmentFeedbackLoopLayout-07313  
  If the [`attachmentFeedbackLoopLayout`](#features-attachmentFeedbackLoopLayout) feature is not enabled, `newLayout` **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`
- [](#VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-09550)VUID-VkImageMemoryBarrier2-srcQueueFamilyIndex-09550  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ` then `image` **must** have been created with either `VK_IMAGE_USAGE_STORAGE_BIT`, or with both `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` and either of `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier2-dynamicRenderingLocalRead-09551)VUID-VkImageMemoryBarrier2-dynamicRenderingLocalRead-09551  
  If the [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead) feature is not enabled, `oldLayout` **must** not be `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ`
- [](#VUID-VkImageMemoryBarrier2-dynamicRenderingLocalRead-09552)VUID-VkImageMemoryBarrier2-dynamicRenderingLocalRead-09552  
  If the [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead) feature is not enabled, `newLayout` **must** not be `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ`

<!--THE END-->

- [](#VUID-VkImageMemoryBarrier2-subresourceRange-01486)VUID-VkImageMemoryBarrier2-subresourceRange-01486  
  `subresourceRange.baseMipLevel` **must** be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkImageMemoryBarrier2-subresourceRange-01724)VUID-VkImageMemoryBarrier2-subresourceRange-01724  
  If `subresourceRange.levelCount` is not `VK_REMAINING_MIP_LEVELS`, `subresourceRange.baseMipLevel` + `subresourceRange.levelCount` **must** be less than or equal to the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkImageMemoryBarrier2-subresourceRange-01488)VUID-VkImageMemoryBarrier2-subresourceRange-01488  
  If `image` is not a 3D image or was created without `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set, or the [`maintenance9`](#features-maintenance9) feature is not enabled, `subresourceRange.baseArrayLayer` **must** be less than the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkImageMemoryBarrier2-maintenance9-10798)VUID-VkImageMemoryBarrier2-maintenance9-10798  
  If the [`maintenance9`](#features-maintenance9) feature is enabled and `image` is a 3D image created with `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set, `subresourceRange.baseArrayLayer` **must** be less than the depth computed from `baseMipLevel` and `extent.depth` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created, according to the formula defined in [Image Mip Level Sizing](#resources-image-mip-level-sizing)
- [](#VUID-VkImageMemoryBarrier2-maintenance9-10799)VUID-VkImageMemoryBarrier2-maintenance9-10799  
  If the [`maintenance9`](#features-maintenance9) feature is enabled and `image` is a 3D image created with `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set and either `subresourceRange.baseArrayLayer` is not equal to 0 or `subresourceRange.layerCount` is not equal to `VK_REMAINING_ARRAY_LAYERS`, `subresourceRange.levelCount` **must** be 1
- [](#VUID-VkImageMemoryBarrier2-subresourceRange-01725)VUID-VkImageMemoryBarrier2-subresourceRange-01725  
  If `image` is not a 3D image or was created without `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set, or the [`maintenance9`](#features-maintenance9) feature is not enabled, and `subresourceRange.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, `subresourceRange.baseArrayLayer` + `subresourceRange.layerCount` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkImageMemoryBarrier2-maintenance9-10800)VUID-VkImageMemoryBarrier2-maintenance9-10800  
  If the [`maintenance9`](#features-maintenance9) feature is enabled, `subresourceRange.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, and `image` is a 3D image created with `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set, `subresourceRange.baseArrayLayer` + `subresourceRange.layerCount` **must** be less than or equal to the depth computed from `baseMipLevel` and `extent.depth` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created, according to the formula defined in [Image Mip Level Sizing](#resources-image-mip-level-sizing)
- [](#VUID-VkImageMemoryBarrier2-image-01932)VUID-VkImageMemoryBarrier2-image-01932  
  If `image` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkImageMemoryBarrier2-image-09241)VUID-VkImageMemoryBarrier2-image-09241  
  If `image` has a color format that is single-plane, then the `aspectMask` member of `subresourceRange` **must** be `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-VkImageMemoryBarrier2-image-09242)VUID-VkImageMemoryBarrier2-image-09242  
  If `image` has a color format and is not *disjoint*, then the `aspectMask` member of `subresourceRange` **must** be `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-VkImageMemoryBarrier2-image-01672)VUID-VkImageMemoryBarrier2-image-01672  
  If `image` has a [multi-planar format](#formats-multiplanar) and the image is *disjoint*, then the `aspectMask` member of `subresourceRange` **must** include at least one [multi-planar aspect mask](#formats-multiplanar-image-aspect) bit or `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-VkImageMemoryBarrier2-image-03320)VUID-VkImageMemoryBarrier2-image-03320  
  If `image` has a depth/stencil format with both depth and stencil and the [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts) feature is not enabled, then the `aspectMask` member of `subresourceRange` **must** include both `VK_IMAGE_ASPECT_DEPTH_BIT` and `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-VkImageMemoryBarrier2-image-03319)VUID-VkImageMemoryBarrier2-image-03319  
  If `image` has a depth/stencil format with both depth and stencil and the [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts) feature is enabled, then the `aspectMask` member of `subresourceRange` **must** include either or both `VK_IMAGE_ASPECT_DEPTH_BIT` and `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-VkImageMemoryBarrier2-image-10749)VUID-VkImageMemoryBarrier2-image-10749  
  If `image` has a depth-only format then the `aspectMask` member of `subresourceRange` **must** be `VK_IMAGE_ASPECT_DEPTH_BIT`
- [](#VUID-VkImageMemoryBarrier2-image-10750)VUID-VkImageMemoryBarrier2-image-10750  
  If `image` has a stencil-only format then the `aspectMask` member of `subresourceRange` **must** be `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-VkImageMemoryBarrier2-aspectMask-08702)VUID-VkImageMemoryBarrier2-aspectMask-08702  
  If the `aspectMask` member of `subresourceRange` includes `VK_IMAGE_ASPECT_DEPTH_BIT`, `oldLayout` and `newLayout` **must** not be one of `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkImageMemoryBarrier2-aspectMask-08703)VUID-VkImageMemoryBarrier2-aspectMask-08703  
  If the `aspectMask` member of `subresourceRange` includes `VK_IMAGE_ASPECT_STENCIL_BIT`, `oldLayout` and `newLayout` **must** not be one of `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`
- [](#VUID-VkImageMemoryBarrier2-subresourceRange-09601)VUID-VkImageMemoryBarrier2-subresourceRange-09601  
  `subresourceRange.aspectMask` **must** be valid for the `format` the `image` was created with
- [](#VUID-VkImageMemoryBarrier2-srcStageMask-03854)VUID-VkImageMemoryBarrier2-srcStageMask-03854  
  If either `srcStageMask` or `dstStageMask` includes `VK_PIPELINE_STAGE_2_HOST_BIT`, `srcQueueFamilyIndex` and `dstQueueFamilyIndex` **must** be equal
- [](#VUID-VkImageMemoryBarrier2-srcStageMask-03855)VUID-VkImageMemoryBarrier2-srcStageMask-03855  
  If `srcStageMask` includes `VK_PIPELINE_STAGE_2_HOST_BIT`, and `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-image-layout-transitions), `oldLayout` **must** be one of `VK_IMAGE_LAYOUT_PREINITIALIZED`, `VK_IMAGE_LAYOUT_ZERO_INITIALIZED_EXT`, `VK_IMAGE_LAYOUT_UNDEFINED`, or `VK_IMAGE_LAYOUT_GENERAL`

Valid Usage (Implicit)

- [](#VUID-VkImageMemoryBarrier2-sType-sType)VUID-VkImageMemoryBarrier2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER_2`
- [](#VUID-VkImageMemoryBarrier2-pNext-pNext)VUID-VkImageMemoryBarrier2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkExternalMemoryAcquireUnmodifiedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryAcquireUnmodifiedEXT.html), [VkMemoryBarrierAccessFlags3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrierAccessFlags3KHR.html), or [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleLocationsInfoEXT.html)
- [](#VUID-VkImageMemoryBarrier2-sType-unique)VUID-VkImageMemoryBarrier2-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkImageMemoryBarrier2-srcStageMask-parameter)VUID-VkImageMemoryBarrier2-srcStageMask-parameter  
  `srcStageMask` **must** be a valid combination of [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html) values
- [](#VUID-VkImageMemoryBarrier2-srcAccessMask-parameter)VUID-VkImageMemoryBarrier2-srcAccessMask-parameter  
  `srcAccessMask` **must** be a valid combination of [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html) values
- [](#VUID-VkImageMemoryBarrier2-dstStageMask-parameter)VUID-VkImageMemoryBarrier2-dstStageMask-parameter  
  `dstStageMask` **must** be a valid combination of [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html) values
- [](#VUID-VkImageMemoryBarrier2-dstAccessMask-parameter)VUID-VkImageMemoryBarrier2-dstAccessMask-parameter  
  `dstAccessMask` **must** be a valid combination of [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html) values
- [](#VUID-VkImageMemoryBarrier2-oldLayout-parameter)VUID-VkImageMemoryBarrier2-oldLayout-parameter  
  `oldLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-VkImageMemoryBarrier2-newLayout-parameter)VUID-VkImageMemoryBarrier2-newLayout-parameter  
  `newLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-VkImageMemoryBarrier2-image-parameter)VUID-VkImageMemoryBarrier2-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-VkImageMemoryBarrier2-subresourceRange-parameter)VUID-VkImageMemoryBarrier2-subresourceRange-parameter  
  `subresourceRange` **must** be a valid [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html) structure

## [](#_see_also)See Also

[VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkAccessFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags2.html), [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfo.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html), [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageMemoryBarrier2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0