# VK_KHR_maintenance5(3) Manual Page

## Name

VK_KHR_maintenance5 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

471

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

     [Version 1.1](#versions-1.1)  
     and  
     [VK_KHR_dynamic_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering.html)  
or  
[Version 1.3](#versions-1.3)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_VERSION_1_1

- Interacts with VK_VERSION_1_2

- Interacts with VK_VERSION_1_3

- Interacts with VK_EXT_attachment_feedback_loop_layout

- Interacts with VK_EXT_buffer_device_address

- Interacts with VK_EXT_conditional_rendering

- Interacts with VK_EXT_descriptor_buffer

- Interacts with VK_EXT_fragment_density_map

- Interacts with VK_EXT_graphics_pipeline_library

- Interacts with VK_EXT_opacity_micromap

- Interacts with VK_EXT_pipeline_creation_cache_control

- Interacts with VK_EXT_pipeline_protected_access

- Interacts with VK_EXT_transform_feedback

- Interacts with VK_KHR_acceleration_structure

- Interacts with VK_KHR_buffer_device_address

- Interacts with VK_KHR_device_group

- Interacts with VK_KHR_dynamic_rendering

- Interacts with VK_KHR_fragment_shading_rate

- Interacts with VK_KHR_pipeline_executable_properties

- Interacts with VK_KHR_pipeline_library

- Interacts with VK_KHR_ray_tracing_pipeline

- Interacts with VK_KHR_video_decode_queue

- Interacts with VK_KHR_video_encode_queue

- Interacts with VK_NV_device_generated_commands

- Interacts with VK_NV_displacement_micromap

- Interacts with VK_NV_ray_tracing

- Interacts with VK_NV_ray_tracing_motion_blur

## <a href="#_contact" class="anchor"></a>Contact

- Stu Smith <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_maintenance5%5D%20@stu-s%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_maintenance5%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>stu-s</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_maintenance5](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_maintenance5.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-05-02

**Interactions and External Dependencies**  
**Contributors**  
- Stu Smith, AMD

- Tobias Hector, AMD

- Shahbaz Youssefi, Google

- Slawomir Cygan, Intel

- Lionel Landwerlin, Intel

- James Fitzpatrick, Imagination Technologies

- Andrew Garrard, Imagination Technologies

- Ralph Potter, Samsung

- Pan Gao, Huawei

- Jan-Harald Fredriksen, ARM

- Jon Leech, Khronos

- Mike Blumenkrantz, Valve

## <a href="#_description" class="anchor"></a>Description

`VK_KHR_maintenance5` adds a collection of minor features, none of which
would warrant an entire extension of their own.

The new features are as follows:

- A new `VK_FORMAT_A1B5G5R5_UNORM_PACK16_KHR` format

- A new `VK_FORMAT_A8_UNORM_KHR` format

- A property to indicate that multisample coverage operations are
  performed after sample counting in EarlyFragmentTests mode

- Relax VkBufferView creation requirements by allowing subsets of the
  associated VkBuffer usage using `VkBufferUsageFlags2CreateInfoKHR`

- A new command
  [vkCmdBindIndexBuffer2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer2KHR.html), allowing a
  range of memory to be bound as an index buffer

- [vkGetDeviceProcAddr](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceProcAddr.html) must return `NULL` for
  supported core functions beyond the version requested by the
  application.

- A property to indicate that the sample mask test is performed after
  sample counting in EarlyFragmentTests mode

- `vkCmdBindVertexBuffers2` now supports using `VK_WHOLE_SIZE` in the
  `pSizes` parameter.

- A default size of 1.0 is used if `PointSize` is not written

- Shader modules are deprecated - applications can now pass
  [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html) as a chained
  struct to pipeline creation via
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)

- A function
  [vkGetRenderingAreaGranularityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRenderingAreaGranularityKHR.html)
  to query the optimal render area for a dynamic rendering instance.

- A property to indicate that depth/stencil texturing operations with
  `VK_COMPONENT_SWIZZLE_ONE` have defined behavior

- Add
  [vkGetImageSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout2KHR.html)
  and a new function
  [vkGetDeviceImageSubresourceLayoutKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageSubresourceLayoutKHR.html)
  to allow the application to query the image memory layout without
  having to create an image object and query it.

- Allow `VK_REMAINING_ARRAY_LAYERS` as the `layerCount` member of
  [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html)

- Adds stronger guarantees for propagation of `VK_ERROR_DEVICE_LOST`
  return values

- A property to indicate whether `PointSize` controls the final
  rasterization of polygons if <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-polygonmode"
  target="_blank" rel="noopener">polygon mode</a> is
  `VK_POLYGON_MODE_POINT`

- Two properties to indicate the non-strict line rasterization algorithm
  used

- Two new flags words
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html) and
  [VkBufferUsageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits2KHR.html)

