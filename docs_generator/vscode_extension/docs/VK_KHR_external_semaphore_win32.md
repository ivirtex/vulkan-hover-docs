# VK\_KHR\_external\_semaphore\_win32(3) Manual Page

## Name

VK\_KHR\_external\_semaphore\_win32 - device extension



## [](#_registered_extension_number)Registered Extension Number

79

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_external\_semaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_semaphore.html)

## [](#_contact)Contact

- James Jones [\[GitHub\]cubanismo](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_semaphore_win32%5D%20%40cubanismo%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_semaphore_win32%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-10-21

**IP Status**

No known IP claims.

**Contributors**

- James Jones, NVIDIA
- Jeff Juliano, NVIDIA
- Carsten Rohde, NVIDIA

## [](#_description)Description

An application using external memory may wish to synchronize access to that memory using semaphores. This extension enables an application to export semaphore payload to and import semaphore payload from Windows handles.

## [](#_new_commands)New Commands

- [vkGetSemaphoreWin32HandleKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSemaphoreWin32HandleKHR.html)
- [vkImportSemaphoreWin32HandleKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkImportSemaphoreWin32HandleKHR.html)

## [](#_new_structures)New Structures

- [VkImportSemaphoreWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportSemaphoreWin32HandleInfoKHR.html)
- [VkSemaphoreGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreGetWin32HandleInfoKHR.html)
- Extending [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreCreateInfo.html):
  
  - [VkExportSemaphoreWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportSemaphoreWin32HandleInfoKHR.html)
- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html):
  
  - [VkD3D12FenceSubmitInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkD3D12FenceSubmitInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_EXTERNAL_SEMAPHORE_WIN32_EXTENSION_NAME`
- `VK_KHR_EXTERNAL_SEMAPHORE_WIN32_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_D3D12_FENCE_SUBMIT_INFO_KHR`
  - `VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_IMPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_SEMAPHORE_GET_WIN32_HANDLE_INFO_KHR`

## [](#_issues)Issues

1\) Do applications need to call `CloseHandle`() on the values returned from [vkGetSemaphoreWin32HandleKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSemaphoreWin32HandleKHR.html) when `handleType` is `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR`?

**RESOLVED**: Yes. A successful get call transfers ownership of the handle to the application. Destroying the semaphore object will not destroy the handle or the handle’s reference to the underlying semaphore resource. Unlike file descriptor opaque handles, win32 opaque handle ownership can not be transferred back to a driver by an import operation.

2\) Should the language regarding KMT/Windows 7 handles be moved to a separate extension so that it can be deprecated over time?

**RESOLVED**: No. Support for them can be deprecated by drivers if they choose, by no longer returning them in the supported handle types of the instance level queries.

3\) Should applications be allowed to specify additional object attributes for shared handles?

**RESOLVED**: Yes. Applications will be allowed to provide similar attributes to those they would to any other handle creation API.

4\) How do applications communicate the desired fence values to use with `D3D12_FENCE`-based Vulkan semaphores?

**RESOLVED**: There are a couple of options. The values for the signaled and reset states could be communicated up front when creating the object and remain static for the life of the Vulkan semaphore, or they could be specified using auxiliary structures when submitting semaphore signal and wait operations, similar to what is done with the keyed mutex extensions. The latter is more flexible and consistent with the keyed mutex usage, but the former is a much simpler API.

Since Vulkan tends to favor flexibility and consistency over simplicity, a new structure specifying D3D12 fence acquire and release values is added to the [vkQueueSubmit](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit.html) function.

## [](#_version_history)Version History

- Revision 1, 2016-10-21 (James Jones)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_external_semaphore_win32)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0