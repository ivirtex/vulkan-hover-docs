# VK_AMD_device_coherent_memory(3) Manual Page

## Name

VK_AMD_device_coherent_memory - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

230

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Tobias Hector <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_device_coherent_memory%5D%20@tobski%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_device_coherent_memory%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>tobski</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-02-04

**Contributors**  
- Ping Fu, AMD

- Timothy Lottes, AMD

- Tobias Hector, AMD

## <a href="#_description" class="anchor"></a>Description

This extension adds the device coherent and device uncached memory
types. Any device accesses to device coherent memory are automatically
made visible to any other device access. Device uncached memory
indicates to applications that caches are disabled for a particular
memory type, which guarantees device coherence.

Device coherent and uncached memory are expected to have lower
performance for general access than non-device coherent memory, but can
be useful in certain scenarios; particularly so for debugging.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceCoherentMemoryFeaturesAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCoherentMemoryFeaturesAMD.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_AMD_DEVICE_COHERENT_MEMORY_EXTENSION_NAME`

- `VK_AMD_DEVICE_COHERENT_MEMORY_SPEC_VERSION`

- Extending [VkMemoryPropertyFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryPropertyFlagBits.html):

  - `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD`

  - `VK_MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COHERENT_MEMORY_FEATURES_AMD`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-02-04 (Tobias Hector)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_AMD_device_coherent_memory"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