- Physical-device-level functions can now be called with any value in
  the valid range for a type beyond the defined enumerants, such that
  applications can avoid checking individual features, extensions, or
  versions before querying supported properties of a particular
  enumerant.

- Clarification that copies between images of any type are allowed,
  treating 1D images as 2D images with a height of 1.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdBindIndexBuffer2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer2KHR.html)

- [vkGetDeviceImageSubresourceLayoutKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageSubresourceLayoutKHR.html)

- [vkGetImageSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout2KHR.html)

- [vkGetRenderingAreaGranularityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRenderingAreaGranularityKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkDeviceImageSubresourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceImageSubresourceInfoKHR.html)

- [VkImageSubresource2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource2KHR.html)

- [VkRenderingAreaInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAreaInfoKHR.html)

- [VkSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout2KHR.html)

- Extending [VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferViewCreateInfo.html),
  [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html),
  [VkPhysicalDeviceExternalBufferInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalBufferInfo.html),
  [VkDescriptorBufferBindingInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBufferBindingInfoEXT.html):

  - [VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2CreateInfoKHR.html)

- Extending
  [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html),
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
  [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html),
  [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html):

  - [VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2CreateInfoKHR.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceMaintenance5FeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance5FeaturesKHR.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceMaintenance5PropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance5PropertiesKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkBufferUsageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits2KHR.html)

- [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkBufferUsageFlags2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2KHR.html)

- [VkPipelineCreateFlags2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2KHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_MAINTENANCE_5_EXTENSION_NAME`

- `VK_KHR_MAINTENANCE_5_SPEC_VERSION`

- Extending [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html):

  - `VK_FORMAT_A1B5G5R5_UNORM_PACK16_KHR`

  - `VK_FORMAT_A8_UNORM_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_BUFFER_USAGE_FLAGS_2_CREATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_DEVICE_IMAGE_SUBRESOURCE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_IMAGE_SUBRESOURCE_2_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_5_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_5_PROPERTIES_KHR`

  - `VK_STRUCTURE_TYPE_PIPELINE_CREATE_FLAGS_2_CREATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_RENDERING_AREA_INFO_KHR`

  - `VK_STRUCTURE_TYPE_SUBRESOURCE_LAYOUT_2_KHR`

If [VK_KHR_dynamic_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering.html) or [Version
1.3](#versions-1.3) and
[VK_EXT_fragment_density_map](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_fragment_density_map.html) is
supported:

- Extending
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html):

  - `VK_PIPELINE_CREATE_2_RENDERING_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT`

If [VK_KHR_dynamic_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering.html) or [Version
1.3](#versions-1.3) and
[VK_KHR_fragment_shading_rate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_fragment_shading_rate.html) is
supported:

- Extending
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html):

  - `VK_PIPELINE_CREATE_2_RENDERING_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

If
[VK_EXT_attachment_feedback_loop_layout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_attachment_feedback_loop_layout.html)
is supported:

- Extending
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html):

  - `VK_PIPELINE_CREATE_2_COLOR_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`

  - `VK_PIPELINE_CREATE_2_DEPTH_STENCIL_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`

If [VK_EXT_conditional_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_conditional_rendering.html) is
supported:

- Extending [VkBufferUsageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits2KHR.html):

  - `VK_BUFFER_USAGE_2_CONDITIONAL_RENDERING_BIT_EXT`

If [VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html) is
supported:

- Extending [VkBufferUsageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits2KHR.html):

  - `VK_BUFFER_USAGE_2_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT`

  - `VK_BUFFER_USAGE_2_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT`

  - `VK_BUFFER_USAGE_2_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT`

- Extending
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html):

  - `VK_PIPELINE_CREATE_2_DESCRIPTOR_BUFFER_BIT_EXT`

If
[VK_EXT_graphics_pipeline_library](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_graphics_pipeline_library.html)
is supported:

- Extending
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html):

  - `VK_PIPELINE_CREATE_2_LINK_TIME_OPTIMIZATION_BIT_EXT`

  - `VK_PIPELINE_CREATE_2_RETAIN_LINK_TIME_OPTIMIZATION_INFO_BIT_EXT`

If [VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html) is supported:

- Extending [VkBufferUsageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits2KHR.html):

  - `VK_BUFFER_USAGE_2_MICROMAP_BUILD_INPUT_READ_ONLY_BIT_EXT`

  - `VK_BUFFER_USAGE_2_MICROMAP_STORAGE_BIT_EXT`

- Extending
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html):

  - `VK_PIPELINE_CREATE_2_RAY_TRACING_OPACITY_MICROMAP_BIT_EXT`

If
[VK_EXT_pipeline_protected_access](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_pipeline_protected_access.html)
is supported:

- Extending
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html):

  - `VK_PIPELINE_CREATE_2_NO_PROTECTED_ACCESS_BIT_EXT`

  - `VK_PIPELINE_CREATE_2_PROTECTED_ACCESS_ONLY_BIT_EXT`

