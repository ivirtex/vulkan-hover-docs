# VK\_KHR\_synchronization2(3) Manual Page

## Name

VK\_KHR\_synchronization2 - device extension



## [](#_registered_extension_number)Registered Extension Number

315

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VK\_EXT\_blend\_operation\_advanced
- Interacts with VK\_EXT\_conditional\_rendering
- Interacts with VK\_EXT\_device\_generated\_commands
- Interacts with VK\_EXT\_fragment\_density\_map
- Interacts with VK\_EXT\_mesh\_shader
- Interacts with VK\_EXT\_transform\_feedback
- Interacts with VK\_KHR\_acceleration\_structure
- Interacts with VK\_KHR\_fragment\_shading\_rate
- Interacts with VK\_KHR\_ray\_tracing\_pipeline
- Interacts with VK\_NV\_device\_generated\_commands
- Interacts with VK\_NV\_mesh\_shader
- Interacts with VK\_NV\_ray\_tracing
- Interacts with VK\_NV\_shading\_rate\_image

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.3](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3-promotions)

## [](#_contact)Contact

- Tobias Hector [\[GitHub\]tobski](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_synchronization2%5D%20%40tobski%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_synchronization2%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-12-03

**Interactions and External Dependencies**

- Interacts with `VK_KHR_create_renderpass2`

**Contributors**

- Tobias Hector

## [](#_description)Description

This extension modifies the original core synchronization APIs to simplify the interface and improve usability of these APIs. It also adds new pipeline stage and access flag types that extend into the 64-bit range, as we have run out within the 32-bit range. The new flags are identical to the old values within the 32-bit range, with new stages and bits beyond that.

Pipeline stages and access flags are now specified together in memory barrier structures, making the connection between the two more obvious. Additionally, scoping the pipeline stages into the barrier structs allows the use of the `MEMORY_READ` and `MEMORY_WRITE` flags without sacrificing precision. The per-stage access flags should be used to disambiguate specific accesses in a given stage or set of stages - for instance, between uniform reads and sampling operations.

Layout transitions have been simplified as well; rather than requiring a different set of layouts for depth/stencil/color attachments, there are generic `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` and `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR` layouts which are contextually applied based on the image format. For example, for a depth format image, `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR` is equivalent to `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL_KHR`. `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR` also functionally replaces `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL`.

Events are now more efficient, because they include memory dependency information when you set them on the device. Previously, this information was only known when waiting on an event, so the dependencies could not be satisfied until the wait occurred. That sometimes meant stalling the pipeline when the wait occurred. The new API provides enough information for implementations to satisfy these dependencies in parallel with other tasks.

Queue submission has been changed to wrap command buffers and semaphores in extensible structures, which incorporate changes from Vulkan 1.1, `VK_KHR_device_group`, and `VK_KHR_timeline_semaphore`. This also adds a pipeline stage to the semaphore signal operation, mirroring the existing pipeline stage specification for wait operations.

Other miscellaneous changes include:

- Events can now be specified as interacting only with the device, allowing more efficient access to the underlying object.
- Image memory barriers that do not perform an image layout transition can be specified by setting `oldLayout` equal to `newLayout`.
  
  - E.g. the old and new layout can both be set to `VK_IMAGE_LAYOUT_UNDEFINED`, without discarding data in the image.
- Queue family ownership transfer parameters are simplified in some cases.
- Extensions with commands or functions with a [VkPipelineStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags.html) or [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) parameter have had those APIs replaced with equivalents using [VkPipelineStageFlags2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2KHR.html).
- The new event and barrier interfaces are now more extensible for future changes.
- Relevant pipeline stage masks can now be specified as empty with the new `VK_PIPELINE_STAGE_NONE_KHR` and `VK_PIPELINE_STAGE_2_NONE_KHR` values.
- [VkMemoryBarrier2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrier2KHR.html) can be chained to [VkSubpassDependency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDependency2.html), overriding the original 32-bit stage and access masks.

## [](#_new_base_types)New Base Types

- [VkFlags64](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFlags64.html)

## [](#_new_commands)New Commands

- [vkCmdPipelineBarrier2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPipelineBarrier2KHR.html)
- [vkCmdResetEvent2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdResetEvent2KHR.html)
- [vkCmdSetEvent2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent2KHR.html)
- [vkCmdWaitEvents2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents2KHR.html)
- [vkCmdWriteTimestamp2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteTimestamp2KHR.html)
- [vkQueueSubmit2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit2KHR.html)

## [](#_new_structures)New Structures

- [VkBufferMemoryBarrier2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferMemoryBarrier2KHR.html)
- [VkCommandBufferSubmitInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferSubmitInfoKHR.html)
- [VkDependencyInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfoKHR.html)
- [VkImageMemoryBarrier2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier2KHR.html)
- [VkSemaphoreSubmitInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreSubmitInfoKHR.html)
- [VkSubmitInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo2KHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceSynchronization2FeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSynchronization2FeaturesKHR.html)
- Extending [VkSubpassDependency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDependency2.html):
  
  - [VkMemoryBarrier2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrier2KHR.html)

## [](#_new_enums)New Enums

- [VkAccessFlagBits2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2KHR.html)
- [VkPipelineStageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2KHR.html)
- [VkSubmitFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkAccessFlags2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags2KHR.html)
- [VkPipelineStageFlags2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2KHR.html)
- [VkSubmitFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SYNCHRONIZATION_2_EXTENSION_NAME`
- `VK_KHR_SYNCHRONIZATION_2_SPEC_VERSION`
- Extending [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html):
  
  - `VK_ACCESS_NONE_KHR`
- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT_KHR`
  - `VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT_KHR`
  - `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT_KHR`
  - `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT_KHR`
  - `VK_ACCESS_2_HOST_READ_BIT_KHR`
  - `VK_ACCESS_2_HOST_WRITE_BIT_KHR`
  - `VK_ACCESS_2_INDEX_READ_BIT_KHR`
  - `VK_ACCESS_2_INDIRECT_COMMAND_READ_BIT_KHR`
  - `VK_ACCESS_2_INPUT_ATTACHMENT_READ_BIT_KHR`
  - `VK_ACCESS_2_MEMORY_READ_BIT_KHR`
  - `VK_ACCESS_2_MEMORY_WRITE_BIT_KHR`
  - `VK_ACCESS_2_NONE_KHR`
  - `VK_ACCESS_2_SHADER_READ_BIT_KHR`
  - `VK_ACCESS_2_SHADER_SAMPLED_READ_BIT_KHR`
  - `VK_ACCESS_2_SHADER_STORAGE_READ_BIT_KHR`
  - `VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT_KHR`
  - `VK_ACCESS_2_SHADER_WRITE_BIT_KHR`
  - `VK_ACCESS_2_TRANSFER_READ_BIT_KHR`
  - `VK_ACCESS_2_TRANSFER_WRITE_BIT_KHR`
  - `VK_ACCESS_2_UNIFORM_READ_BIT_KHR`
  - `VK_ACCESS_2_VERTEX_ATTRIBUTE_READ_BIT_KHR`
- Extending [VkEventCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEventCreateFlagBits.html):
  
  - `VK_EVENT_CREATE_DEVICE_ONLY_BIT_KHR`
- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html):
  
  - `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR`
  - `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`
- Extending [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html):
  
  - `VK_PIPELINE_STAGE_NONE_KHR`
- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_BLIT_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_BOTTOM_OF_PIPE_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_CLEAR_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_COPY_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_HOST_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_INDEX_INPUT_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_NONE_KHR`
  - `VK_PIPELINE_STAGE_2_PRE_RASTERIZATION_SHADERS_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_RESOLVE_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_TOP_OF_PIPE_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_TRANSFER_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_VERTEX_ATTRIBUTE_INPUT_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT_KHR`
  - `VK_PIPELINE_STAGE_2_VERTEX_SHADER_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER_2_KHR`
  - `VK_STRUCTURE_TYPE_COMMAND_BUFFER_SUBMIT_INFO_KHR`
  - `VK_STRUCTURE_TYPE_DEPENDENCY_INFO_KHR`
  - `VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER_2_KHR`
  - `VK_STRUCTURE_TYPE_MEMORY_BARRIER_2_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SYNCHRONIZATION_2_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_SEMAPHORE_SUBMIT_INFO_KHR`
  - `VK_STRUCTURE_TYPE_SUBMIT_INFO_2_KHR`
- Extending [VkSubmitFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitFlagBits.html):
  
  - `VK_SUBMIT_PROTECTED_BIT_KHR`

If [VK\_EXT\_blend\_operation\_advanced](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_blend_operation_advanced.html) is supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT`

If [VK\_EXT\_conditional\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_conditional_rendering.html) is supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_CONDITIONAL_RENDERING_READ_BIT_EXT`
- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`

If [VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html) is supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_EXT`
  - `VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_EXT`
- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_EXT`

If [VK\_EXT\_fragment\_density\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_fragment_density_map.html) is supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_FRAGMENT_DENSITY_MAP_READ_BIT_EXT`
- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

If [VK\_EXT\_mesh\_shader](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_mesh_shader.html) is supported:

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`
  - `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`

If [VK\_EXT\_transform\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_transform_feedback.html) is supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT`
  - `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT`
  - `VK_ACCESS_2_TRANSFORM_FEEDBACK_WRITE_BIT_EXT`
- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`

If [VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html) is supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`
  - `VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`
- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`

If [VK\_KHR\_fragment\_shading\_rate](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_fragment_shading_rate.html) is supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_FRAGMENT_SHADING_RATE_ATTACHMENT_READ_BIT_KHR`
- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

If [VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html) is supported:

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`

If [VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html) is supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_NV`
  - `VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_NV`
- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV`

If [VK\_NV\_mesh\_shader](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_mesh_shader.html) is supported:

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_NV`
  - `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_NV`

If [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) is supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_NV`
  - `VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_NV`
- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_NV`
  - `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_NV`

If [VK\_NV\_shading\_rate\_image](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_shading_rate_image.html) is supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_SHADING_RATE_IMAGE_READ_BIT_NV`
- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_SHADING_RATE_IMAGE_BIT_NV`

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

Vulkan APIs in this extension are included in core Vulkan 1.3, with the KHR suffix omitted. External interactions defined by this extension, such as SPIR-V token names, retain their original names. The original Vulkan API names are still available as aliases of the core functionality.

## [](#_examples)Examples

See [https://github.com/KhronosGroup/Vulkan-Docs/wiki/Synchronization-Examples](https://github.com/KhronosGroup/Vulkan-Docs/wiki/Synchronization-Examples)

## [](#_version_history)Version History

- Revision 1, 2020-12-03 (Tobias Hector)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_synchronization2)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0