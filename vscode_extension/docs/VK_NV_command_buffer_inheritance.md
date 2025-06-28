# VK\_NV\_command\_buffer\_inheritance(3) Manual Page

## Name

VK\_NV\_command\_buffer\_inheritance - device extension



## [](#_registered_extension_number)Registered Extension Number

560

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Lujin Wang [\[GitHub\]lujinwangnv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_command_buffer_inheritance%5D%20%40lujinwangnv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_command_buffer_inheritance%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-02-15

**IP Status**

No known IP claims.

**Contributors**

- Piers Daniell, NVIDIA
- Daniel Story, Nintendo

## [](#_description)Description

This extension allows applications to take advantage of the graphics and compute state that remains valid in the queue between executions of submitted command buffers. This works across both primary and secondary command buffers.

The state inherited includes the previously bound pipeline state, previously bound shader objects, previously bound vertex and index buffers, previously bound descriptor sets and push constants, and all previously set dynamic state.

This extension relaxes the requirement that all that state needs to be bound and set after begin command buffer and before the next draw or dispatch.

By not having to set state that has been inherited applications can save both CPU and GPU cycles by not having to set state redundantly, and also have improved flexibility when reusing secondary command buffers.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceCommandBufferInheritanceFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCommandBufferInheritanceFeaturesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_COMMAND_BUFFER_INHERITANCE_EXTENSION_NAME`
- `VK_NV_COMMAND_BUFFER_INHERITANCE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COMMAND_BUFFER_INHERITANCE_FEATURES_NV`

## [](#_issues)Issues

1\) How can the validation layer know when state is valid at draw or dispatch time if it is inherited at execution time?

**RESOLVED**: Validation of invalid state at draw and dispatch time cannot be done while recording those commands. Instead the validation layer will need to keep track of any unset state when draw and dispatch commands are recorded, but not report an error at that time. It should also keep track of what state is valid at the end of each recorded command buffer. When secondary command buffer execution is recorded the validation layer can update its unset state tracking for that command buffer, and also for draw and dispatch commands recorded after execution of the secondary as they will inherit state from the executed secondary. This can be done recursively so every recorded primary command buffer has a final tally of any unset state used at draw and dispatch time. Finally when the primary is submitted to the queue the validation layer will know the previous primaries submitted to the queue and will know if there is any unset state used and can report the error then.

## [](#_version_history)Version History

- Revision 1, 2024-02-15 (Lujin Wang)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_command_buffer_inheritance)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0