If [VK_EXT_transform_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_transform_feedback.html) is
supported:

- Extending [VkBufferUsageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits2KHR.html):

  - `VK_BUFFER_USAGE_2_TRANSFORM_FEEDBACK_BUFFER_BIT_EXT`

  - `VK_BUFFER_USAGE_2_TRANSFORM_FEEDBACK_COUNTER_BUFFER_BIT_EXT`

If [VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html)
is supported:

- Extending [VkBufferUsageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits2KHR.html):

  - `VK_BUFFER_USAGE_2_ACCELERATION_STRUCTURE_BUILD_INPUT_READ_ONLY_BIT_KHR`

  - `VK_BUFFER_USAGE_2_ACCELERATION_STRUCTURE_STORAGE_BIT_KHR`

If
[VK_KHR_pipeline_executable_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_pipeline_executable_properties.html)
is supported:

- Extending
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html):

  - `VK_PIPELINE_CREATE_2_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR`

  - `VK_PIPELINE_CREATE_2_CAPTURE_STATISTICS_BIT_KHR`

If [VK_KHR_pipeline_library](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_pipeline_library.html) is supported:

- Extending
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html):

  - `VK_PIPELINE_CREATE_2_LIBRARY_BIT_KHR`

If [VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html) is
supported:

- Extending [VkBufferUsageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits2KHR.html):

  - `VK_BUFFER_USAGE_2_SHADER_BINDING_TABLE_BIT_KHR`

- Extending
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html):

  - `VK_PIPELINE_CREATE_2_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`

  - `VK_PIPELINE_CREATE_2_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`

  - `VK_PIPELINE_CREATE_2_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`

  - `VK_PIPELINE_CREATE_2_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`

  - `VK_PIPELINE_CREATE_2_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`

  - `VK_PIPELINE_CREATE_2_RAY_TRACING_SKIP_AABBS_BIT_KHR`

  - `VK_PIPELINE_CREATE_2_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR`

If [VK_KHR_video_decode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_queue.html) is
supported:

- Extending [VkBufferUsageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits2KHR.html):

  - `VK_BUFFER_USAGE_2_VIDEO_DECODE_DST_BIT_KHR`

  - `VK_BUFFER_USAGE_2_VIDEO_DECODE_SRC_BIT_KHR`

If [VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html) is
supported:

- Extending [VkBufferUsageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits2KHR.html):

  - `VK_BUFFER_USAGE_2_VIDEO_ENCODE_DST_BIT_KHR`

  - `VK_BUFFER_USAGE_2_VIDEO_ENCODE_SRC_BIT_KHR`

If
[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html)
is supported:

- Extending
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html):

  - `VK_PIPELINE_CREATE_2_INDIRECT_BINDABLE_BIT_NV`

If [VK_NV_displacement_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_displacement_micromap.html) is
supported:

- Extending
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html):

  - `VK_PIPELINE_CREATE_2_RAY_TRACING_DISPLACEMENT_MICROMAP_BIT_NV`

If [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) is supported:

- Extending [VkBufferUsageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits2KHR.html):

  - `VK_BUFFER_USAGE_2_RAY_TRACING_BIT_NV`

- Extending
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html):

  - `VK_PIPELINE_CREATE_2_DEFER_COMPILE_BIT_NV`

If [VK_NV_ray_tracing_motion_blur](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing_motion_blur.html)
is supported:

- Extending
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html):

  - `VK_PIPELINE_CREATE_2_RAY_TRACING_ALLOW_MOTION_BIT_NV`

If [Version 1.1](#versions-1.1) or
[VK_KHR_device_group](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group.html) is supported:

- Extending
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html):

  - `VK_PIPELINE_CREATE_2_DISPATCH_BASE_BIT_KHR`

  - `VK_PIPELINE_CREATE_2_VIEW_INDEX_FROM_DEVICE_INDEX_BIT_KHR`

If [Version 1.2](#versions-1.2) or
[VK_KHR_buffer_device_address](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_buffer_device_address.html) or
[VK_EXT_buffer_device_address](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_buffer_device_address.html) is
supported:

- Extending [VkBufferUsageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits2KHR.html):

  - `VK_BUFFER_USAGE_2_SHADER_DEVICE_ADDRESS_BIT_KHR`

If [Version 1.3](#versions-1.3) or
[VK_EXT_pipeline_creation_cache_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_pipeline_creation_cache_control.html)
is supported:

- Extending
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html):

  - `VK_PIPELINE_CREATE_2_EARLY_RETURN_ON_FAILURE_BIT_KHR`

  - `VK_PIPELINE_CREATE_2_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT_KHR`

## <a href="#_issues" class="anchor"></a>Issues

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-12-12 (Stu Smith)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_maintenance5"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
