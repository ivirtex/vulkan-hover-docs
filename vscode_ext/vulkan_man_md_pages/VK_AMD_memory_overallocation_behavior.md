# VK_AMD_memory_overallocation_behavior(3) Manual Page

## Name

VK_AMD_memory_overallocation_behavior - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

190

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Martin Dinkov <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_memory_overallocation_behavior%5D%20@mdinkov%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_memory_overallocation_behavior%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mdinkov</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-09-19

**IP Status**  
No known IP claims.

**Contributors**  
- Martin Dinkov, AMD

- Matthaeus Chajdas, AMD

- Daniel Rakos, AMD

- Jon Campbell, AMD

## <a href="#_description" class="anchor"></a>Description

This extension allows controlling whether explicit overallocation beyond
the device memory heap sizes (reported by
[VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties.html))
is allowed or not. Overallocation may lead to performance loss and is
not supported for all platforms.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkDeviceMemoryOverallocationCreateInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryOverallocationCreateInfoAMD.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkMemoryOverallocationBehaviorAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryOverallocationBehaviorAMD.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_AMD_MEMORY_OVERALLOCATION_BEHAVIOR_EXTENSION_NAME`

- `VK_AMD_MEMORY_OVERALLOCATION_BEHAVIOR_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DEVICE_MEMORY_OVERALLOCATION_CREATE_INFO_AMD`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-09-19 (Martin Dinkov)

  - Initial draft.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_AMD_memory_overallocation_behavior"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
