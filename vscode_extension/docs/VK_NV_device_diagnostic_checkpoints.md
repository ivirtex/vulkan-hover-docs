# VK\_NV\_device\_diagnostic\_checkpoints(3) Manual Page

## Name

VK\_NV\_device\_diagnostic\_checkpoints - device extension



## [](#_registered_extension_number)Registered Extension Number

207

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_3
- Interacts with VK\_KHR\_synchronization2

## [](#_contact)Contact

- Nuno Subtil [\[GitHub\]nsubtil](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_device_diagnostic_checkpoints%5D%20%40nsubtil%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_device_diagnostic_checkpoints%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-07-16

**Contributors**

- Oleg Kuznetsov, NVIDIA
- Alex Dunn, NVIDIA
- Jeff Bolz, NVIDIA
- Eric Werness, NVIDIA
- Daniel Koch, NVIDIA

## [](#_description)Description

This extension allows applications to insert markers in the command stream and associate them with custom data.

If a device lost error occurs, the application **may** then query the implementation for the last markers to cross specific implementation-defined pipeline stages, in order to narrow down which commands were executing at the time and might have caused the failure.

## [](#_new_commands)New Commands

- [vkCmdSetCheckpointNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCheckpointNV.html)
- [vkGetQueueCheckpointDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetQueueCheckpointDataNV.html)

If [Vulkan Version 1.3](#versions-1.3) or [VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html) is supported:

- [vkGetQueueCheckpointData2NV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetQueueCheckpointData2NV.html)

## [](#_new_structures)New Structures

- [VkCheckpointDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCheckpointDataNV.html)
- Extending [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties2.html):
  
  - [VkQueueFamilyCheckpointPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyCheckpointPropertiesNV.html)

If [Vulkan Version 1.3](#versions-1.3) or [VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html) is supported:

- [VkCheckpointData2NV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCheckpointData2NV.html)
- Extending [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties2.html):
  
  - [VkQueueFamilyCheckpointProperties2NV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyCheckpointProperties2NV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_DEVICE_DIAGNOSTIC_CHECKPOINTS_EXTENSION_NAME`
- `VK_NV_DEVICE_DIAGNOSTIC_CHECKPOINTS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_CHECKPOINT_DATA_NV`
  - `VK_STRUCTURE_TYPE_QUEUE_FAMILY_CHECKPOINT_PROPERTIES_NV`

If [Vulkan Version 1.3](#versions-1.3) or [VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_CHECKPOINT_DATA_2_NV`
  - `VK_STRUCTURE_TYPE_QUEUE_FAMILY_CHECKPOINT_PROPERTIES_2_NV`

## [](#_version_history)Version History

- Revision 1, 2018-07-16 (Nuno Subtil)
  
  - Internal revisions
- Revision 2, 2018-07-16 (Nuno Subtil)
  
  - ???

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_device_diagnostic_checkpoints)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0