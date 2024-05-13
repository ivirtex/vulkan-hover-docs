# VK_KHR_external_fence_win32(3) Manual Page

## Name

VK_KHR_external_fence_win32 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

115

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_external_fence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_fence.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Jesse Hall <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_fence_win32%5D%20@critsec%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_fence_win32%20extension*"
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
  [`VK_KHR_external_semaphore_win32`](VK_KHR_external_semaphore_win32.html)

## <a href="#_description" class="anchor"></a>Description

An application using external memory may wish to synchronize access to
that memory using fences. This extension enables an application to
export fence payload to and import fence payload from Windows handles.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetFenceWin32HandleKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetFenceWin32HandleKHR.html)

- [vkImportFenceWin32HandleKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkImportFenceWin32HandleKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkFenceGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceGetWin32HandleInfoKHR.html)

- [VkImportFenceWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportFenceWin32HandleInfoKHR.html)

- Extending [VkFenceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceCreateInfo.html):

  - [VkExportFenceWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportFenceWin32HandleInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_EXTERNAL_FENCE_WIN32_EXTENSION_NAME`

- `VK_KHR_EXTERNAL_FENCE_WIN32_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_EXPORT_FENCE_WIN32_HANDLE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_FENCE_GET_WIN32_HANDLE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_IMPORT_FENCE_WIN32_HANDLE_INFO_KHR`

## <a href="#_issues" class="anchor"></a>Issues

This extension borrows concepts, semantics, and language from
[`VK_KHR_external_semaphore_win32`](VK_KHR_external_semaphore_win32.html).
That extension’s issues apply equally to this extension.

1\) Should D3D12 fence handle types be supported, like they are for
semaphores?

**RESOLVED**: No. Doing so would require extending the fence signal and
wait operations to provide values to signal / wait for, like
`VkD3D12FenceSubmitInfoKHR` does. A D3D12 fence can be signaled by
importing it into a [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) instead of a
[VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html), and applications can check status or wait on
the D3D12 fence using non-Vulkan APIs. The convenience of being able to
do these operations on `VkFence` objects does not justify the extra API
complexity.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-05-08 (Jesse Hall)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_external_fence_win32"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
