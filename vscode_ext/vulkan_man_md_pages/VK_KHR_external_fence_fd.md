# VK_KHR_external_fence_fd(3) Manual Page

## Name

VK_KHR_external_fence_fd - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

116

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_external_fence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_fence.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Jesse Hall <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_fence_fd%5D%20@critsec%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_fence_fd%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>critsec</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-05-08

**IP Status**  
No known IP claims.

**Contributors**  
- Jesse Hall, Google

- James Jones, NVIDIA

- Jeff Juliano, NVIDIA

- Cass Everitt, Oculus

- Contributors to
  [`VK_KHR_external_semaphore_fd`](VK_KHR_external_semaphore_fd.html)

## <a href="#_description" class="anchor"></a>Description

An application using external memory may wish to synchronize access to
that memory using fences. This extension enables an application to
export fence payload to and import fence payload from POSIX file
descriptors.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetFenceFdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetFenceFdKHR.html)

- [vkImportFenceFdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkImportFenceFdKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkFenceGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceGetFdInfoKHR.html)

- [VkImportFenceFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportFenceFdInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_EXTERNAL_FENCE_FD_EXTENSION_NAME`

- `VK_KHR_EXTERNAL_FENCE_FD_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_FENCE_GET_FD_INFO_KHR`

  - `VK_STRUCTURE_TYPE_IMPORT_FENCE_FD_INFO_KHR`

## <a href="#_issues" class="anchor"></a>Issues

This extension borrows concepts, semantics, and language from
[`VK_KHR_external_semaphore_fd`](VK_KHR_external_semaphore_fd.html).
That extension’s issues apply equally to this extension.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-05-08 (Jesse Hall)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_external_fence_fd"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
