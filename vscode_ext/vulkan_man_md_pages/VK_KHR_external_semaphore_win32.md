# VK_KHR_external_semaphore_win32(3) Manual Page

## Name

VK_KHR_external_semaphore_win32 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

79

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_external_semaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_semaphore.html)  

## <a href="#_contact" class="anchor"></a>Contact

- James Jones <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_semaphore_win32%5D%20@cubanismo%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_semaphore_win32%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cubanismo</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-10-21

**IP Status**  
No known IP claims.

**Contributors**  
- James Jones, NVIDIA

- Jeff Juliano, NVIDIA

- Carsten Rohde, NVIDIA

## <a href="#_description" class="anchor"></a>Description

An application using external memory may wish to synchronize access to
that memory using semaphores. This extension enables an application to
export semaphore payload to and import semaphore payload from Windows
handles.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetSemaphoreWin32HandleKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSemaphoreWin32HandleKHR.html)

- [vkImportSemaphoreWin32HandleKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkImportSemaphoreWin32HandleKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkImportSemaphoreWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreWin32HandleInfoKHR.html)

- [VkSemaphoreGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreGetWin32HandleInfoKHR.html)

- Extending [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html):

  - [VkExportSemaphoreWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportSemaphoreWin32HandleInfoKHR.html)

- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html):

  - [VkD3D12FenceSubmitInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkD3D12FenceSubmitInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_EXTERNAL_SEMAPHORE_WIN32_EXTENSION_NAME`

- `VK_KHR_EXTERNAL_SEMAPHORE_WIN32_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_D3D12_FENCE_SUBMIT_INFO_KHR`

  - `VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_IMPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_SEMAPHORE_GET_WIN32_HANDLE_INFO_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) Do applications need to call `CloseHandle`() on the values returned
from [vkGetSemaphoreWin32HandleKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSemaphoreWin32HandleKHR.html)
when `handleType` is
`VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR`?

**RESOLVED**: Yes. A successful get call transfers ownership of the
handle to the application. Destroying the semaphore object will not
destroy the handle or the handle’s reference to the underlying semaphore
resource. Unlike file descriptor opaque handles, win32 opaque handle
ownership can not be transferred back to a driver by an import
operation.

2\) Should the language regarding KMT/Windows 7 handles be moved to a
separate extension so that it can be deprecated over time?

**RESOLVED**: No. Support for them can be deprecated by drivers if they
choose, by no longer returning them in the supported handle types of the
instance level queries.

3\) Should applications be allowed to specify additional object
attributes for shared handles?

**RESOLVED**: Yes. Applications will be allowed to provide similar
attributes to those they would to any other handle creation API.

4\) How do applications communicate the desired fence values to use with
`D3D12_FENCE`-based Vulkan semaphores?

**RESOLVED**: There are a couple of options. The values for the signaled
and reset states could be communicated up front when creating the object
and remain static for the life of the Vulkan semaphore, or they could be
specified using auxiliary structures when submitting semaphore signal
and wait operations, similar to what is done with the keyed mutex
extensions. The latter is more flexible and consistent with the keyed
mutex usage, but the former is a much simpler API.

Since Vulkan tends to favor flexibility and consistency over simplicity,
a new structure specifying D3D12 fence acquire and release values is
added to the [vkQueueSubmit](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueSubmit.html) function.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-10-21 (James Jones)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_external_semaphore_win32"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
