# VK\_NV\_scissor\_exclusive(3) Manual Page

## Name

VK\_NV\_scissor\_exclusive - device extension



## [](#_registered_extension_number)Registered Extension Number

206

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Pat Brown [\[GitHub\]nvpbrown](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_scissor_exclusive%5D%20%40nvpbrown%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_scissor_exclusive%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-01-18

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

None

**Contributors**

- Pat Brown, NVIDIA
- Jeff Bolz, NVIDIA
- Piers Daniell, NVIDIA
- Daniel Koch, NVIDIA

## [](#_description)Description

This extension adds support for an exclusive scissor test to Vulkan. The exclusive scissor test behaves like the scissor test, except that the exclusive scissor test fails for pixels inside the corresponding rectangle and passes for pixels outside the rectangle. If the same rectangle is used for both the scissor and exclusive scissor tests, the exclusive scissor test will pass if and only if the scissor test fails.

Version 2 of this extension introduces `VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_ENABLE_NV` and [vkCmdSetExclusiveScissorEnableNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetExclusiveScissorEnableNV.html). Applications that use this dynamic state must ensure the implementation advertises at least `specVersion` `2` of this extension.

## [](#_new_commands)New Commands

- [vkCmdSetExclusiveScissorEnableNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetExclusiveScissorEnableNV.html)
- [vkCmdSetExclusiveScissorNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetExclusiveScissorNV.html)

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceExclusiveScissorFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExclusiveScissorFeaturesNV.html)
- Extending [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportStateCreateInfo.html):
  
  - [VkPipelineViewportExclusiveScissorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportExclusiveScissorStateCreateInfoNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_SCISSOR_EXCLUSIVE_EXTENSION_NAME`
- `VK_NV_SCISSOR_EXCLUSIVE_SPEC_VERSION`
- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_ENABLE_NV`
  - `VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_NV`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXCLUSIVE_SCISSOR_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_EXCLUSIVE_SCISSOR_STATE_CREATE_INFO_NV`

## [](#_issues)Issues

1\) For the scissor test, the viewport state must be created with a matching number of scissor and viewport rectangles. Should we have the same requirement for exclusive scissors?

**RESOLVED**: For exclusive scissors, we relax this requirement and allow an exclusive scissor rectangle count that is either zero or equal to the number of viewport rectangles. If you pass in an exclusive scissor count of zero, the exclusive scissor test is treated as disabled.

## [](#_version_history)Version History

- Revision 2, 2023-01-18 (Piers Daniell)
  
  - Add dynamic state for explicit exclusive scissor enables
- Revision 1, 2018-07-31 (Pat Brown)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_scissor_exclusive).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0