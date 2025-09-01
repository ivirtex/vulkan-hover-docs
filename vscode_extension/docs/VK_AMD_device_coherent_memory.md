# VK\_AMD\_device\_coherent\_memory(3) Manual Page

## Name

VK\_AMD\_device\_coherent\_memory - device extension



## [](#_registered_extension_number)Registered Extension Number

230

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Tobias Hector [\[GitHub\]tobski](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_device_coherent_memory%5D%20%40tobski%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_device_coherent_memory%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-02-04

**Contributors**

- Ping Fu, AMD
- Timothy Lottes, AMD
- Tobias Hector, AMD

## [](#_description)Description

This extension adds the device coherent and device uncached memory types. Any device accesses to device coherent memory are automatically made visible to any other device access. Device uncached memory indicates to applications that caches are disabled for a particular memory type, which guarantees device coherence.

Device coherent and uncached memory are expected to have lower performance for general access than non-device coherent memory, but can be useful in certain scenarios; particularly so for debugging.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceCoherentMemoryFeaturesAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCoherentMemoryFeaturesAMD.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_AMD_DEVICE_COHERENT_MEMORY_EXTENSION_NAME`
- `VK_AMD_DEVICE_COHERENT_MEMORY_SPEC_VERSION`
- Extending [VkMemoryPropertyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryPropertyFlagBits.html):
  
  - `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD`
  - `VK_MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COHERENT_MEMORY_FEATURES_AMD`

## [](#_version_history)Version History

- Revision 1, 2019-02-04 (Tobias Hector)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_AMD_device_coherent_memory).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0