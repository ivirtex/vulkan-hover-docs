# VK\_NV\_memory\_decompression(3) Manual Page

## Name

VK\_NV\_memory\_decompression - device extension



## [](#_registered_extension_number)Registered Extension Number

428

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

- Vikram Kushwaha [\[GitHub\]vkushwaha-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_memory_decompression%5D%20%40vkushwaha-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_memory_decompression%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-01-31

**Contributors**

- Vikram Kushwaha, NVIDIA
- Jeff Bolz, NVIDIA
- Christoph Kubisch, NVIDIA
- Piers Daniell, NVIDIA

## [](#_description)Description

This extension adds support for performing memory to memory decompression.

## [](#_new_commands)New Commands

- [vkCmdDecompressMemoryIndirectCountNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDecompressMemoryIndirectCountNV.html)
- [vkCmdDecompressMemoryNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDecompressMemoryNV.html)

## [](#_new_structures)New Structures

- [VkDecompressMemoryRegionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDecompressMemoryRegionNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceMemoryDecompressionFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryDecompressionFeaturesNV.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceMemoryDecompressionPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryDecompressionPropertiesNV.html)

## [](#_new_enums)New Enums

- [VkMemoryDecompressionMethodFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDecompressionMethodFlagBitsNV.html)

## [](#_new_bitmasks)New Bitmasks

- [VkMemoryDecompressionMethodFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDecompressionMethodFlagsNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_MEMORY_DECOMPRESSION_EXTENSION_NAME`
- `VK_NV_MEMORY_DECOMPRESSION_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_DECOMPRESSION_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_DECOMPRESSION_PROPERTIES_NV`

## [](#_version_history)Version History

- Revision 1, 2022-01-31 (Vikram Kushwaha)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_memory_decompression)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0