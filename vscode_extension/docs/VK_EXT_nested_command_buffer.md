# VK\_EXT\_nested\_command\_buffer(3) Manual Page

## Name

VK\_EXT\_nested\_command\_buffer - device extension



## [](#_registered_extension_number)Registered Extension Number

452

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_nested_command_buffer%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_nested_command_buffer%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-09-18

**Contributors**

- Daniel Story, Nintendo
- Peter Kohaut, NVIDIA
- Shahbaz Youssefi, Google
- Slawomir Grajewski, Intel
- Stu Smith, AMD

## [](#_description)Description

With core Vulkan it is not legal to call [vkCmdExecuteCommands](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdExecuteCommands.html) when recording a secondary command buffer. This extension relaxes that restriction, allowing secondary command buffers to execute other secondary command buffers.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceNestedCommandBufferFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceNestedCommandBufferFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceNestedCommandBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceNestedCommandBufferPropertiesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_NESTED_COMMAND_BUFFER_EXTENSION_NAME`
- `VK_EXT_NESTED_COMMAND_BUFFER_SPEC_VERSION`
- Extending [VkRenderingFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFlagBits.html):
  
  - `VK_RENDERING_CONTENTS_INLINE_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_NESTED_COMMAND_BUFFER_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_NESTED_COMMAND_BUFFER_PROPERTIES_EXT`
- Extending [VkSubpassContents](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassContents.html):
  
  - `VK_SUBPASS_CONTENTS_INLINE_AND_SECONDARY_COMMAND_BUFFERS_EXT`

## [](#_issues)Issues

1\) The Command Buffer Levels property for the Vulkan commands comes from the `cmdbufferlevel` attribute in `vk.xml` for the command, and it is currently not possible to modify this attribute based on whether an extension is enabled. For this extension we want the `cmdbufferlevel` attribute for vkCmdExecuteCommands to be `primary,secondary` when this extension is enabled and `primary` otherwise.

**RESOLVED**: The `cmdbufferlevel` attribute for [vkCmdExecuteCommands](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdExecuteCommands.html) has been changed to `primary,secondary` and a new VUID added to prohibit recording this command in a secondary command buffer unless this extension is enabled.

## [](#_version_history)Version History

- Revision 1, 2023-09-18 (Piers Daniell)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_nested_command_buffer)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0