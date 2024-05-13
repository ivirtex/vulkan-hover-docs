# VK_EXT_external_memory_acquire_unmodified(3) Manual Page

## Name

VK_EXT_external_memory_acquire_unmodified - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

454

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Lina Versace <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_external_memory_acquire_unmodified%5D%20@versalinyaa%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_external_memory_acquire_unmodified%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>versalinyaa</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_external_memory_acquire_unmodified](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_external_memory_acquire_unmodified.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-03-09

**Contributors**  
- Lina Versace, Google

- Chia-I Wu, Google

- James Jones, NVIDIA

- Yiwei Zhang, Google

## <a href="#_description" class="anchor"></a>Description

A memory barrier **may** have a performance penalty when acquiring
ownership of a subresource range from an external queue family. This
extension provides API that **may** reduce the performance penalty if
ownership of the subresource range was previously released to the
external queue family and if the resource’s memory has remained
unmodified between the release and acquire operations.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkBufferMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryBarrier.html),
  [VkBufferMemoryBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryBarrier2.html),
  [VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier.html),
  [VkImageMemoryBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier2.html):

  - [VkExternalMemoryAcquireUnmodifiedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryAcquireUnmodifiedEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_EXTERNAL_MEMORY_ACQUIRE_UNMODIFIED_EXTENSION_NAME`

- `VK_EXT_EXTERNAL_MEMORY_ACQUIRE_UNMODIFIED_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_ACQUIRE_UNMODIFIED_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-03-09 (Lina Versace)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_external_memory_acquire_unmodified"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
