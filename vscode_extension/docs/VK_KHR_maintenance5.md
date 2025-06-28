# VK\_KHR\_maintenance5(3) Manual Page

## Name

VK\_KHR\_maintenance5 - device extension



## [](#_registered_extension_number)Registered Extension Number

471

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [Vulkan Version 1.1](#versions-1.1)  
     and  
     [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html)  
or  
[Vulkan Version 1.3](#versions-1.3)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_2
- Interacts with VK\_VERSION\_1\_3
- Interacts with VK\_VERSION\_1\_4
- Interacts with VK\_ARM\_pipeline\_opacity\_micromap
- Interacts with VK\_EXT\_attachment\_feedback\_loop\_layout
- Interacts with VK\_EXT\_buffer\_device\_address
- Interacts with VK\_EXT\_conditional\_rendering
- Interacts with VK\_EXT\_descriptor\_buffer
- Interacts with VK\_EXT\_fragment\_density\_map
- Interacts with VK\_EXT\_graphics\_pipeline\_library
- Interacts with VK\_EXT\_opacity\_micromap
- Interacts with VK\_EXT\_pipeline\_creation\_cache\_control
- Interacts with VK\_EXT\_pipeline\_protected\_access
- Interacts with VK\_EXT\_transform\_feedback
- Interacts with VK\_KHR\_acceleration\_structure
- Interacts with VK\_KHR\_buffer\_device\_address
- Interacts with VK\_KHR\_dynamic\_rendering
- Interacts with VK\_KHR\_fragment\_shading\_rate
- Interacts with VK\_KHR\_pipeline\_executable\_properties
- Interacts with VK\_KHR\_pipeline\_library
- Interacts with VK\_KHR\_ray\_tracing\_pipeline
- Interacts with VK\_KHR\_video\_decode\_queue
- Interacts with VK\_KHR\_video\_encode\_queue
- Interacts with VK\_NV\_device\_generated\_commands
- Interacts with VK\_NV\_displacement\_micromap
- Interacts with VK\_NV\_ray\_tracing
- Interacts with VK\_NV\_ray\_tracing\_motion\_blur

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4-promotions)

## [](#_contact)Contact

- Stu Smith [\[GitHub\]stu-s](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_maintenance5%5D%20%40stu-s%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_maintenance5%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_maintenance5](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_maintenance5.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

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

## [](#_description)Description

`VK_KHR_maintenance5` adds a collection of minor features, none of which would warrant an entire extension of their own.

The new features are as follows:

- A new `VK_FORMAT_A1B5G5R5_UNORM_PACK16_KHR` format
- A new `VK_FORMAT_A8_UNORM_KHR` format
- A property to indicate that multisample coverage operations are performed after sample counting in EarlyFragmentTests mode
- Relax VkBufferView creation requirements by allowing subsets of the associated VkBuffer usage using `VkBufferUsageFlags2CreateInfoKHR`
- A new command [vkCmdBindIndexBuffer2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindIndexBuffer2KHR.html), allowing a range of memory to be bound as an index buffer
- [vkGetDeviceProcAddr](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceProcAddr.html) must return `NULL` for supported core functions beyond the version requested by the application.
- A property to indicate that the sample mask test is performed after sample counting in EarlyFragmentTests mode
- `vkCmdBindVertexBuffers2` now supports using `VK_WHOLE_SIZE` in the `pSizes` parameter.
- A default size of 1.0 is used if `PointSize` is not written
- Shader modules are deprecated - applications can now pass [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) as a chained structure to pipeline creation via [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html)
- A function [vkGetRenderingAreaGranularityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRenderingAreaGranularityKHR.html) to query the optimal render area for a dynamic rendering instance.
- A property to indicate that depth/stencil texturing operations with `VK_COMPONENT_SWIZZLE_ONE` have defined behavior
- Add [vkGetImageSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2KHR.html) and a new function [vkGetDeviceImageSubresourceLayoutKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageSubresourceLayoutKHR.html) to allow the application to query the image memory layout without having to create an image object and query it.
- Allow `VK_REMAINING_ARRAY_LAYERS` as the `layerCount` member of [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html)
- Adds stronger guarantees for propagation of `VK_ERROR_DEVICE_LOST` return values
- A property to indicate whether `PointSize` controls the final rasterization of polygons if [polygon mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-polygonmode) is `VK_POLYGON_MODE_POINT`
- Two properties to indicate the non-strict line rasterization algorithm used
- Two new flags words [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2KHR.html) and [VkBufferUsageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2KHR.html)
- Physical-device-level functions can now be called with any value in the valid range for a type beyond the defined enumerants, such that applications can avoid checking individual features, extensions, or versions before querying supported properties of a particular enumerant.
- Clarification that copies between images of any type are allowed, treating 1D images as 2D images with a height of 1.

## [](#_new_commands)New Commands

- [vkCmdBindIndexBuffer2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindIndexBuffer2KHR.html)
- [vkGetDeviceImageSubresourceLayoutKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageSubresourceLayoutKHR.html)
- [vkGetImageSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2KHR.html)
- [vkGetRenderingAreaGranularityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRenderingAreaGranularityKHR.html)

## [](#_new_structures)New Structures

- [VkDeviceImageSubresourceInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceImageSubresourceInfoKHR.html)
- [VkImageSubresource2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresource2KHR.html)
- [VkRenderingAreaInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAreaInfoKHR.html)
- [VkSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout2KHR.html)
- Extending [VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferViewCreateInfo.html), [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html), [VkPhysicalDeviceExternalBufferInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalBufferInfo.html), [VkDescriptorBufferBindingInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBufferBindingInfoEXT.html):
  
  - [VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfoKHR.html)
- Extending [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html), [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html), [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html):
  
  - [VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2CreateInfoKHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceMaintenance5FeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance5FeaturesKHR.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceMaintenance5PropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance5PropertiesKHR.html)

## [](#_new_enums)New Enums

- [VkBufferUsageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2KHR.html)
- [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2KHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkBufferUsageFlags2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2KHR.html)
- [VkPipelineCreateFlags2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2KHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_MAINTENANCE_5_EXTENSION_NAME`
- `VK_KHR_MAINTENANCE_5_SPEC_VERSION`
- Extending [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html):
  
  - `VK_BUFFER_USAGE_2_INDEX_BUFFER_BIT_KHR`
  - `VK_BUFFER_USAGE_2_INDIRECT_BUFFER_BIT_KHR`
  - `VK_BUFFER_USAGE_2_STORAGE_BUFFER_BIT_KHR`
  - `VK_BUFFER_USAGE_2_STORAGE_TEXEL_BUFFER_BIT_KHR`
  - `VK_BUFFER_USAGE_2_TRANSFER_DST_BIT_KHR`
  - `VK_BUFFER_USAGE_2_TRANSFER_SRC_BIT_KHR`
  - `VK_BUFFER_USAGE_2_UNIFORM_BUFFER_BIT_KHR`
  - `VK_BUFFER_USAGE_2_UNIFORM_TEXEL_BUFFER_BIT_KHR`
  - `VK_BUFFER_USAGE_2_VERTEX_BUFFER_BIT_KHR`
- Extending [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html):
  
  - `VK_FORMAT_A1B5G5R5_UNORM_PACK16_KHR`
  - `VK_FORMAT_A8_UNORM_KHR`
- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_ALLOW_DERIVATIVES_BIT_KHR`
  - `VK_PIPELINE_CREATE_2_DERIVATIVE_BIT_KHR`
  - `VK_PIPELINE_CREATE_2_DISABLE_OPTIMIZATION_BIT_KHR`
  - `VK_PIPELINE_CREATE_2_DISPATCH_BASE_BIT_KHR`
  - `VK_PIPELINE_CREATE_2_VIEW_INDEX_FROM_DEVICE_INDEX_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BUFFER_USAGE_FLAGS_2_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_DEVICE_IMAGE_SUBRESOURCE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_IMAGE_SUBRESOURCE_2_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_5_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_5_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_PIPELINE_CREATE_FLAGS_2_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_RENDERING_AREA_INFO_KHR`
  - `VK_STRUCTURE_TYPE_SUBRESOURCE_LAYOUT_2_KHR`

If [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html) or [Vulkan Version 1.3](#versions-1.3) and [VK\_EXT\_fragment\_density\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_fragment_density_map.html) is supported:

- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_RENDERING_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT`

If [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html) or [Vulkan Version 1.3](#versions-1.3) and [VK\_KHR\_fragment\_shading\_rate](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_fragment_shading_rate.html) is supported:

- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_RENDERING_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

If [VK\_ARM\_pipeline\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_pipeline_opacity_micromap.html) is supported:

- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_DISALLOW_OPACITY_MICROMAP_BIT_ARM`

If [VK\_EXT\_attachment\_feedback\_loop\_layout](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_attachment_feedback_loop_layout.html) is supported:

- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_COLOR_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`
  - `VK_PIPELINE_CREATE_2_DEPTH_STENCIL_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`

If [VK\_EXT\_conditional\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_conditional_rendering.html) is supported:

- Extending [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html):
  
  - `VK_BUFFER_USAGE_2_CONDITIONAL_RENDERING_BIT_EXT`

If [VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html) is supported:

- Extending [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html):
  
  - `VK_BUFFER_USAGE_2_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT`
  - `VK_BUFFER_USAGE_2_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT`
  - `VK_BUFFER_USAGE_2_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT`
- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_DESCRIPTOR_BUFFER_BIT_EXT`

If [VK\_EXT\_graphics\_pipeline\_library](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_graphics_pipeline_library.html) is supported:

- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_LINK_TIME_OPTIMIZATION_BIT_EXT`
  - `VK_PIPELINE_CREATE_2_RETAIN_LINK_TIME_OPTIMIZATION_INFO_BIT_EXT`

If [VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html) is supported:

- Extending [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html):
  
  - `VK_BUFFER_USAGE_2_MICROMAP_BUILD_INPUT_READ_ONLY_BIT_EXT`
  - `VK_BUFFER_USAGE_2_MICROMAP_STORAGE_BIT_EXT`
- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_RAY_TRACING_OPACITY_MICROMAP_BIT_EXT`

If [VK\_EXT\_transform\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_transform_feedback.html) is supported:

- Extending [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html):
  
  - `VK_BUFFER_USAGE_2_TRANSFORM_FEEDBACK_BUFFER_BIT_EXT`
  - `VK_BUFFER_USAGE_2_TRANSFORM_FEEDBACK_COUNTER_BUFFER_BIT_EXT`

If [VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html) is supported:

- Extending [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html):
  
  - `VK_BUFFER_USAGE_2_ACCELERATION_STRUCTURE_BUILD_INPUT_READ_ONLY_BIT_KHR`
  - `VK_BUFFER_USAGE_2_ACCELERATION_STRUCTURE_STORAGE_BIT_KHR`

If [VK\_KHR\_pipeline\_executable\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_executable_properties.html) is supported:

- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR`
  - `VK_PIPELINE_CREATE_2_CAPTURE_STATISTICS_BIT_KHR`

If [VK\_KHR\_pipeline\_library](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_library.html) is supported:

- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_LIBRARY_BIT_KHR`

If [VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html) is supported:

- Extending [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html):
  
  - `VK_BUFFER_USAGE_2_SHADER_BINDING_TABLE_BIT_KHR`
- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`
  - `VK_PIPELINE_CREATE_2_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`
  - `VK_PIPELINE_CREATE_2_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`
  - `VK_PIPELINE_CREATE_2_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`
  - `VK_PIPELINE_CREATE_2_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`
  - `VK_PIPELINE_CREATE_2_RAY_TRACING_SKIP_AABBS_BIT_KHR`
  - `VK_PIPELINE_CREATE_2_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR`

If [VK\_KHR\_video\_decode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_queue.html) is supported:

- Extending [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html):
  
  - `VK_BUFFER_USAGE_2_VIDEO_DECODE_DST_BIT_KHR`
  - `VK_BUFFER_USAGE_2_VIDEO_DECODE_SRC_BIT_KHR`

If [VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html) is supported:

- Extending [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html):
  
  - `VK_BUFFER_USAGE_2_VIDEO_ENCODE_DST_BIT_KHR`
  - `VK_BUFFER_USAGE_2_VIDEO_ENCODE_SRC_BIT_KHR`

If [VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html) is supported:

- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_INDIRECT_BINDABLE_BIT_NV`

If [VK\_NV\_displacement\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_displacement_micromap.html) is supported:

- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_RAY_TRACING_DISPLACEMENT_MICROMAP_BIT_NV`

If [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) is supported:

- Extending [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html):
  
  - `VK_BUFFER_USAGE_2_RAY_TRACING_BIT_NV`
- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_DEFER_COMPILE_BIT_NV`

If [VK\_NV\_ray\_tracing\_motion\_blur](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing_motion_blur.html) is supported:

- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_RAY_TRACING_ALLOW_MOTION_BIT_NV`

If [Vulkan Version 1.2](#versions-1.2) or [VK\_KHR\_buffer\_device\_address](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_buffer_device_address.html) or [VK\_EXT\_buffer\_device\_address](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_buffer_device_address.html) is supported:

- Extending [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html):
  
  - `VK_BUFFER_USAGE_2_SHADER_DEVICE_ADDRESS_BIT_KHR`

If [Vulkan Version 1.3](#versions-1.3) or [VK\_EXT\_pipeline\_creation\_cache\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_creation_cache_control.html) is supported:

- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_EARLY_RETURN_ON_FAILURE_BIT_KHR`
  - `VK_PIPELINE_CREATE_2_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT_KHR`

If [Vulkan Version 1.4](#versions-1.4) or [VK\_EXT\_pipeline\_protected\_access](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_protected_access.html) is supported:

- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_NO_PROTECTED_ACCESS_BIT_EXT`
  - `VK_PIPELINE_CREATE_2_PROTECTED_ACCESS_ONLY_BIT_EXT`

## [](#_promotion_to_vulkan_1_4)Promotion to Vulkan 1.4

Functionality in this extension is included in core Vulkan 1.4 with the KHR suffix omitted. The original type, enum and command names are still available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 1, 2022-12-12 (Stu Smith)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_maintenance5)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0