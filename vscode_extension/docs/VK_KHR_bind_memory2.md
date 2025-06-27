# VK_KHR_bind_memory2(3) Manual Page

## Name

VK_KHR_bind_memory2 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

158

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
  target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- Tobias Hector <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_bind_memory2%5D%20@tobski%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_bind_memory2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>tobski</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-09-05

**IP Status**  
No known IP claims.

**Contributors**  
- Jeff Bolz, NVIDIA

- Tobias Hector, Imagination Technologies

## <a href="#_description" class="anchor"></a>Description

This extension provides versions of
[vkBindBufferMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindBufferMemory.html) and
[vkBindImageMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindImageMemory.html) that allow multiple bindings
to be performed at once, and are extensible.

This extension also introduces `VK_IMAGE_CREATE_ALIAS_BIT_KHR`, which
allows “identical” images that alias the same memory to interpret the
contents consistently, even across image layout changes.

## <a href="#_promotion_to_vulkan_1_1" class="anchor"></a>Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkBindBufferMemory2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindBufferMemory2KHR.html)

- [vkBindImageMemory2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindImageMemory2KHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkBindBufferMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryInfoKHR.html)

- [VkBindImageMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_BIND_MEMORY_2_EXTENSION_NAME`

- `VK_KHR_BIND_MEMORY_2_SPEC_VERSION`

- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html):

  - `VK_IMAGE_CREATE_ALIAS_BIT_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO_KHR`

  - `VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-05-19 (Tobias Hector)

  - Pulled bind memory functions into their own extension

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_bind_memory2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
