# VK\_EXT\_transform\_feedback(3) Manual Page

## Name

VK\_EXT\_transform\_feedback - device extension



## [](#_registered_extension_number)Registered Extension Number

29

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_special_uses)Special Uses

- [OpenGL / ES support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)
- [D3D support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)
- [Developer tools](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_transform_feedback%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_transform_feedback%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-10-09

**Contributors**

- Baldur Karlsson, Valve
- Boris Zanin, Mobica
- Daniel Rakos, AMD
- Donald Scorgie, Imagination
- Henri Verbeet, CodeWeavers
- Jan-Harald Fredriksen, Arm
- Faith Ekstrand, Intel
- Jeff Bolz, NVIDIA
- Jesse Barker, Unity
- Jesse Hall, Google
- Pierre-Loup Griffais, Valve
- Philip Rebohle, DXVK
- Ruihao Zhang, Qualcomm
- Samuel Pitoiset, Valve
- Slawomir Grajewski, Intel
- Stu Smith, Imagination Technologies

## [](#_description)Description

This extension adds transform feedback to the Vulkan API by exposing the SPIR-V `TransformFeedback` and `GeometryStreams` capabilities to capture vertex, tessellation or geometry shader outputs to one or more buffers. It adds API functionality to bind transform feedback buffers to capture the primitives emitted by the graphics pipeline from SPIR-V outputs decorated for transform feedback. The transform feedback capture can be paused and resumed by way of storing and retrieving a byte counter. The captured data can be drawn again where the vertex count is derived from the byte counter without CPU intervention. If the implementation is capable, a vertex stream other than zero can be rasterized.

All these features are designed to match the full capabilities of OpenGL core transform feedback functionality and beyond. Many of the features are optional to allow base OpenGL ES GPUs to also implement this extension.

The primary purpose of the functionality exposed by this extension is to support translation layers from other 3D APIs. This functionality is not considered forward looking, and is not expected to be promoted to a KHR extension or to core Vulkan. Unless this is needed for translation, it is recommended that developers use alternative techniques of using the GPU to process and capture vertex data.

## [](#_new_commands)New Commands

- [vkCmdBeginQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginQueryIndexedEXT.html)
- [vkCmdBeginTransformFeedbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginTransformFeedbackEXT.html)
- [vkCmdBindTransformFeedbackBuffersEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindTransformFeedbackBuffersEXT.html)
- [vkCmdDrawIndirectByteCountEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndirectByteCountEXT.html)
- [vkCmdEndQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndQueryIndexedEXT.html)
- [vkCmdEndTransformFeedbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndTransformFeedbackEXT.html)

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceTransformFeedbackFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTransformFeedbackFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceTransformFeedbackPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTransformFeedbackPropertiesEXT.html)
- Extending [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html):
  
  - [VkPipelineRasterizationStateStreamCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateStreamCreateInfoEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkPipelineRasterizationStateStreamCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateStreamCreateFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_TRANSFORM_FEEDBACK_EXTENSION_NAME`
- `VK_EXT_TRANSFORM_FEEDBACK_SPEC_VERSION`
- Extending [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html):
  
  - `VK_ACCESS_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT`
  - `VK_ACCESS_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT`
  - `VK_ACCESS_TRANSFORM_FEEDBACK_WRITE_BIT_EXT`
- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html):
  
  - `VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_BUFFER_BIT_EXT`
  - `VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_COUNTER_BUFFER_BIT_EXT`
- Extending [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html):
  
  - `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`
- Extending [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html):
  
  - `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TRANSFORM_FEEDBACK_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TRANSFORM_FEEDBACK_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_STREAM_CREATE_INFO_EXT`

## [](#_issues)Issues

1\) Should we include pause/resume functionality?

**RESOLVED**: Yes, this is needed to ease layering other APIs which have this functionality. To pause use `vkCmdEndTransformFeedbackEXT` and provide valid buffer handles in the `pCounterBuffers` array and offsets in the `pCounterBufferOffsets` array for the implementation to save the resume points. Then to resume use `vkCmdBeginTransformFeedbackEXT` with the previous `pCounterBuffers` and `pCounterBufferOffsets` values. Between the pause and resume there needs to be a memory barrier for the counter buffers with a source access of `VK_ACCESS_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT` at pipeline stage `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT` to a destination access of `VK_ACCESS_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT` at pipeline stage `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`.

2\) How does this interact with multiview?

**RESOLVED**: Transform feedback cannot be made active in a render pass with multiview enabled.

3\) How should queries be done?

**RESOLVED**: There is a new query type `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT`. A query pool created with this type will capture 2 integers - numPrimitivesWritten and numPrimitivesNeeded - for the specified vertex stream output from the last [pre-rasterization shader stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization). The vertex stream output queried is zero by default, but can be specified with the new `vkCmdBeginQueryIndexedEXT` and `vkCmdEndQueryIndexedEXT` commands.

## [](#_version_history)Version History

- Revision 1, 2018-10-09 (Piers Daniell)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_transform_feedback)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0