# VK\_NV\_copy\_memory\_indirect(3) Manual Page

## Name

VK\_NV\_copy\_memory\_indirect - device extension



## [](#_registered_extension_number)Registered Extension Number

427

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

         [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
         or  
         [Vulkan Version 1.1](#versions-1.1)  
     and  
     [VK\_KHR\_buffer\_device\_address](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_buffer_device_address.html)  
or  
[Vulkan Version 1.2](#versions-1.2)

## [](#_contact)Contact

- Vikram Kushwaha [\[GitHub\]vkushwaha-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_copy_memory_indirect%5D%20%40vkushwaha-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_copy_memory_indirect%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-10-14

**Contributors**

- Vikram Kushwaha, NVIDIA
- Jeff Bolz, NVIDIA
- Christoph Kubisch, NVIDIA
- Daniel Koch, NVIDIA

## [](#_description)Description

This extension adds support for performing copies between memory and image regions using indirect parameters that are read by the device from a buffer during execution. This functionality **may** be useful for performing copies where the copy parameters are not known during the command buffer creation time.

## [](#_new_commands)New Commands

- [vkCmdCopyMemoryIndirectNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryIndirectNV.html)
- [vkCmdCopyMemoryToImageIndirectNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryToImageIndirectNV.html)

## [](#_new_structures)New Structures

- [VkCopyMemoryIndirectCommandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryIndirectCommandNV.html)
- [VkCopyMemoryToImageIndirectCommandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageIndirectCommandNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceCopyMemoryIndirectFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCopyMemoryIndirectFeaturesNV.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceCopyMemoryIndirectPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCopyMemoryIndirectPropertiesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_COPY_MEMORY_INDIRECT_EXTENSION_NAME`
- `VK_NV_COPY_MEMORY_INDIRECT_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COPY_MEMORY_INDIRECT_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COPY_MEMORY_INDIRECT_PROPERTIES_NV`

## [](#_version_history)Version History

- Revision 1, 2022-10-14 (Vikram Kushwaha)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_copy_memory_indirect).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0