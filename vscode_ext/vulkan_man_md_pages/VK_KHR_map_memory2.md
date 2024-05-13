# VK_KHR_map_memory2(3) Manual Page

## Name

VK_KHR_map_memory2 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

272

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Faith Ekstrand <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_map_memory2%5D%20@gfxstrand%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_map_memory2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>gfxstrand</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_map_memory2](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_map_memory2.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-03-14

**Interactions and External Dependencies**  
- None

**Contributors**  
- Faith Ekstrand, Collabora

- Tobias Hector, AMD

## <a href="#_description" class="anchor"></a>Description

This extension provides extensible versions of the Vulkan memory map and
unmap entry points. The new entry points are functionally identical to
the core entry points, except that their parameters are specified using
extensible structures that can be used to pass extension-specific
information.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkMapMemory2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkMapMemory2KHR.html)

- [vkUnmapMemory2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUnmapMemory2KHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkMemoryMapInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapInfoKHR.html)

- [VkMemoryUnmapInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryUnmapInfoKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkMemoryUnmapFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryUnmapFlagBitsKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkMemoryUnmapFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryUnmapFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_MAP_MEMORY_2_EXTENSION_NAME`

- `VK_KHR_MAP_MEMORY_2_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_MEMORY_MAP_INFO_KHR`

  - `VK_STRUCTURE_TYPE_MEMORY_UNMAP_INFO_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 0, 2022-08-03 (Faith Ekstrand)

  - Internal revisions

- Revision 1, 2023-03-14

  - Public release

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_map_memory2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
