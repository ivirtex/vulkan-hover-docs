# VK\_KHR\_get\_memory\_requirements2(3) Manual Page

## Name

VK\_KHR\_get\_memory\_requirements2 - device extension



## [](#_registered_extension_number)Registered Extension Number

147

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- Faith Ekstrand [\[GitHub\]gfxstrand](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_get_memory_requirements2%5D%20%40gfxstrand%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_get_memory_requirements2%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-09-05

**IP Status**

No known IP claims.

**Contributors**

- Faith Ekstrand, Intel
- Jeff Bolz, NVIDIA
- Jesse Hall, Google

## [](#_description)Description

This extension provides new queries for memory requirements of images and buffers that can be easily extended by other extensions, without introducing any additional commands. The Vulkan 1.0 [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) and [VkSparseImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryRequirements.html) structures do not include `sType` and `pNext` members. This extension wraps them in new structures with these members, so an application can query a chain of memory requirements structures by constructing the chain and letting the implementation fill them in. A new command is added for each `vkGet*MemoryRequrements` command in core Vulkan 1.0.

## [](#_promotion_to_vulkan_1_1)Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_commands)New Commands

- [vkGetBufferMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferMemoryRequirements2KHR.html)
- [vkGetImageMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageMemoryRequirements2KHR.html)
- [vkGetImageSparseMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSparseMemoryRequirements2KHR.html)

## [](#_new_structures)New Structures

- [VkBufferMemoryRequirementsInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferMemoryRequirementsInfo2KHR.html)
- [VkImageMemoryRequirementsInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryRequirementsInfo2KHR.html)
- [VkImageSparseMemoryRequirementsInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSparseMemoryRequirementsInfo2KHR.html)
- [VkMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2KHR.html)
- [VkSparseImageMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryRequirements2KHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_GET_MEMORY_REQUIREMENTS_2_EXTENSION_NAME`
- `VK_KHR_GET_MEMORY_REQUIREMENTS_2_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2_KHR`
  - `VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2_KHR`
  - `VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2_KHR`
  - `VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2_KHR`
  - `VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2_KHR`

## [](#_version_history)Version History

- Revision 1, 2017-03-23 (Faith Ekstrand)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_get_memory_requirements2).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0