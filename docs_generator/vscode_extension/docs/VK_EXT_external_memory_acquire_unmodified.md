# VK\_EXT\_external\_memory\_acquire\_unmodified(3) Manual Page

## Name

VK\_EXT\_external\_memory\_acquire\_unmodified - device extension



## [](#_registered_extension_number)Registered Extension Number

454

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Lina Versace [\[GitHub\]linyaa-kiwi](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_external_memory_acquire_unmodified%5D%20%40linyaa-kiwi%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_external_memory_acquire_unmodified%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_external\_memory\_acquire\_unmodified](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_external_memory_acquire_unmodified.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-03-09

**Contributors**

- Lina Versace, Google
- Chia-I Wu, Google
- James Jones, NVIDIA
- Yiwei Zhang, Google

## [](#_description)Description

A memory barrier **may** have a performance penalty when acquiring ownership of a subresource range from an external queue family. This extension provides API that **may** reduce the performance penalty if ownership of the subresource range was previously released to the external queue family and if the resource’s memory has remained unmodified between the release and acquire operations.

## [](#_new_structures)New Structures

- Extending [VkBufferMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferMemoryBarrier.html), [VkBufferMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferMemoryBarrier2.html), [VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier.html), [VkImageMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier2.html):
  
  - [VkExternalMemoryAcquireUnmodifiedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryAcquireUnmodifiedEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_EXTERNAL_MEMORY_ACQUIRE_UNMODIFIED_EXTENSION_NAME`
- `VK_EXT_EXTERNAL_MEMORY_ACQUIRE_UNMODIFIED_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_ACQUIRE_UNMODIFIED_EXT`

## [](#_version_history)Version History

- Revision 1, 2023-03-09 (Lina Versace)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_external_memory_acquire_unmodified)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0