# VK_KHR_synchronization2(3) Manual Page

## Name

VK_KHR_synchronization2 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

315

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_AMD_buffer_marker

- Interacts with VK_EXT_blend_operation_advanced

- Interacts with VK_EXT_conditional_rendering

- Interacts with VK_EXT_fragment_density_map

- Interacts with VK_EXT_mesh_shader

- Interacts with VK_EXT_transform_feedback

- Interacts with VK_KHR_acceleration_structure

- Interacts with VK_KHR_fragment_shading_rate

- Interacts with VK_KHR_ray_tracing_pipeline

- Interacts with VK_NV_device_diagnostic_checkpoints

- Interacts with VK_NV_device_generated_commands

- Interacts with VK_NV_mesh_shader

- Interacts with VK_NV_ray_tracing

- Interacts with VK_NV_shading_rate_image

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.3-promotions"
  target="_blank" rel="noopener">Vulkan 1.3</a>

## <a href="#_contact" class="anchor"></a>Contact

- Tobias Hector <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_synchronization2%5D%20@tobski%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_synchronization2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>tobski</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-12-03

**Interactions and External Dependencies**  
- Interacts with
  [`VK_KHR_create_renderpass2`](VK_KHR_create_renderpass2.html)

**Contributors**  
- Tobias Hector

## <a href="#_description" class="anchor"></a>Description

This extension modifies the original core synchronization APIs to
simplify the interface and improve usability of these APIs. It also adds
new pipeline stage and access flag types that extend into the 64-bit
range, as we have run out within the 32-bit range. The new flags are
identical to the old values within the 32-bit range, with new stages and
bits beyond that.

Pipeline stages and access flags are now specified together in memory
barrier structures, making the connection between the two more obvious.
Additionally, scoping the pipeline stages into the barrier structs
allows the use of the `MEMORY_READ` and `MEMORY_WRITE` flags without
sacrificing precision. The per-stage access flags should be used to
disambiguate specific accesses in a given stage or set of stages - for
instance, between uniform reads and sampling operations.

Layout transitions have been simplified as well; rather than requiring a
different set of layouts for depth/stencil/color attachments, there are
generic `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` and
`VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR` layouts which are contextually
applied based on the image format. For example, for a depth format
image, `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR` is equivalent to
`VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL_KHR`.
`VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR` also functionally replaces
`VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL`.

Events are now more efficient, because they include memory dependency
information when you set them on the device. Previously, this
information was only known when waiting on an event, so the dependencies
could not be satisfied until the wait occurred. That sometimes meant
stalling the pipeline when the wait occurred. The new API provides
enough information for implementations to satisfy these dependencies in
parallel with other tasks.

Queue submission has been changed to wrap command buffers and semaphores
in extensible structures, which incorporate changes from Vulkan 1.1,
[`VK_KHR_device_group`](VK_KHR_device_group.html), and
[`VK_KHR_timeline_semaphore`](VK_KHR_timeline_semaphore.html). This also
adds a pipeline stage to the semaphore signal operation, mirroring the
existing pipeline stage specification for wait operations.

Other miscellaneous changes include:

- Events can now be specified as interacting only with the device,
  allowing more efficient access to the underlying object.

- Image memory barriers that do not perform an image layout transition
  can be specified by setting `oldLayout` equal to `newLayout`.

  - E.g. the old and new layout can both be set to
    `VK_IMAGE_LAYOUT_UNDEFINED`, without discarding data in the image.

- Queue family ownership transfer parameters are simplified in some
  cases.

- Extensions with commands or functions with a
  [VkPipelineStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags.html) or
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) parameter have
  had those APIs replaced with equivalents using
  [VkPipelineStageFlags2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags2KHR.html).

- The new event and barrier interfaces are now more extensible for
  future changes.

- Relevant pipeline stage masks can now be specified as empty with the
  new `VK_PIPELINE_STAGE_NONE_KHR` and `VK_PIPELINE_STAGE_2_NONE_KHR`
  values.

- [VkMemoryBarrier2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryBarrier2KHR.html) can be chained to
  [VkSubpassDependency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDependency2.html), overriding the
  original 32-bit stage and access masks.

## <a href="#_new_base_types" class="anchor"></a>New Base Types

- [VkFlags64](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFlags64.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdPipelineBarrier2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPipelineBarrier2KHR.html)

- [vkCmdResetEvent2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdResetEvent2KHR.html)

- [vkCmdSetEvent2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent2KHR.html)

- [vkCmdWaitEvents2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents2KHR.html)

- [vkCmdWriteTimestamp2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteTimestamp2KHR.html)

- [vkQueueSubmit2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueSubmit2KHR.html)

If [VK_AMD_buffer_marker](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_buffer_marker.html) is supported:

- [vkCmdWriteBufferMarker2AMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteBufferMarker2AMD.html)

If
[VK_NV_device_diagnostic_checkpoints](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_diagnostic_checkpoints.html)
is supported:

- [vkGetQueueCheckpointData2NV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetQueueCheckpointData2NV.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkBufferMemoryBarrier2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryBarrier2KHR.html)

- [VkCommandBufferSubmitInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferSubmitInfoKHR.html)

- [VkDependencyInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyInfoKHR.html)

- [VkImageMemoryBarrier2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier2KHR.html)

- [VkSemaphoreSubmitInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreSubmitInfoKHR.html)

- [VkSubmitInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo2KHR.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceSynchronization2FeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSynchronization2FeaturesKHR.html)

- Extending [VkSubpassDependency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDependency2.html):

  - [VkMemoryBarrier2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryBarrier2KHR.html)

If
[VK_NV_device_diagnostic_checkpoints](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_diagnostic_checkpoints.html)
is supported:

- [VkCheckpointData2NV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCheckpointData2NV.html)

- Extending [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties2.html):

  - [VkQueueFamilyCheckpointProperties2NV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyCheckpointProperties2NV.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkAccessFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2KHR.html)

- [VkPipelineStageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2KHR.html)

- [VkSubmitFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitFlagBitsKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkAccessFlags2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlags2KHR.html)

- [VkPipelineStageFlags2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags2KHR.html)

- [VkSubmitFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SYNCHRONIZATION_2_EXTENSION_NAME`

- `VK_KHR_SYNCHRONIZATION_2_SPEC_VERSION`

- Extending [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits.html):

  - `VK_ACCESS_NONE_KHR`

- Extending [VkEventCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEventCreateFlagBits.html):

  - `VK_EVENT_CREATE_DEVICE_ONLY_BIT_KHR`

- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html):

  - `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR`

  - `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`

- Extending [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html):

  - `VK_PIPELINE_STAGE_NONE_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER_2_KHR`

  - `VK_STRUCTURE_TYPE_COMMAND_BUFFER_SUBMIT_INFO_KHR`

  - `VK_STRUCTURE_TYPE_DEPENDENCY_INFO_KHR`

  - `VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER_2_KHR`

  - `VK_STRUCTURE_TYPE_MEMORY_BARRIER_2_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SYNCHRONIZATION_2_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_SEMAPHORE_SUBMIT_INFO_KHR`

  - `VK_STRUCTURE_TYPE_SUBMIT_INFO_2_KHR`

If
[VK_EXT_blend_operation_advanced](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_blend_operation_advanced.html)
is supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2.html):

  - `VK_ACCESS_2_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT`

If [VK_EXT_conditional_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_conditional_rendering.html) is
supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2.html):

  - `VK_ACCESS_2_CONDITIONAL_RENDERING_READ_BIT_EXT`

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html):

  - `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`

If [VK_EXT_fragment_density_map](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_fragment_density_map.html) is
supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2.html):

  - `VK_ACCESS_2_FRAGMENT_DENSITY_MAP_READ_BIT_EXT`

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html):

  - `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

If [VK_EXT_mesh_shader](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_mesh_shader.html) is supported:

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html):

  - `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`

  - `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`

If [VK_EXT_transform_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_transform_feedback.html) is
supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2.html):

  - `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT`

  - `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT`

  - `VK_ACCESS_2_TRANSFORM_FEEDBACK_WRITE_BIT_EXT`

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html):

  - `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`

If [VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html)
is supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2.html):

  - `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`

  - `VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html):

  - `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`

If [VK_KHR_fragment_shading_rate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_fragment_shading_rate.html) is
supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2.html):

  - `VK_ACCESS_2_FRAGMENT_SHADING_RATE_ATTACHMENT_READ_BIT_KHR`

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html):

  - `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

If [VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html) is
supported:

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html):

  - `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`

If
[VK_NV_device_diagnostic_checkpoints](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_diagnostic_checkpoints.html)
is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_CHECKPOINT_DATA_2_NV`

  - `VK_STRUCTURE_TYPE_QUEUE_FAMILY_CHECKPOINT_PROPERTIES_2_NV`

If
[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html)
is supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2.html):

  - `VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_NV`

  - `VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_NV`

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html):

  - `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV`

If [VK_NV_mesh_shader](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_mesh_shader.html) is supported:

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html):

  - `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_NV`

  - `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_NV`

If [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) is supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2.html):

  - `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_NV`

  - `VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_NV`

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html):

  - `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_NV`

  - `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_NV`

If [VK_NV_shading_rate_image](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_shading_rate_image.html) is
supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2.html):

  - `VK_ACCESS_2_SHADING_RATE_IMAGE_READ_BIT_NV`

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html):

  - `VK_PIPELINE_STAGE_2_SHADING_RATE_IMAGE_BIT_NV`

## <a href="#_promotion_to_vulkan_1_3" class="anchor"></a>Promotion to Vulkan 1.3

Functionality in this extension is included in core Vulkan 1.3, with the
KHR suffix omitted. The original type, enum and command names are still
available as aliases of the core functionality.

## <a href="#_examples" class="anchor"></a>Examples

See <a
href="https://github.com/KhronosGroup/Vulkan-Docs/wiki/Synchronization-Examples"
class="bare">https://github.com/KhronosGroup/Vulkan-Docs/wiki/Synchronization-Examples</a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-12-03 (Tobias Hector)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_synchronization2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
