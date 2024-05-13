# VK_KHR_get_memory_requirements2(3) Manual Page

## Name

VK_KHR_get_memory_requirements2 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

147

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

- Faith Ekstrand <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_get_memory_requirements2%5D%20@gfxstrand%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_get_memory_requirements2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>gfxstrand</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-09-05

**IP Status**  
No known IP claims.

**Contributors**  
- Faith Ekstrand, Intel

- Jeff Bolz, NVIDIA

- Jesse Hall, Google

## <a href="#_description" class="anchor"></a>Description

This extension provides new queries for memory requirements of images
and buffers that can be easily extended by other extensions, without
introducing any further entry points. The Vulkan 1.0
[VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) and
[VkSparseImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryRequirements.html)
structures do not include `sType` and `pNext` members. This extension
wraps them in new structures with these members, so an application can
query a chain of memory requirements structures by constructing the
chain and letting the implementation fill them in. A new command is
added for each `vkGet*MemoryRequrements` command in core Vulkan 1.0.

## <a href="#_promotion_to_vulkan_1_1" class="anchor"></a>Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetBufferMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferMemoryRequirements2KHR.html)

- [vkGetImageMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements2KHR.html)

- [vkGetImageSparseMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSparseMemoryRequirements2KHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkBufferMemoryRequirementsInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryRequirementsInfo2KHR.html)

- [VkImageMemoryRequirementsInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryRequirementsInfo2KHR.html)

- [VkImageSparseMemoryRequirementsInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSparseMemoryRequirementsInfo2KHR.html)

- [VkMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2KHR.html)

- [VkSparseImageMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryRequirements2KHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_GET_MEMORY_REQUIREMENTS_2_EXTENSION_NAME`

- `VK_KHR_GET_MEMORY_REQUIREMENTS_2_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2_KHR`

  - `VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2_KHR`

  - `VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2_KHR`

  - `VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2_KHR`

  - `VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-03-23 (Faith Ekstrand)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_get_memory_requirements2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
